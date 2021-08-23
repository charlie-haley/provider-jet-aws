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
// +groupName=neptune.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/neptune/v1alpha1"
)

type NeptuneClusterParameterGroupObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type NeptuneClusterParameterGroupParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Family string `json:"family" tf:"family"`

	Name *string `json:"name,omitempty" tf:"name"`

	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	Parameter []ParameterParameters `json:"parameter,omitempty" tf:"parameter"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type ParameterObservation struct {
}

type ParameterParameters struct {
	ApplyMethod *string `json:"applyMethod,omitempty" tf:"apply_method"`

	Name string `json:"name" tf:"name"`

	Value string `json:"value" tf:"value"`
}

// NeptuneClusterParameterGroupSpec defines the desired state of NeptuneClusterParameterGroup
type NeptuneClusterParameterGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       NeptuneClusterParameterGroupParameters `json:"forProvider"`
}

// NeptuneClusterParameterGroupStatus defines the observed state of NeptuneClusterParameterGroup.
type NeptuneClusterParameterGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          NeptuneClusterParameterGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NeptuneClusterParameterGroup is the Schema for the NeptuneClusterParameterGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type NeptuneClusterParameterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NeptuneClusterParameterGroupSpec   `json:"spec"`
	Status            NeptuneClusterParameterGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NeptuneClusterParameterGroupList contains a list of NeptuneClusterParameterGroups
type NeptuneClusterParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NeptuneClusterParameterGroup `json:"items"`
}

// Repository type metadata.
var (
	NeptuneClusterParameterGroupKind             = "NeptuneClusterParameterGroup"
	NeptuneClusterParameterGroupGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: NeptuneClusterParameterGroupKind}.String()
	NeptuneClusterParameterGroupKindAPIVersion   = NeptuneClusterParameterGroupKind + "." + v1alpha1.GroupVersion.String()
	NeptuneClusterParameterGroupGroupVersionKind = v1alpha1.GroupVersion.WithKind(NeptuneClusterParameterGroupKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&NeptuneClusterParameterGroup{}, &NeptuneClusterParameterGroupList{})
}
