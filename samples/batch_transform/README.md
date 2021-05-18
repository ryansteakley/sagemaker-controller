# Bring Your Own Batch Transform Job Sample

This sample demonstrates how to start batch-transform jobs using your own batch-transform script, packaged in a SageMaker-compatible container, using the Amazon AWS Controllers for Kubernetes (ACK) service controller for Amazon SageMaker.                     

## Prerequisites

This sample assumes that you have already configured an EKS cluster with the ACK operator. It also assumes that you have installed `kubectl` - you can find a link on our [installation page](To do).

In order to follow this script, you must first create a batch-transform script packaged in a Dockerfile that is [compatible with Amazon SageMaker](https://docs.aws.amazon.com/sagemaker/latest/dg/amazon-sagemaker-containers.html). 

You will also need a model in SageMaker for this sample. If you do not have one you must first create a [model](/samples/model/README.md)
## Preparing the Batch-transform Script


### Updating the Batch-transform Specification

In the `my-batch-transform-job.yaml` file, modify the placeholder values with those associated with your account and batch-transform job. The `spec.algorithmSpecification.modelName` should be the model from the previous step. The `spec.roleARN` field should be the ARN of an IAM role which has permissions to access your S3 resources. If you have not yet created a role with these permissions, you can find an example policy at [Amazon SageMaker Roles](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html#sagemaker-roles-createbatch-transformjob-perms).

## Submitting your Batch-transform Job

To submit your prepared batch-transform job specification, apply the specification to your EKS cluster as such:
```
$ kubectl apply -f my-batch-transform-job.yaml
batch-transformjob.sagemaker.services.k8s.aws.amazon.com/my-batch-transform-job created
```

To monitor the batch-transform job status, you can use the following command:
```
$ kubectl get batch-transformjob my-batch-transform-job
```

To monitor the batch-transform job in-depth once it has started, you can see the full status and any additional errors with the following command:
```
$ kubectl describe batch-transformjob my-batch-transform-job
```

To delete the batch-transform job once it has started if errors occurred or for any reason with the following command:
```
$ kubectl delete batch-transformjob my-batch-transform-job
```
