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

var awsRegions = []string{"us-east-1"}

func listFilteredInstances(nameFilter string) {
	for _, region := range awsRegions {
		svc := ec2.New(session.New(&aws.Config{Region: aws.String(region)}))
		fmt.Printf("listing instances with tag %v in: %v\n", nameFilter,region)
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
		fmt.Printf("%+v\n",*resp)
	}
}

func main() {
	listFilteredInstances(os.Args[1])
}


