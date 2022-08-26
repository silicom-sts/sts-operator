/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"io/ioutil"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/pointer"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/go-logr/logr"
	srov1beta1 "github.com/openshift-psap/special-resource-operator/api/v1beta1"
	helmerv1beta1 "github.com/openshift-psap/special-resource-operator/pkg/helmer/api/v1beta1"
	nfdv1 "github.com/openshift/cluster-nfd-operator/api/v1"
	stsv1alpha1 "github.com/silicomdk/sts-operator/api/v1alpha1"
)

// StsOperatorConfigReconciler reconciles a StsOperatorConfig object
type StsOperatorConfigReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

//+kubebuilder:rbac:groups=sts.silicom.com,resources=stsoperatorconfigs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=sts.silicom.com,resources=stsoperatorconfigs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=sts.silicom.com,resources=stsoperatorconfigs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the StsOperatorConfig object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *StsOperatorConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	reqLogger := r.Log.WithValues("Request.Namespace", req.Namespace, "Request.Name", req.Name)
	reqLogger.Info("Reconciling StsOperatorConfig")

	operatorCfgList := &stsv1alpha1.StsOperatorConfigList{}

	err := r.List(ctx, operatorCfgList, &client.ListOptions{})
	if err != nil {
		reqLogger.Error(err, "Failed to get operator config")
		return ctrl.Result{}, err
	}

	if len(operatorCfgList.Items) == 0 {
		reqLogger.Info("No StsOperatorConfig found")
		return ctrl.Result{}, err
	}

	if len(operatorCfgList.Items) > 1 {
		reqLogger.Info("ERROR: There are 2 StsOperatorConfigs found, please remove 1")
		return ctrl.Result{}, err
	}

	operatorCfg := &stsv1alpha1.StsOperatorConfig{}
	if err := r.Get(context.TODO(), client.ObjectKey{
		Namespace: req.NamespacedName.Namespace,
		Name:      req.NamespacedName.Name,
	}, operatorCfg); err != nil {
		return ctrl.Result{}, err
	}

	err = r.DeployNfd(operatorCfg)
	if err != nil {
		reqLogger.Error(err, "Failed to create NFD CR")
		return ctrl.Result{RequeueAfter: 30}, nil
	}

	err = r.DeploySro(operatorCfg)
	if err != nil {
		reqLogger.Error(err, "Failed to deploy SRO requirements")
		return ctrl.Result{RequeueAfter: 30}, nil
	}

	err = r.DeployPlugin(operatorCfg)
	if err != nil {
		reqLogger.Error(err, "Failed to deploy plugin daemons")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *StsOperatorConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&stsv1alpha1.StsOperatorConfig{}).
		Complete(r)
}

