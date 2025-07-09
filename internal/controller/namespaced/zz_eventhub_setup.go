// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	authorizationrule "github.com/upbound/provider-azure/internal/controller/namespaced/eventhub/authorizationrule"
	consumergroup "github.com/upbound/provider-azure/internal/controller/namespaced/eventhub/consumergroup"
	eventhub "github.com/upbound/provider-azure/internal/controller/namespaced/eventhub/eventhub"
	eventhubnamespace "github.com/upbound/provider-azure/internal/controller/namespaced/eventhub/eventhubnamespace"
	namespaceauthorizationrule "github.com/upbound/provider-azure/internal/controller/namespaced/eventhub/namespaceauthorizationrule"
	namespacedisasterrecoveryconfig "github.com/upbound/provider-azure/internal/controller/namespaced/eventhub/namespacedisasterrecoveryconfig"
	namespaceschemagroup "github.com/upbound/provider-azure/internal/controller/namespaced/eventhub/namespaceschemagroup"
)

// Setup_eventhub creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_eventhub(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		authorizationrule.Setup,
		consumergroup.Setup,
		eventhub.Setup,
		eventhubnamespace.Setup,
		namespaceauthorizationrule.Setup,
		namespacedisasterrecoveryconfig.Setup,
		namespaceschemagroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
