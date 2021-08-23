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
func (in *RedshiftEventSubscription) DeepCopyInto(out *RedshiftEventSubscription) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedshiftEventSubscription.
func (in *RedshiftEventSubscription) DeepCopy() *RedshiftEventSubscription {
	if in == nil {
		return nil
	}
	out := new(RedshiftEventSubscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedshiftEventSubscription) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedshiftEventSubscriptionList) DeepCopyInto(out *RedshiftEventSubscriptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedshiftEventSubscription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedshiftEventSubscriptionList.
func (in *RedshiftEventSubscriptionList) DeepCopy() *RedshiftEventSubscriptionList {
	if in == nil {
		return nil
	}
	out := new(RedshiftEventSubscriptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedshiftEventSubscriptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedshiftEventSubscriptionObservation) DeepCopyInto(out *RedshiftEventSubscriptionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedshiftEventSubscriptionObservation.
func (in *RedshiftEventSubscriptionObservation) DeepCopy() *RedshiftEventSubscriptionObservation {
	if in == nil {
		return nil
	}
	out := new(RedshiftEventSubscriptionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedshiftEventSubscriptionParameters) DeepCopyInto(out *RedshiftEventSubscriptionParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EventCategories != nil {
		in, out := &in.EventCategories, &out.EventCategories
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Severity != nil {
		in, out := &in.Severity, &out.Severity
		*out = new(string)
		**out = **in
	}
	if in.SourceIds != nil {
		in, out := &in.SourceIds, &out.SourceIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SourceType != nil {
		in, out := &in.SourceType, &out.SourceType
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedshiftEventSubscriptionParameters.
func (in *RedshiftEventSubscriptionParameters) DeepCopy() *RedshiftEventSubscriptionParameters {
	if in == nil {
		return nil
	}
	out := new(RedshiftEventSubscriptionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedshiftEventSubscriptionSpec) DeepCopyInto(out *RedshiftEventSubscriptionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedshiftEventSubscriptionSpec.
func (in *RedshiftEventSubscriptionSpec) DeepCopy() *RedshiftEventSubscriptionSpec {
	if in == nil {
		return nil
	}
	out := new(RedshiftEventSubscriptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedshiftEventSubscriptionStatus) DeepCopyInto(out *RedshiftEventSubscriptionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedshiftEventSubscriptionStatus.
func (in *RedshiftEventSubscriptionStatus) DeepCopy() *RedshiftEventSubscriptionStatus {
	if in == nil {
		return nil
	}
	out := new(RedshiftEventSubscriptionStatus)
	in.DeepCopyInto(out)
	return out
}
