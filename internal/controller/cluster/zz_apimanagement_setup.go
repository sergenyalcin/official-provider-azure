// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

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
	certificate "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/certificate"
	customdomain "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/customdomain"
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
	policy "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/policy"
	product "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/product"
	productapi "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/productapi"
	productpolicy "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/productpolicy"
	producttag "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/producttag"
	rediscache "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/rediscache"
	subscription "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/subscription"
	tag "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/tag"
	user "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/user"
)

var apimanagementCrdGroup = "apimanagement.azure.upbound.io"

// Setup_apimanagement creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_apimanagement(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "API"}:                        api.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "APIDiagnostic"}:              apidiagnostic.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "APIOperation"}:               apioperation.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "APIOperationPolicy"}:         apioperationpolicy.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "APIOperationTag"}:            apioperationtag.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "APIPolicy"}:                  apipolicy.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "APIRelease"}:                 apirelease.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "APISchema"}:                  apischema.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "APITag"}:                     apitag.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "APIVersionSet"}:              apiversionset.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "AuthorizationServer"}:        authorizationserver.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "Backend"}:                    backend.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "Certificate"}:                certificate.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "CustomDomain"}:               customdomain.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "Diagnostic"}:                 diagnostic.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "EmailTemplate"}:              emailtemplate.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "Gateway"}:                    gateway.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "GatewayAPI"}:                 gatewayapi.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "GlobalSchema"}:               globalschema.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "IdentityProviderAAD"}:        identityprovideraad.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "IdentityProviderFacebook"}:   identityproviderfacebook.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "IdentityProviderGoogle"}:     identityprovidergoogle.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "IdentityProviderMicrosoft"}:  identityprovidermicrosoft.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "IdentityProviderTwitter"}:    identityprovidertwitter.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "Logger"}:                     logger.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "Management"}:                 management.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "NamedValue"}:                 namedvalue.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "NotificationRecipientEmail"}: notificationrecipientemail.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "NotificationRecipientUser"}:  notificationrecipientuser.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "OpenIDConnectProvider"}:      openidconnectprovider.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "Policy"}:                     policy.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "Product"}:                    product.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "ProductAPI"}:                 productapi.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "ProductPolicy"}:              productpolicy.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "ProductTag"}:                 producttag.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "RedisCache"}:                 rediscache.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "Subscription"}:               subscription.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "Tag"}:                        tag.Setup,
		schema.GroupKind{Group: apimanagementCrdGroup, Kind: "User"}:                       user.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
