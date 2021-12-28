package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createRouteTable(ctx *pulumi.Context, vpcId pulumi.IDOutput, igId pulumi.IDOutput) (pulumi.IDOutput, error) {
	rt, err := ec2.NewRouteTable(ctx, "example", &ec2.RouteTableArgs{
		VpcId: vpcId,
		Routes: ec2.RouteTableRouteArray{
			&ec2.RouteTableRouteArgs{
				CidrBlock: pulumi.String("0.0.0.0/0"),
				GatewayId: igId,
			},
			&ec2.RouteTableRouteArgs{
				Ipv6CidrBlock: pulumi.String("::/0"),
				GatewayId:     igId,
			},
		},
		Tags: pulumi.StringMap{
			"Name": pulumi.String("muhammed-pulumi-route-table"),
		},
	})
	if err != nil {
		return pulumi.IDOutput{}, err
	}
	return rt.ID(), nil
}
