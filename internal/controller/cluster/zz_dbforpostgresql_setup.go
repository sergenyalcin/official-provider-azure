// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	activedirectoryadministrator "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/activedirectoryadministrator"
	configuration "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/configuration"
	database "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/database"
	firewallrule "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/firewallrule"
	flexibleserver "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/flexibleserver"
	flexibleserveractivedirectoryadministrator "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/flexibleserveractivedirectoryadministrator"
	flexibleserverconfiguration "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/flexibleserverconfiguration"
	flexibleserverdatabase "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/flexibleserverdatabase"
	flexibleserverfirewallrule "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/flexibleserverfirewallrule"
	server "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/server"
	serverkey "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/serverkey"
	virtualnetworkrule "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/virtualnetworkrule"
)

var dbforpostgresqlCrdGroup = "dbforpostgresql.azure.upbound.io"

// Setup_dbforpostgresql creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dbforpostgresql(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: dbforpostgresqlCrdGroup, Kind: "ActiveDirectoryAdministrator"}:               activedirectoryadministrator.Setup,
		schema.GroupKind{Group: dbforpostgresqlCrdGroup, Kind: "Configuration"}:                              configuration.Setup,
		schema.GroupKind{Group: dbforpostgresqlCrdGroup, Kind: "Database"}:                                   database.Setup,
		schema.GroupKind{Group: dbforpostgresqlCrdGroup, Kind: "FirewallRule"}:                               firewallrule.Setup,
		schema.GroupKind{Group: dbforpostgresqlCrdGroup, Kind: "FlexibleServer"}:                             flexibleserver.Setup,
		schema.GroupKind{Group: dbforpostgresqlCrdGroup, Kind: "FlexibleServerActiveDirectoryAdministrator"}: flexibleserveractivedirectoryadministrator.Setup,
		schema.GroupKind{Group: dbforpostgresqlCrdGroup, Kind: "FlexibleServerConfiguration"}:                flexibleserverconfiguration.Setup,
		schema.GroupKind{Group: dbforpostgresqlCrdGroup, Kind: "FlexibleServerDatabase"}:                     flexibleserverdatabase.Setup,
		schema.GroupKind{Group: dbforpostgresqlCrdGroup, Kind: "FlexibleServerFirewallRule"}:                 flexibleserverfirewallrule.Setup,
		schema.GroupKind{Group: dbforpostgresqlCrdGroup, Kind: "Server"}:                                     server.Setup,
		schema.GroupKind{Group: dbforpostgresqlCrdGroup, Kind: "ServerKey"}:                                  serverkey.Setup,
		schema.GroupKind{Group: dbforpostgresqlCrdGroup, Kind: "VirtualNetworkRule"}:                         virtualnetworkrule.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, dbforpostgresqlCrdGroup, o); err != nil {
		return err
	}
	return nil
}
