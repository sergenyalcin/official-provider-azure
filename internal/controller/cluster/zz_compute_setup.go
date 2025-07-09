// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/dynamiccrd"

	availabilityset "github.com/upbound/provider-azure/internal/controller/cluster/compute/availabilityset"
	capacityreservation "github.com/upbound/provider-azure/internal/controller/cluster/compute/capacityreservation"
	capacityreservationgroup "github.com/upbound/provider-azure/internal/controller/cluster/compute/capacityreservationgroup"
	dedicatedhost "github.com/upbound/provider-azure/internal/controller/cluster/compute/dedicatedhost"
	diskaccess "github.com/upbound/provider-azure/internal/controller/cluster/compute/diskaccess"
	diskencryptionset "github.com/upbound/provider-azure/internal/controller/cluster/compute/diskencryptionset"
	galleryapplication "github.com/upbound/provider-azure/internal/controller/cluster/compute/galleryapplication"
	galleryapplicationversion "github.com/upbound/provider-azure/internal/controller/cluster/compute/galleryapplicationversion"
	image "github.com/upbound/provider-azure/internal/controller/cluster/compute/image"
	linuxvirtualmachine "github.com/upbound/provider-azure/internal/controller/cluster/compute/linuxvirtualmachine"
	linuxvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/cluster/compute/linuxvirtualmachinescaleset"
	manageddisk "github.com/upbound/provider-azure/internal/controller/cluster/compute/manageddisk"
	manageddisksastoken "github.com/upbound/provider-azure/internal/controller/cluster/compute/manageddisksastoken"
	orchestratedvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/cluster/compute/orchestratedvirtualmachinescaleset"
	proximityplacementgroup "github.com/upbound/provider-azure/internal/controller/cluster/compute/proximityplacementgroup"
	sharedimage "github.com/upbound/provider-azure/internal/controller/cluster/compute/sharedimage"
	sharedimagegallery "github.com/upbound/provider-azure/internal/controller/cluster/compute/sharedimagegallery"
	snapshot "github.com/upbound/provider-azure/internal/controller/cluster/compute/snapshot"
	sshpublickey "github.com/upbound/provider-azure/internal/controller/cluster/compute/sshpublickey"
	virtualmachinedatadiskattachment "github.com/upbound/provider-azure/internal/controller/cluster/compute/virtualmachinedatadiskattachment"
	virtualmachineextension "github.com/upbound/provider-azure/internal/controller/cluster/compute/virtualmachineextension"
	virtualmachineruncommand "github.com/upbound/provider-azure/internal/controller/cluster/compute/virtualmachineruncommand"
	windowsvirtualmachine "github.com/upbound/provider-azure/internal/controller/cluster/compute/windowsvirtualmachine"
	windowsvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/cluster/compute/windowsvirtualmachinescaleset"
)

var computeCrdGroup = "compute.azure.upbound.io"

// Setup_compute creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_compute(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: computeCrdGroup, Kind: "AvailabilitySet"}:                    availabilityset.Setup,
		schema.GroupKind{Group: computeCrdGroup, Kind: "CapacityReservation"}:                capacityreservation.Setup,
		schema.GroupKind{Group: computeCrdGroup, Kind: "CapacityReservationGroup"}:           capacityreservationgroup.Setup,
		schema.GroupKind{Group: computeCrdGroup, Kind: "DedicatedHost"}:                      dedicatedhost.Setup,
		schema.GroupKind{Group: computeCrdGroup, Kind: "DiskAccess"}:                         diskaccess.Setup,
		schema.GroupKind{Group: computeCrdGroup, Kind: "DiskEncryptionSet"}:                  diskencryptionset.Setup,
		schema.GroupKind{Group: computeCrdGroup, Kind: "GalleryApplication"}:                 galleryapplication.Setup,
		schema.GroupKind{Group: computeCrdGroup, Kind: "GalleryApplicationVersion"}:          galleryapplicationversion.Setup,
		schema.GroupKind{Group: computeCrdGroup, Kind: "Image"}:                              image.Setup,
		schema.GroupKind{Group: computeCrdGroup, Kind: "LinuxVirtualMachine"}:                linuxvirtualmachine.Setup,
		schema.GroupKind{Group: computeCrdGroup, Kind: "LinuxVirtualMachineScaleSet"}:        linuxvirtualmachinescaleset.Setup,
		schema.GroupKind{Group: computeCrdGroup, Kind: "ManagedDisk"}:                        manageddisk.Setup,
		schema.GroupKind{Group: computeCrdGroup, Kind: "ManagedDiskSASToken"}:                manageddisksastoken.Setup,
		schema.GroupKind{Group: computeCrdGroup, Kind: "OrchestratedVirtualMachineScaleSet"}: orchestratedvirtualmachinescaleset.Setup,
		schema.GroupKind{Group: computeCrdGroup, Kind: "ProximityPlacementGroup"}:            proximityplacementgroup.Setup,
		schema.GroupKind{Group: computeCrdGroup, Kind: "SSHPublicKey"}:                       sshpublickey.Setup,
		schema.GroupKind{Group: computeCrdGroup, Kind: "SharedImage"}:                        sharedimage.Setup,
		schema.GroupKind{Group: computeCrdGroup, Kind: "SharedImageGallery"}:                 sharedimagegallery.Setup,
		schema.GroupKind{Group: computeCrdGroup, Kind: "Snapshot"}:                           snapshot.Setup,
		schema.GroupKind{Group: computeCrdGroup, Kind: "VirtualMachineDataDiskAttachment"}:   virtualmachinedatadiskattachment.Setup,
		schema.GroupKind{Group: computeCrdGroup, Kind: "VirtualMachineExtension"}:            virtualmachineextension.Setup,
		schema.GroupKind{Group: computeCrdGroup, Kind: "VirtualMachineRunCommand"}:           virtualmachineruncommand.Setup,
		schema.GroupKind{Group: computeCrdGroup, Kind: "WindowsVirtualMachine"}:              windowsvirtualmachine.Setup,
		schema.GroupKind{Group: computeCrdGroup, Kind: "WindowsVirtualMachineScaleSet"}:      windowsvirtualmachinescaleset.Setup,
	}
	if err := dynamiccrd.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
