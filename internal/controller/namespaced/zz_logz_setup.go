// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	monitor "github.com/upbound/provider-azure/internal/controller/namespaced/logz/monitor"
	subaccount "github.com/upbound/provider-azure/internal/controller/namespaced/logz/subaccount"
	subaccounttagrule "github.com/upbound/provider-azure/internal/controller/namespaced/logz/subaccounttagrule"
	tagrule "github.com/upbound/provider-azure/internal/controller/namespaced/logz/tagrule"
)

var logzCrdGroup = "logz.azure.m.upbound.io"

// Setup_logz creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_logz(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: logzCrdGroup, Kind: "Monitor"}:           monitor.Setup,
		schema.GroupKind{Group: logzCrdGroup, Kind: "SubAccount"}:        subaccount.Setup,
		schema.GroupKind{Group: logzCrdGroup, Kind: "SubAccountTagRule"}: subaccounttagrule.Setup,
		schema.GroupKind{Group: logzCrdGroup, Kind: "TagRule"}:           tagrule.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, logzCrdGroup, o); err != nil {
		return err
	}
	return nil
}
