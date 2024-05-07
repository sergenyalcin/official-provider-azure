// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type WindowsVirtualMachineGalleryImageReferenceInitParameters struct {

	// The Offer of the Gallery Image. Changing this forces a new resource to be created.
	Offer *string `json:"offer,omitempty" tf:"offer,omitempty"`

	// The Publisher of the Gallery Image. Changing this forces a new resource to be created.
	Publisher *string `json:"publisher,omitempty" tf:"publisher,omitempty"`

	// The SKU of the Gallery Image. Changing this forces a new resource to be created.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// The Version of the Gallery Image. Changing this forces a new resource to be created.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type WindowsVirtualMachineGalleryImageReferenceObservation struct {

	// The Offer of the Gallery Image. Changing this forces a new resource to be created.
	Offer *string `json:"offer,omitempty" tf:"offer,omitempty"`

	// The Publisher of the Gallery Image. Changing this forces a new resource to be created.
	Publisher *string `json:"publisher,omitempty" tf:"publisher,omitempty"`

	// The SKU of the Gallery Image. Changing this forces a new resource to be created.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// The Version of the Gallery Image. Changing this forces a new resource to be created.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type WindowsVirtualMachineGalleryImageReferenceParameters struct {

	// The Offer of the Gallery Image. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Offer *string `json:"offer" tf:"offer,omitempty"`

	// The Publisher of the Gallery Image. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Publisher *string `json:"publisher" tf:"publisher,omitempty"`

	// The SKU of the Gallery Image. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Sku *string `json:"sku" tf:"sku,omitempty"`

	// The Version of the Gallery Image. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Version *string `json:"version" tf:"version,omitempty"`
}

