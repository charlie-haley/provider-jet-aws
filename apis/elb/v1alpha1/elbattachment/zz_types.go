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
// +groupName=elb.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/elb/v1alpha1"
)

type ElbAttachmentObservation struct {
}

type ElbAttachmentParameters struct {
	Elb string `json:"elb" tf:"elb"`

	Instance string `json:"instance" tf:"instance"`
}

// ElbAttachmentSpec defines the desired state of ElbAttachment
type ElbAttachmentSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ElbAttachmentParameters `json:"forProvider"`
}

// ElbAttachmentStatus defines the observed state of ElbAttachment.
type ElbAttachmentStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ElbAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ElbAttachment is the Schema for the ElbAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ElbAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElbAttachmentSpec   `json:"spec"`
	Status            ElbAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElbAttachmentList contains a list of ElbAttachments
type ElbAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElbAttachment `json:"items"`
}

// Repository type metadata.
var (
	ElbAttachmentKind             = "ElbAttachment"
	ElbAttachmentGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: ElbAttachmentKind}.String()
	ElbAttachmentKindAPIVersion   = ElbAttachmentKind + "." + v1alpha1.GroupVersion.String()
	ElbAttachmentGroupVersionKind = v1alpha1.GroupVersion.WithKind(ElbAttachmentKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&ElbAttachment{}, &ElbAttachmentList{})
}
