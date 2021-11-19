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

type RegexMatchSetObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`
}

type RegexMatchSetParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RegexMatchTuple []RegexMatchTupleParameters `json:"regexMatchTuple,omitempty" tf:"regex_match_tuple,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type RegexMatchTupleFieldToMatchObservation struct {
}

type RegexMatchTupleFieldToMatchParameters struct {

	// +kubebuilder:validation:Optional
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type RegexMatchTupleObservation struct {
}

type RegexMatchTupleParameters struct {

	// +kubebuilder:validation:Required
	FieldToMatch []RegexMatchTupleFieldToMatchParameters `json:"fieldToMatch" tf:"field_to_match,omitempty"`

	// +kubebuilder:validation:Required
	RegexPatternSetID *string `json:"regexPatternSetId" tf:"regex_pattern_set_id,omitempty"`

	// +kubebuilder:validation:Required
	TextTransformation *string `json:"textTransformation" tf:"text_transformation,omitempty"`
}

// RegexMatchSetSpec defines the desired state of RegexMatchSet
type RegexMatchSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegexMatchSetParameters `json:"forProvider"`
}

// RegexMatchSetStatus defines the observed state of RegexMatchSet.
type RegexMatchSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegexMatchSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RegexMatchSet is the Schema for the RegexMatchSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type RegexMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegexMatchSetSpec   `json:"spec"`
	Status            RegexMatchSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegexMatchSetList contains a list of RegexMatchSets
type RegexMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegexMatchSet `json:"items"`
}

// Repository type metadata.
var (
	RegexMatchSet_Kind             = "RegexMatchSet"
	RegexMatchSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegexMatchSet_Kind}.String()
	RegexMatchSet_KindAPIVersion   = RegexMatchSet_Kind + "." + CRDGroupVersion.String()
	RegexMatchSet_GroupVersionKind = CRDGroupVersion.WithKind(RegexMatchSet_Kind)
)

func init() {
	SchemeBuilder.Register(&RegexMatchSet{}, &RegexMatchSetList{})
}
