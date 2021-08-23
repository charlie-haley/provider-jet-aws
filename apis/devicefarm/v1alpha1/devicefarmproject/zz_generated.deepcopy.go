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
func (in *DevicefarmProject) DeepCopyInto(out *DevicefarmProject) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicefarmProject.
func (in *DevicefarmProject) DeepCopy() *DevicefarmProject {
	if in == nil {
		return nil
	}
	out := new(DevicefarmProject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevicefarmProject) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicefarmProjectList) DeepCopyInto(out *DevicefarmProjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DevicefarmProject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicefarmProjectList.
func (in *DevicefarmProjectList) DeepCopy() *DevicefarmProjectList {
	if in == nil {
		return nil
	}
	out := new(DevicefarmProjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevicefarmProjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicefarmProjectObservation) DeepCopyInto(out *DevicefarmProjectObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicefarmProjectObservation.
func (in *DevicefarmProjectObservation) DeepCopy() *DevicefarmProjectObservation {
	if in == nil {
		return nil
	}
	out := new(DevicefarmProjectObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicefarmProjectParameters) DeepCopyInto(out *DevicefarmProjectParameters) {
	*out = *in
	if in.DefaultJobTimeoutMinutes != nil {
		in, out := &in.DefaultJobTimeoutMinutes, &out.DefaultJobTimeoutMinutes
		*out = new(int64)
		**out = **in
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicefarmProjectParameters.
func (in *DevicefarmProjectParameters) DeepCopy() *DevicefarmProjectParameters {
	if in == nil {
		return nil
	}
	out := new(DevicefarmProjectParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicefarmProjectSpec) DeepCopyInto(out *DevicefarmProjectSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicefarmProjectSpec.
func (in *DevicefarmProjectSpec) DeepCopy() *DevicefarmProjectSpec {
	if in == nil {
		return nil
	}
	out := new(DevicefarmProjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicefarmProjectStatus) DeepCopyInto(out *DevicefarmProjectStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicefarmProjectStatus.
func (in *DevicefarmProjectStatus) DeepCopy() *DevicefarmProjectStatus {
	if in == nil {
		return nil
	}
	out := new(DevicefarmProjectStatus)
	in.DeepCopyInto(out)
	return out
}
