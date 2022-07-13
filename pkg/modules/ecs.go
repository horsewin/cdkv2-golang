package modules

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	elb "github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type UmaEcsProps struct {
	ClusterName *string
	Vpc         *awsec2.IVpc
}

type UmaEcsResponse struct {
	Cluster awsecs.ICluster
	Service awsecs.IService
	Lb      elb.ILoadBalancerV2
}

func NewEcsTemplate(scope constructs.Construct, id string, props *UmaEcsProps) *UmaEcsResponse {

	// ECS Cluster
	cluster := awsecs.NewCluster(scope, jsii.String("uma-cdk2-cluster"), &awsecs.ClusterProps{
		ClusterName: props.ClusterName,
		Vpc:         *props.Vpc,
	})

	// Create Task Definition
	taskDef := awsecs.NewFargateTaskDefinition(scope, jsii.String("uma-cdk2-taskdef"),
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
	service := awsecs.NewFargateService(scope, jsii.String("uma-cdk2-service"), &awsecs.FargateServiceProps{
		Cluster:        cluster,
		TaskDefinition: taskDef,
	})

	// Create ALB
	lb := elb.NewApplicationLoadBalancer(scope, jsii.String("uma-cdk2-lb"), &elb.ApplicationLoadBalancerProps{
		Vpc:            *props.Vpc,
		InternetFacing: jsii.Bool(true),
	})
	listener := lb.AddListener(jsii.String("uma-cdk2-publicListener"), &elb.BaseApplicationListenerProps{
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

	awscdk.NewCfnOutput(scope, jsii.String("LoadBalancerDNS"), &awscdk.CfnOutputProps{Value: lb.LoadBalancerDnsName()})

	return &UmaEcsResponse{
		Cluster: cluster,
		Service: service,
		Lb:      lb,
	}
}
