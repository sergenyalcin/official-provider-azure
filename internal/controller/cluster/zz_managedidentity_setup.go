// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	federatedidentitycredential "github.com/upbound/provider-azure/internal/controller/cluster/managedidentity/federatedidentitycredential"
	userassignedidentity "github.com/upbound/provider-azure/internal/controller/cluster/managedidentity/userassignedidentity"
)

var managedidentityCrdGroup = "managedidentity.azure.upbound.io"

// Setup_managedidentity creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_managedidentity(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: managedidentityCrdGroup, Kind: "FederatedIdentityCredential"}: federatedidentitycredential.Setup,
		schema.GroupKind{Group: managedidentityCrdGroup, Kind: "UserAssignedIdentity"}:        userassignedidentity.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, managedidentityCrdGroup, o); err != nil {
		return err
	}
	return nil
}
