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
// +groupName=neptune.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/neptune/v1alpha1"
)

type NeptuneClusterSnapshotObservation struct {
	AllocatedStorage int64 `json:"allocatedStorage" tf:"allocated_storage"`

	AvailabilityZones []string `json:"availabilityZones" tf:"availability_zones"`

	DbClusterSnapshotArn string `json:"dbClusterSnapshotArn" tf:"db_cluster_snapshot_arn"`

	Engine string `json:"engine" tf:"engine"`

	EngineVersion string `json:"engineVersion" tf:"engine_version"`

	KmsKeyId string `json:"kmsKeyId" tf:"kms_key_id"`

	LicenseModel string `json:"licenseModel" tf:"license_model"`

	Port int64 `json:"port" tf:"port"`

	SnapshotType string `json:"snapshotType" tf:"snapshot_type"`

	SourceDbClusterSnapshotArn string `json:"sourceDbClusterSnapshotArn" tf:"source_db_cluster_snapshot_arn"`

	Status string `json:"status" tf:"status"`

	StorageEncrypted bool `json:"storageEncrypted" tf:"storage_encrypted"`

	VpcId string `json:"vpcId" tf:"vpc_id"`
}

type NeptuneClusterSnapshotParameters struct {
	DbClusterIdentifier string `json:"dbClusterIdentifier" tf:"db_cluster_identifier"`

	DbClusterSnapshotIdentifier string `json:"dbClusterSnapshotIdentifier" tf:"db_cluster_snapshot_identifier"`
}

// NeptuneClusterSnapshotSpec defines the desired state of NeptuneClusterSnapshot
type NeptuneClusterSnapshotSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       NeptuneClusterSnapshotParameters `json:"forProvider"`
}

// NeptuneClusterSnapshotStatus defines the observed state of NeptuneClusterSnapshot.
type NeptuneClusterSnapshotStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          NeptuneClusterSnapshotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NeptuneClusterSnapshot is the Schema for the NeptuneClusterSnapshots API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type NeptuneClusterSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NeptuneClusterSnapshotSpec   `json:"spec"`
	Status            NeptuneClusterSnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NeptuneClusterSnapshotList contains a list of NeptuneClusterSnapshots
type NeptuneClusterSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NeptuneClusterSnapshot `json:"items"`
}

// Repository type metadata.
var (
	NeptuneClusterSnapshotKind             = "NeptuneClusterSnapshot"
	NeptuneClusterSnapshotGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: NeptuneClusterSnapshotKind}.String()
	NeptuneClusterSnapshotKindAPIVersion   = NeptuneClusterSnapshotKind + "." + v1alpha1.GroupVersion.String()
	NeptuneClusterSnapshotGroupVersionKind = v1alpha1.GroupVersion.WithKind(NeptuneClusterSnapshotKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&NeptuneClusterSnapshot{}, &NeptuneClusterSnapshotList{})
}
