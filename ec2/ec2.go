package ec2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"log"
)

type EC2DataTypes struct {
	instanceTypes []types.InstanceTypeOffering
}
type EC2Images struct {
	ami []types.Image
}

func GetEC2InstanceTypes() []types.InstanceTypeOffering {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		log.Fatal(err)
	}
	client := ec2.NewFromConfig(cfg)
	ec2InstanceTypes := &EC2DataTypes{}
	pagingInstanceTypes(client, nil, ec2InstanceTypes)
	return ec2InstanceTypes.instanceTypes
}

func pagingInstanceTypes(client *ec2.Client, nextToken *string, ec2InstanceTypes *EC2DataTypes) {
	//filterType := "instance-type"
	//instanceName := "t1.micro"
	resp, err := client.DescribeInstanceTypeOfferings(context.TODO(), &ec2.DescribeInstanceTypeOfferingsInput{
		NextToken: nextToken,
		//Filters: []types.Filter{
		//	{
		//		Name:   &filterType,
		//		Values: []string{instanceName},
		//	},
		//},
	})
	if err != nil {
		return
	}
	if resp.NextToken != nil {
		pagingInstanceTypes(client, resp.NextToken, ec2InstanceTypes)
	}
	ec2InstanceTypes.instanceTypes = append(ec2InstanceTypes.instanceTypes, resp.InstanceTypeOfferings...)
}

func GetEC2AMI() []types.Image {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		log.Fatal(err)
	}
	client := ec2.NewFromConfig(cfg)
	ec2Images := &EC2Images{}
	pagingAMI(client, nil, ec2Images)
	return ec2Images.ami
}
func pagingAMI(client *ec2.Client, nextToken *string, ec2Images *EC2Images) {
	filterType := "owner-alias"
	publicFilterType := "is-public"
	ownerName := "amazon"
	resp, err := client.DescribeImages(context.TODO(), &ec2.DescribeImagesInput{
		NextToken: nextToken,
		Filters: []types.Filter{
			{
				Name:   &filterType,
				Values: []string{ownerName},
			},
			{
				Name:   &publicFilterType,
				Values: []string{"true"},
			},
		},
	})
	if err != nil {
		return
	}
	ec2Images.ami = append(ec2Images.ami, resp.Images...)
}
