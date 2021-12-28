package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createSubnet(ctx *pulumi.Context, vpcId pulumi.IDOutput) (pulumi.IDOutput, error) {
	subnet, err := ec2.NewSubnet(ctx, "main", &ec2.SubnetArgs{
		VpcId:               vpcId,
		CidrBlock:           pulumi.String("10.0.1.0/24"),
		MapPublicIpOnLaunch: pulumi.Bool(true),
		AvailabilityZone:    pulumi.String("us-west-2a"),
		Tags: pulumi.StringMap{
			"Name": pulumi.String("muhammed-pulumi-subnet"),
		},
	})
	if err != nil {
		return pulumi.IDOutput{}, err
	}
	return subnet.ID(), nil
}
