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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CustomErrorResponseObservation struct {
}

type CustomErrorResponseParameters struct {

	// +kubebuilder:validation:Optional
	ErrorCachingMinTTL *int64 `json:"errorCachingMinTtl,omitempty" tf:"error_caching_min_ttl,omitempty"`

	// +kubebuilder:validation:Required
	ErrorCode *int64 `json:"errorCode" tf:"error_code,omitempty"`

	// +kubebuilder:validation:Optional
	ResponseCode *int64 `json:"responseCode,omitempty" tf:"response_code,omitempty"`

	// +kubebuilder:validation:Optional
	ResponsePagePath *string `json:"responsePagePath,omitempty" tf:"response_page_path,omitempty"`
}

type CustomHeaderObservation struct {
}

type CustomHeaderParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type CustomOriginConfigObservation struct {
}

type CustomOriginConfigParameters struct {

	// +kubebuilder:validation:Required
	HTTPPort *int64 `json:"httpPort" tf:"http_port,omitempty"`

	// +kubebuilder:validation:Required
	HTTPSPort *int64 `json:"httpsPort" tf:"https_port,omitempty"`

	// +kubebuilder:validation:Optional
	OriginKeepaliveTimeout *int64 `json:"originKeepaliveTimeout,omitempty" tf:"origin_keepalive_timeout,omitempty"`

	// +kubebuilder:validation:Required
	OriginProtocolPolicy *string `json:"originProtocolPolicy" tf:"origin_protocol_policy,omitempty"`

	// +kubebuilder:validation:Optional
	OriginReadTimeout *int64 `json:"originReadTimeout,omitempty" tf:"origin_read_timeout,omitempty"`

	// +kubebuilder:validation:Required
	OriginSSLProtocols []*string `json:"originSslProtocols" tf:"origin_ssl_protocols,omitempty"`
}

type DefaultCacheBehaviorObservation struct {
}

type DefaultCacheBehaviorParameters struct {

	// +kubebuilder:validation:Required
	AllowedMethods []*string `json:"allowedMethods" tf:"allowed_methods,omitempty"`

	// +kubebuilder:validation:Optional
	CachePolicyID *string `json:"cachePolicyId,omitempty" tf:"cache_policy_id,omitempty"`

	// +kubebuilder:validation:Required
	CachedMethods []*string `json:"cachedMethods" tf:"cached_methods,omitempty"`

	// +kubebuilder:validation:Optional
	Compress *bool `json:"compress,omitempty" tf:"compress,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultTTL *int64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// +kubebuilder:validation:Optional
	FieldLevelEncryptionID *string `json:"fieldLevelEncryptionId,omitempty" tf:"field_level_encryption_id,omitempty"`

	// +kubebuilder:validation:Optional
	ForwardedValues []ForwardedValuesParameters `json:"forwardedValues,omitempty" tf:"forwarded_values,omitempty"`

	// +kubebuilder:validation:Optional
	FunctionAssociation []FunctionAssociationParameters `json:"functionAssociation,omitempty" tf:"function_association,omitempty"`

	// +kubebuilder:validation:Optional
	LambdaFunctionAssociation []LambdaFunctionAssociationParameters `json:"lambdaFunctionAssociation,omitempty" tf:"lambda_function_association,omitempty"`

	// +kubebuilder:validation:Optional
	MaxTTL *int64 `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	// +kubebuilder:validation:Optional
	MinTTL *int64 `json:"minTtl,omitempty" tf:"min_ttl,omitempty"`

	// +kubebuilder:validation:Optional
	OriginRequestPolicyID *string `json:"originRequestPolicyId,omitempty" tf:"origin_request_policy_id,omitempty"`

	// +kubebuilder:validation:Optional
	RealtimeLogConfigArn *string `json:"realtimeLogConfigArn,omitempty" tf:"realtime_log_config_arn,omitempty"`

	// +kubebuilder:validation:Optional
	SmoothStreaming *bool `json:"smoothStreaming,omitempty" tf:"smooth_streaming,omitempty"`

	// +kubebuilder:validation:Required
	TargetOriginID *string `json:"targetOriginId" tf:"target_origin_id,omitempty"`

	// +kubebuilder:validation:Optional
	TrustedKeyGroups []*string `json:"trustedKeyGroups,omitempty" tf:"trusted_key_groups,omitempty"`

	// +kubebuilder:validation:Optional
	TrustedSigners []*string `json:"trustedSigners,omitempty" tf:"trusted_signers,omitempty"`

	// +kubebuilder:validation:Required
	ViewerProtocolPolicy *string `json:"viewerProtocolPolicy" tf:"viewer_protocol_policy,omitempty"`
}

type DistributionObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	CallerReference *string `json:"callerReference,omitempty" tf:"caller_reference,omitempty"`

	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	HostedZoneID *string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InProgressValidationBatches *int64 `json:"inProgressValidationBatches,omitempty" tf:"in_progress_validation_batches,omitempty"`

	LastModifiedTime *string `json:"lastModifiedTime,omitempty" tf:"last_modified_time,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	TrustedKeyGroups []TrustedKeyGroupsObservation `json:"trustedKeyGroups,omitempty" tf:"trusted_key_groups,omitempty"`

	TrustedSigners []TrustedSignersObservation `json:"trustedSigners,omitempty" tf:"trusted_signers,omitempty"`
}

type DistributionParameters struct {

	// +kubebuilder:validation:Optional
	Aliases []*string `json:"aliases,omitempty" tf:"aliases,omitempty"`

	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// +kubebuilder:validation:Optional
	CustomErrorResponse []CustomErrorResponseParameters `json:"customErrorResponse,omitempty" tf:"custom_error_response,omitempty"`

	// +kubebuilder:validation:Required
	DefaultCacheBehavior []DefaultCacheBehaviorParameters `json:"defaultCacheBehavior" tf:"default_cache_behavior,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultRootObject *string `json:"defaultRootObject,omitempty" tf:"default_root_object,omitempty"`

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPVersion *string `json:"httpVersion,omitempty" tf:"http_version,omitempty"`

	// +kubebuilder:validation:Optional
	IsIPv6Enabled *bool `json:"isIpv6Enabled,omitempty" tf:"is_ipv6_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	LoggingConfig []LoggingConfigParameters `json:"loggingConfig,omitempty" tf:"logging_config,omitempty"`

	// +kubebuilder:validation:Optional
	OrderedCacheBehavior []OrderedCacheBehaviorParameters `json:"orderedCacheBehavior,omitempty" tf:"ordered_cache_behavior,omitempty"`

	// +kubebuilder:validation:Required
	Origin []OriginParameters `json:"origin" tf:"origin,omitempty"`

	// +kubebuilder:validation:Optional
	OriginGroup []OriginGroupParameters `json:"originGroup,omitempty" tf:"origin_group,omitempty"`

	// +kubebuilder:validation:Optional
	PriceClass *string `json:"priceClass,omitempty" tf:"price_class,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	Restrictions []RestrictionsParameters `json:"restrictions" tf:"restrictions,omitempty"`

	// +kubebuilder:validation:Optional
	RetainOnDelete *bool `json:"retainOnDelete,omitempty" tf:"retain_on_delete,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	ViewerCertificate []ViewerCertificateParameters `json:"viewerCertificate" tf:"viewer_certificate,omitempty"`

	// +kubebuilder:validation:Optional
	WaitForDeployment *bool `json:"waitForDeployment,omitempty" tf:"wait_for_deployment,omitempty"`

	// +kubebuilder:validation:Optional
	WebACLID *string `json:"webAclId,omitempty" tf:"web_acl_id,omitempty"`
}

type FailoverCriteriaObservation struct {
}

type FailoverCriteriaParameters struct {

	// +kubebuilder:validation:Required
	StatusCodes []*int64 `json:"statusCodes" tf:"status_codes,omitempty"`
}

type ForwardedValuesCookiesObservation struct {
}

type ForwardedValuesCookiesParameters struct {

	// +kubebuilder:validation:Required
	Forward *string `json:"forward" tf:"forward,omitempty"`

	// +kubebuilder:validation:Optional
	WhitelistedNames []*string `json:"whitelistedNames,omitempty" tf:"whitelisted_names,omitempty"`
}

type ForwardedValuesObservation struct {
}

type ForwardedValuesParameters struct {

	// +kubebuilder:validation:Required
	Cookies []ForwardedValuesCookiesParameters `json:"cookies" tf:"cookies,omitempty"`

	// +kubebuilder:validation:Optional
	Headers []*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// +kubebuilder:validation:Required
	QueryString *bool `json:"queryString" tf:"query_string,omitempty"`

	// +kubebuilder:validation:Optional
	QueryStringCacheKeys []*string `json:"queryStringCacheKeys,omitempty" tf:"query_string_cache_keys,omitempty"`
}

type FunctionAssociationObservation struct {
}

type FunctionAssociationParameters struct {

	// +kubebuilder:validation:Required
	EventType *string `json:"eventType" tf:"event_type,omitempty"`

	// +kubebuilder:validation:Required
	FunctionArn *string `json:"functionArn" tf:"function_arn,omitempty"`
}

type GeoRestrictionObservation struct {
}

type GeoRestrictionParameters struct {

	// +kubebuilder:validation:Optional
	Locations []*string `json:"locations,omitempty" tf:"locations,omitempty"`

	// +kubebuilder:validation:Required
	RestrictionType *string `json:"restrictionType" tf:"restriction_type,omitempty"`
}

type ItemsObservation struct {
	KeyGroupID *string `json:"keyGroupId,omitempty" tf:"key_group_id,omitempty"`

	KeyPairIds []*string `json:"keyPairIds,omitempty" tf:"key_pair_ids,omitempty"`
}

type ItemsParameters struct {
}

type LambdaFunctionAssociationObservation struct {
}

type LambdaFunctionAssociationParameters struct {

	// +kubebuilder:validation:Required
	EventType *string `json:"eventType" tf:"event_type,omitempty"`

	// +kubebuilder:validation:Optional
	IncludeBody *bool `json:"includeBody,omitempty" tf:"include_body,omitempty"`

	// +kubebuilder:validation:Required
	LambdaArn *string `json:"lambdaArn" tf:"lambda_arn,omitempty"`
}

type LoggingConfigObservation struct {
}

type LoggingConfigParameters struct {

	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Optional
	IncludeCookies *bool `json:"includeCookies,omitempty" tf:"include_cookies,omitempty"`

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type MemberObservation struct {
}

type MemberParameters struct {

	// +kubebuilder:validation:Required
	OriginID *string `json:"originId" tf:"origin_id,omitempty"`
}

type OrderedCacheBehaviorForwardedValuesCookiesObservation struct {
}

type OrderedCacheBehaviorForwardedValuesCookiesParameters struct {

	// +kubebuilder:validation:Required
	Forward *string `json:"forward" tf:"forward,omitempty"`

	// +kubebuilder:validation:Optional
	WhitelistedNames []*string `json:"whitelistedNames,omitempty" tf:"whitelisted_names,omitempty"`
}

type OrderedCacheBehaviorForwardedValuesObservation struct {
}

type OrderedCacheBehaviorForwardedValuesParameters struct {

	// +kubebuilder:validation:Required
	Cookies []OrderedCacheBehaviorForwardedValuesCookiesParameters `json:"cookies" tf:"cookies,omitempty"`

	// +kubebuilder:validation:Optional
	Headers []*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// +kubebuilder:validation:Required
	QueryString *bool `json:"queryString" tf:"query_string,omitempty"`

	// +kubebuilder:validation:Optional
	QueryStringCacheKeys []*string `json:"queryStringCacheKeys,omitempty" tf:"query_string_cache_keys,omitempty"`
}

type OrderedCacheBehaviorFunctionAssociationObservation struct {
}

type OrderedCacheBehaviorFunctionAssociationParameters struct {

	// +kubebuilder:validation:Required
	EventType *string `json:"eventType" tf:"event_type,omitempty"`

	// +kubebuilder:validation:Required
	FunctionArn *string `json:"functionArn" tf:"function_arn,omitempty"`
}

type OrderedCacheBehaviorLambdaFunctionAssociationObservation struct {
}

type OrderedCacheBehaviorLambdaFunctionAssociationParameters struct {

	// +kubebuilder:validation:Required
	EventType *string `json:"eventType" tf:"event_type,omitempty"`

	// +kubebuilder:validation:Optional
	IncludeBody *bool `json:"includeBody,omitempty" tf:"include_body,omitempty"`

	// +kubebuilder:validation:Required
	LambdaArn *string `json:"lambdaArn" tf:"lambda_arn,omitempty"`
}

type OrderedCacheBehaviorObservation struct {
}

type OrderedCacheBehaviorParameters struct {

	// +kubebuilder:validation:Required
	AllowedMethods []*string `json:"allowedMethods" tf:"allowed_methods,omitempty"`

	// +kubebuilder:validation:Optional
	CachePolicyID *string `json:"cachePolicyId,omitempty" tf:"cache_policy_id,omitempty"`

	// +kubebuilder:validation:Required
	CachedMethods []*string `json:"cachedMethods" tf:"cached_methods,omitempty"`

	// +kubebuilder:validation:Optional
	Compress *bool `json:"compress,omitempty" tf:"compress,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultTTL *int64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// +kubebuilder:validation:Optional
	FieldLevelEncryptionID *string `json:"fieldLevelEncryptionId,omitempty" tf:"field_level_encryption_id,omitempty"`

	// +kubebuilder:validation:Optional
	ForwardedValues []OrderedCacheBehaviorForwardedValuesParameters `json:"forwardedValues,omitempty" tf:"forwarded_values,omitempty"`

	// +kubebuilder:validation:Optional
	FunctionAssociation []OrderedCacheBehaviorFunctionAssociationParameters `json:"functionAssociation,omitempty" tf:"function_association,omitempty"`

	// +kubebuilder:validation:Optional
	LambdaFunctionAssociation []OrderedCacheBehaviorLambdaFunctionAssociationParameters `json:"lambdaFunctionAssociation,omitempty" tf:"lambda_function_association,omitempty"`

	// +kubebuilder:validation:Optional
	MaxTTL *int64 `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	// +kubebuilder:validation:Optional
	MinTTL *int64 `json:"minTtl,omitempty" tf:"min_ttl,omitempty"`

	// +kubebuilder:validation:Optional
	OriginRequestPolicyID *string `json:"originRequestPolicyId,omitempty" tf:"origin_request_policy_id,omitempty"`

	// +kubebuilder:validation:Required
	PathPattern *string `json:"pathPattern" tf:"path_pattern,omitempty"`

	// +kubebuilder:validation:Optional
	RealtimeLogConfigArn *string `json:"realtimeLogConfigArn,omitempty" tf:"realtime_log_config_arn,omitempty"`

	// +kubebuilder:validation:Optional
	SmoothStreaming *bool `json:"smoothStreaming,omitempty" tf:"smooth_streaming,omitempty"`

	// +kubebuilder:validation:Required
	TargetOriginID *string `json:"targetOriginId" tf:"target_origin_id,omitempty"`

	// +kubebuilder:validation:Optional
	TrustedKeyGroups []*string `json:"trustedKeyGroups,omitempty" tf:"trusted_key_groups,omitempty"`

	// +kubebuilder:validation:Optional
	TrustedSigners []*string `json:"trustedSigners,omitempty" tf:"trusted_signers,omitempty"`

	// +kubebuilder:validation:Required
	ViewerProtocolPolicy *string `json:"viewerProtocolPolicy" tf:"viewer_protocol_policy,omitempty"`
}

type OriginGroupObservation struct {
}

type OriginGroupParameters struct {

	// +kubebuilder:validation:Required
	FailoverCriteria []FailoverCriteriaParameters `json:"failoverCriteria" tf:"failover_criteria,omitempty"`

	// +kubebuilder:validation:Required
	Member []MemberParameters `json:"member" tf:"member,omitempty"`

	// +kubebuilder:validation:Required
	OriginID *string `json:"originId" tf:"origin_id,omitempty"`
}

type OriginObservation struct {
}

type OriginParameters struct {

	// +kubebuilder:validation:Optional
	ConnectionAttempts *int64 `json:"connectionAttempts,omitempty" tf:"connection_attempts,omitempty"`

	// +kubebuilder:validation:Optional
	ConnectionTimeout *int64 `json:"connectionTimeout,omitempty" tf:"connection_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	CustomHeader []CustomHeaderParameters `json:"customHeader,omitempty" tf:"custom_header,omitempty"`

	// +kubebuilder:validation:Optional
	CustomOriginConfig []CustomOriginConfigParameters `json:"customOriginConfig,omitempty" tf:"custom_origin_config,omitempty"`

	// +kubebuilder:validation:Required
	DomainName *string `json:"domainName" tf:"domain_name,omitempty"`

	// +kubebuilder:validation:Required
	OriginID *string `json:"originId" tf:"origin_id,omitempty"`

	// +kubebuilder:validation:Optional
	OriginPath *string `json:"originPath,omitempty" tf:"origin_path,omitempty"`

	// +kubebuilder:validation:Optional
	OriginShield []OriginShieldParameters `json:"originShield,omitempty" tf:"origin_shield,omitempty"`

	// +kubebuilder:validation:Optional
	S3OriginConfig []S3OriginConfigParameters `json:"s3OriginConfig,omitempty" tf:"s3_origin_config,omitempty"`
}

