// +build integration

package cloudwatch_test

import (
	"testing"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/cloudwatch"
)

func TestPutAndGetMetricData(t *testing.T) {

	cwClient := cloudwatch.New(aws.DefaultCreds(), "us-east-1", nil)

	putInput := cloudwatch.PutMetricDataInput{
		Namespace: aws.String("aws-go-sdk-test"),
		MetricData: []cloudwatch.MetricDatum{
			cloudwatch.MetricDatum{
				MetricName: aws.String("aws-go-sdk-test"),
				Unit:       aws.String("Seconds"),
				Value:      aws.Double(1.1),
				Timestamp:  time.Now(),
			},
		},
	}

	err := cwClient.PutMetricData(&putInput)

	if err != nil {
		t.Fatal("Failed to put metric data: ", err)
	}

	time.Sleep(time.Second * 2)

	getMetric, err := cwClient.GetMetricStatistics(&cloudwatch.GetMetricStatisticsInput{
		Namespace:  aws.String("aws-go-sdk-test"),
		MetricName: aws.String("aws-go-sdk-test"),
		StartTime:  time.Now().AddDate(0, 0, -2),
		Period:     aws.Integer(60 * 60),
		Statistics: []string{"Maximum"},
		EndTime:    time.Now().AddDate(0, 0, 2),
	})
	if err != nil {
		t.Fatal("Failed to get metric data: ", err)
	}

	found := false
	for _, dp := range getMetric.Datapoints {
		if *(dp.Maximum) == 1.1 {
			found = true
			break
		}
	}

	if !found {
		t.Fatal("Failed to find metric data")
	}
}
