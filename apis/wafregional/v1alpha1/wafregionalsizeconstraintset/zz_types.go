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
// +groupName=wafregional.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/wafregional/v1alpha1"
)

type FieldToMatchObservation struct {
}

type FieldToMatchParameters struct {
	Data *string `json:"data,omitempty" tf:"data"`

	Type string `json:"type" tf:"type"`
}

type SizeConstraintsObservation struct {
}

type SizeConstraintsParameters struct {
	ComparisonOperator string `json:"comparisonOperator" tf:"comparison_operator"`

	FieldToMatch []FieldToMatchParameters `json:"fieldToMatch" tf:"field_to_match"`

	Size int64 `json:"size" tf:"size"`

	TextTransformation string `json:"textTransformation" tf:"text_transformation"`
}

type WafregionalSizeConstraintSetObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type WafregionalSizeConstraintSetParameters struct {
	Name string `json:"name" tf:"name"`

	SizeConstraints []SizeConstraintsParameters `json:"sizeConstraints,omitempty" tf:"size_constraints"`
}

// WafregionalSizeConstraintSetSpec defines the desired state of WafregionalSizeConstraintSet
type WafregionalSizeConstraintSetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       WafregionalSizeConstraintSetParameters `json:"forProvider"`
}

// WafregionalSizeConstraintSetStatus defines the observed state of WafregionalSizeConstraintSet.
type WafregionalSizeConstraintSetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          WafregionalSizeConstraintSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WafregionalSizeConstraintSet is the Schema for the WafregionalSizeConstraintSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type WafregionalSizeConstraintSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafregionalSizeConstraintSetSpec   `json:"spec"`
	Status            WafregionalSizeConstraintSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafregionalSizeConstraintSetList contains a list of WafregionalSizeConstraintSets
type WafregionalSizeConstraintSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafregionalSizeConstraintSet `json:"items"`
}

// Repository type metadata.
var (
	WafregionalSizeConstraintSetKind             = "WafregionalSizeConstraintSet"
	WafregionalSizeConstraintSetGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: WafregionalSizeConstraintSetKind}.String()
	WafregionalSizeConstraintSetKindAPIVersion   = WafregionalSizeConstraintSetKind + "." + v1alpha1.GroupVersion.String()
	WafregionalSizeConstraintSetGroupVersionKind = v1alpha1.GroupVersion.WithKind(WafregionalSizeConstraintSetKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&WafregionalSizeConstraintSet{}, &WafregionalSizeConstraintSetList{})
}
