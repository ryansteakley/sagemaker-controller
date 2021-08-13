import boto3

client = boto3.client('sagemaker')
client.list_model_packages(ModelPackage