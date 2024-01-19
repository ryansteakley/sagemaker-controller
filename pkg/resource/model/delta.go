// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package model

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}
	customSetDefaults(a, b)

	if len(a.ko.Spec.Containers) != len(b.ko.Spec.Containers) {
		delta.Add("Spec.Containers", a.ko.Spec.Containers, b.ko.Spec.Containers)
	} else if len(a.ko.Spec.Containers) > 0 {
		if !reflect.DeepEqual(a.ko.Spec.Containers, b.ko.Spec.Containers) {
			delta.Add("Spec.Containers", a.ko.Spec.Containers, b.ko.Spec.Containers)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.EnableNetworkIsolation, b.ko.Spec.EnableNetworkIsolation) {
		delta.Add("Spec.EnableNetworkIsolation", a.ko.Spec.EnableNetworkIsolation, b.ko.Spec.EnableNetworkIsolation)
	} else if a.ko.Spec.EnableNetworkIsolation != nil && b.ko.Spec.EnableNetworkIsolation != nil {
		if *a.ko.Spec.EnableNetworkIsolation != *b.ko.Spec.EnableNetworkIsolation {
			delta.Add("Spec.EnableNetworkIsolation", a.ko.Spec.EnableNetworkIsolation, b.ko.Spec.EnableNetworkIsolation)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ExecutionRoleARN, b.ko.Spec.ExecutionRoleARN) {
		delta.Add("Spec.ExecutionRoleARN", a.ko.Spec.ExecutionRoleARN, b.ko.Spec.ExecutionRoleARN)
	} else if a.ko.Spec.ExecutionRoleARN != nil && b.ko.Spec.ExecutionRoleARN != nil {
		if *a.ko.Spec.ExecutionRoleARN != *b.ko.Spec.ExecutionRoleARN {
			delta.Add("Spec.ExecutionRoleARN", a.ko.Spec.ExecutionRoleARN, b.ko.Spec.ExecutionRoleARN)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.InferenceExecutionConfig, b.ko.Spec.InferenceExecutionConfig) {
		delta.Add("Spec.InferenceExecutionConfig", a.ko.Spec.InferenceExecutionConfig, b.ko.Spec.InferenceExecutionConfig)
	} else if a.ko.Spec.InferenceExecutionConfig != nil && b.ko.Spec.InferenceExecutionConfig != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.InferenceExecutionConfig.Mode, b.ko.Spec.InferenceExecutionConfig.Mode) {
			delta.Add("Spec.InferenceExecutionConfig.Mode", a.ko.Spec.InferenceExecutionConfig.Mode, b.ko.Spec.InferenceExecutionConfig.Mode)
		} else if a.ko.Spec.InferenceExecutionConfig.Mode != nil && b.ko.Spec.InferenceExecutionConfig.Mode != nil {
			if *a.ko.Spec.InferenceExecutionConfig.Mode != *b.ko.Spec.InferenceExecutionConfig.Mode {
				delta.Add("Spec.InferenceExecutionConfig.Mode", a.ko.Spec.InferenceExecutionConfig.Mode, b.ko.Spec.InferenceExecutionConfig.Mode)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ModelName, b.ko.Spec.ModelName) {
		delta.Add("Spec.ModelName", a.ko.Spec.ModelName, b.ko.Spec.ModelName)
	} else if a.ko.Spec.ModelName != nil && b.ko.Spec.ModelName != nil {
		if *a.ko.Spec.ModelName != *b.ko.Spec.ModelName {
			delta.Add("Spec.ModelName", a.ko.Spec.ModelName, b.ko.Spec.ModelName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PrimaryContainer, b.ko.Spec.PrimaryContainer) {
		delta.Add("Spec.PrimaryContainer", a.ko.Spec.PrimaryContainer, b.ko.Spec.PrimaryContainer)
	} else if a.ko.Spec.PrimaryContainer != nil && b.ko.Spec.PrimaryContainer != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.PrimaryContainer.ContainerHostname, b.ko.Spec.PrimaryContainer.ContainerHostname) {
			delta.Add("Spec.PrimaryContainer.ContainerHostname", a.ko.Spec.PrimaryContainer.ContainerHostname, b.ko.Spec.PrimaryContainer.ContainerHostname)
		} else if a.ko.Spec.PrimaryContainer.ContainerHostname != nil && b.ko.Spec.PrimaryContainer.ContainerHostname != nil {
			if *a.ko.Spec.PrimaryContainer.ContainerHostname != *b.ko.Spec.PrimaryContainer.ContainerHostname {
				delta.Add("Spec.PrimaryContainer.ContainerHostname", a.ko.Spec.PrimaryContainer.ContainerHostname, b.ko.Spec.PrimaryContainer.ContainerHostname)
			}
		}
		if len(a.ko.Spec.PrimaryContainer.Environment) != len(b.ko.Spec.PrimaryContainer.Environment) {
			delta.Add("Spec.PrimaryContainer.Environment", a.ko.Spec.PrimaryContainer.Environment, b.ko.Spec.PrimaryContainer.Environment)
		} else if len(a.ko.Spec.PrimaryContainer.Environment) > 0 {
			if !ackcompare.MapStringStringPEqual(a.ko.Spec.PrimaryContainer.Environment, b.ko.Spec.PrimaryContainer.Environment) {
				delta.Add("Spec.PrimaryContainer.Environment", a.ko.Spec.PrimaryContainer.Environment, b.ko.Spec.PrimaryContainer.Environment)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.PrimaryContainer.Image, b.ko.Spec.PrimaryContainer.Image) {
			delta.Add("Spec.PrimaryContainer.Image", a.ko.Spec.PrimaryContainer.Image, b.ko.Spec.PrimaryContainer.Image)
		} else if a.ko.Spec.PrimaryContainer.Image != nil && b.ko.Spec.PrimaryContainer.Image != nil {
			if *a.ko.Spec.PrimaryContainer.Image != *b.ko.Spec.PrimaryContainer.Image {
				delta.Add("Spec.PrimaryContainer.Image", a.ko.Spec.PrimaryContainer.Image, b.ko.Spec.PrimaryContainer.Image)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.PrimaryContainer.ImageConfig, b.ko.Spec.PrimaryContainer.ImageConfig) {
			delta.Add("Spec.PrimaryContainer.ImageConfig", a.ko.Spec.PrimaryContainer.ImageConfig, b.ko.Spec.PrimaryContainer.ImageConfig)
		} else if a.ko.Spec.PrimaryContainer.ImageConfig != nil && b.ko.Spec.PrimaryContainer.ImageConfig != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAccessMode, b.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAccessMode) {
				delta.Add("Spec.PrimaryContainer.ImageConfig.RepositoryAccessMode", a.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAccessMode, b.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAccessMode)
			} else if a.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAccessMode != nil && b.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAccessMode != nil {
				if *a.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAccessMode != *b.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAccessMode {
					delta.Add("Spec.PrimaryContainer.ImageConfig.RepositoryAccessMode", a.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAccessMode, b.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAccessMode)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAuthConfig, b.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAuthConfig) {
				delta.Add("Spec.PrimaryContainer.ImageConfig.RepositoryAuthConfig", a.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAuthConfig, b.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAuthConfig)
			} else if a.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAuthConfig != nil && b.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAuthConfig != nil {
				if ackcompare.HasNilDifference(a.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAuthConfig.RepositoryCredentialsProviderARN, b.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAuthConfig.RepositoryCredentialsProviderARN) {
					delta.Add("Spec.PrimaryContainer.ImageConfig.RepositoryAuthConfig.RepositoryCredentialsProviderARN", a.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAuthConfig.RepositoryCredentialsProviderARN, b.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAuthConfig.RepositoryCredentialsProviderARN)
				} else if a.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAuthConfig.RepositoryCredentialsProviderARN != nil && b.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAuthConfig.RepositoryCredentialsProviderARN != nil {
					if *a.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAuthConfig.RepositoryCredentialsProviderARN != *b.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAuthConfig.RepositoryCredentialsProviderARN {
						delta.Add("Spec.PrimaryContainer.ImageConfig.RepositoryAuthConfig.RepositoryCredentialsProviderARN", a.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAuthConfig.RepositoryCredentialsProviderARN, b.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAuthConfig.RepositoryCredentialsProviderARN)
					}
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.PrimaryContainer.InferenceSpecificationName, b.ko.Spec.PrimaryContainer.InferenceSpecificationName) {
			delta.Add("Spec.PrimaryContainer.InferenceSpecificationName", a.ko.Spec.PrimaryContainer.InferenceSpecificationName, b.ko.Spec.PrimaryContainer.InferenceSpecificationName)
		} else if a.ko.Spec.PrimaryContainer.InferenceSpecificationName != nil && b.ko.Spec.PrimaryContainer.InferenceSpecificationName != nil {
			if *a.ko.Spec.PrimaryContainer.InferenceSpecificationName != *b.ko.Spec.PrimaryContainer.InferenceSpecificationName {
				delta.Add("Spec.PrimaryContainer.InferenceSpecificationName", a.ko.Spec.PrimaryContainer.InferenceSpecificationName, b.ko.Spec.PrimaryContainer.InferenceSpecificationName)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.PrimaryContainer.Mode, b.ko.Spec.PrimaryContainer.Mode) {
			delta.Add("Spec.PrimaryContainer.Mode", a.ko.Spec.PrimaryContainer.Mode, b.ko.Spec.PrimaryContainer.Mode)
		} else if a.ko.Spec.PrimaryContainer.Mode != nil && b.ko.Spec.PrimaryContainer.Mode != nil {
			if *a.ko.Spec.PrimaryContainer.Mode != *b.ko.Spec.PrimaryContainer.Mode {
				delta.Add("Spec.PrimaryContainer.Mode", a.ko.Spec.PrimaryContainer.Mode, b.ko.Spec.PrimaryContainer.Mode)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.PrimaryContainer.ModelDataURL, b.ko.Spec.PrimaryContainer.ModelDataURL) {
			delta.Add("Spec.PrimaryContainer.ModelDataURL", a.ko.Spec.PrimaryContainer.ModelDataURL, b.ko.Spec.PrimaryContainer.ModelDataURL)
		} else if a.ko.Spec.PrimaryContainer.ModelDataURL != nil && b.ko.Spec.PrimaryContainer.ModelDataURL != nil {
			if *a.ko.Spec.PrimaryContainer.ModelDataURL != *b.ko.Spec.PrimaryContainer.ModelDataURL {
				delta.Add("Spec.PrimaryContainer.ModelDataURL", a.ko.Spec.PrimaryContainer.ModelDataURL, b.ko.Spec.PrimaryContainer.ModelDataURL)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.PrimaryContainer.ModelPackageName, b.ko.Spec.PrimaryContainer.ModelPackageName) {
			delta.Add("Spec.PrimaryContainer.ModelPackageName", a.ko.Spec.PrimaryContainer.ModelPackageName, b.ko.Spec.PrimaryContainer.ModelPackageName)
		} else if a.ko.Spec.PrimaryContainer.ModelPackageName != nil && b.ko.Spec.PrimaryContainer.ModelPackageName != nil {
			if *a.ko.Spec.PrimaryContainer.ModelPackageName != *b.ko.Spec.PrimaryContainer.ModelPackageName {
				delta.Add("Spec.PrimaryContainer.ModelPackageName", a.ko.Spec.PrimaryContainer.ModelPackageName, b.ko.Spec.PrimaryContainer.ModelPackageName)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.PrimaryContainer.MultiModelConfig, b.ko.Spec.PrimaryContainer.MultiModelConfig) {
			delta.Add("Spec.PrimaryContainer.MultiModelConfig", a.ko.Spec.PrimaryContainer.MultiModelConfig, b.ko.Spec.PrimaryContainer.MultiModelConfig)
		} else if a.ko.Spec.PrimaryContainer.MultiModelConfig != nil && b.ko.Spec.PrimaryContainer.MultiModelConfig != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.PrimaryContainer.MultiModelConfig.ModelCacheSetting, b.ko.Spec.PrimaryContainer.MultiModelConfig.ModelCacheSetting) {
				delta.Add("Spec.PrimaryContainer.MultiModelConfig.ModelCacheSetting", a.ko.Spec.PrimaryContainer.MultiModelConfig.ModelCacheSetting, b.ko.Spec.PrimaryContainer.MultiModelConfig.ModelCacheSetting)
			} else if a.ko.Spec.PrimaryContainer.MultiModelConfig.ModelCacheSetting != nil && b.ko.Spec.PrimaryContainer.MultiModelConfig.ModelCacheSetting != nil {
				if *a.ko.Spec.PrimaryContainer.MultiModelConfig.ModelCacheSetting != *b.ko.Spec.PrimaryContainer.MultiModelConfig.ModelCacheSetting {
					delta.Add("Spec.PrimaryContainer.MultiModelConfig.ModelCacheSetting", a.ko.Spec.PrimaryContainer.MultiModelConfig.ModelCacheSetting, b.ko.Spec.PrimaryContainer.MultiModelConfig.ModelCacheSetting)
				}
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.VPCConfig, b.ko.Spec.VPCConfig) {
		delta.Add("Spec.VPCConfig", a.ko.Spec.VPCConfig, b.ko.Spec.VPCConfig)
	} else if a.ko.Spec.VPCConfig != nil && b.ko.Spec.VPCConfig != nil {
		if len(a.ko.Spec.VPCConfig.SecurityGroupIDs) != len(b.ko.Spec.VPCConfig.SecurityGroupIDs) {
			delta.Add("Spec.VPCConfig.SecurityGroupIDs", a.ko.Spec.VPCConfig.SecurityGroupIDs, b.ko.Spec.VPCConfig.SecurityGroupIDs)
		} else if len(a.ko.Spec.VPCConfig.SecurityGroupIDs) > 0 {
			if !ackcompare.SliceStringPEqual(a.ko.Spec.VPCConfig.SecurityGroupIDs, b.ko.Spec.VPCConfig.SecurityGroupIDs) {
				delta.Add("Spec.VPCConfig.SecurityGroupIDs", a.ko.Spec.VPCConfig.SecurityGroupIDs, b.ko.Spec.VPCConfig.SecurityGroupIDs)
			}
		}
		if len(a.ko.Spec.VPCConfig.Subnets) != len(b.ko.Spec.VPCConfig.Subnets) {
			delta.Add("Spec.VPCConfig.Subnets", a.ko.Spec.VPCConfig.Subnets, b.ko.Spec.VPCConfig.Subnets)
		} else if len(a.ko.Spec.VPCConfig.Subnets) > 0 {
			if !ackcompare.SliceStringPEqual(a.ko.Spec.VPCConfig.Subnets, b.ko.Spec.VPCConfig.Subnets) {
				delta.Add("Spec.VPCConfig.Subnets", a.ko.Spec.VPCConfig.Subnets, b.ko.Spec.VPCConfig.Subnets)
			}
		}
	}

	return delta
}
