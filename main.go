package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		vpcId, vpcErr := createVPC(ctx)
		if vpcErr != nil {
			return vpcErr
		}

		igId, igErr := createInternetGateway(ctx, vpcId)
		if igErr != nil {
			return igErr
		}

		rtId, rtErr := createRouteTable(ctx, vpcId, igId)
		if rtErr != nil {
			return rtErr
		}

		subnetId, subnetErr := createSubnet(ctx, vpcId)
		if subnetErr != nil {
			return subnetErr
		}

		associateErr := associateRouteTableAndSubnet(ctx, rtId, subnetId)
		if associateErr != nil {
			return associateErr
		}

		sgId, sgErr := createSecurityGroup(ctx, vpcId)
		if sgErr != nil {
			return sgErr
		}

		keyId, keyErr := createKeyPair(ctx)
		if keyErr != nil {
			return keyErr
		}

		_, ec2Err := createInstance(ctx, sgId, keyId)
		if ec2Err != nil {
			return ec2Err
		}

		return nil
	})
}
