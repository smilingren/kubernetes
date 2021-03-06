// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	componentconfig "k8s.io/kubernetes/pkg/apis/componentconfig"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration,
		Convert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration,
		Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration,
		Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration,
	)
}

func autoConvert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration(in *KubeSchedulerConfiguration, out *componentconfig.KubeSchedulerConfiguration, s conversion.Scope) error {
	out.Port = int32(in.Port)
	out.Address = in.Address
	out.AlgorithmProvider = in.AlgorithmProvider
	out.PolicyConfigFile = in.PolicyConfigFile
	if err := v1.Convert_Pointer_bool_To_bool(&in.EnableProfiling, &out.EnableProfiling, s); err != nil {
		return err
	}
	out.EnableContentionProfiling = in.EnableContentionProfiling
	out.ContentType = in.ContentType
	out.KubeAPIQPS = in.KubeAPIQPS
	out.KubeAPIBurst = int32(in.KubeAPIBurst)
	out.SchedulerName = in.SchedulerName
	out.HardPodAffinitySymmetricWeight = in.HardPodAffinitySymmetricWeight
	out.FailureDomains = in.FailureDomains
	if err := Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	out.LockObjectNamespace = in.LockObjectNamespace
	out.LockObjectName = in.LockObjectName
	out.PolicyConfigMapName = in.PolicyConfigMapName
	out.PolicyConfigMapNamespace = in.PolicyConfigMapNamespace
	out.UseLegacyPolicyConfig = in.UseLegacyPolicyConfig
	return nil
}

// Convert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration(in *KubeSchedulerConfiguration, out *componentconfig.KubeSchedulerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration(in, out, s)
}

func autoConvert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(in *componentconfig.KubeSchedulerConfiguration, out *KubeSchedulerConfiguration, s conversion.Scope) error {
	out.Port = int(in.Port)
	out.Address = in.Address
	out.AlgorithmProvider = in.AlgorithmProvider
	out.PolicyConfigFile = in.PolicyConfigFile
	if err := v1.Convert_bool_To_Pointer_bool(&in.EnableProfiling, &out.EnableProfiling, s); err != nil {
		return err
	}
	out.EnableContentionProfiling = in.EnableContentionProfiling
	out.ContentType = in.ContentType
	out.KubeAPIQPS = in.KubeAPIQPS
	out.KubeAPIBurst = int(in.KubeAPIBurst)
	out.SchedulerName = in.SchedulerName
	out.HardPodAffinitySymmetricWeight = in.HardPodAffinitySymmetricWeight
	out.FailureDomains = in.FailureDomains
	if err := Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	out.LockObjectNamespace = in.LockObjectNamespace
	out.LockObjectName = in.LockObjectName
	out.PolicyConfigMapName = in.PolicyConfigMapName
	out.PolicyConfigMapNamespace = in.PolicyConfigMapNamespace
	out.UseLegacyPolicyConfig = in.UseLegacyPolicyConfig
	return nil
}

// Convert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration is an autogenerated conversion function.
func Convert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(in *componentconfig.KubeSchedulerConfiguration, out *KubeSchedulerConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(in *LeaderElectionConfiguration, out *componentconfig.LeaderElectionConfiguration, s conversion.Scope) error {
	if err := v1.Convert_Pointer_bool_To_bool(&in.LeaderElect, &out.LeaderElect, s); err != nil {
		return err
	}
	out.LeaseDuration = in.LeaseDuration
	out.RenewDeadline = in.RenewDeadline
	out.RetryPeriod = in.RetryPeriod
	out.ResourceLock = in.ResourceLock
	return nil
}

// Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(in *LeaderElectionConfiguration, out *componentconfig.LeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(in, out, s)
}

func autoConvert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in *componentconfig.LeaderElectionConfiguration, out *LeaderElectionConfiguration, s conversion.Scope) error {
	if err := v1.Convert_bool_To_Pointer_bool(&in.LeaderElect, &out.LeaderElect, s); err != nil {
		return err
	}
	out.LeaseDuration = in.LeaseDuration
	out.RenewDeadline = in.RenewDeadline
	out.RetryPeriod = in.RetryPeriod
	out.ResourceLock = in.ResourceLock
	return nil
}

// Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration is an autogenerated conversion function.
func Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in *componentconfig.LeaderElectionConfiguration, out *LeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in, out, s)
}
