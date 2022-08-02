//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APNSCredentialObservation) DeepCopyInto(out *APNSCredentialObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APNSCredentialObservation.
func (in *APNSCredentialObservation) DeepCopy() *APNSCredentialObservation {
	if in == nil {
		return nil
	}
	out := new(APNSCredentialObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APNSCredentialParameters) DeepCopyInto(out *APNSCredentialParameters) {
	*out = *in
	if in.ApplicationMode != nil {
		in, out := &in.ApplicationMode, &out.ApplicationMode
		*out = new(string)
		**out = **in
	}
	if in.BundleID != nil {
		in, out := &in.BundleID, &out.BundleID
		*out = new(string)
		**out = **in
	}
	if in.KeyID != nil {
		in, out := &in.KeyID, &out.KeyID
		*out = new(string)
		**out = **in
	}
	if in.TeamID != nil {
		in, out := &in.TeamID, &out.TeamID
		*out = new(string)
		**out = **in
	}
	out.TokenSecretRef = in.TokenSecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APNSCredentialParameters.
func (in *APNSCredentialParameters) DeepCopy() *APNSCredentialParameters {
	if in == nil {
		return nil
	}
	out := new(APNSCredentialParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCMCredentialObservation) DeepCopyInto(out *GCMCredentialObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCMCredentialObservation.
func (in *GCMCredentialObservation) DeepCopy() *GCMCredentialObservation {
	if in == nil {
		return nil
	}
	out := new(GCMCredentialObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCMCredentialParameters) DeepCopyInto(out *GCMCredentialParameters) {
	*out = *in
	out.APIKeySecretRef = in.APIKeySecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCMCredentialParameters.
func (in *GCMCredentialParameters) DeepCopy() *GCMCredentialParameters {
	if in == nil {
		return nil
	}
	out := new(GCMCredentialParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationHub) DeepCopyInto(out *NotificationHub) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationHub.
func (in *NotificationHub) DeepCopy() *NotificationHub {
	if in == nil {
		return nil
	}
	out := new(NotificationHub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NotificationHub) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationHubList) DeepCopyInto(out *NotificationHubList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NotificationHub, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationHubList.
func (in *NotificationHubList) DeepCopy() *NotificationHubList {
	if in == nil {
		return nil
	}
	out := new(NotificationHubList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NotificationHubList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationHubObservation) DeepCopyInto(out *NotificationHubObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationHubObservation.
func (in *NotificationHubObservation) DeepCopy() *NotificationHubObservation {
	if in == nil {
		return nil
	}
	out := new(NotificationHubObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationHubParameters) DeepCopyInto(out *NotificationHubParameters) {
	*out = *in
	if in.APNSCredential != nil {
		in, out := &in.APNSCredential, &out.APNSCredential
		*out = make([]APNSCredentialParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GCMCredential != nil {
		in, out := &in.GCMCredential, &out.GCMCredential
		*out = make([]GCMCredentialParameters, len(*in))
		copy(*out, *in)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.NamespaceName != nil {
		in, out := &in.NamespaceName, &out.NamespaceName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationHubParameters.
func (in *NotificationHubParameters) DeepCopy() *NotificationHubParameters {
	if in == nil {
		return nil
	}
	out := new(NotificationHubParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationHubSpec) DeepCopyInto(out *NotificationHubSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationHubSpec.
func (in *NotificationHubSpec) DeepCopy() *NotificationHubSpec {
	if in == nil {
		return nil
	}
	out := new(NotificationHubSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationHubStatus) DeepCopyInto(out *NotificationHubStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationHubStatus.
func (in *NotificationHubStatus) DeepCopy() *NotificationHubStatus {
	if in == nil {
		return nil
	}
	out := new(NotificationHubStatus)
	in.DeepCopyInto(out)
	return out
}
