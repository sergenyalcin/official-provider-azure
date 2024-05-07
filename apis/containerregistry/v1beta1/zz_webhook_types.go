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

type WebhookInitParameters struct {

	// A list of actions that trigger the Webhook to post notifications. At least one action needs to be specified. Valid values are: push, delete, quarantine, chart_push, chart_delete
	// +listType=set
	Actions []*string `json:"actions,omitempty" tf:"actions,omitempty"`

	// Custom headers that will be added to the webhook notifications request.
	// +mapType=granular
	CustomHeaders map[string]*string `json:"customHeaders,omitempty" tf:"custom_headers,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The Name of Container registry this Webhook belongs to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/containerregistry/v1beta2.Registry
	RegistryName *string `json:"registryName,omitempty" tf:"registry_name,omitempty"`

	// Reference to a Registry in containerregistry to populate registryName.
	// +kubebuilder:validation:Optional
	RegistryNameRef *v1.Reference `json:"registryNameRef,omitempty" tf:"-"`

	// Selector for a Registry in containerregistry to populate registryName.
	// +kubebuilder:validation:Optional
	RegistryNameSelector *v1.Selector `json:"registryNameSelector,omitempty" tf:"-"`

	// Specifies the scope of repositories that can trigger an event. For example, foo:* means events for all tags under repository foo. foo:bar means events for 'foo:bar' only. foo is equivalent to foo:latest. Empty means all events. Defaults to "".
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Specifies the service URI for the Webhook to post notifications.
	ServiceURI *string `json:"serviceUri,omitempty" tf:"service_uri,omitempty"`

	// Specifies if this Webhook triggers notifications or not. Valid values: enabled and disabled. Default is enabled.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type WebhookObservation struct {

	// A list of actions that trigger the Webhook to post notifications. At least one action needs to be specified. Valid values are: push, delete, quarantine, chart_push, chart_delete
	// +listType=set
	Actions []*string `json:"actions,omitempty" tf:"actions,omitempty"`

	// Custom headers that will be added to the webhook notifications request.
	// +mapType=granular
	CustomHeaders map[string]*string `json:"customHeaders,omitempty" tf:"custom_headers,omitempty"`

	// The ID of the Container Registry Webhook.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The Name of Container registry this Webhook belongs to. Changing this forces a new resource to be created.
	RegistryName *string `json:"registryName,omitempty" tf:"registry_name,omitempty"`

	// The name of the resource group in which to create the Container Registry Webhook. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Specifies the scope of repositories that can trigger an event. For example, foo:* means events for all tags under repository foo. foo:bar means events for 'foo:bar' only. foo is equivalent to foo:latest. Empty means all events. Defaults to "".
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Specifies the service URI for the Webhook to post notifications.
	ServiceURI *string `json:"serviceUri,omitempty" tf:"service_uri,omitempty"`

	// Specifies if this Webhook triggers notifications or not. Valid values: enabled and disabled. Default is enabled.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type WebhookParameters struct {

	// A list of actions that trigger the Webhook to post notifications. At least one action needs to be specified. Valid values are: push, delete, quarantine, chart_push, chart_delete
	// +kubebuilder:validation:Optional
	// +listType=set
	Actions []*string `json:"actions,omitempty" tf:"actions,omitempty"`

	// Custom headers that will be added to the webhook notifications request.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	CustomHeaders map[string]*string `json:"customHeaders,omitempty" tf:"custom_headers,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The Name of Container registry this Webhook belongs to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/containerregistry/v1beta2.Registry
	// +kubebuilder:validation:Optional
	RegistryName *string `json:"registryName,omitempty" tf:"registry_name,omitempty"`

	// Reference to a Registry in containerregistry to populate registryName.
	// +kubebuilder:validation:Optional
	RegistryNameRef *v1.Reference `json:"registryNameRef,omitempty" tf:"-"`

	// Selector for a Registry in containerregistry to populate registryName.
	// +kubebuilder:validation:Optional
	RegistryNameSelector *v1.Selector `json:"registryNameSelector,omitempty" tf:"-"`

	// The name of the resource group in which to create the Container Registry Webhook. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the scope of repositories that can trigger an event. For example, foo:* means events for all tags under repository foo. foo:bar means events for 'foo:bar' only. foo is equivalent to foo:latest. Empty means all events. Defaults to "".
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Specifies the service URI for the Webhook to post notifications.
	// +kubebuilder:validation:Optional
	ServiceURI *string `json:"serviceUri,omitempty" tf:"service_uri,omitempty"`

	// Specifies if this Webhook triggers notifications or not. Valid values: enabled and disabled. Default is enabled.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// WebhookSpec defines the desired state of Webhook
type WebhookSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WebhookParameters `json:"forProvider"`
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
	InitProvider WebhookInitParameters `json:"initProvider,omitempty"`
}

// WebhookStatus defines the observed state of Webhook.
type WebhookStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WebhookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Webhook is the Schema for the Webhooks API. Manages an Azure Container Registry Webhook.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Webhook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.actions) || (has(self.initProvider) && has(self.initProvider.actions))",message="spec.forProvider.actions is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceUri) || (has(self.initProvider) && has(self.initProvider.serviceUri))",message="spec.forProvider.serviceUri is a required parameter"
	Spec   WebhookSpec   `json:"spec"`
	Status WebhookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebhookList contains a list of Webhooks
type WebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Webhook `json:"items"`
}

// Repository type metadata.
var (
	Webhook_Kind             = "Webhook"
	Webhook_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Webhook_Kind}.String()
	Webhook_KindAPIVersion   = Webhook_Kind + "." + CRDGroupVersion.String()
	Webhook_GroupVersionKind = CRDGroupVersion.WithKind(Webhook_Kind)
)

func init() {
	SchemeBuilder.Register(&Webhook{}, &WebhookList{})
}
