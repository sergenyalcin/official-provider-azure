// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	authorizationrule "github.com/upbound/provider-azure/internal/controller/cluster/eventhub/authorizationrule"
	consumergroup "github.com/upbound/provider-azure/internal/controller/cluster/eventhub/consumergroup"
	eventhub "github.com/upbound/provider-azure/internal/controller/cluster/eventhub/eventhub"
	eventhubnamespace "github.com/upbound/provider-azure/internal/controller/cluster/eventhub/eventhubnamespace"
	namespaceauthorizationrule "github.com/upbound/provider-azure/internal/controller/cluster/eventhub/namespaceauthorizationrule"
	namespacedisasterrecoveryconfig "github.com/upbound/provider-azure/internal/controller/cluster/eventhub/namespacedisasterrecoveryconfig"
	namespaceschemagroup "github.com/upbound/provider-azure/internal/controller/cluster/eventhub/namespaceschemagroup"
)

var eventhubCrdGroup = "eventhub.azure.upbound.io"

// Setup_eventhub creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_eventhub(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: eventhubCrdGroup, Kind: "AuthorizationRule"}:               authorizationrule.Setup,
		schema.GroupKind{Group: eventhubCrdGroup, Kind: "ConsumerGroup"}:                   consumergroup.Setup,
		schema.GroupKind{Group: eventhubCrdGroup, Kind: "EventHub"}:                        eventhub.Setup,
		schema.GroupKind{Group: eventhubCrdGroup, Kind: "EventHubNamespace"}:               eventhubnamespace.Setup,
		schema.GroupKind{Group: eventhubCrdGroup, Kind: "NamespaceAuthorizationRule"}:      namespaceauthorizationrule.Setup,
		schema.GroupKind{Group: eventhubCrdGroup, Kind: "NamespaceDisasterRecoveryConfig"}: namespacedisasterrecoveryconfig.Setup,
		schema.GroupKind{Group: eventhubCrdGroup, Kind: "NamespaceSchemaGroup"}:            namespaceschemagroup.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, eventhubCrdGroup, o); err != nil {
		return err
	}
	return nil
}
