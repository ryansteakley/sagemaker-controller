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

package hyper_parameter_tuning_job

import (
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
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

	if ackcompare.HasNilDifference(a.ko.Spec.HyperParameterTuningJobConfig, b.ko.Spec.HyperParameterTuningJobConfig) {
		delta.Add("Spec.HyperParameterTuningJobConfig", a.ko.Spec.HyperParameterTuningJobConfig, b.ko.Spec.HyperParameterTuningJobConfig)
	} else if a.ko.Spec.HyperParameterTuningJobConfig != nil && b.ko.Spec.HyperParameterTuningJobConfig != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective, b.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective) {
			delta.Add("Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective", a.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective, b.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective)
		} else if a.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective != nil && b.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.MetricName, b.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.MetricName) {
				delta.Add("Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.MetricName", a.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.MetricName, b.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.MetricName)
			} else if a.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.MetricName != nil && b.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.MetricName != nil {
				if *a.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.MetricName != *b.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.MetricName {
					delta.Add("Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.MetricName", a.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.MetricName, b.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.MetricName)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.Type, b.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.Type) {
				delta.Add("Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.Type", a.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.Type, b.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.Type)
			} else if a.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.Type != nil && b.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.Type != nil {
				if *a.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.Type != *b.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.Type {
					delta.Add("Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.Type", a.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.Type, b.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.Type)
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.HyperParameterTuningJobConfig.ParameterRanges, b.ko.Spec.HyperParameterTuningJobConfig.ParameterRanges) {
			delta.Add("Spec.HyperParameterTuningJobConfig.ParameterRanges", a.ko.Spec.HyperParameterTuningJobConfig.ParameterRanges, b.ko.Spec.HyperParameterTuningJobConfig.ParameterRanges)
		} else if a.ko.Spec.HyperParameterTuningJobConfig.ParameterRanges != nil && b.ko.Spec.HyperParameterTuningJobConfig.ParameterRanges != nil {

		}
		if ackcompare.HasNilDifference(a.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits, b.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits) {
			delta.Add("Spec.HyperParameterTuningJobConfig.ResourceLimits", a.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits, b.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits)
		} else if a.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits != nil && b.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxNumberOfTrainingJobs, b.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxNumberOfTrainingJobs) {
				delta.Add("Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxNumberOfTrainingJobs", a.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxNumberOfTrainingJobs, b.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxNumberOfTrainingJobs)
			} else if a.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxNumberOfTrainingJobs != nil && b.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxNumberOfTrainingJobs != nil {
				if *a.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxNumberOfTrainingJobs != *b.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxNumberOfTrainingJobs {
					delta.Add("Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxNumberOfTrainingJobs", a.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxNumberOfTrainingJobs, b.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxNumberOfTrainingJobs)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxParallelTrainingJobs, b.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxParallelTrainingJobs) {
				delta.Add("Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxParallelTrainingJobs", a.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxParallelTrainingJobs, b.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxParallelTrainingJobs)
			} else if a.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxParallelTrainingJobs != nil && b.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxParallelTrainingJobs != nil {
				if *a.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxParallelTrainingJobs != *b.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxParallelTrainingJobs {
					delta.Add("Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxParallelTrainingJobs", a.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxParallelTrainingJobs, b.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxParallelTrainingJobs)
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.HyperParameterTuningJobConfig.Strategy, b.ko.Spec.HyperParameterTuningJobConfig.Strategy) {
			delta.Add("Spec.HyperParameterTuningJobConfig.Strategy", a.ko.Spec.HyperParameterTuningJobConfig.Strategy, b.ko.Spec.HyperParameterTuningJobConfig.Strategy)
		} else if a.ko.Spec.HyperParameterTuningJobConfig.Strategy != nil && b.ko.Spec.HyperParameterTuningJobConfig.Strategy != nil {
			if *a.ko.Spec.HyperParameterTuningJobConfig.Strategy != *b.ko.Spec.HyperParameterTuningJobConfig.Strategy {
				delta.Add("Spec.HyperParameterTuningJobConfig.Strategy", a.ko.Spec.HyperParameterTuningJobConfig.Strategy, b.ko.Spec.HyperParameterTuningJobConfig.Strategy)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.HyperParameterTuningJobConfig.TrainingJobEarlyStoppingType, b.ko.Spec.HyperParameterTuningJobConfig.TrainingJobEarlyStoppingType) {
			delta.Add("Spec.HyperParameterTuningJobConfig.TrainingJobEarlyStoppingType", a.ko.Spec.HyperParameterTuningJobConfig.TrainingJobEarlyStoppingType, b.ko.Spec.HyperParameterTuningJobConfig.TrainingJobEarlyStoppingType)
		} else if a.ko.Spec.HyperParameterTuningJobConfig.TrainingJobEarlyStoppingType != nil && b.ko.Spec.HyperParameterTuningJobConfig.TrainingJobEarlyStoppingType != nil {
			if *a.ko.Spec.HyperParameterTuningJobConfig.TrainingJobEarlyStoppingType != *b.ko.Spec.HyperParameterTuningJobConfig.TrainingJobEarlyStoppingType {
				delta.Add("Spec.HyperParameterTuningJobConfig.TrainingJobEarlyStoppingType", a.ko.Spec.HyperParameterTuningJobConfig.TrainingJobEarlyStoppingType, b.ko.Spec.HyperParameterTuningJobConfig.TrainingJobEarlyStoppingType)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.HyperParameterTuningJobConfig.TuningJobCompletionCriteria, b.ko.Spec.HyperParameterTuningJobConfig.TuningJobCompletionCriteria) {
			delta.Add("Spec.HyperParameterTuningJobConfig.TuningJobCompletionCriteria", a.ko.Spec.HyperParameterTuningJobConfig.TuningJobCompletionCriteria, b.ko.Spec.HyperParameterTuningJobConfig.TuningJobCompletionCriteria)
		} else if a.ko.Spec.HyperParameterTuningJobConfig.TuningJobCompletionCriteria != nil && b.ko.Spec.HyperParameterTuningJobConfig.TuningJobCompletionCriteria != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.HyperParameterTuningJobConfig.TuningJobCompletionCriteria.TargetObjectiveMetricValue, b.ko.Spec.HyperParameterTuningJobConfig.TuningJobCompletionCriteria.TargetObjectiveMetricValue) {
				delta.Add("Spec.HyperParameterTuningJobConfig.TuningJobCompletionCriteria.TargetObjectiveMetricValue", a.ko.Spec.HyperParameterTuningJobConfig.TuningJobCompletionCriteria.TargetObjectiveMetricValue, b.ko.Spec.HyperParameterTuningJobConfig.TuningJobCompletionCriteria.TargetObjectiveMetricValue)
			} else if a.ko.Spec.HyperParameterTuningJobConfig.TuningJobCompletionCriteria.TargetObjectiveMetricValue != nil && b.ko.Spec.HyperParameterTuningJobConfig.TuningJobCompletionCriteria.TargetObjectiveMetricValue != nil {
				if *a.ko.Spec.HyperParameterTuningJobConfig.TuningJobCompletionCriteria.TargetObjectiveMetricValue != *b.ko.Spec.HyperParameterTuningJobConfig.TuningJobCompletionCriteria.TargetObjectiveMetricValue {
					delta.Add("Spec.HyperParameterTuningJobConfig.TuningJobCompletionCriteria.TargetObjectiveMetricValue", a.ko.Spec.HyperParameterTuningJobConfig.TuningJobCompletionCriteria.TargetObjectiveMetricValue, b.ko.Spec.HyperParameterTuningJobConfig.TuningJobCompletionCriteria.TargetObjectiveMetricValue)
				}
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.HyperParameterTuningJobName, b.ko.Spec.HyperParameterTuningJobName) {
		delta.Add("Spec.HyperParameterTuningJobName", a.ko.Spec.HyperParameterTuningJobName, b.ko.Spec.HyperParameterTuningJobName)
	} else if a.ko.Spec.HyperParameterTuningJobName != nil && b.ko.Spec.HyperParameterTuningJobName != nil {
		if *a.ko.Spec.HyperParameterTuningJobName != *b.ko.Spec.HyperParameterTuningJobName {
			delta.Add("Spec.HyperParameterTuningJobName", a.ko.Spec.HyperParameterTuningJobName, b.ko.Spec.HyperParameterTuningJobName)
		}
	}

	if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition, b.ko.Spec.TrainingJobDefinition) {
		delta.Add("Spec.TrainingJobDefinition", a.ko.Spec.TrainingJobDefinition, b.ko.Spec.TrainingJobDefinition)
	} else if a.ko.Spec.TrainingJobDefinition != nil && b.ko.Spec.TrainingJobDefinition != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.AlgorithmSpecification, b.ko.Spec.TrainingJobDefinition.AlgorithmSpecification) {
			delta.Add("Spec.TrainingJobDefinition.AlgorithmSpecification", a.ko.Spec.TrainingJobDefinition.AlgorithmSpecification, b.ko.Spec.TrainingJobDefinition.AlgorithmSpecification)
		} else if a.ko.Spec.TrainingJobDefinition.AlgorithmSpecification != nil && b.ko.Spec.TrainingJobDefinition.AlgorithmSpecification != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.AlgorithmName, b.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.AlgorithmName) {
				delta.Add("Spec.TrainingJobDefinition.AlgorithmSpecification.AlgorithmName", a.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.AlgorithmName, b.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.AlgorithmName)
			} else if a.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.AlgorithmName != nil && b.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.AlgorithmName != nil {
				if *a.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.AlgorithmName != *b.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.AlgorithmName {
					delta.Add("Spec.TrainingJobDefinition.AlgorithmSpecification.AlgorithmName", a.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.AlgorithmName, b.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.AlgorithmName)
				}
			}

			if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingImage, b.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingImage) {
				delta.Add("Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingImage", a.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingImage, b.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingImage)
			} else if a.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingImage != nil && b.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingImage != nil {
				if *a.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingImage != *b.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingImage {
					delta.Add("Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingImage", a.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingImage, b.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingImage)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingInputMode, b.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingInputMode) {
				delta.Add("Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingInputMode", a.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingInputMode, b.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingInputMode)
			} else if a.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingInputMode != nil && b.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingInputMode != nil {
				if *a.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingInputMode != *b.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingInputMode {
					delta.Add("Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingInputMode", a.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingInputMode, b.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingInputMode)
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.CheckpointConfig, b.ko.Spec.TrainingJobDefinition.CheckpointConfig) {
			delta.Add("Spec.TrainingJobDefinition.CheckpointConfig", a.ko.Spec.TrainingJobDefinition.CheckpointConfig, b.ko.Spec.TrainingJobDefinition.CheckpointConfig)
		} else if a.ko.Spec.TrainingJobDefinition.CheckpointConfig != nil && b.ko.Spec.TrainingJobDefinition.CheckpointConfig != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.CheckpointConfig.LocalPath, b.ko.Spec.TrainingJobDefinition.CheckpointConfig.LocalPath) {
				delta.Add("Spec.TrainingJobDefinition.CheckpointConfig.LocalPath", a.ko.Spec.TrainingJobDefinition.CheckpointConfig.LocalPath, b.ko.Spec.TrainingJobDefinition.CheckpointConfig.LocalPath)
			} else if a.ko.Spec.TrainingJobDefinition.CheckpointConfig.LocalPath != nil && b.ko.Spec.TrainingJobDefinition.CheckpointConfig.LocalPath != nil {
				if *a.ko.Spec.TrainingJobDefinition.CheckpointConfig.LocalPath != *b.ko.Spec.TrainingJobDefinition.CheckpointConfig.LocalPath {
					delta.Add("Spec.TrainingJobDefinition.CheckpointConfig.LocalPath", a.ko.Spec.TrainingJobDefinition.CheckpointConfig.LocalPath, b.ko.Spec.TrainingJobDefinition.CheckpointConfig.LocalPath)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.CheckpointConfig.S3URI, b.ko.Spec.TrainingJobDefinition.CheckpointConfig.S3URI) {
				delta.Add("Spec.TrainingJobDefinition.CheckpointConfig.S3URI", a.ko.Spec.TrainingJobDefinition.CheckpointConfig.S3URI, b.ko.Spec.TrainingJobDefinition.CheckpointConfig.S3URI)
			} else if a.ko.Spec.TrainingJobDefinition.CheckpointConfig.S3URI != nil && b.ko.Spec.TrainingJobDefinition.CheckpointConfig.S3URI != nil {
				if *a.ko.Spec.TrainingJobDefinition.CheckpointConfig.S3URI != *b.ko.Spec.TrainingJobDefinition.CheckpointConfig.S3URI {
					delta.Add("Spec.TrainingJobDefinition.CheckpointConfig.S3URI", a.ko.Spec.TrainingJobDefinition.CheckpointConfig.S3URI, b.ko.Spec.TrainingJobDefinition.CheckpointConfig.S3URI)
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.DefinitionName, b.ko.Spec.TrainingJobDefinition.DefinitionName) {
			delta.Add("Spec.TrainingJobDefinition.DefinitionName", a.ko.Spec.TrainingJobDefinition.DefinitionName, b.ko.Spec.TrainingJobDefinition.DefinitionName)
		} else if a.ko.Spec.TrainingJobDefinition.DefinitionName != nil && b.ko.Spec.TrainingJobDefinition.DefinitionName != nil {
			if *a.ko.Spec.TrainingJobDefinition.DefinitionName != *b.ko.Spec.TrainingJobDefinition.DefinitionName {
				delta.Add("Spec.TrainingJobDefinition.DefinitionName", a.ko.Spec.TrainingJobDefinition.DefinitionName, b.ko.Spec.TrainingJobDefinition.DefinitionName)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.EnableInterContainerTrafficEncryption, b.ko.Spec.TrainingJobDefinition.EnableInterContainerTrafficEncryption) {
			delta.Add("Spec.TrainingJobDefinition.EnableInterContainerTrafficEncryption", a.ko.Spec.TrainingJobDefinition.EnableInterContainerTrafficEncryption, b.ko.Spec.TrainingJobDefinition.EnableInterContainerTrafficEncryption)
		} else if a.ko.Spec.TrainingJobDefinition.EnableInterContainerTrafficEncryption != nil && b.ko.Spec.TrainingJobDefinition.EnableInterContainerTrafficEncryption != nil {
			if *a.ko.Spec.TrainingJobDefinition.EnableInterContainerTrafficEncryption != *b.ko.Spec.TrainingJobDefinition.EnableInterContainerTrafficEncryption {
				delta.Add("Spec.TrainingJobDefinition.EnableInterContainerTrafficEncryption", a.ko.Spec.TrainingJobDefinition.EnableInterContainerTrafficEncryption, b.ko.Spec.TrainingJobDefinition.EnableInterContainerTrafficEncryption)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.EnableManagedSpotTraining, b.ko.Spec.TrainingJobDefinition.EnableManagedSpotTraining) {
			delta.Add("Spec.TrainingJobDefinition.EnableManagedSpotTraining", a.ko.Spec.TrainingJobDefinition.EnableManagedSpotTraining, b.ko.Spec.TrainingJobDefinition.EnableManagedSpotTraining)
		} else if a.ko.Spec.TrainingJobDefinition.EnableManagedSpotTraining != nil && b.ko.Spec.TrainingJobDefinition.EnableManagedSpotTraining != nil {
			if *a.ko.Spec.TrainingJobDefinition.EnableManagedSpotTraining != *b.ko.Spec.TrainingJobDefinition.EnableManagedSpotTraining {
				delta.Add("Spec.TrainingJobDefinition.EnableManagedSpotTraining", a.ko.Spec.TrainingJobDefinition.EnableManagedSpotTraining, b.ko.Spec.TrainingJobDefinition.EnableManagedSpotTraining)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.EnableNetworkIsolation, b.ko.Spec.TrainingJobDefinition.EnableNetworkIsolation) {
			delta.Add("Spec.TrainingJobDefinition.EnableNetworkIsolation", a.ko.Spec.TrainingJobDefinition.EnableNetworkIsolation, b.ko.Spec.TrainingJobDefinition.EnableNetworkIsolation)
		} else if a.ko.Spec.TrainingJobDefinition.EnableNetworkIsolation != nil && b.ko.Spec.TrainingJobDefinition.EnableNetworkIsolation != nil {
			if *a.ko.Spec.TrainingJobDefinition.EnableNetworkIsolation != *b.ko.Spec.TrainingJobDefinition.EnableNetworkIsolation {
				delta.Add("Spec.TrainingJobDefinition.EnableNetworkIsolation", a.ko.Spec.TrainingJobDefinition.EnableNetworkIsolation, b.ko.Spec.TrainingJobDefinition.EnableNetworkIsolation)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.HyperParameterRanges, b.ko.Spec.TrainingJobDefinition.HyperParameterRanges) {
			delta.Add("Spec.TrainingJobDefinition.HyperParameterRanges", a.ko.Spec.TrainingJobDefinition.HyperParameterRanges, b.ko.Spec.TrainingJobDefinition.HyperParameterRanges)
		} else if a.ko.Spec.TrainingJobDefinition.HyperParameterRanges != nil && b.ko.Spec.TrainingJobDefinition.HyperParameterRanges != nil {

		}

		if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.OutputDataConfig, b.ko.Spec.TrainingJobDefinition.OutputDataConfig) {
			delta.Add("Spec.TrainingJobDefinition.OutputDataConfig", a.ko.Spec.TrainingJobDefinition.OutputDataConfig, b.ko.Spec.TrainingJobDefinition.OutputDataConfig)
		} else if a.ko.Spec.TrainingJobDefinition.OutputDataConfig != nil && b.ko.Spec.TrainingJobDefinition.OutputDataConfig != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.OutputDataConfig.KMSKeyID, b.ko.Spec.TrainingJobDefinition.OutputDataConfig.KMSKeyID) {
				delta.Add("Spec.TrainingJobDefinition.OutputDataConfig.KMSKeyID", a.ko.Spec.TrainingJobDefinition.OutputDataConfig.KMSKeyID, b.ko.Spec.TrainingJobDefinition.OutputDataConfig.KMSKeyID)
			} else if a.ko.Spec.TrainingJobDefinition.OutputDataConfig.KMSKeyID != nil && b.ko.Spec.TrainingJobDefinition.OutputDataConfig.KMSKeyID != nil {
				if *a.ko.Spec.TrainingJobDefinition.OutputDataConfig.KMSKeyID != *b.ko.Spec.TrainingJobDefinition.OutputDataConfig.KMSKeyID {
					delta.Add("Spec.TrainingJobDefinition.OutputDataConfig.KMSKeyID", a.ko.Spec.TrainingJobDefinition.OutputDataConfig.KMSKeyID, b.ko.Spec.TrainingJobDefinition.OutputDataConfig.KMSKeyID)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.OutputDataConfig.S3OutputPath, b.ko.Spec.TrainingJobDefinition.OutputDataConfig.S3OutputPath) {
				delta.Add("Spec.TrainingJobDefinition.OutputDataConfig.S3OutputPath", a.ko.Spec.TrainingJobDefinition.OutputDataConfig.S3OutputPath, b.ko.Spec.TrainingJobDefinition.OutputDataConfig.S3OutputPath)
			} else if a.ko.Spec.TrainingJobDefinition.OutputDataConfig.S3OutputPath != nil && b.ko.Spec.TrainingJobDefinition.OutputDataConfig.S3OutputPath != nil {
				if *a.ko.Spec.TrainingJobDefinition.OutputDataConfig.S3OutputPath != *b.ko.Spec.TrainingJobDefinition.OutputDataConfig.S3OutputPath {
					delta.Add("Spec.TrainingJobDefinition.OutputDataConfig.S3OutputPath", a.ko.Spec.TrainingJobDefinition.OutputDataConfig.S3OutputPath, b.ko.Spec.TrainingJobDefinition.OutputDataConfig.S3OutputPath)
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.ResourceConfig, b.ko.Spec.TrainingJobDefinition.ResourceConfig) {
			delta.Add("Spec.TrainingJobDefinition.ResourceConfig", a.ko.Spec.TrainingJobDefinition.ResourceConfig, b.ko.Spec.TrainingJobDefinition.ResourceConfig)
		} else if a.ko.Spec.TrainingJobDefinition.ResourceConfig != nil && b.ko.Spec.TrainingJobDefinition.ResourceConfig != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.ResourceConfig.InstanceCount, b.ko.Spec.TrainingJobDefinition.ResourceConfig.InstanceCount) {
				delta.Add("Spec.TrainingJobDefinition.ResourceConfig.InstanceCount", a.ko.Spec.TrainingJobDefinition.ResourceConfig.InstanceCount, b.ko.Spec.TrainingJobDefinition.ResourceConfig.InstanceCount)
			} else if a.ko.Spec.TrainingJobDefinition.ResourceConfig.InstanceCount != nil && b.ko.Spec.TrainingJobDefinition.ResourceConfig.InstanceCount != nil {
				if *a.ko.Spec.TrainingJobDefinition.ResourceConfig.InstanceCount != *b.ko.Spec.TrainingJobDefinition.ResourceConfig.InstanceCount {
					delta.Add("Spec.TrainingJobDefinition.ResourceConfig.InstanceCount", a.ko.Spec.TrainingJobDefinition.ResourceConfig.InstanceCount, b.ko.Spec.TrainingJobDefinition.ResourceConfig.InstanceCount)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.ResourceConfig.InstanceType, b.ko.Spec.TrainingJobDefinition.ResourceConfig.InstanceType) {
				delta.Add("Spec.TrainingJobDefinition.ResourceConfig.InstanceType", a.ko.Spec.TrainingJobDefinition.ResourceConfig.InstanceType, b.ko.Spec.TrainingJobDefinition.ResourceConfig.InstanceType)
			} else if a.ko.Spec.TrainingJobDefinition.ResourceConfig.InstanceType != nil && b.ko.Spec.TrainingJobDefinition.ResourceConfig.InstanceType != nil {
				if *a.ko.Spec.TrainingJobDefinition.ResourceConfig.InstanceType != *b.ko.Spec.TrainingJobDefinition.ResourceConfig.InstanceType {
					delta.Add("Spec.TrainingJobDefinition.ResourceConfig.InstanceType", a.ko.Spec.TrainingJobDefinition.ResourceConfig.InstanceType, b.ko.Spec.TrainingJobDefinition.ResourceConfig.InstanceType)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.ResourceConfig.VolumeKMSKeyID, b.ko.Spec.TrainingJobDefinition.ResourceConfig.VolumeKMSKeyID) {
				delta.Add("Spec.TrainingJobDefinition.ResourceConfig.VolumeKMSKeyID", a.ko.Spec.TrainingJobDefinition.ResourceConfig.VolumeKMSKeyID, b.ko.Spec.TrainingJobDefinition.ResourceConfig.VolumeKMSKeyID)
			} else if a.ko.Spec.TrainingJobDefinition.ResourceConfig.VolumeKMSKeyID != nil && b.ko.Spec.TrainingJobDefinition.ResourceConfig.VolumeKMSKeyID != nil {
				if *a.ko.Spec.TrainingJobDefinition.ResourceConfig.VolumeKMSKeyID != *b.ko.Spec.TrainingJobDefinition.ResourceConfig.VolumeKMSKeyID {
					delta.Add("Spec.TrainingJobDefinition.ResourceConfig.VolumeKMSKeyID", a.ko.Spec.TrainingJobDefinition.ResourceConfig.VolumeKMSKeyID, b.ko.Spec.TrainingJobDefinition.ResourceConfig.VolumeKMSKeyID)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.ResourceConfig.VolumeSizeInGB, b.ko.Spec.TrainingJobDefinition.ResourceConfig.VolumeSizeInGB) {
				delta.Add("Spec.TrainingJobDefinition.ResourceConfig.VolumeSizeInGB", a.ko.Spec.TrainingJobDefinition.ResourceConfig.VolumeSizeInGB, b.ko.Spec.TrainingJobDefinition.ResourceConfig.VolumeSizeInGB)
			} else if a.ko.Spec.TrainingJobDefinition.ResourceConfig.VolumeSizeInGB != nil && b.ko.Spec.TrainingJobDefinition.ResourceConfig.VolumeSizeInGB != nil {
				if *a.ko.Spec.TrainingJobDefinition.ResourceConfig.VolumeSizeInGB != *b.ko.Spec.TrainingJobDefinition.ResourceConfig.VolumeSizeInGB {
					delta.Add("Spec.TrainingJobDefinition.ResourceConfig.VolumeSizeInGB", a.ko.Spec.TrainingJobDefinition.ResourceConfig.VolumeSizeInGB, b.ko.Spec.TrainingJobDefinition.ResourceConfig.VolumeSizeInGB)
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.RoleARN, b.ko.Spec.TrainingJobDefinition.RoleARN) {
			delta.Add("Spec.TrainingJobDefinition.RoleARN", a.ko.Spec.TrainingJobDefinition.RoleARN, b.ko.Spec.TrainingJobDefinition.RoleARN)
		} else if a.ko.Spec.TrainingJobDefinition.RoleARN != nil && b.ko.Spec.TrainingJobDefinition.RoleARN != nil {
			if *a.ko.Spec.TrainingJobDefinition.RoleARN != *b.ko.Spec.TrainingJobDefinition.RoleARN {
				delta.Add("Spec.TrainingJobDefinition.RoleARN", a.ko.Spec.TrainingJobDefinition.RoleARN, b.ko.Spec.TrainingJobDefinition.RoleARN)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.StaticHyperParameters, b.ko.Spec.TrainingJobDefinition.StaticHyperParameters) {
			delta.Add("Spec.TrainingJobDefinition.StaticHyperParameters", a.ko.Spec.TrainingJobDefinition.StaticHyperParameters, b.ko.Spec.TrainingJobDefinition.StaticHyperParameters)
		} else if a.ko.Spec.TrainingJobDefinition.StaticHyperParameters != nil && b.ko.Spec.TrainingJobDefinition.StaticHyperParameters != nil {
			if !ackcompare.MapStringStringPEqual(a.ko.Spec.TrainingJobDefinition.StaticHyperParameters, b.ko.Spec.TrainingJobDefinition.StaticHyperParameters) {
				delta.Add("Spec.TrainingJobDefinition.StaticHyperParameters", a.ko.Spec.TrainingJobDefinition.StaticHyperParameters, b.ko.Spec.TrainingJobDefinition.StaticHyperParameters)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.StoppingCondition, b.ko.Spec.TrainingJobDefinition.StoppingCondition) {
			delta.Add("Spec.TrainingJobDefinition.StoppingCondition", a.ko.Spec.TrainingJobDefinition.StoppingCondition, b.ko.Spec.TrainingJobDefinition.StoppingCondition)
		} else if a.ko.Spec.TrainingJobDefinition.StoppingCondition != nil && b.ko.Spec.TrainingJobDefinition.StoppingCondition != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.StoppingCondition.MaxRuntimeInSeconds, b.ko.Spec.TrainingJobDefinition.StoppingCondition.MaxRuntimeInSeconds) {
				delta.Add("Spec.TrainingJobDefinition.StoppingCondition.MaxRuntimeInSeconds", a.ko.Spec.TrainingJobDefinition.StoppingCondition.MaxRuntimeInSeconds, b.ko.Spec.TrainingJobDefinition.StoppingCondition.MaxRuntimeInSeconds)
			} else if a.ko.Spec.TrainingJobDefinition.StoppingCondition.MaxRuntimeInSeconds != nil && b.ko.Spec.TrainingJobDefinition.StoppingCondition.MaxRuntimeInSeconds != nil {
				if *a.ko.Spec.TrainingJobDefinition.StoppingCondition.MaxRuntimeInSeconds != *b.ko.Spec.TrainingJobDefinition.StoppingCondition.MaxRuntimeInSeconds {
					delta.Add("Spec.TrainingJobDefinition.StoppingCondition.MaxRuntimeInSeconds", a.ko.Spec.TrainingJobDefinition.StoppingCondition.MaxRuntimeInSeconds, b.ko.Spec.TrainingJobDefinition.StoppingCondition.MaxRuntimeInSeconds)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.StoppingCondition.MaxWaitTimeInSeconds, b.ko.Spec.TrainingJobDefinition.StoppingCondition.MaxWaitTimeInSeconds) {
				delta.Add("Spec.TrainingJobDefinition.StoppingCondition.MaxWaitTimeInSeconds", a.ko.Spec.TrainingJobDefinition.StoppingCondition.MaxWaitTimeInSeconds, b.ko.Spec.TrainingJobDefinition.StoppingCondition.MaxWaitTimeInSeconds)
			} else if a.ko.Spec.TrainingJobDefinition.StoppingCondition.MaxWaitTimeInSeconds != nil && b.ko.Spec.TrainingJobDefinition.StoppingCondition.MaxWaitTimeInSeconds != nil {
				if *a.ko.Spec.TrainingJobDefinition.StoppingCondition.MaxWaitTimeInSeconds != *b.ko.Spec.TrainingJobDefinition.StoppingCondition.MaxWaitTimeInSeconds {
					delta.Add("Spec.TrainingJobDefinition.StoppingCondition.MaxWaitTimeInSeconds", a.ko.Spec.TrainingJobDefinition.StoppingCondition.MaxWaitTimeInSeconds, b.ko.Spec.TrainingJobDefinition.StoppingCondition.MaxWaitTimeInSeconds)
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.TuningObjective, b.ko.Spec.TrainingJobDefinition.TuningObjective) {
			delta.Add("Spec.TrainingJobDefinition.TuningObjective", a.ko.Spec.TrainingJobDefinition.TuningObjective, b.ko.Spec.TrainingJobDefinition.TuningObjective)
		} else if a.ko.Spec.TrainingJobDefinition.TuningObjective != nil && b.ko.Spec.TrainingJobDefinition.TuningObjective != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.TuningObjective.MetricName, b.ko.Spec.TrainingJobDefinition.TuningObjective.MetricName) {
				delta.Add("Spec.TrainingJobDefinition.TuningObjective.MetricName", a.ko.Spec.TrainingJobDefinition.TuningObjective.MetricName, b.ko.Spec.TrainingJobDefinition.TuningObjective.MetricName)
			} else if a.ko.Spec.TrainingJobDefinition.TuningObjective.MetricName != nil && b.ko.Spec.TrainingJobDefinition.TuningObjective.MetricName != nil {
				if *a.ko.Spec.TrainingJobDefinition.TuningObjective.MetricName != *b.ko.Spec.TrainingJobDefinition.TuningObjective.MetricName {
					delta.Add("Spec.TrainingJobDefinition.TuningObjective.MetricName", a.ko.Spec.TrainingJobDefinition.TuningObjective.MetricName, b.ko.Spec.TrainingJobDefinition.TuningObjective.MetricName)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.TuningObjective.Type, b.ko.Spec.TrainingJobDefinition.TuningObjective.Type) {
				delta.Add("Spec.TrainingJobDefinition.TuningObjective.Type", a.ko.Spec.TrainingJobDefinition.TuningObjective.Type, b.ko.Spec.TrainingJobDefinition.TuningObjective.Type)
			} else if a.ko.Spec.TrainingJobDefinition.TuningObjective.Type != nil && b.ko.Spec.TrainingJobDefinition.TuningObjective.Type != nil {
				if *a.ko.Spec.TrainingJobDefinition.TuningObjective.Type != *b.ko.Spec.TrainingJobDefinition.TuningObjective.Type {
					delta.Add("Spec.TrainingJobDefinition.TuningObjective.Type", a.ko.Spec.TrainingJobDefinition.TuningObjective.Type, b.ko.Spec.TrainingJobDefinition.TuningObjective.Type)
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobDefinition.VPCConfig, b.ko.Spec.TrainingJobDefinition.VPCConfig) {
			delta.Add("Spec.TrainingJobDefinition.VPCConfig", a.ko.Spec.TrainingJobDefinition.VPCConfig, b.ko.Spec.TrainingJobDefinition.VPCConfig)
		} else if a.ko.Spec.TrainingJobDefinition.VPCConfig != nil && b.ko.Spec.TrainingJobDefinition.VPCConfig != nil {

			if !ackcompare.SliceStringPEqual(a.ko.Spec.TrainingJobDefinition.VPCConfig.SecurityGroupIDs, b.ko.Spec.TrainingJobDefinition.VPCConfig.SecurityGroupIDs) {
				delta.Add("Spec.TrainingJobDefinition.VPCConfig.SecurityGroupIDs", a.ko.Spec.TrainingJobDefinition.VPCConfig.SecurityGroupIDs, b.ko.Spec.TrainingJobDefinition.VPCConfig.SecurityGroupIDs)
			}

			if !ackcompare.SliceStringPEqual(a.ko.Spec.TrainingJobDefinition.VPCConfig.Subnets, b.ko.Spec.TrainingJobDefinition.VPCConfig.Subnets) {
				delta.Add("Spec.TrainingJobDefinition.VPCConfig.Subnets", a.ko.Spec.TrainingJobDefinition.VPCConfig.Subnets, b.ko.Spec.TrainingJobDefinition.VPCConfig.Subnets)
			}
		}
	}

	if ackcompare.HasNilDifference(a.ko.Spec.WarmStartConfig, b.ko.Spec.WarmStartConfig) {
		delta.Add("Spec.WarmStartConfig", a.ko.Spec.WarmStartConfig, b.ko.Spec.WarmStartConfig)
	} else if a.ko.Spec.WarmStartConfig != nil && b.ko.Spec.WarmStartConfig != nil {

		if ackcompare.HasNilDifference(a.ko.Spec.WarmStartConfig.WarmStartType, b.ko.Spec.WarmStartConfig.WarmStartType) {
			delta.Add("Spec.WarmStartConfig.WarmStartType", a.ko.Spec.WarmStartConfig.WarmStartType, b.ko.Spec.WarmStartConfig.WarmStartType)
		} else if a.ko.Spec.WarmStartConfig.WarmStartType != nil && b.ko.Spec.WarmStartConfig.WarmStartType != nil {
			if *a.ko.Spec.WarmStartConfig.WarmStartType != *b.ko.Spec.WarmStartConfig.WarmStartType {
				delta.Add("Spec.WarmStartConfig.WarmStartType", a.ko.Spec.WarmStartConfig.WarmStartType, b.ko.Spec.WarmStartConfig.WarmStartType)
			}
		}
	}

	return delta
}
