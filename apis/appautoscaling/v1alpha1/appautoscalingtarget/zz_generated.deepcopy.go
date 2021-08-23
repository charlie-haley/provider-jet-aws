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
func (in *AppautoscalingTarget) DeepCopyInto(out *AppautoscalingTarget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppautoscalingTarget.
func (in *AppautoscalingTarget) DeepCopy() *AppautoscalingTarget {
	if in == nil {
		return nil
	}
	out := new(AppautoscalingTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppautoscalingTarget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppautoscalingTargetList) DeepCopyInto(out *AppautoscalingTargetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppautoscalingTarget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppautoscalingTargetList.
func (in *AppautoscalingTargetList) DeepCopy() *AppautoscalingTargetList {
	if in == nil {
		return nil
	}
	out := new(AppautoscalingTargetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppautoscalingTargetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppautoscalingTargetObservation) DeepCopyInto(out *AppautoscalingTargetObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppautoscalingTargetObservation.
func (in *AppautoscalingTargetObservation) DeepCopy() *AppautoscalingTargetObservation {
	if in == nil {
		return nil
	}
	out := new(AppautoscalingTargetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppautoscalingTargetParameters) DeepCopyInto(out *AppautoscalingTargetParameters) {
	*out = *in
	if in.RoleArn != nil {
		in, out := &in.RoleArn, &out.RoleArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppautoscalingTargetParameters.
func (in *AppautoscalingTargetParameters) DeepCopy() *AppautoscalingTargetParameters {
	if in == nil {
		return nil
	}
	out := new(AppautoscalingTargetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppautoscalingTargetSpec) DeepCopyInto(out *AppautoscalingTargetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppautoscalingTargetSpec.
func (in *AppautoscalingTargetSpec) DeepCopy() *AppautoscalingTargetSpec {
	if in == nil {
		return nil
	}
	out := new(AppautoscalingTargetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppautoscalingTargetStatus) DeepCopyInto(out *AppautoscalingTargetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppautoscalingTargetStatus.
func (in *AppautoscalingTargetStatus) DeepCopy() *AppautoscalingTargetStatus {
	if in == nil {
		return nil
	}
	out := new(AppautoscalingTargetStatus)
	in.DeepCopyInto(out)
	return out
}
