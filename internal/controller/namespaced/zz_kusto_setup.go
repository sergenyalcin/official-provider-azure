// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	attacheddatabaseconfiguration "github.com/upbound/provider-azure/internal/controller/namespaced/kusto/attacheddatabaseconfiguration"
	cluster "github.com/upbound/provider-azure/internal/controller/namespaced/kusto/cluster"
	clustermanagedprivateendpoint "github.com/upbound/provider-azure/internal/controller/namespaced/kusto/clustermanagedprivateendpoint"
	clusterprincipalassignment "github.com/upbound/provider-azure/internal/controller/namespaced/kusto/clusterprincipalassignment"
	database "github.com/upbound/provider-azure/internal/controller/namespaced/kusto/database"
	databaseprincipalassignment "github.com/upbound/provider-azure/internal/controller/namespaced/kusto/databaseprincipalassignment"
	eventgriddataconnection "github.com/upbound/provider-azure/internal/controller/namespaced/kusto/eventgriddataconnection"
	eventhubdataconnection "github.com/upbound/provider-azure/internal/controller/namespaced/kusto/eventhubdataconnection"
	iothubdataconnection "github.com/upbound/provider-azure/internal/controller/namespaced/kusto/iothubdataconnection"
)

var kustoCrdGroup = "kusto.azure.m.upbound.io"

// Setup_kusto creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_kusto(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: kustoCrdGroup, Kind: "AttachedDatabaseConfiguration"}: attacheddatabaseconfiguration.Setup,
		schema.GroupKind{Group: kustoCrdGroup, Kind: "Cluster"}:                       cluster.Setup,
		schema.GroupKind{Group: kustoCrdGroup, Kind: "ClusterManagedPrivateEndpoint"}: clustermanagedprivateendpoint.Setup,
		schema.GroupKind{Group: kustoCrdGroup, Kind: "ClusterPrincipalAssignment"}:    clusterprincipalassignment.Setup,
		schema.GroupKind{Group: kustoCrdGroup, Kind: "Database"}:                      database.Setup,
		schema.GroupKind{Group: kustoCrdGroup, Kind: "DatabasePrincipalAssignment"}:   databaseprincipalassignment.Setup,
		schema.GroupKind{Group: kustoCrdGroup, Kind: "EventGridDataConnection"}:       eventgriddataconnection.Setup,
		schema.GroupKind{Group: kustoCrdGroup, Kind: "EventHubDataConnection"}:        eventhubdataconnection.Setup,
		schema.GroupKind{Group: kustoCrdGroup, Kind: "IOTHubDataConnection"}:          iothubdataconnection.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, kustoCrdGroup, o); err != nil {
		return err
	}
	return nil
}
