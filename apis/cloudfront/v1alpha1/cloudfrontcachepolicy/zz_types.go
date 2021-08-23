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
// +groupName=cloudfront.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/cloudfront/v1alpha1"
)

type CloudfrontCachePolicyObservation struct {
}

type CloudfrontCachePolicyParameters struct {
	Comment *string `json:"comment,omitempty" tf:"comment"`

	DefaultTtl *int64 `json:"defaultTtl,omitempty" tf:"default_ttl"`

	Etag *string `json:"etag,omitempty" tf:"etag"`

	MaxTtl *int64 `json:"maxTtl,omitempty" tf:"max_ttl"`

	MinTtl *int64 `json:"minTtl,omitempty" tf:"min_ttl"`

	Name string `json:"name" tf:"name"`

	ParametersInCacheKeyAndForwardedToOrigin []ParametersInCacheKeyAndForwardedToOriginParameters `json:"parametersInCacheKeyAndForwardedToOrigin,omitempty" tf:"parameters_in_cache_key_and_forwarded_to_origin"`
}

type CookiesConfigObservation struct {
}

type CookiesConfigParameters struct {
	CookieBehavior string `json:"cookieBehavior" tf:"cookie_behavior"`

	Cookies []CookiesParameters `json:"cookies,omitempty" tf:"cookies"`
}

type CookiesObservation struct {
}

type CookiesParameters struct {
	Items []string `json:"items,omitempty" tf:"items"`
}

type HeadersConfigObservation struct {
}

type HeadersConfigParameters struct {
	HeaderBehavior *string `json:"headerBehavior,omitempty" tf:"header_behavior"`

	Headers []HeadersParameters `json:"headers,omitempty" tf:"headers"`
}

type HeadersObservation struct {
}

type HeadersParameters struct {
	Items []string `json:"items,omitempty" tf:"items"`
}

type ParametersInCacheKeyAndForwardedToOriginObservation struct {
}

type ParametersInCacheKeyAndForwardedToOriginParameters struct {
	CookiesConfig []CookiesConfigParameters `json:"cookiesConfig" tf:"cookies_config"`

	EnableAcceptEncodingBrotli *bool `json:"enableAcceptEncodingBrotli,omitempty" tf:"enable_accept_encoding_brotli"`

	EnableAcceptEncodingGzip *bool `json:"enableAcceptEncodingGzip,omitempty" tf:"enable_accept_encoding_gzip"`

	HeadersConfig []HeadersConfigParameters `json:"headersConfig" tf:"headers_config"`

	QueryStringsConfig []QueryStringsConfigParameters `json:"queryStringsConfig" tf:"query_strings_config"`
}

type QueryStringsConfigObservation struct {
}

type QueryStringsConfigParameters struct {
	QueryStringBehavior string `json:"queryStringBehavior" tf:"query_string_behavior"`

	QueryStrings []QueryStringsParameters `json:"queryStrings,omitempty" tf:"query_strings"`
}

type QueryStringsObservation struct {
}

type QueryStringsParameters struct {
	Items []string `json:"items,omitempty" tf:"items"`
}

// CloudfrontCachePolicySpec defines the desired state of CloudfrontCachePolicy
type CloudfrontCachePolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CloudfrontCachePolicyParameters `json:"forProvider"`
}

// CloudfrontCachePolicyStatus defines the observed state of CloudfrontCachePolicy.
type CloudfrontCachePolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CloudfrontCachePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudfrontCachePolicy is the Schema for the CloudfrontCachePolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CloudfrontCachePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudfrontCachePolicySpec   `json:"spec"`
	Status            CloudfrontCachePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudfrontCachePolicyList contains a list of CloudfrontCachePolicys
type CloudfrontCachePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudfrontCachePolicy `json:"items"`
}

// Repository type metadata.
var (
	CloudfrontCachePolicyKind             = "CloudfrontCachePolicy"
	CloudfrontCachePolicyGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: CloudfrontCachePolicyKind}.String()
	CloudfrontCachePolicyKindAPIVersion   = CloudfrontCachePolicyKind + "." + v1alpha1.GroupVersion.String()
	CloudfrontCachePolicyGroupVersionKind = v1alpha1.GroupVersion.WithKind(CloudfrontCachePolicyKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&CloudfrontCachePolicy{}, &CloudfrontCachePolicyList{})
}
