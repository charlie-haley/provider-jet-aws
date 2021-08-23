// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfServicePermissionsObservation) DeepCopyInto(out *SelfServicePermissionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfServicePermissionsObservation.
func (in *SelfServicePermissionsObservation) DeepCopy() *SelfServicePermissionsObservation {
	if in == nil {
		return nil
	}
	out := new(SelfServicePermissionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfServicePermissionsParameters) DeepCopyInto(out *SelfServicePermissionsParameters) {
	*out = *in
	if in.ChangeComputeType != nil {
		in, out := &in.ChangeComputeType, &out.ChangeComputeType
		*out = new(bool)
		**out = **in
	}
	if in.IncreaseVolumeSize != nil {
		in, out := &in.IncreaseVolumeSize, &out.IncreaseVolumeSize
		*out = new(bool)
		**out = **in
	}
	if in.RebuildWorkspace != nil {
		in, out := &in.RebuildWorkspace, &out.RebuildWorkspace
		*out = new(bool)
		**out = **in
	}
	if in.RestartWorkspace != nil {
		in, out := &in.RestartWorkspace, &out.RestartWorkspace
		*out = new(bool)
		**out = **in
	}
	if in.SwitchRunningMode != nil {
		in, out := &in.SwitchRunningMode, &out.SwitchRunningMode
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfServicePermissionsParameters.
func (in *SelfServicePermissionsParameters) DeepCopy() *SelfServicePermissionsParameters {
	if in == nil {
		return nil
	}
	out := new(SelfServicePermissionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceAccessPropertiesObservation) DeepCopyInto(out *WorkspaceAccessPropertiesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceAccessPropertiesObservation.
func (in *WorkspaceAccessPropertiesObservation) DeepCopy() *WorkspaceAccessPropertiesObservation {
	if in == nil {
		return nil
	}
	out := new(WorkspaceAccessPropertiesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceAccessPropertiesParameters) DeepCopyInto(out *WorkspaceAccessPropertiesParameters) {
	*out = *in
	if in.DeviceTypeAndroid != nil {
		in, out := &in.DeviceTypeAndroid, &out.DeviceTypeAndroid
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeChromeos != nil {
		in, out := &in.DeviceTypeChromeos, &out.DeviceTypeChromeos
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeIos != nil {
		in, out := &in.DeviceTypeIos, &out.DeviceTypeIos
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeLinux != nil {
		in, out := &in.DeviceTypeLinux, &out.DeviceTypeLinux
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeOsx != nil {
		in, out := &in.DeviceTypeOsx, &out.DeviceTypeOsx
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeWeb != nil {
		in, out := &in.DeviceTypeWeb, &out.DeviceTypeWeb
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeWindows != nil {
		in, out := &in.DeviceTypeWindows, &out.DeviceTypeWindows
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeZeroclient != nil {
		in, out := &in.DeviceTypeZeroclient, &out.DeviceTypeZeroclient
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceAccessPropertiesParameters.
func (in *WorkspaceAccessPropertiesParameters) DeepCopy() *WorkspaceAccessPropertiesParameters {
	if in == nil {
		return nil
	}
	out := new(WorkspaceAccessPropertiesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceCreationPropertiesObservation) DeepCopyInto(out *WorkspaceCreationPropertiesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceCreationPropertiesObservation.
func (in *WorkspaceCreationPropertiesObservation) DeepCopy() *WorkspaceCreationPropertiesObservation {
	if in == nil {
		return nil
	}
	out := new(WorkspaceCreationPropertiesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceCreationPropertiesParameters) DeepCopyInto(out *WorkspaceCreationPropertiesParameters) {
	*out = *in
	if in.CustomSecurityGroupId != nil {
		in, out := &in.CustomSecurityGroupId, &out.CustomSecurityGroupId
		*out = new(string)
		**out = **in
	}
	if in.DefaultOu != nil {
		in, out := &in.DefaultOu, &out.DefaultOu
		*out = new(string)
		**out = **in
	}
	if in.EnableInternetAccess != nil {
		in, out := &in.EnableInternetAccess, &out.EnableInternetAccess
		*out = new(bool)
		**out = **in
	}
	if in.EnableMaintenanceMode != nil {
		in, out := &in.EnableMaintenanceMode, &out.EnableMaintenanceMode
		*out = new(bool)
		**out = **in
	}
	if in.UserEnabledAsLocalAdministrator != nil {
		in, out := &in.UserEnabledAsLocalAdministrator, &out.UserEnabledAsLocalAdministrator
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceCreationPropertiesParameters.
func (in *WorkspaceCreationPropertiesParameters) DeepCopy() *WorkspaceCreationPropertiesParameters {
	if in == nil {
		return nil
	}
	out := new(WorkspaceCreationPropertiesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspacesDirectory) DeepCopyInto(out *WorkspacesDirectory) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspacesDirectory.
func (in *WorkspacesDirectory) DeepCopy() *WorkspacesDirectory {
	if in == nil {
		return nil
	}
	out := new(WorkspacesDirectory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspacesDirectory) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspacesDirectoryList) DeepCopyInto(out *WorkspacesDirectoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkspacesDirectory, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspacesDirectoryList.
func (in *WorkspacesDirectoryList) DeepCopy() *WorkspacesDirectoryList {
	if in == nil {
		return nil
	}
	out := new(WorkspacesDirectoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspacesDirectoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspacesDirectoryObservation) DeepCopyInto(out *WorkspacesDirectoryObservation) {
	*out = *in
	if in.DnsIpAddresses != nil {
		in, out := &in.DnsIpAddresses, &out.DnsIpAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspacesDirectoryObservation.
func (in *WorkspacesDirectoryObservation) DeepCopy() *WorkspacesDirectoryObservation {
	if in == nil {
		return nil
	}
	out := new(WorkspacesDirectoryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspacesDirectoryParameters) DeepCopyInto(out *WorkspacesDirectoryParameters) {
	*out = *in
	if in.IpGroupIds != nil {
		in, out := &in.IpGroupIds, &out.IpGroupIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SelfServicePermissions != nil {
		in, out := &in.SelfServicePermissions, &out.SelfServicePermissions
		*out = make([]SelfServicePermissionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.WorkspaceAccessProperties != nil {
		in, out := &in.WorkspaceAccessProperties, &out.WorkspaceAccessProperties
		*out = make([]WorkspaceAccessPropertiesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WorkspaceCreationProperties != nil {
		in, out := &in.WorkspaceCreationProperties, &out.WorkspaceCreationProperties
		*out = make([]WorkspaceCreationPropertiesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspacesDirectoryParameters.
func (in *WorkspacesDirectoryParameters) DeepCopy() *WorkspacesDirectoryParameters {
	if in == nil {
		return nil
	}
	out := new(WorkspacesDirectoryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspacesDirectorySpec) DeepCopyInto(out *WorkspacesDirectorySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspacesDirectorySpec.
func (in *WorkspacesDirectorySpec) DeepCopy() *WorkspacesDirectorySpec {
	if in == nil {
		return nil
	}
	out := new(WorkspacesDirectorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspacesDirectoryStatus) DeepCopyInto(out *WorkspacesDirectoryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspacesDirectoryStatus.
func (in *WorkspacesDirectoryStatus) DeepCopy() *WorkspacesDirectoryStatus {
	if in == nil {
		return nil
	}
	out := new(WorkspacesDirectoryStatus)
	in.DeepCopyInto(out)
	return out
}
