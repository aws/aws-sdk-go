package mediastoredata_test

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/awstesting/unit"
	"github.com/aws/aws-sdk-go/service/mediastore"
	"github.com/aws/aws-sdk-go/service/mediastoredata"
)

func ExampleMediaStoreData_DescribeEndpoint_shared00() {
	const containerName = "awsgosdkteamintegcontainer"

	sess := unit.Session
	if v := aws.StringValue(sess.Config.Region); len(v) == 0 {
		sess.Config.Region = aws.String("us-east-1")
	}

	ctrlSvc := mediastore.New(sess)
	descResp, err := ctrlSvc.DescribeContainer(&mediastore.DescribeContainerInput{
		ContainerName: aws.String(containerName),
	})
	if err != nil {
		log.Fatalf("failed to get mediastore container endpoint, %v", err)
	}

	dataSvc := mediastoredata.New(sess, &aws.Config{
		Endpoint: descResp.Container.Endpoint,
	})
	_, err = dataSvc.ListItems(&mediastoredata.ListItemsInput{})
	if err != nil {
		log.Fatalf("failed to make medaistoredata API call, %v", err)
	}
}
