// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	monitoractionruleactiongroup "github.com/upbound/provider-azure/internal/controller/namespaced/alertsmanagement/monitoractionruleactiongroup"
	monitoractionrulesuppression "github.com/upbound/provider-azure/internal/controller/namespaced/alertsmanagement/monitoractionrulesuppression"
	monitoralertprocessingruleactiongroup "github.com/upbound/provider-azure/internal/controller/namespaced/alertsmanagement/monitoralertprocessingruleactiongroup"
	monitoralertprocessingrulesuppression "github.com/upbound/provider-azure/internal/controller/namespaced/alertsmanagement/monitoralertprocessingrulesuppression"
	monitorsmartdetectoralertrule "github.com/upbound/provider-azure/internal/controller/namespaced/alertsmanagement/monitorsmartdetectoralertrule"
)

var alertsmanagementCrdGroup = "alertsmanagement.azure.m.upbound.io"

// Setup_alertsmanagement creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_alertsmanagement(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: alertsmanagementCrdGroup, Kind: "MonitorActionRuleActionGroup"}:          monitoractionruleactiongroup.Setup,
		schema.GroupKind{Group: alertsmanagementCrdGroup, Kind: "MonitorActionRuleSuppression"}:          monitoractionrulesuppression.Setup,
		schema.GroupKind{Group: alertsmanagementCrdGroup, Kind: "MonitorAlertProcessingRuleActionGroup"}: monitoralertprocessingruleactiongroup.Setup,
		schema.GroupKind{Group: alertsmanagementCrdGroup, Kind: "MonitorAlertProcessingRuleSuppression"}: monitoralertprocessingrulesuppression.Setup,
		schema.GroupKind{Group: alertsmanagementCrdGroup, Kind: "MonitorSmartDetectorAlertRule"}:         monitorsmartdetectoralertrule.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
