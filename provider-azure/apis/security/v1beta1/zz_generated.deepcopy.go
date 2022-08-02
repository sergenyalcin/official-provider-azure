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
func (in *AdditionalWorkspaceObservation) DeepCopyInto(out *AdditionalWorkspaceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalWorkspaceObservation.
func (in *AdditionalWorkspaceObservation) DeepCopy() *AdditionalWorkspaceObservation {
	if in == nil {
		return nil
	}
	out := new(AdditionalWorkspaceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdditionalWorkspaceParameters) DeepCopyInto(out *AdditionalWorkspaceParameters) {
	*out = *in
	if in.DataTypes != nil {
		in, out := &in.DataTypes, &out.DataTypes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.WorkspaceID != nil {
		in, out := &in.WorkspaceID, &out.WorkspaceID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalWorkspaceParameters.
func (in *AdditionalWorkspaceParameters) DeepCopy() *AdditionalWorkspaceParameters {
	if in == nil {
		return nil
	}
	out := new(AdditionalWorkspaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedThreatProtection) DeepCopyInto(out *AdvancedThreatProtection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedThreatProtection.
func (in *AdvancedThreatProtection) DeepCopy() *AdvancedThreatProtection {
	if in == nil {
		return nil
	}
	out := new(AdvancedThreatProtection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdvancedThreatProtection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedThreatProtectionList) DeepCopyInto(out *AdvancedThreatProtectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AdvancedThreatProtection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedThreatProtectionList.
func (in *AdvancedThreatProtectionList) DeepCopy() *AdvancedThreatProtectionList {
	if in == nil {
		return nil
	}
	out := new(AdvancedThreatProtectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdvancedThreatProtectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedThreatProtectionObservation) DeepCopyInto(out *AdvancedThreatProtectionObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedThreatProtectionObservation.
func (in *AdvancedThreatProtectionObservation) DeepCopy() *AdvancedThreatProtectionObservation {
	if in == nil {
		return nil
	}
	out := new(AdvancedThreatProtectionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedThreatProtectionParameters) DeepCopyInto(out *AdvancedThreatProtectionParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.TargetResourceID != nil {
		in, out := &in.TargetResourceID, &out.TargetResourceID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedThreatProtectionParameters.
func (in *AdvancedThreatProtectionParameters) DeepCopy() *AdvancedThreatProtectionParameters {
	if in == nil {
		return nil
	}
	out := new(AdvancedThreatProtectionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedThreatProtectionSpec) DeepCopyInto(out *AdvancedThreatProtectionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedThreatProtectionSpec.
func (in *AdvancedThreatProtectionSpec) DeepCopy() *AdvancedThreatProtectionSpec {
	if in == nil {
		return nil
	}
	out := new(AdvancedThreatProtectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedThreatProtectionStatus) DeepCopyInto(out *AdvancedThreatProtectionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedThreatProtectionStatus.
func (in *AdvancedThreatProtectionStatus) DeepCopy() *AdvancedThreatProtectionStatus {
	if in == nil {
		return nil
	}
	out := new(AdvancedThreatProtectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowRuleObservation) DeepCopyInto(out *AllowRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowRuleObservation.
func (in *AllowRuleObservation) DeepCopy() *AllowRuleObservation {
	if in == nil {
		return nil
	}
	out := new(AllowRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowRuleParameters) DeepCopyInto(out *AllowRuleParameters) {
	*out = *in
	if in.ConnectionFromIpsNotAllowed != nil {
		in, out := &in.ConnectionFromIpsNotAllowed, &out.ConnectionFromIpsNotAllowed
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ConnectionToIpsNotAllowed != nil {
		in, out := &in.ConnectionToIpsNotAllowed, &out.ConnectionToIpsNotAllowed
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.LocalUsersNotAllowed != nil {
		in, out := &in.LocalUsersNotAllowed, &out.LocalUsersNotAllowed
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ProcessesNotAllowed != nil {
		in, out := &in.ProcessesNotAllowed, &out.ProcessesNotAllowed
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowRuleParameters.
func (in *AllowRuleParameters) DeepCopy() *AllowRuleParameters {
	if in == nil {
		return nil
	}
	out := new(AllowRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTSecurityDeviceGroup) DeepCopyInto(out *IOTSecurityDeviceGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTSecurityDeviceGroup.
func (in *IOTSecurityDeviceGroup) DeepCopy() *IOTSecurityDeviceGroup {
	if in == nil {
		return nil
	}
	out := new(IOTSecurityDeviceGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IOTSecurityDeviceGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTSecurityDeviceGroupList) DeepCopyInto(out *IOTSecurityDeviceGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IOTSecurityDeviceGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTSecurityDeviceGroupList.
func (in *IOTSecurityDeviceGroupList) DeepCopy() *IOTSecurityDeviceGroupList {
	if in == nil {
		return nil
	}
	out := new(IOTSecurityDeviceGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IOTSecurityDeviceGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTSecurityDeviceGroupObservation) DeepCopyInto(out *IOTSecurityDeviceGroupObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTSecurityDeviceGroupObservation.
func (in *IOTSecurityDeviceGroupObservation) DeepCopy() *IOTSecurityDeviceGroupObservation {
	if in == nil {
		return nil
	}
	out := new(IOTSecurityDeviceGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTSecurityDeviceGroupParameters) DeepCopyInto(out *IOTSecurityDeviceGroupParameters) {
	*out = *in
	if in.AllowRule != nil {
		in, out := &in.AllowRule, &out.AllowRule
		*out = make([]AllowRuleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IOTHubID != nil {
		in, out := &in.IOTHubID, &out.IOTHubID
		*out = new(string)
		**out = **in
	}
	if in.IOTHubIDRef != nil {
		in, out := &in.IOTHubIDRef, &out.IOTHubIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.IOTHubIDSelector != nil {
		in, out := &in.IOTHubIDSelector, &out.IOTHubIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.RangeRule != nil {
		in, out := &in.RangeRule, &out.RangeRule
		*out = make([]RangeRuleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTSecurityDeviceGroupParameters.
func (in *IOTSecurityDeviceGroupParameters) DeepCopy() *IOTSecurityDeviceGroupParameters {
	if in == nil {
		return nil
	}
	out := new(IOTSecurityDeviceGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTSecurityDeviceGroupSpec) DeepCopyInto(out *IOTSecurityDeviceGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTSecurityDeviceGroupSpec.
func (in *IOTSecurityDeviceGroupSpec) DeepCopy() *IOTSecurityDeviceGroupSpec {
	if in == nil {
		return nil
	}
	out := new(IOTSecurityDeviceGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTSecurityDeviceGroupStatus) DeepCopyInto(out *IOTSecurityDeviceGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTSecurityDeviceGroupStatus.
func (in *IOTSecurityDeviceGroupStatus) DeepCopy() *IOTSecurityDeviceGroupStatus {
	if in == nil {
		return nil
	}
	out := new(IOTSecurityDeviceGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTSecuritySolution) DeepCopyInto(out *IOTSecuritySolution) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTSecuritySolution.
func (in *IOTSecuritySolution) DeepCopy() *IOTSecuritySolution {
	if in == nil {
		return nil
	}
	out := new(IOTSecuritySolution)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IOTSecuritySolution) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTSecuritySolutionList) DeepCopyInto(out *IOTSecuritySolutionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IOTSecuritySolution, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTSecuritySolutionList.
func (in *IOTSecuritySolutionList) DeepCopy() *IOTSecuritySolutionList {
	if in == nil {
		return nil
	}
	out := new(IOTSecuritySolutionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IOTSecuritySolutionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTSecuritySolutionObservation) DeepCopyInto(out *IOTSecuritySolutionObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTSecuritySolutionObservation.
func (in *IOTSecuritySolutionObservation) DeepCopy() *IOTSecuritySolutionObservation {
	if in == nil {
		return nil
	}
	out := new(IOTSecuritySolutionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTSecuritySolutionParameters) DeepCopyInto(out *IOTSecuritySolutionParameters) {
	*out = *in
	if in.AdditionalWorkspace != nil {
		in, out := &in.AdditionalWorkspace, &out.AdditionalWorkspace
		*out = make([]AdditionalWorkspaceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DisabledDataSources != nil {
		in, out := &in.DisabledDataSources, &out.DisabledDataSources
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EventsToExport != nil {
		in, out := &in.EventsToExport, &out.EventsToExport
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IOTHubIds != nil {
		in, out := &in.IOTHubIds, &out.IOTHubIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.LogAnalyticsWorkspaceID != nil {
		in, out := &in.LogAnalyticsWorkspaceID, &out.LogAnalyticsWorkspaceID
		*out = new(string)
		**out = **in
	}
	if in.LogUnmaskedIpsEnabled != nil {
		in, out := &in.LogUnmaskedIpsEnabled, &out.LogUnmaskedIpsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.QueryForResources != nil {
		in, out := &in.QueryForResources, &out.QueryForResources
		*out = new(string)
		**out = **in
	}
	if in.QuerySubscriptionIds != nil {
		in, out := &in.QuerySubscriptionIds, &out.QuerySubscriptionIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RecommendationsEnabled != nil {
		in, out := &in.RecommendationsEnabled, &out.RecommendationsEnabled
		*out = make([]RecommendationsEnabledParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTSecuritySolutionParameters.
func (in *IOTSecuritySolutionParameters) DeepCopy() *IOTSecuritySolutionParameters {
	if in == nil {
		return nil
	}
	out := new(IOTSecuritySolutionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTSecuritySolutionSpec) DeepCopyInto(out *IOTSecuritySolutionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTSecuritySolutionSpec.
func (in *IOTSecuritySolutionSpec) DeepCopy() *IOTSecuritySolutionSpec {
	if in == nil {
		return nil
	}
	out := new(IOTSecuritySolutionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTSecuritySolutionStatus) DeepCopyInto(out *IOTSecuritySolutionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTSecuritySolutionStatus.
func (in *IOTSecuritySolutionStatus) DeepCopy() *IOTSecuritySolutionStatus {
	if in == nil {
		return nil
	}
	out := new(IOTSecuritySolutionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RangeRuleObservation) DeepCopyInto(out *RangeRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RangeRuleObservation.
func (in *RangeRuleObservation) DeepCopy() *RangeRuleObservation {
	if in == nil {
		return nil
	}
	out := new(RangeRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RangeRuleParameters) DeepCopyInto(out *RangeRuleParameters) {
	*out = *in
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = new(float64)
		**out = **in
	}
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RangeRuleParameters.
func (in *RangeRuleParameters) DeepCopy() *RangeRuleParameters {
	if in == nil {
		return nil
	}
	out := new(RangeRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecommendationsEnabledObservation) DeepCopyInto(out *RecommendationsEnabledObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecommendationsEnabledObservation.
func (in *RecommendationsEnabledObservation) DeepCopy() *RecommendationsEnabledObservation {
	if in == nil {
		return nil
	}
	out := new(RecommendationsEnabledObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecommendationsEnabledParameters) DeepCopyInto(out *RecommendationsEnabledParameters) {
	*out = *in
	if in.AcrAuthentication != nil {
		in, out := &in.AcrAuthentication, &out.AcrAuthentication
		*out = new(bool)
		**out = **in
	}
	if in.AgentSendUnutilizedMsg != nil {
		in, out := &in.AgentSendUnutilizedMsg, &out.AgentSendUnutilizedMsg
		*out = new(bool)
		**out = **in
	}
	if in.Baseline != nil {
		in, out := &in.Baseline, &out.Baseline
		*out = new(bool)
		**out = **in
	}
	if in.EdgeHubMemOptimize != nil {
		in, out := &in.EdgeHubMemOptimize, &out.EdgeHubMemOptimize
		*out = new(bool)
		**out = **in
	}
	if in.EdgeLoggingOption != nil {
		in, out := &in.EdgeLoggingOption, &out.EdgeLoggingOption
		*out = new(bool)
		**out = **in
	}
	if in.IPFilterDenyAll != nil {
		in, out := &in.IPFilterDenyAll, &out.IPFilterDenyAll
		*out = new(bool)
		**out = **in
	}
	if in.IPFilterPermissiveRule != nil {
		in, out := &in.IPFilterPermissiveRule, &out.IPFilterPermissiveRule
		*out = new(bool)
		**out = **in
	}
	if in.InconsistentModuleSettings != nil {
		in, out := &in.InconsistentModuleSettings, &out.InconsistentModuleSettings
		*out = new(bool)
		**out = **in
	}
	if in.InstallAgent != nil {
		in, out := &in.InstallAgent, &out.InstallAgent
		*out = new(bool)
		**out = **in
	}
	if in.OpenPorts != nil {
		in, out := &in.OpenPorts, &out.OpenPorts
		*out = new(bool)
		**out = **in
	}
	if in.PermissiveFirewallPolicy != nil {
		in, out := &in.PermissiveFirewallPolicy, &out.PermissiveFirewallPolicy
		*out = new(bool)
		**out = **in
	}
	if in.PermissiveInputFirewallRules != nil {
		in, out := &in.PermissiveInputFirewallRules, &out.PermissiveInputFirewallRules
		*out = new(bool)
		**out = **in
	}
	if in.PermissiveOutputFirewallRules != nil {
		in, out := &in.PermissiveOutputFirewallRules, &out.PermissiveOutputFirewallRules
		*out = new(bool)
		**out = **in
	}
	if in.PrivilegedDockerOptions != nil {
		in, out := &in.PrivilegedDockerOptions, &out.PrivilegedDockerOptions
		*out = new(bool)
		**out = **in
	}
	if in.SharedCredentials != nil {
		in, out := &in.SharedCredentials, &out.SharedCredentials
		*out = new(bool)
		**out = **in
	}
	if in.VulnerableTLSCipherSuite != nil {
		in, out := &in.VulnerableTLSCipherSuite, &out.VulnerableTLSCipherSuite
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecommendationsEnabledParameters.
func (in *RecommendationsEnabledParameters) DeepCopy() *RecommendationsEnabledParameters {
	if in == nil {
		return nil
	}
	out := new(RecommendationsEnabledParameters)
	in.DeepCopyInto(out)
	return out
}
