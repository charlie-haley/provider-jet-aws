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

type DefaultNetworkACLObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type DefaultNetworkACLParameters struct {

	// +kubebuilder:validation:Required
	DefaultNetworkACLID *string `json:"defaultNetworkAclId" tf:"default_network_acl_id,omitempty"`

	// +kubebuilder:validation:Optional
	Egress []EgressParameters `json:"egress,omitempty" tf:"egress,omitempty"`

	// +kubebuilder:validation:Optional
	Ingress []IngressParameters `json:"ingress,omitempty" tf:"ingress,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EgressObservation struct {
}

type EgressParameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// +kubebuilder:validation:Required
	FromPort *int64 `json:"fromPort" tf:"from_port,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	// +kubebuilder:validation:Optional
	IcmpCode *int64 `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`

	// +kubebuilder:validation:Optional
	IcmpType *int64 `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Required
	RuleNo *int64 `json:"ruleNo" tf:"rule_no,omitempty"`

	// +kubebuilder:validation:Required
	ToPort *int64 `json:"toPort" tf:"to_port,omitempty"`
}

type IngressObservation struct {
}

type IngressParameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// +kubebuilder:validation:Required
	FromPort *int64 `json:"fromPort" tf:"from_port,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	// +kubebuilder:validation:Optional
	IcmpCode *int64 `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`

	// +kubebuilder:validation:Optional
	IcmpType *int64 `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Required
	RuleNo *int64 `json:"ruleNo" tf:"rule_no,omitempty"`

	// +kubebuilder:validation:Required
	ToPort *int64 `json:"toPort" tf:"to_port,omitempty"`
}

// DefaultNetworkACLSpec defines the desired state of DefaultNetworkACL
type DefaultNetworkACLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DefaultNetworkACLParameters `json:"forProvider"`
}

// DefaultNetworkACLStatus defines the observed state of DefaultNetworkACL.
type DefaultNetworkACLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DefaultNetworkACLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultNetworkACL is the Schema for the DefaultNetworkACLs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type DefaultNetworkACL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultNetworkACLSpec   `json:"spec"`
	Status            DefaultNetworkACLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultNetworkACLList contains a list of DefaultNetworkACLs
type DefaultNetworkACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultNetworkACL `json:"items"`
}

// Repository type metadata.
var (
	DefaultNetworkACL_Kind             = "DefaultNetworkACL"
	DefaultNetworkACL_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DefaultNetworkACL_Kind}.String()
	DefaultNetworkACL_KindAPIVersion   = DefaultNetworkACL_Kind + "." + CRDGroupVersion.String()
	DefaultNetworkACL_GroupVersionKind = CRDGroupVersion.WithKind(DefaultNetworkACL_Kind)
)

func init() {
	SchemeBuilder.Register(&DefaultNetworkACL{}, &DefaultNetworkACLList{})
}
