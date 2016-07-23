package main

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"strings"
	"os"
)

var AwsRegions = []string{"us-east-1", "us-west-1", "us-west-2", "eu-west-1", "eu-central-1", "sa-east-1"}

func listFilteredInstances(nameFilter string) []ec2.DescribeInstancesOutput{
	filteredList := []ec2.DescribeInstancesOutput{}
	for _, region := range AwsRegions {
		svc := ec2.New(session.New(&aws.Config{Region: aws.String(region)}))
		fmt.Println("listing VPNs in:", region)
		params := &ec2.DescribeInstancesInput{
			Filters: []*ec2.Filter{
				{
					Name: aws.String("tag:Name"),
					Values: []*string{
						aws.String(strings.Join([]string{"*",nameFilter,"*"},"")),
					},
				},
			},
		}
		resp, err := svc.DescribeInstances(params)
		if err != nil {
			fmt.Println("there was an error listing instnaces in", region, err.Error())
			log.Fatal(err.Error())
		}
		filteredList = append(filteredList,*resp)
	}
	return filteredList
}

func main() {
	filteredInstances := listFilteredInstances(os.Args[1])
	fmt.Printf("%+v\n",filteredInstances)
}
