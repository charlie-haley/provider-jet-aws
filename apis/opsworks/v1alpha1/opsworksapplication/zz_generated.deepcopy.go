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
func (in *AppSourceObservation) DeepCopyInto(out *AppSourceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSourceObservation.
func (in *AppSourceObservation) DeepCopy() *AppSourceObservation {
	if in == nil {
		return nil
	}
	out := new(AppSourceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSourceParameters) DeepCopyInto(out *AppSourceParameters) {
	*out = *in
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(string)
		**out = **in
	}
	if in.SshKey != nil {
		in, out := &in.SshKey, &out.SshKey
		*out = new(string)
		**out = **in
	}
	if in.Url != nil {
		in, out := &in.Url, &out.Url
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSourceParameters.
func (in *AppSourceParameters) DeepCopy() *AppSourceParameters {
	if in == nil {
		return nil
	}
	out := new(AppSourceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentObservation) DeepCopyInto(out *EnvironmentObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentObservation.
func (in *EnvironmentObservation) DeepCopy() *EnvironmentObservation {
	if in == nil {
		return nil
	}
	out := new(EnvironmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentParameters) DeepCopyInto(out *EnvironmentParameters) {
	*out = *in
	if in.Secure != nil {
		in, out := &in.Secure, &out.Secure
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentParameters.
func (in *EnvironmentParameters) DeepCopy() *EnvironmentParameters {
	if in == nil {
		return nil
	}
	out := new(EnvironmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksApplication) DeepCopyInto(out *OpsworksApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksApplication.
func (in *OpsworksApplication) DeepCopy() *OpsworksApplication {
	if in == nil {
		return nil
	}
	out := new(OpsworksApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpsworksApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksApplicationList) DeepCopyInto(out *OpsworksApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpsworksApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksApplicationList.
func (in *OpsworksApplicationList) DeepCopy() *OpsworksApplicationList {
	if in == nil {
		return nil
	}
	out := new(OpsworksApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpsworksApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksApplicationObservation) DeepCopyInto(out *OpsworksApplicationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksApplicationObservation.
func (in *OpsworksApplicationObservation) DeepCopy() *OpsworksApplicationObservation {
	if in == nil {
		return nil
	}
	out := new(OpsworksApplicationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksApplicationParameters) DeepCopyInto(out *OpsworksApplicationParameters) {
	*out = *in
	if in.AppSource != nil {
		in, out := &in.AppSource, &out.AppSource
		*out = make([]AppSourceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AutoBundleOnDeploy != nil {
		in, out := &in.AutoBundleOnDeploy, &out.AutoBundleOnDeploy
		*out = new(string)
		**out = **in
	}
	if in.AwsFlowRubySettings != nil {
		in, out := &in.AwsFlowRubySettings, &out.AwsFlowRubySettings
		*out = new(string)
		**out = **in
	}
	if in.DataSourceArn != nil {
		in, out := &in.DataSourceArn, &out.DataSourceArn
		*out = new(string)
		**out = **in
	}
	if in.DataSourceDatabaseName != nil {
		in, out := &in.DataSourceDatabaseName, &out.DataSourceDatabaseName
		*out = new(string)
		**out = **in
	}
	if in.DataSourceType != nil {
		in, out := &in.DataSourceType, &out.DataSourceType
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DocumentRoot != nil {
		in, out := &in.DocumentRoot, &out.DocumentRoot
		*out = new(string)
		**out = **in
	}
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EnableSsl != nil {
		in, out := &in.EnableSsl, &out.EnableSsl
		*out = new(bool)
		**out = **in
	}
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = make([]EnvironmentParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RailsEnv != nil {
		in, out := &in.RailsEnv, &out.RailsEnv
		*out = new(string)
		**out = **in
	}
	if in.ShortName != nil {
		in, out := &in.ShortName, &out.ShortName
		*out = new(string)
		**out = **in
	}
	if in.SslConfiguration != nil {
		in, out := &in.SslConfiguration, &out.SslConfiguration
		*out = make([]SslConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksApplicationParameters.
func (in *OpsworksApplicationParameters) DeepCopy() *OpsworksApplicationParameters {
	if in == nil {
		return nil
	}
	out := new(OpsworksApplicationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksApplicationSpec) DeepCopyInto(out *OpsworksApplicationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksApplicationSpec.
func (in *OpsworksApplicationSpec) DeepCopy() *OpsworksApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(OpsworksApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksApplicationStatus) DeepCopyInto(out *OpsworksApplicationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksApplicationStatus.
func (in *OpsworksApplicationStatus) DeepCopy() *OpsworksApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(OpsworksApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SslConfigurationObservation) DeepCopyInto(out *SslConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SslConfigurationObservation.
func (in *SslConfigurationObservation) DeepCopy() *SslConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(SslConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SslConfigurationParameters) DeepCopyInto(out *SslConfigurationParameters) {
	*out = *in
	if in.Chain != nil {
		in, out := &in.Chain, &out.Chain
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SslConfigurationParameters.
func (in *SslConfigurationParameters) DeepCopy() *SslConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(SslConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}
