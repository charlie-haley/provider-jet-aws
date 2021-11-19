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

type GatewayDocumentationVersionObservation struct {
}

type GatewayDocumentationVersionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	RestAPIID *string `json:"restApiId" tf:"rest_api_id,omitempty"`

	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

// GatewayDocumentationVersionSpec defines the desired state of GatewayDocumentationVersion
type GatewayDocumentationVersionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayDocumentationVersionParameters `json:"forProvider"`
}

// GatewayDocumentationVersionStatus defines the observed state of GatewayDocumentationVersion.
type GatewayDocumentationVersionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayDocumentationVersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayDocumentationVersion is the Schema for the GatewayDocumentationVersions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type GatewayDocumentationVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewayDocumentationVersionSpec   `json:"spec"`
	Status            GatewayDocumentationVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayDocumentationVersionList contains a list of GatewayDocumentationVersions
type GatewayDocumentationVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GatewayDocumentationVersion `json:"items"`
}

// Repository type metadata.
var (
	GatewayDocumentationVersion_Kind             = "GatewayDocumentationVersion"
	GatewayDocumentationVersion_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GatewayDocumentationVersion_Kind}.String()
	GatewayDocumentationVersion_KindAPIVersion   = GatewayDocumentationVersion_Kind + "." + CRDGroupVersion.String()
	GatewayDocumentationVersion_GroupVersionKind = CRDGroupVersion.WithKind(GatewayDocumentationVersion_Kind)
)

func init() {
	SchemeBuilder.Register(&GatewayDocumentationVersion{}, &GatewayDocumentationVersionList{})
}
