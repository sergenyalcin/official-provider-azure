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

type ContentLinkHashInitParameters struct {

	// Specifies the hash algorithm used to hash the content.
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// Specifies the expected hash value of the content.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ContentLinkHashObservation struct {

	// Specifies the hash algorithm used to hash the content.
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// Specifies the expected hash value of the content.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ContentLinkHashParameters struct {

	// Specifies the hash algorithm used to hash the content.
	// +kubebuilder:validation:Optional
	Algorithm *string `json:"algorithm" tf:"algorithm,omitempty"`

	// Specifies the expected hash value of the content.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type ContentLinkInitParameters struct {

	// A hash block as defined below.
	Hash *ContentLinkHashInitParameters `json:"hash,omitempty" tf:"hash,omitempty"`

	// The URI of the runbook content.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`

	// Specifies the version of the content
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ContentLinkObservation struct {

	// A hash block as defined below.
	Hash *ContentLinkHashObservation `json:"hash,omitempty" tf:"hash,omitempty"`

	// The URI of the runbook content.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`

	// Specifies the version of the content
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ContentLinkParameters struct {

	// A hash block as defined below.
	// +kubebuilder:validation:Optional
	Hash *ContentLinkHashParameters `json:"hash,omitempty" tf:"hash,omitempty"`

	// The URI of the runbook content.
	// +kubebuilder:validation:Optional
	URI *string `json:"uri" tf:"uri,omitempty"`

	// Specifies the version of the content
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type DraftInitParameters struct {

	// A publish_content_link block as defined above.
	ContentLink *ContentLinkInitParameters `json:"contentLink,omitempty" tf:"content_link,omitempty"`

	// Whether the draft in edit mode.
	EditModeEnabled *bool `json:"editModeEnabled,omitempty" tf:"edit_mode_enabled,omitempty"`

	// Specifies the output types of the runbook.
	OutputTypes []*string `json:"outputTypes,omitempty" tf:"output_types,omitempty"`

	// A list of parameters block as defined below.
	Parameters []ParametersInitParameters `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type DraftObservation struct {

	// A publish_content_link block as defined above.
	ContentLink *ContentLinkObservation `json:"contentLink,omitempty" tf:"content_link,omitempty"`

	CreationTime *string `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	// Whether the draft in edit mode.
	EditModeEnabled *bool `json:"editModeEnabled,omitempty" tf:"edit_mode_enabled,omitempty"`

	LastModifiedTime *string `json:"lastModifiedTime,omitempty" tf:"last_modified_time,omitempty"`

	// Specifies the output types of the runbook.
	OutputTypes []*string `json:"outputTypes,omitempty" tf:"output_types,omitempty"`

	// A list of parameters block as defined below.
	Parameters []ParametersObservation `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type DraftParameters struct {

	// A publish_content_link block as defined above.
	// +kubebuilder:validation:Optional
	ContentLink *ContentLinkParameters `json:"contentLink,omitempty" tf:"content_link,omitempty"`

	// Whether the draft in edit mode.
	// +kubebuilder:validation:Optional
	EditModeEnabled *bool `json:"editModeEnabled,omitempty" tf:"edit_mode_enabled,omitempty"`

	// Specifies the output types of the runbook.
	// +kubebuilder:validation:Optional
	OutputTypes []*string `json:"outputTypes,omitempty" tf:"output_types,omitempty"`

	// A list of parameters block as defined below.
	// +kubebuilder:validation:Optional
	Parameters []ParametersParameters `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type JobScheduleInitParameters struct {

	// The UUID of automation runbook job schedule ID.
	JobScheduleID *string `json:"jobScheduleId,omitempty" tf:"job_schedule_id"`

	// A map of key/value pairs corresponding to the arguments that can be passed to the Runbook.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters"`

	// Name of a Hybrid Worker Group the Runbook will be executed on.
	RunOn *string `json:"runOn,omitempty" tf:"run_on"`

	// The name of the Schedule.
	ScheduleName *string `json:"scheduleName,omitempty" tf:"schedule_name"`
}

type JobScheduleObservation struct {

	// The UUID of automation runbook job schedule ID.
	JobScheduleID *string `json:"jobScheduleId,omitempty" tf:"job_schedule_id,omitempty"`

	// A map of key/value pairs corresponding to the arguments that can be passed to the Runbook.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Name of a Hybrid Worker Group the Runbook will be executed on.
	RunOn *string `json:"runOn,omitempty" tf:"run_on,omitempty"`

	// The name of the Schedule.
	ScheduleName *string `json:"scheduleName,omitempty" tf:"schedule_name,omitempty"`
}

type JobScheduleParameters struct {

	// The UUID of automation runbook job schedule ID.
	// +kubebuilder:validation:Optional
	JobScheduleID *string `json:"jobScheduleId,omitempty" tf:"job_schedule_id"`

	// A map of key/value pairs corresponding to the arguments that can be passed to the Runbook.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters"`

	// Name of a Hybrid Worker Group the Runbook will be executed on.
	// +kubebuilder:validation:Optional
	RunOn *string `json:"runOn,omitempty" tf:"run_on"`

	// The name of the Schedule.
	// +kubebuilder:validation:Optional
	ScheduleName *string `json:"scheduleName,omitempty" tf:"schedule_name"`
}

type ParametersInitParameters struct {

	// Specifies the default value of the parameter.
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// The name of the parameter.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Whether this parameter is mandatory.
	Mandatory *bool `json:"mandatory,omitempty" tf:"mandatory,omitempty"`

	// Specifies the position of the parameter.
	Position *float64 `json:"position,omitempty" tf:"position,omitempty"`

	// Specifies the type of this parameter.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ParametersObservation struct {

	// Specifies the default value of the parameter.
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// The name of the parameter.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Whether this parameter is mandatory.
	Mandatory *bool `json:"mandatory,omitempty" tf:"mandatory,omitempty"`

	// Specifies the position of the parameter.
	Position *float64 `json:"position,omitempty" tf:"position,omitempty"`

	// Specifies the type of this parameter.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ParametersParameters struct {

	// Specifies the default value of the parameter.
	// +kubebuilder:validation:Optional
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// The name of the parameter.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// Whether this parameter is mandatory.
	// +kubebuilder:validation:Optional
	Mandatory *bool `json:"mandatory,omitempty" tf:"mandatory,omitempty"`

	// Specifies the position of the parameter.
	// +kubebuilder:validation:Optional
	Position *float64 `json:"position,omitempty" tf:"position,omitempty"`

	// Specifies the type of this parameter.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type PublishContentLinkHashInitParameters struct {

	// Specifies the hash algorithm used to hash the content.
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// Specifies the expected hash value of the content.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PublishContentLinkHashObservation struct {

	// Specifies the hash algorithm used to hash the content.
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// Specifies the expected hash value of the content.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PublishContentLinkHashParameters struct {

	// Specifies the hash algorithm used to hash the content.
	// +kubebuilder:validation:Optional
	Algorithm *string `json:"algorithm" tf:"algorithm,omitempty"`

	// Specifies the expected hash value of the content.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type PublishContentLinkInitParameters struct {

	// A hash block as defined below.
	Hash *PublishContentLinkHashInitParameters `json:"hash,omitempty" tf:"hash,omitempty"`

	// The URI of the runbook content.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`

	// Specifies the version of the content
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type PublishContentLinkObservation struct {

	// A hash block as defined below.
	Hash *PublishContentLinkHashObservation `json:"hash,omitempty" tf:"hash,omitempty"`

	// The URI of the runbook content.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`

	// Specifies the version of the content
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type PublishContentLinkParameters struct {

	// A hash block as defined below.
	// +kubebuilder:validation:Optional
	Hash *PublishContentLinkHashParameters `json:"hash,omitempty" tf:"hash,omitempty"`

	// The URI of the runbook content.
	// +kubebuilder:validation:Optional
	URI *string `json:"uri" tf:"uri,omitempty"`

	// Specifies the version of the content
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type RunBookInitParameters struct {

	// The name of the automation account in which the Runbook is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/automation/v1beta2.Account
	AutomationAccountName *string `json:"automationAccountName,omitempty" tf:"automation_account_name,omitempty"`

	// Reference to a Account in automation to populate automationAccountName.
	// +kubebuilder:validation:Optional
	AutomationAccountNameRef *v1.Reference `json:"automationAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account in automation to populate automationAccountName.
	// +kubebuilder:validation:Optional
	AutomationAccountNameSelector *v1.Selector `json:"automationAccountNameSelector,omitempty" tf:"-"`

	// The desired content of the runbook.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// A description for this credential.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A draft block as defined below.
	Draft *DraftInitParameters `json:"draft,omitempty" tf:"draft,omitempty"`

	// One or more job_schedule block as defined below.
	JobSchedule []JobScheduleInitParameters `json:"jobSchedule,omitempty" tf:"job_schedule,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the activity-level tracing options of the runbook, available only for Graphical runbooks. Possible values are 0 for None, 9 for Basic, and 15 for Detailed. Must turn on Verbose logging in order to see the tracing.
	LogActivityTraceLevel *float64 `json:"logActivityTraceLevel,omitempty" tf:"log_activity_trace_level,omitempty"`

	// Progress log option.
	LogProgress *bool `json:"logProgress,omitempty" tf:"log_progress,omitempty"`

	// Verbose log option.
	LogVerbose *bool `json:"logVerbose,omitempty" tf:"log_verbose,omitempty"`

	// Specifies the name of the Runbook. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One publish_content_link block as defined below.
	PublishContentLink *PublishContentLinkInitParameters `json:"publishContentLink,omitempty" tf:"publish_content_link,omitempty"`

	// The name of the resource group in which the Runbook is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The type of the runbook - can be either Graph, GraphPowerShell, GraphPowerShellWorkflow, PowerShellWorkflow, PowerShell, PowerShell72, Python3, Python2 or Script. Changing this forces a new resource to be created.
	RunBookType *string `json:"runbookType,omitempty" tf:"runbook_type,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RunBookObservation struct {

	// The name of the automation account in which the Runbook is created. Changing this forces a new resource to be created.
	AutomationAccountName *string `json:"automationAccountName,omitempty" tf:"automation_account_name,omitempty"`

	// The desired content of the runbook.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// A description for this credential.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A draft block as defined below.
	Draft *DraftObservation `json:"draft,omitempty" tf:"draft,omitempty"`

	// The Automation Runbook ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more job_schedule block as defined below.
	JobSchedule []JobScheduleObservation `json:"jobSchedule,omitempty" tf:"job_schedule,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the activity-level tracing options of the runbook, available only for Graphical runbooks. Possible values are 0 for None, 9 for Basic, and 15 for Detailed. Must turn on Verbose logging in order to see the tracing.
	LogActivityTraceLevel *float64 `json:"logActivityTraceLevel,omitempty" tf:"log_activity_trace_level,omitempty"`

	// Progress log option.
	LogProgress *bool `json:"logProgress,omitempty" tf:"log_progress,omitempty"`

	// Verbose log option.
	LogVerbose *bool `json:"logVerbose,omitempty" tf:"log_verbose,omitempty"`

	// Specifies the name of the Runbook. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One publish_content_link block as defined below.
	PublishContentLink *PublishContentLinkObservation `json:"publishContentLink,omitempty" tf:"publish_content_link,omitempty"`

	// The name of the resource group in which the Runbook is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The type of the runbook - can be either Graph, GraphPowerShell, GraphPowerShellWorkflow, PowerShellWorkflow, PowerShell, PowerShell72, Python3, Python2 or Script. Changing this forces a new resource to be created.
	RunBookType *string `json:"runbookType,omitempty" tf:"runbook_type,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RunBookParameters struct {

	// The name of the automation account in which the Runbook is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/automation/v1beta2.Account
	// +kubebuilder:validation:Optional
	AutomationAccountName *string `json:"automationAccountName,omitempty" tf:"automation_account_name,omitempty"`

	// Reference to a Account in automation to populate automationAccountName.
	// +kubebuilder:validation:Optional
	AutomationAccountNameRef *v1.Reference `json:"automationAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account in automation to populate automationAccountName.
	// +kubebuilder:validation:Optional
	AutomationAccountNameSelector *v1.Selector `json:"automationAccountNameSelector,omitempty" tf:"-"`

	// The desired content of the runbook.
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// A description for this credential.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A draft block as defined below.
	// +kubebuilder:validation:Optional
	Draft *DraftParameters `json:"draft,omitempty" tf:"draft,omitempty"`

	// One or more job_schedule block as defined below.
	// +kubebuilder:validation:Optional
	JobSchedule []JobScheduleParameters `json:"jobSchedule,omitempty" tf:"job_schedule,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the activity-level tracing options of the runbook, available only for Graphical runbooks. Possible values are 0 for None, 9 for Basic, and 15 for Detailed. Must turn on Verbose logging in order to see the tracing.
	// +kubebuilder:validation:Optional
	LogActivityTraceLevel *float64 `json:"logActivityTraceLevel,omitempty" tf:"log_activity_trace_level,omitempty"`

	// Progress log option.
	// +kubebuilder:validation:Optional
	LogProgress *bool `json:"logProgress,omitempty" tf:"log_progress,omitempty"`

	// Verbose log option.
	// +kubebuilder:validation:Optional
	LogVerbose *bool `json:"logVerbose,omitempty" tf:"log_verbose,omitempty"`

	// Specifies the name of the Runbook. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One publish_content_link block as defined below.
	// +kubebuilder:validation:Optional
	PublishContentLink *PublishContentLinkParameters `json:"publishContentLink,omitempty" tf:"publish_content_link,omitempty"`

	// The name of the resource group in which the Runbook is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The type of the runbook - can be either Graph, GraphPowerShell, GraphPowerShellWorkflow, PowerShellWorkflow, PowerShell, PowerShell72, Python3, Python2 or Script. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	RunBookType *string `json:"runbookType,omitempty" tf:"runbook_type,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// RunBookSpec defines the desired state of RunBook
type RunBookSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RunBookParameters `json:"forProvider"`
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
	InitProvider RunBookInitParameters `json:"initProvider,omitempty"`
}

// RunBookStatus defines the observed state of RunBook.
type RunBookStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RunBookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// RunBook is the Schema for the RunBooks API. Manages a Automation Runbook.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type RunBook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.logProgress) || (has(self.initProvider) && has(self.initProvider.logProgress))",message="spec.forProvider.logProgress is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.logVerbose) || (has(self.initProvider) && has(self.initProvider.logVerbose))",message="spec.forProvider.logVerbose is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.runbookType) || (has(self.initProvider) && has(self.initProvider.runbookType))",message="spec.forProvider.runbookType is a required parameter"
	Spec   RunBookSpec   `json:"spec"`
	Status RunBookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RunBookList contains a list of RunBooks
type RunBookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RunBook `json:"items"`
}

// Repository type metadata.
var (
	RunBook_Kind             = "RunBook"
	RunBook_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RunBook_Kind}.String()
	RunBook_KindAPIVersion   = RunBook_Kind + "." + CRDGroupVersion.String()
	RunBook_GroupVersionKind = CRDGroupVersion.WithKind(RunBook_Kind)
)

func init() {
	SchemeBuilder.Register(&RunBook{}, &RunBookList{})
}
