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
func (in *PinpointSmsChannel) DeepCopyInto(out *PinpointSmsChannel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinpointSmsChannel.
func (in *PinpointSmsChannel) DeepCopy() *PinpointSmsChannel {
	if in == nil {
		return nil
	}
	out := new(PinpointSmsChannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PinpointSmsChannel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinpointSmsChannelList) DeepCopyInto(out *PinpointSmsChannelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PinpointSmsChannel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinpointSmsChannelList.
func (in *PinpointSmsChannelList) DeepCopy() *PinpointSmsChannelList {
	if in == nil {
		return nil
	}
	out := new(PinpointSmsChannelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PinpointSmsChannelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinpointSmsChannelObservation) DeepCopyInto(out *PinpointSmsChannelObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinpointSmsChannelObservation.
func (in *PinpointSmsChannelObservation) DeepCopy() *PinpointSmsChannelObservation {
	if in == nil {
		return nil
	}
	out := new(PinpointSmsChannelObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinpointSmsChannelParameters) DeepCopyInto(out *PinpointSmsChannelParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.SenderId != nil {
		in, out := &in.SenderId, &out.SenderId
		*out = new(string)
		**out = **in
	}
	if in.ShortCode != nil {
		in, out := &in.ShortCode, &out.ShortCode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinpointSmsChannelParameters.
func (in *PinpointSmsChannelParameters) DeepCopy() *PinpointSmsChannelParameters {
	if in == nil {
		return nil
	}
	out := new(PinpointSmsChannelParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinpointSmsChannelSpec) DeepCopyInto(out *PinpointSmsChannelSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinpointSmsChannelSpec.
func (in *PinpointSmsChannelSpec) DeepCopy() *PinpointSmsChannelSpec {
	if in == nil {
		return nil
	}
	out := new(PinpointSmsChannelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinpointSmsChannelStatus) DeepCopyInto(out *PinpointSmsChannelStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinpointSmsChannelStatus.
func (in *PinpointSmsChannelStatus) DeepCopy() *PinpointSmsChannelStatus {
	if in == nil {
		return nil
	}
	out := new(PinpointSmsChannelStatus)
	in.DeepCopyInto(out)
	return out
}
