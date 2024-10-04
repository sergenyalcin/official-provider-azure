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

type CorsInitParameters struct {

	// A list of origins which should be able to make cross-origin calls. * can be used to allow all calls.
	// +listType=set
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`
}

type CorsObservation struct {

	// A list of origins which should be able to make cross-origin calls. * can be used to allow all calls.
	// +listType=set
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`
}

type CorsParameters struct {

	// A list of origins which should be able to make cross-origin calls. * can be used to allow all calls.
	// +kubebuilder:validation:Optional
	// +listType=set
	AllowedOrigins []*string `json:"allowedOrigins" tf:"allowed_origins,omitempty"`
}

type IdentityInitParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this signalR.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this signalR. Possible values are SystemAssigned, UserAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this signalR.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The ID of the SignalR service.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The ID of the SignalR service.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this signalR. Possible values are SystemAssigned, UserAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this signalR.
	// +kubebuilder:validation:Optional
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this signalR. Possible values are SystemAssigned, UserAssigned.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type LiveTraceInitParameters struct {

	// Whether the log category ConnectivityLogs is enabled? Defaults to true
	ConnectivityLogsEnabled *bool `json:"connectivityLogsEnabled,omitempty" tf:"connectivity_logs_enabled,omitempty"`

	// Whether the live trace is enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Whether the log category HttpRequestLogs is enabled? Defaults to true
	HTTPRequestLogsEnabled *bool `json:"httpRequestLogsEnabled,omitempty" tf:"http_request_logs_enabled,omitempty"`

	// Whether the log category MessagingLogs is enabled? Defaults to true
	MessagingLogsEnabled *bool `json:"messagingLogsEnabled,omitempty" tf:"messaging_logs_enabled,omitempty"`
}

type LiveTraceObservation struct {

	// Whether the log category ConnectivityLogs is enabled? Defaults to true
	ConnectivityLogsEnabled *bool `json:"connectivityLogsEnabled,omitempty" tf:"connectivity_logs_enabled,omitempty"`

	// Whether the live trace is enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Whether the log category HttpRequestLogs is enabled? Defaults to true
	HTTPRequestLogsEnabled *bool `json:"httpRequestLogsEnabled,omitempty" tf:"http_request_logs_enabled,omitempty"`

	// Whether the log category MessagingLogs is enabled? Defaults to true
	MessagingLogsEnabled *bool `json:"messagingLogsEnabled,omitempty" tf:"messaging_logs_enabled,omitempty"`
}

type LiveTraceParameters struct {

	// Whether the log category ConnectivityLogs is enabled? Defaults to true
	// +kubebuilder:validation:Optional
	ConnectivityLogsEnabled *bool `json:"connectivityLogsEnabled,omitempty" tf:"connectivity_logs_enabled,omitempty"`

	// Whether the live trace is enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Whether the log category HttpRequestLogs is enabled? Defaults to true
	// +kubebuilder:validation:Optional
	HTTPRequestLogsEnabled *bool `json:"httpRequestLogsEnabled,omitempty" tf:"http_request_logs_enabled,omitempty"`

	// Whether the log category MessagingLogs is enabled? Defaults to true
	// +kubebuilder:validation:Optional
	MessagingLogsEnabled *bool `json:"messagingLogsEnabled,omitempty" tf:"messaging_logs_enabled,omitempty"`
}

