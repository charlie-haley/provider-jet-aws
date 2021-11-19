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

type GlobalClusterMembersObservation struct {
	DBClusterArn *string `json:"dbClusterArn,omitempty" tf:"db_cluster_arn,omitempty"`

	IsWriter *bool `json:"isWriter,omitempty" tf:"is_writer,omitempty"`
}

type GlobalClusterMembersParameters struct {
}

type GlobalClusterObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	GlobalClusterMembers []GlobalClusterMembersObservation `json:"globalClusterMembers,omitempty" tf:"global_cluster_members,omitempty"`

	GlobalClusterResourceID *string `json:"globalClusterResourceId,omitempty" tf:"global_cluster_resource_id,omitempty"`
}

type GlobalClusterParameters struct {

	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// +kubebuilder:validation:Required
	GlobalClusterIdentifier *string `json:"globalClusterIdentifier" tf:"global_cluster_identifier,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SourceDBClusterIdentifier *string `json:"sourceDbClusterIdentifier,omitempty" tf:"source_db_cluster_identifier,omitempty"`

	// +kubebuilder:validation:Optional
	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`
}

// GlobalClusterSpec defines the desired state of GlobalCluster
type GlobalClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GlobalClusterParameters `json:"forProvider"`
}

// GlobalClusterStatus defines the observed state of GlobalCluster.
type GlobalClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GlobalClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalCluster is the Schema for the GlobalClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type GlobalCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlobalClusterSpec   `json:"spec"`
	Status            GlobalClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalClusterList contains a list of GlobalClusters
type GlobalClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalCluster `json:"items"`
}

// Repository type metadata.
var (
	GlobalCluster_Kind             = "GlobalCluster"
	GlobalCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GlobalCluster_Kind}.String()
	GlobalCluster_KindAPIVersion   = GlobalCluster_Kind + "." + CRDGroupVersion.String()
	GlobalCluster_GroupVersionKind = CRDGroupVersion.WithKind(GlobalCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&GlobalCluster{}, &GlobalClusterList{})
}
