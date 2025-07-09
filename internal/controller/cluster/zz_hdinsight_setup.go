// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	hadoopcluster "github.com/upbound/provider-azure/internal/controller/cluster/hdinsight/hadoopcluster"
	hbasecluster "github.com/upbound/provider-azure/internal/controller/cluster/hdinsight/hbasecluster"
	interactivequerycluster "github.com/upbound/provider-azure/internal/controller/cluster/hdinsight/interactivequerycluster"
	kafkacluster "github.com/upbound/provider-azure/internal/controller/cluster/hdinsight/kafkacluster"
	sparkcluster "github.com/upbound/provider-azure/internal/controller/cluster/hdinsight/sparkcluster"
)

var hdinsightCrdGroup = "hdinsight.azure.upbound.io"

// Setup_hdinsight creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_hdinsight(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: hdinsightCrdGroup, Kind: "HBaseCluster"}:            hbasecluster.Setup,
		schema.GroupKind{Group: hdinsightCrdGroup, Kind: "HadoopCluster"}:           hadoopcluster.Setup,
		schema.GroupKind{Group: hdinsightCrdGroup, Kind: "InteractiveQueryCluster"}: interactivequerycluster.Setup,
		schema.GroupKind{Group: hdinsightCrdGroup, Kind: "KafkaCluster"}:            kafkacluster.Setup,
		schema.GroupKind{Group: hdinsightCrdGroup, Kind: "SparkCluster"}:            sparkcluster.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
