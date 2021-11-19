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

type CopyObservation struct {
	Architecture *string `json:"architecture,omitempty" tf:"architecture,omitempty"`

	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	EnaSupport *bool `json:"enaSupport,omitempty" tf:"ena_support,omitempty"`

	Hypervisor *string `json:"hypervisor,omitempty" tf:"hypervisor,omitempty"`

	ImageLocation *string `json:"imageLocation,omitempty" tf:"image_location,omitempty"`

	ImageOwnerAlias *string `json:"imageOwnerAlias,omitempty" tf:"image_owner_alias,omitempty"`

	ImageType *string `json:"imageType,omitempty" tf:"image_type,omitempty"`

	KernelID *string `json:"kernelId,omitempty" tf:"kernel_id,omitempty"`

	ManageEbsSnapshots *bool `json:"manageEbsSnapshots,omitempty" tf:"manage_ebs_snapshots,omitempty"`

	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	Platform *string `json:"platform,omitempty" tf:"platform,omitempty"`

	PlatformDetails *string `json:"platformDetails,omitempty" tf:"platform_details,omitempty"`

	Public *bool `json:"public,omitempty" tf:"public,omitempty"`

	RamdiskID *string `json:"ramdiskId,omitempty" tf:"ramdisk_id,omitempty"`

	RootDeviceName *string `json:"rootDeviceName,omitempty" tf:"root_device_name,omitempty"`

	RootSnapshotID *string `json:"rootSnapshotId,omitempty" tf:"root_snapshot_id,omitempty"`

	SriovNetSupport *string `json:"sriovNetSupport,omitempty" tf:"sriov_net_support,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	UsageOperation *string `json:"usageOperation,omitempty" tf:"usage_operation,omitempty"`

	VirtualizationType *string `json:"virtualizationType,omitempty" tf:"virtualization_type,omitempty"`
}

type CopyParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DestinationOutpostArn *string `json:"destinationOutpostArn,omitempty" tf:"destination_outpost_arn,omitempty"`

	// +kubebuilder:validation:Optional
	EbsBlockDevice []EbsBlockDeviceParameters `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device,omitempty"`

	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// +kubebuilder:validation:Optional
	EphemeralBlockDevice []EphemeralBlockDeviceParameters `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device,omitempty"`

	// +kubebuilder:validation:Optional
	KmsKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	SourceAmiID *string `json:"sourceAmiId" tf:"source_ami_id,omitempty"`

	// +kubebuilder:validation:Required
	SourceAmiRegion *string `json:"sourceAmiRegion" tf:"source_ami_region,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EbsBlockDeviceObservation struct {
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	Iops *int64 `json:"iops,omitempty" tf:"iops,omitempty"`

	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	Throughput *int64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	VolumeSize *int64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type EbsBlockDeviceParameters struct {
}

type EphemeralBlockDeviceObservation struct {
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	VirtualName *string `json:"virtualName,omitempty" tf:"virtual_name,omitempty"`
}

type EphemeralBlockDeviceParameters struct {
}

// CopySpec defines the desired state of Copy
type CopySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CopyParameters `json:"forProvider"`
}

// CopyStatus defines the observed state of Copy.
type CopyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CopyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Copy is the Schema for the Copys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Copy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CopySpec   `json:"spec"`
	Status            CopyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CopyList contains a list of Copys
type CopyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Copy `json:"items"`
}

// Repository type metadata.
var (
	Copy_Kind             = "Copy"
	Copy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Copy_Kind}.String()
	Copy_KindAPIVersion   = Copy_Kind + "." + CRDGroupVersion.String()
	Copy_GroupVersionKind = CRDGroupVersion.WithKind(Copy_Kind)
)

func init() {
	SchemeBuilder.Register(&Copy{}, &CopyList{})
}
