// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	sentinelalertrulefusion "github.com/upbound/provider-azure/internal/controller/namespaced/securityinsights/sentinelalertrulefusion"
	sentinelalertrulemachinelearningbehavioranalytics "github.com/upbound/provider-azure/internal/controller/namespaced/securityinsights/sentinelalertrulemachinelearningbehavioranalytics"
	sentinelalertrulemssecurityincident "github.com/upbound/provider-azure/internal/controller/namespaced/securityinsights/sentinelalertrulemssecurityincident"
	sentinelautomationrule "github.com/upbound/provider-azure/internal/controller/namespaced/securityinsights/sentinelautomationrule"
	sentineldataconnectoriot "github.com/upbound/provider-azure/internal/controller/namespaced/securityinsights/sentineldataconnectoriot"
	sentinelloganalyticsworkspaceonboarding "github.com/upbound/provider-azure/internal/controller/namespaced/securityinsights/sentinelloganalyticsworkspaceonboarding"
	sentinelwatchlist "github.com/upbound/provider-azure/internal/controller/namespaced/securityinsights/sentinelwatchlist"
)

var securityinsightsCrdGroup = "securityinsights.azure.m.upbound.io"

// Setup_securityinsights creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_securityinsights(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: securityinsightsCrdGroup, Kind: "SentinelAlertRuleFusion"}:                           sentinelalertrulefusion.Setup,
		schema.GroupKind{Group: securityinsightsCrdGroup, Kind: "SentinelAlertRuleMSSecurityIncident"}:               sentinelalertrulemssecurityincident.Setup,
		schema.GroupKind{Group: securityinsightsCrdGroup, Kind: "SentinelAlertRuleMachineLearningBehaviorAnalytics"}: sentinelalertrulemachinelearningbehavioranalytics.Setup,
		schema.GroupKind{Group: securityinsightsCrdGroup, Kind: "SentinelAutomationRule"}:                            sentinelautomationrule.Setup,
		schema.GroupKind{Group: securityinsightsCrdGroup, Kind: "SentinelDataConnectorIOT"}:                          sentineldataconnectoriot.Setup,
		schema.GroupKind{Group: securityinsightsCrdGroup, Kind: "SentinelLogAnalyticsWorkspaceOnboarding"}:           sentinelloganalyticsworkspaceonboarding.Setup,
		schema.GroupKind{Group: securityinsightsCrdGroup, Kind: "SentinelWatchlist"}:                                 sentinelwatchlist.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, securityinsightsCrdGroup, o); err != nil {
		return err
	}
	return nil
}
