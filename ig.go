package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createInternetGateway(ctx *pulumi.Context, vpcId pulumi.IDOutput) (pulumi.IDOutput, error) {
	ig, err := ec2.NewInternetGateway(ctx, "gw", &ec2.InternetGatewayArgs{
		VpcId: vpcId,
		Tags: pulumi.StringMap{
			"Name": pulumi.String("muhammed-pulumi-internet-gateway"),
		},
	})

	if err != nil {
		return pulumi.IDOutput{}, err
	}
	return ig.ID(), nil
}
