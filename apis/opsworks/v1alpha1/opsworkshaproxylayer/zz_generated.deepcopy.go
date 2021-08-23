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
func (in *EbsVolumeObservation) DeepCopyInto(out *EbsVolumeObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EbsVolumeObservation.
func (in *EbsVolumeObservation) DeepCopy() *EbsVolumeObservation {
	if in == nil {
		return nil
	}
	out := new(EbsVolumeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EbsVolumeParameters) DeepCopyInto(out *EbsVolumeParameters) {
	*out = *in
	if in.Encrypted != nil {
		in, out := &in.Encrypted, &out.Encrypted
		*out = new(bool)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(int64)
		**out = **in
	}
	if in.RaidLevel != nil {
		in, out := &in.RaidLevel, &out.RaidLevel
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EbsVolumeParameters.
func (in *EbsVolumeParameters) DeepCopy() *EbsVolumeParameters {
	if in == nil {
		return nil
	}
	out := new(EbsVolumeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksHaproxyLayer) DeepCopyInto(out *OpsworksHaproxyLayer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksHaproxyLayer.
func (in *OpsworksHaproxyLayer) DeepCopy() *OpsworksHaproxyLayer {
	if in == nil {
		return nil
	}
	out := new(OpsworksHaproxyLayer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpsworksHaproxyLayer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksHaproxyLayerList) DeepCopyInto(out *OpsworksHaproxyLayerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpsworksHaproxyLayer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksHaproxyLayerList.
func (in *OpsworksHaproxyLayerList) DeepCopy() *OpsworksHaproxyLayerList {
	if in == nil {
		return nil
	}
	out := new(OpsworksHaproxyLayerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpsworksHaproxyLayerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksHaproxyLayerObservation) DeepCopyInto(out *OpsworksHaproxyLayerObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksHaproxyLayerObservation.
func (in *OpsworksHaproxyLayerObservation) DeepCopy() *OpsworksHaproxyLayerObservation {
	if in == nil {
		return nil
	}
	out := new(OpsworksHaproxyLayerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksHaproxyLayerParameters) DeepCopyInto(out *OpsworksHaproxyLayerParameters) {
	*out = *in
	if in.AutoAssignElasticIps != nil {
		in, out := &in.AutoAssignElasticIps, &out.AutoAssignElasticIps
		*out = new(bool)
		**out = **in
	}
	if in.AutoAssignPublicIps != nil {
		in, out := &in.AutoAssignPublicIps, &out.AutoAssignPublicIps
		*out = new(bool)
		**out = **in
	}
	if in.AutoHealing != nil {
		in, out := &in.AutoHealing, &out.AutoHealing
		*out = new(bool)
		**out = **in
	}
	if in.CustomConfigureRecipes != nil {
		in, out := &in.CustomConfigureRecipes, &out.CustomConfigureRecipes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CustomDeployRecipes != nil {
		in, out := &in.CustomDeployRecipes, &out.CustomDeployRecipes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CustomInstanceProfileArn != nil {
		in, out := &in.CustomInstanceProfileArn, &out.CustomInstanceProfileArn
		*out = new(string)
		**out = **in
	}
	if in.CustomJson != nil {
		in, out := &in.CustomJson, &out.CustomJson
		*out = new(string)
		**out = **in
	}
	if in.CustomSecurityGroupIds != nil {
		in, out := &in.CustomSecurityGroupIds, &out.CustomSecurityGroupIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CustomSetupRecipes != nil {
		in, out := &in.CustomSetupRecipes, &out.CustomSetupRecipes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CustomShutdownRecipes != nil {
		in, out := &in.CustomShutdownRecipes, &out.CustomShutdownRecipes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CustomUndeployRecipes != nil {
		in, out := &in.CustomUndeployRecipes, &out.CustomUndeployRecipes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DrainElbOnShutdown != nil {
		in, out := &in.DrainElbOnShutdown, &out.DrainElbOnShutdown
		*out = new(bool)
		**out = **in
	}
	if in.EbsVolume != nil {
		in, out := &in.EbsVolume, &out.EbsVolume
		*out = make([]EbsVolumeParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ElasticLoadBalancer != nil {
		in, out := &in.ElasticLoadBalancer, &out.ElasticLoadBalancer
		*out = new(string)
		**out = **in
	}
	if in.HealthcheckMethod != nil {
		in, out := &in.HealthcheckMethod, &out.HealthcheckMethod
		*out = new(string)
		**out = **in
	}
	if in.HealthcheckUrl != nil {
		in, out := &in.HealthcheckUrl, &out.HealthcheckUrl
		*out = new(string)
		**out = **in
	}
	if in.InstallUpdatesOnBoot != nil {
		in, out := &in.InstallUpdatesOnBoot, &out.InstallUpdatesOnBoot
		*out = new(bool)
		**out = **in
	}
	if in.InstanceShutdownTimeout != nil {
		in, out := &in.InstanceShutdownTimeout, &out.InstanceShutdownTimeout
		*out = new(int64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.StatsEnabled != nil {
		in, out := &in.StatsEnabled, &out.StatsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.StatsUrl != nil {
		in, out := &in.StatsUrl, &out.StatsUrl
		*out = new(string)
		**out = **in
	}
	if in.StatsUser != nil {
		in, out := &in.StatsUser, &out.StatsUser
		*out = new(string)
		**out = **in
	}
	if in.SystemPackages != nil {
		in, out := &in.SystemPackages, &out.SystemPackages
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
	if in.UseEbsOptimizedInstances != nil {
		in, out := &in.UseEbsOptimizedInstances, &out.UseEbsOptimizedInstances
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksHaproxyLayerParameters.
func (in *OpsworksHaproxyLayerParameters) DeepCopy() *OpsworksHaproxyLayerParameters {
	if in == nil {
		return nil
	}
	out := new(OpsworksHaproxyLayerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksHaproxyLayerSpec) DeepCopyInto(out *OpsworksHaproxyLayerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksHaproxyLayerSpec.
func (in *OpsworksHaproxyLayerSpec) DeepCopy() *OpsworksHaproxyLayerSpec {
	if in == nil {
		return nil
	}
	out := new(OpsworksHaproxyLayerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksHaproxyLayerStatus) DeepCopyInto(out *OpsworksHaproxyLayerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksHaproxyLayerStatus.
func (in *OpsworksHaproxyLayerStatus) DeepCopy() *OpsworksHaproxyLayerStatus {
	if in == nil {
		return nil
	}
	out := new(OpsworksHaproxyLayerStatus)
	in.DeepCopyInto(out)
	return out
}
