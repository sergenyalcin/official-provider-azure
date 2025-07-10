// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	endpoint "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/endpoint"
	frontdoorcustomdomain "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/frontdoorcustomdomain"
	frontdoorcustomdomainassociation "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/frontdoorcustomdomainassociation"
	frontdoorendpoint "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/frontdoorendpoint"
	frontdoorfirewallpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/frontdoorfirewallpolicy"
	frontdoororigin "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/frontdoororigin"
	frontdoororigingroup "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/frontdoororigingroup"
	frontdoorprofile "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/frontdoorprofile"
	frontdoorroute "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/frontdoorroute"
	frontdoorrule "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/frontdoorrule"
	frontdoorruleset "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/frontdoorruleset"
	frontdoorsecuritypolicy "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/frontdoorsecuritypolicy"
	profile "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/profile"
)

var cdnCrdGroup = "cdn.azure.m.upbound.io"

// Setup_cdn creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cdn(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: cdnCrdGroup, Kind: "Endpoint"}:                         endpoint.Setup,
		schema.GroupKind{Group: cdnCrdGroup, Kind: "FrontdoorCustomDomain"}:            frontdoorcustomdomain.Setup,
		schema.GroupKind{Group: cdnCrdGroup, Kind: "FrontdoorCustomDomainAssociation"}: frontdoorcustomdomainassociation.Setup,
		schema.GroupKind{Group: cdnCrdGroup, Kind: "FrontdoorEndpoint"}:                frontdoorendpoint.Setup,
		schema.GroupKind{Group: cdnCrdGroup, Kind: "FrontdoorFirewallPolicy"}:          frontdoorfirewallpolicy.Setup,
		schema.GroupKind{Group: cdnCrdGroup, Kind: "FrontdoorOrigin"}:                  frontdoororigin.Setup,
		schema.GroupKind{Group: cdnCrdGroup, Kind: "FrontdoorOriginGroup"}:             frontdoororigingroup.Setup,
		schema.GroupKind{Group: cdnCrdGroup, Kind: "FrontdoorProfile"}:                 frontdoorprofile.Setup,
		schema.GroupKind{Group: cdnCrdGroup, Kind: "FrontdoorRoute"}:                   frontdoorroute.Setup,
		schema.GroupKind{Group: cdnCrdGroup, Kind: "FrontdoorRule"}:                    frontdoorrule.Setup,
		schema.GroupKind{Group: cdnCrdGroup, Kind: "FrontdoorRuleSet"}:                 frontdoorruleset.Setup,
		schema.GroupKind{Group: cdnCrdGroup, Kind: "FrontdoorSecurityPolicy"}:          frontdoorsecuritypolicy.Setup,
		schema.GroupKind{Group: cdnCrdGroup, Kind: "Profile"}:                          profile.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, cdnCrdGroup, o); err != nil {
		return err
	}
	return nil
}