type ServiceInitParameters struct {

	// Whether to enable AAD auth? Defaults to true.
	AADAuthEnabled *bool `json:"aadAuthEnabled,omitempty" tf:"aad_auth_enabled,omitempty"`

	// Specifies if Connectivity Logs are enabled or not. Defaults to false.
	ConnectivityLogsEnabled *bool `json:"connectivityLogsEnabled,omitempty" tf:"connectivity_logs_enabled,omitempty"`

	// A cors block as documented below.
	Cors []CorsInitParameters `json:"cors,omitempty" tf:"cors,omitempty"`

	// Specifies if Http Request Logs are enabled or not. Defaults to false.
	HTTPRequestLogsEnabled *bool `json:"httpRequestLogsEnabled,omitempty" tf:"http_request_logs_enabled,omitempty"`

	// An identity block as defined below.
	Identity *IdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// A live_trace block as defined below.
	LiveTrace *LiveTraceInitParameters `json:"liveTrace,omitempty" tf:"live_trace,omitempty"`

	// Specifies if Live Trace is enabled or not. Defaults to false.
	LiveTraceEnabled *bool `json:"liveTraceEnabled,omitempty" tf:"live_trace_enabled,omitempty"`

	// Whether to enable local auth? Defaults to true.
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// Specifies the supported Azure location where the SignalR service exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies if Messaging Logs are enabled or not. Defaults to false.
	MessagingLogsEnabled *bool `json:"messagingLogsEnabled,omitempty" tf:"messaging_logs_enabled,omitempty"`

	// Whether to enable public network access? Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// Specifies the client connection timeout. Defaults to 30.
	ServerlessConnectionTimeoutInSeconds *float64 `json:"serverlessConnectionTimeoutInSeconds,omitempty" tf:"serverless_connection_timeout_in_seconds,omitempty"`

	// Specifies the service mode. Possible values are Classic, Default and Serverless. Defaults to Default.
	ServiceMode *string `json:"serviceMode,omitempty" tf:"service_mode,omitempty"`

	// A sku block as documented below.
	Sku *SkuInitParameters `json:"sku,omitempty" tf:"sku,omitempty"`

	// Whether to request client certificate during TLS handshake? Defaults to false.
	TLSClientCertEnabled *bool `json:"tlsClientCertEnabled,omitempty" tf:"tls_client_cert_enabled,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// An upstream_endpoint block as documented below. Using this block requires the SignalR service to be Serverless. When creating multiple blocks they will be processed in the order they are defined in.
	UpstreamEndpoint []UpstreamEndpointInitParameters `json:"upstreamEndpoint,omitempty" tf:"upstream_endpoint,omitempty"`
}

type ServiceObservation struct {

	// Whether to enable AAD auth? Defaults to true.
	AADAuthEnabled *bool `json:"aadAuthEnabled,omitempty" tf:"aad_auth_enabled,omitempty"`

	// Specifies if Connectivity Logs are enabled or not. Defaults to false.
	ConnectivityLogsEnabled *bool `json:"connectivityLogsEnabled,omitempty" tf:"connectivity_logs_enabled,omitempty"`

	// A cors block as documented below.
	Cors []CorsObservation `json:"cors,omitempty" tf:"cors,omitempty"`

	// Specifies if Http Request Logs are enabled or not. Defaults to false.
	HTTPRequestLogsEnabled *bool `json:"httpRequestLogsEnabled,omitempty" tf:"http_request_logs_enabled,omitempty"`

	// The FQDN of the SignalR service.
	HostName *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// The ID of the SignalR service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The publicly accessible IP of the SignalR service.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// An identity block as defined below.
	Identity *IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// A live_trace block as defined below.
	LiveTrace *LiveTraceObservation `json:"liveTrace,omitempty" tf:"live_trace,omitempty"`

	// Specifies if Live Trace is enabled or not. Defaults to false.
	LiveTraceEnabled *bool `json:"liveTraceEnabled,omitempty" tf:"live_trace_enabled,omitempty"`

	// Whether to enable local auth? Defaults to true.
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// Specifies the supported Azure location where the SignalR service exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies if Messaging Logs are enabled or not. Defaults to false.
	MessagingLogsEnabled *bool `json:"messagingLogsEnabled,omitempty" tf:"messaging_logs_enabled,omitempty"`

	// Whether to enable public network access? Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The publicly accessible port of the SignalR service which is designed for browser/client use.
	PublicPort *float64 `json:"publicPort,omitempty" tf:"public_port,omitempty"`

	// The name of the resource group in which to create the SignalR service. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The publicly accessible port of the SignalR service which is designed for customer server side use.
	ServerPort *float64 `json:"serverPort,omitempty" tf:"server_port,omitempty"`

	// Specifies the client connection timeout. Defaults to 30.
	ServerlessConnectionTimeoutInSeconds *float64 `json:"serverlessConnectionTimeoutInSeconds,omitempty" tf:"serverless_connection_timeout_in_seconds,omitempty"`

	// Specifies the service mode. Possible values are Classic, Default and Serverless. Defaults to Default.
	ServiceMode *string `json:"serviceMode,omitempty" tf:"service_mode,omitempty"`

	// A sku block as documented below.
	Sku *SkuObservation `json:"sku,omitempty" tf:"sku,omitempty"`

	// Whether to request client certificate during TLS handshake? Defaults to false.
	TLSClientCertEnabled *bool `json:"tlsClientCertEnabled,omitempty" tf:"tls_client_cert_enabled,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// An upstream_endpoint block as documented below. Using this block requires the SignalR service to be Serverless. When creating multiple blocks they will be processed in the order they are defined in.
	UpstreamEndpoint []UpstreamEndpointObservation `json:"upstreamEndpoint,omitempty" tf:"upstream_endpoint,omitempty"`
}

