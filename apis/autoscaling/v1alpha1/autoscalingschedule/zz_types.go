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

// Code generated by terrajet. DO NOT EDIT.

// +kubebuilder:object:generate=true
// +groupName=autoscaling.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/autoscaling/v1alpha1"
)

type AutoscalingScheduleObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type AutoscalingScheduleParameters struct {
	AutoscalingGroupName string `json:"autoscalingGroupName" tf:"autoscaling_group_name"`

	DesiredCapacity *int64 `json:"desiredCapacity,omitempty" tf:"desired_capacity"`

	EndTime *string `json:"endTime,omitempty" tf:"end_time"`

	MaxSize *int64 `json:"maxSize,omitempty" tf:"max_size"`

	MinSize *int64 `json:"minSize,omitempty" tf:"min_size"`

	Recurrence *string `json:"recurrence,omitempty" tf:"recurrence"`

	ScheduledActionName string `json:"scheduledActionName" tf:"scheduled_action_name"`

	StartTime *string `json:"startTime,omitempty" tf:"start_time"`

	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone"`
}

// AutoscalingScheduleSpec defines the desired state of AutoscalingSchedule
type AutoscalingScheduleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AutoscalingScheduleParameters `json:"forProvider"`
}

// AutoscalingScheduleStatus defines the observed state of AutoscalingSchedule.
type AutoscalingScheduleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AutoscalingScheduleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AutoscalingSchedule is the Schema for the AutoscalingSchedules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AutoscalingSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoscalingScheduleSpec   `json:"spec"`
	Status            AutoscalingScheduleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutoscalingScheduleList contains a list of AutoscalingSchedules
type AutoscalingScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoscalingSchedule `json:"items"`
}

// Repository type metadata.
var (
	AutoscalingScheduleKind             = "AutoscalingSchedule"
	AutoscalingScheduleGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: AutoscalingScheduleKind}.String()
	AutoscalingScheduleKindAPIVersion   = AutoscalingScheduleKind + "." + v1alpha1.GroupVersion.String()
	AutoscalingScheduleGroupVersionKind = v1alpha1.GroupVersion.WithKind(AutoscalingScheduleKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&AutoscalingSchedule{}, &AutoscalingScheduleList{})
}
