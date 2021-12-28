package main

//func createNetworkInterface(ctx *pulumi.Context, subnetId pulumi.IDOutput) (pulumi.IDOutput, error) {
//	_, err := ec2.NewNetworkInterface(ctx, "test", &ec2.NetworkInterfaceArgs{
//		SubnetId: pulumi.Any(aws_subnet.Public_a.Id),
//		PrivateIps: pulumi.StringArray{
//			pulumi.String("10.0.0.50"),
//		},
//		SecurityGroups: pulumi.StringArray{
//			pulumi.Any(aws_security_group.Web.Id),
//		},
//		Attachments: ec2.NetworkInterfaceAttachmentArray{
//			&ec2.NetworkInterfaceAttachmentArgs{
//				Instance:    pulumi.Any(aws_instance.Test.Id),
//				DeviceIndex: pulumi.Int(1),
//			},
//		},
//		Tags: pulumi.StringMap{
//			"Name": pulumi.String("muhammed-pulumi-network-interface"),
//		},
//	})
//	if err != nil {
//		return err
//	}
//	return nil
//}
