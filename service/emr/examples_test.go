package emr_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/emr"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleEMR_AddInstanceGroups() {
	svc := emr.New(nil)

	params := &emr.AddInstanceGroupsInput{
		InstanceGroups: []*emr.InstanceGroupConfig{ // Required
			&emr.InstanceGroupConfig{ // Required
				InstanceCount: aws.Long(1),                    // Required
				InstanceRole:  aws.String("InstanceRoleType"), // Required
				InstanceType:  aws.String("InstanceType"),     // Required
				BidPrice:      aws.String("XmlStringMaxLen256"),
				Market:        aws.String("MarketType"),
				Name:          aws.String("XmlStringMaxLen256"),
			},
			// More values...
		},
		JobFlowID: aws.String("XmlStringMaxLen256"), // Required
	}
	resp, err := svc.AddInstanceGroups(params)

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

func ExampleEMR_AddJobFlowSteps() {
	svc := emr.New(nil)

	params := &emr.AddJobFlowStepsInput{
		JobFlowID: aws.String("XmlStringMaxLen256"), // Required
		Steps: []*emr.StepConfig{ // Required
			&emr.StepConfig{ // Required
				HadoopJARStep: &emr.HadoopJARStepConfig{ // Required
					JAR: aws.String("XmlString"), // Required
					Args: []*string{
						aws.String("XmlString"), // Required
						// More values...
					},
					MainClass: aws.String("XmlString"),
					Properties: []*emr.KeyValue{
						&emr.KeyValue{ // Required
							Key:   aws.String("XmlString"),
							Value: aws.String("XmlString"),
						},
						// More values...
					},
				},
				Name:            aws.String("XmlStringMaxLen256"), // Required
				ActionOnFailure: aws.String("ActionOnFailure"),
			},
			// More values...
		},
	}
	resp, err := svc.AddJobFlowSteps(params)

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

func ExampleEMR_AddTags() {
	svc := emr.New(nil)

	params := &emr.AddTagsInput{
		ResourceID: aws.String("ResourceId"), // Required
		Tags: []*emr.Tag{ // Required
			&emr.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.AddTags(params)

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

func ExampleEMR_DescribeCluster() {
	svc := emr.New(nil)

	params := &emr.DescribeClusterInput{
		ClusterID: aws.String("ClusterId"), // Required
	}
	resp, err := svc.DescribeCluster(params)

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

func ExampleEMR_DescribeJobFlows() {
	svc := emr.New(nil)

	params := &emr.DescribeJobFlowsInput{
		CreatedAfter:  aws.Time(time.Now()),
		CreatedBefore: aws.Time(time.Now()),
		JobFlowIDs: []*string{
			aws.String("XmlString"), // Required
			// More values...
		},
		JobFlowStates: []*string{
			aws.String("JobFlowExecutionState"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeJobFlows(params)

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

func ExampleEMR_DescribeStep() {
	svc := emr.New(nil)

	params := &emr.DescribeStepInput{
		ClusterID: aws.String("ClusterId"), // Required
		StepID:    aws.String("StepId"),    // Required
	}
	resp, err := svc.DescribeStep(params)

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

func ExampleEMR_ListBootstrapActions() {
	svc := emr.New(nil)

	params := &emr.ListBootstrapActionsInput{
		ClusterID: aws.String("ClusterId"), // Required
		Marker:    aws.String("Marker"),
	}
	resp, err := svc.ListBootstrapActions(params)

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

func ExampleEMR_ListClusters() {
	svc := emr.New(nil)

	params := &emr.ListClustersInput{
		ClusterStates: []*string{
			aws.String("ClusterState"), // Required
			// More values...
		},
		CreatedAfter:  aws.Time(time.Now()),
		CreatedBefore: aws.Time(time.Now()),
		Marker:        aws.String("Marker"),
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

func ExampleEMR_ListInstanceGroups() {
	svc := emr.New(nil)

	params := &emr.ListInstanceGroupsInput{
		ClusterID: aws.String("ClusterId"), // Required
		Marker:    aws.String("Marker"),
	}
	resp, err := svc.ListInstanceGroups(params)

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

func ExampleEMR_ListInstances() {
	svc := emr.New(nil)

	params := &emr.ListInstancesInput{
		ClusterID:       aws.String("ClusterId"), // Required
		InstanceGroupID: aws.String("InstanceGroupId"),
		InstanceGroupTypes: []*string{
			aws.String("InstanceGroupType"), // Required
			// More values...
		},
		Marker: aws.String("Marker"),
	}
	resp, err := svc.ListInstances(params)

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

func ExampleEMR_ListSteps() {
	svc := emr.New(nil)

	params := &emr.ListStepsInput{
		ClusterID: aws.String("ClusterId"), // Required
		Marker:    aws.String("Marker"),
		StepIDs: []*string{
			aws.String("XmlString"), // Required
			// More values...
		},
		StepStates: []*string{
			aws.String("StepState"), // Required
			// More values...
		},
	}
	resp, err := svc.ListSteps(params)

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

func ExampleEMR_ModifyInstanceGroups() {
	svc := emr.New(nil)

	params := &emr.ModifyInstanceGroupsInput{
		InstanceGroups: []*emr.InstanceGroupModifyConfig{
			&emr.InstanceGroupModifyConfig{ // Required
				InstanceGroupID: aws.String("XmlStringMaxLen256"), // Required
				EC2InstanceIDsToTerminate: []*string{
					aws.String("InstanceId"), // Required
					// More values...
				},
				InstanceCount: aws.Long(1),
			},
			// More values...
		},
	}
	resp, err := svc.ModifyInstanceGroups(params)

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

func ExampleEMR_RemoveTags() {
	svc := emr.New(nil)

	params := &emr.RemoveTagsInput{
		ResourceID: aws.String("ResourceId"), // Required
		TagKeys: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.RemoveTags(params)

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

func ExampleEMR_RunJobFlow() {
	svc := emr.New(nil)

	params := &emr.RunJobFlowInput{
		Instances: &emr.JobFlowInstancesConfig{ // Required
			AdditionalMasterSecurityGroups: []*string{
				aws.String("XmlStringMaxLen256"), // Required
				// More values...
			},
			AdditionalSlaveSecurityGroups: []*string{
				aws.String("XmlStringMaxLen256"), // Required
				// More values...
			},
			EC2KeyName:                    aws.String("XmlStringMaxLen256"),
			EC2SubnetID:                   aws.String("XmlStringMaxLen256"),
			EMRManagedMasterSecurityGroup: aws.String("XmlStringMaxLen256"),
			EMRManagedSlaveSecurityGroup:  aws.String("XmlStringMaxLen256"),
			HadoopVersion:                 aws.String("XmlStringMaxLen256"),
			InstanceCount:                 aws.Long(1),
			InstanceGroups: []*emr.InstanceGroupConfig{
				&emr.InstanceGroupConfig{ // Required
					InstanceCount: aws.Long(1),                    // Required
					InstanceRole:  aws.String("InstanceRoleType"), // Required
					InstanceType:  aws.String("InstanceType"),     // Required
					BidPrice:      aws.String("XmlStringMaxLen256"),
					Market:        aws.String("MarketType"),
					Name:          aws.String("XmlStringMaxLen256"),
				},
				// More values...
			},
			KeepJobFlowAliveWhenNoSteps: aws.Boolean(true),
			MasterInstanceType:          aws.String("InstanceType"),
			Placement: &emr.PlacementType{
				AvailabilityZone: aws.String("XmlString"), // Required
			},
			SlaveInstanceType:    aws.String("InstanceType"),
			TerminationProtected: aws.Boolean(true),
		},
		Name:           aws.String("XmlStringMaxLen256"), // Required
		AMIVersion:     aws.String("XmlStringMaxLen256"),
		AdditionalInfo: aws.String("XmlString"),
		BootstrapActions: []*emr.BootstrapActionConfig{
			&emr.BootstrapActionConfig{ // Required
				Name: aws.String("XmlStringMaxLen256"), // Required
				ScriptBootstrapAction: &emr.ScriptBootstrapActionConfig{ // Required
					Path: aws.String("XmlString"), // Required
					Args: []*string{
						aws.String("XmlString"), // Required
						// More values...
					},
				},
			},
			// More values...
		},
		JobFlowRole: aws.String("XmlString"),
		LogURI:      aws.String("XmlString"),
		NewSupportedProducts: []*emr.SupportedProductConfig{
			&emr.SupportedProductConfig{ // Required
				Args: []*string{
					aws.String("XmlString"), // Required
					// More values...
				},
				Name: aws.String("XmlStringMaxLen256"),
			},
			// More values...
		},
		ServiceRole: aws.String("XmlString"),
		Steps: []*emr.StepConfig{
			&emr.StepConfig{ // Required
				HadoopJARStep: &emr.HadoopJARStepConfig{ // Required
					JAR: aws.String("XmlString"), // Required
					Args: []*string{
						aws.String("XmlString"), // Required
						// More values...
					},
					MainClass: aws.String("XmlString"),
					Properties: []*emr.KeyValue{
						&emr.KeyValue{ // Required
							Key:   aws.String("XmlString"),
							Value: aws.String("XmlString"),
						},
						// More values...
					},
				},
				Name:            aws.String("XmlStringMaxLen256"), // Required
				ActionOnFailure: aws.String("ActionOnFailure"),
			},
			// More values...
		},
		SupportedProducts: []*string{
			aws.String("XmlStringMaxLen256"), // Required
			// More values...
		},
		Tags: []*emr.Tag{
			&emr.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
		VisibleToAllUsers: aws.Boolean(true),
	}
	resp, err := svc.RunJobFlow(params)

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

func ExampleEMR_SetTerminationProtection() {
	svc := emr.New(nil)

	params := &emr.SetTerminationProtectionInput{
		JobFlowIDs: []*string{ // Required
			aws.String("XmlString"), // Required
			// More values...
		},
		TerminationProtected: aws.Boolean(true), // Required
	}
	resp, err := svc.SetTerminationProtection(params)

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

func ExampleEMR_SetVisibleToAllUsers() {
	svc := emr.New(nil)

	params := &emr.SetVisibleToAllUsersInput{
		JobFlowIDs: []*string{ // Required
			aws.String("XmlString"), // Required
			// More values...
		},
		VisibleToAllUsers: aws.Boolean(true), // Required
	}
	resp, err := svc.SetVisibleToAllUsers(params)

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

func ExampleEMR_TerminateJobFlows() {
	svc := emr.New(nil)

	params := &emr.TerminateJobFlowsInput{
		JobFlowIDs: []*string{ // Required
			aws.String("XmlString"), // Required
			// More values...
		},
	}
	resp, err := svc.TerminateJobFlows(params)

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