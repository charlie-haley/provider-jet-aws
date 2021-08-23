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
func (in *LightsailInstancePublicPorts) DeepCopyInto(out *LightsailInstancePublicPorts) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LightsailInstancePublicPorts.
func (in *LightsailInstancePublicPorts) DeepCopy() *LightsailInstancePublicPorts {
	if in == nil {
		return nil
	}
	out := new(LightsailInstancePublicPorts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LightsailInstancePublicPorts) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LightsailInstancePublicPortsList) DeepCopyInto(out *LightsailInstancePublicPortsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LightsailInstancePublicPorts, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LightsailInstancePublicPortsList.
func (in *LightsailInstancePublicPortsList) DeepCopy() *LightsailInstancePublicPortsList {
	if in == nil {
		return nil
	}
	out := new(LightsailInstancePublicPortsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LightsailInstancePublicPortsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LightsailInstancePublicPortsObservation) DeepCopyInto(out *LightsailInstancePublicPortsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LightsailInstancePublicPortsObservation.
func (in *LightsailInstancePublicPortsObservation) DeepCopy() *LightsailInstancePublicPortsObservation {
	if in == nil {
		return nil
	}
	out := new(LightsailInstancePublicPortsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LightsailInstancePublicPortsParameters) DeepCopyInto(out *LightsailInstancePublicPortsParameters) {
	*out = *in
	if in.PortInfo != nil {
		in, out := &in.PortInfo, &out.PortInfo
		*out = make([]PortInfoParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LightsailInstancePublicPortsParameters.
func (in *LightsailInstancePublicPortsParameters) DeepCopy() *LightsailInstancePublicPortsParameters {
	if in == nil {
		return nil
	}
	out := new(LightsailInstancePublicPortsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LightsailInstancePublicPortsSpec) DeepCopyInto(out *LightsailInstancePublicPortsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LightsailInstancePublicPortsSpec.
func (in *LightsailInstancePublicPortsSpec) DeepCopy() *LightsailInstancePublicPortsSpec {
	if in == nil {
		return nil
	}
	out := new(LightsailInstancePublicPortsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LightsailInstancePublicPortsStatus) DeepCopyInto(out *LightsailInstancePublicPortsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LightsailInstancePublicPortsStatus.
func (in *LightsailInstancePublicPortsStatus) DeepCopy() *LightsailInstancePublicPortsStatus {
	if in == nil {
		return nil
	}
	out := new(LightsailInstancePublicPortsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortInfoObservation) DeepCopyInto(out *PortInfoObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortInfoObservation.
func (in *PortInfoObservation) DeepCopy() *PortInfoObservation {
	if in == nil {
		return nil
	}
	out := new(PortInfoObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortInfoParameters) DeepCopyInto(out *PortInfoParameters) {
	*out = *in
	if in.Cidrs != nil {
		in, out := &in.Cidrs, &out.Cidrs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortInfoParameters.
func (in *PortInfoParameters) DeepCopy() *PortInfoParameters {
	if in == nil {
		return nil
	}
	out := new(PortInfoParameters)
	in.DeepCopyInto(out)
	return out
}
