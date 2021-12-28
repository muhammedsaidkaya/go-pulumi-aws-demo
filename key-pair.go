package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createKeyPair(ctx *pulumi.Context) (pulumi.IDOutput, error) {
	fileAsset := pulumi.NewFileAsset("./keys/id_rsa.pub")
	keyPair, err := ec2.NewKeyPair(ctx, "deployer", &ec2.KeyPairArgs{
		PublicKey: pulumi.String(fileAsset.Text()),
	})
	if err != nil {
		return pulumi.IDOutput{}, err
	}
	return keyPair.ID(), nil
}
