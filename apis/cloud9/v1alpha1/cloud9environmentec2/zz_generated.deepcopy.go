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
func (in *Cloud9EnvironmentEc2) DeepCopyInto(out *Cloud9EnvironmentEc2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cloud9EnvironmentEc2.
func (in *Cloud9EnvironmentEc2) DeepCopy() *Cloud9EnvironmentEc2 {
	if in == nil {
		return nil
	}
	out := new(Cloud9EnvironmentEc2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cloud9EnvironmentEc2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cloud9EnvironmentEc2List) DeepCopyInto(out *Cloud9EnvironmentEc2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cloud9EnvironmentEc2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cloud9EnvironmentEc2List.
func (in *Cloud9EnvironmentEc2List) DeepCopy() *Cloud9EnvironmentEc2List {
	if in == nil {
		return nil
	}
	out := new(Cloud9EnvironmentEc2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cloud9EnvironmentEc2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cloud9EnvironmentEc2Observation) DeepCopyInto(out *Cloud9EnvironmentEc2Observation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cloud9EnvironmentEc2Observation.
func (in *Cloud9EnvironmentEc2Observation) DeepCopy() *Cloud9EnvironmentEc2Observation {
	if in == nil {
		return nil
	}
	out := new(Cloud9EnvironmentEc2Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cloud9EnvironmentEc2Parameters) DeepCopyInto(out *Cloud9EnvironmentEc2Parameters) {
	*out = *in
	if in.AutomaticStopTimeMinutes != nil {
		in, out := &in.AutomaticStopTimeMinutes, &out.AutomaticStopTimeMinutes
		*out = new(int64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.OwnerArn != nil {
		in, out := &in.OwnerArn, &out.OwnerArn
		*out = new(string)
		**out = **in
	}
	if in.SubnetId != nil {
		in, out := &in.SubnetId, &out.SubnetId
		*out = new(string)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cloud9EnvironmentEc2Parameters.
func (in *Cloud9EnvironmentEc2Parameters) DeepCopy() *Cloud9EnvironmentEc2Parameters {
	if in == nil {
		return nil
	}
	out := new(Cloud9EnvironmentEc2Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cloud9EnvironmentEc2Spec) DeepCopyInto(out *Cloud9EnvironmentEc2Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cloud9EnvironmentEc2Spec.
func (in *Cloud9EnvironmentEc2Spec) DeepCopy() *Cloud9EnvironmentEc2Spec {
	if in == nil {
		return nil
	}
	out := new(Cloud9EnvironmentEc2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cloud9EnvironmentEc2Status) DeepCopyInto(out *Cloud9EnvironmentEc2Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cloud9EnvironmentEc2Status.
func (in *Cloud9EnvironmentEc2Status) DeepCopy() *Cloud9EnvironmentEc2Status {
	if in == nil {
		return nil
	}
	out := new(Cloud9EnvironmentEc2Status)
	in.DeepCopyInto(out)
	return out
}
