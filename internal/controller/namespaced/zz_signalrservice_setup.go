// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	networkacl "github.com/upbound/provider-azure/internal/controller/namespaced/signalrservice/networkacl"
	service "github.com/upbound/provider-azure/internal/controller/namespaced/signalrservice/service"
	signalrsharedprivatelinkresource "github.com/upbound/provider-azure/internal/controller/namespaced/signalrservice/signalrsharedprivatelinkresource"
	webpubsub "github.com/upbound/provider-azure/internal/controller/namespaced/signalrservice/webpubsub"
	webpubsubhub "github.com/upbound/provider-azure/internal/controller/namespaced/signalrservice/webpubsubhub"
	webpubsubnetworkacl "github.com/upbound/provider-azure/internal/controller/namespaced/signalrservice/webpubsubnetworkacl"
)

var signalrserviceCrdGroup = "signalrservice.azure.m.upbound.io"

// Setup_signalrservice creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_signalrservice(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: signalrserviceCrdGroup, Kind: "NetworkACL"}:                       networkacl.Setup,
		schema.GroupKind{Group: signalrserviceCrdGroup, Kind: "Service"}:                          service.Setup,
		schema.GroupKind{Group: signalrserviceCrdGroup, Kind: "SignalrSharedPrivateLinkResource"}: signalrsharedprivatelinkresource.Setup,
		schema.GroupKind{Group: signalrserviceCrdGroup, Kind: "WebPubsub"}:                        webpubsub.Setup,
		schema.GroupKind{Group: signalrserviceCrdGroup, Kind: "WebPubsubHub"}:                     webpubsubhub.Setup,
		schema.GroupKind{Group: signalrserviceCrdGroup, Kind: "WebPubsubNetworkACL"}:              webpubsubnetworkacl.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
