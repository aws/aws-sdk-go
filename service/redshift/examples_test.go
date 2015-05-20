package redshift_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/redshift"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleRedshift_AuthorizeClusterSecurityGroupIngress() {
	svc := redshift.New(nil)

	params := &redshift.AuthorizeClusterSecurityGroupIngressInput{
		ClusterSecurityGroupName: aws.String("String"), // Required
		CIDRIP:                  aws.String("String"),
		EC2SecurityGroupName:    aws.String("String"),
		EC2SecurityGroupOwnerID: aws.String("String"),
	}
	resp, err := svc.AuthorizeClusterSecurityGroupIngress(params)

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

func ExampleRedshift_AuthorizeSnapshotAccess() {
	svc := redshift.New(nil)

	params := &redshift.AuthorizeSnapshotAccessInput{
		AccountWithRestoreAccess:  aws.String("String"), // Required
		SnapshotIdentifier:        aws.String("String"), // Required
		SnapshotClusterIdentifier: aws.String("String"),
	}
	resp, err := svc.AuthorizeSnapshotAccess(params)

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

func ExampleRedshift_CopyClusterSnapshot() {
	svc := redshift.New(nil)

	params := &redshift.CopyClusterSnapshotInput{
		SourceSnapshotIdentifier:        aws.String("String"), // Required
		TargetSnapshotIdentifier:        aws.String("String"), // Required
		SourceSnapshotClusterIdentifier: aws.String("String"),
	}
	resp, err := svc.CopyClusterSnapshot(params)

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

func ExampleRedshift_CreateCluster() {
	svc := redshift.New(nil)

	params := &redshift.CreateClusterInput{
		ClusterIdentifier:                aws.String("String"), // Required
		MasterUserPassword:               aws.String("String"), // Required
		MasterUsername:                   aws.String("String"), // Required
		NodeType:                         aws.String("String"), // Required
		AllowVersionUpgrade:              aws.Boolean(true),
		AutomatedSnapshotRetentionPeriod: aws.Long(1),
		AvailabilityZone:                 aws.String("String"),
		ClusterParameterGroupName:        aws.String("String"),
		ClusterSecurityGroups: []*string{
			aws.String("String"), // Required
			// More values...
		},
		ClusterSubnetGroupName:         aws.String("String"),
		ClusterType:                    aws.String("String"),
		ClusterVersion:                 aws.String("String"),
		DBName:                         aws.String("String"),
		ElasticIP:                      aws.String("String"),
		Encrypted:                      aws.Boolean(true),
		HSMClientCertificateIdentifier: aws.String("String"),
		HSMConfigurationIdentifier:     aws.String("String"),
		KMSKeyID:                       aws.String("String"),
		NumberOfNodes:                  aws.Long(1),
		Port:                           aws.Long(1),
		PreferredMaintenanceWindow: aws.String("String"),
		PubliclyAccessible:         aws.Boolean(true),
		Tags: []*redshift.Tag{
			&redshift.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
		VPCSecurityGroupIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
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

func ExampleRedshift_CreateClusterParameterGroup() {
	svc := redshift.New(nil)

	params := &redshift.CreateClusterParameterGroupInput{
		Description:          aws.String("String"), // Required
		ParameterGroupFamily: aws.String("String"), // Required
		ParameterGroupName:   aws.String("String"), // Required
		Tags: []*redshift.Tag{
			&redshift.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateClusterParameterGroup(params)

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

func ExampleRedshift_CreateClusterSecurityGroup() {
	svc := redshift.New(nil)

	params := &redshift.CreateClusterSecurityGroupInput{
		ClusterSecurityGroupName: aws.String("String"), // Required
		Description:              aws.String("String"), // Required
		Tags: []*redshift.Tag{
			&redshift.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateClusterSecurityGroup(params)

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

func ExampleRedshift_CreateClusterSnapshot() {
	svc := redshift.New(nil)

	params := &redshift.CreateClusterSnapshotInput{
		ClusterIdentifier:  aws.String("String"), // Required
		SnapshotIdentifier: aws.String("String"), // Required
		Tags: []*redshift.Tag{
			&redshift.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateClusterSnapshot(params)

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

func ExampleRedshift_CreateClusterSubnetGroup() {
	svc := redshift.New(nil)

	params := &redshift.CreateClusterSubnetGroupInput{
		ClusterSubnetGroupName: aws.String("String"), // Required
		Description:            aws.String("String"), // Required
		SubnetIDs: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
		Tags: []*redshift.Tag{
			&redshift.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateClusterSubnetGroup(params)

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

func ExampleRedshift_CreateEventSubscription() {
	svc := redshift.New(nil)

	params := &redshift.CreateEventSubscriptionInput{
		SNSTopicARN:      aws.String("String"), // Required
		SubscriptionName: aws.String("String"), // Required
		Enabled:          aws.Boolean(true),
		EventCategories: []*string{
			aws.String("String"), // Required
			// More values...
		},
		Severity: aws.String("String"),
		SourceIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
		SourceType: aws.String("String"),
		Tags: []*redshift.Tag{
			&redshift.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateEventSubscription(params)

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

func ExampleRedshift_CreateHSMClientCertificate() {
	svc := redshift.New(nil)

	params := &redshift.CreateHSMClientCertificateInput{
		HSMClientCertificateIdentifier: aws.String("String"), // Required
		Tags: []*redshift.Tag{
			&redshift.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateHSMClientCertificate(params)

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

func ExampleRedshift_CreateHSMConfiguration() {
	svc := redshift.New(nil)

	params := &redshift.CreateHSMConfigurationInput{
		Description:                aws.String("String"), // Required
		HSMConfigurationIdentifier: aws.String("String"), // Required
		HSMIPAddress:               aws.String("String"), // Required
		HSMPartitionName:           aws.String("String"), // Required
		HSMPartitionPassword:       aws.String("String"), // Required
		HSMServerPublicCertificate: aws.String("String"), // Required
		Tags: []*redshift.Tag{
			&redshift.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateHSMConfiguration(params)

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

func ExampleRedshift_CreateTags() {
	svc := redshift.New(nil)

	params := &redshift.CreateTagsInput{
		ResourceName: aws.String("String"), // Required
		Tags: []*redshift.Tag{ // Required
			&redshift.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateTags(params)

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

func ExampleRedshift_DeleteCluster() {
	svc := redshift.New(nil)

	params := &redshift.DeleteClusterInput{
		ClusterIdentifier:              aws.String("String"), // Required
		FinalClusterSnapshotIdentifier: aws.String("String"),
		SkipFinalClusterSnapshot:       aws.Boolean(true),
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

func ExampleRedshift_DeleteClusterParameterGroup() {
	svc := redshift.New(nil)

	params := &redshift.DeleteClusterParameterGroupInput{
		ParameterGroupName: aws.String("String"), // Required
	}
	resp, err := svc.DeleteClusterParameterGroup(params)

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

func ExampleRedshift_DeleteClusterSecurityGroup() {
	svc := redshift.New(nil)

	params := &redshift.DeleteClusterSecurityGroupInput{
		ClusterSecurityGroupName: aws.String("String"), // Required
	}
	resp, err := svc.DeleteClusterSecurityGroup(params)

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

func ExampleRedshift_DeleteClusterSnapshot() {
	svc := redshift.New(nil)

	params := &redshift.DeleteClusterSnapshotInput{
		SnapshotIdentifier:        aws.String("String"), // Required
		SnapshotClusterIdentifier: aws.String("String"),
	}
	resp, err := svc.DeleteClusterSnapshot(params)

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

func ExampleRedshift_DeleteClusterSubnetGroup() {
	svc := redshift.New(nil)

	params := &redshift.DeleteClusterSubnetGroupInput{
		ClusterSubnetGroupName: aws.String("String"), // Required
	}
	resp, err := svc.DeleteClusterSubnetGroup(params)

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

func ExampleRedshift_DeleteEventSubscription() {
	svc := redshift.New(nil)

	params := &redshift.DeleteEventSubscriptionInput{
		SubscriptionName: aws.String("String"), // Required
	}
	resp, err := svc.DeleteEventSubscription(params)

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

func ExampleRedshift_DeleteHSMClientCertificate() {
	svc := redshift.New(nil)

	params := &redshift.DeleteHSMClientCertificateInput{
		HSMClientCertificateIdentifier: aws.String("String"), // Required
	}
	resp, err := svc.DeleteHSMClientCertificate(params)

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

func ExampleRedshift_DeleteHSMConfiguration() {
	svc := redshift.New(nil)

	params := &redshift.DeleteHSMConfigurationInput{
		HSMConfigurationIdentifier: aws.String("String"), // Required
	}
	resp, err := svc.DeleteHSMConfiguration(params)

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

func ExampleRedshift_DeleteTags() {
	svc := redshift.New(nil)

	params := &redshift.DeleteTagsInput{
		ResourceName: aws.String("String"), // Required
		TagKeys: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.DeleteTags(params)

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

func ExampleRedshift_DescribeClusterParameterGroups() {
	svc := redshift.New(nil)

	params := &redshift.DescribeClusterParameterGroupsInput{
		Marker:             aws.String("String"),
		MaxRecords:         aws.Long(1),
		ParameterGroupName: aws.String("String"),
		TagKeys: []*string{
			aws.String("String"), // Required
			// More values...
		},
		TagValues: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeClusterParameterGroups(params)

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

func ExampleRedshift_DescribeClusterParameters() {
	svc := redshift.New(nil)

	params := &redshift.DescribeClusterParametersInput{
		ParameterGroupName: aws.String("String"), // Required
		Marker:             aws.String("String"),
		MaxRecords:         aws.Long(1),
		Source:             aws.String("String"),
	}
	resp, err := svc.DescribeClusterParameters(params)

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

func ExampleRedshift_DescribeClusterSecurityGroups() {
	svc := redshift.New(nil)

	params := &redshift.DescribeClusterSecurityGroupsInput{
		ClusterSecurityGroupName: aws.String("String"),
		Marker:     aws.String("String"),
		MaxRecords: aws.Long(1),
		TagKeys: []*string{
			aws.String("String"), // Required
			// More values...
		},
		TagValues: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeClusterSecurityGroups(params)

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

func ExampleRedshift_DescribeClusterSnapshots() {
	svc := redshift.New(nil)

	params := &redshift.DescribeClusterSnapshotsInput{
		ClusterIdentifier:  aws.String("String"),
		EndTime:            aws.Time(time.Now()),
		Marker:             aws.String("String"),
		MaxRecords:         aws.Long(1),
		OwnerAccount:       aws.String("String"),
		SnapshotIdentifier: aws.String("String"),
		SnapshotType:       aws.String("String"),
		StartTime:          aws.Time(time.Now()),
		TagKeys: []*string{
			aws.String("String"), // Required
			// More values...
		},
		TagValues: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeClusterSnapshots(params)

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

func ExampleRedshift_DescribeClusterSubnetGroups() {
	svc := redshift.New(nil)

	params := &redshift.DescribeClusterSubnetGroupsInput{
		ClusterSubnetGroupName: aws.String("String"),
		Marker:                 aws.String("String"),
		MaxRecords:             aws.Long(1),
		TagKeys: []*string{
			aws.String("String"), // Required
			// More values...
		},
		TagValues: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeClusterSubnetGroups(params)

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

func ExampleRedshift_DescribeClusterVersions() {
	svc := redshift.New(nil)

	params := &redshift.DescribeClusterVersionsInput{
		ClusterParameterGroupFamily: aws.String("String"),
		ClusterVersion:              aws.String("String"),
		Marker:                      aws.String("String"),
		MaxRecords:                  aws.Long(1),
	}
	resp, err := svc.DescribeClusterVersions(params)

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

func ExampleRedshift_DescribeClusters() {
	svc := redshift.New(nil)

	params := &redshift.DescribeClustersInput{
		ClusterIdentifier: aws.String("String"),
		Marker:            aws.String("String"),
		MaxRecords:        aws.Long(1),
		TagKeys: []*string{
			aws.String("String"), // Required
			// More values...
		},
		TagValues: []*string{
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

func ExampleRedshift_DescribeDefaultClusterParameters() {
	svc := redshift.New(nil)

	params := &redshift.DescribeDefaultClusterParametersInput{
		ParameterGroupFamily: aws.String("String"), // Required
		Marker:               aws.String("String"),
		MaxRecords:           aws.Long(1),
	}
	resp, err := svc.DescribeDefaultClusterParameters(params)

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

func ExampleRedshift_DescribeEventCategories() {
	svc := redshift.New(nil)

	params := &redshift.DescribeEventCategoriesInput{
		SourceType: aws.String("String"),
	}
	resp, err := svc.DescribeEventCategories(params)

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

func ExampleRedshift_DescribeEventSubscriptions() {
	svc := redshift.New(nil)

	params := &redshift.DescribeEventSubscriptionsInput{
		Marker:           aws.String("String"),
		MaxRecords:       aws.Long(1),
		SubscriptionName: aws.String("String"),
	}
	resp, err := svc.DescribeEventSubscriptions(params)

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

func ExampleRedshift_DescribeEvents() {
	svc := redshift.New(nil)

	params := &redshift.DescribeEventsInput{
		Duration:         aws.Long(1),
		EndTime:          aws.Time(time.Now()),
		Marker:           aws.String("String"),
		MaxRecords:       aws.Long(1),
		SourceIdentifier: aws.String("String"),
		SourceType:       aws.String("SourceType"),
		StartTime:        aws.Time(time.Now()),
	}
	resp, err := svc.DescribeEvents(params)

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

func ExampleRedshift_DescribeHSMClientCertificates() {
	svc := redshift.New(nil)

	params := &redshift.DescribeHSMClientCertificatesInput{
		HSMClientCertificateIdentifier: aws.String("String"),
		Marker:     aws.String("String"),
		MaxRecords: aws.Long(1),
		TagKeys: []*string{
			aws.String("String"), // Required
			// More values...
		},
		TagValues: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeHSMClientCertificates(params)

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

func ExampleRedshift_DescribeHSMConfigurations() {
	svc := redshift.New(nil)

	params := &redshift.DescribeHSMConfigurationsInput{
		HSMConfigurationIdentifier: aws.String("String"),
		Marker:     aws.String("String"),
		MaxRecords: aws.Long(1),
		TagKeys: []*string{
			aws.String("String"), // Required
			// More values...
		},
		TagValues: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeHSMConfigurations(params)

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

func ExampleRedshift_DescribeLoggingStatus() {
	svc := redshift.New(nil)

	params := &redshift.DescribeLoggingStatusInput{
		ClusterIdentifier: aws.String("String"), // Required
	}
	resp, err := svc.DescribeLoggingStatus(params)

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

func ExampleRedshift_DescribeOrderableClusterOptions() {
	svc := redshift.New(nil)

	params := &redshift.DescribeOrderableClusterOptionsInput{
		ClusterVersion: aws.String("String"),
		Marker:         aws.String("String"),
		MaxRecords:     aws.Long(1),
		NodeType:       aws.String("String"),
	}
	resp, err := svc.DescribeOrderableClusterOptions(params)

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

func ExampleRedshift_DescribeReservedNodeOfferings() {
	svc := redshift.New(nil)

	params := &redshift.DescribeReservedNodeOfferingsInput{
		Marker:                 aws.String("String"),
		MaxRecords:             aws.Long(1),
		ReservedNodeOfferingID: aws.String("String"),
	}
	resp, err := svc.DescribeReservedNodeOfferings(params)

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

func ExampleRedshift_DescribeReservedNodes() {
	svc := redshift.New(nil)

	params := &redshift.DescribeReservedNodesInput{
		Marker:         aws.String("String"),
		MaxRecords:     aws.Long(1),
		ReservedNodeID: aws.String("String"),
	}
	resp, err := svc.DescribeReservedNodes(params)

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

func ExampleRedshift_DescribeResize() {
	svc := redshift.New(nil)

	params := &redshift.DescribeResizeInput{
		ClusterIdentifier: aws.String("String"), // Required
	}
	resp, err := svc.DescribeResize(params)

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

func ExampleRedshift_DescribeTags() {
	svc := redshift.New(nil)

	params := &redshift.DescribeTagsInput{
		Marker:       aws.String("String"),
		MaxRecords:   aws.Long(1),
		ResourceName: aws.String("String"),
		ResourceType: aws.String("String"),
		TagKeys: []*string{
			aws.String("String"), // Required
			// More values...
		},
		TagValues: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeTags(params)

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

func ExampleRedshift_DisableLogging() {
	svc := redshift.New(nil)

	params := &redshift.DisableLoggingInput{
		ClusterIdentifier: aws.String("String"), // Required
	}
	resp, err := svc.DisableLogging(params)

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

func ExampleRedshift_DisableSnapshotCopy() {
	svc := redshift.New(nil)

	params := &redshift.DisableSnapshotCopyInput{
		ClusterIdentifier: aws.String("String"), // Required
	}
	resp, err := svc.DisableSnapshotCopy(params)

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

func ExampleRedshift_EnableLogging() {
	svc := redshift.New(nil)

	params := &redshift.EnableLoggingInput{
		BucketName:        aws.String("String"), // Required
		ClusterIdentifier: aws.String("String"), // Required
		S3KeyPrefix:       aws.String("String"),
	}
	resp, err := svc.EnableLogging(params)

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

func ExampleRedshift_EnableSnapshotCopy() {
	svc := redshift.New(nil)

	params := &redshift.EnableSnapshotCopyInput{
		ClusterIdentifier: aws.String("String"), // Required
		DestinationRegion: aws.String("String"), // Required
		RetentionPeriod:   aws.Long(1),
	}
	resp, err := svc.EnableSnapshotCopy(params)

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

func ExampleRedshift_ModifyCluster() {
	svc := redshift.New(nil)

	params := &redshift.ModifyClusterInput{
		ClusterIdentifier:                aws.String("String"), // Required
		AllowVersionUpgrade:              aws.Boolean(true),
		AutomatedSnapshotRetentionPeriod: aws.Long(1),
		ClusterParameterGroupName:        aws.String("String"),
		ClusterSecurityGroups: []*string{
			aws.String("String"), // Required
			// More values...
		},
		ClusterType:                    aws.String("String"),
		ClusterVersion:                 aws.String("String"),
		HSMClientCertificateIdentifier: aws.String("String"),
		HSMConfigurationIdentifier:     aws.String("String"),
		MasterUserPassword:             aws.String("String"),
		NewClusterIdentifier:           aws.String("String"),
		NodeType:                       aws.String("String"),
		NumberOfNodes:                  aws.Long(1),
		PreferredMaintenanceWindow:     aws.String("String"),
		VPCSecurityGroupIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.ModifyCluster(params)

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

func ExampleRedshift_ModifyClusterParameterGroup() {
	svc := redshift.New(nil)

	params := &redshift.ModifyClusterParameterGroupInput{
		ParameterGroupName: aws.String("String"), // Required
		Parameters: []*redshift.Parameter{ // Required
			&redshift.Parameter{ // Required
				AllowedValues:        aws.String("String"),
				DataType:             aws.String("String"),
				Description:          aws.String("String"),
				IsModifiable:         aws.Boolean(true),
				MinimumEngineVersion: aws.String("String"),
				ParameterName:        aws.String("String"),
				ParameterValue:       aws.String("String"),
				Source:               aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.ModifyClusterParameterGroup(params)

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

func ExampleRedshift_ModifyClusterSubnetGroup() {
	svc := redshift.New(nil)

	params := &redshift.ModifyClusterSubnetGroupInput{
		ClusterSubnetGroupName: aws.String("String"), // Required
		SubnetIDs: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
		Description: aws.String("String"),
	}
	resp, err := svc.ModifyClusterSubnetGroup(params)

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

func ExampleRedshift_ModifyEventSubscription() {
	svc := redshift.New(nil)

	params := &redshift.ModifyEventSubscriptionInput{
		SubscriptionName: aws.String("String"), // Required
		Enabled:          aws.Boolean(true),
		EventCategories: []*string{
			aws.String("String"), // Required
			// More values...
		},
		SNSTopicARN: aws.String("String"),
		Severity:    aws.String("String"),
		SourceIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
		SourceType: aws.String("String"),
	}
	resp, err := svc.ModifyEventSubscription(params)

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

func ExampleRedshift_ModifySnapshotCopyRetentionPeriod() {
	svc := redshift.New(nil)

	params := &redshift.ModifySnapshotCopyRetentionPeriodInput{
		ClusterIdentifier: aws.String("String"), // Required
		RetentionPeriod:   aws.Long(1),          // Required
	}
	resp, err := svc.ModifySnapshotCopyRetentionPeriod(params)

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

func ExampleRedshift_PurchaseReservedNodeOffering() {
	svc := redshift.New(nil)

	params := &redshift.PurchaseReservedNodeOfferingInput{
		ReservedNodeOfferingID: aws.String("String"), // Required
		NodeCount:              aws.Long(1),
	}
	resp, err := svc.PurchaseReservedNodeOffering(params)

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

func ExampleRedshift_RebootCluster() {
	svc := redshift.New(nil)

	params := &redshift.RebootClusterInput{
		ClusterIdentifier: aws.String("String"), // Required
	}
	resp, err := svc.RebootCluster(params)

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

func ExampleRedshift_ResetClusterParameterGroup() {
	svc := redshift.New(nil)

	params := &redshift.ResetClusterParameterGroupInput{
		ParameterGroupName: aws.String("String"), // Required
		Parameters: []*redshift.Parameter{
			&redshift.Parameter{ // Required
				AllowedValues:        aws.String("String"),
				DataType:             aws.String("String"),
				Description:          aws.String("String"),
				IsModifiable:         aws.Boolean(true),
				MinimumEngineVersion: aws.String("String"),
				ParameterName:        aws.String("String"),
				ParameterValue:       aws.String("String"),
				Source:               aws.String("String"),
			},
			// More values...
		},
		ResetAllParameters: aws.Boolean(true),
	}
	resp, err := svc.ResetClusterParameterGroup(params)

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

func ExampleRedshift_RestoreFromClusterSnapshot() {
	svc := redshift.New(nil)

	params := &redshift.RestoreFromClusterSnapshotInput{
		ClusterIdentifier:                aws.String("String"), // Required
		SnapshotIdentifier:               aws.String("String"), // Required
		AllowVersionUpgrade:              aws.Boolean(true),
		AutomatedSnapshotRetentionPeriod: aws.Long(1),
		AvailabilityZone:                 aws.String("String"),
		ClusterParameterGroupName:        aws.String("String"),
		ClusterSecurityGroups: []*string{
			aws.String("String"), // Required
			// More values...
		},
		ClusterSubnetGroupName:         aws.String("String"),
		ElasticIP:                      aws.String("String"),
		HSMClientCertificateIdentifier: aws.String("String"),
		HSMConfigurationIdentifier:     aws.String("String"),
		KMSKeyID:                       aws.String("String"),
		OwnerAccount:                   aws.String("String"),
		Port:                           aws.Long(1),
		PreferredMaintenanceWindow: aws.String("String"),
		PubliclyAccessible:         aws.Boolean(true),
		SnapshotClusterIdentifier:  aws.String("String"),
		VPCSecurityGroupIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.RestoreFromClusterSnapshot(params)

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

func ExampleRedshift_RevokeClusterSecurityGroupIngress() {
	svc := redshift.New(nil)

	params := &redshift.RevokeClusterSecurityGroupIngressInput{
		ClusterSecurityGroupName: aws.String("String"), // Required
		CIDRIP:                  aws.String("String"),
		EC2SecurityGroupName:    aws.String("String"),
		EC2SecurityGroupOwnerID: aws.String("String"),
	}
	resp, err := svc.RevokeClusterSecurityGroupIngress(params)

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

func ExampleRedshift_RevokeSnapshotAccess() {
	svc := redshift.New(nil)

	params := &redshift.RevokeSnapshotAccessInput{
		AccountWithRestoreAccess:  aws.String("String"), // Required
		SnapshotIdentifier:        aws.String("String"), // Required
		SnapshotClusterIdentifier: aws.String("String"),
	}
	resp, err := svc.RevokeSnapshotAccess(params)

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

func ExampleRedshift_RotateEncryptionKey() {
	svc := redshift.New(nil)

	params := &redshift.RotateEncryptionKeyInput{
		ClusterIdentifier: aws.String("String"), // Required
	}
	resp, err := svc.RotateEncryptionKey(params)

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