type ServiceParameters struct {

	// Whether to enable AAD auth? Defaults to true.
	// +kubebuilder:validation:Optional
	AADAuthEnabled *bool `json:"aadAuthEnabled,omitempty" tf:"aad_auth_enabled,omitempty"`

	// Specifies if Connectivity Logs are enabled or not. Defaults to false.
	// +kubebuilder:validation:Optional
	ConnectivityLogsEnabled *bool `json:"connectivityLogsEnabled,omitempty" tf:"connectivity_logs_enabled,omitempty"`

	// A cors block as documented below.
	// +kubebuilder:validation:Optional
	Cors []CorsParameters `json:"cors,omitempty" tf:"cors,omitempty"`

	// Specifies if Http Request Logs are enabled or not. Defaults to false.
	// +kubebuilder:validation:Optional
	HTTPRequestLogsEnabled *bool `json:"httpRequestLogsEnabled,omitempty" tf:"http_request_logs_enabled,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity *IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// A live_trace block as defined below.
	// +kubebuilder:validation:Optional
	LiveTrace *LiveTraceParameters `json:"liveTrace,omitempty" tf:"live_trace,omitempty"`

	// Specifies if Live Trace is enabled or not. Defaults to false.
	// +kubebuilder:validation:Optional
	LiveTraceEnabled *bool `json:"liveTraceEnabled,omitempty" tf:"live_trace_enabled,omitempty"`

	// Whether to enable local auth? Defaults to true.
	// +kubebuilder:validation:Optional
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// Specifies the supported Azure location where the SignalR service exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies if Messaging Logs are enabled or not. Defaults to false.
	// +kubebuilder:validation:Optional
	MessagingLogsEnabled *bool `json:"messagingLogsEnabled,omitempty" tf:"messaging_logs_enabled,omitempty"`

	// Whether to enable public network access? Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the resource group in which to create the SignalR service. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the client connection timeout. Defaults to 30.
	// +kubebuilder:validation:Optional
	ServerlessConnectionTimeoutInSeconds *float64 `json:"serverlessConnectionTimeoutInSeconds,omitempty" tf:"serverless_connection_timeout_in_seconds,omitempty"`

	// Specifies the service mode. Possible values are Classic, Default and Serverless. Defaults to Default.
	// +kubebuilder:validation:Optional
	ServiceMode *string `json:"serviceMode,omitempty" tf:"service_mode,omitempty"`

	// A sku block as documented below.
	// +kubebuilder:validation:Optional
	Sku *SkuParameters `json:"sku,omitempty" tf:"sku,omitempty"`

	// Whether to request client certificate during TLS handshake? Defaults to false.
	// +kubebuilder:validation:Optional
	TLSClientCertEnabled *bool `json:"tlsClientCertEnabled,omitempty" tf:"tls_client_cert_enabled,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// An upstream_endpoint block as documented below. Using this block requires the SignalR service to be Serverless. When creating multiple blocks they will be processed in the order they are defined in.
	// +kubebuilder:validation:Optional
	UpstreamEndpoint []UpstreamEndpointParameters `json:"upstreamEndpoint,omitempty" tf:"upstream_endpoint,omitempty"`
}

