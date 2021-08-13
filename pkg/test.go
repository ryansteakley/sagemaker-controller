package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
)

// NewSession returns a new session object. By default the returned session is
// created using pod IRSA environment variables. If assumeRoleARN is not empty,
// NewSession will call STS::AssumeRole and use the returned credentials to create
// the session.
func main() {
	awsCfg := aws.Config{
		Region:              aws.String("us-west-2"),
		STSRegionalEndpoint: endpoints.RegionalSTSEndpoint,
	}
	endpointURL := "https://api-fips.sagemaker.us-west-2.amazonaws.com"
	pointerURL := &endpointURL
	if pointerURL != nil {
		fmt.Println("endpoint not none")
		awsCfg.Endpoint = pointerURL
	}
	if svcsdk.EndpointsID == "api.sagemaker" {
		fmt.Println("SAD")
	}
	sess, err := session.NewSession(&awsCfg)
	if err != nil {
		fmt.Println(err)
	}
	assumeRoleARN := ""
	//"arn:aws:iam::024825446905:role/ack-sage-role-bugbash-may-14"
	if assumeRoleARN != "" {
		// call STS::AssumeRole
		creds := stscreds.NewCredentials(sess, string(assumeRoleARN))
		// recreate session with the new credentials
		awsCfg.Credentials = creds

		sess, err = session.NewSession(&awsCfg)
		if err != nil {
			fmt.Println(err)
		}
	}
	svc := svcsdk.New(sess)
	res, err := svc.CreateEndpoint(&svcsdk.CreateEndpointInput{
		EndpointConfigName: aws.String("privatelink-conifg"),
		EndpointName:       aws.String("aws-sdk-go-model-endpoint"),
	})
	fmt.Println(res)
	fmt.Println(err)
}
