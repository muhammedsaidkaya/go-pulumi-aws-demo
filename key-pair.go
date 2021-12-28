package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"os"
)

func createKeyPair(ctx *pulumi.Context) (pulumi.IDOutput, error) {
	//fileAsset := pulumi.NewFileAsset("./keys/id_rsa.pub")
	dat, fileErr := os.ReadFile("./keys/id_rsa.pub")
	if fileErr != nil {
		return pulumi.IDOutput{}, fileErr
	}
	keyPair, err := ec2.NewKeyPair(ctx, "deployer", &ec2.KeyPairArgs{
		PublicKey: pulumi.String(string(dat)),
	})
	if err != nil {
		return pulumi.IDOutput{}, err
	}
	return keyPair.ID(), nil
}
