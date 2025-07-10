// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	botchannelalexa "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botchannelalexa"
	botchanneldirectline "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botchanneldirectline"
	botchannelline "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botchannelline"
	botchannelmsteams "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botchannelmsteams"
	botchannelslack "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botchannelslack"
	botchannelsms "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botchannelsms"
	botchannelsregistration "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botchannelsregistration"
	botchannelwebchat "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botchannelwebchat"
	botconnection "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botconnection"
	botwebapp "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botwebapp"
)

var botserviceCrdGroup = "botservice.azure.m.upbound.io"

// Setup_botservice creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_botservice(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: botserviceCrdGroup, Kind: "BotChannelAlexa"}:         botchannelalexa.Setup,
		schema.GroupKind{Group: botserviceCrdGroup, Kind: "BotChannelDirectLine"}:    botchanneldirectline.Setup,
		schema.GroupKind{Group: botserviceCrdGroup, Kind: "BotChannelLine"}:          botchannelline.Setup,
		schema.GroupKind{Group: botserviceCrdGroup, Kind: "BotChannelMSTeams"}:       botchannelmsteams.Setup,
		schema.GroupKind{Group: botserviceCrdGroup, Kind: "BotChannelSMS"}:           botchannelsms.Setup,
		schema.GroupKind{Group: botserviceCrdGroup, Kind: "BotChannelSlack"}:         botchannelslack.Setup,
		schema.GroupKind{Group: botserviceCrdGroup, Kind: "BotChannelWebChat"}:       botchannelwebchat.Setup,
		schema.GroupKind{Group: botserviceCrdGroup, Kind: "BotChannelsRegistration"}: botchannelsregistration.Setup,
		schema.GroupKind{Group: botserviceCrdGroup, Kind: "BotConnection"}:           botconnection.Setup,
		schema.GroupKind{Group: botserviceCrdGroup, Kind: "BotWebApp"}:               botwebapp.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, botserviceCrdGroup, o); err != nil {
		return err
	}
	return nil
}
