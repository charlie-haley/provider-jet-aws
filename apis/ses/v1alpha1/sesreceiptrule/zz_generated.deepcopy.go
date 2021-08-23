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
func (in *AddHeaderActionObservation) DeepCopyInto(out *AddHeaderActionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddHeaderActionObservation.
func (in *AddHeaderActionObservation) DeepCopy() *AddHeaderActionObservation {
	if in == nil {
		return nil
	}
	out := new(AddHeaderActionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddHeaderActionParameters) DeepCopyInto(out *AddHeaderActionParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddHeaderActionParameters.
func (in *AddHeaderActionParameters) DeepCopy() *AddHeaderActionParameters {
	if in == nil {
		return nil
	}
	out := new(AddHeaderActionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BounceActionObservation) DeepCopyInto(out *BounceActionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BounceActionObservation.
func (in *BounceActionObservation) DeepCopy() *BounceActionObservation {
	if in == nil {
		return nil
	}
	out := new(BounceActionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BounceActionParameters) DeepCopyInto(out *BounceActionParameters) {
	*out = *in
	if in.StatusCode != nil {
		in, out := &in.StatusCode, &out.StatusCode
		*out = new(string)
		**out = **in
	}
	if in.TopicArn != nil {
		in, out := &in.TopicArn, &out.TopicArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BounceActionParameters.
func (in *BounceActionParameters) DeepCopy() *BounceActionParameters {
	if in == nil {
		return nil
	}
	out := new(BounceActionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LambdaActionObservation) DeepCopyInto(out *LambdaActionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LambdaActionObservation.
func (in *LambdaActionObservation) DeepCopy() *LambdaActionObservation {
	if in == nil {
		return nil
	}
	out := new(LambdaActionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LambdaActionParameters) DeepCopyInto(out *LambdaActionParameters) {
	*out = *in
	if in.InvocationType != nil {
		in, out := &in.InvocationType, &out.InvocationType
		*out = new(string)
		**out = **in
	}
	if in.TopicArn != nil {
		in, out := &in.TopicArn, &out.TopicArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LambdaActionParameters.
func (in *LambdaActionParameters) DeepCopy() *LambdaActionParameters {
	if in == nil {
		return nil
	}
	out := new(LambdaActionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ActionObservation) DeepCopyInto(out *S3ActionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ActionObservation.
func (in *S3ActionObservation) DeepCopy() *S3ActionObservation {
	if in == nil {
		return nil
	}
	out := new(S3ActionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ActionParameters) DeepCopyInto(out *S3ActionParameters) {
	*out = *in
	if in.KmsKeyArn != nil {
		in, out := &in.KmsKeyArn, &out.KmsKeyArn
		*out = new(string)
		**out = **in
	}
	if in.ObjectKeyPrefix != nil {
		in, out := &in.ObjectKeyPrefix, &out.ObjectKeyPrefix
		*out = new(string)
		**out = **in
	}
	if in.TopicArn != nil {
		in, out := &in.TopicArn, &out.TopicArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ActionParameters.
func (in *S3ActionParameters) DeepCopy() *S3ActionParameters {
	if in == nil {
		return nil
	}
	out := new(S3ActionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesReceiptRule) DeepCopyInto(out *SesReceiptRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesReceiptRule.
func (in *SesReceiptRule) DeepCopy() *SesReceiptRule {
	if in == nil {
		return nil
	}
	out := new(SesReceiptRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SesReceiptRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesReceiptRuleList) DeepCopyInto(out *SesReceiptRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SesReceiptRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesReceiptRuleList.
func (in *SesReceiptRuleList) DeepCopy() *SesReceiptRuleList {
	if in == nil {
		return nil
	}
	out := new(SesReceiptRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SesReceiptRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesReceiptRuleObservation) DeepCopyInto(out *SesReceiptRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesReceiptRuleObservation.
func (in *SesReceiptRuleObservation) DeepCopy() *SesReceiptRuleObservation {
	if in == nil {
		return nil
	}
	out := new(SesReceiptRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesReceiptRuleParameters) DeepCopyInto(out *SesReceiptRuleParameters) {
	*out = *in
	if in.AddHeaderAction != nil {
		in, out := &in.AddHeaderAction, &out.AddHeaderAction
		*out = make([]AddHeaderActionParameters, len(*in))
		copy(*out, *in)
	}
	if in.After != nil {
		in, out := &in.After, &out.After
		*out = new(string)
		**out = **in
	}
	if in.BounceAction != nil {
		in, out := &in.BounceAction, &out.BounceAction
		*out = make([]BounceActionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.LambdaAction != nil {
		in, out := &in.LambdaAction, &out.LambdaAction
		*out = make([]LambdaActionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Recipients != nil {
		in, out := &in.Recipients, &out.Recipients
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.S3Action != nil {
		in, out := &in.S3Action, &out.S3Action
		*out = make([]S3ActionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ScanEnabled != nil {
		in, out := &in.ScanEnabled, &out.ScanEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SnsAction != nil {
		in, out := &in.SnsAction, &out.SnsAction
		*out = make([]SnsActionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StopAction != nil {
		in, out := &in.StopAction, &out.StopAction
		*out = make([]StopActionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TlsPolicy != nil {
		in, out := &in.TlsPolicy, &out.TlsPolicy
		*out = new(string)
		**out = **in
	}
	if in.WorkmailAction != nil {
		in, out := &in.WorkmailAction, &out.WorkmailAction
		*out = make([]WorkmailActionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesReceiptRuleParameters.
func (in *SesReceiptRuleParameters) DeepCopy() *SesReceiptRuleParameters {
	if in == nil {
		return nil
	}
	out := new(SesReceiptRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesReceiptRuleSpec) DeepCopyInto(out *SesReceiptRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesReceiptRuleSpec.
func (in *SesReceiptRuleSpec) DeepCopy() *SesReceiptRuleSpec {
	if in == nil {
		return nil
	}
	out := new(SesReceiptRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesReceiptRuleStatus) DeepCopyInto(out *SesReceiptRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesReceiptRuleStatus.
func (in *SesReceiptRuleStatus) DeepCopy() *SesReceiptRuleStatus {
	if in == nil {
		return nil
	}
	out := new(SesReceiptRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsActionObservation) DeepCopyInto(out *SnsActionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsActionObservation.
func (in *SnsActionObservation) DeepCopy() *SnsActionObservation {
	if in == nil {
		return nil
	}
	out := new(SnsActionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsActionParameters) DeepCopyInto(out *SnsActionParameters) {
	*out = *in
	if in.Encoding != nil {
		in, out := &in.Encoding, &out.Encoding
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsActionParameters.
func (in *SnsActionParameters) DeepCopy() *SnsActionParameters {
	if in == nil {
		return nil
	}
	out := new(SnsActionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StopActionObservation) DeepCopyInto(out *StopActionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StopActionObservation.
func (in *StopActionObservation) DeepCopy() *StopActionObservation {
	if in == nil {
		return nil
	}
	out := new(StopActionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StopActionParameters) DeepCopyInto(out *StopActionParameters) {
	*out = *in
	if in.TopicArn != nil {
		in, out := &in.TopicArn, &out.TopicArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StopActionParameters.
func (in *StopActionParameters) DeepCopy() *StopActionParameters {
	if in == nil {
		return nil
	}
	out := new(StopActionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkmailActionObservation) DeepCopyInto(out *WorkmailActionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkmailActionObservation.
func (in *WorkmailActionObservation) DeepCopy() *WorkmailActionObservation {
	if in == nil {
		return nil
	}
	out := new(WorkmailActionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkmailActionParameters) DeepCopyInto(out *WorkmailActionParameters) {
	*out = *in
	if in.TopicArn != nil {
		in, out := &in.TopicArn, &out.TopicArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkmailActionParameters.
func (in *WorkmailActionParameters) DeepCopy() *WorkmailActionParameters {
	if in == nil {
		return nil
	}
	out := new(WorkmailActionParameters)
	in.DeepCopyInto(out)
	return out
}
