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
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/go-logr/logr"
	stsv1alpha1 "github.com/silicomdk/sts-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// StsConfigReconciler reconciles a StsConfig object
type StsConfigReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

type StsConfigTemplate struct {
	*stsv1alpha1.StsConfig
	NodeName  string
	EnableGPS bool
}

//+kubebuilder:rbac:groups=sts.silicom.com,resources=stsconfigs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=sts.silicom.com,resources=stsconfigs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=sts.silicom.com,resources=stsconfigs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the StsConfig object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.9.2/pkg/reconcile
func (r *StsConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	var objects []client.Object
	reqLogger := r.Log.WithValues("Request.Namespace", req.Namespace, "Request.Name", req.Name)
	reqLogger.Info("Reconciling StsConfig")

	stsConfigList := &stsv1alpha1.StsConfigList{}
	err := r.List(ctx, stsConfigList)
	if err != nil {
		if errors.IsNotFound(err) {
			reqLogger.Error(err, "NOT FOUND: Reconciling StsConfig")
			return ctrl.Result{}, nil
		}
		reqLogger.Error(err, "NOT FOUND2: Reconciling StsConfig")
		return ctrl.Result{}, err
	}

	content, err := ioutil.ReadFile("/assets/sts-deployment.yaml")
	if err != nil {
		reqLogger.Error(err, "ERROR1: Reconciling StsConfig")
		return ctrl.Result{}, err
	}

	t, err := template.New("asset").Option("missingkey=error").Parse(string(content))
	if err != nil {
		reqLogger.Error(err, "ERROR2: Reconciling StsConfig")
		return ctrl.Result{}, err
	}

	for _, stsConfig := range stsConfigList.Items {

		nodeList := &v1.NodeList{}
		err := r.List(ctx, nodeList, client.MatchingLabels(map[string]string{"feature.node.kubernetes.io/pci-0200_8086_1591_1374_02d8.present": "true"}))
		if err != nil {
			if errors.IsNotFound(err) {
				reqLogger.Error(err, "NOT FOUND: Reconciling StsConfig")
				return ctrl.Result{}, nil
			}
			reqLogger.Error(err, "NOT FOUND2: Reconciling StsConfig")
			return ctrl.Result{}, err
		}

		cfgTemplate := &StsConfigTemplate{}
		for _, node := range nodeList.Items {

			var buff bytes.Buffer

			reqLogger.Info(fmt.Sprintf("NODE: %s", node.Name))

			cfgTemplate.EnableGPS = stsConfig.Spec.Mode == "gm"
			cfgTemplate.NodeName = node.Name
			cfgTemplate.StsConfig = &stsConfig

			err = t.Execute(&buff, cfgTemplate)
			if err != nil {
				reqLogger.Error(err, "ERROR3: Reconciling StsConfig")
				return ctrl.Result{}, err
			}

			rx := regexp.MustCompile("\n-{3}")
			objectsDefs := rx.Split(buff.String(), -1)

			for _, objectDef := range objectsDefs {
				obj := unstructured.Unstructured{}
				r := strings.NewReader(objectDef)
				decoder := yaml.NewYAMLOrJSONDecoder(r, 4096)
				err := decoder.Decode(&obj)
				if err != nil {
					reqLogger.Error(err, "ERROR4: Reconciling StsConfig")
					return ctrl.Result{}, err
				}

				objects = append(objects, &obj)
			}

			for _, obj := range objects {
				reqLogger.Info(fmt.Sprintf("Create or update: %v\n", obj))

				gvk := obj.GetObjectKind().GroupVersionKind()
				old := &unstructured.Unstructured{}
				old.SetGroupVersionKind(gvk)
				key := client.ObjectKeyFromObject(obj)
				if err := r.Get(ctx, key, old); err != nil {

					if err := r.Create(ctx, obj); err != nil {
						reqLogger.Error(err, "Create failed", "key", key, "GroupVersionKind", gvk)
						return ctrl.Result{}, err
					}
					reqLogger.Info("Object created")
				} else {
					if !equality.Semantic.DeepDerivative(obj, old) {
						obj.SetResourceVersion(old.GetResourceVersion())
						if err := r.Update(ctx, obj); err != nil {
							reqLogger.Error(err, "Update failed", "key", key, "GroupVersionKind", gvk)
							return ctrl.Result{}, err
						}
						reqLogger.Info("Object updated", "key", key, "GroupVersionKind", gvk)
					} else {
						reqLogger.Info("Object has not changed", "key", key, "GroupVersionKind", gvk)
					}
				}
			}
		}
	}
	return ctrl.Result{}, nil
}

// syncStsConfig synchronizes StsConfig CR
func (r *StsConfigReconciler) syncStsConfig(ctx context.Context, StsConfigList *stsv1alpha1.StsConfigList, nodeList *corev1.NodeList) error {

	fmt.Println("-------------------------------------------------------->syncStsConfig")
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *StsConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	ctrl.NewControllerManagedBy(mgr). // Create the Controller
						For(&stsv1alpha1.StsConfig{}). // StsConfig is the Application API
						Owns(&v1.Pod{}).               // StsConfig owns Pods created by it
						Complete(r)
	return nil
}