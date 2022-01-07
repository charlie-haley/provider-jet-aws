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

type CriterionObservation struct {
}

type CriterionParameters struct {

	// +kubebuilder:validation:Optional
	Eq []*string `json:"eq,omitempty" tf:"eq,omitempty"`

	// +kubebuilder:validation:Optional
	EqExactMatch []*string `json:"eqExactMatch,omitempty" tf:"eq_exact_match,omitempty"`

	// +kubebuilder:validation:Required
	Field *string `json:"field" tf:"field,omitempty"`

	// +kubebuilder:validation:Optional
	Gt *string `json:"gt,omitempty" tf:"gt,omitempty"`

	// +kubebuilder:validation:Optional
	Gte *string `json:"gte,omitempty" tf:"gte,omitempty"`

	// +kubebuilder:validation:Optional
	Lt *string `json:"lt,omitempty" tf:"lt,omitempty"`

	// +kubebuilder:validation:Optional
	Lte *string `json:"lte,omitempty" tf:"lte,omitempty"`

	// +kubebuilder:validation:Optional
	Neq []*string `json:"neq,omitempty" tf:"neq,omitempty"`
}

type FindingCriteriaObservation struct {
}

type FindingCriteriaParameters struct {

	// +kubebuilder:validation:Optional
	Criterion []CriterionParameters `json:"criterion,omitempty" tf:"criterion,omitempty"`
}

type FindingsFilterObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type FindingsFilterParameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	FindingCriteria []FindingCriteriaParameters `json:"findingCriteria" tf:"finding_criteria,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Position *int64 `json:"position,omitempty" tf:"position,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// FindingsFilterSpec defines the desired state of FindingsFilter
type FindingsFilterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FindingsFilterParameters `json:"forProvider"`
}

// FindingsFilterStatus defines the observed state of FindingsFilter.
type FindingsFilterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FindingsFilterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FindingsFilter is the Schema for the FindingsFilters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type FindingsFilter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FindingsFilterSpec   `json:"spec"`
	Status            FindingsFilterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FindingsFilterList contains a list of FindingsFilters
type FindingsFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FindingsFilter `json:"items"`
}

// Repository type metadata.
var (
	FindingsFilter_Kind             = "FindingsFilter"
	FindingsFilter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FindingsFilter_Kind}.String()
	FindingsFilter_KindAPIVersion   = FindingsFilter_Kind + "." + CRDGroupVersion.String()
	FindingsFilter_GroupVersionKind = CRDGroupVersion.WithKind(FindingsFilter_Kind)
)

func init() {
	SchemeBuilder.Register(&FindingsFilter{}, &FindingsFilterList{})
}
