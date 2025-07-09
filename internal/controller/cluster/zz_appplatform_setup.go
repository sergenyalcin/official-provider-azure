// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

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
)

var appplatformCrdGroup = "appplatform.azure.upbound.io"

// Setup_appplatform creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_appplatform(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: appplatformCrdGroup, Kind: "SpringCloudAPIPortal"}:              springcloudapiportal.Setup,
		schema.GroupKind{Group: appplatformCrdGroup, Kind: "SpringCloudAPIPortalCustomDomain"}:  springcloudapiportalcustomdomain.Setup,
		schema.GroupKind{Group: appplatformCrdGroup, Kind: "SpringCloudAccelerator"}:            springcloudaccelerator.Setup,
		schema.GroupKind{Group: appplatformCrdGroup, Kind: "SpringCloudActiveDeployment"}:       springcloudactivedeployment.Setup,
		schema.GroupKind{Group: appplatformCrdGroup, Kind: "SpringCloudApp"}:                    springcloudapp.Setup,
		schema.GroupKind{Group: appplatformCrdGroup, Kind: "SpringCloudAppCosmosDBAssociation"}: springcloudappcosmosdbassociation.Setup,
		schema.GroupKind{Group: appplatformCrdGroup, Kind: "SpringCloudAppMySQLAssociation"}:    springcloudappmysqlassociation.Setup,
		schema.GroupKind{Group: appplatformCrdGroup, Kind: "SpringCloudAppRedisAssociation"}:    springcloudappredisassociation.Setup,
		schema.GroupKind{Group: appplatformCrdGroup, Kind: "SpringCloudBuildDeployment"}:        springcloudbuilddeployment.Setup,
		schema.GroupKind{Group: appplatformCrdGroup, Kind: "SpringCloudBuildPackBinding"}:       springcloudbuildpackbinding.Setup,
		schema.GroupKind{Group: appplatformCrdGroup, Kind: "SpringCloudBuilder"}:                springcloudbuilder.Setup,
		schema.GroupKind{Group: appplatformCrdGroup, Kind: "SpringCloudCertificate"}:            springcloudcertificate.Setup,
		schema.GroupKind{Group: appplatformCrdGroup, Kind: "SpringCloudConfigurationService"}:   springcloudconfigurationservice.Setup,
		schema.GroupKind{Group: appplatformCrdGroup, Kind: "SpringCloudContainerDeployment"}:    springcloudcontainerdeployment.Setup,
		schema.GroupKind{Group: appplatformCrdGroup, Kind: "SpringCloudCustomDomain"}:           springcloudcustomdomain.Setup,
		schema.GroupKind{Group: appplatformCrdGroup, Kind: "SpringCloudCustomizedAccelerator"}:  springcloudcustomizedaccelerator.Setup,
		schema.GroupKind{Group: appplatformCrdGroup, Kind: "SpringCloudDevToolPortal"}:          springclouddevtoolportal.Setup,
		schema.GroupKind{Group: appplatformCrdGroup, Kind: "SpringCloudGateway"}:                springcloudgateway.Setup,
		schema.GroupKind{Group: appplatformCrdGroup, Kind: "SpringCloudGatewayCustomDomain"}:    springcloudgatewaycustomdomain.Setup,
		schema.GroupKind{Group: appplatformCrdGroup, Kind: "SpringCloudJavaDeployment"}:         springcloudjavadeployment.Setup,
		schema.GroupKind{Group: appplatformCrdGroup, Kind: "SpringCloudService"}:                springcloudservice.Setup,
		schema.GroupKind{Group: appplatformCrdGroup, Kind: "SpringCloudStorage"}:                springcloudstorage.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
