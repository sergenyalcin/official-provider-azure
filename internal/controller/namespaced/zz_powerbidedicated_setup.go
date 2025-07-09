// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	powerbiembedded "github.com/upbound/provider-azure/internal/controller/namespaced/powerbidedicated/powerbiembedded"
)

var powerbidedicatedCrdGroup = "powerbidedicated.azure.m.upbound.io"

// Setup_powerbidedicated creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_powerbidedicated(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: powerbidedicatedCrdGroup, Kind: "PowerBIEmbedded"}: powerbiembedded.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
