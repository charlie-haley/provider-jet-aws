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
func (in *Route53ResolverFirewallDomainList) DeepCopyInto(out *Route53ResolverFirewallDomainList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route53ResolverFirewallDomainList.
func (in *Route53ResolverFirewallDomainList) DeepCopy() *Route53ResolverFirewallDomainList {
	if in == nil {
		return nil
	}
	out := new(Route53ResolverFirewallDomainList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Route53ResolverFirewallDomainList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route53ResolverFirewallDomainListList) DeepCopyInto(out *Route53ResolverFirewallDomainListList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Route53ResolverFirewallDomainList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route53ResolverFirewallDomainListList.
func (in *Route53ResolverFirewallDomainListList) DeepCopy() *Route53ResolverFirewallDomainListList {
	if in == nil {
		return nil
	}
	out := new(Route53ResolverFirewallDomainListList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Route53ResolverFirewallDomainListList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route53ResolverFirewallDomainListObservation) DeepCopyInto(out *Route53ResolverFirewallDomainListObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route53ResolverFirewallDomainListObservation.
func (in *Route53ResolverFirewallDomainListObservation) DeepCopy() *Route53ResolverFirewallDomainListObservation {
	if in == nil {
		return nil
	}
	out := new(Route53ResolverFirewallDomainListObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route53ResolverFirewallDomainListParameters) DeepCopyInto(out *Route53ResolverFirewallDomainListParameters) {
	*out = *in
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route53ResolverFirewallDomainListParameters.
func (in *Route53ResolverFirewallDomainListParameters) DeepCopy() *Route53ResolverFirewallDomainListParameters {
	if in == nil {
		return nil
	}
	out := new(Route53ResolverFirewallDomainListParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route53ResolverFirewallDomainListSpec) DeepCopyInto(out *Route53ResolverFirewallDomainListSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route53ResolverFirewallDomainListSpec.
func (in *Route53ResolverFirewallDomainListSpec) DeepCopy() *Route53ResolverFirewallDomainListSpec {
	if in == nil {
		return nil
	}
	out := new(Route53ResolverFirewallDomainListSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route53ResolverFirewallDomainListStatus) DeepCopyInto(out *Route53ResolverFirewallDomainListStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route53ResolverFirewallDomainListStatus.
func (in *Route53ResolverFirewallDomainListStatus) DeepCopy() *Route53ResolverFirewallDomainListStatus {
	if in == nil {
		return nil
	}
	out := new(Route53ResolverFirewallDomainListStatus)
	in.DeepCopyInto(out)
	return out
}
