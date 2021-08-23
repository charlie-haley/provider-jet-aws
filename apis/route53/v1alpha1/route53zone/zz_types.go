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
// +groupName=route53.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/route53/v1alpha1"
)

type Route53ZoneObservation struct {
	NameServers []string `json:"nameServers" tf:"name_servers"`

	ZoneId string `json:"zoneId" tf:"zone_id"`
}

type Route53ZoneParameters struct {
	Comment *string `json:"comment,omitempty" tf:"comment"`

	DelegationSetId *string `json:"delegationSetId,omitempty" tf:"delegation_set_id"`

	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy"`

	Name string `json:"name" tf:"name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	Vpc []VpcParameters `json:"vpc,omitempty" tf:"vpc"`
}

type VpcObservation struct {
}

type VpcParameters struct {
	VpcId string `json:"vpcId" tf:"vpc_id"`

	VpcRegion *string `json:"vpcRegion,omitempty" tf:"vpc_region"`
}

// Route53ZoneSpec defines the desired state of Route53Zone
type Route53ZoneSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Route53ZoneParameters `json:"forProvider"`
}

// Route53ZoneStatus defines the observed state of Route53Zone.
type Route53ZoneStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Route53ZoneObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Route53Zone is the Schema for the Route53Zones API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Route53Zone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Route53ZoneSpec   `json:"spec"`
	Status            Route53ZoneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route53ZoneList contains a list of Route53Zones
type Route53ZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route53Zone `json:"items"`
}

// Repository type metadata.
var (
	Route53ZoneKind             = "Route53Zone"
	Route53ZoneGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: Route53ZoneKind}.String()
	Route53ZoneKindAPIVersion   = Route53ZoneKind + "." + v1alpha1.GroupVersion.String()
	Route53ZoneGroupVersionKind = v1alpha1.GroupVersion.WithKind(Route53ZoneKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&Route53Zone{}, &Route53ZoneList{})
}
