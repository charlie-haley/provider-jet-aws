/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type InstancePublicPortsObservation struct {
}

type InstancePublicPortsParameters struct {

	// +kubebuilder:validation:Required
	InstanceName *string `json:"instanceName" tf:"instance_name,omitempty"`

	// +kubebuilder:validation:Required
	PortInfo []PortInfoParameters `json:"portInfo" tf:"port_info,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type PortInfoObservation struct {
}

type PortInfoParameters struct {

	// +kubebuilder:validation:Optional
	Cidrs []*string `json:"cidrs,omitempty" tf:"cidrs,omitempty"`

	// +kubebuilder:validation:Required
	FromPort *int64 `json:"fromPort" tf:"from_port,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Required
	ToPort *int64 `json:"toPort" tf:"to_port,omitempty"`
}

// InstancePublicPortsSpec defines the desired state of InstancePublicPorts
type InstancePublicPortsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstancePublicPortsParameters `json:"forProvider"`
}

// InstancePublicPortsStatus defines the observed state of InstancePublicPorts.
type InstancePublicPortsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstancePublicPortsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InstancePublicPorts is the Schema for the InstancePublicPortss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type InstancePublicPorts struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstancePublicPortsSpec   `json:"spec"`
	Status            InstancePublicPortsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstancePublicPortsList contains a list of InstancePublicPortss
type InstancePublicPortsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstancePublicPorts `json:"items"`
}

// Repository type metadata.
var (
	InstancePublicPorts_Kind             = "InstancePublicPorts"
	InstancePublicPorts_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstancePublicPorts_Kind}.String()
	InstancePublicPorts_KindAPIVersion   = InstancePublicPorts_Kind + "." + CRDGroupVersion.String()
	InstancePublicPorts_GroupVersionKind = CRDGroupVersion.WithKind(InstancePublicPorts_Kind)
)

func init() {
	SchemeBuilder.Register(&InstancePublicPorts{}, &InstancePublicPortsList{})
}
