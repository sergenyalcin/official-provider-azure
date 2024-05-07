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

type AuthorizationInitParameters struct {

	// The authentication Parameter value.
	Parameter *string `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// The authentication Scheme name.
	Scheme *string `json:"scheme,omitempty" tf:"scheme,omitempty"`
}

type AuthorizationObservation struct {

	// The authentication Parameter value.
	Parameter *string `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// The authentication Scheme name.
	Scheme *string `json:"scheme,omitempty" tf:"scheme,omitempty"`
}

type AuthorizationParameters struct {

	// The authentication Parameter value.
	// +kubebuilder:validation:Optional
	Parameter *string `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// The authentication Scheme name.
	// +kubebuilder:validation:Optional
	Scheme *string `json:"scheme,omitempty" tf:"scheme,omitempty"`
}

type BackendInitParameters struct {

	// A credentials block as documented below.
	Credentials *CredentialsInitParameters `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// The description of the backend.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The protocol used by the backend host. Possible values are http or soap.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// A proxy block as documented below.
	Proxy *BackendProxyInitParameters `json:"proxy,omitempty" tf:"proxy,omitempty"`

	// The management URI of the backend host in an external system. This URI can be the ARM Resource ID of Logic Apps, Function Apps or API Apps, or the management endpoint of a Service Fabric cluster.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// A service_fabric_cluster block as documented below.
	ServiceFabricCluster *ServiceFabricClusterInitParameters `json:"serviceFabricCluster,omitempty" tf:"service_fabric_cluster,omitempty"`

	// A tls block as documented below.
	TLS *TLSInitParameters `json:"tls,omitempty" tf:"tls,omitempty"`

	// The title of the backend.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`

	// The URL of the backend host.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type BackendObservation struct {

	// The Name of the API Management Service where this backend should be created. Changing this forces a new resource to be created.
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// A credentials block as documented below.
	Credentials *CredentialsObservation `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// The description of the backend.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the API Management API.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The protocol used by the backend host. Possible values are http or soap.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// A proxy block as documented below.
	Proxy *BackendProxyObservation `json:"proxy,omitempty" tf:"proxy,omitempty"`

	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The management URI of the backend host in an external system. This URI can be the ARM Resource ID of Logic Apps, Function Apps or API Apps, or the management endpoint of a Service Fabric cluster.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// A service_fabric_cluster block as documented below.
	ServiceFabricCluster *ServiceFabricClusterObservation `json:"serviceFabricCluster,omitempty" tf:"service_fabric_cluster,omitempty"`

	// A tls block as documented below.
	TLS *TLSObservation `json:"tls,omitempty" tf:"tls,omitempty"`

	// The title of the backend.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`

	// The URL of the backend host.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type BackendParameters struct {

	// The Name of the API Management Service where this backend should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta2.Management
	// +kubebuilder:validation:Optional
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// A credentials block as documented below.
	// +kubebuilder:validation:Optional
	Credentials *CredentialsParameters `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// The description of the backend.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The protocol used by the backend host. Possible values are http or soap.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// A proxy block as documented below.
	// +kubebuilder:validation:Optional
	Proxy *BackendProxyParameters `json:"proxy,omitempty" tf:"proxy,omitempty"`

	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The management URI of the backend host in an external system. This URI can be the ARM Resource ID of Logic Apps, Function Apps or API Apps, or the management endpoint of a Service Fabric cluster.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// A service_fabric_cluster block as documented below.
	// +kubebuilder:validation:Optional
	ServiceFabricCluster *ServiceFabricClusterParameters `json:"serviceFabricCluster,omitempty" tf:"service_fabric_cluster,omitempty"`

	// A tls block as documented below.
	// +kubebuilder:validation:Optional
	TLS *TLSParameters `json:"tls,omitempty" tf:"tls,omitempty"`

	// The title of the backend.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`

	// The URL of the backend host.
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type BackendProxyInitParameters struct {

	// The URL of the proxy server.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// The username to connect to the proxy server.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type BackendProxyObservation struct {

	// The URL of the proxy server.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// The username to connect to the proxy server.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type BackendProxyParameters struct {

	// The password to connect to the proxy server.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// The URL of the proxy server.
	// +kubebuilder:validation:Optional
	URL *string `json:"url" tf:"url,omitempty"`

	// The username to connect to the proxy server.
	// +kubebuilder:validation:Optional
	Username *string `json:"username" tf:"username,omitempty"`
}

type CredentialsInitParameters struct {

	// An authorization block as defined below.
	Authorization *AuthorizationInitParameters `json:"authorization,omitempty" tf:"authorization,omitempty"`

	// A list of client certificate thumbprints to present to the backend host. The certificates must exist within the API Management Service.
	Certificate []*string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// A mapping of header parameters to pass to the backend host. The keys are the header names and the values are a comma separated string of header values. This is converted to a list before being passed to the API.
	// +mapType=granular
	Header map[string]*string `json:"header,omitempty" tf:"header,omitempty"`

	// A mapping of query parameters to pass to the backend host. The keys are the query names and the values are a comma separated string of query values. This is converted to a list before being passed to the API.
	// +mapType=granular
	Query map[string]*string `json:"query,omitempty" tf:"query,omitempty"`
}

type CredentialsObservation struct {

	// An authorization block as defined below.
	Authorization *AuthorizationObservation `json:"authorization,omitempty" tf:"authorization,omitempty"`

	// A list of client certificate thumbprints to present to the backend host. The certificates must exist within the API Management Service.
	Certificate []*string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// A mapping of header parameters to pass to the backend host. The keys are the header names and the values are a comma separated string of header values. This is converted to a list before being passed to the API.
	// +mapType=granular
	Header map[string]*string `json:"header,omitempty" tf:"header,omitempty"`

	// A mapping of query parameters to pass to the backend host. The keys are the query names and the values are a comma separated string of query values. This is converted to a list before being passed to the API.
	// +mapType=granular
	Query map[string]*string `json:"query,omitempty" tf:"query,omitempty"`
}

type CredentialsParameters struct {

	// An authorization block as defined below.
	// +kubebuilder:validation:Optional
	Authorization *AuthorizationParameters `json:"authorization,omitempty" tf:"authorization,omitempty"`

	// A list of client certificate thumbprints to present to the backend host. The certificates must exist within the API Management Service.
	// +kubebuilder:validation:Optional
	Certificate []*string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// A mapping of header parameters to pass to the backend host. The keys are the header names and the values are a comma separated string of header values. This is converted to a list before being passed to the API.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Header map[string]*string `json:"header,omitempty" tf:"header,omitempty"`

	// A mapping of query parameters to pass to the backend host. The keys are the query names and the values are a comma separated string of query values. This is converted to a list before being passed to the API.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Query map[string]*string `json:"query,omitempty" tf:"query,omitempty"`
}

type ServerX509NameInitParameters struct {

	// The thumbprint for the issuer of the certificate.
	IssuerCertificateThumbprint *string `json:"issuerCertificateThumbprint,omitempty" tf:"issuer_certificate_thumbprint,omitempty"`

	// The common name of the certificate.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ServerX509NameObservation struct {

	// The thumbprint for the issuer of the certificate.
	IssuerCertificateThumbprint *string `json:"issuerCertificateThumbprint,omitempty" tf:"issuer_certificate_thumbprint,omitempty"`

	// The common name of the certificate.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ServerX509NameParameters struct {

	// The thumbprint for the issuer of the certificate.
	// +kubebuilder:validation:Optional
	IssuerCertificateThumbprint *string `json:"issuerCertificateThumbprint" tf:"issuer_certificate_thumbprint,omitempty"`

	// The common name of the certificate.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type ServiceFabricClusterInitParameters struct {

	// The client certificate resource id for the management endpoint.
	ClientCertificateID *string `json:"clientCertificateId,omitempty" tf:"client_certificate_id,omitempty"`

	// The client certificate thumbprint for the management endpoint.
	ClientCertificateThumbprint *string `json:"clientCertificateThumbprint,omitempty" tf:"client_certificate_thumbprint,omitempty"`

	// A list of cluster management endpoints.
	// +listType=set
	ManagementEndpoints []*string `json:"managementEndpoints,omitempty" tf:"management_endpoints,omitempty"`

	// The maximum number of retries when attempting resolve the partition.
	MaxPartitionResolutionRetries *float64 `json:"maxPartitionResolutionRetries,omitempty" tf:"max_partition_resolution_retries,omitempty"`

	// A list of thumbprints of the server certificates of the Service Fabric cluster.
	// +listType=set
	ServerCertificateThumbprints []*string `json:"serverCertificateThumbprints,omitempty" tf:"server_certificate_thumbprints,omitempty"`

	// One or more server_x509_name blocks as documented below.
	ServerX509Name []ServerX509NameInitParameters `json:"serverX509Name,omitempty" tf:"server_x509_name,omitempty"`
}

type ServiceFabricClusterObservation struct {

	// The client certificate resource id for the management endpoint.
	ClientCertificateID *string `json:"clientCertificateId,omitempty" tf:"client_certificate_id,omitempty"`

	// The client certificate thumbprint for the management endpoint.
	ClientCertificateThumbprint *string `json:"clientCertificateThumbprint,omitempty" tf:"client_certificate_thumbprint,omitempty"`

	// A list of cluster management endpoints.
	// +listType=set
	ManagementEndpoints []*string `json:"managementEndpoints,omitempty" tf:"management_endpoints,omitempty"`

	// The maximum number of retries when attempting resolve the partition.
	MaxPartitionResolutionRetries *float64 `json:"maxPartitionResolutionRetries,omitempty" tf:"max_partition_resolution_retries,omitempty"`

	// A list of thumbprints of the server certificates of the Service Fabric cluster.
	// +listType=set
	ServerCertificateThumbprints []*string `json:"serverCertificateThumbprints,omitempty" tf:"server_certificate_thumbprints,omitempty"`

	// One or more server_x509_name blocks as documented below.
	ServerX509Name []ServerX509NameObservation `json:"serverX509Name,omitempty" tf:"server_x509_name,omitempty"`
}

type ServiceFabricClusterParameters struct {

	// The client certificate resource id for the management endpoint.
	// +kubebuilder:validation:Optional
	ClientCertificateID *string `json:"clientCertificateId,omitempty" tf:"client_certificate_id,omitempty"`

	// The client certificate thumbprint for the management endpoint.
	// +kubebuilder:validation:Optional
	ClientCertificateThumbprint *string `json:"clientCertificateThumbprint,omitempty" tf:"client_certificate_thumbprint,omitempty"`

	// A list of cluster management endpoints.
	// +kubebuilder:validation:Optional
	// +listType=set
	ManagementEndpoints []*string `json:"managementEndpoints" tf:"management_endpoints,omitempty"`

	// The maximum number of retries when attempting resolve the partition.
	// +kubebuilder:validation:Optional
	MaxPartitionResolutionRetries *float64 `json:"maxPartitionResolutionRetries" tf:"max_partition_resolution_retries,omitempty"`

	// A list of thumbprints of the server certificates of the Service Fabric cluster.
	// +kubebuilder:validation:Optional
	// +listType=set
	ServerCertificateThumbprints []*string `json:"serverCertificateThumbprints,omitempty" tf:"server_certificate_thumbprints,omitempty"`

	// One or more server_x509_name blocks as documented below.
	// +kubebuilder:validation:Optional
	ServerX509Name []ServerX509NameParameters `json:"serverX509Name,omitempty" tf:"server_x509_name,omitempty"`
}

type TLSInitParameters struct {

	// Flag indicating whether SSL certificate chain validation should be done when using self-signed certificates for the backend host.
	ValidateCertificateChain *bool `json:"validateCertificateChain,omitempty" tf:"validate_certificate_chain,omitempty"`

	// Flag indicating whether SSL certificate name validation should be done when using self-signed certificates for the backend host.
	ValidateCertificateName *bool `json:"validateCertificateName,omitempty" tf:"validate_certificate_name,omitempty"`
}

type TLSObservation struct {

	// Flag indicating whether SSL certificate chain validation should be done when using self-signed certificates for the backend host.
	ValidateCertificateChain *bool `json:"validateCertificateChain,omitempty" tf:"validate_certificate_chain,omitempty"`

	// Flag indicating whether SSL certificate name validation should be done when using self-signed certificates for the backend host.
	ValidateCertificateName *bool `json:"validateCertificateName,omitempty" tf:"validate_certificate_name,omitempty"`
}

type TLSParameters struct {

	// Flag indicating whether SSL certificate chain validation should be done when using self-signed certificates for the backend host.
	// +kubebuilder:validation:Optional
	ValidateCertificateChain *bool `json:"validateCertificateChain,omitempty" tf:"validate_certificate_chain,omitempty"`

	// Flag indicating whether SSL certificate name validation should be done when using self-signed certificates for the backend host.
	// +kubebuilder:validation:Optional
	ValidateCertificateName *bool `json:"validateCertificateName,omitempty" tf:"validate_certificate_name,omitempty"`
}

// BackendSpec defines the desired state of Backend
type BackendSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackendParameters `json:"forProvider"`
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
	InitProvider BackendInitParameters `json:"initProvider,omitempty"`
}

// BackendStatus defines the observed state of Backend.
type BackendStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackendObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Backend is the Schema for the Backends API. Manages a backend within an API Management Service.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Backend struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocol) || (has(self.initProvider) && has(self.initProvider.protocol))",message="spec.forProvider.protocol is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.url) || (has(self.initProvider) && has(self.initProvider.url))",message="spec.forProvider.url is a required parameter"
	Spec   BackendSpec   `json:"spec"`
	Status BackendStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackendList contains a list of Backends
type BackendList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Backend `json:"items"`
}

// Repository type metadata.
var (
	Backend_Kind             = "Backend"
	Backend_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Backend_Kind}.String()
	Backend_KindAPIVersion   = Backend_Kind + "." + CRDGroupVersion.String()
	Backend_GroupVersionKind = CRDGroupVersion.WithKind(Backend_Kind)
)

func init() {
	SchemeBuilder.Register(&Backend{}, &BackendList{})
}
