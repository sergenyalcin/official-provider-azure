// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	service "github.com/upbound/provider-azure/internal/controller/cluster/search/service"
	sharedprivatelinkservice "github.com/upbound/provider-azure/internal/controller/cluster/search/sharedprivatelinkservice"
)

var searchCrdGroup = "search.azure.upbound.io"

// Setup_search creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_search(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: searchCrdGroup, Kind: "Service"}:                  service.Setup,
		schema.GroupKind{Group: searchCrdGroup, Kind: "SharedPrivateLinkService"}: sharedprivatelinkservice.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, searchCrdGroup, o); err != nil {
		return err
	}
	return nil
}
