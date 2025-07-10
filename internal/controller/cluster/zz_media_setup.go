// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	asset "github.com/upbound/provider-azure/internal/controller/cluster/media/asset"
	assetfilter "github.com/upbound/provider-azure/internal/controller/cluster/media/assetfilter"
	contentkeypolicy "github.com/upbound/provider-azure/internal/controller/cluster/media/contentkeypolicy"
	job "github.com/upbound/provider-azure/internal/controller/cluster/media/job"
	liveevent "github.com/upbound/provider-azure/internal/controller/cluster/media/liveevent"
	liveeventoutput "github.com/upbound/provider-azure/internal/controller/cluster/media/liveeventoutput"
	servicesaccount "github.com/upbound/provider-azure/internal/controller/cluster/media/servicesaccount"
	servicesaccountfilter "github.com/upbound/provider-azure/internal/controller/cluster/media/servicesaccountfilter"
	streamingendpoint "github.com/upbound/provider-azure/internal/controller/cluster/media/streamingendpoint"
	streaminglocator "github.com/upbound/provider-azure/internal/controller/cluster/media/streaminglocator"
	streamingpolicy "github.com/upbound/provider-azure/internal/controller/cluster/media/streamingpolicy"
	transform "github.com/upbound/provider-azure/internal/controller/cluster/media/transform"
)

var mediaCrdGroup = "media.azure.upbound.io"

// Setup_media creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_media(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: mediaCrdGroup, Kind: "Asset"}:                 asset.Setup,
		schema.GroupKind{Group: mediaCrdGroup, Kind: "AssetFilter"}:           assetfilter.Setup,
		schema.GroupKind{Group: mediaCrdGroup, Kind: "ContentKeyPolicy"}:      contentkeypolicy.Setup,
		schema.GroupKind{Group: mediaCrdGroup, Kind: "Job"}:                   job.Setup,
		schema.GroupKind{Group: mediaCrdGroup, Kind: "LiveEvent"}:             liveevent.Setup,
		schema.GroupKind{Group: mediaCrdGroup, Kind: "LiveEventOutput"}:       liveeventoutput.Setup,
		schema.GroupKind{Group: mediaCrdGroup, Kind: "ServicesAccount"}:       servicesaccount.Setup,
		schema.GroupKind{Group: mediaCrdGroup, Kind: "ServicesAccountFilter"}: servicesaccountfilter.Setup,
		schema.GroupKind{Group: mediaCrdGroup, Kind: "StreamingEndpoint"}:     streamingendpoint.Setup,
		schema.GroupKind{Group: mediaCrdGroup, Kind: "StreamingLocator"}:      streaminglocator.Setup,
		schema.GroupKind{Group: mediaCrdGroup, Kind: "StreamingPolicy"}:       streamingpolicy.Setup,
		schema.GroupKind{Group: mediaCrdGroup, Kind: "Transform"}:             transform.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, mediaCrdGroup, o); err != nil {
		return err
	}
	return nil
}
