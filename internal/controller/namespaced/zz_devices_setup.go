// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	iothub "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothub"
	iothubcertificate "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubcertificate"
	iothubconsumergroup "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubconsumergroup"
	iothubdps "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubdps"
	iothubdpscertificate "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubdpscertificate"
	iothubdpssharedaccesspolicy "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubdpssharedaccesspolicy"
	iothubendpointeventhub "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubendpointeventhub"
	iothubendpointservicebusqueue "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubendpointservicebusqueue"
	iothubendpointservicebustopic "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubendpointservicebustopic"
	iothubendpointstoragecontainer "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubendpointstoragecontainer"
	iothubenrichment "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubenrichment"
	iothubfallbackroute "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubfallbackroute"
	iothubroute "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubroute"
	iothubsharedaccesspolicy "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubsharedaccesspolicy"
)

var devicesCrdGroup = "devices.azure.m.upbound.io"

// Setup_devices creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_devices(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: devicesCrdGroup, Kind: "IOTHub"}:                         iothub.Setup,
		schema.GroupKind{Group: devicesCrdGroup, Kind: "IOTHubCertificate"}:              iothubcertificate.Setup,
		schema.GroupKind{Group: devicesCrdGroup, Kind: "IOTHubConsumerGroup"}:            iothubconsumergroup.Setup,
		schema.GroupKind{Group: devicesCrdGroup, Kind: "IOTHubDPS"}:                      iothubdps.Setup,
		schema.GroupKind{Group: devicesCrdGroup, Kind: "IOTHubDPSCertificate"}:           iothubdpscertificate.Setup,
		schema.GroupKind{Group: devicesCrdGroup, Kind: "IOTHubDPSSharedAccessPolicy"}:    iothubdpssharedaccesspolicy.Setup,
		schema.GroupKind{Group: devicesCrdGroup, Kind: "IOTHubEndpointEventHub"}:         iothubendpointeventhub.Setup,
		schema.GroupKind{Group: devicesCrdGroup, Kind: "IOTHubEndpointServiceBusQueue"}:  iothubendpointservicebusqueue.Setup,
		schema.GroupKind{Group: devicesCrdGroup, Kind: "IOTHubEndpointServiceBusTopic"}:  iothubendpointservicebustopic.Setup,
		schema.GroupKind{Group: devicesCrdGroup, Kind: "IOTHubEndpointStorageContainer"}: iothubendpointstoragecontainer.Setup,
		schema.GroupKind{Group: devicesCrdGroup, Kind: "IOTHubEnrichment"}:               iothubenrichment.Setup,
		schema.GroupKind{Group: devicesCrdGroup, Kind: "IOTHubFallbackRoute"}:            iothubfallbackroute.Setup,
		schema.GroupKind{Group: devicesCrdGroup, Kind: "IOTHubRoute"}:                    iothubroute.Setup,
		schema.GroupKind{Group: devicesCrdGroup, Kind: "IOTHubSharedAccessPolicy"}:       iothubsharedaccesspolicy.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, devicesCrdGroup, o); err != nil {
		return err
	}
	return nil
}
