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
func (in *WafregionalWebAclAssociation) DeepCopyInto(out *WafregionalWebAclAssociation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalWebAclAssociation.
func (in *WafregionalWebAclAssociation) DeepCopy() *WafregionalWebAclAssociation {
	if in == nil {
		return nil
	}
	out := new(WafregionalWebAclAssociation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WafregionalWebAclAssociation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalWebAclAssociationList) DeepCopyInto(out *WafregionalWebAclAssociationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WafregionalWebAclAssociation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalWebAclAssociationList.
func (in *WafregionalWebAclAssociationList) DeepCopy() *WafregionalWebAclAssociationList {
	if in == nil {
		return nil
	}
	out := new(WafregionalWebAclAssociationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WafregionalWebAclAssociationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalWebAclAssociationObservation) DeepCopyInto(out *WafregionalWebAclAssociationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalWebAclAssociationObservation.
func (in *WafregionalWebAclAssociationObservation) DeepCopy() *WafregionalWebAclAssociationObservation {
	if in == nil {
		return nil
	}
	out := new(WafregionalWebAclAssociationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalWebAclAssociationParameters) DeepCopyInto(out *WafregionalWebAclAssociationParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalWebAclAssociationParameters.
func (in *WafregionalWebAclAssociationParameters) DeepCopy() *WafregionalWebAclAssociationParameters {
	if in == nil {
		return nil
	}
	out := new(WafregionalWebAclAssociationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalWebAclAssociationSpec) DeepCopyInto(out *WafregionalWebAclAssociationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalWebAclAssociationSpec.
func (in *WafregionalWebAclAssociationSpec) DeepCopy() *WafregionalWebAclAssociationSpec {
	if in == nil {
		return nil
	}
	out := new(WafregionalWebAclAssociationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalWebAclAssociationStatus) DeepCopyInto(out *WafregionalWebAclAssociationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalWebAclAssociationStatus.
func (in *WafregionalWebAclAssociationStatus) DeepCopy() *WafregionalWebAclAssociationStatus {
	if in == nil {
		return nil
	}
	out := new(WafregionalWebAclAssociationStatus)
	in.DeepCopyInto(out)
	return out
}
