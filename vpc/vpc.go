package vpc

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"log"
)

func GetVpc() []types.Vpc {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		log.Fatal(err)
	}
	svc := ec2.NewFromConfig(cfg)
	resp, err := svc.DescribeVpcs(context.TODO(), &ec2.DescribeVpcsInput{})
	if err != nil {
		log.Fatal(err)
	}
	return resp.Vpcs
}

func GetVpcSubnets() []types.Subnet {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		log.Fatal(err)
	}
	svc := ec2.NewFromConfig(cfg)
	filterType := "availability-zone"
	zoneName := "us-west-2a"
	resp, err := svc.DescribeSubnets(context.TODO(), &ec2.DescribeSubnetsInput{
		Filters: []types.Filter{
			{
				Name:   &filterType,
				Values: []string{zoneName},
			},
		},
		MaxResults: nil,
		NextToken:  nil,
		SubnetIds:  nil,
	})
	return resp.Subnets
}
