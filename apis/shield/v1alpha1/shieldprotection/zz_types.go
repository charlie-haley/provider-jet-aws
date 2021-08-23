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
// +groupName=shield.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/shield/v1alpha1"
)

type ShieldProtectionObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type ShieldProtectionParameters struct {
	Name string `json:"name" tf:"name"`

	ResourceArn string `json:"resourceArn" tf:"resource_arn"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// ShieldProtectionSpec defines the desired state of ShieldProtection
type ShieldProtectionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ShieldProtectionParameters `json:"forProvider"`
}

// ShieldProtectionStatus defines the observed state of ShieldProtection.
type ShieldProtectionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ShieldProtectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ShieldProtection is the Schema for the ShieldProtections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ShieldProtection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ShieldProtectionSpec   `json:"spec"`
	Status            ShieldProtectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ShieldProtectionList contains a list of ShieldProtections
type ShieldProtectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ShieldProtection `json:"items"`
}

// Repository type metadata.
var (
	ShieldProtectionKind             = "ShieldProtection"
	ShieldProtectionGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: ShieldProtectionKind}.String()
	ShieldProtectionKindAPIVersion   = ShieldProtectionKind + "." + v1alpha1.GroupVersion.String()
	ShieldProtectionGroupVersionKind = v1alpha1.GroupVersion.WithKind(ShieldProtectionKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&ShieldProtection{}, &ShieldProtectionList{})
}
