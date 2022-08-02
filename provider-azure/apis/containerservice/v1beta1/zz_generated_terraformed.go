/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this KubernetesCluster
func (mg *KubernetesCluster) GetTerraformResourceType() string {
	return "azurerm_kubernetes_cluster"
}

// GetConnectionDetailsMapping for this KubernetesCluster
func (tr *KubernetesCluster) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"azure_active_directory_role_based_access_control[*].server_app_secret": "spec.forProvider.azureActiveDirectoryRoleBasedAccessControl[*].serverAppSecretSecretRef", "http_proxy_config[*].trusted_ca": "spec.forProvider.httpProxyConfig[*].trustedCaSecretRef", "kube_admin_config[*]": "status.atProvider.kubeAdminConfig[*]", "kube_admin_config[*].client_certificate": "status.atProvider.kubeAdminConfig[*].clientCertificate", "kube_admin_config[*].client_key": "status.atProvider.kubeAdminConfig[*].clientKey", "kube_admin_config[*].cluster_ca_certificate": "status.atProvider.kubeAdminConfig[*].clusterCaCertificate", "kube_admin_config[*].password": "status.atProvider.kubeAdminConfig[*].password", "kube_admin_config_raw": "status.atProvider.kubeAdminConfigRaw", "kube_config[*]": "status.atProvider.kubeConfig[*]", "kube_config[*].client_certificate": "status.atProvider.kubeConfig[*].clientCertificate", "kube_config[*].client_key": "status.atProvider.kubeConfig[*].clientKey", "kube_config[*].cluster_ca_certificate": "status.atProvider.kubeConfig[*].clusterCaCertificate", "kube_config[*].password": "status.atProvider.kubeConfig[*].password", "kube_config_raw": "status.atProvider.kubeConfigRaw", "service_principal[*].client_secret": "spec.forProvider.servicePrincipal[*].clientSecretSecretRef", "windows_profile[*].admin_password": "spec.forProvider.windowsProfile[*].adminPasswordSecretRef"}
}

// GetObservation of this KubernetesCluster
func (tr *KubernetesCluster) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this KubernetesCluster
func (tr *KubernetesCluster) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this KubernetesCluster
func (tr *KubernetesCluster) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this KubernetesCluster
func (tr *KubernetesCluster) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this KubernetesCluster
func (tr *KubernetesCluster) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this KubernetesCluster using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *KubernetesCluster) LateInitialize(attrs []byte) (bool, error) {
	params := &KubernetesClusterParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}
	opts = append(opts, resource.WithNameFilter("KubeletIdentity"))

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *KubernetesCluster) GetTerraformSchemaVersion() int {
	return 2
}

// GetTerraformResourceType returns Terraform resource type for this KubernetesClusterNodePool
func (mg *KubernetesClusterNodePool) GetTerraformResourceType() string {
	return "azurerm_kubernetes_cluster_node_pool"
}

// GetConnectionDetailsMapping for this KubernetesClusterNodePool
func (tr *KubernetesClusterNodePool) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this KubernetesClusterNodePool
func (tr *KubernetesClusterNodePool) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this KubernetesClusterNodePool
func (tr *KubernetesClusterNodePool) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this KubernetesClusterNodePool
func (tr *KubernetesClusterNodePool) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this KubernetesClusterNodePool
func (tr *KubernetesClusterNodePool) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this KubernetesClusterNodePool
func (tr *KubernetesClusterNodePool) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this KubernetesClusterNodePool using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *KubernetesClusterNodePool) LateInitialize(attrs []byte) (bool, error) {
	params := &KubernetesClusterNodePoolParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *KubernetesClusterNodePool) GetTerraformSchemaVersion() int {
	return 1
}
