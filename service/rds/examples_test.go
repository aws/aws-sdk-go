package rds_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/rds"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleRDS_AddSourceIdentifierToSubscription() {
	svc := rds.New(nil)

	params := &rds.AddSourceIdentifierToSubscriptionInput{
		SourceIdentifier: aws.String("String"), // Required
		SubscriptionName: aws.String("String"), // Required
	}
	resp, err := svc.AddSourceIdentifierToSubscription(params)

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

func ExampleRDS_AddTagsToResource() {
	svc := rds.New(nil)

	params := &rds.AddTagsToResourceInput{
		ResourceName: aws.String("String"), // Required
		Tags: []*rds.Tag{ // Required
			&rds.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.AddTagsToResource(params)

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

func ExampleRDS_ApplyPendingMaintenanceAction() {
	svc := rds.New(nil)

	params := &rds.ApplyPendingMaintenanceActionInput{
		ApplyAction:        aws.String("String"), // Required
		OptInType:          aws.String("String"), // Required
		ResourceIdentifier: aws.String("String"), // Required
	}
	resp, err := svc.ApplyPendingMaintenanceAction(params)

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

func ExampleRDS_AuthorizeDBSecurityGroupIngress() {
	svc := rds.New(nil)

	params := &rds.AuthorizeDBSecurityGroupIngressInput{
		DBSecurityGroupName:     aws.String("String"), // Required
		CIDRIP:                  aws.String("String"),
		EC2SecurityGroupID:      aws.String("String"),
		EC2SecurityGroupName:    aws.String("String"),
		EC2SecurityGroupOwnerID: aws.String("String"),
	}
	resp, err := svc.AuthorizeDBSecurityGroupIngress(params)

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

func ExampleRDS_CopyDBParameterGroup() {
	svc := rds.New(nil)

	params := &rds.CopyDBParameterGroupInput{
		SourceDBParameterGroupIdentifier:  aws.String("String"), // Required
		TargetDBParameterGroupDescription: aws.String("String"), // Required
		TargetDBParameterGroupIdentifier:  aws.String("String"), // Required
		Tags: []*rds.Tag{
			&rds.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CopyDBParameterGroup(params)

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

func ExampleRDS_CopyDBSnapshot() {
	svc := rds.New(nil)

	params := &rds.CopyDBSnapshotInput{
		SourceDBSnapshotIdentifier: aws.String("String"), // Required
		TargetDBSnapshotIdentifier: aws.String("String"), // Required
		Tags: []*rds.Tag{
			&rds.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CopyDBSnapshot(params)

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

func ExampleRDS_CopyOptionGroup() {
	svc := rds.New(nil)

	params := &rds.CopyOptionGroupInput{
		SourceOptionGroupIdentifier:  aws.String("String"), // Required
		TargetOptionGroupDescription: aws.String("String"), // Required
		TargetOptionGroupIdentifier:  aws.String("String"), // Required
		Tags: []*rds.Tag{
			&rds.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CopyOptionGroup(params)

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

func ExampleRDS_CreateDBInstance() {
	svc := rds.New(nil)

	params := &rds.CreateDBInstanceInput{
		AllocatedStorage:        aws.Long(1),          // Required
		DBInstanceClass:         aws.String("String"), // Required
		DBInstanceIdentifier:    aws.String("String"), // Required
		Engine:                  aws.String("String"), // Required
		MasterUserPassword:      aws.String("String"), // Required
		MasterUsername:          aws.String("String"), // Required
		AutoMinorVersionUpgrade: aws.Boolean(true),
		AvailabilityZone:        aws.String("String"),
		BackupRetentionPeriod:   aws.Long(1),
		CharacterSetName:        aws.String("String"),
		DBName:                  aws.String("String"),
		DBParameterGroupName:    aws.String("String"),
		DBSecurityGroups: []*string{
			aws.String("String"), // Required
			// More values...
		},
		DBSubnetGroupName: aws.String("String"),
		EngineVersion:     aws.String("String"),
		IOPS:              aws.Long(1),
		KMSKeyID:          aws.String("String"),
		LicenseModel:      aws.String("String"),
		MultiAZ:           aws.Boolean(true),
		OptionGroupName:   aws.String("String"),
		Port:              aws.Long(1),
		PreferredBackupWindow:      aws.String("String"),
		PreferredMaintenanceWindow: aws.String("String"),
		PubliclyAccessible:         aws.Boolean(true),
		StorageEncrypted:           aws.Boolean(true),
		StorageType:                aws.String("String"),
		TDECredentialARN:           aws.String("String"),
		TDECredentialPassword:      aws.String("String"),
		Tags: []*rds.Tag{
			&rds.Tag{ // Required
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
	resp, err := svc.CreateDBInstance(params)

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

func ExampleRDS_CreateDBInstanceReadReplica() {
	svc := rds.New(nil)

	params := &rds.CreateDBInstanceReadReplicaInput{
		DBInstanceIdentifier:       aws.String("String"), // Required
		SourceDBInstanceIdentifier: aws.String("String"), // Required
		AutoMinorVersionUpgrade:    aws.Boolean(true),
		AvailabilityZone:           aws.String("String"),
		DBInstanceClass:            aws.String("String"),
		DBSubnetGroupName:          aws.String("String"),
		IOPS:                       aws.Long(1),
		OptionGroupName:            aws.String("String"),
		Port:                       aws.Long(1),
		PubliclyAccessible:         aws.Boolean(true),
		StorageType:                aws.String("String"),
		Tags: []*rds.Tag{
			&rds.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateDBInstanceReadReplica(params)

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

func ExampleRDS_CreateDBParameterGroup() {
	svc := rds.New(nil)

	params := &rds.CreateDBParameterGroupInput{
		DBParameterGroupFamily: aws.String("String"), // Required
		DBParameterGroupName:   aws.String("String"), // Required
		Description:            aws.String("String"), // Required
		Tags: []*rds.Tag{
			&rds.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateDBParameterGroup(params)

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

func ExampleRDS_CreateDBSecurityGroup() {
	svc := rds.New(nil)

	params := &rds.CreateDBSecurityGroupInput{
		DBSecurityGroupDescription: aws.String("String"), // Required
		DBSecurityGroupName:        aws.String("String"), // Required
		Tags: []*rds.Tag{
			&rds.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateDBSecurityGroup(params)

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

func ExampleRDS_CreateDBSnapshot() {
	svc := rds.New(nil)

	params := &rds.CreateDBSnapshotInput{
		DBInstanceIdentifier: aws.String("String"), // Required
		DBSnapshotIdentifier: aws.String("String"), // Required
		Tags: []*rds.Tag{
			&rds.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateDBSnapshot(params)

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

func ExampleRDS_CreateDBSubnetGroup() {
	svc := rds.New(nil)

	params := &rds.CreateDBSubnetGroupInput{
		DBSubnetGroupDescription: aws.String("String"), // Required
		DBSubnetGroupName:        aws.String("String"), // Required
		SubnetIDs: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
		Tags: []*rds.Tag{
			&rds.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateDBSubnetGroup(params)

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

func ExampleRDS_CreateEventSubscription() {
	svc := rds.New(nil)

	params := &rds.CreateEventSubscriptionInput{
		SNSTopicARN:      aws.String("String"), // Required
		SubscriptionName: aws.String("String"), // Required
		Enabled:          aws.Boolean(true),
		EventCategories: []*string{
			aws.String("String"), // Required
			// More values...
		},
		SourceIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
		SourceType: aws.String("String"),
		Tags: []*rds.Tag{
			&rds.Tag{ // Required
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

func ExampleRDS_CreateOptionGroup() {
	svc := rds.New(nil)

	params := &rds.CreateOptionGroupInput{
		EngineName:             aws.String("String"), // Required
		MajorEngineVersion:     aws.String("String"), // Required
		OptionGroupDescription: aws.String("String"), // Required
		OptionGroupName:        aws.String("String"), // Required
		Tags: []*rds.Tag{
			&rds.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateOptionGroup(params)

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

func ExampleRDS_DeleteDBInstance() {
	svc := rds.New(nil)

	params := &rds.DeleteDBInstanceInput{
		DBInstanceIdentifier:      aws.String("String"), // Required
		FinalDBSnapshotIdentifier: aws.String("String"),
		SkipFinalSnapshot:         aws.Boolean(true),
	}
	resp, err := svc.DeleteDBInstance(params)

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

func ExampleRDS_DeleteDBParameterGroup() {
	svc := rds.New(nil)

	params := &rds.DeleteDBParameterGroupInput{
		DBParameterGroupName: aws.String("String"), // Required
	}
	resp, err := svc.DeleteDBParameterGroup(params)

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

func ExampleRDS_DeleteDBSecurityGroup() {
	svc := rds.New(nil)

	params := &rds.DeleteDBSecurityGroupInput{
		DBSecurityGroupName: aws.String("String"), // Required
	}
	resp, err := svc.DeleteDBSecurityGroup(params)

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

func ExampleRDS_DeleteDBSnapshot() {
	svc := rds.New(nil)

	params := &rds.DeleteDBSnapshotInput{
		DBSnapshotIdentifier: aws.String("String"), // Required
	}
	resp, err := svc.DeleteDBSnapshot(params)

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

func ExampleRDS_DeleteDBSubnetGroup() {
	svc := rds.New(nil)

	params := &rds.DeleteDBSubnetGroupInput{
		DBSubnetGroupName: aws.String("String"), // Required
	}
	resp, err := svc.DeleteDBSubnetGroup(params)

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

func ExampleRDS_DeleteEventSubscription() {
	svc := rds.New(nil)

	params := &rds.DeleteEventSubscriptionInput{
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

func ExampleRDS_DeleteOptionGroup() {
	svc := rds.New(nil)

	params := &rds.DeleteOptionGroupInput{
		OptionGroupName: aws.String("String"), // Required
	}
	resp, err := svc.DeleteOptionGroup(params)

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

func ExampleRDS_DescribeDBEngineVersions() {
	svc := rds.New(nil)

	params := &rds.DescribeDBEngineVersionsInput{
		DBParameterGroupFamily: aws.String("String"),
		DefaultOnly:            aws.Boolean(true),
		Engine:                 aws.String("String"),
		EngineVersion:          aws.String("String"),
		Filters: []*rds.Filter{
			&rds.Filter{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		ListSupportedCharacterSets: aws.Boolean(true),
		Marker:     aws.String("String"),
		MaxRecords: aws.Long(1),
	}
	resp, err := svc.DescribeDBEngineVersions(params)

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

func ExampleRDS_DescribeDBInstances() {
	svc := rds.New(nil)

	params := &rds.DescribeDBInstancesInput{
		DBInstanceIdentifier: aws.String("String"),
		Filters: []*rds.Filter{
			&rds.Filter{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:     aws.String("String"),
		MaxRecords: aws.Long(1),
	}
	resp, err := svc.DescribeDBInstances(params)

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

func ExampleRDS_DescribeDBLogFiles() {
	svc := rds.New(nil)

	params := &rds.DescribeDBLogFilesInput{
		DBInstanceIdentifier: aws.String("String"), // Required
		FileLastWritten:      aws.Long(1),
		FileSize:             aws.Long(1),
		FilenameContains:     aws.String("String"),
		Filters: []*rds.Filter{
			&rds.Filter{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:     aws.String("String"),
		MaxRecords: aws.Long(1),
	}
	resp, err := svc.DescribeDBLogFiles(params)

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

func ExampleRDS_DescribeDBParameterGroups() {
	svc := rds.New(nil)

	params := &rds.DescribeDBParameterGroupsInput{
		DBParameterGroupName: aws.String("String"),
		Filters: []*rds.Filter{
			&rds.Filter{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:     aws.String("String"),
		MaxRecords: aws.Long(1),
	}
	resp, err := svc.DescribeDBParameterGroups(params)

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

func ExampleRDS_DescribeDBParameters() {
	svc := rds.New(nil)

	params := &rds.DescribeDBParametersInput{
		DBParameterGroupName: aws.String("String"), // Required
		Filters: []*rds.Filter{
			&rds.Filter{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:     aws.String("String"),
		MaxRecords: aws.Long(1),
		Source:     aws.String("String"),
	}
	resp, err := svc.DescribeDBParameters(params)

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

func ExampleRDS_DescribeDBSecurityGroups() {
	svc := rds.New(nil)

	params := &rds.DescribeDBSecurityGroupsInput{
		DBSecurityGroupName: aws.String("String"),
		Filters: []*rds.Filter{
			&rds.Filter{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:     aws.String("String"),
		MaxRecords: aws.Long(1),
	}
	resp, err := svc.DescribeDBSecurityGroups(params)

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

func ExampleRDS_DescribeDBSnapshots() {
	svc := rds.New(nil)

	params := &rds.DescribeDBSnapshotsInput{
		DBInstanceIdentifier: aws.String("String"),
		DBSnapshotIdentifier: aws.String("String"),
		Filters: []*rds.Filter{
			&rds.Filter{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:       aws.String("String"),
		MaxRecords:   aws.Long(1),
		SnapshotType: aws.String("String"),
	}
	resp, err := svc.DescribeDBSnapshots(params)

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

func ExampleRDS_DescribeDBSubnetGroups() {
	svc := rds.New(nil)

	params := &rds.DescribeDBSubnetGroupsInput{
		DBSubnetGroupName: aws.String("String"),
		Filters: []*rds.Filter{
			&rds.Filter{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:     aws.String("String"),
		MaxRecords: aws.Long(1),
	}
	resp, err := svc.DescribeDBSubnetGroups(params)

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

func ExampleRDS_DescribeEngineDefaultParameters() {
	svc := rds.New(nil)

	params := &rds.DescribeEngineDefaultParametersInput{
		DBParameterGroupFamily: aws.String("String"), // Required
		Filters: []*rds.Filter{
			&rds.Filter{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:     aws.String("String"),
		MaxRecords: aws.Long(1),
	}
	resp, err := svc.DescribeEngineDefaultParameters(params)

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

func ExampleRDS_DescribeEventCategories() {
	svc := rds.New(nil)

	params := &rds.DescribeEventCategoriesInput{
		Filters: []*rds.Filter{
			&rds.Filter{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
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

func ExampleRDS_DescribeEventSubscriptions() {
	svc := rds.New(nil)

	params := &rds.DescribeEventSubscriptionsInput{
		Filters: []*rds.Filter{
			&rds.Filter{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
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

func ExampleRDS_DescribeEvents() {
	svc := rds.New(nil)

	params := &rds.DescribeEventsInput{
		Duration: aws.Long(1),
		EndTime:  aws.Time(time.Now()),
		EventCategories: []*string{
			aws.String("String"), // Required
			// More values...
		},
		Filters: []*rds.Filter{
			&rds.Filter{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
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

func ExampleRDS_DescribeOptionGroupOptions() {
	svc := rds.New(nil)

	params := &rds.DescribeOptionGroupOptionsInput{
		EngineName: aws.String("String"), // Required
		Filters: []*rds.Filter{
			&rds.Filter{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		MajorEngineVersion: aws.String("String"),
		Marker:             aws.String("String"),
		MaxRecords:         aws.Long(1),
	}
	resp, err := svc.DescribeOptionGroupOptions(params)

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

func ExampleRDS_DescribeOptionGroups() {
	svc := rds.New(nil)

	params := &rds.DescribeOptionGroupsInput{
		EngineName: aws.String("String"),
		Filters: []*rds.Filter{
			&rds.Filter{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		MajorEngineVersion: aws.String("String"),
		Marker:             aws.String("String"),
		MaxRecords:         aws.Long(1),
		OptionGroupName:    aws.String("String"),
	}
	resp, err := svc.DescribeOptionGroups(params)

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

func ExampleRDS_DescribeOrderableDBInstanceOptions() {
	svc := rds.New(nil)

	params := &rds.DescribeOrderableDBInstanceOptionsInput{
		Engine:          aws.String("String"), // Required
		DBInstanceClass: aws.String("String"),
		EngineVersion:   aws.String("String"),
		Filters: []*rds.Filter{
			&rds.Filter{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		LicenseModel: aws.String("String"),
		Marker:       aws.String("String"),
		MaxRecords:   aws.Long(1),
		VPC:          aws.Boolean(true),
	}
	resp, err := svc.DescribeOrderableDBInstanceOptions(params)

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

func ExampleRDS_DescribePendingMaintenanceActions() {
	svc := rds.New(nil)

	params := &rds.DescribePendingMaintenanceActionsInput{
		Filters: []*rds.Filter{
			&rds.Filter{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:             aws.String("String"),
		MaxRecords:         aws.Long(1),
		ResourceIdentifier: aws.String("String"),
	}
	resp, err := svc.DescribePendingMaintenanceActions(params)

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

func ExampleRDS_DescribeReservedDBInstances() {
	svc := rds.New(nil)

	params := &rds.DescribeReservedDBInstancesInput{
		DBInstanceClass: aws.String("String"),
		Duration:        aws.String("String"),
		Filters: []*rds.Filter{
			&rds.Filter{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:                        aws.String("String"),
		MaxRecords:                    aws.Long(1),
		MultiAZ:                       aws.Boolean(true),
		OfferingType:                  aws.String("String"),
		ProductDescription:            aws.String("String"),
		ReservedDBInstanceID:          aws.String("String"),
		ReservedDBInstancesOfferingID: aws.String("String"),
	}
	resp, err := svc.DescribeReservedDBInstances(params)

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

func ExampleRDS_DescribeReservedDBInstancesOfferings() {
	svc := rds.New(nil)

	params := &rds.DescribeReservedDBInstancesOfferingsInput{
		DBInstanceClass: aws.String("String"),
		Duration:        aws.String("String"),
		Filters: []*rds.Filter{
			&rds.Filter{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:                        aws.String("String"),
		MaxRecords:                    aws.Long(1),
		MultiAZ:                       aws.Boolean(true),
		OfferingType:                  aws.String("String"),
		ProductDescription:            aws.String("String"),
		ReservedDBInstancesOfferingID: aws.String("String"),
	}
	resp, err := svc.DescribeReservedDBInstancesOfferings(params)

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

func ExampleRDS_DownloadDBLogFilePortion() {
	svc := rds.New(nil)

	params := &rds.DownloadDBLogFilePortionInput{
		DBInstanceIdentifier: aws.String("String"), // Required
		LogFileName:          aws.String("String"), // Required
		Marker:               aws.String("String"),
		NumberOfLines:        aws.Long(1),
	}
	resp, err := svc.DownloadDBLogFilePortion(params)

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

func ExampleRDS_ListTagsForResource() {
	svc := rds.New(nil)

	params := &rds.ListTagsForResourceInput{
		ResourceName: aws.String("String"), // Required
		Filters: []*rds.Filter{
			&rds.Filter{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
	}
	resp, err := svc.ListTagsForResource(params)

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

func ExampleRDS_ModifyDBInstance() {
	svc := rds.New(nil)

	params := &rds.ModifyDBInstanceInput{
		DBInstanceIdentifier:     aws.String("String"), // Required
		AllocatedStorage:         aws.Long(1),
		AllowMajorVersionUpgrade: aws.Boolean(true),
		ApplyImmediately:         aws.Boolean(true),
		AutoMinorVersionUpgrade:  aws.Boolean(true),
		BackupRetentionPeriod:    aws.Long(1),
		DBInstanceClass:          aws.String("String"),
		DBParameterGroupName:     aws.String("String"),
		DBSecurityGroups: []*string{
			aws.String("String"), // Required
			// More values...
		},
		EngineVersion:              aws.String("String"),
		IOPS:                       aws.Long(1),
		MasterUserPassword:         aws.String("String"),
		MultiAZ:                    aws.Boolean(true),
		NewDBInstanceIdentifier:    aws.String("String"),
		OptionGroupName:            aws.String("String"),
		PreferredBackupWindow:      aws.String("String"),
		PreferredMaintenanceWindow: aws.String("String"),
		StorageType:                aws.String("String"),
		TDECredentialARN:           aws.String("String"),
		TDECredentialPassword:      aws.String("String"),
		VPCSecurityGroupIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.ModifyDBInstance(params)

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

func ExampleRDS_ModifyDBParameterGroup() {
	svc := rds.New(nil)

	params := &rds.ModifyDBParameterGroupInput{
		DBParameterGroupName: aws.String("String"), // Required
		Parameters: []*rds.Parameter{ // Required
			&rds.Parameter{ // Required
				AllowedValues:        aws.String("String"),
				ApplyMethod:          aws.String("ApplyMethod"),
				ApplyType:            aws.String("String"),
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
	resp, err := svc.ModifyDBParameterGroup(params)

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

func ExampleRDS_ModifyDBSubnetGroup() {
	svc := rds.New(nil)

	params := &rds.ModifyDBSubnetGroupInput{
		DBSubnetGroupName: aws.String("String"), // Required
		SubnetIDs: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
		DBSubnetGroupDescription: aws.String("String"),
	}
	resp, err := svc.ModifyDBSubnetGroup(params)

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

func ExampleRDS_ModifyEventSubscription() {
	svc := rds.New(nil)

	params := &rds.ModifyEventSubscriptionInput{
		SubscriptionName: aws.String("String"), // Required
		Enabled:          aws.Boolean(true),
		EventCategories: []*string{
			aws.String("String"), // Required
			// More values...
		},
		SNSTopicARN: aws.String("String"),
		SourceType:  aws.String("String"),
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

func ExampleRDS_ModifyOptionGroup() {
	svc := rds.New(nil)

	params := &rds.ModifyOptionGroupInput{
		OptionGroupName:  aws.String("String"), // Required
		ApplyImmediately: aws.Boolean(true),
		OptionsToInclude: []*rds.OptionConfiguration{
			&rds.OptionConfiguration{ // Required
				OptionName: aws.String("String"), // Required
				DBSecurityGroupMemberships: []*string{
					aws.String("String"), // Required
					// More values...
				},
				OptionSettings: []*rds.OptionSetting{
					&rds.OptionSetting{ // Required
						AllowedValues: aws.String("String"),
						ApplyType:     aws.String("String"),
						DataType:      aws.String("String"),
						DefaultValue:  aws.String("String"),
						Description:   aws.String("String"),
						IsCollection:  aws.Boolean(true),
						IsModifiable:  aws.Boolean(true),
						Name:          aws.String("String"),
						Value:         aws.String("String"),
					},
					// More values...
				},
				Port: aws.Long(1),
				VPCSecurityGroupMemberships: []*string{
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		OptionsToRemove: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.ModifyOptionGroup(params)

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

func ExampleRDS_PromoteReadReplica() {
	svc := rds.New(nil)

	params := &rds.PromoteReadReplicaInput{
		DBInstanceIdentifier:  aws.String("String"), // Required
		BackupRetentionPeriod: aws.Long(1),
		PreferredBackupWindow: aws.String("String"),
	}
	resp, err := svc.PromoteReadReplica(params)

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

func ExampleRDS_PurchaseReservedDBInstancesOffering() {
	svc := rds.New(nil)

	params := &rds.PurchaseReservedDBInstancesOfferingInput{
		ReservedDBInstancesOfferingID: aws.String("String"), // Required
		DBInstanceCount:               aws.Long(1),
		ReservedDBInstanceID:          aws.String("String"),
		Tags: []*rds.Tag{
			&rds.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.PurchaseReservedDBInstancesOffering(params)

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

func ExampleRDS_RebootDBInstance() {
	svc := rds.New(nil)

	params := &rds.RebootDBInstanceInput{
		DBInstanceIdentifier: aws.String("String"), // Required
		ForceFailover:        aws.Boolean(true),
	}
	resp, err := svc.RebootDBInstance(params)

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

func ExampleRDS_RemoveSourceIdentifierFromSubscription() {
	svc := rds.New(nil)

	params := &rds.RemoveSourceIdentifierFromSubscriptionInput{
		SourceIdentifier: aws.String("String"), // Required
		SubscriptionName: aws.String("String"), // Required
	}
	resp, err := svc.RemoveSourceIdentifierFromSubscription(params)

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

func ExampleRDS_RemoveTagsFromResource() {
	svc := rds.New(nil)

	params := &rds.RemoveTagsFromResourceInput{
		ResourceName: aws.String("String"), // Required
		TagKeys: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.RemoveTagsFromResource(params)

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

func ExampleRDS_ResetDBParameterGroup() {
	svc := rds.New(nil)

	params := &rds.ResetDBParameterGroupInput{
		DBParameterGroupName: aws.String("String"), // Required
		Parameters: []*rds.Parameter{
			&rds.Parameter{ // Required
				AllowedValues:        aws.String("String"),
				ApplyMethod:          aws.String("ApplyMethod"),
				ApplyType:            aws.String("String"),
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
	resp, err := svc.ResetDBParameterGroup(params)

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

func ExampleRDS_RestoreDBInstanceFromDBSnapshot() {
	svc := rds.New(nil)

	params := &rds.RestoreDBInstanceFromDBSnapshotInput{
		DBInstanceIdentifier:    aws.String("String"), // Required
		DBSnapshotIdentifier:    aws.String("String"), // Required
		AutoMinorVersionUpgrade: aws.Boolean(true),
		AvailabilityZone:        aws.String("String"),
		DBInstanceClass:         aws.String("String"),
		DBName:                  aws.String("String"),
		DBSubnetGroupName:       aws.String("String"),
		Engine:                  aws.String("String"),
		IOPS:                    aws.Long(1),
		LicenseModel:            aws.String("String"),
		MultiAZ:                 aws.Boolean(true),
		OptionGroupName:         aws.String("String"),
		Port:                    aws.Long(1),
		PubliclyAccessible:      aws.Boolean(true),
		StorageType:             aws.String("String"),
		TDECredentialARN:        aws.String("String"),
		TDECredentialPassword:   aws.String("String"),
		Tags: []*rds.Tag{
			&rds.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.RestoreDBInstanceFromDBSnapshot(params)

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

func ExampleRDS_RestoreDBInstanceToPointInTime() {
	svc := rds.New(nil)

	params := &rds.RestoreDBInstanceToPointInTimeInput{
		SourceDBInstanceIdentifier: aws.String("String"), // Required
		TargetDBInstanceIdentifier: aws.String("String"), // Required
		AutoMinorVersionUpgrade:    aws.Boolean(true),
		AvailabilityZone:           aws.String("String"),
		DBInstanceClass:            aws.String("String"),
		DBName:                     aws.String("String"),
		DBSubnetGroupName:          aws.String("String"),
		Engine:                     aws.String("String"),
		IOPS:                       aws.Long(1),
		LicenseModel:               aws.String("String"),
		MultiAZ:                    aws.Boolean(true),
		OptionGroupName:            aws.String("String"),
		Port:                       aws.Long(1),
		PubliclyAccessible:         aws.Boolean(true),
		RestoreTime:                aws.Time(time.Now()),
		StorageType:                aws.String("String"),
		TDECredentialARN:           aws.String("String"),
		TDECredentialPassword:      aws.String("String"),
		Tags: []*rds.Tag{
			&rds.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
		UseLatestRestorableTime: aws.Boolean(true),
	}
	resp, err := svc.RestoreDBInstanceToPointInTime(params)

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

func ExampleRDS_RevokeDBSecurityGroupIngress() {
	svc := rds.New(nil)

	params := &rds.RevokeDBSecurityGroupIngressInput{
		DBSecurityGroupName:     aws.String("String"), // Required
		CIDRIP:                  aws.String("String"),
		EC2SecurityGroupID:      aws.String("String"),
		EC2SecurityGroupName:    aws.String("String"),
		EC2SecurityGroupOwnerID: aws.String("String"),
	}
	resp, err := svc.RevokeDBSecurityGroupIngress(params)

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