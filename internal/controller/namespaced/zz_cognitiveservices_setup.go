// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	account "github.com/upbound/provider-azure/internal/controller/namespaced/cognitiveservices/account"
	deployment "github.com/upbound/provider-azure/internal/controller/namespaced/cognitiveservices/deployment"
)

var cognitiveservicesCrdGroup = "cognitiveservices.azure.m.upbound.io"

// Setup_cognitiveservices creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cognitiveservices(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: cognitiveservicesCrdGroup, Kind: "Account"}:    account.Setup,
		schema.GroupKind{Group: cognitiveservicesCrdGroup, Kind: "Deployment"}: deployment.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, cognitiveservicesCrdGroup, o); err != nil {
		return err
	}
	return nil
}
