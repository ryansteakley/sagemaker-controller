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

package transform_job

import (
	"context"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/sagemaker-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.SageMaker{}
	_ = &svcapitypes.TransformJob{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer exit(err)
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.DescribeTransformJobOutput
	resp, err = rm.sdkapi.DescribeTransformJobWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeTransformJob", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "ValidationException" && strings.HasPrefix(awsErr.Message(), "Could not find requested job with name") {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.BatchStrategy != nil {
		ko.Spec.BatchStrategy = resp.BatchStrategy
	} else {
		ko.Spec.BatchStrategy = nil
	}
	if resp.DataProcessing != nil {
		f3 := &svcapitypes.DataProcessing{}
		if resp.DataProcessing.InputFilter != nil {
			f3.InputFilter = resp.DataProcessing.InputFilter
		}
		if resp.DataProcessing.JoinSource != nil {
			f3.JoinSource = resp.DataProcessing.JoinSource
		}
		if resp.DataProcessing.OutputFilter != nil {
			f3.OutputFilter = resp.DataProcessing.OutputFilter
		}
		ko.Spec.DataProcessing = f3
	} else {
		ko.Spec.DataProcessing = nil
	}
	if resp.Environment != nil {
		f4 := map[string]*string{}
		for f4key, f4valiter := range resp.Environment {
			var f4val string
			f4val = *f4valiter
			f4[f4key] = &f4val
		}
		ko.Spec.Environment = f4
	} else {
		ko.Spec.Environment = nil
	}
	if resp.ExperimentConfig != nil {
		f5 := &svcapitypes.ExperimentConfig{}
		if resp.ExperimentConfig.ExperimentName != nil {
			f5.ExperimentName = resp.ExperimentConfig.ExperimentName
		}
		if resp.ExperimentConfig.TrialComponentDisplayName != nil {
			f5.TrialComponentDisplayName = resp.ExperimentConfig.TrialComponentDisplayName
		}
		if resp.ExperimentConfig.TrialName != nil {
			f5.TrialName = resp.ExperimentConfig.TrialName
		}
		ko.Spec.ExperimentConfig = f5
	} else {
		ko.Spec.ExperimentConfig = nil
	}
	if resp.FailureReason != nil {
		ko.Status.FailureReason = resp.FailureReason
	} else {
		ko.Status.FailureReason = nil
	}
	if resp.MaxConcurrentTransforms != nil {
		ko.Spec.MaxConcurrentTransforms = resp.MaxConcurrentTransforms
	} else {
		ko.Spec.MaxConcurrentTransforms = nil
	}
	if resp.MaxPayloadInMB != nil {
		ko.Spec.MaxPayloadInMB = resp.MaxPayloadInMB
	} else {
		ko.Spec.MaxPayloadInMB = nil
	}
	if resp.ModelClientConfig != nil {
		f10 := &svcapitypes.ModelClientConfig{}
		if resp.ModelClientConfig.InvocationsMaxRetries != nil {
			f10.InvocationsMaxRetries = resp.ModelClientConfig.InvocationsMaxRetries
		}
		if resp.ModelClientConfig.InvocationsTimeoutInSeconds != nil {
			f10.InvocationsTimeoutInSeconds = resp.ModelClientConfig.InvocationsTimeoutInSeconds
		}
		ko.Spec.ModelClientConfig = f10
	} else {
		ko.Spec.ModelClientConfig = nil
	}
	if resp.ModelName != nil {
		ko.Spec.ModelName = resp.ModelName
	} else {
		ko.Spec.ModelName = nil
	}
	if resp.TransformInput != nil {
		f13 := &svcapitypes.TransformInput{}
		if resp.TransformInput.CompressionType != nil {
			f13.CompressionType = resp.TransformInput.CompressionType
		}
		if resp.TransformInput.ContentType != nil {
			f13.ContentType = resp.TransformInput.ContentType
		}
		if resp.TransformInput.DataSource != nil {
			f13f2 := &svcapitypes.TransformDataSource{}
			if resp.TransformInput.DataSource.S3DataSource != nil {
				f13f2f0 := &svcapitypes.TransformS3DataSource{}
				if resp.TransformInput.DataSource.S3DataSource.S3DataType != nil {
					f13f2f0.S3DataType = resp.TransformInput.DataSource.S3DataSource.S3DataType
				}
				if resp.TransformInput.DataSource.S3DataSource.S3Uri != nil {
					f13f2f0.S3URI = resp.TransformInput.DataSource.S3DataSource.S3Uri
				}
				f13f2.S3DataSource = f13f2f0
			}
			f13.DataSource = f13f2
		}
		if resp.TransformInput.SplitType != nil {
			f13.SplitType = resp.TransformInput.SplitType
		}
		ko.Spec.TransformInput = f13
	} else {
		ko.Spec.TransformInput = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.TransformJobArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.TransformJobArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.TransformJobName != nil {
		ko.Spec.TransformJobName = resp.TransformJobName
	} else {
		ko.Spec.TransformJobName = nil
	}
	if resp.TransformJobStatus != nil {
		ko.Status.TransformJobStatus = resp.TransformJobStatus
	} else {
		ko.Status.TransformJobStatus = nil
	}
	if resp.TransformOutput != nil {
		f17 := &svcapitypes.TransformOutput{}
		if resp.TransformOutput.Accept != nil {
			f17.Accept = resp.TransformOutput.Accept
		}
		if resp.TransformOutput.AssembleWith != nil {
			f17.AssembleWith = resp.TransformOutput.AssembleWith
		}
		if resp.TransformOutput.KmsKeyId != nil {
			f17.KMSKeyID = resp.TransformOutput.KmsKeyId
		}
		if resp.TransformOutput.S3OutputPath != nil {
			f17.S3OutputPath = resp.TransformOutput.S3OutputPath
		}
		ko.Spec.TransformOutput = f17
	} else {
		ko.Spec.TransformOutput = nil
	}
	if resp.TransformResources != nil {
		f18 := &svcapitypes.TransformResources{}
		if resp.TransformResources.InstanceCount != nil {
			f18.InstanceCount = resp.TransformResources.InstanceCount
		}
		if resp.TransformResources.InstanceType != nil {
			f18.InstanceType = resp.TransformResources.InstanceType
		}
		if resp.TransformResources.VolumeKmsKeyId != nil {
			f18.VolumeKMSKeyID = resp.TransformResources.VolumeKmsKeyId
		}
		ko.Spec.TransformResources = f18
	} else {
		ko.Spec.TransformResources = nil
	}

	rm.setStatusDefaults(ko)
	rm.customSetOutput(r, resp.TransformJobStatus, ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.TransformJobName == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.DescribeTransformJobInput, error) {
	res := &svcsdk.DescribeTransformJobInput{}

	if r.ko.Spec.TransformJobName != nil {
		res.SetTransformJobName(*r.ko.Spec.TransformJobName)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer exit(err)
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateTransformJobOutput
	_ = resp
	resp, err = rm.sdkapi.CreateTransformJobWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateTransformJob", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.TransformJobArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.TransformJobArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}

	rm.setStatusDefaults(ko)
	rm.customSetOutput(desired, aws.String(svcsdk.TransformJobStatusInProgress), ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateTransformJobInput, error) {
	res := &svcsdk.CreateTransformJobInput{}

	if r.ko.Spec.BatchStrategy != nil {
		res.SetBatchStrategy(*r.ko.Spec.BatchStrategy)
	}
	if r.ko.Spec.DataProcessing != nil {
		f1 := &svcsdk.DataProcessing{}
		if r.ko.Spec.DataProcessing.InputFilter != nil {
			f1.SetInputFilter(*r.ko.Spec.DataProcessing.InputFilter)
		}
		if r.ko.Spec.DataProcessing.JoinSource != nil {
			f1.SetJoinSource(*r.ko.Spec.DataProcessing.JoinSource)
		}
		if r.ko.Spec.DataProcessing.OutputFilter != nil {
			f1.SetOutputFilter(*r.ko.Spec.DataProcessing.OutputFilter)
		}
		res.SetDataProcessing(f1)
	}
	if r.ko.Spec.Environment != nil {
		f2 := map[string]*string{}
		for f2key, f2valiter := range r.ko.Spec.Environment {
			var f2val string
			f2val = *f2valiter
			f2[f2key] = &f2val
		}
		res.SetEnvironment(f2)
	}
	if r.ko.Spec.ExperimentConfig != nil {
		f3 := &svcsdk.ExperimentConfig{}
		if r.ko.Spec.ExperimentConfig.ExperimentName != nil {
			f3.SetExperimentName(*r.ko.Spec.ExperimentConfig.ExperimentName)
		}
		if r.ko.Spec.ExperimentConfig.TrialComponentDisplayName != nil {
			f3.SetTrialComponentDisplayName(*r.ko.Spec.ExperimentConfig.TrialComponentDisplayName)
		}
		if r.ko.Spec.ExperimentConfig.TrialName != nil {
			f3.SetTrialName(*r.ko.Spec.ExperimentConfig.TrialName)
		}
		res.SetExperimentConfig(f3)
	}
	if r.ko.Spec.MaxConcurrentTransforms != nil {
		res.SetMaxConcurrentTransforms(*r.ko.Spec.MaxConcurrentTransforms)
	}
	if r.ko.Spec.MaxPayloadInMB != nil {
		res.SetMaxPayloadInMB(*r.ko.Spec.MaxPayloadInMB)
	}
	if r.ko.Spec.ModelClientConfig != nil {
		f6 := &svcsdk.ModelClientConfig{}
		if r.ko.Spec.ModelClientConfig.InvocationsMaxRetries != nil {
			f6.SetInvocationsMaxRetries(*r.ko.Spec.ModelClientConfig.InvocationsMaxRetries)
		}
		if r.ko.Spec.ModelClientConfig.InvocationsTimeoutInSeconds != nil {
			f6.SetInvocationsTimeoutInSeconds(*r.ko.Spec.ModelClientConfig.InvocationsTimeoutInSeconds)
		}
		res.SetModelClientConfig(f6)
	}
	if r.ko.Spec.ModelName != nil {
		res.SetModelName(*r.ko.Spec.ModelName)
	}
	if r.ko.Spec.Tags != nil {
		f8 := []*svcsdk.Tag{}
		for _, f8iter := range r.ko.Spec.Tags {
			f8elem := &svcsdk.Tag{}
			if f8iter.Key != nil {
				f8elem.SetKey(*f8iter.Key)
			}
			if f8iter.Value != nil {
				f8elem.SetValue(*f8iter.Value)
			}
			f8 = append(f8, f8elem)
		}
		res.SetTags(f8)
	}
	if r.ko.Spec.TransformInput != nil {
		f9 := &svcsdk.TransformInput{}
		if r.ko.Spec.TransformInput.CompressionType != nil {
			f9.SetCompressionType(*r.ko.Spec.TransformInput.CompressionType)
		}
		if r.ko.Spec.TransformInput.ContentType != nil {
			f9.SetContentType(*r.ko.Spec.TransformInput.ContentType)
		}
		if r.ko.Spec.TransformInput.DataSource != nil {
			f9f2 := &svcsdk.TransformDataSource{}
			if r.ko.Spec.TransformInput.DataSource.S3DataSource != nil {
				f9f2f0 := &svcsdk.TransformS3DataSource{}
				if r.ko.Spec.TransformInput.DataSource.S3DataSource.S3DataType != nil {
					f9f2f0.SetS3DataType(*r.ko.Spec.TransformInput.DataSource.S3DataSource.S3DataType)
				}
				if r.ko.Spec.TransformInput.DataSource.S3DataSource.S3URI != nil {
					f9f2f0.SetS3Uri(*r.ko.Spec.TransformInput.DataSource.S3DataSource.S3URI)
				}
				f9f2.SetS3DataSource(f9f2f0)
			}
			f9.SetDataSource(f9f2)
		}
		if r.ko.Spec.TransformInput.SplitType != nil {
			f9.SetSplitType(*r.ko.Spec.TransformInput.SplitType)
		}
		res.SetTransformInput(f9)
	}
	if r.ko.Spec.TransformJobName != nil {
		res.SetTransformJobName(*r.ko.Spec.TransformJobName)
	}
	if r.ko.Spec.TransformOutput != nil {
		f11 := &svcsdk.TransformOutput{}
		if r.ko.Spec.TransformOutput.Accept != nil {
			f11.SetAccept(*r.ko.Spec.TransformOutput.Accept)
		}
		if r.ko.Spec.TransformOutput.AssembleWith != nil {
			f11.SetAssembleWith(*r.ko.Spec.TransformOutput.AssembleWith)
		}
		if r.ko.Spec.TransformOutput.KMSKeyID != nil {
			f11.SetKmsKeyId(*r.ko.Spec.TransformOutput.KMSKeyID)
		}
		if r.ko.Spec.TransformOutput.S3OutputPath != nil {
			f11.SetS3OutputPath(*r.ko.Spec.TransformOutput.S3OutputPath)
		}
		res.SetTransformOutput(f11)
	}
	if r.ko.Spec.TransformResources != nil {
		f12 := &svcsdk.TransformResources{}
		if r.ko.Spec.TransformResources.InstanceCount != nil {
			f12.SetInstanceCount(*r.ko.Spec.TransformResources.InstanceCount)
		}
		if r.ko.Spec.TransformResources.InstanceType != nil {
			f12.SetInstanceType(*r.ko.Spec.TransformResources.InstanceType)
		}
		if r.ko.Spec.TransformResources.VolumeKMSKeyID != nil {
			f12.SetVolumeKmsKeyId(*r.ko.Spec.TransformResources.VolumeKMSKeyID)
		}
		res.SetTransformResources(f12)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {
	// TODO(jaypipes): Figure this out...
	return nil, ackerr.NotImplemented
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer exit(err)
	// Call StopTranformJob only if the job is InProgress, otherwise just return nil to mark the
	// resource Unmanaged
	latestStatus := r.ko.Status.TransformJobStatus
	if latestStatus != nil && *latestStatus != svcsdk.TransformJobStatusInProgress {
		return r, err
	}
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.StopTransformJobOutput
	_ = resp
	resp, err = rm.sdkapi.StopTransformJobWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "StopTransformJob", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.StopTransformJobInput, error) {
	res := &svcsdk.StopTransformJobInput{}

	if r.ko.Spec.TransformJobName != nil {
		res.SetTransformJobName(*r.ko.Spec.TransformJobName)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.TransformJob,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}

	if rm.terminalAWSError(err) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		terminalCondition.Status = corev1.ConditionTrue
		awsErr, _ := ackerr.AWSError(err)
		errorMessage := awsErr.Message()
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Message()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	if err == nil {
		return false
	}
	awsErr, ok := ackerr.AWSError(err)
	if !ok {
		return false
	}
	switch awsErr.Code() {
	case "ResourceLimitExceeded",
		"ResourceNotFound",
		"ResourceInUse",
		"OptInRequired",
		"InvalidParameterCombination",
		"InvalidParameterValue",
		"MissingParameter",
		"MissingAction",
		"InvalidClientTokenId",
		"InvalidQueryParameter",
		"MalformedQueryString",
		"InvalidAction",
		"UnrecognizedClientException":
		return true
	default:
		return false
	}
}
