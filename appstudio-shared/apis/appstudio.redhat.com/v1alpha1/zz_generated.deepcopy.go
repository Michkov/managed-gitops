//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutomatedPromotionConfiguration) DeepCopyInto(out *AutomatedPromotionConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutomatedPromotionConfiguration.
func (in *AutomatedPromotionConfiguration) DeepCopy() *AutomatedPromotionConfiguration {
	if in == nil {
		return nil
	}
	out := new(AutomatedPromotionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingComponent) DeepCopyInto(out *BindingComponent) {
	*out = *in
	in.Configuration.DeepCopyInto(&out.Configuration)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingComponent.
func (in *BindingComponent) DeepCopy() *BindingComponent {
	if in == nil {
		return nil
	}
	out := new(BindingComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingComponentConfiguration) DeepCopyInto(out *BindingComponentConfiguration) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]EnvVarPair, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingComponentConfiguration.
func (in *BindingComponentConfiguration) DeepCopy() *BindingComponentConfiguration {
	if in == nil {
		return nil
	}
	out := new(BindingComponentConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingComponentGitOpsRepository) DeepCopyInto(out *BindingComponentGitOpsRepository) {
	*out = *in
	if in.GeneratedResources != nil {
		in, out := &in.GeneratedResources, &out.GeneratedResources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingComponentGitOpsRepository.
func (in *BindingComponentGitOpsRepository) DeepCopy() *BindingComponentGitOpsRepository {
	if in == nil {
		return nil
	}
	out := new(BindingComponentGitOpsRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingComponentStatus) DeepCopyInto(out *BindingComponentStatus) {
	*out = *in
	in.GitOpsRepository.DeepCopyInto(&out.GitOpsRepository)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingComponentStatus.
func (in *BindingComponentStatus) DeepCopy() *BindingComponentStatus {
	if in == nil {
		return nil
	}
	out := new(BindingComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingStatusGitOpsDeployment) DeepCopyInto(out *BindingStatusGitOpsDeployment) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingStatusGitOpsDeployment.
func (in *BindingStatusGitOpsDeployment) DeepCopy() *BindingStatusGitOpsDeployment {
	if in == nil {
		return nil
	}
	out := new(BindingStatusGitOpsDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvVarPair) DeepCopyInto(out *EnvVarPair) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVarPair.
func (in *EnvVarPair) DeepCopy() *EnvVarPair {
	if in == nil {
		return nil
	}
	out := new(EnvVarPair)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Environment) DeepCopyInto(out *Environment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Environment.
func (in *Environment) DeepCopy() *Environment {
	if in == nil {
		return nil
	}
	out := new(Environment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Environment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentConfiguration) DeepCopyInto(out *EnvironmentConfiguration) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]EnvVarPair, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentConfiguration.
func (in *EnvironmentConfiguration) DeepCopy() *EnvironmentConfiguration {
	if in == nil {
		return nil
	}
	out := new(EnvironmentConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentList) DeepCopyInto(out *EnvironmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Environment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentList.
func (in *EnvironmentList) DeepCopy() *EnvironmentList {
	if in == nil {
		return nil
	}
	out := new(EnvironmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvironmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpec) DeepCopyInto(out *EnvironmentSpec) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Configuration.DeepCopyInto(&out.Configuration)
	if in.UnstableConfigurationFields != nil {
		in, out := &in.UnstableConfigurationFields, &out.UnstableConfigurationFields
		*out = new(UnstableEnvironmentConfiguration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpec.
func (in *EnvironmentSpec) DeepCopy() *EnvironmentSpec {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentStatus) DeepCopyInto(out *EnvironmentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentStatus.
func (in *EnvironmentStatus) DeepCopy() *EnvironmentStatus {
	if in == nil {
		return nil
	}
	out := new(EnvironmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesClusterCredentials) DeepCopyInto(out *KubernetesClusterCredentials) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesClusterCredentials.
func (in *KubernetesClusterCredentials) DeepCopy() *KubernetesClusterCredentials {
	if in == nil {
		return nil
	}
	out := new(KubernetesClusterCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManualPromotionConfiguration) DeepCopyInto(out *ManualPromotionConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManualPromotionConfiguration.
func (in *ManualPromotionConfiguration) DeepCopy() *ManualPromotionConfiguration {
	if in == nil {
		return nil
	}
	out := new(ManualPromotionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromotionRun) DeepCopyInto(out *PromotionRun) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromotionRun.
func (in *PromotionRun) DeepCopy() *PromotionRun {
	if in == nil {
		return nil
	}
	out := new(PromotionRun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PromotionRun) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromotionRunCondition) DeepCopyInto(out *PromotionRunCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromotionRunCondition.
func (in *PromotionRunCondition) DeepCopy() *PromotionRunCondition {
	if in == nil {
		return nil
	}
	out := new(PromotionRunCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromotionRunEnvironmentStatus) DeepCopyInto(out *PromotionRunEnvironmentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromotionRunEnvironmentStatus.
func (in *PromotionRunEnvironmentStatus) DeepCopy() *PromotionRunEnvironmentStatus {
	if in == nil {
		return nil
	}
	out := new(PromotionRunEnvironmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromotionRunList) DeepCopyInto(out *PromotionRunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PromotionRun, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromotionRunList.
func (in *PromotionRunList) DeepCopy() *PromotionRunList {
	if in == nil {
		return nil
	}
	out := new(PromotionRunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PromotionRunList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromotionRunSpec) DeepCopyInto(out *PromotionRunSpec) {
	*out = *in
	out.ManualPromotion = in.ManualPromotion
	out.AutomatedPromotion = in.AutomatedPromotion
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromotionRunSpec.
func (in *PromotionRunSpec) DeepCopy() *PromotionRunSpec {
	if in == nil {
		return nil
	}
	out := new(PromotionRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromotionRunStatus) DeepCopyInto(out *PromotionRunStatus) {
	*out = *in
	if in.EnvironmentStatus != nil {
		in, out := &in.EnvironmentStatus, &out.EnvironmentStatus
		*out = make([]PromotionRunEnvironmentStatus, len(*in))
		copy(*out, *in)
	}
	if in.ActiveBindings != nil {
		in, out := &in.ActiveBindings, &out.ActiveBindings
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.PromotionStartTime.DeepCopyInto(&out.PromotionStartTime)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]PromotionRunCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromotionRunStatus.
func (in *PromotionRunStatus) DeepCopy() *PromotionRunStatus {
	if in == nil {
		return nil
	}
	out := new(PromotionRunStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Snapshot) DeepCopyInto(out *Snapshot) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Snapshot.
func (in *Snapshot) DeepCopy() *Snapshot {
	if in == nil {
		return nil
	}
	out := new(Snapshot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Snapshot) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotArtifacts) DeepCopyInto(out *SnapshotArtifacts) {
	*out = *in
	if in.UnstableFields != nil {
		in, out := &in.UnstableFields, &out.UnstableFields
		*out = new(v1.JSON)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotArtifacts.
func (in *SnapshotArtifacts) DeepCopy() *SnapshotArtifacts {
	if in == nil {
		return nil
	}
	out := new(SnapshotArtifacts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotComponent) DeepCopyInto(out *SnapshotComponent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotComponent.
func (in *SnapshotComponent) DeepCopy() *SnapshotComponent {
	if in == nil {
		return nil
	}
	out := new(SnapshotComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotEnvironmentBinding) DeepCopyInto(out *SnapshotEnvironmentBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotEnvironmentBinding.
func (in *SnapshotEnvironmentBinding) DeepCopy() *SnapshotEnvironmentBinding {
	if in == nil {
		return nil
	}
	out := new(SnapshotEnvironmentBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnapshotEnvironmentBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotEnvironmentBindingList) DeepCopyInto(out *SnapshotEnvironmentBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SnapshotEnvironmentBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotEnvironmentBindingList.
func (in *SnapshotEnvironmentBindingList) DeepCopy() *SnapshotEnvironmentBindingList {
	if in == nil {
		return nil
	}
	out := new(SnapshotEnvironmentBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnapshotEnvironmentBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotEnvironmentBindingSpec) DeepCopyInto(out *SnapshotEnvironmentBindingSpec) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]BindingComponent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotEnvironmentBindingSpec.
func (in *SnapshotEnvironmentBindingSpec) DeepCopy() *SnapshotEnvironmentBindingSpec {
	if in == nil {
		return nil
	}
	out := new(SnapshotEnvironmentBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotEnvironmentBindingStatus) DeepCopyInto(out *SnapshotEnvironmentBindingStatus) {
	*out = *in
	if in.GitOpsDeployments != nil {
		in, out := &in.GitOpsDeployments, &out.GitOpsDeployments
		*out = make([]BindingStatusGitOpsDeployment, len(*in))
		copy(*out, *in)
	}
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]BindingComponentStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GitOpsRepoConditions != nil {
		in, out := &in.GitOpsRepoConditions, &out.GitOpsRepoConditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotEnvironmentBindingStatus.
func (in *SnapshotEnvironmentBindingStatus) DeepCopy() *SnapshotEnvironmentBindingStatus {
	if in == nil {
		return nil
	}
	out := new(SnapshotEnvironmentBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotList) DeepCopyInto(out *SnapshotList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Snapshot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotList.
func (in *SnapshotList) DeepCopy() *SnapshotList {
	if in == nil {
		return nil
	}
	out := new(SnapshotList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnapshotList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotSpec) DeepCopyInto(out *SnapshotSpec) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]SnapshotComponent, len(*in))
		copy(*out, *in)
	}
	in.Artifacts.DeepCopyInto(&out.Artifacts)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotSpec.
func (in *SnapshotSpec) DeepCopy() *SnapshotSpec {
	if in == nil {
		return nil
	}
	out := new(SnapshotSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotStatus) DeepCopyInto(out *SnapshotStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotStatus.
func (in *SnapshotStatus) DeepCopy() *SnapshotStatus {
	if in == nil {
		return nil
	}
	out := new(SnapshotStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnstableEnvironmentConfiguration) DeepCopyInto(out *UnstableEnvironmentConfiguration) {
	*out = *in
	out.KubernetesClusterCredentials = in.KubernetesClusterCredentials
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnstableEnvironmentConfiguration.
func (in *UnstableEnvironmentConfiguration) DeepCopy() *UnstableEnvironmentConfiguration {
	if in == nil {
		return nil
	}
	out := new(UnstableEnvironmentConfiguration)
	in.DeepCopyInto(out)
	return out
}
