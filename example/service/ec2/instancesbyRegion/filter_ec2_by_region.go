package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func fetchRegion() []string {

	var regions []string
	awsSession := session.Must(session.NewSession(&aws.Config{}))

	svc := ec2.New(awsSession)
	awsRegions, err := svc.DescribeRegions(&ec2.DescribeRegionsInput{})
	if err != nil {
		fmt.Println("Error", err)
	}
	for _, region := range awsRegions.Regions {
		regions = append(regions, *region.RegionName)
	}

	return regions
}

func (i *commandLineArumentType) String() string {
	return "my string representation"
}

func (i *commandLineArumentType) Set(value string) error {
	*i = append(*i, value)
	return nil
}

type commandLineArumentType []string

func parseArguments() ([]*string, []string) {
	var StateArguments, RegionArguments commandLineArumentType
	States := []*string{}
	Regions := []string{}
	flag.Var(&StateArguments, "state", "state list")
	flag.Var(&RegionArguments, "region", "region list")
	flag.Parse()
	if flag.NFlag() != 0 {
		for i := 0; i < len(StateArguments); i++ {
			States = append(States, &StateArguments[i])
		}

		for i := 0; i < len(RegionArguments); i++ {
			Regions = append(Regions, RegionArguments[i])
		}
	}

	return States, Regions
}

func usage() string {

	message := "\n\nMissing mandatory flag 'state'. Please use like below  Example :\n\n"
	message = message + "To fetch the stopped instance of all region use below:\n"
	message = message + "\t./filter_ec2_by_region --state running --state stopped \n"
	message = message + "To fetch the stopped and running instance  for  region us-west-1 and eu-west-1 use below:\n"
	message = message + "\t./filter_ec2_by_region --state running --state stopped --region us-west-1 --region=eu-west-1\n"
	return message

}
func main() {

	Regions := []string{}
	States := []*string{}
	States, Regions = parseArguments()
	if len(States) == 0 {
		fmt.Fprintf(os.Stderr, "error: %v\n", usage())
		os.Exit(1)
	}
	InstanceCriteria := " "
	for _, State := range States {
		InstanceCriteria = InstanceCriteria + "[" + *State + "]"
	}

	if len(Regions) == 0 {
		Regions = fetchRegion()
	}

	for _, region := range Regions {
		sess := session.Must(session.NewSession(&aws.Config{
			Region: aws.String(region),
		}))

		ec2Svc := ec2.New(sess)
		params := &ec2.DescribeInstancesInput{
			Filters: []*ec2.Filter{
				&ec2.Filter{
					Name:   aws.String("instance-state-name"),
					Values: States,
				},
			},
		}

		result, err := ec2Svc.DescribeInstances(params)
		if err != nil {
			fmt.Println("Error", err)
		} else {
			fmt.Printf("\n\n\nFetching instace details  for region: %s with criteria: %s**\n ", region, InstanceCriteria)
			if len(result.Reservations) == 0 {
				fmt.Printf("There is no instance for the for region %s with the matching Criteria:%s  \n", region, InstanceCriteria)
			}
			for _, reservation := range result.Reservations {

				fmt.Println("printing instance details.....")
				for _, instance := range reservation.Instances {
					fmt.Println("instance id " + *instance.InstanceId)
					fmt.Println("current State " + *instance.State.Name)
				}
			}
			fmt.Printf("done for region %s **** \n", region)
		}
	}
}
