package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

//https://pkg.go.dev/github.com/aws/aws-cdk-go/awscdk/v2@v2.30.0/awslambda

type Lambda struct {
	awscdk.StackProps
	handler awslambda.Handler
}

//func NewLambdaFunction(scope constructs.Construct, id string, props) {
//	awslambda.NewFunction(scope)
//	//fn := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
//	//	runtime: lambda.runtime_NODEJS_16_X(),
//	//	handler: jsii.String("index.handler"),
//	//	code: lambda.code.fromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//	//})
//}
