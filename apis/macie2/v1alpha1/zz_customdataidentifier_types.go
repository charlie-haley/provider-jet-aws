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

type CustomDataIdentifierObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type CustomDataIdentifierParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	IgnoreWords []*string `json:"ignoreWords,omitempty" tf:"ignore_words,omitempty"`

	// +kubebuilder:validation:Optional
	Keywords []*string `json:"keywords,omitempty" tf:"keywords,omitempty"`

	// +kubebuilder:validation:Optional
	MaximumMatchDistance *int64 `json:"maximumMatchDistance,omitempty" tf:"maximum_match_distance,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Regex *string `json:"regex,omitempty" tf:"regex,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// CustomDataIdentifierSpec defines the desired state of CustomDataIdentifier
type CustomDataIdentifierSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CustomDataIdentifierParameters `json:"forProvider"`
}

// CustomDataIdentifierStatus defines the observed state of CustomDataIdentifier.
type CustomDataIdentifierStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CustomDataIdentifierObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CustomDataIdentifier is the Schema for the CustomDataIdentifiers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type CustomDataIdentifier struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CustomDataIdentifierSpec   `json:"spec"`
	Status            CustomDataIdentifierStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CustomDataIdentifierList contains a list of CustomDataIdentifiers
type CustomDataIdentifierList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomDataIdentifier `json:"items"`
}

// Repository type metadata.
var (
	CustomDataIdentifier_Kind             = "CustomDataIdentifier"
	CustomDataIdentifier_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CustomDataIdentifier_Kind}.String()
	CustomDataIdentifier_KindAPIVersion   = CustomDataIdentifier_Kind + "." + CRDGroupVersion.String()
	CustomDataIdentifier_GroupVersionKind = CRDGroupVersion.WithKind(CustomDataIdentifier_Kind)
)

func init() {
	SchemeBuilder.Register(&CustomDataIdentifier{}, &CustomDataIdentifierList{})
}
