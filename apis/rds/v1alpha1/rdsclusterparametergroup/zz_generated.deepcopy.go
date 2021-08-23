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
func (in *ParameterObservation) DeepCopyInto(out *ParameterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterObservation.
func (in *ParameterObservation) DeepCopy() *ParameterObservation {
	if in == nil {
		return nil
	}
	out := new(ParameterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterParameters) DeepCopyInto(out *ParameterParameters) {
	*out = *in
	if in.ApplyMethod != nil {
		in, out := &in.ApplyMethod, &out.ApplyMethod
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterParameters.
func (in *ParameterParameters) DeepCopy() *ParameterParameters {
	if in == nil {
		return nil
	}
	out := new(ParameterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RdsClusterParameterGroup) DeepCopyInto(out *RdsClusterParameterGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RdsClusterParameterGroup.
func (in *RdsClusterParameterGroup) DeepCopy() *RdsClusterParameterGroup {
	if in == nil {
		return nil
	}
	out := new(RdsClusterParameterGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RdsClusterParameterGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RdsClusterParameterGroupList) DeepCopyInto(out *RdsClusterParameterGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RdsClusterParameterGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RdsClusterParameterGroupList.
func (in *RdsClusterParameterGroupList) DeepCopy() *RdsClusterParameterGroupList {
	if in == nil {
		return nil
	}
	out := new(RdsClusterParameterGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RdsClusterParameterGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RdsClusterParameterGroupObservation) DeepCopyInto(out *RdsClusterParameterGroupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RdsClusterParameterGroupObservation.
func (in *RdsClusterParameterGroupObservation) DeepCopy() *RdsClusterParameterGroupObservation {
	if in == nil {
		return nil
	}
	out := new(RdsClusterParameterGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RdsClusterParameterGroupParameters) DeepCopyInto(out *RdsClusterParameterGroupParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamePrefix != nil {
		in, out := &in.NamePrefix, &out.NamePrefix
		*out = new(string)
		**out = **in
	}
	if in.Parameter != nil {
		in, out := &in.Parameter, &out.Parameter
		*out = make([]ParameterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RdsClusterParameterGroupParameters.
func (in *RdsClusterParameterGroupParameters) DeepCopy() *RdsClusterParameterGroupParameters {
	if in == nil {
		return nil
	}
	out := new(RdsClusterParameterGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RdsClusterParameterGroupSpec) DeepCopyInto(out *RdsClusterParameterGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RdsClusterParameterGroupSpec.
func (in *RdsClusterParameterGroupSpec) DeepCopy() *RdsClusterParameterGroupSpec {
	if in == nil {
		return nil
	}
	out := new(RdsClusterParameterGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RdsClusterParameterGroupStatus) DeepCopyInto(out *RdsClusterParameterGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RdsClusterParameterGroupStatus.
func (in *RdsClusterParameterGroupStatus) DeepCopy() *RdsClusterParameterGroupStatus {
	if in == nil {
		return nil
	}
	out := new(RdsClusterParameterGroupStatus)
	in.DeepCopyInto(out)
	return out
}
