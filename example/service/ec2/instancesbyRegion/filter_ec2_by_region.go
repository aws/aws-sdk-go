package main

import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
    "github.com/aws/aws-sdk-go/aws"
    "os"
)

func fetchRegion()[]string{

	regions := []string{}  
	sess1 := session.Must(session.NewSession(&aws.Config{
			Region: aws.String("us-west-1"),
	}))

	svc := ec2.New(sess1)
        awsregions, err := svc.DescribeRegions(&ec2.DescribeRegionsInput{})
 	if err != nil {
        	fmt.Println("Error", err)
    	}  
        for _, region := range awsregions.Regions {
                 regions=append(regions,*region.RegionName) 
         }

 	return regions
}

func main() {

	States:= []*string{}
	if len(os.Args)>1  {
		for i := 1; i < len(os.Args); i++ {
			States=append(States,&os.Args[i])
			
		}	
	} else {
		States= []*string{
				 aws.String("running"),
				 aws.String("pending"),
				 aws.String("stopped"),
				 }
	       }					
    	  

        regions:=fetchRegion()
	for _, region := range regions {
		sess := session.Must(session.NewSession(&aws.Config{
			Region: aws.String(region),
		}))
   	 	
   		ec2Svc := ec2.New(sess)
		params := &ec2.DescribeInstancesInput{
        	Filters: []*ec2.Filter{
            		&ec2.Filter{
                		Name: aws.String("instance-state-name"),
                		Values: States,
           		 },
        		},
    		}
    
    	result, err := ec2Svc.DescribeInstances(params)
    	if err != nil {
        	fmt.Println("Error", err)
    	} else {
		fmt.Println("\n\n\nFetching instace details  for region: ", region )
              	if len(result.Reservations)==0{
			fmt.Println("There is no instance for the for region: ", region )
                }
	      	for _, reservation := range result.Reservations {
	   	  
		    fmt.Println("printing instance details.....")
		    for _, instance := range reservation.Instances {
			fmt.Println("instance id "+*instance.InstanceId)
			fmt.Println("current State "+*instance.State.Name)
		    }		
	        }
	       fmt.Printf("done for region %s **** \n", region )     
          } 
     }
 }

