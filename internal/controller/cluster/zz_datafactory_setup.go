// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	customdataset "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/customdataset"
	dataflow "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/dataflow"
	datasetazureblob "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/datasetazureblob"
	datasetbinary "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/datasetbinary"
	datasetcosmosdbsqlapi "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/datasetcosmosdbsqlapi"
	datasetdelimitedtext "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/datasetdelimitedtext"
	datasethttp "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/datasethttp"
	datasetjson "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/datasetjson"
	datasetmysql "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/datasetmysql"
	datasetparquet "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/datasetparquet"
	datasetpostgresql "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/datasetpostgresql"
	datasetsnowflake "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/datasetsnowflake"
	datasetsqlservertable "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/datasetsqlservertable"
	factory "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/factory"
	integrationruntimeazure "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/integrationruntimeazure"
	integrationruntimeazuressis "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/integrationruntimeazuressis"
	integrationruntimemanaged "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/integrationruntimemanaged"
	integrationruntimeselfhosted "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/integrationruntimeselfhosted"
	linkedcustomservice "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedcustomservice"
	linkedserviceazureblobstorage "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedserviceazureblobstorage"
	linkedserviceazuredatabricks "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedserviceazuredatabricks"
	linkedserviceazurefilestorage "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedserviceazurefilestorage"
	linkedserviceazurefunction "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedserviceazurefunction"
	linkedserviceazuresearch "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedserviceazuresearch"
	linkedserviceazuresqldatabase "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedserviceazuresqldatabase"
	linkedserviceazuretablestorage "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedserviceazuretablestorage"
	linkedservicecosmosdb "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedservicecosmosdb"
	linkedservicecosmosdbmongoapi "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedservicecosmosdbmongoapi"
	linkedservicedatalakestoragegen2 "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedservicedatalakestoragegen2"
	linkedservicekeyvault "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedservicekeyvault"
	linkedservicekusto "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedservicekusto"
	linkedservicemysql "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedservicemysql"
	linkedserviceodata "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedserviceodata"
	linkedserviceodbc "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedserviceodbc"
	linkedservicepostgresql "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedservicepostgresql"
	linkedservicesftp "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedservicesftp"
	linkedservicesnowflake "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedservicesnowflake"
	linkedservicesqlserver "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedservicesqlserver"
	linkedservicesynapse "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedservicesynapse"
	linkedserviceweb "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedserviceweb"
	managedprivateendpoint "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/managedprivateendpoint"
	pipeline "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/pipeline"
	triggerblobevent "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/triggerblobevent"
	triggercustomevent "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/triggercustomevent"
	triggerschedule "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/triggerschedule"
)

var datafactoryCrdGroup = "datafactory.azure.upbound.io"

// Setup_datafactory creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_datafactory(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "CustomDataSet"}:                    customdataset.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "DataFlow"}:                         dataflow.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "DataSetAzureBlob"}:                 datasetazureblob.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "DataSetBinary"}:                    datasetbinary.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "DataSetCosmosDBSQLAPI"}:            datasetcosmosdbsqlapi.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "DataSetDelimitedText"}:             datasetdelimitedtext.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "DataSetHTTP"}:                      datasethttp.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "DataSetJSON"}:                      datasetjson.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "DataSetMySQL"}:                     datasetmysql.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "DataSetParquet"}:                   datasetparquet.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "DataSetPostgreSQL"}:                datasetpostgresql.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "DataSetSQLServerTable"}:            datasetsqlservertable.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "DataSetSnowflake"}:                 datasetsnowflake.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "Factory"}:                          factory.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "IntegrationRuntimeAzure"}:          integrationruntimeazure.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "IntegrationRuntimeAzureSSIS"}:      integrationruntimeazuressis.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "IntegrationRuntimeManaged"}:        integrationruntimemanaged.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "IntegrationRuntimeSelfHosted"}:     integrationruntimeselfhosted.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "LinkedCustomService"}:              linkedcustomservice.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "LinkedServiceAzureBlobStorage"}:    linkedserviceazureblobstorage.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "LinkedServiceAzureDatabricks"}:     linkedserviceazuredatabricks.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "LinkedServiceAzureFileStorage"}:    linkedserviceazurefilestorage.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "LinkedServiceAzureFunction"}:       linkedserviceazurefunction.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "LinkedServiceAzureSQLDatabase"}:    linkedserviceazuresqldatabase.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "LinkedServiceAzureSearch"}:         linkedserviceazuresearch.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "LinkedServiceAzureTableStorage"}:   linkedserviceazuretablestorage.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "LinkedServiceCosmosDB"}:            linkedservicecosmosdb.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "LinkedServiceCosmosDBMongoapi"}:    linkedservicecosmosdbmongoapi.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "LinkedServiceDataLakeStorageGen2"}: linkedservicedatalakestoragegen2.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "LinkedServiceKeyVault"}:            linkedservicekeyvault.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "LinkedServiceKusto"}:               linkedservicekusto.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "LinkedServiceMySQL"}:               linkedservicemysql.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "LinkedServiceOData"}:               linkedserviceodata.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "LinkedServiceOdbc"}:                linkedserviceodbc.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "LinkedServicePostgreSQL"}:          linkedservicepostgresql.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "LinkedServiceSFTP"}:                linkedservicesftp.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "LinkedServiceSQLServer"}:           linkedservicesqlserver.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "LinkedServiceSnowflake"}:           linkedservicesnowflake.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "LinkedServiceSynapse"}:             linkedservicesynapse.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "LinkedServiceWeb"}:                 linkedserviceweb.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "ManagedPrivateEndpoint"}:           managedprivateendpoint.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "Pipeline"}:                         pipeline.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "TriggerBlobEvent"}:                 triggerblobevent.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "TriggerCustomEvent"}:               triggercustomevent.Setup,
		schema.GroupKind{Group: datafactoryCrdGroup, Kind: "TriggerSchedule"}:                  triggerschedule.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
