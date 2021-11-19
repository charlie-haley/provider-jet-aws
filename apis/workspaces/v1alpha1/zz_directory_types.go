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

type DirectoryObservation struct {
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	CustomerUserName *string `json:"customerUserName,omitempty" tf:"customer_user_name,omitempty"`

	DNSIPAddresses []*string `json:"dnsIpAddresses,omitempty" tf:"dns_ip_addresses,omitempty"`

	DirectoryName *string `json:"directoryName,omitempty" tf:"directory_name,omitempty"`

	DirectoryType *string `json:"directoryType,omitempty" tf:"directory_type,omitempty"`

	IamRoleID *string `json:"iamRoleId,omitempty" tf:"iam_role_id,omitempty"`

	RegistrationCode *string `json:"registrationCode,omitempty" tf:"registration_code,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	WorkspaceSecurityGroupID *string `json:"workspaceSecurityGroupId,omitempty" tf:"workspace_security_group_id,omitempty"`
}

type DirectoryParameters struct {

	// +kubebuilder:validation:Required
	DirectoryID *string `json:"directoryId" tf:"directory_id,omitempty"`

	// +kubebuilder:validation:Optional
	IPGroupIds []*string `json:"ipGroupIds,omitempty" tf:"ip_group_ids,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SelfServicePermissions []SelfServicePermissionsParameters `json:"selfServicePermissions,omitempty" tf:"self_service_permissions,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	WorkspaceAccessProperties []WorkspaceAccessPropertiesParameters `json:"workspaceAccessProperties,omitempty" tf:"workspace_access_properties,omitempty"`

	// +kubebuilder:validation:Optional
	WorkspaceCreationProperties []WorkspaceCreationPropertiesParameters `json:"workspaceCreationProperties,omitempty" tf:"workspace_creation_properties,omitempty"`
}

type SelfServicePermissionsObservation struct {
}

type SelfServicePermissionsParameters struct {

	// +kubebuilder:validation:Optional
	ChangeComputeType *bool `json:"changeComputeType,omitempty" tf:"change_compute_type,omitempty"`

	// +kubebuilder:validation:Optional
	IncreaseVolumeSize *bool `json:"increaseVolumeSize,omitempty" tf:"increase_volume_size,omitempty"`

	// +kubebuilder:validation:Optional
	RebuildWorkspace *bool `json:"rebuildWorkspace,omitempty" tf:"rebuild_workspace,omitempty"`

	// +kubebuilder:validation:Optional
	RestartWorkspace *bool `json:"restartWorkspace,omitempty" tf:"restart_workspace,omitempty"`

	// +kubebuilder:validation:Optional
	SwitchRunningMode *bool `json:"switchRunningMode,omitempty" tf:"switch_running_mode,omitempty"`
}

type WorkspaceAccessPropertiesObservation struct {
}

type WorkspaceAccessPropertiesParameters struct {

	// +kubebuilder:validation:Optional
	DeviceTypeAndroid *string `json:"deviceTypeAndroid,omitempty" tf:"device_type_android,omitempty"`

	// +kubebuilder:validation:Optional
	DeviceTypeChromeos *string `json:"deviceTypeChromeos,omitempty" tf:"device_type_chromeos,omitempty"`

	// +kubebuilder:validation:Optional
	DeviceTypeIos *string `json:"deviceTypeIos,omitempty" tf:"device_type_ios,omitempty"`

	// +kubebuilder:validation:Optional
	DeviceTypeLinux *string `json:"deviceTypeLinux,omitempty" tf:"device_type_linux,omitempty"`

	// +kubebuilder:validation:Optional
	DeviceTypeOsx *string `json:"deviceTypeOsx,omitempty" tf:"device_type_osx,omitempty"`

	// +kubebuilder:validation:Optional
	DeviceTypeWeb *string `json:"deviceTypeWeb,omitempty" tf:"device_type_web,omitempty"`

	// +kubebuilder:validation:Optional
	DeviceTypeWindows *string `json:"deviceTypeWindows,omitempty" tf:"device_type_windows,omitempty"`

	// +kubebuilder:validation:Optional
	DeviceTypeZeroclient *string `json:"deviceTypeZeroclient,omitempty" tf:"device_type_zeroclient,omitempty"`
}

type WorkspaceCreationPropertiesObservation struct {
}

type WorkspaceCreationPropertiesParameters struct {

	// +kubebuilder:validation:Optional
	CustomSecurityGroupID *string `json:"customSecurityGroupId,omitempty" tf:"custom_security_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultOu *string `json:"defaultOu,omitempty" tf:"default_ou,omitempty"`

	// +kubebuilder:validation:Optional
	EnableInternetAccess *bool `json:"enableInternetAccess,omitempty" tf:"enable_internet_access,omitempty"`

	// +kubebuilder:validation:Optional
	EnableMaintenanceMode *bool `json:"enableMaintenanceMode,omitempty" tf:"enable_maintenance_mode,omitempty"`

	// +kubebuilder:validation:Optional
	UserEnabledAsLocalAdministrator *bool `json:"userEnabledAsLocalAdministrator,omitempty" tf:"user_enabled_as_local_administrator,omitempty"`
}

// DirectorySpec defines the desired state of Directory
type DirectorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DirectoryParameters `json:"forProvider"`
}

// DirectoryStatus defines the observed state of Directory.
type DirectoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DirectoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Directory is the Schema for the Directorys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Directory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DirectorySpec   `json:"spec"`
	Status            DirectoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DirectoryList contains a list of Directorys
type DirectoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Directory `json:"items"`
}

// Repository type metadata.
var (
	Directory_Kind             = "Directory"
	Directory_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Directory_Kind}.String()
	Directory_KindAPIVersion   = Directory_Kind + "." + CRDGroupVersion.String()
	Directory_GroupVersionKind = CRDGroupVersion.WithKind(Directory_Kind)
)

func init() {
	SchemeBuilder.Register(&Directory{}, &DirectoryList{})
}
