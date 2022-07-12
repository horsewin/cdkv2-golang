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
	VpcName           *string
	CidrRange         *string
	AvailabilityZones *[]*string
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

	// props check
	var cidrProps string
	if props.CidrRange != "" {
		cidrProps = props.CidrRange
	}

	// VPC
	vpc := awsec2.NewVpc(stack, jsii.String("Vpc"), &awsec2.VpcProps{
		Cidr:              jsii.String(cidrProps),
		VpcName:           props.VpcName,
		NatGateways:       jsii.Number(1),
		AvailabilityZones: props.AvailabilityZones,
	})

	return &UmaNetworkStackResponse{
		Stack: stack,
		Vpc:   vpc,
	}
}
