package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createInstance(ctx *pulumi.Context, subnetId pulumi.IDOutput, sgId pulumi.IDOutput, keyId pulumi.IDOutput) (pulumi.IDOutput, error) {
	instance, err := ec2.NewInstance(ctx, "web", &ec2.InstanceArgs{
		Ami: pulumi.String("ami-0686851c4e7b1a8e1"),
		//ami-0f81e6e71078b75b6 ubuntu 20.04
		//ami-0686851c4e7b1a8e1 centos-7
		InstanceType:             pulumi.String("t3.large"),
		AssociatePublicIpAddress: pulumi.Bool(true),
		VpcSecurityGroupIds: pulumi.StringArray{
			sgId,
		},
		KeyName:  keyId,
		SubnetId: subnetId,
		Tags: pulumi.StringMap{
			"Name":          pulumi.String("muhammed-pulumi-instance"),
			"Desired-State": pulumi.String("Ignore"),
		},
	})
	if err != nil {
		return pulumi.IDOutput{}, err
	}
	return instance.ID(), nil
}
