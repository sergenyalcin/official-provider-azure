// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	applicationinsights "github.com/upbound/provider-azure/internal/controller/namespaced/insights/applicationinsights"
	applicationinsightsanalyticsitem "github.com/upbound/provider-azure/internal/controller/namespaced/insights/applicationinsightsanalyticsitem"
	applicationinsightsapikey "github.com/upbound/provider-azure/internal/controller/namespaced/insights/applicationinsightsapikey"
	applicationinsightssmartdetectionrule "github.com/upbound/provider-azure/internal/controller/namespaced/insights/applicationinsightssmartdetectionrule"
	applicationinsightsstandardwebtest "github.com/upbound/provider-azure/internal/controller/namespaced/insights/applicationinsightsstandardwebtest"
	applicationinsightswebtest "github.com/upbound/provider-azure/internal/controller/namespaced/insights/applicationinsightswebtest"
	applicationinsightsworkbook "github.com/upbound/provider-azure/internal/controller/namespaced/insights/applicationinsightsworkbook"
	applicationinsightsworkbooktemplate "github.com/upbound/provider-azure/internal/controller/namespaced/insights/applicationinsightsworkbooktemplate"
	monitoractiongroup "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitoractiongroup"
	monitoractivitylogalert "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitoractivitylogalert"
	monitorautoscalesetting "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitorautoscalesetting"
	monitordatacollectionendpoint "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitordatacollectionendpoint"
	monitordatacollectionrule "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitordatacollectionrule"
	monitordatacollectionruleassociation "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitordatacollectionruleassociation"
	monitordiagnosticsetting "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitordiagnosticsetting"
	monitormetricalert "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitormetricalert"
	monitorprivatelinkscope "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitorprivatelinkscope"
	monitorprivatelinkscopedservice "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitorprivatelinkscopedservice"
	monitorscheduledqueryrulesalert "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitorscheduledqueryrulesalert"
	monitorscheduledqueryrulesalertv2 "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitorscheduledqueryrulesalertv2"
	monitorscheduledqueryruleslog "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitorscheduledqueryruleslog"
)

var insightsCrdGroup = "insights.azure.m.upbound.io"

// Setup_insights creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_insights(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: insightsCrdGroup, Kind: "ApplicationInsights"}:                   applicationinsights.Setup,
		schema.GroupKind{Group: insightsCrdGroup, Kind: "ApplicationInsightsAPIKey"}:             applicationinsightsapikey.Setup,
		schema.GroupKind{Group: insightsCrdGroup, Kind: "ApplicationInsightsAnalyticsItem"}:      applicationinsightsanalyticsitem.Setup,
		schema.GroupKind{Group: insightsCrdGroup, Kind: "ApplicationInsightsSmartDetectionRule"}: applicationinsightssmartdetectionrule.Setup,
		schema.GroupKind{Group: insightsCrdGroup, Kind: "ApplicationInsightsStandardWebTest"}:    applicationinsightsstandardwebtest.Setup,
		schema.GroupKind{Group: insightsCrdGroup, Kind: "ApplicationInsightsWebTest"}:            applicationinsightswebtest.Setup,
		schema.GroupKind{Group: insightsCrdGroup, Kind: "ApplicationInsightsWorkbook"}:           applicationinsightsworkbook.Setup,
		schema.GroupKind{Group: insightsCrdGroup, Kind: "ApplicationInsightsWorkbookTemplate"}:   applicationinsightsworkbooktemplate.Setup,
		schema.GroupKind{Group: insightsCrdGroup, Kind: "MonitorActionGroup"}:                    monitoractiongroup.Setup,
		schema.GroupKind{Group: insightsCrdGroup, Kind: "MonitorActivityLogAlert"}:               monitoractivitylogalert.Setup,
		schema.GroupKind{Group: insightsCrdGroup, Kind: "MonitorAutoscaleSetting"}:               monitorautoscalesetting.Setup,
		schema.GroupKind{Group: insightsCrdGroup, Kind: "MonitorDataCollectionEndpoint"}:         monitordatacollectionendpoint.Setup,
		schema.GroupKind{Group: insightsCrdGroup, Kind: "MonitorDataCollectionRule"}:             monitordatacollectionrule.Setup,
		schema.GroupKind{Group: insightsCrdGroup, Kind: "MonitorDataCollectionRuleAssociation"}:  monitordatacollectionruleassociation.Setup,
		schema.GroupKind{Group: insightsCrdGroup, Kind: "MonitorDiagnosticSetting"}:              monitordiagnosticsetting.Setup,
		schema.GroupKind{Group: insightsCrdGroup, Kind: "MonitorMetricAlert"}:                    monitormetricalert.Setup,
		schema.GroupKind{Group: insightsCrdGroup, Kind: "MonitorPrivateLinkScope"}:               monitorprivatelinkscope.Setup,
		schema.GroupKind{Group: insightsCrdGroup, Kind: "MonitorPrivateLinkScopedService"}:       monitorprivatelinkscopedservice.Setup,
		schema.GroupKind{Group: insightsCrdGroup, Kind: "MonitorScheduledQueryRulesAlert"}:       monitorscheduledqueryrulesalert.Setup,
		schema.GroupKind{Group: insightsCrdGroup, Kind: "MonitorScheduledQueryRulesAlertV2"}:     monitorscheduledqueryrulesalertv2.Setup,
		schema.GroupKind{Group: insightsCrdGroup, Kind: "MonitorScheduledQueryRulesLog"}:         monitorscheduledqueryruleslog.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, insightsCrdGroup, o); err != nil {
		return err
	}
	return nil
}
