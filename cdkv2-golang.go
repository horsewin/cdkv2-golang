package main

import (
	"cdkv2-golang/pkg/resources"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"os"
)

type Cdkv2GolangStackProps struct {
	awscdk.StackProps
}

func NewCdkv2GolangStack(scope constructs.Construct, id string, props *Cdkv2GolangStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// example resource
	awssqs.NewQueue(stack, jsii.String("uma-queue"), &awssqs.QueueProps{
		VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	})

	lambdaFn := awslambda.NewFunction(stack, jsii.String("Singleton"), &awslambda.FunctionProps{
		Code:    awslambda.NewAssetCode(jsii.String("lambda/main"), nil),
		Handler: jsii.String("handler.main"),
		Timeout: awscdk.Duration_Seconds(jsii.Number(300)),
		Runtime: awslambda.Runtime_GO_1_X(),
	})

	// Run every day at 6PM UTC
	// See https://docs.aws.amazon.com/lambda/latest/dg/tutorial-scheduled-events-schedule-expressions.html
	rule := awsevents.NewRule(stack, jsii.String("Rule"), &awsevents.RuleProps{
		Schedule: awsevents.Schedule_Expression(jsii.String("cron(0 18 ? * MON-FRI *)")),
	})

	rule.AddTarget(awseventstargets.NewLambdaFunction(lambdaFn, nil))

	return stack
}

func main() {
	app := awscdk.NewApp(nil)
	sprops := awscdk.StackProps{
		Env: env(),
	}

	NewCdkv2GolangStack(app, "UmaStack", &Cdkv2GolangStackProps{
		sprops,
	})

	// The code that defines your stack goes here
	resources.NewUmaNetworkStack(app, "Network", &resources.UmaNetworkStackProps{
		StackProps: sprops,
		Name:       jsii.String("uma-vpc"),
		//AvailabilityZones: jsii.Strings("ap-northeast-1c", "ap-northeast-1d"),
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
