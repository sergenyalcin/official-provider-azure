// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	firewallrule "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/firewallrule"
	integrationruntimeazure "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/integrationruntimeazure"
	integrationruntimeselfhosted "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/integrationruntimeselfhosted"
	linkedservice "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/linkedservice"
	managedprivateendpoint "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/managedprivateendpoint"
	privatelinkhub "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/privatelinkhub"
	roleassignment "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/roleassignment"
	sparkpool "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/sparkpool"
	sqlpool "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/sqlpool"
	sqlpoolextendedauditingpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/sqlpoolextendedauditingpolicy"
	sqlpoolsecurityalertpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/sqlpoolsecurityalertpolicy"
	sqlpoolworkloadclassifier "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/sqlpoolworkloadclassifier"
	sqlpoolworkloadgroup "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/sqlpoolworkloadgroup"
	workspace "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/workspace"
	workspaceaadadmin "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/workspaceaadadmin"
	workspaceextendedauditingpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/workspaceextendedauditingpolicy"
	workspacesecurityalertpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/workspacesecurityalertpolicy"
	workspacesqlaadadmin "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/workspacesqlaadadmin"
	workspacevulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/workspacevulnerabilityassessment"
)

var synapseCrdGroup = "synapse.azure.m.upbound.io"

// Setup_synapse creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_synapse(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: synapseCrdGroup, Kind: "FirewallRule"}:                     firewallrule.Setup,
		schema.GroupKind{Group: synapseCrdGroup, Kind: "IntegrationRuntimeAzure"}:          integrationruntimeazure.Setup,
		schema.GroupKind{Group: synapseCrdGroup, Kind: "IntegrationRuntimeSelfHosted"}:     integrationruntimeselfhosted.Setup,
		schema.GroupKind{Group: synapseCrdGroup, Kind: "LinkedService"}:                    linkedservice.Setup,
		schema.GroupKind{Group: synapseCrdGroup, Kind: "ManagedPrivateEndpoint"}:           managedprivateendpoint.Setup,
		schema.GroupKind{Group: synapseCrdGroup, Kind: "PrivateLinkHub"}:                   privatelinkhub.Setup,
		schema.GroupKind{Group: synapseCrdGroup, Kind: "RoleAssignment"}:                   roleassignment.Setup,
		schema.GroupKind{Group: synapseCrdGroup, Kind: "SQLPool"}:                          sqlpool.Setup,
		schema.GroupKind{Group: synapseCrdGroup, Kind: "SQLPoolExtendedAuditingPolicy"}:    sqlpoolextendedauditingpolicy.Setup,
		schema.GroupKind{Group: synapseCrdGroup, Kind: "SQLPoolSecurityAlertPolicy"}:       sqlpoolsecurityalertpolicy.Setup,
		schema.GroupKind{Group: synapseCrdGroup, Kind: "SQLPoolWorkloadClassifier"}:        sqlpoolworkloadclassifier.Setup,
		schema.GroupKind{Group: synapseCrdGroup, Kind: "SQLPoolWorkloadGroup"}:             sqlpoolworkloadgroup.Setup,
		schema.GroupKind{Group: synapseCrdGroup, Kind: "SparkPool"}:                        sparkpool.Setup,
		schema.GroupKind{Group: synapseCrdGroup, Kind: "Workspace"}:                        workspace.Setup,
		schema.GroupKind{Group: synapseCrdGroup, Kind: "WorkspaceAADAdmin"}:                workspaceaadadmin.Setup,
		schema.GroupKind{Group: synapseCrdGroup, Kind: "WorkspaceExtendedAuditingPolicy"}:  workspaceextendedauditingpolicy.Setup,
		schema.GroupKind{Group: synapseCrdGroup, Kind: "WorkspaceSQLAADAdmin"}:             workspacesqlaadadmin.Setup,
		schema.GroupKind{Group: synapseCrdGroup, Kind: "WorkspaceSecurityAlertPolicy"}:     workspacesecurityalertpolicy.Setup,
		schema.GroupKind{Group: synapseCrdGroup, Kind: "WorkspaceVulnerabilityAssessment"}: workspacevulnerabilityassessment.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
