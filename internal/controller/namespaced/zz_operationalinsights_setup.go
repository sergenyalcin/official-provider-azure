// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	loganalyticsdataexportrule "github.com/upbound/provider-azure/internal/controller/namespaced/operationalinsights/loganalyticsdataexportrule"
	loganalyticsdatasourcewindowsevent "github.com/upbound/provider-azure/internal/controller/namespaced/operationalinsights/loganalyticsdatasourcewindowsevent"
	loganalyticsdatasourcewindowsperformancecounter "github.com/upbound/provider-azure/internal/controller/namespaced/operationalinsights/loganalyticsdatasourcewindowsperformancecounter"
	loganalyticslinkedservice "github.com/upbound/provider-azure/internal/controller/namespaced/operationalinsights/loganalyticslinkedservice"
	loganalyticslinkedstorageaccount "github.com/upbound/provider-azure/internal/controller/namespaced/operationalinsights/loganalyticslinkedstorageaccount"
	loganalyticsquerypack "github.com/upbound/provider-azure/internal/controller/namespaced/operationalinsights/loganalyticsquerypack"
	loganalyticsquerypackquery "github.com/upbound/provider-azure/internal/controller/namespaced/operationalinsights/loganalyticsquerypackquery"
	loganalyticssavedsearch "github.com/upbound/provider-azure/internal/controller/namespaced/operationalinsights/loganalyticssavedsearch"
	workspace "github.com/upbound/provider-azure/internal/controller/namespaced/operationalinsights/workspace"
)

var operationalinsightsCrdGroup = "operationalinsights.azure.m.upbound.io"

// Setup_operationalinsights creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_operationalinsights(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: operationalinsightsCrdGroup, Kind: "LogAnalyticsDataExportRule"}:                      loganalyticsdataexportrule.Setup,
		schema.GroupKind{Group: operationalinsightsCrdGroup, Kind: "LogAnalyticsDataSourceWindowsEvent"}:              loganalyticsdatasourcewindowsevent.Setup,
		schema.GroupKind{Group: operationalinsightsCrdGroup, Kind: "LogAnalyticsDataSourceWindowsPerformanceCounter"}: loganalyticsdatasourcewindowsperformancecounter.Setup,
		schema.GroupKind{Group: operationalinsightsCrdGroup, Kind: "LogAnalyticsLinkedService"}:                       loganalyticslinkedservice.Setup,
		schema.GroupKind{Group: operationalinsightsCrdGroup, Kind: "LogAnalyticsLinkedStorageAccount"}:                loganalyticslinkedstorageaccount.Setup,
		schema.GroupKind{Group: operationalinsightsCrdGroup, Kind: "LogAnalyticsQueryPack"}:                           loganalyticsquerypack.Setup,
		schema.GroupKind{Group: operationalinsightsCrdGroup, Kind: "LogAnalyticsQueryPackQuery"}:                      loganalyticsquerypackquery.Setup,
		schema.GroupKind{Group: operationalinsightsCrdGroup, Kind: "LogAnalyticsSavedSearch"}:                         loganalyticssavedsearch.Setup,
		schema.GroupKind{Group: operationalinsightsCrdGroup, Kind: "Workspace"}:                                       workspace.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
