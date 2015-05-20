package ecs_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/ecs"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleECS_CreateCluster() {
	svc := ecs.New(nil)

	params := &ecs.CreateClusterInput{
		ClusterName: aws.String("String"),
	}
	resp, err := svc.CreateCluster(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleECS_CreateService() {
	svc := ecs.New(nil)

	params := &ecs.CreateServiceInput{
		ServiceName:  aws.String("String"), // Required
		ClientToken:  aws.String("String"),
		Cluster:      aws.String("String"),
		DesiredCount: aws.Long(1),
		LoadBalancers: []*ecs.LoadBalancer{
			&ecs.LoadBalancer{ // Required
				ContainerName:    aws.String("String"),
				ContainerPort:    aws.Long(1),
				LoadBalancerName: aws.String("String"),
			},
			// More values...
		},
		Role:           aws.String("String"),
		TaskDefinition: aws.String("String"),
	}
	resp, err := svc.CreateService(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleECS_DeleteCluster() {
	svc := ecs.New(nil)

	params := &ecs.DeleteClusterInput{
		Cluster: aws.String("String"), // Required
	}
	resp, err := svc.DeleteCluster(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleECS_DeleteService() {
	svc := ecs.New(nil)

	params := &ecs.DeleteServiceInput{
		Service: aws.String("String"), // Required
		Cluster: aws.String("String"),
	}
	resp, err := svc.DeleteService(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleECS_DeregisterContainerInstance() {
	svc := ecs.New(nil)

	params := &ecs.DeregisterContainerInstanceInput{
		ContainerInstance: aws.String("String"), // Required
		Cluster:           aws.String("String"),
		Force:             aws.Boolean(true),
	}
	resp, err := svc.DeregisterContainerInstance(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleECS_DeregisterTaskDefinition() {
	svc := ecs.New(nil)

	params := &ecs.DeregisterTaskDefinitionInput{
		TaskDefinition: aws.String("String"), // Required
	}
	resp, err := svc.DeregisterTaskDefinition(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleECS_DescribeClusters() {
	svc := ecs.New(nil)

	params := &ecs.DescribeClustersInput{
		Clusters: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeClusters(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleECS_DescribeContainerInstances() {
	svc := ecs.New(nil)

	params := &ecs.DescribeContainerInstancesInput{
		ContainerInstances: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
		Cluster: aws.String("String"),
	}
	resp, err := svc.DescribeContainerInstances(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleECS_DescribeServices() {
	svc := ecs.New(nil)

	params := &ecs.DescribeServicesInput{
		Services: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
		Cluster: aws.String("String"),
	}
	resp, err := svc.DescribeServices(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleECS_DescribeTaskDefinition() {
	svc := ecs.New(nil)

	params := &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String("String"), // Required
	}
	resp, err := svc.DescribeTaskDefinition(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleECS_DescribeTasks() {
	svc := ecs.New(nil)

	params := &ecs.DescribeTasksInput{
		Tasks: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
		Cluster: aws.String("String"),
	}
	resp, err := svc.DescribeTasks(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleECS_DiscoverPollEndpoint() {
	svc := ecs.New(nil)

	params := &ecs.DiscoverPollEndpointInput{
		Cluster:           aws.String("String"),
		ContainerInstance: aws.String("String"),
	}
	resp, err := svc.DiscoverPollEndpoint(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleECS_ListClusters() {
	svc := ecs.New(nil)

	params := &ecs.ListClustersInput{
		MaxResults: aws.Long(1),
		NextToken:  aws.String("String"),
	}
	resp, err := svc.ListClusters(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleECS_ListContainerInstances() {
	svc := ecs.New(nil)

	params := &ecs.ListContainerInstancesInput{
		Cluster:    aws.String("String"),
		MaxResults: aws.Long(1),
		NextToken:  aws.String("String"),
	}
	resp, err := svc.ListContainerInstances(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleECS_ListServices() {
	svc := ecs.New(nil)

	params := &ecs.ListServicesInput{
		Cluster:    aws.String("String"),
		MaxResults: aws.Long(1),
		NextToken:  aws.String("String"),
	}
	resp, err := svc.ListServices(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleECS_ListTaskDefinitionFamilies() {
	svc := ecs.New(nil)

	params := &ecs.ListTaskDefinitionFamiliesInput{
		FamilyPrefix: aws.String("String"),
		MaxResults:   aws.Long(1),
		NextToken:    aws.String("String"),
	}
	resp, err := svc.ListTaskDefinitionFamilies(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleECS_ListTaskDefinitions() {
	svc := ecs.New(nil)

	params := &ecs.ListTaskDefinitionsInput{
		FamilyPrefix: aws.String("String"),
		MaxResults:   aws.Long(1),
		NextToken:    aws.String("String"),
	}
	resp, err := svc.ListTaskDefinitions(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleECS_ListTasks() {
	svc := ecs.New(nil)

	params := &ecs.ListTasksInput{
		Cluster:           aws.String("String"),
		ContainerInstance: aws.String("String"),
		Family:            aws.String("String"),
		MaxResults:        aws.Long(1),
		NextToken:         aws.String("String"),
		ServiceName:       aws.String("String"),
		StartedBy:         aws.String("String"),
	}
	resp, err := svc.ListTasks(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleECS_RegisterContainerInstance() {
	svc := ecs.New(nil)

	params := &ecs.RegisterContainerInstanceInput{
		Cluster:                           aws.String("String"),
		InstanceIdentityDocument:          aws.String("String"),
		InstanceIdentityDocumentSignature: aws.String("String"),
		TotalResources: []*ecs.Resource{
			&ecs.Resource{ // Required
				DoubleValue:  aws.Double(1.0),
				IntegerValue: aws.Long(1),
				LongValue:    aws.Long(1),
				Name:         aws.String("String"),
				StringSetValue: []*string{
					aws.String("String"), // Required
					// More values...
				},
				Type: aws.String("String"),
			},
			// More values...
		},
		VersionInfo: &ecs.VersionInfo{
			AgentHash:     aws.String("String"),
			AgentVersion:  aws.String("String"),
			DockerVersion: aws.String("String"),
		},
	}
	resp, err := svc.RegisterContainerInstance(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleECS_RegisterTaskDefinition() {
	svc := ecs.New(nil)

	params := &ecs.RegisterTaskDefinitionInput{
		ContainerDefinitions: []*ecs.ContainerDefinition{ // Required
			&ecs.ContainerDefinition{ // Required
				CPU: aws.Long(1),
				Command: []*string{
					aws.String("String"), // Required
					// More values...
				},
				EntryPoint: []*string{
					aws.String("String"), // Required
					// More values...
				},
				Environment: []*ecs.KeyValuePair{
					&ecs.KeyValuePair{ // Required
						Name:  aws.String("String"),
						Value: aws.String("String"),
					},
					// More values...
				},
				Essential: aws.Boolean(true),
				Image:     aws.String("String"),
				Links: []*string{
					aws.String("String"), // Required
					// More values...
				},
				Memory: aws.Long(1),
				MountPoints: []*ecs.MountPoint{
					&ecs.MountPoint{ // Required
						ContainerPath: aws.String("String"),
						ReadOnly:      aws.Boolean(true),
						SourceVolume:  aws.String("String"),
					},
					// More values...
				},
				Name: aws.String("String"),
				PortMappings: []*ecs.PortMapping{
					&ecs.PortMapping{ // Required
						ContainerPort: aws.Long(1),
						HostPort:      aws.Long(1),
					},
					// More values...
				},
				VolumesFrom: []*ecs.VolumeFrom{
					&ecs.VolumeFrom{ // Required
						ReadOnly:        aws.Boolean(true),
						SourceContainer: aws.String("String"),
					},
					// More values...
				},
			},
			// More values...
		},
		Family: aws.String("String"), // Required
		Volumes: []*ecs.Volume{
			&ecs.Volume{ // Required
				Host: &ecs.HostVolumeProperties{
					SourcePath: aws.String("String"),
				},
				Name: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.RegisterTaskDefinition(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleECS_RunTask() {
	svc := ecs.New(nil)

	params := &ecs.RunTaskInput{
		TaskDefinition: aws.String("String"), // Required
		Cluster:        aws.String("String"),
		Count:          aws.Long(1),
		Overrides: &ecs.TaskOverride{
			ContainerOverrides: []*ecs.ContainerOverride{
				&ecs.ContainerOverride{ // Required
					Command: []*string{
						aws.String("String"), // Required
						// More values...
					},
					Name: aws.String("String"),
				},
				// More values...
			},
		},
		StartedBy: aws.String("String"),
	}
	resp, err := svc.RunTask(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleECS_StartTask() {
	svc := ecs.New(nil)

	params := &ecs.StartTaskInput{
		ContainerInstances: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
		TaskDefinition: aws.String("String"), // Required
		Cluster:        aws.String("String"),
		Overrides: &ecs.TaskOverride{
			ContainerOverrides: []*ecs.ContainerOverride{
				&ecs.ContainerOverride{ // Required
					Command: []*string{
						aws.String("String"), // Required
						// More values...
					},
					Name: aws.String("String"),
				},
				// More values...
			},
		},
		StartedBy: aws.String("String"),
	}
	resp, err := svc.StartTask(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleECS_StopTask() {
	svc := ecs.New(nil)

	params := &ecs.StopTaskInput{
		Task:    aws.String("String"), // Required
		Cluster: aws.String("String"),
	}
	resp, err := svc.StopTask(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleECS_SubmitContainerStateChange() {
	svc := ecs.New(nil)

	params := &ecs.SubmitContainerStateChangeInput{
		Cluster:       aws.String("String"),
		ContainerName: aws.String("String"),
		ExitCode:      aws.Long(1),
		NetworkBindings: []*ecs.NetworkBinding{
			&ecs.NetworkBinding{ // Required
				BindIP:        aws.String("String"),
				ContainerPort: aws.Long(1),
				HostPort:      aws.Long(1),
			},
			// More values...
		},
		Reason: aws.String("String"),
		Status: aws.String("String"),
		Task:   aws.String("String"),
	}
	resp, err := svc.SubmitContainerStateChange(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleECS_SubmitTaskStateChange() {
	svc := ecs.New(nil)

	params := &ecs.SubmitTaskStateChangeInput{
		Cluster: aws.String("String"),
		Reason:  aws.String("String"),
		Status:  aws.String("String"),
		Task:    aws.String("String"),
	}
	resp, err := svc.SubmitTaskStateChange(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleECS_UpdateService() {
	svc := ecs.New(nil)

	params := &ecs.UpdateServiceInput{
		Service:        aws.String("String"), // Required
		Cluster:        aws.String("String"),
		DesiredCount:   aws.Long(1),
		TaskDefinition: aws.String("String"),
	}
	resp, err := svc.UpdateService(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}