// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	namespaceauthorizationrule "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/namespaceauthorizationrule"
	namespacedisasterrecoveryconfig "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/namespacedisasterrecoveryconfig"
	namespacenetworkruleset "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/namespacenetworkruleset"
	queue "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/queue"
	queueauthorizationrule "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/queueauthorizationrule"
	servicebusnamespace "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/servicebusnamespace"
	subscription "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/subscription"
	subscriptionrule "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/subscriptionrule"
	topic "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/topic"
	topicauthorizationrule "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/topicauthorizationrule"
)

var servicebusCrdGroup = "servicebus.azure.upbound.io"

// Setup_servicebus creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_servicebus(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: servicebusCrdGroup, Kind: "NamespaceAuthorizationRule"}:      namespaceauthorizationrule.Setup,
		schema.GroupKind{Group: servicebusCrdGroup, Kind: "NamespaceDisasterRecoveryConfig"}: namespacedisasterrecoveryconfig.Setup,
		schema.GroupKind{Group: servicebusCrdGroup, Kind: "NamespaceNetworkRuleSet"}:         namespacenetworkruleset.Setup,
		schema.GroupKind{Group: servicebusCrdGroup, Kind: "Queue"}:                           queue.Setup,
		schema.GroupKind{Group: servicebusCrdGroup, Kind: "QueueAuthorizationRule"}:          queueauthorizationrule.Setup,
		schema.GroupKind{Group: servicebusCrdGroup, Kind: "ServiceBusNamespace"}:             servicebusnamespace.Setup,
		schema.GroupKind{Group: servicebusCrdGroup, Kind: "Subscription"}:                    subscription.Setup,
		schema.GroupKind{Group: servicebusCrdGroup, Kind: "SubscriptionRule"}:                subscriptionrule.Setup,
		schema.GroupKind{Group: servicebusCrdGroup, Kind: "Topic"}:                           topic.Setup,
		schema.GroupKind{Group: servicebusCrdGroup, Kind: "TopicAuthorizationRule"}:          topicauthorizationrule.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
