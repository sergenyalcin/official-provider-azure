// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	account "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/account"
	cassandracluster "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/cassandracluster"
	cassandradatacenter "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/cassandradatacenter"
	cassandrakeyspace "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/cassandrakeyspace"
	cassandratable "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/cassandratable"
	gremlindatabase "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/gremlindatabase"
	gremlingraph "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/gremlingraph"
	mongocollection "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/mongocollection"
	mongodatabase "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/mongodatabase"
	mongoroledefinition "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/mongoroledefinition"
	mongouserdefinition "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/mongouserdefinition"
	sqlcontainer "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/sqlcontainer"
	sqldatabase "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/sqldatabase"
	sqldedicatedgateway "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/sqldedicatedgateway"
	sqlfunction "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/sqlfunction"
	sqlroleassignment "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/sqlroleassignment"
	sqlroledefinition "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/sqlroledefinition"
	sqlstoredprocedure "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/sqlstoredprocedure"
	sqltrigger "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/sqltrigger"
	table "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/table"
)

var cosmosdbCrdGroup = "cosmosdb.azure.m.upbound.io"

// Setup_cosmosdb creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cosmosdb(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: cosmosdbCrdGroup, Kind: "Account"}:             account.Setup,
		schema.GroupKind{Group: cosmosdbCrdGroup, Kind: "CassandraCluster"}:    cassandracluster.Setup,
		schema.GroupKind{Group: cosmosdbCrdGroup, Kind: "CassandraDatacenter"}: cassandradatacenter.Setup,
		schema.GroupKind{Group: cosmosdbCrdGroup, Kind: "CassandraKeySpace"}:   cassandrakeyspace.Setup,
		schema.GroupKind{Group: cosmosdbCrdGroup, Kind: "CassandraTable"}:      cassandratable.Setup,
		schema.GroupKind{Group: cosmosdbCrdGroup, Kind: "GremlinDatabase"}:     gremlindatabase.Setup,
		schema.GroupKind{Group: cosmosdbCrdGroup, Kind: "GremlinGraph"}:        gremlingraph.Setup,
		schema.GroupKind{Group: cosmosdbCrdGroup, Kind: "MongoCollection"}:     mongocollection.Setup,
		schema.GroupKind{Group: cosmosdbCrdGroup, Kind: "MongoDatabase"}:       mongodatabase.Setup,
		schema.GroupKind{Group: cosmosdbCrdGroup, Kind: "MongoRoleDefinition"}: mongoroledefinition.Setup,
		schema.GroupKind{Group: cosmosdbCrdGroup, Kind: "MongoUserDefinition"}: mongouserdefinition.Setup,
		schema.GroupKind{Group: cosmosdbCrdGroup, Kind: "SQLContainer"}:        sqlcontainer.Setup,
		schema.GroupKind{Group: cosmosdbCrdGroup, Kind: "SQLDatabase"}:         sqldatabase.Setup,
		schema.GroupKind{Group: cosmosdbCrdGroup, Kind: "SQLDedicatedGateway"}: sqldedicatedgateway.Setup,
		schema.GroupKind{Group: cosmosdbCrdGroup, Kind: "SQLFunction"}:         sqlfunction.Setup,
		schema.GroupKind{Group: cosmosdbCrdGroup, Kind: "SQLRoleAssignment"}:   sqlroleassignment.Setup,
		schema.GroupKind{Group: cosmosdbCrdGroup, Kind: "SQLRoleDefinition"}:   sqlroledefinition.Setup,
		schema.GroupKind{Group: cosmosdbCrdGroup, Kind: "SQLStoredProcedure"}:  sqlstoredprocedure.Setup,
		schema.GroupKind{Group: cosmosdbCrdGroup, Kind: "SQLTrigger"}:          sqltrigger.Setup,
		schema.GroupKind{Group: cosmosdbCrdGroup, Kind: "Table"}:               table.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
