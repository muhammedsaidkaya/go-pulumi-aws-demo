package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func associateRouteTableAndSubnet(ctx *pulumi.Context, routeTableId pulumi.IDOutput, subnetId pulumi.IDOutput) error {
	_, err := ec2.NewRouteTableAssociation(ctx, "routeTableAssociation", &ec2.RouteTableAssociationArgs{
		SubnetId:     subnetId,
		RouteTableId: routeTableId,
	})
	if err != nil {
		return err
	}
	return nil
}
