package elasticache_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/elasticache"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleElastiCache_AddTagsToResource() {
	svc := elasticache.New(nil)

	params := &elasticache.AddTagsToResourceInput{
		ResourceName: aws.String("String"), // Required
		Tags: []*elasticache.Tag{ // Required
			&elasticache.Tag{ // Required
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

func ExampleElastiCache_AuthorizeCacheSecurityGroupIngress() {
	svc := elasticache.New(nil)

	params := &elasticache.AuthorizeCacheSecurityGroupIngressInput{
		CacheSecurityGroupName:  aws.String("String"), // Required
		EC2SecurityGroupName:    aws.String("String"), // Required
		EC2SecurityGroupOwnerID: aws.String("String"), // Required
	}
	resp, err := svc.AuthorizeCacheSecurityGroupIngress(params)

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

func ExampleElastiCache_CopySnapshot() {
	svc := elasticache.New(nil)

	params := &elasticache.CopySnapshotInput{
		SourceSnapshotName: aws.String("String"), // Required
		TargetSnapshotName: aws.String("String"), // Required
	}
	resp, err := svc.CopySnapshot(params)

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

func ExampleElastiCache_CreateCacheCluster() {
	svc := elasticache.New(nil)

	params := &elasticache.CreateCacheClusterInput{
		CacheClusterID:          aws.String("String"), // Required
		AZMode:                  aws.String("AZMode"),
		AutoMinorVersionUpgrade: aws.Boolean(true),
		CacheNodeType:           aws.String("String"),
		CacheParameterGroupName: aws.String("String"),
		CacheSecurityGroupNames: []*string{
			aws.String("String"), // Required
			// More values...
		},
		CacheSubnetGroupName: aws.String("String"),
		Engine:               aws.String("String"),
		EngineVersion:        aws.String("String"),
		NotificationTopicARN: aws.String("String"),
		NumCacheNodes:        aws.Long(1),
		Port:                 aws.Long(1),
		PreferredAvailabilityZone: aws.String("String"),
		PreferredAvailabilityZones: []*string{
			aws.String("String"), // Required
			// More values...
		},
		PreferredMaintenanceWindow: aws.String("String"),
		ReplicationGroupID:         aws.String("String"),
		SecurityGroupIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
		SnapshotARNs: []*string{
			aws.String("String"), // Required
			// More values...
		},
		SnapshotName:           aws.String("String"),
		SnapshotRetentionLimit: aws.Long(1),
		SnapshotWindow:         aws.String("String"),
		Tags: []*elasticache.Tag{
			&elasticache.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateCacheCluster(params)

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

func ExampleElastiCache_CreateCacheParameterGroup() {
	svc := elasticache.New(nil)

	params := &elasticache.CreateCacheParameterGroupInput{
		CacheParameterGroupFamily: aws.String("String"), // Required
		CacheParameterGroupName:   aws.String("String"), // Required
		Description:               aws.String("String"), // Required
	}
	resp, err := svc.CreateCacheParameterGroup(params)

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

func ExampleElastiCache_CreateCacheSecurityGroup() {
	svc := elasticache.New(nil)

	params := &elasticache.CreateCacheSecurityGroupInput{
		CacheSecurityGroupName: aws.String("String"), // Required
		Description:            aws.String("String"), // Required
	}
	resp, err := svc.CreateCacheSecurityGroup(params)

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

func ExampleElastiCache_CreateCacheSubnetGroup() {
	svc := elasticache.New(nil)

	params := &elasticache.CreateCacheSubnetGroupInput{
		CacheSubnetGroupDescription: aws.String("String"), // Required
		CacheSubnetGroupName:        aws.String("String"), // Required
		SubnetIDs: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.CreateCacheSubnetGroup(params)

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

func ExampleElastiCache_CreateReplicationGroup() {
	svc := elasticache.New(nil)

	params := &elasticache.CreateReplicationGroupInput{
		ReplicationGroupDescription: aws.String("String"), // Required
		ReplicationGroupID:          aws.String("String"), // Required
		AutoMinorVersionUpgrade:     aws.Boolean(true),
		AutomaticFailoverEnabled:    aws.Boolean(true),
		CacheNodeType:               aws.String("String"),
		CacheParameterGroupName:     aws.String("String"),
		CacheSecurityGroupNames: []*string{
			aws.String("String"), // Required
			// More values...
		},
		CacheSubnetGroupName: aws.String("String"),
		Engine:               aws.String("String"),
		EngineVersion:        aws.String("String"),
		NotificationTopicARN: aws.String("String"),
		NumCacheClusters:     aws.Long(1),
		Port:                 aws.Long(1),
		PreferredCacheClusterAZs: []*string{
			aws.String("String"), // Required
			// More values...
		},
		PreferredMaintenanceWindow: aws.String("String"),
		PrimaryClusterID:           aws.String("String"),
		SecurityGroupIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
		SnapshotARNs: []*string{
			aws.String("String"), // Required
			// More values...
		},
		SnapshotName:           aws.String("String"),
		SnapshotRetentionLimit: aws.Long(1),
		SnapshotWindow:         aws.String("String"),
		Tags: []*elasticache.Tag{
			&elasticache.Tag{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateReplicationGroup(params)

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

func ExampleElastiCache_CreateSnapshot() {
	svc := elasticache.New(nil)

	params := &elasticache.CreateSnapshotInput{
		CacheClusterID: aws.String("String"), // Required
		SnapshotName:   aws.String("String"), // Required
	}
	resp, err := svc.CreateSnapshot(params)

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

func ExampleElastiCache_DeleteCacheCluster() {
	svc := elasticache.New(nil)

	params := &elasticache.DeleteCacheClusterInput{
		CacheClusterID:          aws.String("String"), // Required
		FinalSnapshotIdentifier: aws.String("String"),
	}
	resp, err := svc.DeleteCacheCluster(params)

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

func ExampleElastiCache_DeleteCacheParameterGroup() {
	svc := elasticache.New(nil)

	params := &elasticache.DeleteCacheParameterGroupInput{
		CacheParameterGroupName: aws.String("String"), // Required
	}
	resp, err := svc.DeleteCacheParameterGroup(params)

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

func ExampleElastiCache_DeleteCacheSecurityGroup() {
	svc := elasticache.New(nil)

	params := &elasticache.DeleteCacheSecurityGroupInput{
		CacheSecurityGroupName: aws.String("String"), // Required
	}
	resp, err := svc.DeleteCacheSecurityGroup(params)

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

func ExampleElastiCache_DeleteCacheSubnetGroup() {
	svc := elasticache.New(nil)

	params := &elasticache.DeleteCacheSubnetGroupInput{
		CacheSubnetGroupName: aws.String("String"), // Required
	}
	resp, err := svc.DeleteCacheSubnetGroup(params)

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

func ExampleElastiCache_DeleteReplicationGroup() {
	svc := elasticache.New(nil)

	params := &elasticache.DeleteReplicationGroupInput{
		ReplicationGroupID:      aws.String("String"), // Required
		FinalSnapshotIdentifier: aws.String("String"),
		RetainPrimaryCluster:    aws.Boolean(true),
	}
	resp, err := svc.DeleteReplicationGroup(params)

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

func ExampleElastiCache_DeleteSnapshot() {
	svc := elasticache.New(nil)

	params := &elasticache.DeleteSnapshotInput{
		SnapshotName: aws.String("String"), // Required
	}
	resp, err := svc.DeleteSnapshot(params)

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

func ExampleElastiCache_DescribeCacheClusters() {
	svc := elasticache.New(nil)

	params := &elasticache.DescribeCacheClustersInput{
		CacheClusterID:    aws.String("String"),
		Marker:            aws.String("String"),
		MaxRecords:        aws.Long(1),
		ShowCacheNodeInfo: aws.Boolean(true),
	}
	resp, err := svc.DescribeCacheClusters(params)

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

func ExampleElastiCache_DescribeCacheEngineVersions() {
	svc := elasticache.New(nil)

	params := &elasticache.DescribeCacheEngineVersionsInput{
		CacheParameterGroupFamily: aws.String("String"),
		DefaultOnly:               aws.Boolean(true),
		Engine:                    aws.String("String"),
		EngineVersion:             aws.String("String"),
		Marker:                    aws.String("String"),
		MaxRecords:                aws.Long(1),
	}
	resp, err := svc.DescribeCacheEngineVersions(params)

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

func ExampleElastiCache_DescribeCacheParameterGroups() {
	svc := elasticache.New(nil)

	params := &elasticache.DescribeCacheParameterGroupsInput{
		CacheParameterGroupName: aws.String("String"),
		Marker:                  aws.String("String"),
		MaxRecords:              aws.Long(1),
	}
	resp, err := svc.DescribeCacheParameterGroups(params)

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

func ExampleElastiCache_DescribeCacheParameters() {
	svc := elasticache.New(nil)

	params := &elasticache.DescribeCacheParametersInput{
		CacheParameterGroupName: aws.String("String"), // Required
		Marker:                  aws.String("String"),
		MaxRecords:              aws.Long(1),
		Source:                  aws.String("String"),
	}
	resp, err := svc.DescribeCacheParameters(params)

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

func ExampleElastiCache_DescribeCacheSecurityGroups() {
	svc := elasticache.New(nil)

	params := &elasticache.DescribeCacheSecurityGroupsInput{
		CacheSecurityGroupName: aws.String("String"),
		Marker:                 aws.String("String"),
		MaxRecords:             aws.Long(1),
	}
	resp, err := svc.DescribeCacheSecurityGroups(params)

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

func ExampleElastiCache_DescribeCacheSubnetGroups() {
	svc := elasticache.New(nil)

	params := &elasticache.DescribeCacheSubnetGroupsInput{
		CacheSubnetGroupName: aws.String("String"),
		Marker:               aws.String("String"),
		MaxRecords:           aws.Long(1),
	}
	resp, err := svc.DescribeCacheSubnetGroups(params)

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

func ExampleElastiCache_DescribeEngineDefaultParameters() {
	svc := elasticache.New(nil)

	params := &elasticache.DescribeEngineDefaultParametersInput{
		CacheParameterGroupFamily: aws.String("String"), // Required
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

func ExampleElastiCache_DescribeEvents() {
	svc := elasticache.New(nil)

	params := &elasticache.DescribeEventsInput{
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

func ExampleElastiCache_DescribeReplicationGroups() {
	svc := elasticache.New(nil)

	params := &elasticache.DescribeReplicationGroupsInput{
		Marker:             aws.String("String"),
		MaxRecords:         aws.Long(1),
		ReplicationGroupID: aws.String("String"),
	}
	resp, err := svc.DescribeReplicationGroups(params)

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

func ExampleElastiCache_DescribeReservedCacheNodes() {
	svc := elasticache.New(nil)

	params := &elasticache.DescribeReservedCacheNodesInput{
		CacheNodeType:                aws.String("String"),
		Duration:                     aws.String("String"),
		Marker:                       aws.String("String"),
		MaxRecords:                   aws.Long(1),
		OfferingType:                 aws.String("String"),
		ProductDescription:           aws.String("String"),
		ReservedCacheNodeID:          aws.String("String"),
		ReservedCacheNodesOfferingID: aws.String("String"),
	}
	resp, err := svc.DescribeReservedCacheNodes(params)

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

func ExampleElastiCache_DescribeReservedCacheNodesOfferings() {
	svc := elasticache.New(nil)

	params := &elasticache.DescribeReservedCacheNodesOfferingsInput{
		CacheNodeType:                aws.String("String"),
		Duration:                     aws.String("String"),
		Marker:                       aws.String("String"),
		MaxRecords:                   aws.Long(1),
		OfferingType:                 aws.String("String"),
		ProductDescription:           aws.String("String"),
		ReservedCacheNodesOfferingID: aws.String("String"),
	}
	resp, err := svc.DescribeReservedCacheNodesOfferings(params)

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

func ExampleElastiCache_DescribeSnapshots() {
	svc := elasticache.New(nil)

	params := &elasticache.DescribeSnapshotsInput{
		CacheClusterID: aws.String("String"),
		Marker:         aws.String("String"),
		MaxRecords:     aws.Long(1),
		SnapshotName:   aws.String("String"),
		SnapshotSource: aws.String("String"),
	}
	resp, err := svc.DescribeSnapshots(params)

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

func ExampleElastiCache_ListTagsForResource() {
	svc := elasticache.New(nil)

	params := &elasticache.ListTagsForResourceInput{
		ResourceName: aws.String("String"), // Required
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

func ExampleElastiCache_ModifyCacheCluster() {
	svc := elasticache.New(nil)

	params := &elasticache.ModifyCacheClusterInput{
		CacheClusterID:          aws.String("String"), // Required
		AZMode:                  aws.String("AZMode"),
		ApplyImmediately:        aws.Boolean(true),
		AutoMinorVersionUpgrade: aws.Boolean(true),
		CacheNodeIDsToRemove: []*string{
			aws.String("String"), // Required
			// More values...
		},
		CacheParameterGroupName: aws.String("String"),
		CacheSecurityGroupNames: []*string{
			aws.String("String"), // Required
			// More values...
		},
		EngineVersion: aws.String("String"),
		NewAvailabilityZones: []*string{
			aws.String("String"), // Required
			// More values...
		},
		NotificationTopicARN:       aws.String("String"),
		NotificationTopicStatus:    aws.String("String"),
		NumCacheNodes:              aws.Long(1),
		PreferredMaintenanceWindow: aws.String("String"),
		SecurityGroupIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
		SnapshotRetentionLimit: aws.Long(1),
		SnapshotWindow:         aws.String("String"),
	}
	resp, err := svc.ModifyCacheCluster(params)

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

func ExampleElastiCache_ModifyCacheParameterGroup() {
	svc := elasticache.New(nil)

	params := &elasticache.ModifyCacheParameterGroupInput{
		CacheParameterGroupName: aws.String("String"), // Required
		ParameterNameValues: []*elasticache.ParameterNameValue{ // Required
			&elasticache.ParameterNameValue{ // Required
				ParameterName:  aws.String("String"),
				ParameterValue: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.ModifyCacheParameterGroup(params)

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

func ExampleElastiCache_ModifyCacheSubnetGroup() {
	svc := elasticache.New(nil)

	params := &elasticache.ModifyCacheSubnetGroupInput{
		CacheSubnetGroupName:        aws.String("String"), // Required
		CacheSubnetGroupDescription: aws.String("String"),
		SubnetIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.ModifyCacheSubnetGroup(params)

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

func ExampleElastiCache_ModifyReplicationGroup() {
	svc := elasticache.New(nil)

	params := &elasticache.ModifyReplicationGroupInput{
		ReplicationGroupID:       aws.String("String"), // Required
		ApplyImmediately:         aws.Boolean(true),
		AutoMinorVersionUpgrade:  aws.Boolean(true),
		AutomaticFailoverEnabled: aws.Boolean(true),
		CacheParameterGroupName:  aws.String("String"),
		CacheSecurityGroupNames: []*string{
			aws.String("String"), // Required
			// More values...
		},
		EngineVersion:               aws.String("String"),
		NotificationTopicARN:        aws.String("String"),
		NotificationTopicStatus:     aws.String("String"),
		PreferredMaintenanceWindow:  aws.String("String"),
		PrimaryClusterID:            aws.String("String"),
		ReplicationGroupDescription: aws.String("String"),
		SecurityGroupIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
		SnapshotRetentionLimit: aws.Long(1),
		SnapshotWindow:         aws.String("String"),
		SnapshottingClusterID:  aws.String("String"),
	}
	resp, err := svc.ModifyReplicationGroup(params)

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

func ExampleElastiCache_PurchaseReservedCacheNodesOffering() {
	svc := elasticache.New(nil)

	params := &elasticache.PurchaseReservedCacheNodesOfferingInput{
		ReservedCacheNodesOfferingID: aws.String("String"), // Required
		CacheNodeCount:               aws.Long(1),
		ReservedCacheNodeID:          aws.String("String"),
	}
	resp, err := svc.PurchaseReservedCacheNodesOffering(params)

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

func ExampleElastiCache_RebootCacheCluster() {
	svc := elasticache.New(nil)

	params := &elasticache.RebootCacheClusterInput{
		CacheClusterID: aws.String("String"), // Required
		CacheNodeIDsToReboot: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.RebootCacheCluster(params)

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

func ExampleElastiCache_RemoveTagsFromResource() {
	svc := elasticache.New(nil)

	params := &elasticache.RemoveTagsFromResourceInput{
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

func ExampleElastiCache_ResetCacheParameterGroup() {
	svc := elasticache.New(nil)

	params := &elasticache.ResetCacheParameterGroupInput{
		CacheParameterGroupName: aws.String("String"), // Required
		ParameterNameValues: []*elasticache.ParameterNameValue{ // Required
			&elasticache.ParameterNameValue{ // Required
				ParameterName:  aws.String("String"),
				ParameterValue: aws.String("String"),
			},
			// More values...
		},
		ResetAllParameters: aws.Boolean(true),
	}
	resp, err := svc.ResetCacheParameterGroup(params)

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

func ExampleElastiCache_RevokeCacheSecurityGroupIngress() {
	svc := elasticache.New(nil)

	params := &elasticache.RevokeCacheSecurityGroupIngressInput{
		CacheSecurityGroupName:  aws.String("String"), // Required
		EC2SecurityGroupName:    aws.String("String"), // Required
		EC2SecurityGroupOwnerID: aws.String("String"), // Required
	}
	resp, err := svc.RevokeCacheSecurityGroupIngress(params)

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