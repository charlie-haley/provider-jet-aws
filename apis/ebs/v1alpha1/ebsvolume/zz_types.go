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
// +groupName=ebs.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/ebs/v1alpha1"
)

type EbsVolumeObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type EbsVolumeParameters struct {
	AvailabilityZone string `json:"availabilityZone" tf:"availability_zone"`

	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted"`

	Iops *int64 `json:"iops,omitempty" tf:"iops"`

	KmsKeyId *string `json:"kmsKeyId,omitempty" tf:"kms_key_id"`

	MultiAttachEnabled *bool `json:"multiAttachEnabled,omitempty" tf:"multi_attach_enabled"`

	OutpostArn *string `json:"outpostArn,omitempty" tf:"outpost_arn"`

	Size *int64 `json:"size,omitempty" tf:"size"`

	SnapshotId *string `json:"snapshotId,omitempty" tf:"snapshot_id"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	Throughput *int64 `json:"throughput,omitempty" tf:"throughput"`

	Type *string `json:"type,omitempty" tf:"type"`
}

// EbsVolumeSpec defines the desired state of EbsVolume
type EbsVolumeSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EbsVolumeParameters `json:"forProvider"`
}

// EbsVolumeStatus defines the observed state of EbsVolume.
type EbsVolumeStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EbsVolumeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EbsVolume is the Schema for the EbsVolumes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EbsVolume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EbsVolumeSpec   `json:"spec"`
	Status            EbsVolumeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EbsVolumeList contains a list of EbsVolumes
type EbsVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EbsVolume `json:"items"`
}

// Repository type metadata.
var (
	EbsVolumeKind             = "EbsVolume"
	EbsVolumeGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: EbsVolumeKind}.String()
	EbsVolumeKindAPIVersion   = EbsVolumeKind + "." + v1alpha1.GroupVersion.String()
	EbsVolumeGroupVersionKind = v1alpha1.GroupVersion.WithKind(EbsVolumeKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&EbsVolume{}, &EbsVolumeList{})
}
