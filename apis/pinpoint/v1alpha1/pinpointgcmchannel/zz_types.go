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

// +kubebuilder:object:generate=true
// +groupName=pinpoint.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/pinpoint/v1alpha1"
)

type PinpointGcmChannelObservation struct {
}

type PinpointGcmChannelParameters struct {
	ApiKey string `json:"apiKey" tf:"api_key"`

	ApplicationId string `json:"applicationId" tf:"application_id"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
}

// PinpointGcmChannelSpec defines the desired state of PinpointGcmChannel
type PinpointGcmChannelSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       PinpointGcmChannelParameters `json:"forProvider"`
}

// PinpointGcmChannelStatus defines the observed state of PinpointGcmChannel.
type PinpointGcmChannelStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          PinpointGcmChannelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PinpointGcmChannel is the Schema for the PinpointGcmChannels API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type PinpointGcmChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PinpointGcmChannelSpec   `json:"spec"`
	Status            PinpointGcmChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PinpointGcmChannelList contains a list of PinpointGcmChannels
type PinpointGcmChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PinpointGcmChannel `json:"items"`
}

// Repository type metadata.
var (
	PinpointGcmChannelKind             = "PinpointGcmChannel"
	PinpointGcmChannelGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: PinpointGcmChannelKind}.String()
	PinpointGcmChannelKindAPIVersion   = PinpointGcmChannelKind + "." + v1alpha1.GroupVersion.String()
	PinpointGcmChannelGroupVersionKind = v1alpha1.GroupVersion.WithKind(PinpointGcmChannelKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&PinpointGcmChannel{}, &PinpointGcmChannelList{})
}
