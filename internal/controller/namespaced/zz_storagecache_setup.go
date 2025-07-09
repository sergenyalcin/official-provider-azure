// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	hpccache "github.com/upbound/provider-azure/internal/controller/namespaced/storagecache/hpccache"
	hpccacheaccesspolicy "github.com/upbound/provider-azure/internal/controller/namespaced/storagecache/hpccacheaccesspolicy"
	hpccacheblobnfstarget "github.com/upbound/provider-azure/internal/controller/namespaced/storagecache/hpccacheblobnfstarget"
	hpccacheblobtarget "github.com/upbound/provider-azure/internal/controller/namespaced/storagecache/hpccacheblobtarget"
	hpccachenfstarget "github.com/upbound/provider-azure/internal/controller/namespaced/storagecache/hpccachenfstarget"
)

var storagecacheCrdGroup = "storagecache.azure.m.upbound.io"

// Setup_storagecache creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_storagecache(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: storagecacheCrdGroup, Kind: "HPCCache"}:              hpccache.Setup,
		schema.GroupKind{Group: storagecacheCrdGroup, Kind: "HPCCacheAccessPolicy"}:  hpccacheaccesspolicy.Setup,
		schema.GroupKind{Group: storagecacheCrdGroup, Kind: "HPCCacheBlobNFSTarget"}: hpccacheblobnfstarget.Setup,
		schema.GroupKind{Group: storagecacheCrdGroup, Kind: "HPCCacheBlobTarget"}:    hpccacheblobtarget.Setup,
		schema.GroupKind{Group: storagecacheCrdGroup, Kind: "HPCCacheNFSTarget"}:     hpccachenfstarget.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
