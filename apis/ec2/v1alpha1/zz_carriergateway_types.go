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

type CarrierGatewayObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type CarrierGatewayParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcidRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcidSelector,omitempty" tf:"-"`
}

// CarrierGatewaySpec defines the desired state of CarrierGateway
type CarrierGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CarrierGatewayParameters `json:"forProvider"`
}

// CarrierGatewayStatus defines the observed state of CarrierGateway.
type CarrierGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CarrierGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CarrierGateway is the Schema for the CarrierGateways API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type CarrierGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CarrierGatewaySpec   `json:"spec"`
	Status            CarrierGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CarrierGatewayList contains a list of CarrierGateways
type CarrierGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CarrierGateway `json:"items"`
}

// Repository type metadata.
var (
	CarrierGateway_Kind             = "CarrierGateway"
	CarrierGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CarrierGateway_Kind}.String()
	CarrierGateway_KindAPIVersion   = CarrierGateway_Kind + "." + CRDGroupVersion.String()
	CarrierGateway_GroupVersionKind = CRDGroupVersion.WithKind(CarrierGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&CarrierGateway{}, &CarrierGatewayList{})
}
