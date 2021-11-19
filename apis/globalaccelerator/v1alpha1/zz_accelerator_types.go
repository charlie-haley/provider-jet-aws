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

type AcceleratorObservation struct {
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	HostedZoneID *string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id,omitempty"`

	IPSets []IPSetsObservation `json:"ipSets,omitempty" tf:"ip_sets,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type AcceleratorParameters struct {

	// +kubebuilder:validation:Optional
	Attributes []AttributesParameters `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	IPAddressType *string `json:"ipAddressType,omitempty" tf:"ip_address_type,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AttributesObservation struct {
}

type AttributesParameters struct {

	// +kubebuilder:validation:Optional
	FlowLogsEnabled *bool `json:"flowLogsEnabled,omitempty" tf:"flow_logs_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	FlowLogsS3Bucket *string `json:"flowLogsS3Bucket,omitempty" tf:"flow_logs_s3_bucket,omitempty"`

	// +kubebuilder:validation:Optional
	FlowLogsS3Prefix *string `json:"flowLogsS3Prefix,omitempty" tf:"flow_logs_s3_prefix,omitempty"`
}

type IPSetsObservation struct {
	IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`

	IPFamily *string `json:"ipFamily,omitempty" tf:"ip_family,omitempty"`
}

type IPSetsParameters struct {
}

// AcceleratorSpec defines the desired state of Accelerator
type AcceleratorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AcceleratorParameters `json:"forProvider"`
}

// AcceleratorStatus defines the observed state of Accelerator.
type AcceleratorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AcceleratorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Accelerator is the Schema for the Accelerators API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Accelerator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AcceleratorSpec   `json:"spec"`
	Status            AcceleratorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AcceleratorList contains a list of Accelerators
type AcceleratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Accelerator `json:"items"`
}

// Repository type metadata.
var (
	Accelerator_Kind             = "Accelerator"
	Accelerator_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Accelerator_Kind}.String()
	Accelerator_KindAPIVersion   = Accelerator_Kind + "." + CRDGroupVersion.String()
	Accelerator_GroupVersionKind = CRDGroupVersion.WithKind(Accelerator_Kind)
)

func init() {
	SchemeBuilder.Register(&Accelerator{}, &AcceleratorList{})
}
