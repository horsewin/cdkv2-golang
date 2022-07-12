package resources

import (
	"cdkv2-golang/pkg/modules"
	awscdk "github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type UmaApplicationStackProps struct {
	awscdk.StackProps
	AppName *string
	Vpc     *awsec2.IVpc
}

type UmaApplicationStackResponse struct {
	Stack awscdk.Stack
}

func NewUmaApplicationStack(scope constructs.Construct, id string, props *UmaApplicationStackProps) *UmaApplicationStackResponse {
	// boilerplate
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// Create each app
	modules.NewEcsTemplate(stack, "uma-ecs", &modules.UmaEcsProps{
		ClusterName: jsii.String("uma-cluster"),
		Vpc:         props.Vpc,
	})

	return &UmaApplicationStackResponse{
		stack,
	}
}
