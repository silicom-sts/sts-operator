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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type TsyncStatus struct {
	Mode   string `json:"mode"`
	Status string `json:"status"`
	Time   string `json:"time"`
}

type GPSStatus struct {
	Time   string `json:"time"`
	Lat    string `json:"lat"`
	Lon    string `json:"lon"`
	Active int    `json:"active"`
	Device string `json:"device"`
	Mode   int    `json:"mode"`
}

// StsNodeSpec defines the desired state of StsNode
type StsNodeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Dummy int `json:"dummy"`
}

type StsNodeInterfaceStatus struct {
	EthName string `json:"ethName"`
	EthPort int    `json:"ethPort"`
	Status  string `json:"status,omitempty"`
	Mode    string `json:"mode,omitempty"`
	PciAddr string `json:"pciAddr"`
}

// StsNodeStatus defines the observed state of StsNode
type StsNodeStatus struct {
	EthInterfaces   []StsNodeInterfaceStatus `json:"ethInterfaces,omitempty"`
	DriverAvailable bool                     `json:"driverAvailable,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// StsNode is the Schema for the stsnodes API
type StsNode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StsNodeSpec   `json:"spec,omitempty"`
	Status StsNodeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// StsNodeList contains a list of StsNode
type StsNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StsNode `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StsNode{}, &StsNodeList{})
}
