// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	activedirectoryadministrator "github.com/upbound/provider-azure/internal/controller/cluster/dbformysql/activedirectoryadministrator"
	configuration "github.com/upbound/provider-azure/internal/controller/cluster/dbformysql/configuration"
	database "github.com/upbound/provider-azure/internal/controller/cluster/dbformysql/database"
	firewallrule "github.com/upbound/provider-azure/internal/controller/cluster/dbformysql/firewallrule"
	flexibledatabase "github.com/upbound/provider-azure/internal/controller/cluster/dbformysql/flexibledatabase"
	flexibleserver "github.com/upbound/provider-azure/internal/controller/cluster/dbformysql/flexibleserver"
	flexibleserverconfiguration "github.com/upbound/provider-azure/internal/controller/cluster/dbformysql/flexibleserverconfiguration"
	flexibleserverfirewallrule "github.com/upbound/provider-azure/internal/controller/cluster/dbformysql/flexibleserverfirewallrule"
	server "github.com/upbound/provider-azure/internal/controller/cluster/dbformysql/server"
	virtualnetworkrule "github.com/upbound/provider-azure/internal/controller/cluster/dbformysql/virtualnetworkrule"
)

var dbformysqlCrdGroup = "dbformysql.azure.upbound.io"

// Setup_dbformysql creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dbformysql(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: dbformysqlCrdGroup, Kind: "ActiveDirectoryAdministrator"}: activedirectoryadministrator.Setup,
		schema.GroupKind{Group: dbformysqlCrdGroup, Kind: "Configuration"}:                configuration.Setup,
		schema.GroupKind{Group: dbformysqlCrdGroup, Kind: "Database"}:                     database.Setup,
		schema.GroupKind{Group: dbformysqlCrdGroup, Kind: "FirewallRule"}:                 firewallrule.Setup,
		schema.GroupKind{Group: dbformysqlCrdGroup, Kind: "FlexibleDatabase"}:             flexibledatabase.Setup,
		schema.GroupKind{Group: dbformysqlCrdGroup, Kind: "FlexibleServer"}:               flexibleserver.Setup,
		schema.GroupKind{Group: dbformysqlCrdGroup, Kind: "FlexibleServerConfiguration"}:  flexibleserverconfiguration.Setup,
		schema.GroupKind{Group: dbformysqlCrdGroup, Kind: "FlexibleServerFirewallRule"}:   flexibleserverfirewallrule.Setup,
		schema.GroupKind{Group: dbformysqlCrdGroup, Kind: "Server"}:                       server.Setup,
		schema.GroupKind{Group: dbformysqlCrdGroup, Kind: "VirtualNetworkRule"}:           virtualnetworkrule.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
