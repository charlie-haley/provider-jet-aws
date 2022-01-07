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
func (in *AudioCodecOptionsObservation) DeepCopyInto(out *AudioCodecOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AudioCodecOptionsObservation.
func (in *AudioCodecOptionsObservation) DeepCopy() *AudioCodecOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(AudioCodecOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AudioCodecOptionsParameters) DeepCopyInto(out *AudioCodecOptionsParameters) {
	*out = *in
	if in.BitDepth != nil {
		in, out := &in.BitDepth, &out.BitDepth
		*out = new(string)
		**out = **in
	}
	if in.BitOrder != nil {
		in, out := &in.BitOrder, &out.BitOrder
		*out = new(string)
		**out = **in
	}
	if in.Profile != nil {
		in, out := &in.Profile, &out.Profile
		*out = new(string)
		**out = **in
	}
	if in.Signed != nil {
		in, out := &in.Signed, &out.Signed
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AudioCodecOptionsParameters.
func (in *AudioCodecOptionsParameters) DeepCopy() *AudioCodecOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(AudioCodecOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AudioObservation) DeepCopyInto(out *AudioObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AudioObservation.
func (in *AudioObservation) DeepCopy() *AudioObservation {
	if in == nil {
		return nil
	}
	out := new(AudioObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AudioParameters) DeepCopyInto(out *AudioParameters) {
	*out = *in
	if in.AudioPackingMode != nil {
		in, out := &in.AudioPackingMode, &out.AudioPackingMode
		*out = new(string)
		**out = **in
	}
	if in.BitRate != nil {
		in, out := &in.BitRate, &out.BitRate
		*out = new(string)
		**out = **in
	}
	if in.Channels != nil {
		in, out := &in.Channels, &out.Channels
		*out = new(string)
		**out = **in
	}
	if in.Codec != nil {
		in, out := &in.Codec, &out.Codec
		*out = new(string)
		**out = **in
	}
	if in.SampleRate != nil {
		in, out := &in.SampleRate, &out.SampleRate
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AudioParameters.
func (in *AudioParameters) DeepCopy() *AudioParameters {
	if in == nil {
		return nil
	}
	out := new(AudioParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentConfigObservation) DeepCopyInto(out *ContentConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentConfigObservation.
func (in *ContentConfigObservation) DeepCopy() *ContentConfigObservation {
	if in == nil {
		return nil
	}
	out := new(ContentConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentConfigParameters) DeepCopyInto(out *ContentConfigParameters) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentConfigParameters.
func (in *ContentConfigParameters) DeepCopy() *ContentConfigParameters {
	if in == nil {
		return nil
	}
	out := new(ContentConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentConfigPermissionsObservation) DeepCopyInto(out *ContentConfigPermissionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentConfigPermissionsObservation.
func (in *ContentConfigPermissionsObservation) DeepCopy() *ContentConfigPermissionsObservation {
	if in == nil {
		return nil
	}
	out := new(ContentConfigPermissionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentConfigPermissionsParameters) DeepCopyInto(out *ContentConfigPermissionsParameters) {
	*out = *in
	if in.Access != nil {
		in, out := &in.Access, &out.Access
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Grantee != nil {
		in, out := &in.Grantee, &out.Grantee
		*out = new(string)
		**out = **in
	}
	if in.GranteeType != nil {
		in, out := &in.GranteeType, &out.GranteeType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentConfigPermissionsParameters.
func (in *ContentConfigPermissionsParameters) DeepCopy() *ContentConfigPermissionsParameters {
	if in == nil {
		return nil
	}
	out := new(ContentConfigPermissionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationsObservation) DeepCopyInto(out *NotificationsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationsObservation.
func (in *NotificationsObservation) DeepCopy() *NotificationsObservation {
	if in == nil {
		return nil
	}
	out := new(NotificationsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationsParameters) DeepCopyInto(out *NotificationsParameters) {
	*out = *in
	if in.Completed != nil {
		in, out := &in.Completed, &out.Completed
		*out = new(string)
		**out = **in
	}
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = new(string)
		**out = **in
	}
	if in.Progressing != nil {
		in, out := &in.Progressing, &out.Progressing
		*out = new(string)
		**out = **in
	}
	if in.Warning != nil {
		in, out := &in.Warning, &out.Warning
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationsParameters.
func (in *NotificationsParameters) DeepCopy() *NotificationsParameters {
	if in == nil {
		return nil
	}
	out := new(NotificationsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pipeline) DeepCopyInto(out *Pipeline) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pipeline.
func (in *Pipeline) DeepCopy() *Pipeline {
	if in == nil {
		return nil
	}
	out := new(Pipeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pipeline) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineList) DeepCopyInto(out *PipelineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pipeline, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineList.
func (in *PipelineList) DeepCopy() *PipelineList {
	if in == nil {
		return nil
	}
	out := new(PipelineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PipelineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineObservation) DeepCopyInto(out *PipelineObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineObservation.
func (in *PipelineObservation) DeepCopy() *PipelineObservation {
	if in == nil {
		return nil
	}
	out := new(PipelineObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineParameters) DeepCopyInto(out *PipelineParameters) {
	*out = *in
	if in.AwsKMSKeyArn != nil {
		in, out := &in.AwsKMSKeyArn, &out.AwsKMSKeyArn
		*out = new(string)
		**out = **in
	}
	if in.ContentConfig != nil {
		in, out := &in.ContentConfig, &out.ContentConfig
		*out = make([]ContentConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ContentConfigPermissions != nil {
		in, out := &in.ContentConfigPermissions, &out.ContentConfigPermissions
		*out = make([]ContentConfigPermissionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InputBucket != nil {
		in, out := &in.InputBucket, &out.InputBucket
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Notifications != nil {
		in, out := &in.Notifications, &out.Notifications
		*out = make([]NotificationsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OutputBucket != nil {
		in, out := &in.OutputBucket, &out.OutputBucket
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.ThumbnailConfig != nil {
		in, out := &in.ThumbnailConfig, &out.ThumbnailConfig
		*out = make([]ThumbnailConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ThumbnailConfigPermissions != nil {
		in, out := &in.ThumbnailConfigPermissions, &out.ThumbnailConfigPermissions
		*out = make([]ThumbnailConfigPermissionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineParameters.
func (in *PipelineParameters) DeepCopy() *PipelineParameters {
	if in == nil {
		return nil
	}
	out := new(PipelineParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpec) DeepCopyInto(out *PipelineSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpec.
func (in *PipelineSpec) DeepCopy() *PipelineSpec {
	if in == nil {
		return nil
	}
	out := new(PipelineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineStatus) DeepCopyInto(out *PipelineStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineStatus.
func (in *PipelineStatus) DeepCopy() *PipelineStatus {
	if in == nil {
		return nil
	}
	out := new(PipelineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Preset) DeepCopyInto(out *Preset) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Preset.
func (in *Preset) DeepCopy() *Preset {
	if in == nil {
		return nil
	}
	out := new(Preset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Preset) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PresetList) DeepCopyInto(out *PresetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Preset, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PresetList.
func (in *PresetList) DeepCopy() *PresetList {
	if in == nil {
		return nil
	}
	out := new(PresetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PresetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PresetObservation) DeepCopyInto(out *PresetObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PresetObservation.
func (in *PresetObservation) DeepCopy() *PresetObservation {
	if in == nil {
		return nil
	}
	out := new(PresetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PresetParameters) DeepCopyInto(out *PresetParameters) {
	*out = *in
	if in.Audio != nil {
		in, out := &in.Audio, &out.Audio
		*out = make([]AudioParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AudioCodecOptions != nil {
		in, out := &in.AudioCodecOptions, &out.AudioCodecOptions
		*out = make([]AudioCodecOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Container != nil {
		in, out := &in.Container, &out.Container
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Thumbnails != nil {
		in, out := &in.Thumbnails, &out.Thumbnails
		*out = make([]ThumbnailsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Video != nil {
		in, out := &in.Video, &out.Video
		*out = make([]VideoParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VideoCodecOptions != nil {
		in, out := &in.VideoCodecOptions, &out.VideoCodecOptions
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.VideoWatermarks != nil {
		in, out := &in.VideoWatermarks, &out.VideoWatermarks
		*out = make([]VideoWatermarksParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PresetParameters.
func (in *PresetParameters) DeepCopy() *PresetParameters {
	if in == nil {
		return nil
	}
	out := new(PresetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PresetSpec) DeepCopyInto(out *PresetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PresetSpec.
func (in *PresetSpec) DeepCopy() *PresetSpec {
	if in == nil {
		return nil
	}
	out := new(PresetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PresetStatus) DeepCopyInto(out *PresetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PresetStatus.
func (in *PresetStatus) DeepCopy() *PresetStatus {
	if in == nil {
		return nil
	}
	out := new(PresetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThumbnailConfigObservation) DeepCopyInto(out *ThumbnailConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThumbnailConfigObservation.
func (in *ThumbnailConfigObservation) DeepCopy() *ThumbnailConfigObservation {
	if in == nil {
		return nil
	}
	out := new(ThumbnailConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThumbnailConfigParameters) DeepCopyInto(out *ThumbnailConfigParameters) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThumbnailConfigParameters.
func (in *ThumbnailConfigParameters) DeepCopy() *ThumbnailConfigParameters {
	if in == nil {
		return nil
	}
	out := new(ThumbnailConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThumbnailConfigPermissionsObservation) DeepCopyInto(out *ThumbnailConfigPermissionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThumbnailConfigPermissionsObservation.
func (in *ThumbnailConfigPermissionsObservation) DeepCopy() *ThumbnailConfigPermissionsObservation {
	if in == nil {
		return nil
	}
	out := new(ThumbnailConfigPermissionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThumbnailConfigPermissionsParameters) DeepCopyInto(out *ThumbnailConfigPermissionsParameters) {
	*out = *in
	if in.Access != nil {
		in, out := &in.Access, &out.Access
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Grantee != nil {
		in, out := &in.Grantee, &out.Grantee
		*out = new(string)
		**out = **in
	}
	if in.GranteeType != nil {
		in, out := &in.GranteeType, &out.GranteeType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThumbnailConfigPermissionsParameters.
func (in *ThumbnailConfigPermissionsParameters) DeepCopy() *ThumbnailConfigPermissionsParameters {
	if in == nil {
		return nil
	}
	out := new(ThumbnailConfigPermissionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThumbnailsObservation) DeepCopyInto(out *ThumbnailsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThumbnailsObservation.
func (in *ThumbnailsObservation) DeepCopy() *ThumbnailsObservation {
	if in == nil {
		return nil
	}
	out := new(ThumbnailsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThumbnailsParameters) DeepCopyInto(out *ThumbnailsParameters) {
	*out = *in
	if in.AspectRatio != nil {
		in, out := &in.AspectRatio, &out.AspectRatio
		*out = new(string)
		**out = **in
	}
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(string)
		**out = **in
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(string)
		**out = **in
	}
	if in.MaxHeight != nil {
		in, out := &in.MaxHeight, &out.MaxHeight
		*out = new(string)
		**out = **in
	}
	if in.MaxWidth != nil {
		in, out := &in.MaxWidth, &out.MaxWidth
		*out = new(string)
		**out = **in
	}
	if in.PaddingPolicy != nil {
		in, out := &in.PaddingPolicy, &out.PaddingPolicy
		*out = new(string)
		**out = **in
	}
	if in.Resolution != nil {
		in, out := &in.Resolution, &out.Resolution
		*out = new(string)
		**out = **in
	}
	if in.SizingPolicy != nil {
		in, out := &in.SizingPolicy, &out.SizingPolicy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThumbnailsParameters.
func (in *ThumbnailsParameters) DeepCopy() *ThumbnailsParameters {
	if in == nil {
		return nil
	}
	out := new(ThumbnailsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VideoObservation) DeepCopyInto(out *VideoObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VideoObservation.
func (in *VideoObservation) DeepCopy() *VideoObservation {
	if in == nil {
		return nil
	}
	out := new(VideoObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VideoParameters) DeepCopyInto(out *VideoParameters) {
	*out = *in
	if in.AspectRatio != nil {
		in, out := &in.AspectRatio, &out.AspectRatio
		*out = new(string)
		**out = **in
	}
	if in.BitRate != nil {
		in, out := &in.BitRate, &out.BitRate
		*out = new(string)
		**out = **in
	}
	if in.Codec != nil {
		in, out := &in.Codec, &out.Codec
		*out = new(string)
		**out = **in
	}
	if in.DisplayAspectRatio != nil {
		in, out := &in.DisplayAspectRatio, &out.DisplayAspectRatio
		*out = new(string)
		**out = **in
	}
	if in.FixedGop != nil {
		in, out := &in.FixedGop, &out.FixedGop
		*out = new(string)
		**out = **in
	}
	if in.FrameRate != nil {
		in, out := &in.FrameRate, &out.FrameRate
		*out = new(string)
		**out = **in
	}
	if in.KeyframesMaxDist != nil {
		in, out := &in.KeyframesMaxDist, &out.KeyframesMaxDist
		*out = new(string)
		**out = **in
	}
	if in.MaxFrameRate != nil {
		in, out := &in.MaxFrameRate, &out.MaxFrameRate
		*out = new(string)
		**out = **in
	}
	if in.MaxHeight != nil {
		in, out := &in.MaxHeight, &out.MaxHeight
		*out = new(string)
		**out = **in
	}
	if in.MaxWidth != nil {
		in, out := &in.MaxWidth, &out.MaxWidth
		*out = new(string)
		**out = **in
	}
	if in.PaddingPolicy != nil {
		in, out := &in.PaddingPolicy, &out.PaddingPolicy
		*out = new(string)
		**out = **in
	}
	if in.Resolution != nil {
		in, out := &in.Resolution, &out.Resolution
		*out = new(string)
		**out = **in
	}
	if in.SizingPolicy != nil {
		in, out := &in.SizingPolicy, &out.SizingPolicy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VideoParameters.
func (in *VideoParameters) DeepCopy() *VideoParameters {
	if in == nil {
		return nil
	}
	out := new(VideoParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VideoWatermarksObservation) DeepCopyInto(out *VideoWatermarksObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VideoWatermarksObservation.
func (in *VideoWatermarksObservation) DeepCopy() *VideoWatermarksObservation {
	if in == nil {
		return nil
	}
	out := new(VideoWatermarksObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VideoWatermarksParameters) DeepCopyInto(out *VideoWatermarksParameters) {
	*out = *in
	if in.HorizontalAlign != nil {
		in, out := &in.HorizontalAlign, &out.HorizontalAlign
		*out = new(string)
		**out = **in
	}
	if in.HorizontalOffset != nil {
		in, out := &in.HorizontalOffset, &out.HorizontalOffset
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MaxHeight != nil {
		in, out := &in.MaxHeight, &out.MaxHeight
		*out = new(string)
		**out = **in
	}
	if in.MaxWidth != nil {
		in, out := &in.MaxWidth, &out.MaxWidth
		*out = new(string)
		**out = **in
	}
	if in.Opacity != nil {
		in, out := &in.Opacity, &out.Opacity
		*out = new(string)
		**out = **in
	}
	if in.SizingPolicy != nil {
		in, out := &in.SizingPolicy, &out.SizingPolicy
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
	if in.VerticalAlign != nil {
		in, out := &in.VerticalAlign, &out.VerticalAlign
		*out = new(string)
		**out = **in
	}
	if in.VerticalOffset != nil {
		in, out := &in.VerticalOffset, &out.VerticalOffset
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VideoWatermarksParameters.
func (in *VideoWatermarksParameters) DeepCopy() *VideoWatermarksParameters {
	if in == nil {
		return nil
	}
	out := new(VideoWatermarksParameters)
	in.DeepCopyInto(out)
	return out
}
