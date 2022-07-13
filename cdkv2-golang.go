package main

import (
	"cdkv2-golang/pkg/resources"
	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
	"os"
)

func main() {
	app := cdk.NewApp(nil)
	sprops := cdk.StackProps{
		Env: env(),
	}

	// The code that defines your stack goes here
	resources.NewUmaNetworkStack(app, "Network", &resources.UmaNetworkStackProps{
		StackProps: sprops,
		VpcName:    jsii.String("uma-cdk2-vpc"),
		//AvailabilityZones: jsii.Strings("ap-northeast-1c", "ap-northeast-1d"),
	})

	app.Synth(nil)
}

func env() *cdk.Environment {
	return &cdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
