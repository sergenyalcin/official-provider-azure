// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	configuration "github.com/upbound/provider-azure/internal/controller/cluster/dbformariadb/configuration"
	database "github.com/upbound/provider-azure/internal/controller/cluster/dbformariadb/database"
	firewallrule "github.com/upbound/provider-azure/internal/controller/cluster/dbformariadb/firewallrule"
	server "github.com/upbound/provider-azure/internal/controller/cluster/dbformariadb/server"
	virtualnetworkrule "github.com/upbound/provider-azure/internal/controller/cluster/dbformariadb/virtualnetworkrule"
)

var dbformariadbCrdGroup = "dbformariadb.azure.upbound.io"

// Setup_dbformariadb creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dbformariadb(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: dbformariadbCrdGroup, Kind: "Configuration"}:      configuration.Setup,
		schema.GroupKind{Group: dbformariadbCrdGroup, Kind: "Database"}:           database.Setup,
		schema.GroupKind{Group: dbformariadbCrdGroup, Kind: "FirewallRule"}:       firewallrule.Setup,
		schema.GroupKind{Group: dbformariadbCrdGroup, Kind: "Server"}:             server.Setup,
		schema.GroupKind{Group: dbformariadbCrdGroup, Kind: "VirtualNetworkRule"}: virtualnetworkrule.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, dbformariadbCrdGroup, o); err != nil {
		return err
	}
	return nil
}
