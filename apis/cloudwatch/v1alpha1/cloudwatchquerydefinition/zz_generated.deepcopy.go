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
func (in *CloudwatchQueryDefinition) DeepCopyInto(out *CloudwatchQueryDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchQueryDefinition.
func (in *CloudwatchQueryDefinition) DeepCopy() *CloudwatchQueryDefinition {
	if in == nil {
		return nil
	}
	out := new(CloudwatchQueryDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudwatchQueryDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchQueryDefinitionList) DeepCopyInto(out *CloudwatchQueryDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudwatchQueryDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchQueryDefinitionList.
func (in *CloudwatchQueryDefinitionList) DeepCopy() *CloudwatchQueryDefinitionList {
	if in == nil {
		return nil
	}
	out := new(CloudwatchQueryDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudwatchQueryDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchQueryDefinitionObservation) DeepCopyInto(out *CloudwatchQueryDefinitionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchQueryDefinitionObservation.
func (in *CloudwatchQueryDefinitionObservation) DeepCopy() *CloudwatchQueryDefinitionObservation {
	if in == nil {
		return nil
	}
	out := new(CloudwatchQueryDefinitionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchQueryDefinitionParameters) DeepCopyInto(out *CloudwatchQueryDefinitionParameters) {
	*out = *in
	if in.LogGroupNames != nil {
		in, out := &in.LogGroupNames, &out.LogGroupNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchQueryDefinitionParameters.
func (in *CloudwatchQueryDefinitionParameters) DeepCopy() *CloudwatchQueryDefinitionParameters {
	if in == nil {
		return nil
	}
	out := new(CloudwatchQueryDefinitionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchQueryDefinitionSpec) DeepCopyInto(out *CloudwatchQueryDefinitionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchQueryDefinitionSpec.
func (in *CloudwatchQueryDefinitionSpec) DeepCopy() *CloudwatchQueryDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(CloudwatchQueryDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchQueryDefinitionStatus) DeepCopyInto(out *CloudwatchQueryDefinitionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchQueryDefinitionStatus.
func (in *CloudwatchQueryDefinitionStatus) DeepCopy() *CloudwatchQueryDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(CloudwatchQueryDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}
