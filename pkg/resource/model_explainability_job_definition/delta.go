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

package model_explainability_job_definition

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

	if ackcompare.HasNilDifference(a.ko.Spec.JobDefinitionName, b.ko.Spec.JobDefinitionName) {
		delta.Add("Spec.JobDefinitionName", a.ko.Spec.JobDefinitionName, b.ko.Spec.JobDefinitionName)
	} else if a.ko.Spec.JobDefinitionName != nil && b.ko.Spec.JobDefinitionName != nil {
		if *a.ko.Spec.JobDefinitionName != *b.ko.Spec.JobDefinitionName {
			delta.Add("Spec.JobDefinitionName", a.ko.Spec.JobDefinitionName, b.ko.Spec.JobDefinitionName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.JobResources, b.ko.Spec.JobResources) {
		delta.Add("Spec.JobResources", a.ko.Spec.JobResources, b.ko.Spec.JobResources)
	} else if a.ko.Spec.JobResources != nil && b.ko.Spec.JobResources != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.JobResources.ClusterConfig, b.ko.Spec.JobResources.ClusterConfig) {
			delta.Add("Spec.JobResources.ClusterConfig", a.ko.Spec.JobResources.ClusterConfig, b.ko.Spec.JobResources.ClusterConfig)
		} else if a.ko.Spec.JobResources.ClusterConfig != nil && b.ko.Spec.JobResources.ClusterConfig != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.JobResources.ClusterConfig.InstanceCount, b.ko.Spec.JobResources.ClusterConfig.InstanceCount) {
				delta.Add("Spec.JobResources.ClusterConfig.InstanceCount", a.ko.Spec.JobResources.ClusterConfig.InstanceCount, b.ko.Spec.JobResources.ClusterConfig.InstanceCount)
			} else if a.ko.Spec.JobResources.ClusterConfig.InstanceCount != nil && b.ko.Spec.JobResources.ClusterConfig.InstanceCount != nil {
				if *a.ko.Spec.JobResources.ClusterConfig.InstanceCount != *b.ko.Spec.JobResources.ClusterConfig.InstanceCount {
					delta.Add("Spec.JobResources.ClusterConfig.InstanceCount", a.ko.Spec.JobResources.ClusterConfig.InstanceCount, b.ko.Spec.JobResources.ClusterConfig.InstanceCount)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.JobResources.ClusterConfig.InstanceType, b.ko.Spec.JobResources.ClusterConfig.InstanceType) {
				delta.Add("Spec.JobResources.ClusterConfig.InstanceType", a.ko.Spec.JobResources.ClusterConfig.InstanceType, b.ko.Spec.JobResources.ClusterConfig.InstanceType)
			} else if a.ko.Spec.JobResources.ClusterConfig.InstanceType != nil && b.ko.Spec.JobResources.ClusterConfig.InstanceType != nil {
				if *a.ko.Spec.JobResources.ClusterConfig.InstanceType != *b.ko.Spec.JobResources.ClusterConfig.InstanceType {
					delta.Add("Spec.JobResources.ClusterConfig.InstanceType", a.ko.Spec.JobResources.ClusterConfig.InstanceType, b.ko.Spec.JobResources.ClusterConfig.InstanceType)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.JobResources.ClusterConfig.VolumeKMSKeyID, b.ko.Spec.JobResources.ClusterConfig.VolumeKMSKeyID) {
				delta.Add("Spec.JobResources.ClusterConfig.VolumeKMSKeyID", a.ko.Spec.JobResources.ClusterConfig.VolumeKMSKeyID, b.ko.Spec.JobResources.ClusterConfig.VolumeKMSKeyID)
			} else if a.ko.Spec.JobResources.ClusterConfig.VolumeKMSKeyID != nil && b.ko.Spec.JobResources.ClusterConfig.VolumeKMSKeyID != nil {
				if *a.ko.Spec.JobResources.ClusterConfig.VolumeKMSKeyID != *b.ko.Spec.JobResources.ClusterConfig.VolumeKMSKeyID {
					delta.Add("Spec.JobResources.ClusterConfig.VolumeKMSKeyID", a.ko.Spec.JobResources.ClusterConfig.VolumeKMSKeyID, b.ko.Spec.JobResources.ClusterConfig.VolumeKMSKeyID)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.JobResources.ClusterConfig.VolumeSizeInGB, b.ko.Spec.JobResources.ClusterConfig.VolumeSizeInGB) {
				delta.Add("Spec.JobResources.ClusterConfig.VolumeSizeInGB", a.ko.Spec.JobResources.ClusterConfig.VolumeSizeInGB, b.ko.Spec.JobResources.ClusterConfig.VolumeSizeInGB)
			} else if a.ko.Spec.JobResources.ClusterConfig.VolumeSizeInGB != nil && b.ko.Spec.JobResources.ClusterConfig.VolumeSizeInGB != nil {
				if *a.ko.Spec.JobResources.ClusterConfig.VolumeSizeInGB != *b.ko.Spec.JobResources.ClusterConfig.VolumeSizeInGB {
					delta.Add("Spec.JobResources.ClusterConfig.VolumeSizeInGB", a.ko.Spec.JobResources.ClusterConfig.VolumeSizeInGB, b.ko.Spec.JobResources.ClusterConfig.VolumeSizeInGB)
				}
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ModelExplainabilityAppSpecification, b.ko.Spec.ModelExplainabilityAppSpecification) {
		delta.Add("Spec.ModelExplainabilityAppSpecification", a.ko.Spec.ModelExplainabilityAppSpecification, b.ko.Spec.ModelExplainabilityAppSpecification)
	} else if a.ko.Spec.ModelExplainabilityAppSpecification != nil && b.ko.Spec.ModelExplainabilityAppSpecification != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.ModelExplainabilityAppSpecification.ConfigURI, b.ko.Spec.ModelExplainabilityAppSpecification.ConfigURI) {
			delta.Add("Spec.ModelExplainabilityAppSpecification.ConfigURI", a.ko.Spec.ModelExplainabilityAppSpecification.ConfigURI, b.ko.Spec.ModelExplainabilityAppSpecification.ConfigURI)
		} else if a.ko.Spec.ModelExplainabilityAppSpecification.ConfigURI != nil && b.ko.Spec.ModelExplainabilityAppSpecification.ConfigURI != nil {
			if *a.ko.Spec.ModelExplainabilityAppSpecification.ConfigURI != *b.ko.Spec.ModelExplainabilityAppSpecification.ConfigURI {
				delta.Add("Spec.ModelExplainabilityAppSpecification.ConfigURI", a.ko.Spec.ModelExplainabilityAppSpecification.ConfigURI, b.ko.Spec.ModelExplainabilityAppSpecification.ConfigURI)
			}
		}
		if len(a.ko.Spec.ModelExplainabilityAppSpecification.Environment) != len(b.ko.Spec.ModelExplainabilityAppSpecification.Environment) {
			delta.Add("Spec.ModelExplainabilityAppSpecification.Environment", a.ko.Spec.ModelExplainabilityAppSpecification.Environment, b.ko.Spec.ModelExplainabilityAppSpecification.Environment)
		} else if len(a.ko.Spec.ModelExplainabilityAppSpecification.Environment) > 0 {
			if !ackcompare.MapStringStringPEqual(a.ko.Spec.ModelExplainabilityAppSpecification.Environment, b.ko.Spec.ModelExplainabilityAppSpecification.Environment) {
				delta.Add("Spec.ModelExplainabilityAppSpecification.Environment", a.ko.Spec.ModelExplainabilityAppSpecification.Environment, b.ko.Spec.ModelExplainabilityAppSpecification.Environment)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ModelExplainabilityAppSpecification.ImageURI, b.ko.Spec.ModelExplainabilityAppSpecification.ImageURI) {
			delta.Add("Spec.ModelExplainabilityAppSpecification.ImageURI", a.ko.Spec.ModelExplainabilityAppSpecification.ImageURI, b.ko.Spec.ModelExplainabilityAppSpecification.ImageURI)
		} else if a.ko.Spec.ModelExplainabilityAppSpecification.ImageURI != nil && b.ko.Spec.ModelExplainabilityAppSpecification.ImageURI != nil {
			if *a.ko.Spec.ModelExplainabilityAppSpecification.ImageURI != *b.ko.Spec.ModelExplainabilityAppSpecification.ImageURI {
				delta.Add("Spec.ModelExplainabilityAppSpecification.ImageURI", a.ko.Spec.ModelExplainabilityAppSpecification.ImageURI, b.ko.Spec.ModelExplainabilityAppSpecification.ImageURI)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ModelExplainabilityBaselineConfig, b.ko.Spec.ModelExplainabilityBaselineConfig) {
		delta.Add("Spec.ModelExplainabilityBaselineConfig", a.ko.Spec.ModelExplainabilityBaselineConfig, b.ko.Spec.ModelExplainabilityBaselineConfig)
	} else if a.ko.Spec.ModelExplainabilityBaselineConfig != nil && b.ko.Spec.ModelExplainabilityBaselineConfig != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.ModelExplainabilityBaselineConfig.BaseliningJobName, b.ko.Spec.ModelExplainabilityBaselineConfig.BaseliningJobName) {
			delta.Add("Spec.ModelExplainabilityBaselineConfig.BaseliningJobName", a.ko.Spec.ModelExplainabilityBaselineConfig.BaseliningJobName, b.ko.Spec.ModelExplainabilityBaselineConfig.BaseliningJobName)
		} else if a.ko.Spec.ModelExplainabilityBaselineConfig.BaseliningJobName != nil && b.ko.Spec.ModelExplainabilityBaselineConfig.BaseliningJobName != nil {
			if *a.ko.Spec.ModelExplainabilityBaselineConfig.BaseliningJobName != *b.ko.Spec.ModelExplainabilityBaselineConfig.BaseliningJobName {
				delta.Add("Spec.ModelExplainabilityBaselineConfig.BaseliningJobName", a.ko.Spec.ModelExplainabilityBaselineConfig.BaseliningJobName, b.ko.Spec.ModelExplainabilityBaselineConfig.BaseliningJobName)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ModelExplainabilityBaselineConfig.ConstraintsResource, b.ko.Spec.ModelExplainabilityBaselineConfig.ConstraintsResource) {
			delta.Add("Spec.ModelExplainabilityBaselineConfig.ConstraintsResource", a.ko.Spec.ModelExplainabilityBaselineConfig.ConstraintsResource, b.ko.Spec.ModelExplainabilityBaselineConfig.ConstraintsResource)
		} else if a.ko.Spec.ModelExplainabilityBaselineConfig.ConstraintsResource != nil && b.ko.Spec.ModelExplainabilityBaselineConfig.ConstraintsResource != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.ModelExplainabilityBaselineConfig.ConstraintsResource.S3URI, b.ko.Spec.ModelExplainabilityBaselineConfig.ConstraintsResource.S3URI) {
				delta.Add("Spec.ModelExplainabilityBaselineConfig.ConstraintsResource.S3URI", a.ko.Spec.ModelExplainabilityBaselineConfig.ConstraintsResource.S3URI, b.ko.Spec.ModelExplainabilityBaselineConfig.ConstraintsResource.S3URI)
			} else if a.ko.Spec.ModelExplainabilityBaselineConfig.ConstraintsResource.S3URI != nil && b.ko.Spec.ModelExplainabilityBaselineConfig.ConstraintsResource.S3URI != nil {
				if *a.ko.Spec.ModelExplainabilityBaselineConfig.ConstraintsResource.S3URI != *b.ko.Spec.ModelExplainabilityBaselineConfig.ConstraintsResource.S3URI {
					delta.Add("Spec.ModelExplainabilityBaselineConfig.ConstraintsResource.S3URI", a.ko.Spec.ModelExplainabilityBaselineConfig.ConstraintsResource.S3URI, b.ko.Spec.ModelExplainabilityBaselineConfig.ConstraintsResource.S3URI)
				}
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ModelExplainabilityJobInput, b.ko.Spec.ModelExplainabilityJobInput) {
		delta.Add("Spec.ModelExplainabilityJobInput", a.ko.Spec.ModelExplainabilityJobInput, b.ko.Spec.ModelExplainabilityJobInput)
	} else if a.ko.Spec.ModelExplainabilityJobInput != nil && b.ko.Spec.ModelExplainabilityJobInput != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.ModelExplainabilityJobInput.EndpointInput, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput) {
			delta.Add("Spec.ModelExplainabilityJobInput.EndpointInput", a.ko.Spec.ModelExplainabilityJobInput.EndpointInput, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput)
		} else if a.ko.Spec.ModelExplainabilityJobInput.EndpointInput != nil && b.ko.Spec.ModelExplainabilityJobInput.EndpointInput != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.EndTimeOffset, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.EndTimeOffset) {
				delta.Add("Spec.ModelExplainabilityJobInput.EndpointInput.EndTimeOffset", a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.EndTimeOffset, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.EndTimeOffset)
			} else if a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.EndTimeOffset != nil && b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.EndTimeOffset != nil {
				if *a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.EndTimeOffset != *b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.EndTimeOffset {
					delta.Add("Spec.ModelExplainabilityJobInput.EndpointInput.EndTimeOffset", a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.EndTimeOffset, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.EndTimeOffset)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.EndpointName, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.EndpointName) {
				delta.Add("Spec.ModelExplainabilityJobInput.EndpointInput.EndpointName", a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.EndpointName, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.EndpointName)
			} else if a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.EndpointName != nil && b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.EndpointName != nil {
				if *a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.EndpointName != *b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.EndpointName {
					delta.Add("Spec.ModelExplainabilityJobInput.EndpointInput.EndpointName", a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.EndpointName, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.EndpointName)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.FeaturesAttribute, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.FeaturesAttribute) {
				delta.Add("Spec.ModelExplainabilityJobInput.EndpointInput.FeaturesAttribute", a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.FeaturesAttribute, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.FeaturesAttribute)
			} else if a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.FeaturesAttribute != nil && b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.FeaturesAttribute != nil {
				if *a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.FeaturesAttribute != *b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.FeaturesAttribute {
					delta.Add("Spec.ModelExplainabilityJobInput.EndpointInput.FeaturesAttribute", a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.FeaturesAttribute, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.FeaturesAttribute)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.InferenceAttribute, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.InferenceAttribute) {
				delta.Add("Spec.ModelExplainabilityJobInput.EndpointInput.InferenceAttribute", a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.InferenceAttribute, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.InferenceAttribute)
			} else if a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.InferenceAttribute != nil && b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.InferenceAttribute != nil {
				if *a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.InferenceAttribute != *b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.InferenceAttribute {
					delta.Add("Spec.ModelExplainabilityJobInput.EndpointInput.InferenceAttribute", a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.InferenceAttribute, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.InferenceAttribute)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.LocalPath, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.LocalPath) {
				delta.Add("Spec.ModelExplainabilityJobInput.EndpointInput.LocalPath", a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.LocalPath, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.LocalPath)
			} else if a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.LocalPath != nil && b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.LocalPath != nil {
				if *a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.LocalPath != *b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.LocalPath {
					delta.Add("Spec.ModelExplainabilityJobInput.EndpointInput.LocalPath", a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.LocalPath, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.LocalPath)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.ProbabilityAttribute, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.ProbabilityAttribute) {
				delta.Add("Spec.ModelExplainabilityJobInput.EndpointInput.ProbabilityAttribute", a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.ProbabilityAttribute, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.ProbabilityAttribute)
			} else if a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.ProbabilityAttribute != nil && b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.ProbabilityAttribute != nil {
				if *a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.ProbabilityAttribute != *b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.ProbabilityAttribute {
					delta.Add("Spec.ModelExplainabilityJobInput.EndpointInput.ProbabilityAttribute", a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.ProbabilityAttribute, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.ProbabilityAttribute)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.ProbabilityThresholdAttribute, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.ProbabilityThresholdAttribute) {
				delta.Add("Spec.ModelExplainabilityJobInput.EndpointInput.ProbabilityThresholdAttribute", a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.ProbabilityThresholdAttribute, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.ProbabilityThresholdAttribute)
			} else if a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.ProbabilityThresholdAttribute != nil && b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.ProbabilityThresholdAttribute != nil {
				if *a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.ProbabilityThresholdAttribute != *b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.ProbabilityThresholdAttribute {
					delta.Add("Spec.ModelExplainabilityJobInput.EndpointInput.ProbabilityThresholdAttribute", a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.ProbabilityThresholdAttribute, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.ProbabilityThresholdAttribute)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.S3DataDistributionType, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.S3DataDistributionType) {
				delta.Add("Spec.ModelExplainabilityJobInput.EndpointInput.S3DataDistributionType", a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.S3DataDistributionType, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.S3DataDistributionType)
			} else if a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.S3DataDistributionType != nil && b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.S3DataDistributionType != nil {
				if *a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.S3DataDistributionType != *b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.S3DataDistributionType {
					delta.Add("Spec.ModelExplainabilityJobInput.EndpointInput.S3DataDistributionType", a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.S3DataDistributionType, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.S3DataDistributionType)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.S3InputMode, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.S3InputMode) {
				delta.Add("Spec.ModelExplainabilityJobInput.EndpointInput.S3InputMode", a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.S3InputMode, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.S3InputMode)
			} else if a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.S3InputMode != nil && b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.S3InputMode != nil {
				if *a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.S3InputMode != *b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.S3InputMode {
					delta.Add("Spec.ModelExplainabilityJobInput.EndpointInput.S3InputMode", a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.S3InputMode, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.S3InputMode)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.StartTimeOffset, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.StartTimeOffset) {
				delta.Add("Spec.ModelExplainabilityJobInput.EndpointInput.StartTimeOffset", a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.StartTimeOffset, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.StartTimeOffset)
			} else if a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.StartTimeOffset != nil && b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.StartTimeOffset != nil {
				if *a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.StartTimeOffset != *b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.StartTimeOffset {
					delta.Add("Spec.ModelExplainabilityJobInput.EndpointInput.StartTimeOffset", a.ko.Spec.ModelExplainabilityJobInput.EndpointInput.StartTimeOffset, b.ko.Spec.ModelExplainabilityJobInput.EndpointInput.StartTimeOffset)
				}
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ModelExplainabilityJobOutputConfig, b.ko.Spec.ModelExplainabilityJobOutputConfig) {
		delta.Add("Spec.ModelExplainabilityJobOutputConfig", a.ko.Spec.ModelExplainabilityJobOutputConfig, b.ko.Spec.ModelExplainabilityJobOutputConfig)
	} else if a.ko.Spec.ModelExplainabilityJobOutputConfig != nil && b.ko.Spec.ModelExplainabilityJobOutputConfig != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.ModelExplainabilityJobOutputConfig.KMSKeyID, b.ko.Spec.ModelExplainabilityJobOutputConfig.KMSKeyID) {
			delta.Add("Spec.ModelExplainabilityJobOutputConfig.KMSKeyID", a.ko.Spec.ModelExplainabilityJobOutputConfig.KMSKeyID, b.ko.Spec.ModelExplainabilityJobOutputConfig.KMSKeyID)
		} else if a.ko.Spec.ModelExplainabilityJobOutputConfig.KMSKeyID != nil && b.ko.Spec.ModelExplainabilityJobOutputConfig.KMSKeyID != nil {
			if *a.ko.Spec.ModelExplainabilityJobOutputConfig.KMSKeyID != *b.ko.Spec.ModelExplainabilityJobOutputConfig.KMSKeyID {
				delta.Add("Spec.ModelExplainabilityJobOutputConfig.KMSKeyID", a.ko.Spec.ModelExplainabilityJobOutputConfig.KMSKeyID, b.ko.Spec.ModelExplainabilityJobOutputConfig.KMSKeyID)
			}
		}
		if len(a.ko.Spec.ModelExplainabilityJobOutputConfig.MonitoringOutputs) != len(b.ko.Spec.ModelExplainabilityJobOutputConfig.MonitoringOutputs) {
			delta.Add("Spec.ModelExplainabilityJobOutputConfig.MonitoringOutputs", a.ko.Spec.ModelExplainabilityJobOutputConfig.MonitoringOutputs, b.ko.Spec.ModelExplainabilityJobOutputConfig.MonitoringOutputs)
		} else if len(a.ko.Spec.ModelExplainabilityJobOutputConfig.MonitoringOutputs) > 0 {
			if !reflect.DeepEqual(a.ko.Spec.ModelExplainabilityJobOutputConfig.MonitoringOutputs, b.ko.Spec.ModelExplainabilityJobOutputConfig.MonitoringOutputs) {
				delta.Add("Spec.ModelExplainabilityJobOutputConfig.MonitoringOutputs", a.ko.Spec.ModelExplainabilityJobOutputConfig.MonitoringOutputs, b.ko.Spec.ModelExplainabilityJobOutputConfig.MonitoringOutputs)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.NetworkConfig, b.ko.Spec.NetworkConfig) {
		delta.Add("Spec.NetworkConfig", a.ko.Spec.NetworkConfig, b.ko.Spec.NetworkConfig)
	} else if a.ko.Spec.NetworkConfig != nil && b.ko.Spec.NetworkConfig != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.NetworkConfig.EnableInterContainerTrafficEncryption, b.ko.Spec.NetworkConfig.EnableInterContainerTrafficEncryption) {
			delta.Add("Spec.NetworkConfig.EnableInterContainerTrafficEncryption", a.ko.Spec.NetworkConfig.EnableInterContainerTrafficEncryption, b.ko.Spec.NetworkConfig.EnableInterContainerTrafficEncryption)
		} else if a.ko.Spec.NetworkConfig.EnableInterContainerTrafficEncryption != nil && b.ko.Spec.NetworkConfig.EnableInterContainerTrafficEncryption != nil {
			if *a.ko.Spec.NetworkConfig.EnableInterContainerTrafficEncryption != *b.ko.Spec.NetworkConfig.EnableInterContainerTrafficEncryption {
				delta.Add("Spec.NetworkConfig.EnableInterContainerTrafficEncryption", a.ko.Spec.NetworkConfig.EnableInterContainerTrafficEncryption, b.ko.Spec.NetworkConfig.EnableInterContainerTrafficEncryption)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.NetworkConfig.EnableNetworkIsolation, b.ko.Spec.NetworkConfig.EnableNetworkIsolation) {
			delta.Add("Spec.NetworkConfig.EnableNetworkIsolation", a.ko.Spec.NetworkConfig.EnableNetworkIsolation, b.ko.Spec.NetworkConfig.EnableNetworkIsolation)
		} else if a.ko.Spec.NetworkConfig.EnableNetworkIsolation != nil && b.ko.Spec.NetworkConfig.EnableNetworkIsolation != nil {
			if *a.ko.Spec.NetworkConfig.EnableNetworkIsolation != *b.ko.Spec.NetworkConfig.EnableNetworkIsolation {
				delta.Add("Spec.NetworkConfig.EnableNetworkIsolation", a.ko.Spec.NetworkConfig.EnableNetworkIsolation, b.ko.Spec.NetworkConfig.EnableNetworkIsolation)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.NetworkConfig.VPCConfig, b.ko.Spec.NetworkConfig.VPCConfig) {
			delta.Add("Spec.NetworkConfig.VPCConfig", a.ko.Spec.NetworkConfig.VPCConfig, b.ko.Spec.NetworkConfig.VPCConfig)
		} else if a.ko.Spec.NetworkConfig.VPCConfig != nil && b.ko.Spec.NetworkConfig.VPCConfig != nil {
			if len(a.ko.Spec.NetworkConfig.VPCConfig.SecurityGroupIDs) != len(b.ko.Spec.NetworkConfig.VPCConfig.SecurityGroupIDs) {
				delta.Add("Spec.NetworkConfig.VPCConfig.SecurityGroupIDs", a.ko.Spec.NetworkConfig.VPCConfig.SecurityGroupIDs, b.ko.Spec.NetworkConfig.VPCConfig.SecurityGroupIDs)
			} else if len(a.ko.Spec.NetworkConfig.VPCConfig.SecurityGroupIDs) > 0 {
				if !ackcompare.SliceStringPEqual(a.ko.Spec.NetworkConfig.VPCConfig.SecurityGroupIDs, b.ko.Spec.NetworkConfig.VPCConfig.SecurityGroupIDs) {
					delta.Add("Spec.NetworkConfig.VPCConfig.SecurityGroupIDs", a.ko.Spec.NetworkConfig.VPCConfig.SecurityGroupIDs, b.ko.Spec.NetworkConfig.VPCConfig.SecurityGroupIDs)
				}
			}
			if len(a.ko.Spec.NetworkConfig.VPCConfig.Subnets) != len(b.ko.Spec.NetworkConfig.VPCConfig.Subnets) {
				delta.Add("Spec.NetworkConfig.VPCConfig.Subnets", a.ko.Spec.NetworkConfig.VPCConfig.Subnets, b.ko.Spec.NetworkConfig.VPCConfig.Subnets)
			} else if len(a.ko.Spec.NetworkConfig.VPCConfig.Subnets) > 0 {
				if !ackcompare.SliceStringPEqual(a.ko.Spec.NetworkConfig.VPCConfig.Subnets, b.ko.Spec.NetworkConfig.VPCConfig.Subnets) {
					delta.Add("Spec.NetworkConfig.VPCConfig.Subnets", a.ko.Spec.NetworkConfig.VPCConfig.Subnets, b.ko.Spec.NetworkConfig.VPCConfig.Subnets)
				}
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RoleARN, b.ko.Spec.RoleARN) {
		delta.Add("Spec.RoleARN", a.ko.Spec.RoleARN, b.ko.Spec.RoleARN)
	} else if a.ko.Spec.RoleARN != nil && b.ko.Spec.RoleARN != nil {
		if *a.ko.Spec.RoleARN != *b.ko.Spec.RoleARN {
			delta.Add("Spec.RoleARN", a.ko.Spec.RoleARN, b.ko.Spec.RoleARN)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.StoppingCondition, b.ko.Spec.StoppingCondition) {
		delta.Add("Spec.StoppingCondition", a.ko.Spec.StoppingCondition, b.ko.Spec.StoppingCondition)
	} else if a.ko.Spec.StoppingCondition != nil && b.ko.Spec.StoppingCondition != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.StoppingCondition.MaxRuntimeInSeconds, b.ko.Spec.StoppingCondition.MaxRuntimeInSeconds) {
			delta.Add("Spec.StoppingCondition.MaxRuntimeInSeconds", a.ko.Spec.StoppingCondition.MaxRuntimeInSeconds, b.ko.Spec.StoppingCondition.MaxRuntimeInSeconds)
		} else if a.ko.Spec.StoppingCondition.MaxRuntimeInSeconds != nil && b.ko.Spec.StoppingCondition.MaxRuntimeInSeconds != nil {
			if *a.ko.Spec.StoppingCondition.MaxRuntimeInSeconds != *b.ko.Spec.StoppingCondition.MaxRuntimeInSeconds {
				delta.Add("Spec.StoppingCondition.MaxRuntimeInSeconds", a.ko.Spec.StoppingCondition.MaxRuntimeInSeconds, b.ko.Spec.StoppingCondition.MaxRuntimeInSeconds)
			}
		}
	}

	return delta
}
