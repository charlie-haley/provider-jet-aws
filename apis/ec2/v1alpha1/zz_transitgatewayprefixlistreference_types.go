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

type TransitGatewayPrefixListReferenceObservation struct {
	PrefixListOwnerID *string `json:"prefixListOwnerId,omitempty" tf:"prefix_list_owner_id,omitempty"`
}

type TransitGatewayPrefixListReferenceParameters struct {

	// +kubebuilder:validation:Optional
	Blackhole *bool `json:"blackhole,omitempty" tf:"blackhole,omitempty"`

	// +kubebuilder:validation:Required
	PrefixListID *string `json:"prefixListId" tf:"prefix_list_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentID *string `json:"transitGatewayAttachmentId,omitempty" tf:"transit_gateway_attachment_id,omitempty"`

	// +kubebuilder:validation:Required
	TransitGatewayRouteTableID *string `json:"transitGatewayRouteTableId" tf:"transit_gateway_route_table_id,omitempty"`
}

// TransitGatewayPrefixListReferenceSpec defines the desired state of TransitGatewayPrefixListReference
type TransitGatewayPrefixListReferenceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayPrefixListReferenceParameters `json:"forProvider"`
}

// TransitGatewayPrefixListReferenceStatus defines the observed state of TransitGatewayPrefixListReference.
type TransitGatewayPrefixListReferenceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayPrefixListReferenceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayPrefixListReference is the Schema for the TransitGatewayPrefixListReferences API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type TransitGatewayPrefixListReference struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewayPrefixListReferenceSpec   `json:"spec"`
	Status            TransitGatewayPrefixListReferenceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayPrefixListReferenceList contains a list of TransitGatewayPrefixListReferences
type TransitGatewayPrefixListReferenceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayPrefixListReference `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayPrefixListReference_Kind             = "TransitGatewayPrefixListReference"
	TransitGatewayPrefixListReference_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayPrefixListReference_Kind}.String()
	TransitGatewayPrefixListReference_KindAPIVersion   = TransitGatewayPrefixListReference_Kind + "." + CRDGroupVersion.String()
	TransitGatewayPrefixListReference_GroupVersionKind = CRDGroupVersion.WithKind(TransitGatewayPrefixListReference_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayPrefixListReference{}, &TransitGatewayPrefixListReferenceList{})
}
