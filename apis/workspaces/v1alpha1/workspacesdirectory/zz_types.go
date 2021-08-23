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
// +groupName=workspaces.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/workspaces/v1alpha1"
)

type SelfServicePermissionsObservation struct {
}

type SelfServicePermissionsParameters struct {
	ChangeComputeType *bool `json:"changeComputeType,omitempty" tf:"change_compute_type"`

	IncreaseVolumeSize *bool `json:"increaseVolumeSize,omitempty" tf:"increase_volume_size"`

	RebuildWorkspace *bool `json:"rebuildWorkspace,omitempty" tf:"rebuild_workspace"`

	RestartWorkspace *bool `json:"restartWorkspace,omitempty" tf:"restart_workspace"`

	SwitchRunningMode *bool `json:"switchRunningMode,omitempty" tf:"switch_running_mode"`
}

type WorkspaceAccessPropertiesObservation struct {
}

type WorkspaceAccessPropertiesParameters struct {
	DeviceTypeAndroid *string `json:"deviceTypeAndroid,omitempty" tf:"device_type_android"`

	DeviceTypeChromeos *string `json:"deviceTypeChromeos,omitempty" tf:"device_type_chromeos"`

	DeviceTypeIos *string `json:"deviceTypeIos,omitempty" tf:"device_type_ios"`

	DeviceTypeLinux *string `json:"deviceTypeLinux,omitempty" tf:"device_type_linux"`

	DeviceTypeOsx *string `json:"deviceTypeOsx,omitempty" tf:"device_type_osx"`

	DeviceTypeWeb *string `json:"deviceTypeWeb,omitempty" tf:"device_type_web"`

	DeviceTypeWindows *string `json:"deviceTypeWindows,omitempty" tf:"device_type_windows"`

	DeviceTypeZeroclient *string `json:"deviceTypeZeroclient,omitempty" tf:"device_type_zeroclient"`
}

type WorkspaceCreationPropertiesObservation struct {
}

type WorkspaceCreationPropertiesParameters struct {
	CustomSecurityGroupId *string `json:"customSecurityGroupId,omitempty" tf:"custom_security_group_id"`

	DefaultOu *string `json:"defaultOu,omitempty" tf:"default_ou"`

	EnableInternetAccess *bool `json:"enableInternetAccess,omitempty" tf:"enable_internet_access"`

	EnableMaintenanceMode *bool `json:"enableMaintenanceMode,omitempty" tf:"enable_maintenance_mode"`

	UserEnabledAsLocalAdministrator *bool `json:"userEnabledAsLocalAdministrator,omitempty" tf:"user_enabled_as_local_administrator"`
}

type WorkspacesDirectoryObservation struct {
	Alias string `json:"alias" tf:"alias"`

	CustomerUserName string `json:"customerUserName" tf:"customer_user_name"`

	DirectoryName string `json:"directoryName" tf:"directory_name"`

	DirectoryType string `json:"directoryType" tf:"directory_type"`

	DnsIpAddresses []string `json:"dnsIpAddresses" tf:"dns_ip_addresses"`

	IamRoleId string `json:"iamRoleId" tf:"iam_role_id"`

	RegistrationCode string `json:"registrationCode" tf:"registration_code"`

	WorkspaceSecurityGroupId string `json:"workspaceSecurityGroupId" tf:"workspace_security_group_id"`
}

type WorkspacesDirectoryParameters struct {
	DirectoryId string `json:"directoryId" tf:"directory_id"`

	IpGroupIds []string `json:"ipGroupIds,omitempty" tf:"ip_group_ids"`

	SelfServicePermissions []SelfServicePermissionsParameters `json:"selfServicePermissions,omitempty" tf:"self_service_permissions"`

	SubnetIds []string `json:"subnetIds,omitempty" tf:"subnet_ids"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	WorkspaceAccessProperties []WorkspaceAccessPropertiesParameters `json:"workspaceAccessProperties,omitempty" tf:"workspace_access_properties"`

	WorkspaceCreationProperties []WorkspaceCreationPropertiesParameters `json:"workspaceCreationProperties,omitempty" tf:"workspace_creation_properties"`
}

// WorkspacesDirectorySpec defines the desired state of WorkspacesDirectory
type WorkspacesDirectorySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       WorkspacesDirectoryParameters `json:"forProvider"`
}

// WorkspacesDirectoryStatus defines the observed state of WorkspacesDirectory.
type WorkspacesDirectoryStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          WorkspacesDirectoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspacesDirectory is the Schema for the WorkspacesDirectorys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type WorkspacesDirectory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkspacesDirectorySpec   `json:"spec"`
	Status            WorkspacesDirectoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspacesDirectoryList contains a list of WorkspacesDirectorys
type WorkspacesDirectoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkspacesDirectory `json:"items"`
}

// Repository type metadata.
var (
	WorkspacesDirectoryKind             = "WorkspacesDirectory"
	WorkspacesDirectoryGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: WorkspacesDirectoryKind}.String()
	WorkspacesDirectoryKindAPIVersion   = WorkspacesDirectoryKind + "." + v1alpha1.GroupVersion.String()
	WorkspacesDirectoryGroupVersionKind = v1alpha1.GroupVersion.WithKind(WorkspacesDirectoryKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&WorkspacesDirectory{}, &WorkspacesDirectoryList{})
}
