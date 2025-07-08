// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

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

// Setup_media creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_media(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		asset.Setup,
		assetfilter.Setup,
		contentkeypolicy.Setup,
		job.Setup,
		liveevent.Setup,
		liveeventoutput.Setup,
		servicesaccount.Setup,
		servicesaccountfilter.Setup,
		streamingendpoint.Setup,
		streaminglocator.Setup,
		streamingpolicy.Setup,
		transform.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
