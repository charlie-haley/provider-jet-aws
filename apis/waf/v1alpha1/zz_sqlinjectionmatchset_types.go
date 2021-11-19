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

type SQLInjectionMatchTuplesFieldToMatchObservation struct {
}

type SQLInjectionMatchTuplesFieldToMatchParameters struct {

	// +kubebuilder:validation:Optional
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type SQLInjectionMatchTuplesObservation struct {
}

type SQLInjectionMatchTuplesParameters struct {

	// +kubebuilder:validation:Required
	FieldToMatch []SQLInjectionMatchTuplesFieldToMatchParameters `json:"fieldToMatch" tf:"field_to_match,omitempty"`

	// +kubebuilder:validation:Required
	TextTransformation *string `json:"textTransformation" tf:"text_transformation,omitempty"`
}

type SqlInjectionMatchSetObservation struct {
}

type SqlInjectionMatchSetParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SQLInjectionMatchTuples []SQLInjectionMatchTuplesParameters `json:"sqlInjectionMatchTuples,omitempty" tf:"sql_injection_match_tuples,omitempty"`
}

// SqlInjectionMatchSetSpec defines the desired state of SqlInjectionMatchSet
type SqlInjectionMatchSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SqlInjectionMatchSetParameters `json:"forProvider"`
}

// SqlInjectionMatchSetStatus defines the observed state of SqlInjectionMatchSet.
type SqlInjectionMatchSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SqlInjectionMatchSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SqlInjectionMatchSet is the Schema for the SqlInjectionMatchSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type SqlInjectionMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlInjectionMatchSetSpec   `json:"spec"`
	Status            SqlInjectionMatchSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SqlInjectionMatchSetList contains a list of SqlInjectionMatchSets
type SqlInjectionMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlInjectionMatchSet `json:"items"`
}

// Repository type metadata.
var (
	SqlInjectionMatchSet_Kind             = "SqlInjectionMatchSet"
	SqlInjectionMatchSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SqlInjectionMatchSet_Kind}.String()
	SqlInjectionMatchSet_KindAPIVersion   = SqlInjectionMatchSet_Kind + "." + CRDGroupVersion.String()
	SqlInjectionMatchSet_GroupVersionKind = CRDGroupVersion.WithKind(SqlInjectionMatchSet_Kind)
)

func init() {
	SchemeBuilder.Register(&SqlInjectionMatchSet{}, &SqlInjectionMatchSetList{})
}
