// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	eventrelaynamespace "github.com/upbound/provider-azure/internal/controller/cluster/relay/eventrelaynamespace"
	hybridconnection "github.com/upbound/provider-azure/internal/controller/cluster/relay/hybridconnection"
	hybridconnectionauthorizationrule "github.com/upbound/provider-azure/internal/controller/cluster/relay/hybridconnectionauthorizationrule"
	namespaceauthorizationrule "github.com/upbound/provider-azure/internal/controller/cluster/relay/namespaceauthorizationrule"
)

var relayCrdGroup = "relay.azure.upbound.io"

// Setup_relay creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_relay(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: relayCrdGroup, Kind: "EventRelayNamespace"}:               eventrelaynamespace.Setup,
		schema.GroupKind{Group: relayCrdGroup, Kind: "HybridConnection"}:                  hybridconnection.Setup,
		schema.GroupKind{Group: relayCrdGroup, Kind: "HybridConnectionAuthorizationRule"}: hybridconnectionauthorizationrule.Setup,
		schema.GroupKind{Group: relayCrdGroup, Kind: "NamespaceAuthorizationRule"}:        namespaceauthorizationrule.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, relayCrdGroup, o); err != nil {
		return err
	}
	return nil
}