func (r *StsOperatorConfigReconciler) DeploySro(operatorCfg *stsv1alpha1.StsOperatorConfig) error {

	if !operatorCfg.Spec.Sro.Build {
		r.Log.Info("Build of SRO CR is disabled")
		return nil
	}

	if len(operatorCfg.Spec.Sro.Chart.Repository.URL) < 1 {
		operatorCfg.Spec.Sro.Chart.Repository.URL = "http://ice-driver-src"
	}

	if len(operatorCfg.Spec.Sro.Chart.Name) < 1 {
		operatorCfg.Spec.Sro.Chart.Name = "ice-special-resource"
	}

	if len(operatorCfg.Spec.Sro.Chart.Repository.Name) < 1 {
		operatorCfg.Spec.Sro.Chart.Repository.Name = "ice-special-resource"
	}

	svc := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ice-driver-src",
			Namespace: operatorCfg.Namespace,
			Labels: map[string]string{
				"app": "ice-driver-src",
			},
			Annotations: map[string]string{
				"openshift.io/scc": "sts-silicom",
			},
			OwnerReferences: []metav1.OwnerReference{{
				Kind:       operatorCfg.Kind,
				APIVersion: operatorCfg.APIVersion,
				Name:       operatorCfg.Name,
				UID:        operatorCfg.UID,
			}},
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{"app": "ice-driver-src"},
			Ports: []v1.ServicePort{
				{
					Name:       "ice-driver-src",
					Port:       int32(operatorCfg.Spec.Sro.SrcSvcPort),
					TargetPort: intstr.FromInt(operatorCfg.Spec.Sro.SrcSvcPort),
					Protocol:   "TCP",
				},
			},
		},
	}

	ctrl.SetControllerReference(operatorCfg, svc, r.Scheme)

	if err := r.Get(context.TODO(), client.ObjectKey{
		Namespace: svc.Namespace,
		Name:      svc.Name,
	}, svc); err != nil {

		err = r.Create(context.TODO(), svc)
		if err != nil {
			panic(err)
		}
	} else {
		r.Update(context.TODO(), svc)
		if err != nil {
			panic(err)
		}
	}

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ice-driver-src",
			Namespace: operatorCfg.Namespace,
			OwnerReferences: []metav1.OwnerReference{{
				Kind:       operatorCfg.Kind,
				APIVersion: operatorCfg.APIVersion,
				Name:       operatorCfg.Name,
				UID:        operatorCfg.UID,
			}},
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "ice-driver-src",
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "ice-driver-src",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:            "ice-driver-src",
							Image:           operatorCfg.Spec.Sro.SrcImage,
							ImagePullPolicy: "Always",
							Ports: []v1.ContainerPort{
								{
									Name:          "http",
									Protocol:      v1.ProtocolTCP,
									ContainerPort: int32(operatorCfg.Spec.Sro.SrcSvcPort),
								},
							},
						},
					},
				},
			},
		},
	}

	ctrl.SetControllerReference(operatorCfg, deployment, r.Scheme)

	if err := r.Get(context.TODO(), client.ObjectKey{
		Namespace: deployment.Namespace,
		Name:      deployment.Name,
	}, deployment); err != nil {

		err = r.Create(context.TODO(), deployment)
		if err != nil {
			panic(err)
		}
	} else {
		r.Update(context.TODO(), deployment)
		if err != nil {
			panic(err)
		}
	}

	sr := &srov1beta1.SpecialResource{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ice-special-resource",
			Namespace: operatorCfg.Spec.Sro.Namespace,
		},
		Spec: srov1beta1.SpecialResourceSpec{
			Debug:        false,
			Namespace:    operatorCfg.Spec.Sro.Namespace,
			NodeSelector: map[string]string{"feature.node.kubernetes.io/custom-intel.e810_c.devices": "true"},
			Chart: helmerv1beta1.HelmChart{
				Version: operatorCfg.Spec.Sro.Chart.Version,
				Name:    operatorCfg.Spec.Sro.Chart.Name,
				Repository: helmerv1beta1.HelmRepo{
					Name: operatorCfg.Spec.Sro.Chart.Repository.Name,
					URL: fmt.Sprintf("%s.%s.svc:%d",
						operatorCfg.Spec.Sro.Chart.Repository.URL,
						operatorCfg.Namespace,
						operatorCfg.Spec.Sro.SrcSvcPort),
				},
			},
			Set: unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":           "Values",
					"apiVersion":     "sro.openshift.io/v1beta1",
					"driverRegistry": fmt.Sprintf("%s/%s", operatorCfg.Spec.Sro.DriverRegistry, operatorCfg.Spec.Sro.Namespace),
					"buildArgs": []map[string]interface{}{
						{
							"name":  "ICE_VERSION",
							"value": operatorCfg.Spec.Sro.IceVersion,
						},
						{
							"name":  "ICE_SRC",
							"value": operatorCfg.Spec.Sro.SrcImage,
						},
					},
					"runArgs": map[string]interface{}{
						"platform": "openshift-container-platform",
						"buildIce": operatorCfg.Spec.Sro.Build,
					},
				},
			},
		},
	}

	ctrl.SetControllerReference(operatorCfg, sr, r.Scheme)

	if err := r.Get(context.TODO(), client.ObjectKey{
		Namespace: operatorCfg.Spec.Sro.Namespace,
		Name:      sr.Name,
	}, sr); err != nil {

		err = r.Create(context.TODO(), sr)
		if err != nil {
			panic(err)
		}
	} else {
		r.Update(context.TODO(), sr)
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func (r *StsOperatorConfigReconciler) DeployNfd(operatorCfg *stsv1alpha1.StsOperatorConfig) error {
	content, err := ioutil.ReadFile("/assets/nfd-discovery.yaml")
	if err != nil {
		r.Log.Error(err, "Loading nfd-discovery.yaml file")
		return err
	}

	nfdOperand := &nfdv1.NodeFeatureDiscovery{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "nfd-sts-silicom",
			Namespace: operatorCfg.Namespace,
			OwnerReferences: []metav1.OwnerReference{{
				Kind:       operatorCfg.Kind,
				APIVersion: operatorCfg.APIVersion,
				Name:       operatorCfg.Name,
				UID:        operatorCfg.UID,
			}},
		},
		Spec: nfdv1.NodeFeatureDiscoverySpec{
			Operand: nfdv1.OperandSpec{
				Namespace: operatorCfg.Namespace,
			},
			WorkerConfig: &nfdv1.ConfigMap{
				ConfigData: string(content),
			},
		},
	}

	ctrl.SetControllerReference(operatorCfg, nfdOperand, r.Scheme)

	if err := r.Get(context.TODO(), client.ObjectKey{
		Namespace: nfdOperand.Namespace,
		Name:      nfdOperand.Name,
	}, nfdOperand); err != nil {

		err = r.Create(context.TODO(), nfdOperand)
		if err != nil {
			panic(err)
		}
	} else {
		r.Update(context.TODO(), nfdOperand)
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func (r *StsOperatorConfigReconciler) DeployPlugin(operatorCfg *stsv1alpha1.StsOperatorConfig) error {

	daemonset := &appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "sts-plugin",
			Namespace: operatorCfg.Namespace,
			Labels: map[string]string{
				"app": "sts-plugin",
			},
			OwnerReferences: []metav1.OwnerReference{{
				Kind:       operatorCfg.Kind,
				APIVersion: operatorCfg.APIVersion,
				Name:       operatorCfg.Name,
				UID:        operatorCfg.UID,
			}},
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "sts-plugin",
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "sts-plugin",
					},
				},
				Spec: v1.PodSpec{
					NodeSelector: map[string]string{
						"feature.node.kubernetes.io/custom-silicom.sts.devices": "true",
					},
					DNSPolicy:          v1.DNSClusterFirstWithHostNet,
					ServiceAccountName: "sts-plugin",
					HostNetwork:        true,
					Volumes: []v1.Volume{{
						Name: "devfs",
						VolumeSource: v1.VolumeSource{
							HostPath: &v1.HostPathVolumeSource{
								Path: "/dev",
							},
						}},
					},
					Containers: []v1.Container{
						{
							Name:            "sts-plugin",
							Image:           operatorCfg.Spec.Images.StsPlugin,
							ImagePullPolicy: "Always",
							SecurityContext: &v1.SecurityContext{
								Privileged: pointer.Bool(true),
							},
							VolumeMounts: []v1.VolumeMount{{
								Name:      "devfs",
								MountPath: "/dev",
							}},
							Env: []v1.EnvVar{
								{
									Name:  "GPS_SVC_PORT",
									Value: "2947",
								},
								{
									Name: "NODE_NAME",
									ValueFrom: &v1.EnvVarSource{
										FieldRef: &v1.ObjectFieldSelector{
											FieldPath: "spec.nodeName",
										},
									},
								},
								{
									Name: "NAMESPACE",
									ValueFrom: &v1.EnvVarSource{
										FieldRef: &v1.ObjectFieldSelector{
											FieldPath: "metadata.namespace",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	ctrl.SetControllerReference(operatorCfg, daemonset, r.Scheme)

	if err := r.Get(context.TODO(), client.ObjectKey{
		Namespace: daemonset.Namespace,
		Name:      daemonset.Name,
	}, daemonset); err != nil {

		err = r.Create(context.TODO(), daemonset)
		if err != nil {
			panic(err)
		}
	} else {
		r.Update(context.TODO(), daemonset)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