type WindowsVirtualMachineInboundNATRuleInitParameters struct {

	// The Backend Port associated with this NAT Rule. Changing this forces a new resource to be created.
	BackendPort *float64 `json:"backendPort,omitempty" tf:"backend_port,omitempty"`

	// The Protocol used for this NAT Rule. Possible values are Tcp and Udp.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type WindowsVirtualMachineInboundNATRuleObservation struct {

	// The Backend Port associated with this NAT Rule. Changing this forces a new resource to be created.
	BackendPort *float64 `json:"backendPort,omitempty" tf:"backend_port,omitempty"`

	// The frontend port associated with this Inbound NAT Rule.
	FrontendPort *float64 `json:"frontendPort,omitempty" tf:"frontend_port,omitempty"`

	// The Protocol used for this NAT Rule. Possible values are Tcp and Udp.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type WindowsVirtualMachineInboundNATRuleParameters struct {

	// The Backend Port associated with this NAT Rule. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	BackendPort *float64 `json:"backendPort" tf:"backend_port,omitempty"`

	// The Protocol used for this NAT Rule. Possible values are Tcp and Udp.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

type WindowsVirtualMachineInitParameters struct {

	// Can this Virtual Machine be claimed by users? Defaults to true.
	AllowClaim *bool `json:"allowClaim,omitempty" tf:"allow_claim,omitempty"`

	// Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
	DisallowPublicIPAddress *bool `json:"disallowPublicIpAddress,omitempty" tf:"disallow_public_ip_address,omitempty"`

	// A gallery_image_reference block as defined below.
	GalleryImageReference *WindowsVirtualMachineGalleryImageReferenceInitParameters `json:"galleryImageReference,omitempty" tf:"gallery_image_reference,omitempty"`

	// One or more inbound_nat_rule blocks as defined below. Changing this forces a new resource to be created.
	InboundNATRule []WindowsVirtualMachineInboundNATRuleInitParameters `json:"inboundNatRule,omitempty" tf:"inbound_nat_rule,omitempty"`

	// Specifies the name of the Dev Test Lab in which the Virtual Machine should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devtestlab/v1beta1.Lab
	LabName *string `json:"labName,omitempty" tf:"lab_name,omitempty"`

	// Reference to a Lab in devtestlab to populate labName.
	// +kubebuilder:validation:Optional
	LabNameRef *v1.Reference `json:"labNameRef,omitempty" tf:"-"`

	// Selector for a Lab in devtestlab to populate labName.
	// +kubebuilder:validation:Optional
	LabNameSelector *v1.Selector `json:"labNameSelector,omitempty" tf:"-"`

	// The name of a Subnet within the Dev Test Virtual Network where this machine should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	LabSubnetName *string `json:"labSubnetName,omitempty" tf:"lab_subnet_name,omitempty"`

	// Reference to a Subnet in network to populate labSubnetName.
	// +kubebuilder:validation:Optional
	LabSubnetNameRef *v1.Reference `json:"labSubnetNameRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate labSubnetName.
	// +kubebuilder:validation:Optional
	LabSubnetNameSelector *v1.Selector `json:"labSubnetNameSelector,omitempty" tf:"-"`

	// The ID of the Dev Test Virtual Network where this Virtual Machine should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devtestlab/v1beta2.VirtualNetwork
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	LabVirtualNetworkID *string `json:"labVirtualNetworkId,omitempty" tf:"lab_virtual_network_id,omitempty"`

	// Reference to a VirtualNetwork in devtestlab to populate labVirtualNetworkId.
	// +kubebuilder:validation:Optional
	LabVirtualNetworkIDRef *v1.Reference `json:"labVirtualNetworkIdRef,omitempty" tf:"-"`

	// Selector for a VirtualNetwork in devtestlab to populate labVirtualNetworkId.
	// +kubebuilder:validation:Optional
	LabVirtualNetworkIDSelector *v1.Selector `json:"labVirtualNetworkIdSelector,omitempty" tf:"-"`

	// Specifies the supported Azure location where the Dev Test Lab exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name of the Dev Test Machine. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Any notes about the Virtual Machine.
	Notes *string `json:"notes,omitempty" tf:"notes,omitempty"`

	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The Machine Size to use for this Virtual Machine, such as Standard_F2. Changing this forces a new resource to be created.
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// The type of Storage to use on this Virtual Machine. Possible values are Standard and Premium. Changing this forces a new resource to be created.
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type WindowsVirtualMachineObservation struct {

	// Can this Virtual Machine be claimed by users? Defaults to true.
	AllowClaim *bool `json:"allowClaim,omitempty" tf:"allow_claim,omitempty"`

	// Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
	DisallowPublicIPAddress *bool `json:"disallowPublicIpAddress,omitempty" tf:"disallow_public_ip_address,omitempty"`

	// The FQDN of the Virtual Machine.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// A gallery_image_reference block as defined below.
	GalleryImageReference *WindowsVirtualMachineGalleryImageReferenceObservation `json:"galleryImageReference,omitempty" tf:"gallery_image_reference,omitempty"`

	// The ID of the Virtual Machine.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more inbound_nat_rule blocks as defined below. Changing this forces a new resource to be created.
	InboundNATRule []WindowsVirtualMachineInboundNATRuleObservation `json:"inboundNatRule,omitempty" tf:"inbound_nat_rule,omitempty"`

	// Specifies the name of the Dev Test Lab in which the Virtual Machine should be created. Changing this forces a new resource to be created.
	LabName *string `json:"labName,omitempty" tf:"lab_name,omitempty"`

	// The name of a Subnet within the Dev Test Virtual Network where this machine should exist. Changing this forces a new resource to be created.
	LabSubnetName *string `json:"labSubnetName,omitempty" tf:"lab_subnet_name,omitempty"`

	// The ID of the Dev Test Virtual Network where this Virtual Machine should be created. Changing this forces a new resource to be created.
	LabVirtualNetworkID *string `json:"labVirtualNetworkId,omitempty" tf:"lab_virtual_network_id,omitempty"`

	// Specifies the supported Azure location where the Dev Test Lab exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name of the Dev Test Machine. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Any notes about the Virtual Machine.
	Notes *string `json:"notes,omitempty" tf:"notes,omitempty"`

	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The Machine Size to use for this Virtual Machine, such as Standard_F2. Changing this forces a new resource to be created.
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// The type of Storage to use on this Virtual Machine. Possible values are Standard and Premium. Changing this forces a new resource to be created.
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The unique immutable identifier of the Virtual Machine.
	UniqueIdentifier *string `json:"uniqueIdentifier,omitempty" tf:"unique_identifier,omitempty"`

	// The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type WindowsVirtualMachineParameters struct {

	// Can this Virtual Machine be claimed by users? Defaults to true.
	// +kubebuilder:validation:Optional
	AllowClaim *bool `json:"allowClaim,omitempty" tf:"allow_claim,omitempty"`

	// Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DisallowPublicIPAddress *bool `json:"disallowPublicIpAddress,omitempty" tf:"disallow_public_ip_address,omitempty"`

	// A gallery_image_reference block as defined below.
	// +kubebuilder:validation:Optional
	GalleryImageReference *WindowsVirtualMachineGalleryImageReferenceParameters `json:"galleryImageReference,omitempty" tf:"gallery_image_reference,omitempty"`

	// One or more inbound_nat_rule blocks as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	InboundNATRule []WindowsVirtualMachineInboundNATRuleParameters `json:"inboundNatRule,omitempty" tf:"inbound_nat_rule,omitempty"`

	// Specifies the name of the Dev Test Lab in which the Virtual Machine should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devtestlab/v1beta1.Lab
	// +kubebuilder:validation:Optional
	LabName *string `json:"labName,omitempty" tf:"lab_name,omitempty"`

	// Reference to a Lab in devtestlab to populate labName.
	// +kubebuilder:validation:Optional
	LabNameRef *v1.Reference `json:"labNameRef,omitempty" tf:"-"`

	// Selector for a Lab in devtestlab to populate labName.
	// +kubebuilder:validation:Optional
	LabNameSelector *v1.Selector `json:"labNameSelector,omitempty" tf:"-"`

	// The name of a Subnet within the Dev Test Virtual Network where this machine should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +kubebuilder:validation:Optional
	LabSubnetName *string `json:"labSubnetName,omitempty" tf:"lab_subnet_name,omitempty"`

	// Reference to a Subnet in network to populate labSubnetName.
	// +kubebuilder:validation:Optional
	LabSubnetNameRef *v1.Reference `json:"labSubnetNameRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate labSubnetName.
	// +kubebuilder:validation:Optional
	LabSubnetNameSelector *v1.Selector `json:"labSubnetNameSelector,omitempty" tf:"-"`

	// The ID of the Dev Test Virtual Network where this Virtual Machine should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devtestlab/v1beta2.VirtualNetwork
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LabVirtualNetworkID *string `json:"labVirtualNetworkId,omitempty" tf:"lab_virtual_network_id,omitempty"`

	// Reference to a VirtualNetwork in devtestlab to populate labVirtualNetworkId.
	// +kubebuilder:validation:Optional
	LabVirtualNetworkIDRef *v1.Reference `json:"labVirtualNetworkIdRef,omitempty" tf:"-"`

	// Selector for a VirtualNetwork in devtestlab to populate labVirtualNetworkId.
	// +kubebuilder:validation:Optional
	LabVirtualNetworkIDSelector *v1.Selector `json:"labVirtualNetworkIdSelector,omitempty" tf:"-"`

	// Specifies the supported Azure location where the Dev Test Lab exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name of the Dev Test Machine. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Any notes about the Virtual Machine.
	// +kubebuilder:validation:Optional
	Notes *string `json:"notes,omitempty" tf:"notes,omitempty"`

	// The Password associated with the username used to login to this Virtual Machine. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The Machine Size to use for this Virtual Machine, such as Standard_F2. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// The type of Storage to use on this Virtual Machine. Possible values are Standard and Premium. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// WindowsVirtualMachineSpec defines the desired state of WindowsVirtualMachine
type WindowsVirtualMachineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WindowsVirtualMachineParameters `json:"forProvider"`
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
	InitProvider WindowsVirtualMachineInitParameters `json:"initProvider,omitempty"`
}

// WindowsVirtualMachineStatus defines the observed state of WindowsVirtualMachine.
type WindowsVirtualMachineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WindowsVirtualMachineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// WindowsVirtualMachine is the Schema for the WindowsVirtualMachines API. Manages a Windows Virtual Machine within a Dev Test Lab.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type WindowsVirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.galleryImageReference) || (has(self.initProvider) && has(self.initProvider.galleryImageReference))",message="spec.forProvider.galleryImageReference is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.passwordSecretRef)",message="spec.forProvider.passwordSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.size) || (has(self.initProvider) && has(self.initProvider.size))",message="spec.forProvider.size is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storageType) || (has(self.initProvider) && has(self.initProvider.storageType))",message="spec.forProvider.storageType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.username) || (has(self.initProvider) && has(self.initProvider.username))",message="spec.forProvider.username is a required parameter"
	Spec   WindowsVirtualMachineSpec   `json:"spec"`
	Status WindowsVirtualMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WindowsVirtualMachineList contains a list of WindowsVirtualMachines
type WindowsVirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WindowsVirtualMachine `json:"items"`
}

// Repository type metadata.
var (
	WindowsVirtualMachine_Kind             = "WindowsVirtualMachine"
	WindowsVirtualMachine_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WindowsVirtualMachine_Kind}.String()
	WindowsVirtualMachine_KindAPIVersion   = WindowsVirtualMachine_Kind + "." + CRDGroupVersion.String()
	WindowsVirtualMachine_GroupVersionKind = CRDGroupVersion.WithKind(WindowsVirtualMachine_Kind)
)

func init() {
	SchemeBuilder.Register(&WindowsVirtualMachine{}, &WindowsVirtualMachineList{})
}
