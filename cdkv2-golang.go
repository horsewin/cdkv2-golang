package main

import (
	"cdkv2-golang/pkg/resources"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
	"os"
)

type Cdkv2GolangStackProps struct {
	awscdk.StackProps
}

func main() {
	app := awscdk.NewApp(nil)
	sprops := awscdk.StackProps{
		Env: env(),
	}

	network := resources.NewUmaNetworkStack(app, "Network", &resources.UmaNetworkStackProps{
		StackProps:        sprops,
		VpcName:           jsii.String("uma-vpc"),
		AvailabilityZones: jsii.Strings("ap-northeast-1c", "ap-northeast-1d"),
	})

	resources.NewUmaApplicationStack(app, "Application", &resources.UmaAppliationStackProps{
		StackProps: sprops,
		AppName:    jsii.String("uma-app"),
		Vpc:        &network.Vpc,
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
