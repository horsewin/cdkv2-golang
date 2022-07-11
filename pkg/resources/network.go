package resources

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type UmaNetworkStackProps struct {
	awscdk.StackProps
	// Golangでは外に公開するモジュールはPascal-caseである必要がある
	VpcName *string
}

type UmaNetworkStackResponse struct {
	Stack awscdk.Stack
	Vpc   awsec2.IVpc
}

func NewUmaNetworkStack(scope constructs.Construct, id string, props *UmaNetworkStackProps) *UmaNetworkStackResponse {
	// boilerplate
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// VPC
	vpc := awsec2.NewVpc(stack, jsii.String("Vpc"), &awsec2.VpcProps{
		Cidr:        jsii.String("10.10.0.0/16"),
		VpcName:     props.VpcName,
		NatGateways: jsii.Number(1),
		MaxAzs:      jsii.Number(2),
	})

	return &UmaNetworkStackResponse{
		Stack: stack,
		Vpc:   vpc,
	}
}
