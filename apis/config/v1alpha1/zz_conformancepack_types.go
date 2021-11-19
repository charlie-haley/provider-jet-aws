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

type ConformancePackObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`
}

type ConformancePackParameters struct {

	// +kubebuilder:validation:Optional
	DeliveryS3Bucket *string `json:"deliveryS3Bucket,omitempty" tf:"delivery_s3_bucket,omitempty"`

	// +kubebuilder:validation:Optional
	DeliveryS3KeyPrefix *string `json:"deliveryS3KeyPrefix,omitempty" tf:"delivery_s3_key_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	InputParameter []InputParameterParameters `json:"inputParameter,omitempty" tf:"input_parameter,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	TemplateBody *string `json:"templateBody,omitempty" tf:"template_body,omitempty"`

	// +kubebuilder:validation:Optional
	TemplateS3URI *string `json:"templateS3Uri,omitempty" tf:"template_s3_uri,omitempty"`
}

type InputParameterObservation struct {
}

type InputParameterParameters struct {

	// +kubebuilder:validation:Required
	ParameterName *string `json:"parameterName" tf:"parameter_name,omitempty"`

	// +kubebuilder:validation:Required
	ParameterValue *string `json:"parameterValue" tf:"parameter_value,omitempty"`
}

// ConformancePackSpec defines the desired state of ConformancePack
type ConformancePackSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConformancePackParameters `json:"forProvider"`
}

// ConformancePackStatus defines the observed state of ConformancePack.
type ConformancePackStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConformancePackObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConformancePack is the Schema for the ConformancePacks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type ConformancePack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConformancePackSpec   `json:"spec"`
	Status            ConformancePackStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConformancePackList contains a list of ConformancePacks
type ConformancePackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConformancePack `json:"items"`
}

// Repository type metadata.
var (
	ConformancePack_Kind             = "ConformancePack"
	ConformancePack_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConformancePack_Kind}.String()
	ConformancePack_KindAPIVersion   = ConformancePack_Kind + "." + CRDGroupVersion.String()
	ConformancePack_GroupVersionKind = CRDGroupVersion.WithKind(ConformancePack_Kind)
)

func init() {
	SchemeBuilder.Register(&ConformancePack{}, &ConformancePackList{})
}
