// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	configuration "github.com/upbound/provider-azure/internal/controller/namespaced/dbformariadb/configuration"
	database "github.com/upbound/provider-azure/internal/controller/namespaced/dbformariadb/database"
	firewallrule "github.com/upbound/provider-azure/internal/controller/namespaced/dbformariadb/firewallrule"
	server "github.com/upbound/provider-azure/internal/controller/namespaced/dbformariadb/server"
	virtualnetworkrule "github.com/upbound/provider-azure/internal/controller/namespaced/dbformariadb/virtualnetworkrule"
)

// Setup_dbformariadb creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dbformariadb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		configuration.Setup,
		database.Setup,
		firewallrule.Setup,
		server.Setup,
		virtualnetworkrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
