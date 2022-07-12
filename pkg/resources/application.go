package resources

import (
	awscdk "github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	elb "github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type UmaAppliationStackProps struct {
	awscdk.StackProps
	AppName *string
	Vpc     *awsec2.IVpc
}

type UmaApplicationStackResponse struct {
	Stack awscdk.Stack
}

func NewUmaApplicationStack(scope constructs.Construct, id string, props *UmaAppliationStackProps) *UmaApplicationStackResponse {
	// boilerplate
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// ECS Cluster
	cluster := awsecs.NewCluster(stack, jsii.String("FargoECSCluster"), &awsecs.ClusterProps{
		ClusterName: props.AppName,
		Vpc:         *props.Vpc,
	})

	// Create Task Definition
	taskDef := awsecs.NewFargateTaskDefinition(stack, jsii.String("FargoTaskDef"),
		&awsecs.FargateTaskDefinitionProps{
			MemoryLimitMiB: jsii.Number(512),
			Cpu:            jsii.Number(256),
		})

	container := taskDef.AddContainer(jsii.String("FargoContainer"), &awsecs.ContainerDefinitionOptions{
		Image: awsecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample"),
			&awsecs.RepositoryImageProps{}),
	})
	container.AddPortMappings(&awsecs.PortMapping{
		ContainerPort: jsii.Number(80),
		Protocol:      awsecs.Protocol_TCP,
	})

	// Create Fargate Service
	service := awsecs.NewFargateService(stack, jsii.String("FargoService"), &awsecs.FargateServiceProps{
		Cluster:        cluster,
		TaskDefinition: taskDef,
	})

	// Create ALB
	lb := elb.NewApplicationLoadBalancer(stack, jsii.String("LB"), &elb.ApplicationLoadBalancerProps{
		Vpc:            *props.Vpc,
		InternetFacing: jsii.Bool(true),
	})
	listener := lb.AddListener(jsii.String("PublicListener"), &elb.BaseApplicationListenerProps{
		Port: jsii.Number(80),
		Open: jsii.Bool(true),
	})

	// Attach ALB to Fargate Service
	listener.AddTargets(jsii.String("Fargo"), &elb.AddApplicationTargetsProps{
		Port: jsii.Number(80),
		Targets: &[]elb.IApplicationLoadBalancerTarget{
			service.LoadBalancerTarget(&awsecs.LoadBalancerTargetOptions{
				ContainerName: jsii.String("FargoContainer"),
				ContainerPort: jsii.Number(80),
			}),
		},
	})

	awscdk.NewCfnOutput(stack, jsii.String("LoadBalancerDNS"), &awscdk.CfnOutputProps{Value: lb.LoadBalancerDnsName()})

	return &UmaApplicationStackResponse{
		stack,
	}
}
