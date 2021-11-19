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

type AnalyzerObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type AnalyzerParameters struct {

	// +kubebuilder:validation:Required
	AnalyzerName *string `json:"analyzerName" tf:"analyzer_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// AnalyzerSpec defines the desired state of Analyzer
type AnalyzerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AnalyzerParameters `json:"forProvider"`
}

// AnalyzerStatus defines the observed state of Analyzer.
type AnalyzerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AnalyzerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Analyzer is the Schema for the Analyzers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Analyzer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AnalyzerSpec   `json:"spec"`
	Status            AnalyzerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AnalyzerList contains a list of Analyzers
type AnalyzerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Analyzer `json:"items"`
}

// Repository type metadata.
var (
	Analyzer_Kind             = "Analyzer"
	Analyzer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Analyzer_Kind}.String()
	Analyzer_KindAPIVersion   = Analyzer_Kind + "." + CRDGroupVersion.String()
	Analyzer_GroupVersionKind = CRDGroupVersion.WithKind(Analyzer_Kind)
)

func init() {
	SchemeBuilder.Register(&Analyzer{}, &AnalyzerList{})
}
