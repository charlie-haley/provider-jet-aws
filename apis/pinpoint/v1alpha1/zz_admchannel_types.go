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

type AdmChannelObservation struct {
}

type AdmChannelParameters struct {

	// +kubebuilder:validation:Required
	ApplicationID *string `json:"applicationId" tf:"application_id,omitempty"`

	// +kubebuilder:validation:Required
	ClientIDSecretRef v1.SecretKeySelector `json:"clientIdSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// AdmChannelSpec defines the desired state of AdmChannel
type AdmChannelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AdmChannelParameters `json:"forProvider"`
}

// AdmChannelStatus defines the observed state of AdmChannel.
type AdmChannelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AdmChannelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AdmChannel is the Schema for the AdmChannels API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type AdmChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AdmChannelSpec   `json:"spec"`
	Status            AdmChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AdmChannelList contains a list of AdmChannels
type AdmChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdmChannel `json:"items"`
}

// Repository type metadata.
var (
	AdmChannel_Kind             = "AdmChannel"
	AdmChannel_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AdmChannel_Kind}.String()
	AdmChannel_KindAPIVersion   = AdmChannel_Kind + "." + CRDGroupVersion.String()
	AdmChannel_GroupVersionKind = CRDGroupVersion.WithKind(AdmChannel_Kind)
)

func init() {
	SchemeBuilder.Register(&AdmChannel{}, &AdmChannelList{})
}
