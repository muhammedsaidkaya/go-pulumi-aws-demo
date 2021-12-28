package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createSecurityGroup(ctx *pulumi.Context, vpcId pulumi.IDOutput) (pulumi.IDOutput, error) {
	securityGroup, err := ec2.NewSecurityGroup(ctx, "allowTls", &ec2.SecurityGroupArgs{
		Description: pulumi.String("Allow All"),
		VpcId:       vpcId,
		Ingress: ec2.SecurityGroupIngressArray{
			&ec2.SecurityGroupIngressArgs{
				Description: pulumi.String("All from VPC"),
				FromPort:    pulumi.Int(0),
				ToPort:      pulumi.Int(65535),
				Protocol:    pulumi.String("tcp"),
				CidrBlocks: pulumi.StringArray{
					pulumi.String("0.0.0.0/0"),
				},
				Ipv6CidrBlocks: pulumi.StringArray{
					pulumi.String("::/0"),
				},
			},
		},
		Egress: ec2.SecurityGroupEgressArray{
			&ec2.SecurityGroupEgressArgs{
				FromPort: pulumi.Int(0),
				ToPort:   pulumi.Int(0),
				Protocol: pulumi.String("-1"),
				CidrBlocks: pulumi.StringArray{
					pulumi.String("0.0.0.0/0"),
				},
				Ipv6CidrBlocks: pulumi.StringArray{
					pulumi.String("::/0"),
				},
			},
		},
		Tags: pulumi.StringMap{
			"Name": pulumi.String("muhammed-pulumi-security-group"),
		},
	})
	if err != nil {
		return pulumi.IDOutput{}, err
	}
	return securityGroup.ID(), nil
}