type OriginShieldObservation struct {
}

type OriginShieldParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	OriginShieldRegion *string `json:"originShieldRegion" tf:"origin_shield_region,omitempty"`
}

type RestrictionsObservation struct {
}

type RestrictionsParameters struct {

	// +kubebuilder:validation:Required
	GeoRestriction []GeoRestrictionParameters `json:"geoRestriction" tf:"geo_restriction,omitempty"`
}

type S3OriginConfigObservation struct {
}

type S3OriginConfigParameters struct {

	// +kubebuilder:validation:Required
	OriginAccessIdentity *string `json:"originAccessIdentity" tf:"origin_access_identity,omitempty"`
}

type TrustedKeyGroupsObservation struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	Items []ItemsObservation `json:"items,omitempty" tf:"items,omitempty"`
}

type TrustedKeyGroupsParameters struct {
}

type TrustedSignersItemsObservation struct {
	AwsAccountNumber *string `json:"awsAccountNumber,omitempty" tf:"aws_account_number,omitempty"`

	KeyPairIds []*string `json:"keyPairIds,omitempty" tf:"key_pair_ids,omitempty"`
}

type TrustedSignersItemsParameters struct {
}

type TrustedSignersObservation struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	Items []TrustedSignersItemsObservation `json:"items,omitempty" tf:"items,omitempty"`
}

