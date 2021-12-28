package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createVPC(ctx *pulumi.Context) (pulumi.IDOutput, error) {
	vpc, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
		CidrBlock:          pulumi.String("10.0.0.0/16"),
		EnableDnsHostnames: pulumi.Bool(true),
		Tags: pulumi.StringMap{
			"Name": pulumi.String("muhammed-pulumi-vpc"),
		},
	})
	if err != nil {
		return pulumi.IDOutput{}, err
	}
	return vpc.ID(), nil
}
