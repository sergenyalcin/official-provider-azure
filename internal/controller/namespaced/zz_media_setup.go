// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	asset "github.com/upbound/provider-azure/internal/controller/namespaced/media/asset"
	assetfilter "github.com/upbound/provider-azure/internal/controller/namespaced/media/assetfilter"
	contentkeypolicy "github.com/upbound/provider-azure/internal/controller/namespaced/media/contentkeypolicy"
	job "github.com/upbound/provider-azure/internal/controller/namespaced/media/job"
	liveevent "github.com/upbound/provider-azure/internal/controller/namespaced/media/liveevent"
	liveeventoutput "github.com/upbound/provider-azure/internal/controller/namespaced/media/liveeventoutput"
	servicesaccount "github.com/upbound/provider-azure/internal/controller/namespaced/media/servicesaccount"
	servicesaccountfilter "github.com/upbound/provider-azure/internal/controller/namespaced/media/servicesaccountfilter"
	streamingendpoint "github.com/upbound/provider-azure/internal/controller/namespaced/media/streamingendpoint"
	streaminglocator "github.com/upbound/provider-azure/internal/controller/namespaced/media/streaminglocator"
	streamingpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/media/streamingpolicy"
	transform "github.com/upbound/provider-azure/internal/controller/namespaced/media/transform"
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