type TrustedSignersParameters struct {
}

type ViewerCertificateObservation struct {
}

type ViewerCertificateParameters struct {

	// +kubebuilder:validation:Optional
	AcmCertificateArn *string `json:"acmCertificateArn,omitempty" tf:"acm_certificate_arn,omitempty"`

	// +kubebuilder:validation:Optional
	CloudfrontDefaultCertificate *bool `json:"cloudfrontDefaultCertificate,omitempty" tf:"cloudfront_default_certificate,omitempty"`

	// +kubebuilder:validation:Optional
	IAMCertificateID *string `json:"iamCertificateId,omitempty" tf:"iam_certificate_id,omitempty"`

	// +kubebuilder:validation:Optional
	MinimumProtocolVersion *string `json:"minimumProtocolVersion,omitempty" tf:"minimum_protocol_version,omitempty"`

	// +kubebuilder:validation:Optional
	SSLSupportMethod *string `json:"sslSupportMethod,omitempty" tf:"ssl_support_method,omitempty"`
}

// DistributionSpec defines the desired state of Distribution
type DistributionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DistributionParameters `json:"forProvider"`
}

// DistributionStatus defines the observed state of Distribution.
type DistributionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DistributionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Distribution is the Schema for the Distributions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Distribution struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DistributionSpec   `json:"spec"`
	Status            DistributionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DistributionList contains a list of Distributions
type DistributionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Distribution `json:"items"`
}

// Repository type metadata.
var (
	Distribution_Kind             = "Distribution"
	Distribution_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Distribution_Kind}.String()
	Distribution_KindAPIVersion   = Distribution_Kind + "." + CRDGroupVersion.String()
	Distribution_GroupVersionKind = CRDGroupVersion.WithKind(Distribution_Kind)
)

func init() {
	SchemeBuilder.Register(&Distribution{}, &DistributionList{})
}
