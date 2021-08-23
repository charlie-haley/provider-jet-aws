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
// +groupName=ec2.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1"
)

type Ec2CarrierGatewayObservation struct {
	Arn string `json:"arn" tf:"arn"`

	OwnerId string `json:"ownerId" tf:"owner_id"`
}

type Ec2CarrierGatewayParameters struct {
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	VpcId string `json:"vpcId" tf:"vpc_id"`
}

// Ec2CarrierGatewaySpec defines the desired state of Ec2CarrierGateway
type Ec2CarrierGatewaySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Ec2CarrierGatewayParameters `json:"forProvider"`
}

// Ec2CarrierGatewayStatus defines the observed state of Ec2CarrierGateway.
type Ec2CarrierGatewayStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Ec2CarrierGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2CarrierGateway is the Schema for the Ec2CarrierGateways API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Ec2CarrierGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2CarrierGatewaySpec   `json:"spec"`
	Status            Ec2CarrierGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2CarrierGatewayList contains a list of Ec2CarrierGateways
type Ec2CarrierGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2CarrierGateway `json:"items"`
}

// Repository type metadata.
var (
	Ec2CarrierGatewayKind             = "Ec2CarrierGateway"
	Ec2CarrierGatewayGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: Ec2CarrierGatewayKind}.String()
	Ec2CarrierGatewayKindAPIVersion   = Ec2CarrierGatewayKind + "." + v1alpha1.GroupVersion.String()
	Ec2CarrierGatewayGroupVersionKind = v1alpha1.GroupVersion.WithKind(Ec2CarrierGatewayKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&Ec2CarrierGateway{}, &Ec2CarrierGatewayList{})
}
