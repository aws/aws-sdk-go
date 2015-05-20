package storagegateway_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/storagegateway"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleStorageGateway_ActivateGateway() {
	svc := storagegateway.New(nil)

	params := &storagegateway.ActivateGatewayInput{
		ActivationKey:     aws.String("ActivationKey"),   // Required
		GatewayName:       aws.String("GatewayName"),     // Required
		GatewayRegion:     aws.String("RegionId"),        // Required
		GatewayTimezone:   aws.String("GatewayTimezone"), // Required
		GatewayType:       aws.String("GatewayType"),
		MediumChangerType: aws.String("MediumChangerType"),
		TapeDriveType:     aws.String("TapeDriveType"),
	}
	resp, err := svc.ActivateGateway(params)

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

func ExampleStorageGateway_AddCache() {
	svc := storagegateway.New(nil)

	params := &storagegateway.AddCacheInput{
		DiskIDs: []*string{ // Required
			aws.String("DiskId"), // Required
			// More values...
		},
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.AddCache(params)

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

func ExampleStorageGateway_AddUploadBuffer() {
	svc := storagegateway.New(nil)

	params := &storagegateway.AddUploadBufferInput{
		DiskIDs: []*string{ // Required
			aws.String("DiskId"), // Required
			// More values...
		},
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.AddUploadBuffer(params)

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

func ExampleStorageGateway_AddWorkingStorage() {
	svc := storagegateway.New(nil)

	params := &storagegateway.AddWorkingStorageInput{
		DiskIDs: []*string{ // Required
			aws.String("DiskId"), // Required
			// More values...
		},
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.AddWorkingStorage(params)

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

func ExampleStorageGateway_CancelArchival() {
	svc := storagegateway.New(nil)

	params := &storagegateway.CancelArchivalInput{
		GatewayARN: aws.String("GatewayARN"), // Required
		TapeARN:    aws.String("TapeARN"),    // Required
	}
	resp, err := svc.CancelArchival(params)

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

func ExampleStorageGateway_CancelRetrieval() {
	svc := storagegateway.New(nil)

	params := &storagegateway.CancelRetrievalInput{
		GatewayARN: aws.String("GatewayARN"), // Required
		TapeARN:    aws.String("TapeARN"),    // Required
	}
	resp, err := svc.CancelRetrieval(params)

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

func ExampleStorageGateway_CreateCachediSCSIVolume() {
	svc := storagegateway.New(nil)

	params := &storagegateway.CreateCachediSCSIVolumeInput{
		ClientToken:        aws.String("ClientToken"),        // Required
		GatewayARN:         aws.String("GatewayARN"),         // Required
		NetworkInterfaceID: aws.String("NetworkInterfaceId"), // Required
		TargetName:         aws.String("TargetName"),         // Required
		VolumeSizeInBytes:  aws.Long(1),                      // Required
		SnapshotID:         aws.String("SnapshotId"),
	}
	resp, err := svc.CreateCachediSCSIVolume(params)

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

func ExampleStorageGateway_CreateSnapshot() {
	svc := storagegateway.New(nil)

	params := &storagegateway.CreateSnapshotInput{
		SnapshotDescription: aws.String("SnapshotDescription"), // Required
		VolumeARN:           aws.String("VolumeARN"),           // Required
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

func ExampleStorageGateway_CreateSnapshotFromVolumeRecoveryPoint() {
	svc := storagegateway.New(nil)

	params := &storagegateway.CreateSnapshotFromVolumeRecoveryPointInput{
		SnapshotDescription: aws.String("SnapshotDescription"), // Required
		VolumeARN:           aws.String("VolumeARN"),           // Required
	}
	resp, err := svc.CreateSnapshotFromVolumeRecoveryPoint(params)

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

func ExampleStorageGateway_CreateStorediSCSIVolume() {
	svc := storagegateway.New(nil)

	params := &storagegateway.CreateStorediSCSIVolumeInput{
		DiskID:               aws.String("DiskId"),             // Required
		GatewayARN:           aws.String("GatewayARN"),         // Required
		NetworkInterfaceID:   aws.String("NetworkInterfaceId"), // Required
		PreserveExistingData: aws.Boolean(true),                // Required
		TargetName:           aws.String("TargetName"),         // Required
		SnapshotID:           aws.String("SnapshotId"),
	}
	resp, err := svc.CreateStorediSCSIVolume(params)

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

func ExampleStorageGateway_CreateTapes() {
	svc := storagegateway.New(nil)

	params := &storagegateway.CreateTapesInput{
		ClientToken:       aws.String("ClientToken"),       // Required
		GatewayARN:        aws.String("GatewayARN"),        // Required
		NumTapesToCreate:  aws.Long(1),                     // Required
		TapeBarcodePrefix: aws.String("TapeBarcodePrefix"), // Required
		TapeSizeInBytes:   aws.Long(1),                     // Required
	}
	resp, err := svc.CreateTapes(params)

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

func ExampleStorageGateway_DeleteBandwidthRateLimit() {
	svc := storagegateway.New(nil)

	params := &storagegateway.DeleteBandwidthRateLimitInput{
		BandwidthType: aws.String("BandwidthType"), // Required
		GatewayARN:    aws.String("GatewayARN"),    // Required
	}
	resp, err := svc.DeleteBandwidthRateLimit(params)

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

func ExampleStorageGateway_DeleteChapCredentials() {
	svc := storagegateway.New(nil)

	params := &storagegateway.DeleteChapCredentialsInput{
		InitiatorName: aws.String("IqnName"),   // Required
		TargetARN:     aws.String("TargetARN"), // Required
	}
	resp, err := svc.DeleteChapCredentials(params)

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

func ExampleStorageGateway_DeleteGateway() {
	svc := storagegateway.New(nil)

	params := &storagegateway.DeleteGatewayInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.DeleteGateway(params)

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

func ExampleStorageGateway_DeleteSnapshotSchedule() {
	svc := storagegateway.New(nil)

	params := &storagegateway.DeleteSnapshotScheduleInput{
		VolumeARN: aws.String("VolumeARN"), // Required
	}
	resp, err := svc.DeleteSnapshotSchedule(params)

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

func ExampleStorageGateway_DeleteTape() {
	svc := storagegateway.New(nil)

	params := &storagegateway.DeleteTapeInput{
		GatewayARN: aws.String("GatewayARN"), // Required
		TapeARN:    aws.String("TapeARN"),    // Required
	}
	resp, err := svc.DeleteTape(params)

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

func ExampleStorageGateway_DeleteTapeArchive() {
	svc := storagegateway.New(nil)

	params := &storagegateway.DeleteTapeArchiveInput{
		TapeARN: aws.String("TapeARN"), // Required
	}
	resp, err := svc.DeleteTapeArchive(params)

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

func ExampleStorageGateway_DeleteVolume() {
	svc := storagegateway.New(nil)

	params := &storagegateway.DeleteVolumeInput{
		VolumeARN: aws.String("VolumeARN"), // Required
	}
	resp, err := svc.DeleteVolume(params)

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

func ExampleStorageGateway_DescribeBandwidthRateLimit() {
	svc := storagegateway.New(nil)

	params := &storagegateway.DescribeBandwidthRateLimitInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.DescribeBandwidthRateLimit(params)

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

func ExampleStorageGateway_DescribeCache() {
	svc := storagegateway.New(nil)

	params := &storagegateway.DescribeCacheInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.DescribeCache(params)

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

func ExampleStorageGateway_DescribeCachediSCSIVolumes() {
	svc := storagegateway.New(nil)

	params := &storagegateway.DescribeCachediSCSIVolumesInput{
		VolumeARNs: []*string{ // Required
			aws.String("VolumeARN"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeCachediSCSIVolumes(params)

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

func ExampleStorageGateway_DescribeChapCredentials() {
	svc := storagegateway.New(nil)

	params := &storagegateway.DescribeChapCredentialsInput{
		TargetARN: aws.String("TargetARN"), // Required
	}
	resp, err := svc.DescribeChapCredentials(params)

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

func ExampleStorageGateway_DescribeGatewayInformation() {
	svc := storagegateway.New(nil)

	params := &storagegateway.DescribeGatewayInformationInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.DescribeGatewayInformation(params)

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

func ExampleStorageGateway_DescribeMaintenanceStartTime() {
	svc := storagegateway.New(nil)

	params := &storagegateway.DescribeMaintenanceStartTimeInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.DescribeMaintenanceStartTime(params)

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

func ExampleStorageGateway_DescribeSnapshotSchedule() {
	svc := storagegateway.New(nil)

	params := &storagegateway.DescribeSnapshotScheduleInput{
		VolumeARN: aws.String("VolumeARN"), // Required
	}
	resp, err := svc.DescribeSnapshotSchedule(params)

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

func ExampleStorageGateway_DescribeStorediSCSIVolumes() {
	svc := storagegateway.New(nil)

	params := &storagegateway.DescribeStorediSCSIVolumesInput{
		VolumeARNs: []*string{ // Required
			aws.String("VolumeARN"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeStorediSCSIVolumes(params)

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

func ExampleStorageGateway_DescribeTapeArchives() {
	svc := storagegateway.New(nil)

	params := &storagegateway.DescribeTapeArchivesInput{
		Limit:  aws.Long(1),
		Marker: aws.String("Marker"),
		TapeARNs: []*string{
			aws.String("TapeARN"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeTapeArchives(params)

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

func ExampleStorageGateway_DescribeTapeRecoveryPoints() {
	svc := storagegateway.New(nil)

	params := &storagegateway.DescribeTapeRecoveryPointsInput{
		GatewayARN: aws.String("GatewayARN"), // Required
		Limit:      aws.Long(1),
		Marker:     aws.String("Marker"),
	}
	resp, err := svc.DescribeTapeRecoveryPoints(params)

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

func ExampleStorageGateway_DescribeTapes() {
	svc := storagegateway.New(nil)

	params := &storagegateway.DescribeTapesInput{
		GatewayARN: aws.String("GatewayARN"), // Required
		Limit:      aws.Long(1),
		Marker:     aws.String("Marker"),
		TapeARNs: []*string{
			aws.String("TapeARN"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeTapes(params)

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

func ExampleStorageGateway_DescribeUploadBuffer() {
	svc := storagegateway.New(nil)

	params := &storagegateway.DescribeUploadBufferInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.DescribeUploadBuffer(params)

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

func ExampleStorageGateway_DescribeVTLDevices() {
	svc := storagegateway.New(nil)

	params := &storagegateway.DescribeVTLDevicesInput{
		GatewayARN: aws.String("GatewayARN"), // Required
		Limit:      aws.Long(1),
		Marker:     aws.String("Marker"),
		VTLDeviceARNs: []*string{
			aws.String("VTLDeviceARN"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeVTLDevices(params)

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

func ExampleStorageGateway_DescribeWorkingStorage() {
	svc := storagegateway.New(nil)

	params := &storagegateway.DescribeWorkingStorageInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.DescribeWorkingStorage(params)

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

func ExampleStorageGateway_DisableGateway() {
	svc := storagegateway.New(nil)

	params := &storagegateway.DisableGatewayInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.DisableGateway(params)

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

func ExampleStorageGateway_ListGateways() {
	svc := storagegateway.New(nil)

	params := &storagegateway.ListGatewaysInput{
		Limit:  aws.Long(1),
		Marker: aws.String("Marker"),
	}
	resp, err := svc.ListGateways(params)

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

func ExampleStorageGateway_ListLocalDisks() {
	svc := storagegateway.New(nil)

	params := &storagegateway.ListLocalDisksInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.ListLocalDisks(params)

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

func ExampleStorageGateway_ListVolumeRecoveryPoints() {
	svc := storagegateway.New(nil)

	params := &storagegateway.ListVolumeRecoveryPointsInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.ListVolumeRecoveryPoints(params)

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

func ExampleStorageGateway_ListVolumes() {
	svc := storagegateway.New(nil)

	params := &storagegateway.ListVolumesInput{
		GatewayARN: aws.String("GatewayARN"), // Required
		Limit:      aws.Long(1),
		Marker:     aws.String("Marker"),
	}
	resp, err := svc.ListVolumes(params)

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

func ExampleStorageGateway_ResetCache() {
	svc := storagegateway.New(nil)

	params := &storagegateway.ResetCacheInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.ResetCache(params)

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

func ExampleStorageGateway_RetrieveTapeArchive() {
	svc := storagegateway.New(nil)

	params := &storagegateway.RetrieveTapeArchiveInput{
		GatewayARN: aws.String("GatewayARN"), // Required
		TapeARN:    aws.String("TapeARN"),    // Required
	}
	resp, err := svc.RetrieveTapeArchive(params)

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

func ExampleStorageGateway_RetrieveTapeRecoveryPoint() {
	svc := storagegateway.New(nil)

	params := &storagegateway.RetrieveTapeRecoveryPointInput{
		GatewayARN: aws.String("GatewayARN"), // Required
		TapeARN:    aws.String("TapeARN"),    // Required
	}
	resp, err := svc.RetrieveTapeRecoveryPoint(params)

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

func ExampleStorageGateway_ShutdownGateway() {
	svc := storagegateway.New(nil)

	params := &storagegateway.ShutdownGatewayInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.ShutdownGateway(params)

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

func ExampleStorageGateway_StartGateway() {
	svc := storagegateway.New(nil)

	params := &storagegateway.StartGatewayInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.StartGateway(params)

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

func ExampleStorageGateway_UpdateBandwidthRateLimit() {
	svc := storagegateway.New(nil)

	params := &storagegateway.UpdateBandwidthRateLimitInput{
		GatewayARN:                           aws.String("GatewayARN"), // Required
		AverageDownloadRateLimitInBitsPerSec: aws.Long(1),
		AverageUploadRateLimitInBitsPerSec:   aws.Long(1),
	}
	resp, err := svc.UpdateBandwidthRateLimit(params)

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

func ExampleStorageGateway_UpdateChapCredentials() {
	svc := storagegateway.New(nil)

	params := &storagegateway.UpdateChapCredentialsInput{
		InitiatorName:                 aws.String("IqnName"),    // Required
		SecretToAuthenticateInitiator: aws.String("ChapSecret"), // Required
		TargetARN:                     aws.String("TargetARN"),  // Required
		SecretToAuthenticateTarget:    aws.String("ChapSecret"),
	}
	resp, err := svc.UpdateChapCredentials(params)

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

func ExampleStorageGateway_UpdateGatewayInformation() {
	svc := storagegateway.New(nil)

	params := &storagegateway.UpdateGatewayInformationInput{
		GatewayARN:      aws.String("GatewayARN"), // Required
		GatewayName:     aws.String("GatewayName"),
		GatewayTimezone: aws.String("GatewayTimezone"),
	}
	resp, err := svc.UpdateGatewayInformation(params)

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

func ExampleStorageGateway_UpdateGatewaySoftwareNow() {
	svc := storagegateway.New(nil)

	params := &storagegateway.UpdateGatewaySoftwareNowInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.UpdateGatewaySoftwareNow(params)

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

func ExampleStorageGateway_UpdateMaintenanceStartTime() {
	svc := storagegateway.New(nil)

	params := &storagegateway.UpdateMaintenanceStartTimeInput{
		DayOfWeek:    aws.Long(1),              // Required
		GatewayARN:   aws.String("GatewayARN"), // Required
		HourOfDay:    aws.Long(1),              // Required
		MinuteOfHour: aws.Long(1),              // Required
	}
	resp, err := svc.UpdateMaintenanceStartTime(params)

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

func ExampleStorageGateway_UpdateSnapshotSchedule() {
	svc := storagegateway.New(nil)

	params := &storagegateway.UpdateSnapshotScheduleInput{
		RecurrenceInHours: aws.Long(1),             // Required
		StartAt:           aws.Long(1),             // Required
		VolumeARN:         aws.String("VolumeARN"), // Required
		Description:       aws.String("Description"),
	}
	resp, err := svc.UpdateSnapshotSchedule(params)

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

func ExampleStorageGateway_UpdateVTLDeviceType() {
	svc := storagegateway.New(nil)

	params := &storagegateway.UpdateVTLDeviceTypeInput{
		DeviceType:   aws.String("DeviceType"),   // Required
		VTLDeviceARN: aws.String("VTLDeviceARN"), // Required
	}
	resp, err := svc.UpdateVTLDeviceType(params)

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