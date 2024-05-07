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

type LogAnalyticsLinkedStorageAccountInitParameters struct {

	// The data source type which should be used for this Log Analytics Linked Storage Account. Possible values are CustomLogs, AzureWatson, Query, Ingestion and Alerts. Changing this forces a new Log Analytics Linked Storage Account to be created.
	DataSourceType *string `json:"dataSourceType,omitempty" tf:"data_source_type,omitempty"`

	// The name of the Resource Group where the Log Analytics Linked Storage Account should exist. Changing this forces a new Log Analytics Linked Storage Account to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The storage account resource ids to be linked.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +listType=set
	StorageAccountIds []*string `json:"storageAccountIds,omitempty" tf:"storage_account_ids,omitempty"`

	// References to Account in storage to populate storageAccountIds.
	// +kubebuilder:validation:Optional
	StorageAccountIdsRefs []v1.Reference `json:"storageAccountIdsRefs,omitempty" tf:"-"`

	// Selector for a list of Account in storage to populate storageAccountIds.
	// +kubebuilder:validation:Optional
	StorageAccountIdsSelector *v1.Selector `json:"storageAccountIdsSelector,omitempty" tf:"-"`

	// The resource ID of the Log Analytics Workspace. Changing this forces a new Log Analytics Linked Storage Account to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/operationalinsights/v1beta2.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	WorkspaceResourceID *string `json:"workspaceResourceId,omitempty" tf:"workspace_resource_id,omitempty"`

	// Reference to a Workspace in operationalinsights to populate workspaceResourceId.
	// +kubebuilder:validation:Optional
	WorkspaceResourceIDRef *v1.Reference `json:"workspaceResourceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in operationalinsights to populate workspaceResourceId.
	// +kubebuilder:validation:Optional
	WorkspaceResourceIDSelector *v1.Selector `json:"workspaceResourceIdSelector,omitempty" tf:"-"`
}

type LogAnalyticsLinkedStorageAccountObservation struct {

	// The data source type which should be used for this Log Analytics Linked Storage Account. Possible values are CustomLogs, AzureWatson, Query, Ingestion and Alerts. Changing this forces a new Log Analytics Linked Storage Account to be created.
	DataSourceType *string `json:"dataSourceType,omitempty" tf:"data_source_type,omitempty"`

	// The ID of the Log Analytics Linked Storage Account.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Resource Group where the Log Analytics Linked Storage Account should exist. Changing this forces a new Log Analytics Linked Storage Account to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The storage account resource ids to be linked.
	// +listType=set
	StorageAccountIds []*string `json:"storageAccountIds,omitempty" tf:"storage_account_ids,omitempty"`

	// The resource ID of the Log Analytics Workspace. Changing this forces a new Log Analytics Linked Storage Account to be created.
	WorkspaceResourceID *string `json:"workspaceResourceId,omitempty" tf:"workspace_resource_id,omitempty"`
}

type LogAnalyticsLinkedStorageAccountParameters struct {

	// The data source type which should be used for this Log Analytics Linked Storage Account. Possible values are CustomLogs, AzureWatson, Query, Ingestion and Alerts. Changing this forces a new Log Analytics Linked Storage Account to be created.
	// +kubebuilder:validation:Optional
	DataSourceType *string `json:"dataSourceType,omitempty" tf:"data_source_type,omitempty"`

	// The name of the Resource Group where the Log Analytics Linked Storage Account should exist. Changing this forces a new Log Analytics Linked Storage Account to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The storage account resource ids to be linked.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	// +listType=set
	StorageAccountIds []*string `json:"storageAccountIds,omitempty" tf:"storage_account_ids,omitempty"`

	// References to Account in storage to populate storageAccountIds.
	// +kubebuilder:validation:Optional
	StorageAccountIdsRefs []v1.Reference `json:"storageAccountIdsRefs,omitempty" tf:"-"`

	// Selector for a list of Account in storage to populate storageAccountIds.
	// +kubebuilder:validation:Optional
	StorageAccountIdsSelector *v1.Selector `json:"storageAccountIdsSelector,omitempty" tf:"-"`

	// The resource ID of the Log Analytics Workspace. Changing this forces a new Log Analytics Linked Storage Account to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/operationalinsights/v1beta2.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	WorkspaceResourceID *string `json:"workspaceResourceId,omitempty" tf:"workspace_resource_id,omitempty"`

	// Reference to a Workspace in operationalinsights to populate workspaceResourceId.
	// +kubebuilder:validation:Optional
	WorkspaceResourceIDRef *v1.Reference `json:"workspaceResourceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in operationalinsights to populate workspaceResourceId.
	// +kubebuilder:validation:Optional
	WorkspaceResourceIDSelector *v1.Selector `json:"workspaceResourceIdSelector,omitempty" tf:"-"`
}

// LogAnalyticsLinkedStorageAccountSpec defines the desired state of LogAnalyticsLinkedStorageAccount
type LogAnalyticsLinkedStorageAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogAnalyticsLinkedStorageAccountParameters `json:"forProvider"`
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
	InitProvider LogAnalyticsLinkedStorageAccountInitParameters `json:"initProvider,omitempty"`
}

// LogAnalyticsLinkedStorageAccountStatus defines the observed state of LogAnalyticsLinkedStorageAccount.
type LogAnalyticsLinkedStorageAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogAnalyticsLinkedStorageAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LogAnalyticsLinkedStorageAccount is the Schema for the LogAnalyticsLinkedStorageAccounts API. Manages a Log Analytics Linked Storage Account.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LogAnalyticsLinkedStorageAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dataSourceType) || (has(self.initProvider) && has(self.initProvider.dataSourceType))",message="spec.forProvider.dataSourceType is a required parameter"
	Spec   LogAnalyticsLinkedStorageAccountSpec   `json:"spec"`
	Status LogAnalyticsLinkedStorageAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsLinkedStorageAccountList contains a list of LogAnalyticsLinkedStorageAccounts
type LogAnalyticsLinkedStorageAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogAnalyticsLinkedStorageAccount `json:"items"`
}

// Repository type metadata.
var (
	LogAnalyticsLinkedStorageAccount_Kind             = "LogAnalyticsLinkedStorageAccount"
	LogAnalyticsLinkedStorageAccount_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LogAnalyticsLinkedStorageAccount_Kind}.String()
	LogAnalyticsLinkedStorageAccount_KindAPIVersion   = LogAnalyticsLinkedStorageAccount_Kind + "." + CRDGroupVersion.String()
	LogAnalyticsLinkedStorageAccount_GroupVersionKind = CRDGroupVersion.WithKind(LogAnalyticsLinkedStorageAccount_Kind)
)

func init() {
	SchemeBuilder.Register(&LogAnalyticsLinkedStorageAccount{}, &LogAnalyticsLinkedStorageAccountList{})
}
