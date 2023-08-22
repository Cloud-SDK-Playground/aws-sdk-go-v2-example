package eks

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"log"
)

type AddonVersions struct {
	addonInfo []types.AddonInfo
}

func GetEKSClusterVersion() []string {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		log.Fatal(err)
	}
	client := eks.NewFromConfig(cfg)
	eksAddon := &AddonVersions{}
	result := []string{}
	versionCheck := make(map[string]bool)
	pagingEKSClusterVersion(client, nil, eksAddon)
	for _, val := range eksAddon.addonInfo {
		for _, value := range val.AddonVersions {
			for _, test := range value.Compatibilities {
				if _, ok := versionCheck[*test.ClusterVersion]; !ok {
					versionCheck[*test.ClusterVersion] = true
					result = append(result, *test.ClusterVersion)
				}
			}
		}
	}
	return result
}

func pagingEKSClusterVersion(client *eks.Client, nextToken *string, eksAddon *AddonVersions) {
	resp, err := client.DescribeAddonVersions(context.TODO(), &eks.DescribeAddonVersionsInput{
		NextToken: nextToken,
	})
	if err != nil {
		return
	}
	if resp.NextToken != nil {
		pagingEKSClusterVersion(client, resp.NextToken, eksAddon)
	}
	eksAddon.addonInfo = append(eksAddon.addonInfo, resp.Addons...)
}
