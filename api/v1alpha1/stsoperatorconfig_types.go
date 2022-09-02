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

package v1alpha1

import (
	helmerv1beta1 "github.com/openshift-psap/special-resource-operator/pkg/helmer/api/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// StsOperatorConfigSpec defines the desired state of StsOperatorConfig
type StsOperatorConfigSpec struct {

	// +kubebuilder:validation:Required
	// +kubebuilder:default={"tsyncd":"quay.io/silicom/tsyncd:2.1.2.8"}
	Images StsImages `json:"images,omitempty"`

	// +kubebuilder:validation:Required
	// +kubebuilder:default={"build": false}
	Sro SroCfg `json:"sro,omitempty"`
}

type SroCfg struct {
	// +kubebuilder:default:bool=false
	// +kubebuilder:validation:Optional
	Build bool `json:"build,omitempty"`

	// +kubebuilder:default:string="1.8.3"
	// +kubebuilder:validation:Optional
	IceVersion string `json:"iceVersion,omitempty"`

	// +kubebuilder:validation:Optional
	Chart helmerv1beta1.HelmChart `json:"chart,omitempty"`

	// +kubebuilder:default:string="openshift-operators"
	// +kubebuilder:validation:Optional
	Namespace string `json:"namespace,omitempty"`

	// +kubebuilder:default:string="quay.io/silicom/ice-driver-src:1.8.3.2"
	// +kubebuilder:validation:Optional
	SrcImage string `json:"srcImage,omitempty"`

	// +kubebuilder:default:int32=3000
	// +kubebuilder:validation:Optional
	SrcSvcPort int `json:"srcSvcPort,omitempty"`

	// +kubebuilder:default:string="image-registry.openshift-image-registry.svc:5000"
	// +kubebuilder:validation:Optional
	DriverRegistry string `json:"driverRegistry,omitempty"`
}

type StsImages struct {

	// +kubebuilder:default:string="quay.io/silicom/tsyncd:2.1.2.8"
	// +kubebuilder:validation:Required
	Tsyncd string `json:"tsyncd,omitempty"`

	// +kubebuilder:default:string="quay.io/silicom/grpc-tsyncd:2.1.2.8"
	// +kubebuilder:validation:Required
	GrpcTsyncd string `json:"grpcTsyncd,omitempty"`

	// +kubebuilder:default:string="quay.io/silicom/sts-plugin:0.0.9"
	// +kubebuilder:validation:Required
	StsPlugin string `json:"stsPlugin,omitempty"`

	// +kubebuilder:default:string="quay.io/silicom/gpsd:3.23.1"
	// +kubebuilder:validation:Required
	Gpsd string `json:"gpsd,omitempty"`

	// +kubebuilder:default:string="quay.io/silicom/tsync_extts:1.0.0"
	// +kubebuilder:validation:Required
	TsyncExtts string `json:"tsyncExtts,omitempty"`

	// +kubebuilder:default:string="quay.io/silicom/phc2sys:3.1.1"
	// +kubebuilder:validation:Required
	Phc2sys string `json:"phc2sys,omitempty"`
}

// StsOperatorConfigStatus defines the observed state of StsOperatorConfig
type StsOperatorConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// StsOperatorConfig is the Schema for the stsoperatorconfigs API
type StsOperatorConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Spec   StsOperatorConfigSpec   `json:"spec,omitempty"`
	Status StsOperatorConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// StsOperatorConfigList contains a list of StsOperatorConfig
type StsOperatorConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StsOperatorConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StsOperatorConfig{}, &StsOperatorConfigList{})
}
