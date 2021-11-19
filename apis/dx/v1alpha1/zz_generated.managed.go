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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this BgpPeer.
func (mg *BgpPeer) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this BgpPeer.
func (mg *BgpPeer) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this BgpPeer.
func (mg *BgpPeer) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this BgpPeer.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *BgpPeer) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this BgpPeer.
func (mg *BgpPeer) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this BgpPeer.
func (mg *BgpPeer) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this BgpPeer.
func (mg *BgpPeer) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this BgpPeer.
func (mg *BgpPeer) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this BgpPeer.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *BgpPeer) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this BgpPeer.
func (mg *BgpPeer) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Connection.
func (mg *Connection) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Connection.
func (mg *Connection) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this Connection.
func (mg *Connection) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Connection.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Connection) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this Connection.
func (mg *Connection) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Connection.
func (mg *Connection) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Connection.
func (mg *Connection) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this Connection.
func (mg *Connection) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Connection.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Connection) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this Connection.
func (mg *Connection) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ConnectionAssociation.
func (mg *ConnectionAssociation) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ConnectionAssociation.
func (mg *ConnectionAssociation) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ConnectionAssociation.
func (mg *ConnectionAssociation) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ConnectionAssociation.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ConnectionAssociation) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this ConnectionAssociation.
func (mg *ConnectionAssociation) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ConnectionAssociation.
func (mg *ConnectionAssociation) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ConnectionAssociation.
func (mg *ConnectionAssociation) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ConnectionAssociation.
func (mg *ConnectionAssociation) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ConnectionAssociation.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ConnectionAssociation) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this ConnectionAssociation.
func (mg *ConnectionAssociation) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Gateway.
func (mg *Gateway) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Gateway.
func (mg *Gateway) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this Gateway.
func (mg *Gateway) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Gateway.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Gateway) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this Gateway.
func (mg *Gateway) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Gateway.
func (mg *Gateway) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Gateway.
func (mg *Gateway) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this Gateway.
func (mg *Gateway) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Gateway.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Gateway) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this Gateway.
func (mg *Gateway) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this GatewayAssociation.
func (mg *GatewayAssociation) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this GatewayAssociation.
func (mg *GatewayAssociation) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this GatewayAssociation.
func (mg *GatewayAssociation) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this GatewayAssociation.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *GatewayAssociation) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this GatewayAssociation.
func (mg *GatewayAssociation) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this GatewayAssociation.
func (mg *GatewayAssociation) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this GatewayAssociation.
func (mg *GatewayAssociation) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this GatewayAssociation.
func (mg *GatewayAssociation) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this GatewayAssociation.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *GatewayAssociation) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this GatewayAssociation.
func (mg *GatewayAssociation) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this GatewayAssociationProposal.
func (mg *GatewayAssociationProposal) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this GatewayAssociationProposal.
func (mg *GatewayAssociationProposal) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this GatewayAssociationProposal.
func (mg *GatewayAssociationProposal) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this GatewayAssociationProposal.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *GatewayAssociationProposal) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this GatewayAssociationProposal.
func (mg *GatewayAssociationProposal) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this GatewayAssociationProposal.
func (mg *GatewayAssociationProposal) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this GatewayAssociationProposal.
func (mg *GatewayAssociationProposal) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this GatewayAssociationProposal.
func (mg *GatewayAssociationProposal) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this GatewayAssociationProposal.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *GatewayAssociationProposal) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this GatewayAssociationProposal.
func (mg *GatewayAssociationProposal) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this HostedPrivateVirtualInterface.
func (mg *HostedPrivateVirtualInterface) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this HostedPrivateVirtualInterface.
func (mg *HostedPrivateVirtualInterface) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this HostedPrivateVirtualInterface.
func (mg *HostedPrivateVirtualInterface) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this HostedPrivateVirtualInterface.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *HostedPrivateVirtualInterface) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this HostedPrivateVirtualInterface.
func (mg *HostedPrivateVirtualInterface) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this HostedPrivateVirtualInterface.
func (mg *HostedPrivateVirtualInterface) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this HostedPrivateVirtualInterface.
func (mg *HostedPrivateVirtualInterface) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this HostedPrivateVirtualInterface.
func (mg *HostedPrivateVirtualInterface) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this HostedPrivateVirtualInterface.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *HostedPrivateVirtualInterface) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this HostedPrivateVirtualInterface.
func (mg *HostedPrivateVirtualInterface) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this HostedPrivateVirtualInterfaceAccepter.
func (mg *HostedPrivateVirtualInterfaceAccepter) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this HostedPrivateVirtualInterfaceAccepter.
func (mg *HostedPrivateVirtualInterfaceAccepter) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this HostedPrivateVirtualInterfaceAccepter.
func (mg *HostedPrivateVirtualInterfaceAccepter) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this HostedPrivateVirtualInterfaceAccepter.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *HostedPrivateVirtualInterfaceAccepter) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this HostedPrivateVirtualInterfaceAccepter.
func (mg *HostedPrivateVirtualInterfaceAccepter) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this HostedPrivateVirtualInterfaceAccepter.
func (mg *HostedPrivateVirtualInterfaceAccepter) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this HostedPrivateVirtualInterfaceAccepter.
func (mg *HostedPrivateVirtualInterfaceAccepter) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this HostedPrivateVirtualInterfaceAccepter.
func (mg *HostedPrivateVirtualInterfaceAccepter) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this HostedPrivateVirtualInterfaceAccepter.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *HostedPrivateVirtualInterfaceAccepter) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this HostedPrivateVirtualInterfaceAccepter.
func (mg *HostedPrivateVirtualInterfaceAccepter) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this HostedPublicVirtualInterface.
func (mg *HostedPublicVirtualInterface) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this HostedPublicVirtualInterface.
func (mg *HostedPublicVirtualInterface) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this HostedPublicVirtualInterface.
func (mg *HostedPublicVirtualInterface) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this HostedPublicVirtualInterface.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *HostedPublicVirtualInterface) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this HostedPublicVirtualInterface.
func (mg *HostedPublicVirtualInterface) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this HostedPublicVirtualInterface.
func (mg *HostedPublicVirtualInterface) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this HostedPublicVirtualInterface.
func (mg *HostedPublicVirtualInterface) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this HostedPublicVirtualInterface.
func (mg *HostedPublicVirtualInterface) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this HostedPublicVirtualInterface.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *HostedPublicVirtualInterface) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this HostedPublicVirtualInterface.
func (mg *HostedPublicVirtualInterface) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this HostedPublicVirtualInterfaceAccepter.
func (mg *HostedPublicVirtualInterfaceAccepter) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this HostedPublicVirtualInterfaceAccepter.
func (mg *HostedPublicVirtualInterfaceAccepter) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this HostedPublicVirtualInterfaceAccepter.
func (mg *HostedPublicVirtualInterfaceAccepter) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this HostedPublicVirtualInterfaceAccepter.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *HostedPublicVirtualInterfaceAccepter) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this HostedPublicVirtualInterfaceAccepter.
func (mg *HostedPublicVirtualInterfaceAccepter) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this HostedPublicVirtualInterfaceAccepter.
func (mg *HostedPublicVirtualInterfaceAccepter) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this HostedPublicVirtualInterfaceAccepter.
func (mg *HostedPublicVirtualInterfaceAccepter) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this HostedPublicVirtualInterfaceAccepter.
func (mg *HostedPublicVirtualInterfaceAccepter) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this HostedPublicVirtualInterfaceAccepter.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *HostedPublicVirtualInterfaceAccepter) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this HostedPublicVirtualInterfaceAccepter.
func (mg *HostedPublicVirtualInterfaceAccepter) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this HostedTransitVirtualInterface.
func (mg *HostedTransitVirtualInterface) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this HostedTransitVirtualInterface.
func (mg *HostedTransitVirtualInterface) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this HostedTransitVirtualInterface.
func (mg *HostedTransitVirtualInterface) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this HostedTransitVirtualInterface.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *HostedTransitVirtualInterface) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this HostedTransitVirtualInterface.
func (mg *HostedTransitVirtualInterface) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this HostedTransitVirtualInterface.
func (mg *HostedTransitVirtualInterface) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this HostedTransitVirtualInterface.
func (mg *HostedTransitVirtualInterface) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this HostedTransitVirtualInterface.
func (mg *HostedTransitVirtualInterface) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this HostedTransitVirtualInterface.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *HostedTransitVirtualInterface) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this HostedTransitVirtualInterface.
func (mg *HostedTransitVirtualInterface) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this HostedTransitVirtualInterfaceAccepter.
func (mg *HostedTransitVirtualInterfaceAccepter) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this HostedTransitVirtualInterfaceAccepter.
func (mg *HostedTransitVirtualInterfaceAccepter) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this HostedTransitVirtualInterfaceAccepter.
func (mg *HostedTransitVirtualInterfaceAccepter) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this HostedTransitVirtualInterfaceAccepter.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *HostedTransitVirtualInterfaceAccepter) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this HostedTransitVirtualInterfaceAccepter.
func (mg *HostedTransitVirtualInterfaceAccepter) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this HostedTransitVirtualInterfaceAccepter.
func (mg *HostedTransitVirtualInterfaceAccepter) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this HostedTransitVirtualInterfaceAccepter.
func (mg *HostedTransitVirtualInterfaceAccepter) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this HostedTransitVirtualInterfaceAccepter.
func (mg *HostedTransitVirtualInterfaceAccepter) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this HostedTransitVirtualInterfaceAccepter.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *HostedTransitVirtualInterfaceAccepter) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this HostedTransitVirtualInterfaceAccepter.
func (mg *HostedTransitVirtualInterfaceAccepter) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Lag.
func (mg *Lag) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Lag.
func (mg *Lag) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this Lag.
func (mg *Lag) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Lag.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Lag) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this Lag.
func (mg *Lag) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Lag.
func (mg *Lag) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Lag.
func (mg *Lag) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this Lag.
func (mg *Lag) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Lag.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Lag) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this Lag.
func (mg *Lag) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this PrivateVirtualInterface.
func (mg *PrivateVirtualInterface) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this PrivateVirtualInterface.
func (mg *PrivateVirtualInterface) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this PrivateVirtualInterface.
func (mg *PrivateVirtualInterface) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this PrivateVirtualInterface.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *PrivateVirtualInterface) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this PrivateVirtualInterface.
func (mg *PrivateVirtualInterface) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this PrivateVirtualInterface.
func (mg *PrivateVirtualInterface) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this PrivateVirtualInterface.
func (mg *PrivateVirtualInterface) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this PrivateVirtualInterface.
func (mg *PrivateVirtualInterface) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this PrivateVirtualInterface.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *PrivateVirtualInterface) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this PrivateVirtualInterface.
func (mg *PrivateVirtualInterface) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this PublicVirtualInterface.
func (mg *PublicVirtualInterface) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this PublicVirtualInterface.
func (mg *PublicVirtualInterface) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this PublicVirtualInterface.
func (mg *PublicVirtualInterface) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this PublicVirtualInterface.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *PublicVirtualInterface) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this PublicVirtualInterface.
func (mg *PublicVirtualInterface) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this PublicVirtualInterface.
func (mg *PublicVirtualInterface) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this PublicVirtualInterface.
func (mg *PublicVirtualInterface) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this PublicVirtualInterface.
func (mg *PublicVirtualInterface) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this PublicVirtualInterface.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *PublicVirtualInterface) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this PublicVirtualInterface.
func (mg *PublicVirtualInterface) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this TransitVirtualInterface.
func (mg *TransitVirtualInterface) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this TransitVirtualInterface.
func (mg *TransitVirtualInterface) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this TransitVirtualInterface.
func (mg *TransitVirtualInterface) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this TransitVirtualInterface.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *TransitVirtualInterface) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this TransitVirtualInterface.
func (mg *TransitVirtualInterface) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this TransitVirtualInterface.
func (mg *TransitVirtualInterface) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this TransitVirtualInterface.
func (mg *TransitVirtualInterface) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this TransitVirtualInterface.
func (mg *TransitVirtualInterface) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this TransitVirtualInterface.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *TransitVirtualInterface) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this TransitVirtualInterface.
func (mg *TransitVirtualInterface) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
