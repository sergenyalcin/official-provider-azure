// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	monitoractionruleactiongroup "github.com/upbound/provider-azure/internal/controller/cluster/alertsmanagement/monitoractionruleactiongroup"
	monitoractionrulesuppression "github.com/upbound/provider-azure/internal/controller/cluster/alertsmanagement/monitoractionrulesuppression"
	monitoralertprocessingruleactiongroup "github.com/upbound/provider-azure/internal/controller/cluster/alertsmanagement/monitoralertprocessingruleactiongroup"
	monitoralertprocessingrulesuppression "github.com/upbound/provider-azure/internal/controller/cluster/alertsmanagement/monitoralertprocessingrulesuppression"
	monitorsmartdetectoralertrule "github.com/upbound/provider-azure/internal/controller/cluster/alertsmanagement/monitorsmartdetectoralertrule"
	api "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/api"
	apidiagnostic "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/apidiagnostic"
	apioperation "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/apioperation"
	apioperationpolicy "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/apioperationpolicy"
	apioperationtag "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/apioperationtag"
	apipolicy "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/apipolicy"
	apirelease "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/apirelease"
	apischema "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/apischema"
	apitag "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/apitag"
	apiversionset "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/apiversionset"
	authorizationserver "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/authorizationserver"
	backend "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/backend"
	diagnostic "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/diagnostic"
	emailtemplate "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/emailtemplate"
	gateway "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/gateway"
	gatewayapi "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/gatewayapi"
	globalschema "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/globalschema"
	identityprovideraad "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/identityprovideraad"
	identityproviderfacebook "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/identityproviderfacebook"
	identityprovidergoogle "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/identityprovidergoogle"
	identityprovidermicrosoft "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/identityprovidermicrosoft"
	identityprovidertwitter "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/identityprovidertwitter"
	logger "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/logger"
	management "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/management"
	namedvalue "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/namedvalue"
	notificationrecipientemail "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/notificationrecipientemail"
	notificationrecipientuser "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/notificationrecipientuser"
	openidconnectprovider "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/openidconnectprovider"
	product "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/product"
	productapi "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/productapi"
	productpolicy "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/productpolicy"
	producttag "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/producttag"
	tag "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/tag"
	user "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/user"
	springcloudaccelerator "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudaccelerator"
	springcloudactivedeployment "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudactivedeployment"
	springcloudapiportal "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudapiportal"
	springcloudapiportalcustomdomain "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudapiportalcustomdomain"
	springcloudapp "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudapp"
	springcloudappcosmosdbassociation "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudappcosmosdbassociation"
	springcloudappmysqlassociation "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudappmysqlassociation"
	springcloudappredisassociation "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudappredisassociation"
	springcloudbuilddeployment "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudbuilddeployment"
	springcloudbuilder "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudbuilder"
	springcloudbuildpackbinding "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudbuildpackbinding"
	springcloudcertificate "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudcertificate"
	springcloudconfigurationservice "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudconfigurationservice"
	springcloudcontainerdeployment "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudcontainerdeployment"
	springcloudcustomdomain "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudcustomdomain"
	springcloudcustomizedaccelerator "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudcustomizedaccelerator"
	springclouddevtoolportal "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springclouddevtoolportal"
	springcloudgateway "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudgateway"
	springcloudgatewaycustomdomain "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudgatewaycustomdomain"
	springcloudjavadeployment "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudjavadeployment"
	springcloudservice "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudservice"
	springcloudstorage "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudstorage"
	provider "github.com/upbound/provider-azure/internal/controller/cluster/attestation/provider"
	managementgrouppolicyassignment "github.com/upbound/provider-azure/internal/controller/cluster/authorization/managementgrouppolicyassignment"
	managementgrouppolicyexemption "github.com/upbound/provider-azure/internal/controller/cluster/authorization/managementgrouppolicyexemption"
	managementlock "github.com/upbound/provider-azure/internal/controller/cluster/authorization/managementlock"
	pimactiveroleassignment "github.com/upbound/provider-azure/internal/controller/cluster/authorization/pimactiveroleassignment"
	pimeligibleroleassignment "github.com/upbound/provider-azure/internal/controller/cluster/authorization/pimeligibleroleassignment"
	policydefinition "github.com/upbound/provider-azure/internal/controller/cluster/authorization/policydefinition"
	policysetdefinition "github.com/upbound/provider-azure/internal/controller/cluster/authorization/policysetdefinition"
	resourcegrouppolicyassignment "github.com/upbound/provider-azure/internal/controller/cluster/authorization/resourcegrouppolicyassignment"
	resourcegrouppolicyexemption "github.com/upbound/provider-azure/internal/controller/cluster/authorization/resourcegrouppolicyexemption"
	resourcepolicyassignment "github.com/upbound/provider-azure/internal/controller/cluster/authorization/resourcepolicyassignment"
	resourcepolicyexemption "github.com/upbound/provider-azure/internal/controller/cluster/authorization/resourcepolicyexemption"
	roledefinition "github.com/upbound/provider-azure/internal/controller/cluster/authorization/roledefinition"
	subscriptionpolicyassignment "github.com/upbound/provider-azure/internal/controller/cluster/authorization/subscriptionpolicyassignment"
	subscriptionpolicyexemption "github.com/upbound/provider-azure/internal/controller/cluster/authorization/subscriptionpolicyexemption"
	trustedaccessrolebinding "github.com/upbound/provider-azure/internal/controller/cluster/authorization/trustedaccessrolebinding"
	connection "github.com/upbound/provider-azure/internal/controller/cluster/automation/connection"
	connectionclassiccertificate "github.com/upbound/provider-azure/internal/controller/cluster/automation/connectionclassiccertificate"
	connectiontype "github.com/upbound/provider-azure/internal/controller/cluster/automation/connectiontype"
	credential "github.com/upbound/provider-azure/internal/controller/cluster/automation/credential"
	hybridrunbookworkergroup "github.com/upbound/provider-azure/internal/controller/cluster/automation/hybridrunbookworkergroup"
	module "github.com/upbound/provider-azure/internal/controller/cluster/automation/module"
	runbook "github.com/upbound/provider-azure/internal/controller/cluster/automation/runbook"
	variablebool "github.com/upbound/provider-azure/internal/controller/cluster/automation/variablebool"
	variabledatetime "github.com/upbound/provider-azure/internal/controller/cluster/automation/variabledatetime"
	variableint "github.com/upbound/provider-azure/internal/controller/cluster/automation/variableint"
	variablestring "github.com/upbound/provider-azure/internal/controller/cluster/automation/variablestring"
	resourcegroup "github.com/upbound/provider-azure/internal/controller/cluster/azure/resourcegroup"
	resourceproviderregistration "github.com/upbound/provider-azure/internal/controller/cluster/azure/resourceproviderregistration"
	botchannelalexa "github.com/upbound/provider-azure/internal/controller/cluster/botservice/botchannelalexa"
	botchanneldirectline "github.com/upbound/provider-azure/internal/controller/cluster/botservice/botchanneldirectline"
	botchannelline "github.com/upbound/provider-azure/internal/controller/cluster/botservice/botchannelline"
	botchannelmsteams "github.com/upbound/provider-azure/internal/controller/cluster/botservice/botchannelmsteams"
	botchannelslack "github.com/upbound/provider-azure/internal/controller/cluster/botservice/botchannelslack"
	botchannelsms "github.com/upbound/provider-azure/internal/controller/cluster/botservice/botchannelsms"
	botchannelsregistration "github.com/upbound/provider-azure/internal/controller/cluster/botservice/botchannelsregistration"
	botchannelwebchat "github.com/upbound/provider-azure/internal/controller/cluster/botservice/botchannelwebchat"
	botconnection "github.com/upbound/provider-azure/internal/controller/cluster/botservice/botconnection"
	botwebapp "github.com/upbound/provider-azure/internal/controller/cluster/botservice/botwebapp"
	rediscachecache "github.com/upbound/provider-azure/internal/controller/cluster/cache/rediscache"
	rediscacheaccesspolicy "github.com/upbound/provider-azure/internal/controller/cluster/cache/rediscacheaccesspolicy"
	rediscacheaccesspolicyassignment "github.com/upbound/provider-azure/internal/controller/cluster/cache/rediscacheaccesspolicyassignment"
	redisenterprisecluster "github.com/upbound/provider-azure/internal/controller/cluster/cache/redisenterprisecluster"
	redisenterprisedatabase "github.com/upbound/provider-azure/internal/controller/cluster/cache/redisenterprisedatabase"
	redisfirewallrule "github.com/upbound/provider-azure/internal/controller/cluster/cache/redisfirewallrule"
	redislinkedserver "github.com/upbound/provider-azure/internal/controller/cluster/cache/redislinkedserver"
	endpoint "github.com/upbound/provider-azure/internal/controller/cluster/cdn/endpoint"
	frontdoorcustomdomain "github.com/upbound/provider-azure/internal/controller/cluster/cdn/frontdoorcustomdomain"
	frontdoorcustomdomainassociation "github.com/upbound/provider-azure/internal/controller/cluster/cdn/frontdoorcustomdomainassociation"
	frontdoorendpoint "github.com/upbound/provider-azure/internal/controller/cluster/cdn/frontdoorendpoint"
	frontdoororigin "github.com/upbound/provider-azure/internal/controller/cluster/cdn/frontdoororigin"
	frontdoororigingroup "github.com/upbound/provider-azure/internal/controller/cluster/cdn/frontdoororigingroup"
	frontdoorprofile "github.com/upbound/provider-azure/internal/controller/cluster/cdn/frontdoorprofile"
	frontdoorroute "github.com/upbound/provider-azure/internal/controller/cluster/cdn/frontdoorroute"
	frontdoorrule "github.com/upbound/provider-azure/internal/controller/cluster/cdn/frontdoorrule"
	frontdoorruleset "github.com/upbound/provider-azure/internal/controller/cluster/cdn/frontdoorruleset"
	frontdoorsecuritypolicy "github.com/upbound/provider-azure/internal/controller/cluster/cdn/frontdoorsecuritypolicy"
	appservicecertificateorder "github.com/upbound/provider-azure/internal/controller/cluster/certificateregistration/appservicecertificateorder"
	deployment "github.com/upbound/provider-azure/internal/controller/cluster/cognitiveservices/deployment"
	availabilityset "github.com/upbound/provider-azure/internal/controller/cluster/compute/availabilityset"
	capacityreservation "github.com/upbound/provider-azure/internal/controller/cluster/compute/capacityreservation"
	capacityreservationgroup "github.com/upbound/provider-azure/internal/controller/cluster/compute/capacityreservationgroup"
	dedicatedhost "github.com/upbound/provider-azure/internal/controller/cluster/compute/dedicatedhost"
	diskaccess "github.com/upbound/provider-azure/internal/controller/cluster/compute/diskaccess"
	diskencryptionset "github.com/upbound/provider-azure/internal/controller/cluster/compute/diskencryptionset"
	galleryapplication "github.com/upbound/provider-azure/internal/controller/cluster/compute/galleryapplication"
	galleryapplicationversion "github.com/upbound/provider-azure/internal/controller/cluster/compute/galleryapplicationversion"
	image "github.com/upbound/provider-azure/internal/controller/cluster/compute/image"
	linuxvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/cluster/compute/linuxvirtualmachinescaleset"
	manageddisk "github.com/upbound/provider-azure/internal/controller/cluster/compute/manageddisk"
	manageddisksastoken "github.com/upbound/provider-azure/internal/controller/cluster/compute/manageddisksastoken"
	orchestratedvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/cluster/compute/orchestratedvirtualmachinescaleset"
	proximityplacementgroup "github.com/upbound/provider-azure/internal/controller/cluster/compute/proximityplacementgroup"
	sharedimage "github.com/upbound/provider-azure/internal/controller/cluster/compute/sharedimage"
	sharedimagegallery "github.com/upbound/provider-azure/internal/controller/cluster/compute/sharedimagegallery"
	sshpublickey "github.com/upbound/provider-azure/internal/controller/cluster/compute/sshpublickey"
	virtualmachinedatadiskattachment "github.com/upbound/provider-azure/internal/controller/cluster/compute/virtualmachinedatadiskattachment"
	virtualmachineextension "github.com/upbound/provider-azure/internal/controller/cluster/compute/virtualmachineextension"
	virtualmachineruncommand "github.com/upbound/provider-azure/internal/controller/cluster/compute/virtualmachineruncommand"
	windowsvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/cluster/compute/windowsvirtualmachinescaleset"
	ledger "github.com/upbound/provider-azure/internal/controller/cluster/confidentialledger/ledger"
	budgetmanagementgroup "github.com/upbound/provider-azure/internal/controller/cluster/consumption/budgetmanagementgroup"
	budgetresourcegroup "github.com/upbound/provider-azure/internal/controller/cluster/consumption/budgetresourcegroup"
	budgetsubscription "github.com/upbound/provider-azure/internal/controller/cluster/consumption/budgetsubscription"
	containerapp "github.com/upbound/provider-azure/internal/controller/cluster/containerapp/containerapp"
	customdomaincontainerapp "github.com/upbound/provider-azure/internal/controller/cluster/containerapp/customdomain"
	environment "github.com/upbound/provider-azure/internal/controller/cluster/containerapp/environment"
	environmentcertificate "github.com/upbound/provider-azure/internal/controller/cluster/containerapp/environmentcertificate"
	environmentcustomdomain "github.com/upbound/provider-azure/internal/controller/cluster/containerapp/environmentcustomdomain"
	environmentdaprcomponent "github.com/upbound/provider-azure/internal/controller/cluster/containerapp/environmentdaprcomponent"
	environmentstorage "github.com/upbound/provider-azure/internal/controller/cluster/containerapp/environmentstorage"
	agentpool "github.com/upbound/provider-azure/internal/controller/cluster/containerregistry/agentpool"
	containerconnectedregistry "github.com/upbound/provider-azure/internal/controller/cluster/containerregistry/containerconnectedregistry"
	registry "github.com/upbound/provider-azure/internal/controller/cluster/containerregistry/registry"
	scopemap "github.com/upbound/provider-azure/internal/controller/cluster/containerregistry/scopemap"
	token "github.com/upbound/provider-azure/internal/controller/cluster/containerregistry/token"
	tokenpassword "github.com/upbound/provider-azure/internal/controller/cluster/containerregistry/tokenpassword"
	webhookcontainerregistry "github.com/upbound/provider-azure/internal/controller/cluster/containerregistry/webhook"
	kubernetescluster "github.com/upbound/provider-azure/internal/controller/cluster/containerservice/kubernetescluster"
	kubernetesclusterextension "github.com/upbound/provider-azure/internal/controller/cluster/containerservice/kubernetesclusterextension"
	kubernetesclusternodepool "github.com/upbound/provider-azure/internal/controller/cluster/containerservice/kubernetesclusternodepool"
	kubernetesfleetmanager "github.com/upbound/provider-azure/internal/controller/cluster/containerservice/kubernetesfleetmanager"
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
	costanomalyalert "github.com/upbound/provider-azure/internal/controller/cluster/costmanagement/costanomalyalert"
	resourcegroupcostmanagementexport "github.com/upbound/provider-azure/internal/controller/cluster/costmanagement/resourcegroupcostmanagementexport"
	subscriptioncostmanagementexport "github.com/upbound/provider-azure/internal/controller/cluster/costmanagement/subscriptioncostmanagementexport"
	customprovider "github.com/upbound/provider-azure/internal/controller/cluster/customproviders/customprovider"
	device "github.com/upbound/provider-azure/internal/controller/cluster/databoxedge/device"
	accessconnector "github.com/upbound/provider-azure/internal/controller/cluster/databricks/accessconnector"
	workspacecustomermanagedkey "github.com/upbound/provider-azure/internal/controller/cluster/databricks/workspacecustomermanagedkey"
	workspacerootdbfscustomermanagedkey "github.com/upbound/provider-azure/internal/controller/cluster/databricks/workspacerootdbfscustomermanagedkey"
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
	integrationruntimeazuressis "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/integrationruntimeazuressis"
	integrationruntimemanaged "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/integrationruntimemanaged"
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
	pipeline "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/pipeline"
	triggerblobevent "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/triggerblobevent"
	triggercustomevent "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/triggercustomevent"
	triggerschedule "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/triggerschedule"
	databasemigrationproject "github.com/upbound/provider-azure/internal/controller/cluster/datamigration/databasemigrationproject"
	databasemigrationservice "github.com/upbound/provider-azure/internal/controller/cluster/datamigration/databasemigrationservice"
	backupinstanceblobstorage "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backupinstanceblobstorage"
	backupinstancedisk "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backupinstancedisk"
	backupinstancekubernetescluster "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backupinstancekubernetescluster"
	backupinstancepostgresql "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backupinstancepostgresql"
	backuppolicyblobstorage "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backuppolicyblobstorage"
	backuppolicydisk "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backuppolicydisk"
	backuppolicykubernetescluster "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backuppolicykubernetescluster"
	backuppolicypostgresql "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backuppolicypostgresql"
	backupvault "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backupvault"
	resourceguard "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/resourceguard"
	datasetblobstorage "github.com/upbound/provider-azure/internal/controller/cluster/datashare/datasetblobstorage"
	datasetdatalakegen2 "github.com/upbound/provider-azure/internal/controller/cluster/datashare/datasetdatalakegen2"
	datasetkustocluster "github.com/upbound/provider-azure/internal/controller/cluster/datashare/datasetkustocluster"
	datasetkustodatabase "github.com/upbound/provider-azure/internal/controller/cluster/datashare/datasetkustodatabase"
	datashare "github.com/upbound/provider-azure/internal/controller/cluster/datashare/datashare"
	flexibledatabase "github.com/upbound/provider-azure/internal/controller/cluster/dbformysql/flexibledatabase"
	activedirectoryadministratordbforpostgresql "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/activedirectoryadministrator"
	configurationdbforpostgresql "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/configuration"
	flexibleserverdbforpostgresql "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/flexibleserver"
	flexibleserveractivedirectoryadministrator "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/flexibleserveractivedirectoryadministrator"
	flexibleserverconfigurationdbforpostgresql "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/flexibleserverconfiguration"
	flexibleserverdatabase "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/flexibleserverdatabase"
	flexibleserverfirewallruledbforpostgresql "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/flexibleserverfirewallrule"
	serverkey "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/serverkey"
	virtualnetworkruledbforpostgresql "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/virtualnetworkrule"
	iothub "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothub"
	iothubcertificate "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubcertificate"
	iothubconsumergroup "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubconsumergroup"
	iothubdps "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubdps"
	iothubdpscertificate "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubdpscertificate"
	iothubdpssharedaccesspolicy "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubdpssharedaccesspolicy"
	iothubendpointeventhub "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubendpointeventhub"
	iothubendpointservicebusqueue "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubendpointservicebusqueue"
	iothubendpointservicebustopic "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubendpointservicebustopic"
	iothubendpointstoragecontainer "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubendpointstoragecontainer"
	iothubenrichment "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubenrichment"
	iothubfallbackroute "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubfallbackroute"
	iothubroute "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubroute"
	iothubsharedaccesspolicy "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubsharedaccesspolicy"
	iothubdeviceupdateaccount "github.com/upbound/provider-azure/internal/controller/cluster/deviceupdate/iothubdeviceupdateaccount"
	iothubdeviceupdateinstance "github.com/upbound/provider-azure/internal/controller/cluster/deviceupdate/iothubdeviceupdateinstance"
	globalvmshutdownschedule "github.com/upbound/provider-azure/internal/controller/cluster/devtestlab/globalvmshutdownschedule"
	lab "github.com/upbound/provider-azure/internal/controller/cluster/devtestlab/lab"
	linuxvirtualmachinedevtestlab "github.com/upbound/provider-azure/internal/controller/cluster/devtestlab/linuxvirtualmachine"
	policydevtestlab "github.com/upbound/provider-azure/internal/controller/cluster/devtestlab/policy"
	scheduledevtestlab "github.com/upbound/provider-azure/internal/controller/cluster/devtestlab/schedule"
	windowsvirtualmachinedevtestlab "github.com/upbound/provider-azure/internal/controller/cluster/devtestlab/windowsvirtualmachine"
	instance "github.com/upbound/provider-azure/internal/controller/cluster/digitaltwins/instance"
	cloudelasticsearch "github.com/upbound/provider-azure/internal/controller/cluster/elastic/cloudelasticsearch"
	domain "github.com/upbound/provider-azure/internal/controller/cluster/eventgrid/domain"
	domaintopic "github.com/upbound/provider-azure/internal/controller/cluster/eventgrid/domaintopic"
	eventsubscription "github.com/upbound/provider-azure/internal/controller/cluster/eventgrid/eventsubscription"
	systemtopic "github.com/upbound/provider-azure/internal/controller/cluster/eventgrid/systemtopic"
	consumergroup "github.com/upbound/provider-azure/internal/controller/cluster/eventhub/consumergroup"
	eventhub "github.com/upbound/provider-azure/internal/controller/cluster/eventhub/eventhub"
	eventhubnamespace "github.com/upbound/provider-azure/internal/controller/cluster/eventhub/eventhubnamespace"
	namespaceschemagroup "github.com/upbound/provider-azure/internal/controller/cluster/eventhub/namespaceschemagroup"
	serverfluidrelay "github.com/upbound/provider-azure/internal/controller/cluster/fluidrelay/server"
	policyvirtualmachineconfigurationassignment "github.com/upbound/provider-azure/internal/controller/cluster/guestconfiguration/policyvirtualmachineconfigurationassignment"
	hadoopcluster "github.com/upbound/provider-azure/internal/controller/cluster/hdinsight/hadoopcluster"
	hbasecluster "github.com/upbound/provider-azure/internal/controller/cluster/hdinsight/hbasecluster"
	interactivequerycluster "github.com/upbound/provider-azure/internal/controller/cluster/hdinsight/interactivequerycluster"
	kafkacluster "github.com/upbound/provider-azure/internal/controller/cluster/hdinsight/kafkacluster"
	sparkcluster "github.com/upbound/provider-azure/internal/controller/cluster/hdinsight/sparkcluster"
	healthbot "github.com/upbound/provider-azure/internal/controller/cluster/healthbot/healthbot"
	healthcaredicomservice "github.com/upbound/provider-azure/internal/controller/cluster/healthcareapis/healthcaredicomservice"
	healthcarefhirservice "github.com/upbound/provider-azure/internal/controller/cluster/healthcareapis/healthcarefhirservice"
	healthcaremedtechservice "github.com/upbound/provider-azure/internal/controller/cluster/healthcareapis/healthcaremedtechservice"
	healthcaremedtechservicefhirdestination "github.com/upbound/provider-azure/internal/controller/cluster/healthcareapis/healthcaremedtechservicefhirdestination"
	healthcareservice "github.com/upbound/provider-azure/internal/controller/cluster/healthcareapis/healthcareservice"
	healthcareworkspace "github.com/upbound/provider-azure/internal/controller/cluster/healthcareapis/healthcareworkspace"
	applicationinsights "github.com/upbound/provider-azure/internal/controller/cluster/insights/applicationinsights"
	applicationinsightsanalyticsitem "github.com/upbound/provider-azure/internal/controller/cluster/insights/applicationinsightsanalyticsitem"
	applicationinsightsapikey "github.com/upbound/provider-azure/internal/controller/cluster/insights/applicationinsightsapikey"
	applicationinsightssmartdetectionrule "github.com/upbound/provider-azure/internal/controller/cluster/insights/applicationinsightssmartdetectionrule"
	applicationinsightsstandardwebtest "github.com/upbound/provider-azure/internal/controller/cluster/insights/applicationinsightsstandardwebtest"
	applicationinsightswebtest "github.com/upbound/provider-azure/internal/controller/cluster/insights/applicationinsightswebtest"
	applicationinsightsworkbook "github.com/upbound/provider-azure/internal/controller/cluster/insights/applicationinsightsworkbook"
	applicationinsightsworkbooktemplate "github.com/upbound/provider-azure/internal/controller/cluster/insights/applicationinsightsworkbooktemplate"
	monitoractiongroup "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitoractiongroup"
	monitoractivitylogalert "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitoractivitylogalert"
	monitorautoscalesetting "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitorautoscalesetting"
	monitordatacollectionendpoint "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitordatacollectionendpoint"
	monitordatacollectionrule "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitordatacollectionrule"
	monitordatacollectionruleassociation "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitordatacollectionruleassociation"
	monitordiagnosticsetting "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitordiagnosticsetting"
	monitormetricalert "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitormetricalert"
	monitorprivatelinkscope "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitorprivatelinkscope"
	monitorprivatelinkscopedservice "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitorprivatelinkscopedservice"
	monitorscheduledqueryrulesalert "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitorscheduledqueryrulesalert"
	monitorscheduledqueryrulesalertv2 "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitorscheduledqueryrulesalertv2"
	monitorscheduledqueryruleslog "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitorscheduledqueryruleslog"
	application "github.com/upbound/provider-azure/internal/controller/cluster/iotcentral/application"
	applicationnetworkruleset "github.com/upbound/provider-azure/internal/controller/cluster/iotcentral/applicationnetworkruleset"
	accesspolicy "github.com/upbound/provider-azure/internal/controller/cluster/keyvault/accesspolicy"
	certificatekeyvault "github.com/upbound/provider-azure/internal/controller/cluster/keyvault/certificate"
	certificatecontacts "github.com/upbound/provider-azure/internal/controller/cluster/keyvault/certificatecontacts"
	certificateissuer "github.com/upbound/provider-azure/internal/controller/cluster/keyvault/certificateissuer"
	key "github.com/upbound/provider-azure/internal/controller/cluster/keyvault/key"
	managedhardwaresecuritymodule "github.com/upbound/provider-azure/internal/controller/cluster/keyvault/managedhardwaresecuritymodule"
	managedstorageaccount "github.com/upbound/provider-azure/internal/controller/cluster/keyvault/managedstorageaccount"
	managedstorageaccountsastokendefinition "github.com/upbound/provider-azure/internal/controller/cluster/keyvault/managedstorageaccountsastokendefinition"
	secret "github.com/upbound/provider-azure/internal/controller/cluster/keyvault/secret"
	attacheddatabaseconfiguration "github.com/upbound/provider-azure/internal/controller/cluster/kusto/attacheddatabaseconfiguration"
	clustermanagedprivateendpoint "github.com/upbound/provider-azure/internal/controller/cluster/kusto/clustermanagedprivateendpoint"
	clusterprincipalassignment "github.com/upbound/provider-azure/internal/controller/cluster/kusto/clusterprincipalassignment"
	databasekusto "github.com/upbound/provider-azure/internal/controller/cluster/kusto/database"
	databaseprincipalassignment "github.com/upbound/provider-azure/internal/controller/cluster/kusto/databaseprincipalassignment"
	eventgriddataconnection "github.com/upbound/provider-azure/internal/controller/cluster/kusto/eventgriddataconnection"
	eventhubdataconnection "github.com/upbound/provider-azure/internal/controller/cluster/kusto/eventhubdataconnection"
	iothubdataconnection "github.com/upbound/provider-azure/internal/controller/cluster/kusto/iothubdataconnection"
	labservicelab "github.com/upbound/provider-azure/internal/controller/cluster/labservices/labservicelab"
	labserviceplan "github.com/upbound/provider-azure/internal/controller/cluster/labservices/labserviceplan"
	loadtest "github.com/upbound/provider-azure/internal/controller/cluster/loadtestservice/loadtest"
	appactioncustom "github.com/upbound/provider-azure/internal/controller/cluster/logic/appactioncustom"
	appactionhttp "github.com/upbound/provider-azure/internal/controller/cluster/logic/appactionhttp"
	appintegrationaccount "github.com/upbound/provider-azure/internal/controller/cluster/logic/appintegrationaccount"
	appintegrationaccountbatchconfiguration "github.com/upbound/provider-azure/internal/controller/cluster/logic/appintegrationaccountbatchconfiguration"
	appintegrationaccountpartner "github.com/upbound/provider-azure/internal/controller/cluster/logic/appintegrationaccountpartner"
	appintegrationaccountschema "github.com/upbound/provider-azure/internal/controller/cluster/logic/appintegrationaccountschema"
	appintegrationaccountsession "github.com/upbound/provider-azure/internal/controller/cluster/logic/appintegrationaccountsession"
	apptriggercustom "github.com/upbound/provider-azure/internal/controller/cluster/logic/apptriggercustom"
	apptriggerhttprequest "github.com/upbound/provider-azure/internal/controller/cluster/logic/apptriggerhttprequest"
	apptriggerrecurrence "github.com/upbound/provider-azure/internal/controller/cluster/logic/apptriggerrecurrence"
	appworkflow "github.com/upbound/provider-azure/internal/controller/cluster/logic/appworkflow"
	integrationserviceenvironment "github.com/upbound/provider-azure/internal/controller/cluster/logic/integrationserviceenvironment"
	monitor "github.com/upbound/provider-azure/internal/controller/cluster/logz/monitor"
	subaccount "github.com/upbound/provider-azure/internal/controller/cluster/logz/subaccount"
	subaccounttagrule "github.com/upbound/provider-azure/internal/controller/cluster/logz/subaccounttagrule"
	tagrule "github.com/upbound/provider-azure/internal/controller/cluster/logz/tagrule"
	computecluster "github.com/upbound/provider-azure/internal/controller/cluster/machinelearningservices/computecluster"
	computeinstance "github.com/upbound/provider-azure/internal/controller/cluster/machinelearningservices/computeinstance"
	synapsespark "github.com/upbound/provider-azure/internal/controller/cluster/machinelearningservices/synapsespark"
	maintenanceassignmentdedicatedhost "github.com/upbound/provider-azure/internal/controller/cluster/maintenance/maintenanceassignmentdedicatedhost"
	maintenanceassignmentvirtualmachine "github.com/upbound/provider-azure/internal/controller/cluster/maintenance/maintenanceassignmentvirtualmachine"
	maintenanceconfiguration "github.com/upbound/provider-azure/internal/controller/cluster/maintenance/maintenanceconfiguration"
	federatedidentitycredential "github.com/upbound/provider-azure/internal/controller/cluster/managedidentity/federatedidentitycredential"
	userassignedidentity "github.com/upbound/provider-azure/internal/controller/cluster/managedidentity/userassignedidentity"
	managementgroup "github.com/upbound/provider-azure/internal/controller/cluster/management/managementgroup"
	managementgroupsubscriptionassociation "github.com/upbound/provider-azure/internal/controller/cluster/management/managementgroupsubscriptionassociation"
	creator "github.com/upbound/provider-azure/internal/controller/cluster/maps/creator"
	marketplaceagreement "github.com/upbound/provider-azure/internal/controller/cluster/marketplaceordering/marketplaceagreement"
	asset "github.com/upbound/provider-azure/internal/controller/cluster/media/asset"
	assetfilter "github.com/upbound/provider-azure/internal/controller/cluster/media/assetfilter"
	contentkeypolicy "github.com/upbound/provider-azure/internal/controller/cluster/media/contentkeypolicy"
	liveevent "github.com/upbound/provider-azure/internal/controller/cluster/media/liveevent"
	liveeventoutput "github.com/upbound/provider-azure/internal/controller/cluster/media/liveeventoutput"
	servicesaccount "github.com/upbound/provider-azure/internal/controller/cluster/media/servicesaccount"
	servicesaccountfilter "github.com/upbound/provider-azure/internal/controller/cluster/media/servicesaccountfilter"
	streamingendpoint "github.com/upbound/provider-azure/internal/controller/cluster/media/streamingendpoint"
	streaminglocator "github.com/upbound/provider-azure/internal/controller/cluster/media/streaminglocator"
	streamingpolicy "github.com/upbound/provider-azure/internal/controller/cluster/media/streamingpolicy"
	transform "github.com/upbound/provider-azure/internal/controller/cluster/media/transform"
	spatialanchorsaccount "github.com/upbound/provider-azure/internal/controller/cluster/mixedreality/spatialanchorsaccount"
	pool "github.com/upbound/provider-azure/internal/controller/cluster/netapp/pool"
	snapshotnetapp "github.com/upbound/provider-azure/internal/controller/cluster/netapp/snapshot"
	snapshotpolicy "github.com/upbound/provider-azure/internal/controller/cluster/netapp/snapshotpolicy"
	volume "github.com/upbound/provider-azure/internal/controller/cluster/netapp/volume"
	applicationgateway "github.com/upbound/provider-azure/internal/controller/cluster/network/applicationgateway"
	applicationsecuritygroup "github.com/upbound/provider-azure/internal/controller/cluster/network/applicationsecuritygroup"
	bastionhost "github.com/upbound/provider-azure/internal/controller/cluster/network/bastionhost"
	connectionmonitor "github.com/upbound/provider-azure/internal/controller/cluster/network/connectionmonitor"
	ddosprotectionplan "github.com/upbound/provider-azure/internal/controller/cluster/network/ddosprotectionplan"
	dnsaaaarecord "github.com/upbound/provider-azure/internal/controller/cluster/network/dnsaaaarecord"
	dnsarecord "github.com/upbound/provider-azure/internal/controller/cluster/network/dnsarecord"
	dnscaarecord "github.com/upbound/provider-azure/internal/controller/cluster/network/dnscaarecord"
	dnscnamerecord "github.com/upbound/provider-azure/internal/controller/cluster/network/dnscnamerecord"
	dnsmxrecord "github.com/upbound/provider-azure/internal/controller/cluster/network/dnsmxrecord"
	dnsnsrecord "github.com/upbound/provider-azure/internal/controller/cluster/network/dnsnsrecord"
	dnsptrrecord "github.com/upbound/provider-azure/internal/controller/cluster/network/dnsptrrecord"
	dnssrvrecord "github.com/upbound/provider-azure/internal/controller/cluster/network/dnssrvrecord"
	dnstxtrecord "github.com/upbound/provider-azure/internal/controller/cluster/network/dnstxtrecord"
	dnszone "github.com/upbound/provider-azure/internal/controller/cluster/network/dnszone"
	expressroutecircuit "github.com/upbound/provider-azure/internal/controller/cluster/network/expressroutecircuit"
	expressroutecircuitauthorization "github.com/upbound/provider-azure/internal/controller/cluster/network/expressroutecircuitauthorization"
	expressroutecircuitconnection "github.com/upbound/provider-azure/internal/controller/cluster/network/expressroutecircuitconnection"
	expressroutecircuitpeering "github.com/upbound/provider-azure/internal/controller/cluster/network/expressroutecircuitpeering"
	expressrouteconnection "github.com/upbound/provider-azure/internal/controller/cluster/network/expressrouteconnection"
	expressroutegateway "github.com/upbound/provider-azure/internal/controller/cluster/network/expressroutegateway"
	expressrouteport "github.com/upbound/provider-azure/internal/controller/cluster/network/expressrouteport"
	firewall "github.com/upbound/provider-azure/internal/controller/cluster/network/firewall"
	firewallapplicationrulecollection "github.com/upbound/provider-azure/internal/controller/cluster/network/firewallapplicationrulecollection"
	firewallnatrulecollection "github.com/upbound/provider-azure/internal/controller/cluster/network/firewallnatrulecollection"
	firewallnetworkrulecollection "github.com/upbound/provider-azure/internal/controller/cluster/network/firewallnetworkrulecollection"
	firewallpolicy "github.com/upbound/provider-azure/internal/controller/cluster/network/firewallpolicy"
	firewallpolicyrulecollectiongroup "github.com/upbound/provider-azure/internal/controller/cluster/network/firewallpolicyrulecollectiongroup"
	frontdoor "github.com/upbound/provider-azure/internal/controller/cluster/network/frontdoor"
	frontdoorcustomhttpsconfiguration "github.com/upbound/provider-azure/internal/controller/cluster/network/frontdoorcustomhttpsconfiguration"
	frontdoorfirewallpolicynetwork "github.com/upbound/provider-azure/internal/controller/cluster/network/frontdoorfirewallpolicy"
	frontdoorrulesengine "github.com/upbound/provider-azure/internal/controller/cluster/network/frontdoorrulesengine"
	ipgroup "github.com/upbound/provider-azure/internal/controller/cluster/network/ipgroup"
	loadbalancer "github.com/upbound/provider-azure/internal/controller/cluster/network/loadbalancer"
	loadbalancerbackendaddresspool "github.com/upbound/provider-azure/internal/controller/cluster/network/loadbalancerbackendaddresspool"
	loadbalancerbackendaddresspooladdress "github.com/upbound/provider-azure/internal/controller/cluster/network/loadbalancerbackendaddresspooladdress"
	loadbalancernatpool "github.com/upbound/provider-azure/internal/controller/cluster/network/loadbalancernatpool"
	loadbalancernatrule "github.com/upbound/provider-azure/internal/controller/cluster/network/loadbalancernatrule"
	loadbalanceroutboundrule "github.com/upbound/provider-azure/internal/controller/cluster/network/loadbalanceroutboundrule"
	loadbalancerprobe "github.com/upbound/provider-azure/internal/controller/cluster/network/loadbalancerprobe"
	loadbalancerrule "github.com/upbound/provider-azure/internal/controller/cluster/network/loadbalancerrule"
	localnetworkgateway "github.com/upbound/provider-azure/internal/controller/cluster/network/localnetworkgateway"
	manager "github.com/upbound/provider-azure/internal/controller/cluster/network/manager"
	managermanagementgroupconnection "github.com/upbound/provider-azure/internal/controller/cluster/network/managermanagementgroupconnection"
	managernetworkgroup "github.com/upbound/provider-azure/internal/controller/cluster/network/managernetworkgroup"
	managerstaticmember "github.com/upbound/provider-azure/internal/controller/cluster/network/managerstaticmember"
	managersubscriptionconnection "github.com/upbound/provider-azure/internal/controller/cluster/network/managersubscriptionconnection"
	natgateway "github.com/upbound/provider-azure/internal/controller/cluster/network/natgateway"
	natgatewaypublicipassociation "github.com/upbound/provider-azure/internal/controller/cluster/network/natgatewaypublicipassociation"
	natgatewaypublicipprefixassociation "github.com/upbound/provider-azure/internal/controller/cluster/network/natgatewaypublicipprefixassociation"
	networkinterface "github.com/upbound/provider-azure/internal/controller/cluster/network/networkinterface"
	networkinterfaceapplicationsecuritygroupassociation "github.com/upbound/provider-azure/internal/controller/cluster/network/networkinterfaceapplicationsecuritygroupassociation"
	networkinterfacebackendaddresspoolassociation "github.com/upbound/provider-azure/internal/controller/cluster/network/networkinterfacebackendaddresspoolassociation"
	networkinterfacenatruleassociation "github.com/upbound/provider-azure/internal/controller/cluster/network/networkinterfacenatruleassociation"
	networkinterfacesecuritygroupassociation "github.com/upbound/provider-azure/internal/controller/cluster/network/networkinterfacesecuritygroupassociation"
	packetcapture "github.com/upbound/provider-azure/internal/controller/cluster/network/packetcapture"
	pointtositevpngateway "github.com/upbound/provider-azure/internal/controller/cluster/network/pointtositevpngateway"
	privatednsaaaarecord "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednsaaaarecord"
	privatednsarecord "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednsarecord"
	privatednscnamerecord "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednscnamerecord"
	privatednsmxrecord "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednsmxrecord"
	privatednsptrrecord "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednsptrrecord"
	privatednsresolver "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednsresolver"
	privatednsresolverdnsforwardingruleset "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednsresolverdnsforwardingruleset"
	privatednsresolverforwardingrule "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednsresolverforwardingrule"
	privatednsresolverinboundendpoint "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednsresolverinboundendpoint"
	privatednsresolveroutboundendpoint "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednsresolveroutboundendpoint"
	privatednssrvrecord "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednssrvrecord"
	privatednstxtrecord "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednstxtrecord"
	privatednszone "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednszone"
	privatednszonevirtualnetworklink "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednszonevirtualnetworklink"
	privateendpoint "github.com/upbound/provider-azure/internal/controller/cluster/network/privateendpoint"
	privateendpointapplicationsecuritygroupassociation "github.com/upbound/provider-azure/internal/controller/cluster/network/privateendpointapplicationsecuritygroupassociation"
	privatelinkservice "github.com/upbound/provider-azure/internal/controller/cluster/network/privatelinkservice"
	profilenetwork "github.com/upbound/provider-azure/internal/controller/cluster/network/profile"
	publicip "github.com/upbound/provider-azure/internal/controller/cluster/network/publicip"
	publicipprefix "github.com/upbound/provider-azure/internal/controller/cluster/network/publicipprefix"
	route "github.com/upbound/provider-azure/internal/controller/cluster/network/route"
	routefilter "github.com/upbound/provider-azure/internal/controller/cluster/network/routefilter"
	routemap "github.com/upbound/provider-azure/internal/controller/cluster/network/routemap"
	routeserver "github.com/upbound/provider-azure/internal/controller/cluster/network/routeserver"
	routeserverbgpconnection "github.com/upbound/provider-azure/internal/controller/cluster/network/routeserverbgpconnection"
	routetable "github.com/upbound/provider-azure/internal/controller/cluster/network/routetable"
	securitygroup "github.com/upbound/provider-azure/internal/controller/cluster/network/securitygroup"
	securityrule "github.com/upbound/provider-azure/internal/controller/cluster/network/securityrule"
	subnet "github.com/upbound/provider-azure/internal/controller/cluster/network/subnet"
	subnetnatgatewayassociation "github.com/upbound/provider-azure/internal/controller/cluster/network/subnetnatgatewayassociation"
	subnetnetworksecuritygroupassociation "github.com/upbound/provider-azure/internal/controller/cluster/network/subnetnetworksecuritygroupassociation"
	subnetroutetableassociation "github.com/upbound/provider-azure/internal/controller/cluster/network/subnetroutetableassociation"
	subnetserviceendpointstoragepolicy "github.com/upbound/provider-azure/internal/controller/cluster/network/subnetserviceendpointstoragepolicy"
	trafficmanagerazureendpoint "github.com/upbound/provider-azure/internal/controller/cluster/network/trafficmanagerazureendpoint"
	trafficmanagerexternalendpoint "github.com/upbound/provider-azure/internal/controller/cluster/network/trafficmanagerexternalendpoint"
	trafficmanagernestedendpoint "github.com/upbound/provider-azure/internal/controller/cluster/network/trafficmanagernestedendpoint"
	trafficmanagerprofile "github.com/upbound/provider-azure/internal/controller/cluster/network/trafficmanagerprofile"
	virtualhub "github.com/upbound/provider-azure/internal/controller/cluster/network/virtualhub"
	virtualhubconnection "github.com/upbound/provider-azure/internal/controller/cluster/network/virtualhubconnection"
	virtualhubip "github.com/upbound/provider-azure/internal/controller/cluster/network/virtualhubip"
	virtualhubroutetable "github.com/upbound/provider-azure/internal/controller/cluster/network/virtualhubroutetable"
	virtualhubroutetableroute "github.com/upbound/provider-azure/internal/controller/cluster/network/virtualhubroutetableroute"
	virtualhubsecuritypartnerprovider "github.com/upbound/provider-azure/internal/controller/cluster/network/virtualhubsecuritypartnerprovider"
	virtualnetworknetwork "github.com/upbound/provider-azure/internal/controller/cluster/network/virtualnetwork"
	virtualnetworkgateway "github.com/upbound/provider-azure/internal/controller/cluster/network/virtualnetworkgateway"
	virtualnetworkgatewayconnection "github.com/upbound/provider-azure/internal/controller/cluster/network/virtualnetworkgatewayconnection"
	virtualnetworkpeering "github.com/upbound/provider-azure/internal/controller/cluster/network/virtualnetworkpeering"
	virtualwan "github.com/upbound/provider-azure/internal/controller/cluster/network/virtualwan"
	vpngateway "github.com/upbound/provider-azure/internal/controller/cluster/network/vpngateway"
	vpngatewayconnection "github.com/upbound/provider-azure/internal/controller/cluster/network/vpngatewayconnection"
	vpnserverconfiguration "github.com/upbound/provider-azure/internal/controller/cluster/network/vpnserverconfiguration"
	vpnserverconfigurationpolicygroup "github.com/upbound/provider-azure/internal/controller/cluster/network/vpnserverconfigurationpolicygroup"
	vpnsite "github.com/upbound/provider-azure/internal/controller/cluster/network/vpnsite"
	watcher "github.com/upbound/provider-azure/internal/controller/cluster/network/watcher"
	watcherflowlog "github.com/upbound/provider-azure/internal/controller/cluster/network/watcherflowlog"
	webapplicationfirewallpolicy "github.com/upbound/provider-azure/internal/controller/cluster/network/webapplicationfirewallpolicy"
	authorizationrulenotificationhubs "github.com/upbound/provider-azure/internal/controller/cluster/notificationhubs/authorizationrule"
	notificationhub "github.com/upbound/provider-azure/internal/controller/cluster/notificationhubs/notificationhub"
	notificationhubnamespace "github.com/upbound/provider-azure/internal/controller/cluster/notificationhubs/notificationhubnamespace"
	loganalyticsdataexportrule "github.com/upbound/provider-azure/internal/controller/cluster/operationalinsights/loganalyticsdataexportrule"
	loganalyticsdatasourcewindowsevent "github.com/upbound/provider-azure/internal/controller/cluster/operationalinsights/loganalyticsdatasourcewindowsevent"
	loganalyticsdatasourcewindowsperformancecounter "github.com/upbound/provider-azure/internal/controller/cluster/operationalinsights/loganalyticsdatasourcewindowsperformancecounter"
	loganalyticslinkedservice "github.com/upbound/provider-azure/internal/controller/cluster/operationalinsights/loganalyticslinkedservice"
	loganalyticslinkedstorageaccount "github.com/upbound/provider-azure/internal/controller/cluster/operationalinsights/loganalyticslinkedstorageaccount"
	loganalyticsquerypack "github.com/upbound/provider-azure/internal/controller/cluster/operationalinsights/loganalyticsquerypack"
	loganalyticsquerypackquery "github.com/upbound/provider-azure/internal/controller/cluster/operationalinsights/loganalyticsquerypackquery"
	loganalyticssavedsearch "github.com/upbound/provider-azure/internal/controller/cluster/operationalinsights/loganalyticssavedsearch"
	loganalyticssolution "github.com/upbound/provider-azure/internal/controller/cluster/operationsmanagement/loganalyticssolution"
	contactprofile "github.com/upbound/provider-azure/internal/controller/cluster/orbital/contactprofile"
	spacecraft "github.com/upbound/provider-azure/internal/controller/cluster/orbital/spacecraft"
	resourcepolicyremediation "github.com/upbound/provider-azure/internal/controller/cluster/policyinsights/resourcepolicyremediation"
	subscriptionpolicyremediation "github.com/upbound/provider-azure/internal/controller/cluster/policyinsights/subscriptionpolicyremediation"
	dashboard "github.com/upbound/provider-azure/internal/controller/cluster/portal/dashboard"
	powerbiembedded "github.com/upbound/provider-azure/internal/controller/cluster/powerbidedicated/powerbiembedded"
	providerconfig "github.com/upbound/provider-azure/internal/controller/cluster/providerconfig"
	backupcontainerstorageaccount "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/backupcontainerstorageaccount"
	backuppolicyfileshare "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/backuppolicyfileshare"
	backuppolicyvm "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/backuppolicyvm"
	backuppolicyvmworkload "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/backuppolicyvmworkload"
	backupprotectedfileshare "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/backupprotectedfileshare"
	backupprotectedvm "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/backupprotectedvm"
	siterecoveryfabric "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/siterecoveryfabric"
	siterecoverynetworkmapping "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/siterecoverynetworkmapping"
	siterecoveryprotectioncontainer "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/siterecoveryprotectioncontainer"
	siterecoveryprotectioncontainermapping "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/siterecoveryprotectioncontainermapping"
	siterecoveryreplicationpolicy "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/siterecoveryreplicationpolicy"
	vaultrecoveryservices "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/vault"
	eventrelaynamespace "github.com/upbound/provider-azure/internal/controller/cluster/relay/eventrelaynamespace"
	hybridconnection "github.com/upbound/provider-azure/internal/controller/cluster/relay/hybridconnection"
	hybridconnectionauthorizationrule "github.com/upbound/provider-azure/internal/controller/cluster/relay/hybridconnectionauthorizationrule"
	resourcedeploymentscriptazurecli "github.com/upbound/provider-azure/internal/controller/cluster/resources/resourcedeploymentscriptazurecli"
	resourcedeploymentscriptazurepowershell "github.com/upbound/provider-azure/internal/controller/cluster/resources/resourcedeploymentscriptazurepowershell"
	resourcegrouptemplatedeployment "github.com/upbound/provider-azure/internal/controller/cluster/resources/resourcegrouptemplatedeployment"
	subscriptiontemplatedeployment "github.com/upbound/provider-azure/internal/controller/cluster/resources/subscriptiontemplatedeployment"
	sharedprivatelinkservice "github.com/upbound/provider-azure/internal/controller/cluster/search/sharedprivatelinkservice"
	advancedthreatprotection "github.com/upbound/provider-azure/internal/controller/cluster/security/advancedthreatprotection"
	iotsecuritydevicegroup "github.com/upbound/provider-azure/internal/controller/cluster/security/iotsecuritydevicegroup"
	iotsecuritysolution "github.com/upbound/provider-azure/internal/controller/cluster/security/iotsecuritysolution"
	securitycenterassessment "github.com/upbound/provider-azure/internal/controller/cluster/security/securitycenterassessment"
	securitycenterassessmentpolicy "github.com/upbound/provider-azure/internal/controller/cluster/security/securitycenterassessmentpolicy"
	securitycenterautoprovisioning "github.com/upbound/provider-azure/internal/controller/cluster/security/securitycenterautoprovisioning"
	securitycentercontact "github.com/upbound/provider-azure/internal/controller/cluster/security/securitycentercontact"
	securitycenterservervulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/cluster/security/securitycenterservervulnerabilityassessment"
	securitycenterservervulnerabilityassessmentvirtualmachine "github.com/upbound/provider-azure/internal/controller/cluster/security/securitycenterservervulnerabilityassessmentvirtualmachine"
	securitycentersetting "github.com/upbound/provider-azure/internal/controller/cluster/security/securitycentersetting"
	securitycentersubscriptionpricing "github.com/upbound/provider-azure/internal/controller/cluster/security/securitycentersubscriptionpricing"
	securitycenterworkspace "github.com/upbound/provider-azure/internal/controller/cluster/security/securitycenterworkspace"
	storagedefender "github.com/upbound/provider-azure/internal/controller/cluster/security/storagedefender"
	sentinelalertrulefusion "github.com/upbound/provider-azure/internal/controller/cluster/securityinsights/sentinelalertrulefusion"
	sentinelalertrulemachinelearningbehavioranalytics "github.com/upbound/provider-azure/internal/controller/cluster/securityinsights/sentinelalertrulemachinelearningbehavioranalytics"
	sentinelalertrulemssecurityincident "github.com/upbound/provider-azure/internal/controller/cluster/securityinsights/sentinelalertrulemssecurityincident"
	sentinelautomationrule "github.com/upbound/provider-azure/internal/controller/cluster/securityinsights/sentinelautomationrule"
	sentineldataconnectoriot "github.com/upbound/provider-azure/internal/controller/cluster/securityinsights/sentineldataconnectoriot"
	sentinelloganalyticsworkspaceonboarding "github.com/upbound/provider-azure/internal/controller/cluster/securityinsights/sentinelloganalyticsworkspaceonboarding"
	sentinelwatchlist "github.com/upbound/provider-azure/internal/controller/cluster/securityinsights/sentinelwatchlist"
	namespaceauthorizationruleservicebus "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/namespaceauthorizationrule"
	namespacedisasterrecoveryconfigservicebus "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/namespacedisasterrecoveryconfig"
	namespacenetworkruleset "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/namespacenetworkruleset"
	queueauthorizationrule "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/queueauthorizationrule"
	servicebusnamespace "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/servicebusnamespace"
	subscriptionservicebus "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/subscription"
	subscriptionrule "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/subscriptionrule"
	topicservicebus "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/topic"
	topicauthorizationrule "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/topicauthorizationrule"
	managedcluster "github.com/upbound/provider-azure/internal/controller/cluster/servicefabric/managedcluster"
	springcloudconnection "github.com/upbound/provider-azure/internal/controller/cluster/servicelinker/springcloudconnection"
	networkacl "github.com/upbound/provider-azure/internal/controller/cluster/signalrservice/networkacl"
	servicesignalrservice "github.com/upbound/provider-azure/internal/controller/cluster/signalrservice/service"
	signalrsharedprivatelinkresource "github.com/upbound/provider-azure/internal/controller/cluster/signalrservice/signalrsharedprivatelinkresource"
	webpubsub "github.com/upbound/provider-azure/internal/controller/cluster/signalrservice/webpubsub"
	webpubsubhub "github.com/upbound/provider-azure/internal/controller/cluster/signalrservice/webpubsubhub"
	webpubsubnetworkacl "github.com/upbound/provider-azure/internal/controller/cluster/signalrservice/webpubsubnetworkacl"
	managedapplicationdefinition "github.com/upbound/provider-azure/internal/controller/cluster/solutions/managedapplicationdefinition"
	cloudapplicationliveview "github.com/upbound/provider-azure/internal/controller/cluster/spring/cloudapplicationliveview"
	mssqldatabase "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqldatabase"
	mssqldatabaseextendedauditingpolicy "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqldatabaseextendedauditingpolicy"
	mssqldatabasevulnerabilityassessmentrulebaseline "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqldatabasevulnerabilityassessmentrulebaseline"
	mssqlelasticpool "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlelasticpool"
	mssqlfailovergroup "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlfailovergroup"
	mssqlfirewallrule "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlfirewallrule"
	mssqljobagent "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqljobagent"
	mssqljobcredential "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqljobcredential"
	mssqlmanageddatabase "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlmanageddatabase"
	mssqlmanagedinstance "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlmanagedinstance"
	mssqlmanagedinstanceactivedirectoryadministrator "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlmanagedinstanceactivedirectoryadministrator"
	mssqlmanagedinstancefailovergroup "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlmanagedinstancefailovergroup"
	mssqlmanagedinstancetransparentdataencryption "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlmanagedinstancetransparentdataencryption"
	mssqlmanagedinstancevulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlmanagedinstancevulnerabilityassessment"
	mssqloutboundfirewallrule "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqloutboundfirewallrule"
	mssqlserver "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlserver"
	mssqlserverdnsalias "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlserverdnsalias"
	mssqlservermicrosoftsupportauditingpolicy "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlservermicrosoftsupportauditingpolicy"
	mssqlserversecurityalertpolicy "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlserversecurityalertpolicy"
	mssqlservertransparentdataencryption "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlservertransparentdataencryption"
	mssqlservervulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlservervulnerabilityassessment"
	mssqlvirtualnetworkrule "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlvirtualnetworkrule"
	accountstorage "github.com/upbound/provider-azure/internal/controller/cluster/storage/account"
	accountlocaluser "github.com/upbound/provider-azure/internal/controller/cluster/storage/accountlocaluser"
	accountnetworkrules "github.com/upbound/provider-azure/internal/controller/cluster/storage/accountnetworkrules"
	blob "github.com/upbound/provider-azure/internal/controller/cluster/storage/blob"
	blobinventorypolicy "github.com/upbound/provider-azure/internal/controller/cluster/storage/blobinventorypolicy"
	container "github.com/upbound/provider-azure/internal/controller/cluster/storage/container"
	containerimmutabilitypolicy "github.com/upbound/provider-azure/internal/controller/cluster/storage/containerimmutabilitypolicy"
	datalakegen2filesystem "github.com/upbound/provider-azure/internal/controller/cluster/storage/datalakegen2filesystem"
	datalakegen2path "github.com/upbound/provider-azure/internal/controller/cluster/storage/datalakegen2path"
	encryptionscope "github.com/upbound/provider-azure/internal/controller/cluster/storage/encryptionscope"
	managementpolicy "github.com/upbound/provider-azure/internal/controller/cluster/storage/managementpolicy"
	objectreplication "github.com/upbound/provider-azure/internal/controller/cluster/storage/objectreplication"
	queuestorage "github.com/upbound/provider-azure/internal/controller/cluster/storage/queue"
	share "github.com/upbound/provider-azure/internal/controller/cluster/storage/share"
	sharedirectory "github.com/upbound/provider-azure/internal/controller/cluster/storage/sharedirectory"
	tablestorage "github.com/upbound/provider-azure/internal/controller/cluster/storage/table"
	tableentity "github.com/upbound/provider-azure/internal/controller/cluster/storage/tableentity"
	hpccache "github.com/upbound/provider-azure/internal/controller/cluster/storagecache/hpccache"
	hpccacheaccesspolicy "github.com/upbound/provider-azure/internal/controller/cluster/storagecache/hpccacheaccesspolicy"
	hpccacheblobnfstarget "github.com/upbound/provider-azure/internal/controller/cluster/storagecache/hpccacheblobnfstarget"
	hpccacheblobtarget "github.com/upbound/provider-azure/internal/controller/cluster/storagecache/hpccacheblobtarget"
	hpccachenfstarget "github.com/upbound/provider-azure/internal/controller/cluster/storagecache/hpccachenfstarget"
	diskpool "github.com/upbound/provider-azure/internal/controller/cluster/storagepool/diskpool"
	storagesync "github.com/upbound/provider-azure/internal/controller/cluster/storagesync/storagesync"
	clusterstreamanalytics "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/cluster"
	functionjavascriptuda "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/functionjavascriptuda"
	jobstreamanalytics "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/job"
	outputblob "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/outputblob"
	outputeventhub "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/outputeventhub"
	outputfunction "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/outputfunction"
	outputmssql "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/outputmssql"
	outputpowerbi "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/outputpowerbi"
	outputservicebusqueue "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/outputservicebusqueue"
	outputservicebustopic "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/outputservicebustopic"
	outputsynapse "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/outputsynapse"
	outputtable "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/outputtable"
	referenceinputblob "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/referenceinputblob"
	referenceinputmssql "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/referenceinputmssql"
	streaminputblob "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/streaminputblob"
	streaminputeventhub "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/streaminputeventhub"
	streaminputiothub "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/streaminputiothub"
	firewallrulesynapse "github.com/upbound/provider-azure/internal/controller/cluster/synapse/firewallrule"
	integrationruntimeazuresynapse "github.com/upbound/provider-azure/internal/controller/cluster/synapse/integrationruntimeazure"
	integrationruntimeselfhostedsynapse "github.com/upbound/provider-azure/internal/controller/cluster/synapse/integrationruntimeselfhosted"
	linkedservice "github.com/upbound/provider-azure/internal/controller/cluster/synapse/linkedservice"
	managedprivateendpointsynapse "github.com/upbound/provider-azure/internal/controller/cluster/synapse/managedprivateendpoint"
	privatelinkhub "github.com/upbound/provider-azure/internal/controller/cluster/synapse/privatelinkhub"
	roleassignmentsynapse "github.com/upbound/provider-azure/internal/controller/cluster/synapse/roleassignment"
	sparkpool "github.com/upbound/provider-azure/internal/controller/cluster/synapse/sparkpool"
	sqlpool "github.com/upbound/provider-azure/internal/controller/cluster/synapse/sqlpool"
	sqlpoolextendedauditingpolicy "github.com/upbound/provider-azure/internal/controller/cluster/synapse/sqlpoolextendedauditingpolicy"
	sqlpoolsecurityalertpolicy "github.com/upbound/provider-azure/internal/controller/cluster/synapse/sqlpoolsecurityalertpolicy"
	sqlpoolworkloadclassifier "github.com/upbound/provider-azure/internal/controller/cluster/synapse/sqlpoolworkloadclassifier"
	sqlpoolworkloadgroup "github.com/upbound/provider-azure/internal/controller/cluster/synapse/sqlpoolworkloadgroup"
	workspacesynapse "github.com/upbound/provider-azure/internal/controller/cluster/synapse/workspace"
	workspaceaadadmin "github.com/upbound/provider-azure/internal/controller/cluster/synapse/workspaceaadadmin"
	workspaceextendedauditingpolicy "github.com/upbound/provider-azure/internal/controller/cluster/synapse/workspaceextendedauditingpolicy"
	workspacesecurityalertpolicy "github.com/upbound/provider-azure/internal/controller/cluster/synapse/workspacesecurityalertpolicy"
	workspacesqlaadadmin "github.com/upbound/provider-azure/internal/controller/cluster/synapse/workspacesqlaadadmin"
	workspacevulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/cluster/synapse/workspacevulnerabilityassessment"
	eventsourceeventhub "github.com/upbound/provider-azure/internal/controller/cluster/timeseriesinsights/eventsourceeventhub"
	eventsourceiothub "github.com/upbound/provider-azure/internal/controller/cluster/timeseriesinsights/eventsourceiothub"
	gen2environment "github.com/upbound/provider-azure/internal/controller/cluster/timeseriesinsights/gen2environment"
	referencedataset "github.com/upbound/provider-azure/internal/controller/cluster/timeseriesinsights/referencedataset"
	standardenvironment "github.com/upbound/provider-azure/internal/controller/cluster/timeseriesinsights/standardenvironment"
	appactiveslot "github.com/upbound/provider-azure/internal/controller/cluster/web/appactiveslot"
	apphybridconnection "github.com/upbound/provider-azure/internal/controller/cluster/web/apphybridconnection"
	appserviceplan "github.com/upbound/provider-azure/internal/controller/cluster/web/appserviceplan"
	functionapp "github.com/upbound/provider-azure/internal/controller/cluster/web/functionapp"
	functionappactiveslot "github.com/upbound/provider-azure/internal/controller/cluster/web/functionappactiveslot"
	functionappfunction "github.com/upbound/provider-azure/internal/controller/cluster/web/functionappfunction"
	functionapphybridconnection "github.com/upbound/provider-azure/internal/controller/cluster/web/functionapphybridconnection"
	functionappslot "github.com/upbound/provider-azure/internal/controller/cluster/web/functionappslot"
	linuxfunctionapp "github.com/upbound/provider-azure/internal/controller/cluster/web/linuxfunctionapp"
	linuxfunctionappslot "github.com/upbound/provider-azure/internal/controller/cluster/web/linuxfunctionappslot"
	linuxwebapp "github.com/upbound/provider-azure/internal/controller/cluster/web/linuxwebapp"
	linuxwebappslot "github.com/upbound/provider-azure/internal/controller/cluster/web/linuxwebappslot"
	serviceplan "github.com/upbound/provider-azure/internal/controller/cluster/web/serviceplan"
	sourcecontroltoken "github.com/upbound/provider-azure/internal/controller/cluster/web/sourcecontroltoken"
	staticsite "github.com/upbound/provider-azure/internal/controller/cluster/web/staticsite"
	windowsfunctionapp "github.com/upbound/provider-azure/internal/controller/cluster/web/windowsfunctionapp"
	windowsfunctionappslot "github.com/upbound/provider-azure/internal/controller/cluster/web/windowsfunctionappslot"
	windowswebapp "github.com/upbound/provider-azure/internal/controller/cluster/web/windowswebapp"
	windowswebappslot "github.com/upbound/provider-azure/internal/controller/cluster/web/windowswebappslot"
)

