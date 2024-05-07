// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PolicyInitParameters struct {

	// The name of the VPN Server Configuration Policy member.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The attribute type of the VPN Server Configuration Policy member. Possible values are AADGroupId, CertificateGroupId and RadiusAzureGroupId.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The value of the attribute that is used for the VPN Server Configuration Policy member.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PolicyObservation struct {

	// The name of the VPN Server Configuration Policy member.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The attribute type of the VPN Server Configuration Policy member. Possible values are AADGroupId, CertificateGroupId and RadiusAzureGroupId.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The value of the attribute that is used for the VPN Server Configuration Policy member.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PolicyParameters struct {

	// The name of the VPN Server Configuration Policy member.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The attribute type of the VPN Server Configuration Policy member. Possible values are AADGroupId, CertificateGroupId and RadiusAzureGroupId.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// The value of the attribute that is used for the VPN Server Configuration Policy member.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type VPNServerConfigurationPolicyGroupInitParameters struct {

	// Is this a default VPN Server Configuration Policy Group? Defaults to false. Changing this forces a new resource to be created.
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// One or more policy blocks as documented below.
	Policy []PolicyInitParameters `json:"policy,omitempty" tf:"policy,omitempty"`

	// The priority of this VPN Server Configuration Policy Group. Defaults to 0.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`
}

type VPNServerConfigurationPolicyGroupObservation struct {

	// The ID of the VPN Server Configuration Policy Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Is this a default VPN Server Configuration Policy Group? Defaults to false. Changing this forces a new resource to be created.
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// One or more policy blocks as documented below.
	Policy []PolicyObservation `json:"policy,omitempty" tf:"policy,omitempty"`

	// The priority of this VPN Server Configuration Policy Group. Defaults to 0.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The ID of the VPN Server Configuration that the VPN Server Configuration Policy Group belongs to. Changing this forces a new resource to be created.
	VPNServerConfigurationID *string `json:"vpnServerConfigurationId,omitempty" tf:"vpn_server_configuration_id,omitempty"`
}

type VPNServerConfigurationPolicyGroupParameters struct {

	// Is this a default VPN Server Configuration Policy Group? Defaults to false. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// One or more policy blocks as documented below.
	// +kubebuilder:validation:Optional
	Policy []PolicyParameters `json:"policy,omitempty" tf:"policy,omitempty"`

	// The priority of this VPN Server Configuration Policy Group. Defaults to 0.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The ID of the VPN Server Configuration that the VPN Server Configuration Policy Group belongs to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta2.VPNServerConfiguration
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPNServerConfigurationID *string `json:"vpnServerConfigurationId,omitempty" tf:"vpn_server_configuration_id,omitempty"`

	// Reference to a VPNServerConfiguration in network to populate vpnServerConfigurationId.
	// +kubebuilder:validation:Optional
	VPNServerConfigurationIDRef *v1.Reference `json:"vpnServerConfigurationIdRef,omitempty" tf:"-"`

	// Selector for a VPNServerConfiguration in network to populate vpnServerConfigurationId.
	// +kubebuilder:validation:Optional
	VPNServerConfigurationIDSelector *v1.Selector `json:"vpnServerConfigurationIdSelector,omitempty" tf:"-"`
}

// VPNServerConfigurationPolicyGroupSpec defines the desired state of VPNServerConfigurationPolicyGroup
type VPNServerConfigurationPolicyGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPNServerConfigurationPolicyGroupParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider VPNServerConfigurationPolicyGroupInitParameters `json:"initProvider,omitempty"`
}

// VPNServerConfigurationPolicyGroupStatus defines the observed state of VPNServerConfigurationPolicyGroup.
type VPNServerConfigurationPolicyGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPNServerConfigurationPolicyGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VPNServerConfigurationPolicyGroup is the Schema for the VPNServerConfigurationPolicyGroups API. Manages a VPN Server Configuration Policy Group.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VPNServerConfigurationPolicyGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policy) || (has(self.initProvider) && has(self.initProvider.policy))",message="spec.forProvider.policy is a required parameter"
	Spec   VPNServerConfigurationPolicyGroupSpec   `json:"spec"`
	Status VPNServerConfigurationPolicyGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPNServerConfigurationPolicyGroupList contains a list of VPNServerConfigurationPolicyGroups
type VPNServerConfigurationPolicyGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPNServerConfigurationPolicyGroup `json:"items"`
}

// Repository type metadata.
var (
	VPNServerConfigurationPolicyGroup_Kind             = "VPNServerConfigurationPolicyGroup"
	VPNServerConfigurationPolicyGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPNServerConfigurationPolicyGroup_Kind}.String()
	VPNServerConfigurationPolicyGroup_KindAPIVersion   = VPNServerConfigurationPolicyGroup_Kind + "." + CRDGroupVersion.String()
	VPNServerConfigurationPolicyGroup_GroupVersionKind = CRDGroupVersion.WithKind(VPNServerConfigurationPolicyGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&VPNServerConfigurationPolicyGroup{}, &VPNServerConfigurationPolicyGroupList{})
}
