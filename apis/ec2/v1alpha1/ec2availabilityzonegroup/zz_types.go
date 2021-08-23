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
// +groupName=ec2.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1"
)

type Ec2AvailabilityZoneGroupObservation struct {
}

type Ec2AvailabilityZoneGroupParameters struct {
	GroupName string `json:"groupName" tf:"group_name"`

	OptInStatus string `json:"optInStatus" tf:"opt_in_status"`
}

// Ec2AvailabilityZoneGroupSpec defines the desired state of Ec2AvailabilityZoneGroup
type Ec2AvailabilityZoneGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Ec2AvailabilityZoneGroupParameters `json:"forProvider"`
}

// Ec2AvailabilityZoneGroupStatus defines the observed state of Ec2AvailabilityZoneGroup.
type Ec2AvailabilityZoneGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Ec2AvailabilityZoneGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2AvailabilityZoneGroup is the Schema for the Ec2AvailabilityZoneGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Ec2AvailabilityZoneGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2AvailabilityZoneGroupSpec   `json:"spec"`
	Status            Ec2AvailabilityZoneGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2AvailabilityZoneGroupList contains a list of Ec2AvailabilityZoneGroups
type Ec2AvailabilityZoneGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2AvailabilityZoneGroup `json:"items"`
}

// Repository type metadata.
var (
	Ec2AvailabilityZoneGroupKind             = "Ec2AvailabilityZoneGroup"
	Ec2AvailabilityZoneGroupGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: Ec2AvailabilityZoneGroupKind}.String()
	Ec2AvailabilityZoneGroupKindAPIVersion   = Ec2AvailabilityZoneGroupKind + "." + v1alpha1.GroupVersion.String()
	Ec2AvailabilityZoneGroupGroupVersionKind = v1alpha1.GroupVersion.WithKind(Ec2AvailabilityZoneGroupKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&Ec2AvailabilityZoneGroup{}, &Ec2AvailabilityZoneGroupList{})
}
