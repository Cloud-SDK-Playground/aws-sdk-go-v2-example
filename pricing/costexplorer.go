package pricing

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
)

func GetCostUsage() *costexplorer.GetCostAndUsageOutput {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	svc := costexplorer.NewFromConfig(cfg)
	resp, err := svc.GetCostAndUsage(context.TODO(), &costexplorer.GetCostAndUsageInput{
		// Filter:      getFilter(),
		Granularity: "DAILY",
		TimePeriod:  getPeriod(),
		Metrics:     []string{"BlendedCost"},
		GroupBy: []types.GroupDefinition{
			{
				Type: "DIMENSION",
				Key:  aws.String("SERVICE"),
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func getPeriod() *types.DateInterval {
	now := time.Now()
	then := now.AddDate(0, 0, -7)
	dateRange := &types.DateInterval{
		End:   aws.String(now.Format("2006-01-02")),
		Start: aws.String(then.Format("2006-01-02")),
	}
	return dateRange
}

func getFilter() *types.Expression {
	expression := &types.Expression{
		Dimensions: &types.DimensionValues{
			Key:    "REGION",
			Values: []string{"us-east-1"},
		},
	}
	return expression
}