type SkuInitParameters struct {

	// Specifies the number of units associated with this SignalR service. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 200, 300, 400, 500, 600, 700, 800, 900 and 1000.
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// Specifies which tier to use. Valid values are Free_F1, Standard_S1, Premium_P1 and Premium_P2.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SkuObservation struct {

	// Specifies the number of units associated with this SignalR service. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 200, 300, 400, 500, 600, 700, 800, 900 and 1000.
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// Specifies which tier to use. Valid values are Free_F1, Standard_S1, Premium_P1 and Premium_P2.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SkuParameters struct {

	// Specifies the number of units associated with this SignalR service. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 200, 300, 400, 500, 600, 700, 800, 900 and 1000.
	// +kubebuilder:validation:Optional
	Capacity *float64 `json:"capacity" tf:"capacity,omitempty"`

	// Specifies which tier to use. Valid values are Free_F1, Standard_S1, Premium_P1 and Premium_P2.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type UpstreamEndpointInitParameters struct {

	// The categories to match on, or * for all.
	CategoryPattern []*string `json:"categoryPattern,omitempty" tf:"category_pattern,omitempty"`

	// The events to match on, or * for all.
	EventPattern []*string `json:"eventPattern,omitempty" tf:"event_pattern,omitempty"`

	// The hubs to match on, or * for all.
	HubPattern []*string `json:"hubPattern,omitempty" tf:"hub_pattern,omitempty"`

	// The upstream URL Template. This can be a url or a template such as http://host.com/{hub}/api/{category}/{event}.
	URLTemplate *string `json:"urlTemplate,omitempty" tf:"url_template,omitempty"`

	// Specifies the Managed Identity IDs to be assigned to this signalR upstream setting by using resource uuid as both system assigned and user assigned identity is supported.
	UserAssignedIdentityID *string `json:"userAssignedIdentityId,omitempty" tf:"user_assigned_identity_id,omitempty"`
}

type UpstreamEndpointObservation struct {

	// The categories to match on, or * for all.
	CategoryPattern []*string `json:"categoryPattern,omitempty" tf:"category_pattern,omitempty"`

	// The events to match on, or * for all.
	EventPattern []*string `json:"eventPattern,omitempty" tf:"event_pattern,omitempty"`

	// The hubs to match on, or * for all.
	HubPattern []*string `json:"hubPattern,omitempty" tf:"hub_pattern,omitempty"`

	// The upstream URL Template. This can be a url or a template such as http://host.com/{hub}/api/{category}/{event}.
	URLTemplate *string `json:"urlTemplate,omitempty" tf:"url_template,omitempty"`

	// Specifies the Managed Identity IDs to be assigned to this signalR upstream setting by using resource uuid as both system assigned and user assigned identity is supported.
	UserAssignedIdentityID *string `json:"userAssignedIdentityId,omitempty" tf:"user_assigned_identity_id,omitempty"`
}

type UpstreamEndpointParameters struct {

	// The categories to match on, or * for all.
	// +kubebuilder:validation:Optional
	CategoryPattern []*string `json:"categoryPattern" tf:"category_pattern,omitempty"`

	// The events to match on, or * for all.
	// +kubebuilder:validation:Optional
	EventPattern []*string `json:"eventPattern" tf:"event_pattern,omitempty"`

	// The hubs to match on, or * for all.
	// +kubebuilder:validation:Optional
	HubPattern []*string `json:"hubPattern" tf:"hub_pattern,omitempty"`

	// The upstream URL Template. This can be a url or a template such as http://host.com/{hub}/api/{category}/{event}.
	// +kubebuilder:validation:Optional
	URLTemplate *string `json:"urlTemplate" tf:"url_template,omitempty"`

	// Specifies the Managed Identity IDs to be assigned to this signalR upstream setting by using resource uuid as both system assigned and user assigned identity is supported.
	// +kubebuilder:validation:Optional
	UserAssignedIdentityID *string `json:"userAssignedIdentityId,omitempty" tf:"user_assigned_identity_id,omitempty"`
}

// ServiceSpec defines the desired state of Service
type ServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceParameters `json:"forProvider"`
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
	InitProvider ServiceInitParameters `json:"initProvider,omitempty"`
}

// ServiceStatus defines the observed state of Service.
type ServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Service is the Schema for the Services API. Manages an Azure SignalR service.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Service struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sku) || (has(self.initProvider) && has(self.initProvider.sku))",message="spec.forProvider.sku is a required parameter"
	Spec   ServiceSpec   `json:"spec"`
	Status ServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceList contains a list of Services
type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Service `json:"items"`
}

// Repository type metadata.
var (
	Service_Kind             = "Service"
	Service_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Service_Kind}.String()
	Service_KindAPIVersion   = Service_Kind + "." + CRDGroupVersion.String()
	Service_GroupVersionKind = CRDGroupVersion.WithKind(Service_Kind)
)

func init() {
	SchemeBuilder.Register(&Service{}, &ServiceList{})
}
