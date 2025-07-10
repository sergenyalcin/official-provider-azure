// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	account "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/account"
	cassandracluster "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/cassandracluster"
	cassandradatacenter "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/cassandradatacenter"
	cassandrakeyspace "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/cassandrakeyspace"
	cassandratable "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/cassandratable"
	gremlindatabase "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/gremlindatabase"
	gremlingraph "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/gremlingraph"
	mongocollection "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/mongocollection"
	mongodatabase "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/mongodatabase"
	mongoroledefinition "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/mongoroledefinition"
	mongouserdefinition "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/mongouserdefinition"
	sqlcontainer "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/sqlcontainer"
	sqldatabase "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/sqldatabase"
	sqldedicatedgateway "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/sqldedicatedgateway"
	sqlfunction "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/sqlfunction"
	sqlroleassignment "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/sqlroleassignment"
	sqlroledefinition "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/sqlroledefinition"
	sqlstoredprocedure "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/sqlstoredprocedure"
	sqltrigger "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/sqltrigger"
	table "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/table"
)

var cosmosdbCrdGroup = "cosmosdb.azure.upbound.io"

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
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, cosmosdbCrdGroup, o); err != nil {
		return err
	}
	return nil
}