var monolithCrdGroup = "azure.upbound.io"

// Setup_monolith creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_monolith(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: monolithCrdGroup, Kind: "API"}:                                                       api.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "APIDiagnostic"}:                                             apidiagnostic.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "APIOperation"}:                                              apioperation.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "APIOperationPolicy"}:                                        apioperationpolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "APIOperationTag"}:                                           apioperationtag.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "APIPolicy"}:                                                 apipolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "APIRelease"}:                                                apirelease.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "APISchema"}:                                                 apischema.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "APITag"}:                                                    apitag.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "APIVersionSet"}:                                             apiversionset.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AccessConnector"}:                                           accessconnector.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AccessPolicy"}:                                              accesspolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Account"}:                                                   accountstorage.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AccountLocalUser"}:                                          accountlocaluser.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AccountNetworkRules"}:                                       accountnetworkrules.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ActiveDirectoryAdministrator"}:                              activedirectoryadministratordbforpostgresql.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AdvancedThreatProtection"}:                                  advancedthreatprotection.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AgentPool"}:                                                 agentpool.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AppActionCustom"}:                                           appactioncustom.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AppActionHTTP"}:                                             appactionhttp.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AppActiveSlot"}:                                             appactiveslot.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AppHybridConnection"}:                                       apphybridconnection.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AppIntegrationAccount"}:                                     appintegrationaccount.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AppIntegrationAccountBatchConfiguration"}:                   appintegrationaccountbatchconfiguration.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AppIntegrationAccountPartner"}:                              appintegrationaccountpartner.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AppIntegrationAccountSchema"}:                               appintegrationaccountschema.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AppIntegrationAccountSession"}:                              appintegrationaccountsession.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AppServiceCertificateOrder"}:                                appservicecertificateorder.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AppServicePlan"}:                                            appserviceplan.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AppTriggerCustom"}:                                          apptriggercustom.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AppTriggerHTTPRequest"}:                                     apptriggerhttprequest.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AppTriggerRecurrence"}:                                      apptriggerrecurrence.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AppWorkflow"}:                                               appworkflow.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Application"}:                                               application.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ApplicationGateway"}:                                        applicationgateway.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ApplicationInsights"}:                                       applicationinsights.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ApplicationInsightsAPIKey"}:                                 applicationinsightsapikey.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ApplicationInsightsAnalyticsItem"}:                          applicationinsightsanalyticsitem.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ApplicationInsightsSmartDetectionRule"}:                     applicationinsightssmartdetectionrule.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ApplicationInsightsStandardWebTest"}:                        applicationinsightsstandardwebtest.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ApplicationInsightsWebTest"}:                                applicationinsightswebtest.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ApplicationInsightsWorkbook"}:                               applicationinsightsworkbook.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ApplicationInsightsWorkbookTemplate"}:                       applicationinsightsworkbooktemplate.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ApplicationNetworkRuleSet"}:                                 applicationnetworkruleset.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ApplicationSecurityGroup"}:                                  applicationsecuritygroup.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Asset"}:                                                     asset.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AssetFilter"}:                                               assetfilter.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AttachedDatabaseConfiguration"}:                             attacheddatabaseconfiguration.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AuthorizationRule"}:                                         authorizationrulenotificationhubs.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AuthorizationServer"}:                                       authorizationserver.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "AvailabilitySet"}:                                           availabilityset.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Backend"}:                                                   backend.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BackupContainerStorageAccount"}:                             backupcontainerstorageaccount.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BackupInstanceBlobStorage"}:                                 backupinstanceblobstorage.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BackupInstanceDisk"}:                                        backupinstancedisk.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BackupInstanceKubernetesCluster"}:                           backupinstancekubernetescluster.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BackupInstancePostgreSQL"}:                                  backupinstancepostgresql.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BackupPolicyBlobStorage"}:                                   backuppolicyblobstorage.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BackupPolicyDisk"}:                                          backuppolicydisk.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BackupPolicyFileShare"}:                                     backuppolicyfileshare.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BackupPolicyKubernetesCluster"}:                             backuppolicykubernetescluster.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BackupPolicyPostgreSQL"}:                                    backuppolicypostgresql.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BackupPolicyVM"}:                                            backuppolicyvm.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BackupPolicyVMWorkload"}:                                    backuppolicyvmworkload.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BackupProtectedFileShare"}:                                  backupprotectedfileshare.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BackupProtectedVM"}:                                         backupprotectedvm.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BackupVault"}:                                               backupvault.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BastionHost"}:                                               bastionhost.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Blob"}:                                                      blob.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BlobInventoryPolicy"}:                                       blobinventorypolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BotChannelAlexa"}:                                           botchannelalexa.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BotChannelDirectLine"}:                                      botchanneldirectline.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BotChannelLine"}:                                            botchannelline.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BotChannelMSTeams"}:                                         botchannelmsteams.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BotChannelSMS"}:                                             botchannelsms.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BotChannelSlack"}:                                           botchannelslack.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BotChannelWebChat"}:                                         botchannelwebchat.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BotChannelsRegistration"}:                                   botchannelsregistration.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BotConnection"}:                                             botconnection.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BotWebApp"}:                                                 botwebapp.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BudgetManagementGroup"}:                                     budgetmanagementgroup.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BudgetResourceGroup"}:                                       budgetresourcegroup.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "BudgetSubscription"}:                                        budgetsubscription.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "CapacityReservation"}:                                       capacityreservation.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "CapacityReservationGroup"}:                                  capacityreservationgroup.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "CassandraCluster"}:                                          cassandracluster.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "CassandraDatacenter"}:                                       cassandradatacenter.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "CassandraKeySpace"}:                                         cassandrakeyspace.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "CassandraTable"}:                                            cassandratable.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Certificate"}:                                               certificatekeyvault.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "CertificateContacts"}:                                       certificatecontacts.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "CertificateIssuer"}:                                         certificateissuer.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "CloudApplicationLiveView"}:                                  cloudapplicationliveview.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "CloudElasticsearch"}:                                        cloudelasticsearch.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Cluster"}:                                                   clusterstreamanalytics.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ClusterManagedPrivateEndpoint"}:                             clustermanagedprivateendpoint.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ClusterPrincipalAssignment"}:                                clusterprincipalassignment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ComputeCluster"}:                                            computecluster.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ComputeInstance"}:                                           computeinstance.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Configuration"}:                                             configurationdbforpostgresql.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Connection"}:                                                connection.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ConnectionClassicCertificate"}:                              connectionclassiccertificate.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ConnectionMonitor"}:                                         connectionmonitor.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ConnectionType"}:                                            connectiontype.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ConsumerGroup"}:                                             consumergroup.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ContactProfile"}:                                            contactprofile.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Container"}:                                                 container.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ContainerApp"}:                                              containerapp.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ContainerConnectedRegistry"}:                                containerconnectedregistry.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ContainerImmutabilityPolicy"}:                               containerimmutabilitypolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ContentKeyPolicy"}:                                          contentkeypolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "CostAnomalyAlert"}:                                          costanomalyalert.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Creator"}:                                                   creator.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Credential"}:                                                credential.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "CustomDataSet"}:                                             customdataset.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "CustomDomain"}:                                              customdomaincontainerapp.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "CustomProvider"}:                                            customprovider.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DDoSProtectionPlan"}:                                        ddosprotectionplan.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DNSAAAARecord"}:                                             dnsaaaarecord.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DNSARecord"}:                                                dnsarecord.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DNSCAARecord"}:                                              dnscaarecord.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DNSCNAMERecord"}:                                            dnscnamerecord.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DNSMXRecord"}:                                               dnsmxrecord.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DNSNSRecord"}:                                               dnsnsrecord.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DNSPTRRecord"}:                                              dnsptrrecord.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DNSSRVRecord"}:                                              dnssrvrecord.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DNSTXTRecord"}:                                              dnstxtrecord.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DNSZone"}:                                                   dnszone.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Dashboard"}:                                                 dashboard.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DataFlow"}:                                                  dataflow.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DataLakeGen2FileSystem"}:                                    datalakegen2filesystem.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DataLakeGen2Path"}:                                          datalakegen2path.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DataSetAzureBlob"}:                                          datasetazureblob.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DataSetBinary"}:                                             datasetbinary.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DataSetBlobStorage"}:                                        datasetblobstorage.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DataSetCosmosDBSQLAPI"}:                                     datasetcosmosdbsqlapi.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DataSetDataLakeGen2"}:                                       datasetdatalakegen2.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DataSetDelimitedText"}:                                      datasetdelimitedtext.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DataSetHTTP"}:                                               datasethttp.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DataSetJSON"}:                                               datasetjson.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DataSetKustoCluster"}:                                       datasetkustocluster.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DataSetKustoDatabase"}:                                      datasetkustodatabase.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DataSetMySQL"}:                                              datasetmysql.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DataSetParquet"}:                                            datasetparquet.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DataSetPostgreSQL"}:                                         datasetpostgresql.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DataSetSQLServerTable"}:                                     datasetsqlservertable.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DataSetSnowflake"}:                                          datasetsnowflake.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DataShare"}:                                                 datashare.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Database"}:                                                  databasekusto.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DatabaseMigrationProject"}:                                  databasemigrationproject.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DatabaseMigrationService"}:                                  databasemigrationservice.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DatabasePrincipalAssignment"}:                               databaseprincipalassignment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DedicatedHost"}:                                             dedicatedhost.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Deployment"}:                                                deployment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Device"}:                                                    device.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Diagnostic"}:                                                diagnostic.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DiskAccess"}:                                                diskaccess.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DiskEncryptionSet"}:                                         diskencryptionset.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DiskPool"}:                                                  diskpool.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Domain"}:                                                    domain.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "DomainTopic"}:                                               domaintopic.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "EmailTemplate"}:                                             emailtemplate.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "EncryptionScope"}:                                           encryptionscope.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Endpoint"}:                                                  endpoint.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Environment"}:                                               environment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "EnvironmentCertificate"}:                                    environmentcertificate.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "EnvironmentCustomDomain"}:                                   environmentcustomdomain.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "EnvironmentDaprComponent"}:                                  environmentdaprcomponent.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "EnvironmentStorage"}:                                        environmentstorage.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "EventGridDataConnection"}:                                   eventgriddataconnection.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "EventHub"}:                                                  eventhub.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "EventHubDataConnection"}:                                    eventhubdataconnection.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "EventHubNamespace"}:                                         eventhubnamespace.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "EventRelayNamespace"}:                                       eventrelaynamespace.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "EventSourceEventHub"}:                                       eventsourceeventhub.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "EventSourceIOTHub"}:                                         eventsourceiothub.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "EventSubscription"}:                                         eventsubscription.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ExpressRouteCircuit"}:                                       expressroutecircuit.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ExpressRouteCircuitAuthorization"}:                          expressroutecircuitauthorization.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ExpressRouteCircuitConnection"}:                             expressroutecircuitconnection.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ExpressRouteCircuitPeering"}:                                expressroutecircuitpeering.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ExpressRouteConnection"}:                                    expressrouteconnection.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ExpressRouteGateway"}:                                       expressroutegateway.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ExpressRoutePort"}:                                          expressrouteport.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Factory"}:                                                   factory.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FederatedIdentityCredential"}:                               federatedidentitycredential.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Firewall"}:                                                  firewall.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FirewallApplicationRuleCollection"}:                         firewallapplicationrulecollection.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FirewallNATRuleCollection"}:                                 firewallnatrulecollection.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FirewallNetworkRuleCollection"}:                             firewallnetworkrulecollection.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FirewallPolicy"}:                                            firewallpolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FirewallPolicyRuleCollectionGroup"}:                         firewallpolicyrulecollectiongroup.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FirewallRule"}:                                              firewallrulesynapse.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FlexibleDatabase"}:                                          flexibledatabase.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FlexibleServer"}:                                            flexibleserverdbforpostgresql.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FlexibleServerActiveDirectoryAdministrator"}:                flexibleserveractivedirectoryadministrator.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FlexibleServerConfiguration"}:                               flexibleserverconfigurationdbforpostgresql.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FlexibleServerDatabase"}:                                    flexibleserverdatabase.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FlexibleServerFirewallRule"}:                                flexibleserverfirewallruledbforpostgresql.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FrontDoor"}:                                                 frontdoor.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FrontdoorCustomDomain"}:                                     frontdoorcustomdomain.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FrontdoorCustomDomainAssociation"}:                          frontdoorcustomdomainassociation.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FrontdoorCustomHTTPSConfiguration"}:                         frontdoorcustomhttpsconfiguration.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FrontdoorEndpoint"}:                                         frontdoorendpoint.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FrontdoorFirewallPolicy"}:                                   frontdoorfirewallpolicynetwork.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FrontdoorOrigin"}:                                           frontdoororigin.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FrontdoorOriginGroup"}:                                      frontdoororigingroup.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FrontdoorProfile"}:                                          frontdoorprofile.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FrontdoorRoute"}:                                            frontdoorroute.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FrontdoorRule"}:                                             frontdoorrule.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FrontdoorRuleSet"}:                                          frontdoorruleset.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FrontdoorRulesEngine"}:                                      frontdoorrulesengine.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FrontdoorSecurityPolicy"}:                                   frontdoorsecuritypolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FunctionApp"}:                                               functionapp.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FunctionAppActiveSlot"}:                                     functionappactiveslot.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FunctionAppFunction"}:                                       functionappfunction.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FunctionAppHybridConnection"}:                               functionapphybridconnection.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FunctionAppSlot"}:                                           functionappslot.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "FunctionJavascriptUda"}:                                     functionjavascriptuda.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "GalleryApplication"}:                                        galleryapplication.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "GalleryApplicationVersion"}:                                 galleryapplicationversion.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Gateway"}:                                                   gateway.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "GatewayAPI"}:                                                gatewayapi.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Gen2Environment"}:                                           gen2environment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "GlobalSchema"}:                                              globalschema.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "GlobalVMShutdownSchedule"}:                                  globalvmshutdownschedule.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "GremlinDatabase"}:                                           gremlindatabase.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "GremlinGraph"}:                                              gremlingraph.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "HBaseCluster"}:                                              hbasecluster.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "HPCCache"}:                                                  hpccache.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "HPCCacheAccessPolicy"}:                                      hpccacheaccesspolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "HPCCacheBlobNFSTarget"}:                                     hpccacheblobnfstarget.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "HPCCacheBlobTarget"}:                                        hpccacheblobtarget.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "HPCCacheNFSTarget"}:                                         hpccachenfstarget.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "HadoopCluster"}:                                             hadoopcluster.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "HealthBot"}:                                                 healthbot.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "HealthcareDICOMService"}:                                    healthcaredicomservice.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "HealthcareFHIRService"}:                                     healthcarefhirservice.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "HealthcareMedtechService"}:                                  healthcaremedtechservice.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "HealthcareMedtechServiceFHIRDestination"}:                   healthcaremedtechservicefhirdestination.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "HealthcareService"}:                                         healthcareservice.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "HealthcareWorkspace"}:                                       healthcareworkspace.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "HybridConnection"}:                                          hybridconnection.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "HybridConnectionAuthorizationRule"}:                         hybridconnectionauthorizationrule.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "HybridRunBookWorkerGroup"}:                                  hybridrunbookworkergroup.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IOTHub"}:                                                    iothub.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IOTHubCertificate"}:                                         iothubcertificate.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IOTHubConsumerGroup"}:                                       iothubconsumergroup.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IOTHubDPS"}:                                                 iothubdps.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IOTHubDPSCertificate"}:                                      iothubdpscertificate.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IOTHubDPSSharedAccessPolicy"}:                               iothubdpssharedaccesspolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IOTHubDataConnection"}:                                      iothubdataconnection.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IOTHubDeviceUpdateAccount"}:                                 iothubdeviceupdateaccount.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IOTHubDeviceUpdateInstance"}:                                iothubdeviceupdateinstance.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IOTHubEndpointEventHub"}:                                    iothubendpointeventhub.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IOTHubEndpointServiceBusQueue"}:                             iothubendpointservicebusqueue.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IOTHubEndpointServiceBusTopic"}:                             iothubendpointservicebustopic.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IOTHubEndpointStorageContainer"}:                            iothubendpointstoragecontainer.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IOTHubEnrichment"}:                                          iothubenrichment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IOTHubFallbackRoute"}:                                       iothubfallbackroute.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IOTHubRoute"}:                                               iothubroute.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IOTHubSharedAccessPolicy"}:                                  iothubsharedaccesspolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IOTSecurityDeviceGroup"}:                                    iotsecuritydevicegroup.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IOTSecuritySolution"}:                                       iotsecuritysolution.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IPGroup"}:                                                   ipgroup.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IdentityProviderAAD"}:                                       identityprovideraad.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IdentityProviderFacebook"}:                                  identityproviderfacebook.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IdentityProviderGoogle"}:                                    identityprovidergoogle.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IdentityProviderMicrosoft"}:                                 identityprovidermicrosoft.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IdentityProviderTwitter"}:                                   identityprovidertwitter.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Image"}:                                                     image.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Instance"}:                                                  instance.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IntegrationRuntimeAzure"}:                                   integrationruntimeazuresynapse.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IntegrationRuntimeAzureSSIS"}:                               integrationruntimeazuressis.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IntegrationRuntimeManaged"}:                                 integrationruntimemanaged.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IntegrationRuntimeSelfHosted"}:                              integrationruntimeselfhostedsynapse.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "IntegrationServiceEnvironment"}:                             integrationserviceenvironment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "InteractiveQueryCluster"}:                                   interactivequerycluster.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Job"}:                                                       jobstreamanalytics.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "KafkaCluster"}:                                              kafkacluster.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Key"}:                                                       key.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "KubernetesCluster"}:                                         kubernetescluster.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "KubernetesClusterExtension"}:                                kubernetesclusterextension.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "KubernetesClusterNodePool"}:                                 kubernetesclusternodepool.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "KubernetesFleetManager"}:                                    kubernetesfleetmanager.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Lab"}:                                                       lab.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LabServiceLab"}:                                             labservicelab.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LabServicePlan"}:                                            labserviceplan.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Ledger"}:                                                    ledger.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinkedCustomService"}:                                       linkedcustomservice.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinkedService"}:                                             linkedservice.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinkedServiceAzureBlobStorage"}:                             linkedserviceazureblobstorage.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinkedServiceAzureDatabricks"}:                              linkedserviceazuredatabricks.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinkedServiceAzureFileStorage"}:                             linkedserviceazurefilestorage.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinkedServiceAzureFunction"}:                                linkedserviceazurefunction.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinkedServiceAzureSQLDatabase"}:                             linkedserviceazuresqldatabase.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinkedServiceAzureSearch"}:                                  linkedserviceazuresearch.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinkedServiceAzureTableStorage"}:                            linkedserviceazuretablestorage.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinkedServiceCosmosDB"}:                                     linkedservicecosmosdb.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinkedServiceCosmosDBMongoapi"}:                             linkedservicecosmosdbmongoapi.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinkedServiceDataLakeStorageGen2"}:                          linkedservicedatalakestoragegen2.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinkedServiceKeyVault"}:                                     linkedservicekeyvault.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinkedServiceKusto"}:                                        linkedservicekusto.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinkedServiceMySQL"}:                                        linkedservicemysql.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinkedServiceOData"}:                                        linkedserviceodata.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinkedServiceOdbc"}:                                         linkedserviceodbc.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinkedServicePostgreSQL"}:                                   linkedservicepostgresql.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinkedServiceSFTP"}:                                         linkedservicesftp.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinkedServiceSQLServer"}:                                    linkedservicesqlserver.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinkedServiceSnowflake"}:                                    linkedservicesnowflake.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinkedServiceSynapse"}:                                      linkedservicesynapse.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinkedServiceWeb"}:                                          linkedserviceweb.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinuxFunctionApp"}:                                          linuxfunctionapp.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinuxFunctionAppSlot"}:                                      linuxfunctionappslot.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinuxVirtualMachine"}:                                       linuxvirtualmachinedevtestlab.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinuxVirtualMachineScaleSet"}:                               linuxvirtualmachinescaleset.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinuxWebApp"}:                                               linuxwebapp.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LinuxWebAppSlot"}:                                           linuxwebappslot.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LiveEvent"}:                                                 liveevent.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LiveEventOutput"}:                                           liveeventoutput.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LoadBalancer"}:                                              loadbalancer.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LoadBalancerBackendAddressPool"}:                            loadbalancerbackendaddresspool.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LoadBalancerBackendAddressPoolAddress"}:                     loadbalancerbackendaddresspooladdress.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LoadBalancerNatPool"}:                                       loadbalancernatpool.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LoadBalancerNatRule"}:                                       loadbalancernatrule.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LoadBalancerOutboundRule"}:                                  loadbalanceroutboundrule.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LoadBalancerProbe"}:                                         loadbalancerprobe.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LoadBalancerRule"}:                                          loadbalancerrule.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LoadTest"}:                                                  loadtest.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LocalNetworkGateway"}:                                       localnetworkgateway.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LogAnalyticsDataExportRule"}:                                loganalyticsdataexportrule.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LogAnalyticsDataSourceWindowsEvent"}:                        loganalyticsdatasourcewindowsevent.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LogAnalyticsDataSourceWindowsPerformanceCounter"}:           loganalyticsdatasourcewindowsperformancecounter.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LogAnalyticsLinkedService"}:                                 loganalyticslinkedservice.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LogAnalyticsLinkedStorageAccount"}:                          loganalyticslinkedstorageaccount.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LogAnalyticsQueryPack"}:                                     loganalyticsquerypack.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LogAnalyticsQueryPackQuery"}:                                loganalyticsquerypackquery.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LogAnalyticsSavedSearch"}:                                   loganalyticssavedsearch.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "LogAnalyticsSolution"}:                                      loganalyticssolution.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Logger"}:                                                    logger.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MSSQLDatabase"}:                                             mssqldatabase.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MSSQLDatabaseExtendedAuditingPolicy"}:                       mssqldatabaseextendedauditingpolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MSSQLDatabaseVulnerabilityAssessmentRuleBaseline"}:          mssqldatabasevulnerabilityassessmentrulebaseline.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MSSQLElasticPool"}:                                          mssqlelasticpool.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MSSQLFailoverGroup"}:                                        mssqlfailovergroup.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MSSQLFirewallRule"}:                                         mssqlfirewallrule.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MSSQLJobAgent"}:                                             mssqljobagent.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MSSQLJobCredential"}:                                        mssqljobcredential.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MSSQLManagedDatabase"}:                                      mssqlmanageddatabase.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MSSQLManagedInstance"}:                                      mssqlmanagedinstance.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MSSQLManagedInstanceActiveDirectoryAdministrator"}:          mssqlmanagedinstanceactivedirectoryadministrator.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MSSQLManagedInstanceFailoverGroup"}:                         mssqlmanagedinstancefailovergroup.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MSSQLManagedInstanceTransparentDataEncryption"}:             mssqlmanagedinstancetransparentdataencryption.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MSSQLManagedInstanceVulnerabilityAssessment"}:               mssqlmanagedinstancevulnerabilityassessment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MSSQLOutboundFirewallRule"}:                                 mssqloutboundfirewallrule.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MSSQLServer"}:                                               mssqlserver.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MSSQLServerDNSAlias"}:                                       mssqlserverdnsalias.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MSSQLServerMicrosoftSupportAuditingPolicy"}:                 mssqlservermicrosoftsupportauditingpolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MSSQLServerSecurityAlertPolicy"}:                            mssqlserversecurityalertpolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MSSQLServerTransparentDataEncryption"}:                      mssqlservertransparentdataencryption.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MSSQLServerVulnerabilityAssessment"}:                        mssqlservervulnerabilityassessment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MSSQLVirtualNetworkRule"}:                                   mssqlvirtualnetworkrule.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MaintenanceAssignmentDedicatedHost"}:                        maintenanceassignmentdedicatedhost.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MaintenanceAssignmentVirtualMachine"}:                       maintenanceassignmentvirtualmachine.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MaintenanceConfiguration"}:                                  maintenanceconfiguration.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ManagedApplicationDefinition"}:                              managedapplicationdefinition.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ManagedCluster"}:                                            managedcluster.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ManagedDisk"}:                                               manageddisk.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ManagedDiskSASToken"}:                                       manageddisksastoken.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ManagedHardwareSecurityModule"}:                             managedhardwaresecuritymodule.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ManagedPrivateEndpoint"}:                                    managedprivateendpointsynapse.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ManagedStorageAccount"}:                                     managedstorageaccount.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ManagedStorageAccountSASTokenDefinition"}:                   managedstorageaccountsastokendefinition.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Management"}:                                                management.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ManagementGroup"}:                                           managementgroup.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ManagementGroupPolicyAssignment"}:                           managementgrouppolicyassignment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ManagementGroupPolicyExemption"}:                            managementgrouppolicyexemption.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ManagementGroupSubscriptionAssociation"}:                    managementgroupsubscriptionassociation.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ManagementLock"}:                                            managementlock.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ManagementPolicy"}:                                          managementpolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Manager"}:                                                   manager.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ManagerManagementGroupConnection"}:                          managermanagementgroupconnection.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ManagerNetworkGroup"}:                                       managernetworkgroup.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ManagerStaticMember"}:                                       managerstaticmember.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ManagerSubscriptionConnection"}:                             managersubscriptionconnection.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MarketplaceAgreement"}:                                      marketplaceagreement.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Module"}:                                                    module.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MongoCollection"}:                                           mongocollection.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MongoDatabase"}:                                             mongodatabase.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MongoRoleDefinition"}:                                       mongoroledefinition.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MongoUserDefinition"}:                                       mongouserdefinition.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Monitor"}:                                                   monitor.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MonitorActionGroup"}:                                        monitoractiongroup.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MonitorActionRuleActionGroup"}:                              monitoractionruleactiongroup.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MonitorActionRuleSuppression"}:                              monitoractionrulesuppression.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MonitorActivityLogAlert"}:                                   monitoractivitylogalert.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MonitorAlertProcessingRuleActionGroup"}:                     monitoralertprocessingruleactiongroup.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MonitorAlertProcessingRuleSuppression"}:                     monitoralertprocessingrulesuppression.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MonitorAutoscaleSetting"}:                                   monitorautoscalesetting.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MonitorDataCollectionEndpoint"}:                             monitordatacollectionendpoint.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MonitorDataCollectionRule"}:                                 monitordatacollectionrule.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MonitorDataCollectionRuleAssociation"}:                      monitordatacollectionruleassociation.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MonitorDiagnosticSetting"}:                                  monitordiagnosticsetting.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MonitorMetricAlert"}:                                        monitormetricalert.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MonitorPrivateLinkScope"}:                                   monitorprivatelinkscope.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MonitorPrivateLinkScopedService"}:                           monitorprivatelinkscopedservice.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MonitorScheduledQueryRulesAlert"}:                           monitorscheduledqueryrulesalert.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MonitorScheduledQueryRulesAlertV2"}:                         monitorscheduledqueryrulesalertv2.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MonitorScheduledQueryRulesLog"}:                             monitorscheduledqueryruleslog.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "MonitorSmartDetectorAlertRule"}:                             monitorsmartdetectoralertrule.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "NATGateway"}:                                                natgateway.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "NATGatewayPublicIPAssociation"}:                             natgatewaypublicipassociation.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "NATGatewayPublicIPPrefixAssociation"}:                       natgatewaypublicipprefixassociation.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "NamedValue"}:                                                namedvalue.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "NamespaceAuthorizationRule"}:                                namespaceauthorizationruleservicebus.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "NamespaceDisasterRecoveryConfig"}:                           namespacedisasterrecoveryconfigservicebus.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "NamespaceNetworkRuleSet"}:                                   namespacenetworkruleset.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "NamespaceSchemaGroup"}:                                      namespaceschemagroup.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "NetworkACL"}:                                                networkacl.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "NetworkInterface"}:                                          networkinterface.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "NetworkInterfaceApplicationSecurityGroupAssociation"}:       networkinterfaceapplicationsecuritygroupassociation.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "NetworkInterfaceBackendAddressPoolAssociation"}:             networkinterfacebackendaddresspoolassociation.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "NetworkInterfaceNatRuleAssociation"}:                        networkinterfacenatruleassociation.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "NetworkInterfaceSecurityGroupAssociation"}:                  networkinterfacesecuritygroupassociation.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "NotificationHub"}:                                           notificationhub.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "NotificationHubNamespace"}:                                  notificationhubnamespace.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "NotificationRecipientEmail"}:                                notificationrecipientemail.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "NotificationRecipientUser"}:                                 notificationrecipientuser.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ObjectReplication"}:                                         objectreplication.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "OpenIDConnectProvider"}:                                     openidconnectprovider.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "OrchestratedVirtualMachineScaleSet"}:                        orchestratedvirtualmachinescaleset.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "OutputBlob"}:                                                outputblob.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "OutputEventHub"}:                                            outputeventhub.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "OutputFunction"}:                                            outputfunction.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "OutputMSSQL"}:                                               outputmssql.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "OutputPowerBI"}:                                             outputpowerbi.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "OutputServiceBusQueue"}:                                     outputservicebusqueue.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "OutputServiceBusTopic"}:                                     outputservicebustopic.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "OutputSynapse"}:                                             outputsynapse.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "OutputTable"}:                                               outputtable.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PacketCapture"}:                                             packetcapture.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PimActiveRoleAssignment"}:                                   pimactiveroleassignment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PimEligibleRoleAssignment"}:                                 pimeligibleroleassignment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Pipeline"}:                                                  pipeline.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PointToSiteVPNGateway"}:                                     pointtositevpngateway.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Policy"}:                                                    policydevtestlab.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PolicyDefinition"}:                                          policydefinition.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PolicySetDefinition"}:                                       policysetdefinition.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PolicyVirtualMachineConfigurationAssignment"}:               policyvirtualmachineconfigurationassignment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Pool"}:                                                      pool.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PowerBIEmbedded"}:                                           powerbiembedded.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PrivateDNSAAAARecord"}:                                      privatednsaaaarecord.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PrivateDNSARecord"}:                                         privatednsarecord.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PrivateDNSCNAMERecord"}:                                     privatednscnamerecord.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PrivateDNSMXRecord"}:                                        privatednsmxrecord.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PrivateDNSPTRRecord"}:                                       privatednsptrrecord.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PrivateDNSResolver"}:                                        privatednsresolver.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PrivateDNSResolverDNSForwardingRuleset"}:                    privatednsresolverdnsforwardingruleset.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PrivateDNSResolverForwardingRule"}:                          privatednsresolverforwardingrule.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PrivateDNSResolverInboundEndpoint"}:                         privatednsresolverinboundendpoint.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PrivateDNSResolverOutboundEndpoint"}:                        privatednsresolveroutboundendpoint.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PrivateDNSSRVRecord"}:                                       privatednssrvrecord.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PrivateDNSTXTRecord"}:                                       privatednstxtrecord.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PrivateDNSZone"}:                                            privatednszone.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PrivateDNSZoneVirtualNetworkLink"}:                          privatednszonevirtualnetworklink.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PrivateEndpoint"}:                                           privateendpoint.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PrivateEndpointApplicationSecurityGroupAssociation"}:        privateendpointapplicationsecuritygroupassociation.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PrivateLinkHub"}:                                            privatelinkhub.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PrivateLinkService"}:                                        privatelinkservice.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Product"}:                                                   product.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ProductAPI"}:                                                productapi.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ProductPolicy"}:                                             productpolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ProductTag"}:                                                producttag.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Profile"}:                                                   profilenetwork.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Provider"}:                                                  provider.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ProviderConfig"}:                                            providerconfig.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ProximityPlacementGroup"}:                                   proximityplacementgroup.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PublicIP"}:                                                  publicip.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "PublicIPPrefix"}:                                            publicipprefix.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Queue"}:                                                     queuestorage.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "QueueAuthorizationRule"}:                                    queueauthorizationrule.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "RedisCache"}:                                                rediscachecache.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "RedisCacheAccessPolicy"}:                                    rediscacheaccesspolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "RedisCacheAccessPolicyAssignment"}:                          rediscacheaccesspolicyassignment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "RedisEnterpriseCluster"}:                                    redisenterprisecluster.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "RedisEnterpriseDatabase"}:                                   redisenterprisedatabase.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "RedisFirewallRule"}:                                         redisfirewallrule.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "RedisLinkedServer"}:                                         redislinkedserver.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ReferenceDataSet"}:                                          referencedataset.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ReferenceInputBlob"}:                                        referenceinputblob.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ReferenceInputMSSQL"}:                                       referenceinputmssql.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Registry"}:                                                  registry.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ResourceDeploymentScriptAzureCli"}:                          resourcedeploymentscriptazurecli.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ResourceDeploymentScriptAzurePowerShell"}:                   resourcedeploymentscriptazurepowershell.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ResourceGroup"}:                                             resourcegroup.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ResourceGroupCostManagementExport"}:                         resourcegroupcostmanagementexport.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ResourceGroupPolicyAssignment"}:                             resourcegrouppolicyassignment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ResourceGroupPolicyExemption"}:                              resourcegrouppolicyexemption.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ResourceGroupTemplateDeployment"}:                           resourcegrouptemplatedeployment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ResourceGuard"}:                                             resourceguard.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ResourcePolicyAssignment"}:                                  resourcepolicyassignment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ResourcePolicyExemption"}:                                   resourcepolicyexemption.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ResourcePolicyRemediation"}:                                 resourcepolicyremediation.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ResourceProviderRegistration"}:                              resourceproviderregistration.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "RoleAssignment"}:                                            roleassignmentsynapse.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "RoleDefinition"}:                                            roledefinition.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Route"}:                                                     route.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "RouteFilter"}:                                               routefilter.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "RouteMap"}:                                                  routemap.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "RouteServer"}:                                               routeserver.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "RouteServerBGPConnection"}:                                  routeserverbgpconnection.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "RouteTable"}:                                                routetable.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "RunBook"}:                                                   runbook.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SQLContainer"}:                                              sqlcontainer.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SQLDatabase"}:                                               sqldatabase.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SQLDedicatedGateway"}:                                       sqldedicatedgateway.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SQLFunction"}:                                               sqlfunction.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SQLPool"}:                                                   sqlpool.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SQLPoolExtendedAuditingPolicy"}:                             sqlpoolextendedauditingpolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SQLPoolSecurityAlertPolicy"}:                                sqlpoolsecurityalertpolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SQLPoolWorkloadClassifier"}:                                 sqlpoolworkloadclassifier.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SQLPoolWorkloadGroup"}:                                      sqlpoolworkloadgroup.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SQLRoleAssignment"}:                                         sqlroleassignment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SQLRoleDefinition"}:                                         sqlroledefinition.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SQLStoredProcedure"}:                                        sqlstoredprocedure.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SQLTrigger"}:                                                sqltrigger.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SSHPublicKey"}:                                              sshpublickey.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Schedule"}:                                                  scheduledevtestlab.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ScopeMap"}:                                                  scopemap.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Secret"}:                                                    secret.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SecurityCenterAssessment"}:                                  securitycenterassessment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SecurityCenterAssessmentPolicy"}:                            securitycenterassessmentpolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SecurityCenterAutoProvisioning"}:                            securitycenterautoprovisioning.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SecurityCenterContact"}:                                     securitycentercontact.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SecurityCenterServerVulnerabilityAssessment"}:               securitycenterservervulnerabilityassessment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SecurityCenterServerVulnerabilityAssessmentVirtualMachine"}: securitycenterservervulnerabilityassessmentvirtualmachine.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SecurityCenterSetting"}:                                     securitycentersetting.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SecurityCenterSubscriptionPricing"}:                         securitycentersubscriptionpricing.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SecurityCenterWorkspace"}:                                   securitycenterworkspace.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SecurityGroup"}:                                             securitygroup.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SecurityRule"}:                                              securityrule.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SentinelAlertRuleFusion"}:                                   sentinelalertrulefusion.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SentinelAlertRuleMSSecurityIncident"}:                       sentinelalertrulemssecurityincident.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SentinelAlertRuleMachineLearningBehaviorAnalytics"}:         sentinelalertrulemachinelearningbehavioranalytics.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SentinelAutomationRule"}:                                    sentinelautomationrule.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SentinelDataConnectorIOT"}:                                  sentineldataconnectoriot.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SentinelLogAnalyticsWorkspaceOnboarding"}:                   sentinelloganalyticsworkspaceonboarding.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SentinelWatchlist"}:                                         sentinelwatchlist.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Server"}:                                                    serverfluidrelay.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ServerKey"}:                                                 serverkey.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Service"}:                                                   servicesignalrservice.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ServiceBusNamespace"}:                                       servicebusnamespace.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ServicePlan"}:                                               serviceplan.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ServicesAccount"}:                                           servicesaccount.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ServicesAccountFilter"}:                                     servicesaccountfilter.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Share"}:                                                     share.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "ShareDirectory"}:                                            sharedirectory.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SharedImage"}:                                               sharedimage.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SharedImageGallery"}:                                        sharedimagegallery.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SharedPrivateLinkService"}:                                  sharedprivatelinkservice.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SignalrSharedPrivateLinkResource"}:                          signalrsharedprivatelinkresource.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SiteRecoveryFabric"}:                                        siterecoveryfabric.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SiteRecoveryNetworkMapping"}:                                siterecoverynetworkmapping.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SiteRecoveryProtectionContainer"}:                           siterecoveryprotectioncontainer.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SiteRecoveryProtectionContainerMapping"}:                    siterecoveryprotectioncontainermapping.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SiteRecoveryReplicationPolicy"}:                             siterecoveryreplicationpolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Snapshot"}:                                                  snapshotnetapp.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SnapshotPolicy"}:                                            snapshotpolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SourceControlToken"}:                                        sourcecontroltoken.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Spacecraft"}:                                                spacecraft.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SparkCluster"}:                                              sparkcluster.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SparkPool"}:                                                 sparkpool.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SpatialAnchorsAccount"}:                                     spatialanchorsaccount.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SpringCloudAPIPortal"}:                                      springcloudapiportal.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SpringCloudAPIPortalCustomDomain"}:                          springcloudapiportalcustomdomain.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SpringCloudAccelerator"}:                                    springcloudaccelerator.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SpringCloudActiveDeployment"}:                               springcloudactivedeployment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SpringCloudApp"}:                                            springcloudapp.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SpringCloudAppCosmosDBAssociation"}:                         springcloudappcosmosdbassociation.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SpringCloudAppMySQLAssociation"}:                            springcloudappmysqlassociation.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SpringCloudAppRedisAssociation"}:                            springcloudappredisassociation.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SpringCloudBuildDeployment"}:                                springcloudbuilddeployment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SpringCloudBuildPackBinding"}:                               springcloudbuildpackbinding.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SpringCloudBuilder"}:                                        springcloudbuilder.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SpringCloudCertificate"}:                                    springcloudcertificate.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SpringCloudConfigurationService"}:                           springcloudconfigurationservice.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SpringCloudConnection"}:                                     springcloudconnection.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SpringCloudContainerDeployment"}:                            springcloudcontainerdeployment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SpringCloudCustomDomain"}:                                   springcloudcustomdomain.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SpringCloudCustomizedAccelerator"}:                          springcloudcustomizedaccelerator.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SpringCloudDevToolPortal"}:                                  springclouddevtoolportal.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SpringCloudGateway"}:                                        springcloudgateway.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SpringCloudGatewayCustomDomain"}:                            springcloudgatewaycustomdomain.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SpringCloudJavaDeployment"}:                                 springcloudjavadeployment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SpringCloudService"}:                                        springcloudservice.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SpringCloudStorage"}:                                        springcloudstorage.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "StandardEnvironment"}:                                       standardenvironment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "StaticSite"}:                                                staticsite.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "StorageDefender"}:                                           storagedefender.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "StorageSync"}:                                               storagesync.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "StreamInputBlob"}:                                           streaminputblob.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "StreamInputEventHub"}:                                       streaminputeventhub.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "StreamInputIOTHub"}:                                         streaminputiothub.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "StreamingEndpoint"}:                                         streamingendpoint.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "StreamingLocator"}:                                          streaminglocator.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "StreamingPolicy"}:                                           streamingpolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SubAccount"}:                                                subaccount.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SubAccountTagRule"}:                                         subaccounttagrule.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Subnet"}:                                                    subnet.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SubnetNATGatewayAssociation"}:                               subnetnatgatewayassociation.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SubnetNetworkSecurityGroupAssociation"}:                     subnetnetworksecuritygroupassociation.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SubnetRouteTableAssociation"}:                               subnetroutetableassociation.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SubnetServiceEndpointStoragePolicy"}:                        subnetserviceendpointstoragepolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Subscription"}:                                              subscriptionservicebus.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SubscriptionCostManagementExport"}:                          subscriptioncostmanagementexport.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SubscriptionPolicyAssignment"}:                              subscriptionpolicyassignment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SubscriptionPolicyExemption"}:                               subscriptionpolicyexemption.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SubscriptionPolicyRemediation"}:                             subscriptionpolicyremediation.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SubscriptionRule"}:                                          subscriptionrule.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SubscriptionTemplateDeployment"}:                            subscriptiontemplatedeployment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SynapseSpark"}:                                              synapsespark.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "SystemTopic"}:                                               systemtopic.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Table"}:                                                     tablestorage.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "TableEntity"}:                                               tableentity.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Tag"}:                                                       tag.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "TagRule"}:                                                   tagrule.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Token"}:                                                     token.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "TokenPassword"}:                                             tokenpassword.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Topic"}:                                                     topicservicebus.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "TopicAuthorizationRule"}:                                    topicauthorizationrule.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "TrafficManagerAzureEndpoint"}:                               trafficmanagerazureendpoint.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "TrafficManagerExternalEndpoint"}:                            trafficmanagerexternalendpoint.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "TrafficManagerNestedEndpoint"}:                              trafficmanagernestedendpoint.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "TrafficManagerProfile"}:                                     trafficmanagerprofile.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Transform"}:                                                 transform.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "TriggerBlobEvent"}:                                          triggerblobevent.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "TriggerCustomEvent"}:                                        triggercustomevent.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "TriggerSchedule"}:                                           triggerschedule.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "TrustedAccessRoleBinding"}:                                  trustedaccessrolebinding.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "User"}:                                                      user.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "UserAssignedIdentity"}:                                      userassignedidentity.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "VPNGateway"}:                                                vpngateway.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "VPNGatewayConnection"}:                                      vpngatewayconnection.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "VPNServerConfiguration"}:                                    vpnserverconfiguration.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "VPNServerConfigurationPolicyGroup"}:                         vpnserverconfigurationpolicygroup.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "VPNSite"}:                                                   vpnsite.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "VariableBool"}:                                              variablebool.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "VariableDateTime"}:                                          variabledatetime.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "VariableInt"}:                                               variableint.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "VariableString"}:                                            variablestring.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Vault"}:                                                     vaultrecoveryservices.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "VirtualHub"}:                                                virtualhub.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "VirtualHubConnection"}:                                      virtualhubconnection.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "VirtualHubIP"}:                                              virtualhubip.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "VirtualHubRouteTable"}:                                      virtualhubroutetable.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "VirtualHubRouteTableRoute"}:                                 virtualhubroutetableroute.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "VirtualHubSecurityPartnerProvider"}:                         virtualhubsecuritypartnerprovider.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "VirtualMachineDataDiskAttachment"}:                          virtualmachinedatadiskattachment.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "VirtualMachineExtension"}:                                   virtualmachineextension.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "VirtualMachineRunCommand"}:                                  virtualmachineruncommand.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "VirtualNetwork"}:                                            virtualnetworknetwork.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "VirtualNetworkGateway"}:                                     virtualnetworkgateway.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "VirtualNetworkGatewayConnection"}:                           virtualnetworkgatewayconnection.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "VirtualNetworkPeering"}:                                     virtualnetworkpeering.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "VirtualNetworkRule"}:                                        virtualnetworkruledbforpostgresql.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "VirtualWAN"}:                                                virtualwan.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Volume"}:                                                    volume.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Watcher"}:                                                   watcher.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "WatcherFlowLog"}:                                            watcherflowlog.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "WebApplicationFirewallPolicy"}:                              webapplicationfirewallpolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "WebPubsub"}:                                                 webpubsub.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "WebPubsubHub"}:                                              webpubsubhub.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "WebPubsubNetworkACL"}:                                       webpubsubnetworkacl.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Webhook"}:                                                   webhookcontainerregistry.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "WindowsFunctionApp"}:                                        windowsfunctionapp.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "WindowsFunctionAppSlot"}:                                    windowsfunctionappslot.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "WindowsVirtualMachine"}:                                     windowsvirtualmachinedevtestlab.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "WindowsVirtualMachineScaleSet"}:                             windowsvirtualmachinescaleset.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "WindowsWebApp"}:                                             windowswebapp.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "WindowsWebAppSlot"}:                                         windowswebappslot.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "Workspace"}:                                                 workspacesynapse.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "WorkspaceAADAdmin"}:                                         workspaceaadadmin.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "WorkspaceCustomerManagedKey"}:                               workspacecustomermanagedkey.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "WorkspaceExtendedAuditingPolicy"}:                           workspaceextendedauditingpolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "WorkspaceRootDbfsCustomerManagedKey"}:                       workspacerootdbfscustomermanagedkey.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "WorkspaceSQLAADAdmin"}:                                      workspacesqlaadadmin.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "WorkspaceSecurityAlertPolicy"}:                              workspacesecurityalertpolicy.Setup,
		schema.GroupKind{Group: monolithCrdGroup, Kind: "WorkspaceVulnerabilityAssessment"}:                          workspacevulnerabilityassessment.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, monolithCrdGroup, o); err != nil {
		return err
	}
	return nil
}
