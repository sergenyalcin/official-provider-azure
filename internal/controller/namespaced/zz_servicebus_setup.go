// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	namespaceauthorizationrule "github.com/upbound/provider-azure/internal/controller/namespaced/servicebus/namespaceauthorizationrule"
	namespacedisasterrecoveryconfig "github.com/upbound/provider-azure/internal/controller/namespaced/servicebus/namespacedisasterrecoveryconfig"
	namespacenetworkruleset "github.com/upbound/provider-azure/internal/controller/namespaced/servicebus/namespacenetworkruleset"
	queue "github.com/upbound/provider-azure/internal/controller/namespaced/servicebus/queue"
	queueauthorizationrule "github.com/upbound/provider-azure/internal/controller/namespaced/servicebus/queueauthorizationrule"
	servicebusnamespace "github.com/upbound/provider-azure/internal/controller/namespaced/servicebus/servicebusnamespace"
	subscription "github.com/upbound/provider-azure/internal/controller/namespaced/servicebus/subscription"
	subscriptionrule "github.com/upbound/provider-azure/internal/controller/namespaced/servicebus/subscriptionrule"
	topic "github.com/upbound/provider-azure/internal/controller/namespaced/servicebus/topic"
	topicauthorizationrule "github.com/upbound/provider-azure/internal/controller/namespaced/servicebus/topicauthorizationrule"
)

// Setup_servicebus creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_servicebus(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		namespaceauthorizationrule.Setup,
		namespacedisasterrecoveryconfig.Setup,
		namespacenetworkruleset.Setup,
		queue.Setup,
		queueauthorizationrule.Setup,
		servicebusnamespace.Setup,
		subscription.Setup,
		subscriptionrule.Setup,
		topic.Setup,
		topicauthorizationrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
