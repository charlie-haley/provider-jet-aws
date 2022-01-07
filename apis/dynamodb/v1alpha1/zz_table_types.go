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

type AttributeObservation struct {
}

type AttributeParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type GlobalSecondaryIndexObservation struct {
}

type GlobalSecondaryIndexParameters struct {

	// +kubebuilder:validation:Required
	HashKey *string `json:"hashKey" tf:"hash_key,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NonKeyAttributes []*string `json:"nonKeyAttributes,omitempty" tf:"non_key_attributes,omitempty"`

	// +kubebuilder:validation:Required
	ProjectionType *string `json:"projectionType" tf:"projection_type,omitempty"`

	// +kubebuilder:validation:Optional
	RangeKey *string `json:"rangeKey,omitempty" tf:"range_key,omitempty"`

	// +kubebuilder:validation:Optional
	ReadCapacity *int64 `json:"readCapacity,omitempty" tf:"read_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	WriteCapacity *int64 `json:"writeCapacity,omitempty" tf:"write_capacity,omitempty"`
}

type LocalSecondaryIndexObservation struct {
}

type LocalSecondaryIndexParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NonKeyAttributes []*string `json:"nonKeyAttributes,omitempty" tf:"non_key_attributes,omitempty"`

	// +kubebuilder:validation:Required
	ProjectionType *string `json:"projectionType" tf:"projection_type,omitempty"`

	// +kubebuilder:validation:Required
	RangeKey *string `json:"rangeKey" tf:"range_key,omitempty"`
}

type PointInTimeRecoveryObservation struct {
}

type PointInTimeRecoveryParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type ServerSideEncryptionObservation struct {
}

type ServerSideEncryptionParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
}

type TTLObservation struct {
}

type TTLParameters struct {

	// +kubebuilder:validation:Required
	AttributeName *string `json:"attributeName" tf:"attribute_name,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
}

type TableObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	StreamArn *string `json:"streamArn,omitempty" tf:"stream_arn,omitempty"`

	StreamLabel *string `json:"streamLabel,omitempty" tf:"stream_label,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type TableParameters struct {

	// +kubebuilder:validation:Required
	Attribute []AttributeParameters `json:"attribute" tf:"attribute,omitempty"`

	// +kubebuilder:validation:Optional
	BillingMode *string `json:"billingMode,omitempty" tf:"billing_mode,omitempty"`

	// +kubebuilder:validation:Optional
	GlobalSecondaryIndex []GlobalSecondaryIndexParameters `json:"globalSecondaryIndex,omitempty" tf:"global_secondary_index,omitempty"`

	// +kubebuilder:validation:Required
	HashKey *string `json:"hashKey" tf:"hash_key,omitempty"`

	// +kubebuilder:validation:Optional
	LocalSecondaryIndex []LocalSecondaryIndexParameters `json:"localSecondaryIndex,omitempty" tf:"local_secondary_index,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PointInTimeRecovery []PointInTimeRecoveryParameters `json:"pointInTimeRecovery,omitempty" tf:"point_in_time_recovery,omitempty"`

	// +kubebuilder:validation:Optional
	RangeKey *string `json:"rangeKey,omitempty" tf:"range_key,omitempty"`

	// +kubebuilder:validation:Optional
	ReadCapacity *int64 `json:"readCapacity,omitempty" tf:"read_capacity,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Replica []TableReplicaParameters `json:"replica,omitempty" tf:"replica,omitempty"`

	// +kubebuilder:validation:Optional
	ServerSideEncryption []ServerSideEncryptionParameters `json:"serverSideEncryption,omitempty" tf:"server_side_encryption,omitempty"`

	// +kubebuilder:validation:Optional
	StreamEnabled *bool `json:"streamEnabled,omitempty" tf:"stream_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	StreamViewType *string `json:"streamViewType,omitempty" tf:"stream_view_type,omitempty"`

	// +kubebuilder:validation:Optional
	TTL []TTLParameters `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	WriteCapacity *int64 `json:"writeCapacity,omitempty" tf:"write_capacity,omitempty"`
}

type TableReplicaObservation struct {
}

type TableReplicaParameters struct {

	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// +kubebuilder:validation:Required
	RegionName *string `json:"regionName" tf:"region_name,omitempty"`
}

// TableSpec defines the desired state of Table
type TableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TableParameters `json:"forProvider"`
}

// TableStatus defines the observed state of Table.
type TableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Table is the Schema for the Tables API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Table struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TableSpec   `json:"spec"`
	Status            TableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableList contains a list of Tables
type TableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Table `json:"items"`
}

// Repository type metadata.
var (
	Table_Kind             = "Table"
	Table_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Table_Kind}.String()
	Table_KindAPIVersion   = Table_Kind + "." + CRDGroupVersion.String()
	Table_GroupVersionKind = CRDGroupVersion.WithKind(Table_Kind)
)

func init() {
	SchemeBuilder.Register(&Table{}, &TableList{})
}
