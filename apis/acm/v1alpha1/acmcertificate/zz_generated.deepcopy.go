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
func (in *AcmCertificate) DeepCopyInto(out *AcmCertificate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmCertificate.
func (in *AcmCertificate) DeepCopy() *AcmCertificate {
	if in == nil {
		return nil
	}
	out := new(AcmCertificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AcmCertificate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmCertificateList) DeepCopyInto(out *AcmCertificateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AcmCertificate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmCertificateList.
func (in *AcmCertificateList) DeepCopy() *AcmCertificateList {
	if in == nil {
		return nil
	}
	out := new(AcmCertificateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AcmCertificateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmCertificateObservation) DeepCopyInto(out *AcmCertificateObservation) {
	*out = *in
	if in.DomainValidationOptions != nil {
		in, out := &in.DomainValidationOptions, &out.DomainValidationOptions
		*out = make([]DomainValidationOptionsObservation, len(*in))
		copy(*out, *in)
	}
	if in.ValidationEmails != nil {
		in, out := &in.ValidationEmails, &out.ValidationEmails
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmCertificateObservation.
func (in *AcmCertificateObservation) DeepCopy() *AcmCertificateObservation {
	if in == nil {
		return nil
	}
	out := new(AcmCertificateObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmCertificateParameters) DeepCopyInto(out *AcmCertificateParameters) {
	*out = *in
	if in.CertificateAuthorityArn != nil {
		in, out := &in.CertificateAuthorityArn, &out.CertificateAuthorityArn
		*out = new(string)
		**out = **in
	}
	if in.CertificateBody != nil {
		in, out := &in.CertificateBody, &out.CertificateBody
		*out = new(string)
		**out = **in
	}
	if in.CertificateChain != nil {
		in, out := &in.CertificateChain, &out.CertificateChain
		*out = new(string)
		**out = **in
	}
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]OptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PrivateKey != nil {
		in, out := &in.PrivateKey, &out.PrivateKey
		*out = new(string)
		**out = **in
	}
	if in.SubjectAlternativeNames != nil {
		in, out := &in.SubjectAlternativeNames, &out.SubjectAlternativeNames
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
	if in.ValidationMethod != nil {
		in, out := &in.ValidationMethod, &out.ValidationMethod
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmCertificateParameters.
func (in *AcmCertificateParameters) DeepCopy() *AcmCertificateParameters {
	if in == nil {
		return nil
	}
	out := new(AcmCertificateParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmCertificateSpec) DeepCopyInto(out *AcmCertificateSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmCertificateSpec.
func (in *AcmCertificateSpec) DeepCopy() *AcmCertificateSpec {
	if in == nil {
		return nil
	}
	out := new(AcmCertificateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmCertificateStatus) DeepCopyInto(out *AcmCertificateStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmCertificateStatus.
func (in *AcmCertificateStatus) DeepCopy() *AcmCertificateStatus {
	if in == nil {
		return nil
	}
	out := new(AcmCertificateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainValidationOptionsObservation) DeepCopyInto(out *DomainValidationOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainValidationOptionsObservation.
func (in *DomainValidationOptionsObservation) DeepCopy() *DomainValidationOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(DomainValidationOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainValidationOptionsParameters) DeepCopyInto(out *DomainValidationOptionsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainValidationOptionsParameters.
func (in *DomainValidationOptionsParameters) DeepCopy() *DomainValidationOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(DomainValidationOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionsObservation) DeepCopyInto(out *OptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionsObservation.
func (in *OptionsObservation) DeepCopy() *OptionsObservation {
	if in == nil {
		return nil
	}
	out := new(OptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionsParameters) DeepCopyInto(out *OptionsParameters) {
	*out = *in
	if in.CertificateTransparencyLoggingPreference != nil {
		in, out := &in.CertificateTransparencyLoggingPreference, &out.CertificateTransparencyLoggingPreference
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionsParameters.
func (in *OptionsParameters) DeepCopy() *OptionsParameters {
	if in == nil {
		return nil
	}
	out := new(OptionsParameters)
	in.DeepCopyInto(out)
	return out
}
