// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	mssqldatabase "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqldatabase"
	mssqldatabaseextendedauditingpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqldatabaseextendedauditingpolicy"
	mssqldatabasevulnerabilityassessmentrulebaseline "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqldatabasevulnerabilityassessmentrulebaseline"
	mssqlelasticpool "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlelasticpool"
	mssqlfailovergroup "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlfailovergroup"
	mssqlfirewallrule "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlfirewallrule"
	mssqljobagent "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqljobagent"
	mssqljobcredential "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqljobcredential"
	mssqlmanageddatabase "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlmanageddatabase"
	mssqlmanagedinstance "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlmanagedinstance"
	mssqlmanagedinstanceactivedirectoryadministrator "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlmanagedinstanceactivedirectoryadministrator"
	mssqlmanagedinstancefailovergroup "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlmanagedinstancefailovergroup"
	mssqlmanagedinstancetransparentdataencryption "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlmanagedinstancetransparentdataencryption"
	mssqlmanagedinstancevulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlmanagedinstancevulnerabilityassessment"
	mssqloutboundfirewallrule "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqloutboundfirewallrule"
	mssqlserver "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlserver"
	mssqlserverdnsalias "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlserverdnsalias"
	mssqlservermicrosoftsupportauditingpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlservermicrosoftsupportauditingpolicy"
	mssqlserversecurityalertpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlserversecurityalertpolicy"
	mssqlservertransparentdataencryption "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlservertransparentdataencryption"
	mssqlservervulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlservervulnerabilityassessment"
	mssqlvirtualnetworkrule "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlvirtualnetworkrule"
)

var sqlCrdGroup = "sql.azure.m.upbound.io"

// Setup_sql creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_sql(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: sqlCrdGroup, Kind: "MSSQLDatabase"}:                                    mssqldatabase.Setup,
		schema.GroupKind{Group: sqlCrdGroup, Kind: "MSSQLDatabaseExtendedAuditingPolicy"}:              mssqldatabaseextendedauditingpolicy.Setup,
		schema.GroupKind{Group: sqlCrdGroup, Kind: "MSSQLDatabaseVulnerabilityAssessmentRuleBaseline"}: mssqldatabasevulnerabilityassessmentrulebaseline.Setup,
		schema.GroupKind{Group: sqlCrdGroup, Kind: "MSSQLElasticPool"}:                                 mssqlelasticpool.Setup,
		schema.GroupKind{Group: sqlCrdGroup, Kind: "MSSQLFailoverGroup"}:                               mssqlfailovergroup.Setup,
		schema.GroupKind{Group: sqlCrdGroup, Kind: "MSSQLFirewallRule"}:                                mssqlfirewallrule.Setup,
		schema.GroupKind{Group: sqlCrdGroup, Kind: "MSSQLJobAgent"}:                                    mssqljobagent.Setup,
		schema.GroupKind{Group: sqlCrdGroup, Kind: "MSSQLJobCredential"}:                               mssqljobcredential.Setup,
		schema.GroupKind{Group: sqlCrdGroup, Kind: "MSSQLManagedDatabase"}:                             mssqlmanageddatabase.Setup,
		schema.GroupKind{Group: sqlCrdGroup, Kind: "MSSQLManagedInstance"}:                             mssqlmanagedinstance.Setup,
		schema.GroupKind{Group: sqlCrdGroup, Kind: "MSSQLManagedInstanceActiveDirectoryAdministrator"}: mssqlmanagedinstanceactivedirectoryadministrator.Setup,
		schema.GroupKind{Group: sqlCrdGroup, Kind: "MSSQLManagedInstanceFailoverGroup"}:                mssqlmanagedinstancefailovergroup.Setup,
		schema.GroupKind{Group: sqlCrdGroup, Kind: "MSSQLManagedInstanceTransparentDataEncryption"}:    mssqlmanagedinstancetransparentdataencryption.Setup,
		schema.GroupKind{Group: sqlCrdGroup, Kind: "MSSQLManagedInstanceVulnerabilityAssessment"}:      mssqlmanagedinstancevulnerabilityassessment.Setup,
		schema.GroupKind{Group: sqlCrdGroup, Kind: "MSSQLOutboundFirewallRule"}:                        mssqloutboundfirewallrule.Setup,
		schema.GroupKind{Group: sqlCrdGroup, Kind: "MSSQLServer"}:                                      mssqlserver.Setup,
		schema.GroupKind{Group: sqlCrdGroup, Kind: "MSSQLServerDNSAlias"}:                              mssqlserverdnsalias.Setup,
		schema.GroupKind{Group: sqlCrdGroup, Kind: "MSSQLServerMicrosoftSupportAuditingPolicy"}:        mssqlservermicrosoftsupportauditingpolicy.Setup,
		schema.GroupKind{Group: sqlCrdGroup, Kind: "MSSQLServerSecurityAlertPolicy"}:                   mssqlserversecurityalertpolicy.Setup,
		schema.GroupKind{Group: sqlCrdGroup, Kind: "MSSQLServerTransparentDataEncryption"}:             mssqlservertransparentdataencryption.Setup,
		schema.GroupKind{Group: sqlCrdGroup, Kind: "MSSQLServerVulnerabilityAssessment"}:               mssqlservervulnerabilityassessment.Setup,
		schema.GroupKind{Group: sqlCrdGroup, Kind: "MSSQLVirtualNetworkRule"}:                          mssqlvirtualnetworkrule.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, sqlCrdGroup, o); err != nil {
		return err
	}
	return nil
}
