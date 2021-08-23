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
func (in *DatasyncLocationS3) DeepCopyInto(out *DatasyncLocationS3) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasyncLocationS3.
func (in *DatasyncLocationS3) DeepCopy() *DatasyncLocationS3 {
	if in == nil {
		return nil
	}
	out := new(DatasyncLocationS3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatasyncLocationS3) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasyncLocationS3List) DeepCopyInto(out *DatasyncLocationS3List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DatasyncLocationS3, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasyncLocationS3List.
func (in *DatasyncLocationS3List) DeepCopy() *DatasyncLocationS3List {
	if in == nil {
		return nil
	}
	out := new(DatasyncLocationS3List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatasyncLocationS3List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasyncLocationS3Observation) DeepCopyInto(out *DatasyncLocationS3Observation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasyncLocationS3Observation.
func (in *DatasyncLocationS3Observation) DeepCopy() *DatasyncLocationS3Observation {
	if in == nil {
		return nil
	}
	out := new(DatasyncLocationS3Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasyncLocationS3Parameters) DeepCopyInto(out *DatasyncLocationS3Parameters) {
	*out = *in
	if in.AgentArns != nil {
		in, out := &in.AgentArns, &out.AgentArns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.S3Config != nil {
		in, out := &in.S3Config, &out.S3Config
		*out = make([]S3ConfigParameters, len(*in))
		copy(*out, *in)
	}
	if in.S3StorageClass != nil {
		in, out := &in.S3StorageClass, &out.S3StorageClass
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasyncLocationS3Parameters.
func (in *DatasyncLocationS3Parameters) DeepCopy() *DatasyncLocationS3Parameters {
	if in == nil {
		return nil
	}
	out := new(DatasyncLocationS3Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasyncLocationS3Spec) DeepCopyInto(out *DatasyncLocationS3Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasyncLocationS3Spec.
func (in *DatasyncLocationS3Spec) DeepCopy() *DatasyncLocationS3Spec {
	if in == nil {
		return nil
	}
	out := new(DatasyncLocationS3Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasyncLocationS3Status) DeepCopyInto(out *DatasyncLocationS3Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasyncLocationS3Status.
func (in *DatasyncLocationS3Status) DeepCopy() *DatasyncLocationS3Status {
	if in == nil {
		return nil
	}
	out := new(DatasyncLocationS3Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ConfigObservation) DeepCopyInto(out *S3ConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ConfigObservation.
func (in *S3ConfigObservation) DeepCopy() *S3ConfigObservation {
	if in == nil {
		return nil
	}
	out := new(S3ConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ConfigParameters) DeepCopyInto(out *S3ConfigParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ConfigParameters.
func (in *S3ConfigParameters) DeepCopy() *S3ConfigParameters {
	if in == nil {
		return nil
	}
	out := new(S3ConfigParameters)
	in.DeepCopyInto(out)
	return out
}
