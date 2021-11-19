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

type MountTargetObservation struct {
	AvailabilityZoneID *string `json:"availabilityZoneId,omitempty" tf:"availability_zone_id,omitempty"`

	AvailabilityZoneName *string `json:"availabilityZoneName,omitempty" tf:"availability_zone_name,omitempty"`

	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	FileSystemArn *string `json:"fileSystemArn,omitempty" tf:"file_system_arn,omitempty"`

	MountTargetDNSName *string `json:"mountTargetDnsName,omitempty" tf:"mount_target_dns_name,omitempty"`

	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`
}

type MountTargetParameters struct {

	// +kubebuilder:validation:Required
	FileSystemID *string `json:"fileSystemId" tf:"file_system_id,omitempty"`

	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// +kubebuilder:validation:Required
	SubnetID *string `json:"subnetId" tf:"subnet_id,omitempty"`
}

// MountTargetSpec defines the desired state of MountTarget
type MountTargetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MountTargetParameters `json:"forProvider"`
}

// MountTargetStatus defines the observed state of MountTarget.
type MountTargetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MountTargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MountTarget is the Schema for the MountTargets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type MountTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MountTargetSpec   `json:"spec"`
	Status            MountTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MountTargetList contains a list of MountTargets
type MountTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MountTarget `json:"items"`
}

// Repository type metadata.
var (
	MountTarget_Kind             = "MountTarget"
	MountTarget_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MountTarget_Kind}.String()
	MountTarget_KindAPIVersion   = MountTarget_Kind + "." + CRDGroupVersion.String()
	MountTarget_GroupVersionKind = CRDGroupVersion.WithKind(MountTarget_Kind)
)

func init() {
	SchemeBuilder.Register(&MountTarget{}, &MountTargetList{})
}
