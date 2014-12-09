// Package storagegateway provides a client for AWS Storage Gateway.
package storagegateway

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// StorageGateway is a client for AWS Storage Gateway.
type StorageGateway struct {
	client aws.Client
}

// New returns a new StorageGateway client.
func New(key, secret, region string, client *http.Client) *StorageGateway {
	if client == nil {
		client = http.DefaultClient
	}

	return &StorageGateway{
		client: &aws.JSONClient{
			Signer: &aws.V4Signer{
				Key:     key,
				Secret:  secret,
				Service: "storagegateway",
				Region:  region,
				IncludeXAmzContentSha256: true,
			},
			Client:       client,
			Endpoint:     endpoints.Lookup("storagegateway", region),
			JSONVersion:  "1.1",
			TargetPrefix: "StorageGateway_20130630",
		},
	}
}

// ActivateGateway this operation activates the gateway you previously
// deployed on your host. For more information, see Activate the AWS
// Storage Gateway . In the activation process, you specify information
// such as the region you want to use for storing snapshots, the time zone
// for scheduled snapshots the gateway snapshot schedule window, an
// activation key, and a name for your gateway. The activation process also
// associates your gateway with your account; for more information, see
// UpdateGatewayInformation
func (c *StorageGateway) ActivateGateway(req ActivateGatewayInput) (resp *ActivateGatewayOutput, err error) {
	resp = &ActivateGatewayOutput{}
	err = c.client.Do("ActivateGateway", "POST", "/", req, resp)
	return
}

// AddCache this operation configures one or more gateway local disks as
// cache for a cached-volume gateway. This operation is supported only for
// the gateway-cached volume architecture (see Storage Gateway Concepts In
// the request, you specify the gateway Amazon Resource Name to which you
// want to add cache, and one or more disk IDs that you want to configure
// as cache.
func (c *StorageGateway) AddCache(req AddCacheInput) (resp *AddCacheOutput, err error) {
	resp = &AddCacheOutput{}
	err = c.client.Do("AddCache", "POST", "/", req, resp)
	return
}

// AddUploadBuffer this operation configures one or more gateway local
// disks as upload buffer for a specified gateway. This operation is
// supported for both the gateway-stored and gateway-cached volume
// architectures. In the request, you specify the gateway Amazon Resource
// Name to which you want to add upload buffer, and one or more disk IDs
// that you want to configure as upload buffer.
func (c *StorageGateway) AddUploadBuffer(req AddUploadBufferInput) (resp *AddUploadBufferOutput, err error) {
	resp = &AddUploadBufferOutput{}
	err = c.client.Do("AddUploadBuffer", "POST", "/", req, resp)
	return
}

// AddWorkingStorage this operation configures one or more gateway local
// disks as working storage for a gateway. This operation is supported only
// for the gateway-stored volume architecture. This operation is deprecated
// method in cached-volumes API version (20120630). Use AddUploadBuffer
// instead. In the request, you specify the gateway Amazon Resource Name to
// which you want to add working storage, and one or more disk IDs that you
// want to configure as working storage.
func (c *StorageGateway) AddWorkingStorage(req AddWorkingStorageInput) (resp *AddWorkingStorageOutput, err error) {
	resp = &AddWorkingStorageOutput{}
	err = c.client.Do("AddWorkingStorage", "POST", "/", req, resp)
	return
}

// CancelArchival cancels archiving of a virtual tape to the virtual tape
// shelf after the archiving process is initiated.
func (c *StorageGateway) CancelArchival(req CancelArchivalInput) (resp *CancelArchivalOutput, err error) {
	resp = &CancelArchivalOutput{}
	err = c.client.Do("CancelArchival", "POST", "/", req, resp)
	return
}

// CancelRetrieval cancels retrieval of a virtual tape from the virtual
// tape shelf to a gateway after the retrieval process is initiated. The
// virtual tape is returned to the
func (c *StorageGateway) CancelRetrieval(req CancelRetrievalInput) (resp *CancelRetrievalOutput, err error) {
	resp = &CancelRetrievalOutput{}
	err = c.client.Do("CancelRetrieval", "POST", "/", req, resp)
	return
}

// CreateCachediSCSIVolume this operation creates a cached volume on a
// specified cached gateway. This operation is supported only for the
// gateway-cached volume architecture. In the request, you must specify the
// gateway, size of the volume in bytes, the iSCSI target name, an IP
// address on which to expose the target, and a unique client token. In
// response, AWS Storage Gateway creates the volume and returns information
// about it such as the volume Amazon Resource Name its size, and the iSCSI
// target ARN that initiators can use to connect to the volume target.
func (c *StorageGateway) CreateCachediSCSIVolume(req CreateCachediSCSIVolumeInput) (resp *CreateCachediSCSIVolumeOutput, err error) {
	resp = &CreateCachediSCSIVolumeOutput{}
	err = c.client.Do("CreateCachediSCSIVolume", "POST", "/", req, resp)
	return
}

// CreateSnapshot this operation initiates a snapshot of a volume. AWS
// Storage Gateway provides the ability to back up point-in-time snapshots
// of your data to Amazon Simple Storage (S3) for durable off-site
// recovery, as well as import the data to an Amazon Elastic Block Store
// volume in Amazon Elastic Compute Cloud (EC2). You can take snapshots of
// your gateway volume on a scheduled or ad-hoc basis. This API enables you
// to take ad-hoc snapshot. For more information, see Working With
// Snapshots in the AWS Storage Gateway Console In the CreateSnapshot
// request you identify the volume by providing its Amazon Resource Name
// You must also provide description for the snapshot. When AWS Storage
// Gateway takes the snapshot of specified volume, the snapshot and
// description appears in the AWS Storage Gateway Console. In response, AWS
// Storage Gateway returns you a snapshot ID. You can use this snapshot ID
// to check the snapshot progress or later use it when you want to create a
// volume from a snapshot.
func (c *StorageGateway) CreateSnapshot(req CreateSnapshotInput) (resp *CreateSnapshotOutput, err error) {
	resp = &CreateSnapshotOutput{}
	err = c.client.Do("CreateSnapshot", "POST", "/", req, resp)
	return
}

// CreateSnapshotFromVolumeRecoveryPoint this operation initiates a
// snapshot of a gateway from a volume recovery point. This operation is
// supported only for the gateway-cached volume architecture (see A volume
// recovery point is a point in time at which all data of the volume is
// consistent and from which you can create a snapshot. To get a list of
// volume recovery point for gateway-cached volumes, use
// ListVolumeRecoveryPoints In the CreateSnapshotFromVolumeRecoveryPoint
// request, you identify the volume by providing its Amazon Resource Name
// You must also provide a description for the snapshot. When AWS Storage
// Gateway takes a snapshot of the specified volume, the snapshot and its
// description appear in the AWS Storage Gateway console. In response, AWS
// Storage Gateway returns you a snapshot ID. You can use this snapshot ID
// to check the snapshot progress or later use it when you want to create a
// volume from a snapshot.
func (c *StorageGateway) CreateSnapshotFromVolumeRecoveryPoint(req CreateSnapshotFromVolumeRecoveryPointInput) (resp *CreateSnapshotFromVolumeRecoveryPointOutput, err error) {
	resp = &CreateSnapshotFromVolumeRecoveryPointOutput{}
	err = c.client.Do("CreateSnapshotFromVolumeRecoveryPoint", "POST", "/", req, resp)
	return
}

// CreateStorediSCSIVolume this operation creates a volume on a specified
// gateway. This operation is supported only for the gateway-stored volume
// architecture. The size of the volume to create is inferred from the disk
// size. You can choose to preserve existing data on the disk, create
// volume from an existing snapshot, or create an empty volume. If you
// choose to create an empty gateway volume, then any existing data on the
// disk is erased. In the request you must specify the gateway and the disk
// information on which you are creating the volume. In response, AWS
// Storage Gateway creates the volume and returns volume information such
// as the volume Amazon Resource Name its size, and the iSCSI target ARN
// that initiators can use to connect to the volume target.
func (c *StorageGateway) CreateStorediSCSIVolume(req CreateStorediSCSIVolumeInput) (resp *CreateStorediSCSIVolumeOutput, err error) {
	resp = &CreateStorediSCSIVolumeOutput{}
	err = c.client.Do("CreateStorediSCSIVolume", "POST", "/", req, resp)
	return
}

// CreateTapes creates one or more virtual tapes. You write data to the
// virtual tapes and then archive the tapes.
func (c *StorageGateway) CreateTapes(req CreateTapesInput) (resp *CreateTapesOutput, err error) {
	resp = &CreateTapesOutput{}
	err = c.client.Do("CreateTapes", "POST", "/", req, resp)
	return
}

// DeleteBandwidthRateLimit this operation deletes the bandwidth rate
// limits of a gateway. You can delete either the upload and download
// bandwidth rate limit, or you can delete both. If you delete only one of
// the limits, the other limit remains unchanged. To specify which gateway
// to work with, use the Amazon Resource Name of the gateway in your
// request.
func (c *StorageGateway) DeleteBandwidthRateLimit(req DeleteBandwidthRateLimitInput) (resp *DeleteBandwidthRateLimitOutput, err error) {
	resp = &DeleteBandwidthRateLimitOutput{}
	err = c.client.Do("DeleteBandwidthRateLimit", "POST", "/", req, resp)
	return
}

// DeleteChapCredentials this operation deletes Challenge-Handshake
// Authentication Protocol credentials for a specified iSCSI target and
// initiator pair.
func (c *StorageGateway) DeleteChapCredentials(req DeleteChapCredentialsInput) (resp *DeleteChapCredentialsOutput, err error) {
	resp = &DeleteChapCredentialsOutput{}
	err = c.client.Do("DeleteChapCredentials", "POST", "/", req, resp)
	return
}

// DeleteGateway this operation deletes a gateway. To specify which gateway
// to delete, use the Amazon Resource Name of the gateway in your request.
// The operation deletes the gateway; however, it does not delete the
// gateway virtual machine from your host computer. After you delete a
// gateway, you cannot reactivate it. Completed snapshots of the gateway
// volumes are not deleted upon deleting the gateway, however, pending
// snapshots will not complete. After you delete a gateway, your next step
// is to remove it from your environment. You no longer pay software
// charges after the gateway is deleted; however, your existing Amazon EBS
// snapshots persist and you will continue to be billed for these
// snapshots. You can choose to remove all remaining Amazon EBS snapshots
// by canceling your Amazon EC2 subscription. If you prefer not to cancel
// your Amazon EC2 subscription, you can delete your snapshots using the
// Amazon EC2 console. For more information, see the AWS Storage Gateway
// Detail Page .
func (c *StorageGateway) DeleteGateway(req DeleteGatewayInput) (resp *DeleteGatewayOutput, err error) {
	resp = &DeleteGatewayOutput{}
	err = c.client.Do("DeleteGateway", "POST", "/", req, resp)
	return
}

// DeleteSnapshotSchedule this operation deletes a snapshot of a volume.
// You can take snapshots of your gateway volumes on a scheduled or ad-hoc
// basis. This API enables you to delete a snapshot schedule for a volume.
// For more information, see Working with Snapshots . In the
// DeleteSnapshotSchedule request, you identify the volume by providing its
// Amazon Resource Name
func (c *StorageGateway) DeleteSnapshotSchedule(req DeleteSnapshotScheduleInput) (resp *DeleteSnapshotScheduleOutput, err error) {
	resp = &DeleteSnapshotScheduleOutput{}
	err = c.client.Do("DeleteSnapshotSchedule", "POST", "/", req, resp)
	return
}

// DeleteTape is undocumented.
func (c *StorageGateway) DeleteTape(req DeleteTapeInput) (resp *DeleteTapeOutput, err error) {
	resp = &DeleteTapeOutput{}
	err = c.client.Do("DeleteTape", "POST", "/", req, resp)
	return
}

// DeleteTapeArchive deletes the specified virtual tape from the virtual
// tape shelf
func (c *StorageGateway) DeleteTapeArchive(req DeleteTapeArchiveInput) (resp *DeleteTapeArchiveOutput, err error) {
	resp = &DeleteTapeArchiveOutput{}
	err = c.client.Do("DeleteTapeArchive", "POST", "/", req, resp)
	return
}

// DeleteVolume this operation delete the specified gateway volume that you
// previously created using the CreateStorediSCSIVolume For gateway-stored
// volumes, the local disk that was configured as the storage volume is not
// deleted. You can reuse the local disk to create another storage volume.
// Before you delete a gateway volume, make sure there are no iSCSI
// connections to the volume you are deleting. You should also make sure
// there is no snapshot in progress. You can use the Amazon Elastic Compute
// Cloud (Amazon EC2) API to query snapshots on the volume you are deleting
// and check the snapshot status. For more information, go to
// DescribeSnapshots in the Amazon Elastic Compute Cloud API Reference In
// the request, you must provide the Amazon Resource Name of the storage
// volume you want to delete.
func (c *StorageGateway) DeleteVolume(req DeleteVolumeInput) (resp *DeleteVolumeOutput, err error) {
	resp = &DeleteVolumeOutput{}
	err = c.client.Do("DeleteVolume", "POST", "/", req, resp)
	return
}

// DescribeBandwidthRateLimit this operation returns the bandwidth rate
// limits of a gateway. By default, these limits are not set, which means
// no bandwidth rate limiting is in effect. This operation only returns a
// value for a bandwidth rate limit only if the limit is set. If no limits
// are set for the gateway, then this operation returns only the gateway
// ARN in the response body. To specify which gateway to describe, use the
// Amazon Resource Name of the gateway in your request.
func (c *StorageGateway) DescribeBandwidthRateLimit(req DescribeBandwidthRateLimitInput) (resp *DescribeBandwidthRateLimitOutput, err error) {
	resp = &DescribeBandwidthRateLimitOutput{}
	err = c.client.Do("DescribeBandwidthRateLimit", "POST", "/", req, resp)
	return
}

// DescribeCache this operation returns information about the cache of a
// gateway. This operation is supported only for the gateway-cached volume
// architecture. The response includes disk IDs that are configured as
// cache, and it includes the amount of cache allocated and used.
func (c *StorageGateway) DescribeCache(req DescribeCacheInput) (resp *DescribeCacheOutput, err error) {
	resp = &DescribeCacheOutput{}
	err = c.client.Do("DescribeCache", "POST", "/", req, resp)
	return
}

// DescribeCachediSCSIVolumes this operation returns a description of the
// gateway volumes specified in the request. This operation is supported
// only for the gateway-cached volume architecture. The list of gateway
// volumes in the request must be from one gateway. In the response Amazon
// Storage Gateway returns volume information sorted by volume Amazon
// Resource Name
func (c *StorageGateway) DescribeCachediSCSIVolumes(req DescribeCachediSCSIVolumesInput) (resp *DescribeCachediSCSIVolumesOutput, err error) {
	resp = &DescribeCachediSCSIVolumesOutput{}
	err = c.client.Do("DescribeCachediSCSIVolumes", "POST", "/", req, resp)
	return
}

// DescribeChapCredentials this operation returns an array of
// Challenge-Handshake Authentication Protocol credentials information for
// a specified iSCSI target, one for each target-initiator pair.
func (c *StorageGateway) DescribeChapCredentials(req DescribeChapCredentialsInput) (resp *DescribeChapCredentialsOutput, err error) {
	resp = &DescribeChapCredentialsOutput{}
	err = c.client.Do("DescribeChapCredentials", "POST", "/", req, resp)
	return
}

// DescribeGatewayInformation this operation returns metadata about a
// gateway such as its name, network interfaces, configured time zone, and
// the state (whether the gateway is running or not). To specify which
// gateway to describe, use the Amazon Resource Name of the gateway in your
// request.
func (c *StorageGateway) DescribeGatewayInformation(req DescribeGatewayInformationInput) (resp *DescribeGatewayInformationOutput, err error) {
	resp = &DescribeGatewayInformationOutput{}
	err = c.client.Do("DescribeGatewayInformation", "POST", "/", req, resp)
	return
}

// DescribeMaintenanceStartTime this operation returns your gateway's
// weekly maintenance start time including the day and time of the week.
// Note that values are in terms of the gateway's time zone.
func (c *StorageGateway) DescribeMaintenanceStartTime(req DescribeMaintenanceStartTimeInput) (resp *DescribeMaintenanceStartTimeOutput, err error) {
	resp = &DescribeMaintenanceStartTimeOutput{}
	err = c.client.Do("DescribeMaintenanceStartTime", "POST", "/", req, resp)
	return
}

// DescribeSnapshotSchedule this operation describes the snapshot schedule
// for the specified gateway volume. The snapshot schedule information
// includes intervals at which snapshots are automatically initiated on the
// volume.
func (c *StorageGateway) DescribeSnapshotSchedule(req DescribeSnapshotScheduleInput) (resp *DescribeSnapshotScheduleOutput, err error) {
	resp = &DescribeSnapshotScheduleOutput{}
	err = c.client.Do("DescribeSnapshotSchedule", "POST", "/", req, resp)
	return
}

// DescribeStorediSCSIVolumes this operation returns description of the
// gateway volumes specified in the request. The list of gateway volumes in
// the request must be from one gateway. In the response Amazon Storage
// Gateway returns volume information sorted by volume ARNs.
func (c *StorageGateway) DescribeStorediSCSIVolumes(req DescribeStorediSCSIVolumesInput) (resp *DescribeStorediSCSIVolumesOutput, err error) {
	resp = &DescribeStorediSCSIVolumesOutput{}
	err = c.client.Do("DescribeStorediSCSIVolumes", "POST", "/", req, resp)
	return
}

// DescribeTapeArchives returns a description of specified virtual tapes in
// the virtual tape shelf If a specific TapeARN is not specified, AWS
// Storage Gateway returns a description of all virtual tapes found in the
// VTS associated with your account.
func (c *StorageGateway) DescribeTapeArchives(req DescribeTapeArchivesInput) (resp *DescribeTapeArchivesOutput, err error) {
	resp = &DescribeTapeArchivesOutput{}
	err = c.client.Do("DescribeTapeArchives", "POST", "/", req, resp)
	return
}

// DescribeTapeRecoveryPoints returns a list of virtual tape recovery
// points that are available for the specified gateway-VTL. A recovery
// point is a point in time view of a virtual tape at which all the data on
// the virtual tape is consistent. If your gateway crashes, virtual tapes
// that have recovery points can be recovered to a new gateway.
func (c *StorageGateway) DescribeTapeRecoveryPoints(req DescribeTapeRecoveryPointsInput) (resp *DescribeTapeRecoveryPointsOutput, err error) {
	resp = &DescribeTapeRecoveryPointsOutput{}
	err = c.client.Do("DescribeTapeRecoveryPoints", "POST", "/", req, resp)
	return
}

// DescribeTapes returns a description of the specified Amazon Resource
// Name of virtual tapes. If a TapeARN is not specified, returns a
// description of all virtual tapes associated with the specified gateway.
func (c *StorageGateway) DescribeTapes(req DescribeTapesInput) (resp *DescribeTapesOutput, err error) {
	resp = &DescribeTapesOutput{}
	err = c.client.Do("DescribeTapes", "POST", "/", req, resp)
	return
}

// DescribeUploadBuffer this operation returns information about the upload
// buffer of a gateway. This operation is supported for both the
// gateway-stored and gateway-cached volume architectures. The response
// includes disk IDs that are configured as upload buffer space, and it
// includes the amount of upload buffer space allocated and used.
func (c *StorageGateway) DescribeUploadBuffer(req DescribeUploadBufferInput) (resp *DescribeUploadBufferOutput, err error) {
	resp = &DescribeUploadBufferOutput{}
	err = c.client.Do("DescribeUploadBuffer", "POST", "/", req, resp)
	return
}

// DescribeVTLDevices returns a description of virtual tape library devices
// for the specified gateway. In the response, AWS Storage Gateway returns
// VTL device information. The list of VTL devices must be from one
// gateway.
func (c *StorageGateway) DescribeVTLDevices(req DescribeVTLDevicesInput) (resp *DescribeVTLDevicesOutput, err error) {
	resp = &DescribeVTLDevicesOutput{}
	err = c.client.Do("DescribeVTLDevices", "POST", "/", req, resp)
	return
}

// DescribeWorkingStorage this operation returns information about the
// working storage of a gateway. This operation is supported only for the
// gateway-stored volume architecture. This operation is deprecated in
// cached-volumes API version (20120630). Use DescribeUploadBuffer instead.
// The response includes disk IDs that are configured as working storage,
// and it includes the amount of working storage allocated and used.
func (c *StorageGateway) DescribeWorkingStorage(req DescribeWorkingStorageInput) (resp *DescribeWorkingStorageOutput, err error) {
	resp = &DescribeWorkingStorageOutput{}
	err = c.client.Do("DescribeWorkingStorage", "POST", "/", req, resp)
	return
}

// DisableGateway disables a gateway when the gateway is no longer
// functioning. For example, if your gateway VM is damaged, you can disable
// the gateway so you can recover virtual tapes. Use this operation for a
// gateway-VTL that is not reachable or not functioning. Once a gateway is
// disabled it cannot be enabled.
func (c *StorageGateway) DisableGateway(req DisableGatewayInput) (resp *DisableGatewayOutput, err error) {
	resp = &DisableGatewayOutput{}
	err = c.client.Do("DisableGateway", "POST", "/", req, resp)
	return
}

// ListGateways this operation lists gateways owned by an AWS account in a
// region specified in the request. The returned list is ordered by gateway
// Amazon Resource Name By default, the operation returns a maximum of 100
// gateways. This operation supports pagination that allows you to
// optionally reduce the number of gateways returned in a response. If you
// have more gateways than are returned in a response-that is, the response
// returns only a truncated list of your gateways-the response contains a
// marker that you can specify in your next request to fetch the next page
// of gateways.
func (c *StorageGateway) ListGateways(req ListGatewaysInput) (resp *ListGatewaysOutput, err error) {
	resp = &ListGatewaysOutput{}
	err = c.client.Do("ListGateways", "POST", "/", req, resp)
	return
}

// ListLocalDisks this operation returns a list of the local disks of a
// gateway. To specify which gateway to describe you use the Amazon
// Resource Name of the gateway in the body of the request. The request
// returns all disks, specifying which are configured as working storage,
// stored volume or not configured at all.
func (c *StorageGateway) ListLocalDisks(req ListLocalDisksInput) (resp *ListLocalDisksOutput, err error) {
	resp = &ListLocalDisksOutput{}
	err = c.client.Do("ListLocalDisks", "POST", "/", req, resp)
	return
}

// ListVolumeRecoveryPoints this operation lists the recovery points for a
// specified gateway. This operation is supported only for the
// gateway-cached volume architecture. Each gateway-cached volume has one
// recovery point. A volume recovery point is a point in time at which all
// data of the volume is consistent and from which you can create a
// snapshot. To create a snapshot from a volume recovery point use the
// CreateSnapshotFromVolumeRecoveryPoint operation.
func (c *StorageGateway) ListVolumeRecoveryPoints(req ListVolumeRecoveryPointsInput) (resp *ListVolumeRecoveryPointsOutput, err error) {
	resp = &ListVolumeRecoveryPointsOutput{}
	err = c.client.Do("ListVolumeRecoveryPoints", "POST", "/", req, resp)
	return
}

// ListVolumes this operation lists the iSCSI stored volumes of a gateway.
// Results are sorted by volume The response includes only the volume ARNs.
// If you want additional volume information, use the
// DescribeStorediSCSIVolumes The operation supports pagination. By
// default, the operation returns a maximum of up to 100 volumes. You can
// optionally specify the Limit field in the body to limit the number of
// volumes in the response. If the number of volumes returned in the
// response is truncated, the response includes a Marker field. You can use
// this Marker value in your subsequent request to retrieve the next set of
// volumes.
func (c *StorageGateway) ListVolumes(req ListVolumesInput) (resp *ListVolumesOutput, err error) {
	resp = &ListVolumesOutput{}
	err = c.client.Do("ListVolumes", "POST", "/", req, resp)
	return
}

// RetrieveTapeArchive retrieves an archived virtual tape from the virtual
// tape shelf to a gateway-VTL. Virtual tapes archived in the VTS are not
// associated with any gateway. However after a tape is retrieved, it is
// associated with a gateway, even though it is also listed in the Once a
// tape is successfully retrieved to a gateway, it cannot be retrieved
// again to another gateway. You must archive the tape again before you can
// retrieve it to another gateway.
func (c *StorageGateway) RetrieveTapeArchive(req RetrieveTapeArchiveInput) (resp *RetrieveTapeArchiveOutput, err error) {
	resp = &RetrieveTapeArchiveOutput{}
	err = c.client.Do("RetrieveTapeArchive", "POST", "/", req, resp)
	return
}

// RetrieveTapeRecoveryPoint retrieves the recovery point for the specified
// virtual tape. A recovery point is a point in time view of a virtual tape
// at which all the data on the tape is consistent. If your gateway
// crashes, virtual tapes that have recovery points can be recovered to a
// new gateway.
func (c *StorageGateway) RetrieveTapeRecoveryPoint(req RetrieveTapeRecoveryPointInput) (resp *RetrieveTapeRecoveryPointOutput, err error) {
	resp = &RetrieveTapeRecoveryPointOutput{}
	err = c.client.Do("RetrieveTapeRecoveryPoint", "POST", "/", req, resp)
	return
}

// ShutdownGateway this operation shuts down a gateway. To specify which
// gateway to shut down, use the Amazon Resource Name of the gateway in the
// body of your request. The operation shuts down the gateway service
// component running in the storage gateway's virtual machine and not the
// After the gateway is shutdown, you cannot call any other API except
// StartGateway , DescribeGatewayInformation , and ListGateways . For more
// information, see ActivateGateway . Your applications cannot read from or
// write to the gateway's storage volumes, and there are no snapshots
// taken. If do not intend to use the gateway again, you must delete the
// gateway (using DeleteGateway ) to no longer pay software charges
// associated with the gateway.
func (c *StorageGateway) ShutdownGateway(req ShutdownGatewayInput) (resp *ShutdownGatewayOutput, err error) {
	resp = &ShutdownGatewayOutput{}
	err = c.client.Do("ShutdownGateway", "POST", "/", req, resp)
	return
}

// StartGateway this operation starts a gateway that you previously shut
// down (see ShutdownGateway ). After the gateway starts, you can then make
// other API calls, your applications can read from or write to the
// gateway's storage volumes and you will be able to take snapshot backups.
// To specify which gateway to start, use the Amazon Resource Name of the
// gateway in your request.
func (c *StorageGateway) StartGateway(req StartGatewayInput) (resp *StartGatewayOutput, err error) {
	resp = &StartGatewayOutput{}
	err = c.client.Do("StartGateway", "POST", "/", req, resp)
	return
}

// UpdateBandwidthRateLimit this operation updates the bandwidth rate
// limits of a gateway. You can update both the upload and download
// bandwidth rate limit or specify only one of the two. If you don't set a
// bandwidth rate limit, the existing rate limit remains. By default, a
// gateway's bandwidth rate limits are not set. If you don't set any limit,
// the gateway does not have any limitations on its bandwidth usage and
// could potentially use the maximum available bandwidth. To specify which
// gateway to update, use the Amazon Resource Name of the gateway in your
// request.
func (c *StorageGateway) UpdateBandwidthRateLimit(req UpdateBandwidthRateLimitInput) (resp *UpdateBandwidthRateLimitOutput, err error) {
	resp = &UpdateBandwidthRateLimitOutput{}
	err = c.client.Do("UpdateBandwidthRateLimit", "POST", "/", req, resp)
	return
}

// UpdateChapCredentials this operation updates the Challenge-Handshake
// Authentication Protocol credentials for a specified iSCSI target. By
// default, a gateway does not have enabled; however, for added security,
// you might use it. When you update credentials, all existing connections
// on the target are closed and initiators must reconnect with the new
// credentials.
func (c *StorageGateway) UpdateChapCredentials(req UpdateChapCredentialsInput) (resp *UpdateChapCredentialsOutput, err error) {
	resp = &UpdateChapCredentialsOutput{}
	err = c.client.Do("UpdateChapCredentials", "POST", "/", req, resp)
	return
}

// UpdateGatewayInformation this operation updates a gateway's metadata,
// which includes the gateway's name and time zone. To specify which
// gateway to update, use the Amazon Resource Name of the gateway in your
// request.
func (c *StorageGateway) UpdateGatewayInformation(req UpdateGatewayInformationInput) (resp *UpdateGatewayInformationOutput, err error) {
	resp = &UpdateGatewayInformationOutput{}
	err = c.client.Do("UpdateGatewayInformation", "POST", "/", req, resp)
	return
}

// UpdateGatewaySoftwareNow this operation updates the gateway virtual
// machine software. The request immediately triggers the software update.
// A software update forces a system restart of your gateway. You can
// minimize the chance of any disruption to your applications by increasing
// your iSCSI Initiators' timeouts. For more information about increasing
// iSCSI Initiator timeouts for Windows and Linux, see Customizing Your
// Windows iSCSI Settings and Customizing Your Linux iSCSI Settings ,
// respectively.
func (c *StorageGateway) UpdateGatewaySoftwareNow(req UpdateGatewaySoftwareNowInput) (resp *UpdateGatewaySoftwareNowOutput, err error) {
	resp = &UpdateGatewaySoftwareNowOutput{}
	err = c.client.Do("UpdateGatewaySoftwareNow", "POST", "/", req, resp)
	return
}

// UpdateMaintenanceStartTime this operation updates a gateway's weekly
// maintenance start time information, including day and time of the week.
// The maintenance time is the time in your gateway's time zone.
func (c *StorageGateway) UpdateMaintenanceStartTime(req UpdateMaintenanceStartTimeInput) (resp *UpdateMaintenanceStartTimeOutput, err error) {
	resp = &UpdateMaintenanceStartTimeOutput{}
	err = c.client.Do("UpdateMaintenanceStartTime", "POST", "/", req, resp)
	return
}

// UpdateSnapshotSchedule this operation updates a snapshot schedule
// configured for a gateway volume. The default snapshot schedule for
// volume is once every 24 hours, starting at the creation time of the
// volume. You can use this API to change the snapshot schedule configured
// for the volume. In the request you must identify the gateway volume
// whose snapshot schedule you want to update, and the schedule
// information, including when you want the snapshot to begin on a day and
// the frequency (in hours) of snapshots.
func (c *StorageGateway) UpdateSnapshotSchedule(req UpdateSnapshotScheduleInput) (resp *UpdateSnapshotScheduleOutput, err error) {
	resp = &UpdateSnapshotScheduleOutput{}
	err = c.client.Do("UpdateSnapshotSchedule", "POST", "/", req, resp)
	return
}

// ActivateGatewayInput is undocumented.
type ActivateGatewayInput struct {
	ActivationKey     string `json:"ActivationKey"`
	GatewayName       string `json:"GatewayName"`
	GatewayRegion     string `json:"GatewayRegion"`
	GatewayTimezone   string `json:"GatewayTimezone"`
	GatewayType       string `json:"GatewayType,omitempty"`
	MediumChangerType string `json:"MediumChangerType,omitempty"`
	TapeDriveType     string `json:"TapeDriveType,omitempty"`
}

// ActivateGatewayOutput is undocumented.
type ActivateGatewayOutput struct {
	GatewayARN string `json:"GatewayARN,omitempty"`
}

// AddCacheInput is undocumented.
type AddCacheInput struct {
	DiskIds    []string `json:"DiskIds"`
	GatewayARN string   `json:"GatewayARN"`
}

// AddCacheOutput is undocumented.
type AddCacheOutput struct {
	GatewayARN string `json:"GatewayARN,omitempty"`
}

// AddUploadBufferInput is undocumented.
type AddUploadBufferInput struct {
	DiskIds    []string `json:"DiskIds"`
	GatewayARN string   `json:"GatewayARN"`
}

// AddUploadBufferOutput is undocumented.
type AddUploadBufferOutput struct {
	GatewayARN string `json:"GatewayARN,omitempty"`
}

// AddWorkingStorageInput is undocumented.
type AddWorkingStorageInput struct {
	DiskIds    []string `json:"DiskIds"`
	GatewayARN string   `json:"GatewayARN"`
}

// AddWorkingStorageOutput is undocumented.
type AddWorkingStorageOutput struct {
	GatewayARN string `json:"GatewayARN,omitempty"`
}

// CachediSCSIVolume is undocumented.
type CachediSCSIVolume struct {
	SourceSnapshotID      string                `json:"SourceSnapshotId,omitempty"`
	VolumeARN             string                `json:"VolumeARN,omitempty"`
	VolumeID              string                `json:"VolumeId,omitempty"`
	VolumeProgress        float64               `json:"VolumeProgress,omitempty"`
	VolumeSizeInBytes     int64                 `json:"VolumeSizeInBytes,omitempty"`
	VolumeStatus          string                `json:"VolumeStatus,omitempty"`
	VolumeType            string                `json:"VolumeType,omitempty"`
	VolumeiSCSIAttributes VolumeiSCSIAttributes `json:"VolumeiSCSIAttributes,omitempty"`
}

// CancelArchivalInput is undocumented.
type CancelArchivalInput struct {
	GatewayARN string `json:"GatewayARN"`
	TapeARN    string `json:"TapeARN"`
}

// CancelArchivalOutput is undocumented.
type CancelArchivalOutput struct {
	TapeARN string `json:"TapeARN,omitempty"`
}

// CancelRetrievalInput is undocumented.
type CancelRetrievalInput struct {
	GatewayARN string `json:"GatewayARN"`
	TapeARN    string `json:"TapeARN"`
}

// CancelRetrievalOutput is undocumented.
type CancelRetrievalOutput struct {
	TapeARN string `json:"TapeARN,omitempty"`
}

// ChapInfo is undocumented.
type ChapInfo struct {
	InitiatorName                 string `json:"InitiatorName,omitempty"`
	SecretToAuthenticateInitiator string `json:"SecretToAuthenticateInitiator,omitempty"`
	SecretToAuthenticateTarget    string `json:"SecretToAuthenticateTarget,omitempty"`
	TargetARN                     string `json:"TargetARN,omitempty"`
}

// CreateCachediSCSIVolumeInput is undocumented.
type CreateCachediSCSIVolumeInput struct {
	ClientToken        string `json:"ClientToken"`
	GatewayARN         string `json:"GatewayARN"`
	NetworkInterfaceID string `json:"NetworkInterfaceId"`
	SnapshotID         string `json:"SnapshotId,omitempty"`
	TargetName         string `json:"TargetName"`
	VolumeSizeInBytes  int64  `json:"VolumeSizeInBytes"`
}

// CreateCachediSCSIVolumeOutput is undocumented.
type CreateCachediSCSIVolumeOutput struct {
	TargetARN string `json:"TargetARN,omitempty"`
	VolumeARN string `json:"VolumeARN,omitempty"`
}

// CreateSnapshotFromVolumeRecoveryPointInput is undocumented.
type CreateSnapshotFromVolumeRecoveryPointInput struct {
	SnapshotDescription string `json:"SnapshotDescription"`
	VolumeARN           string `json:"VolumeARN"`
}

// CreateSnapshotFromVolumeRecoveryPointOutput is undocumented.
type CreateSnapshotFromVolumeRecoveryPointOutput struct {
	SnapshotID              string `json:"SnapshotId,omitempty"`
	VolumeARN               string `json:"VolumeARN,omitempty"`
	VolumeRecoveryPointTime string `json:"VolumeRecoveryPointTime,omitempty"`
}

// CreateSnapshotInput is undocumented.
type CreateSnapshotInput struct {
	SnapshotDescription string `json:"SnapshotDescription"`
	VolumeARN           string `json:"VolumeARN"`
}

// CreateSnapshotOutput is undocumented.
type CreateSnapshotOutput struct {
	SnapshotID string `json:"SnapshotId,omitempty"`
	VolumeARN  string `json:"VolumeARN,omitempty"`
}

// CreateStorediSCSIVolumeInput is undocumented.
type CreateStorediSCSIVolumeInput struct {
	DiskID               string `json:"DiskId"`
	GatewayARN           string `json:"GatewayARN"`
	NetworkInterfaceID   string `json:"NetworkInterfaceId"`
	PreserveExistingData bool   `json:"PreserveExistingData"`
	SnapshotID           string `json:"SnapshotId,omitempty"`
	TargetName           string `json:"TargetName"`
}

// CreateStorediSCSIVolumeOutput is undocumented.
type CreateStorediSCSIVolumeOutput struct {
	TargetARN         string `json:"TargetARN,omitempty"`
	VolumeARN         string `json:"VolumeARN,omitempty"`
	VolumeSizeInBytes int64  `json:"VolumeSizeInBytes,omitempty"`
}

// CreateTapesInput is undocumented.
type CreateTapesInput struct {
	ClientToken       string `json:"ClientToken"`
	GatewayARN        string `json:"GatewayARN"`
	NumTapesToCreate  int    `json:"NumTapesToCreate"`
	TapeBarcodePrefix string `json:"TapeBarcodePrefix"`
	TapeSizeInBytes   int64  `json:"TapeSizeInBytes"`
}

// CreateTapesOutput is undocumented.
type CreateTapesOutput struct {
	TapeARNs []string `json:"TapeARNs,omitempty"`
}

// DeleteBandwidthRateLimitInput is undocumented.
type DeleteBandwidthRateLimitInput struct {
	BandwidthType string `json:"BandwidthType"`
	GatewayARN    string `json:"GatewayARN"`
}

// DeleteBandwidthRateLimitOutput is undocumented.
type DeleteBandwidthRateLimitOutput struct {
	GatewayARN string `json:"GatewayARN,omitempty"`
}

// DeleteChapCredentialsInput is undocumented.
type DeleteChapCredentialsInput struct {
	InitiatorName string `json:"InitiatorName"`
	TargetARN     string `json:"TargetARN"`
}

// DeleteChapCredentialsOutput is undocumented.
type DeleteChapCredentialsOutput struct {
	InitiatorName string `json:"InitiatorName,omitempty"`
	TargetARN     string `json:"TargetARN,omitempty"`
}

// DeleteGatewayInput is undocumented.
type DeleteGatewayInput struct {
	GatewayARN string `json:"GatewayARN"`
}

// DeleteGatewayOutput is undocumented.
type DeleteGatewayOutput struct {
	GatewayARN string `json:"GatewayARN,omitempty"`
}

// DeleteSnapshotScheduleInput is undocumented.
type DeleteSnapshotScheduleInput struct {
	VolumeARN string `json:"VolumeARN"`
}

// DeleteSnapshotScheduleOutput is undocumented.
type DeleteSnapshotScheduleOutput struct {
	VolumeARN string `json:"VolumeARN,omitempty"`
}

// DeleteTapeArchiveInput is undocumented.
type DeleteTapeArchiveInput struct {
	TapeARN string `json:"TapeARN"`
}

// DeleteTapeArchiveOutput is undocumented.
type DeleteTapeArchiveOutput struct {
	TapeARN string `json:"TapeARN,omitempty"`
}

// DeleteTapeInput is undocumented.
type DeleteTapeInput struct {
	GatewayARN string `json:"GatewayARN"`
	TapeARN    string `json:"TapeARN"`
}

// DeleteTapeOutput is undocumented.
type DeleteTapeOutput struct {
	TapeARN string `json:"TapeARN,omitempty"`
}

// DeleteVolumeInput is undocumented.
type DeleteVolumeInput struct {
	VolumeARN string `json:"VolumeARN"`
}

// DeleteVolumeOutput is undocumented.
type DeleteVolumeOutput struct {
	VolumeARN string `json:"VolumeARN,omitempty"`
}

// DescribeBandwidthRateLimitInput is undocumented.
type DescribeBandwidthRateLimitInput struct {
	GatewayARN string `json:"GatewayARN"`
}

// DescribeBandwidthRateLimitOutput is undocumented.
type DescribeBandwidthRateLimitOutput struct {
	AverageDownloadRateLimitInBitsPerSec int64  `json:"AverageDownloadRateLimitInBitsPerSec,omitempty"`
	AverageUploadRateLimitInBitsPerSec   int64  `json:"AverageUploadRateLimitInBitsPerSec,omitempty"`
	GatewayARN                           string `json:"GatewayARN,omitempty"`
}

// DescribeCacheInput is undocumented.
type DescribeCacheInput struct {
	GatewayARN string `json:"GatewayARN"`
}

// DescribeCacheOutput is undocumented.
type DescribeCacheOutput struct {
	CacheAllocatedInBytes int64    `json:"CacheAllocatedInBytes,omitempty"`
	CacheDirtyPercentage  float64  `json:"CacheDirtyPercentage,omitempty"`
	CacheHitPercentage    float64  `json:"CacheHitPercentage,omitempty"`
	CacheMissPercentage   float64  `json:"CacheMissPercentage,omitempty"`
	CacheUsedPercentage   float64  `json:"CacheUsedPercentage,omitempty"`
	DiskIds               []string `json:"DiskIds,omitempty"`
	GatewayARN            string   `json:"GatewayARN,omitempty"`
}

// DescribeCachediSCSIVolumesInput is undocumented.
type DescribeCachediSCSIVolumesInput struct {
	VolumeARNs []string `json:"VolumeARNs"`
}

// DescribeCachediSCSIVolumesOutput is undocumented.
type DescribeCachediSCSIVolumesOutput struct {
	CachediSCSIVolumes []CachediSCSIVolume `json:"CachediSCSIVolumes,omitempty"`
}

// DescribeChapCredentialsInput is undocumented.
type DescribeChapCredentialsInput struct {
	TargetARN string `json:"TargetARN"`
}

// DescribeChapCredentialsOutput is undocumented.
type DescribeChapCredentialsOutput struct {
	ChapCredentials []ChapInfo `json:"ChapCredentials,omitempty"`
}

// DescribeGatewayInformationInput is undocumented.
type DescribeGatewayInformationInput struct {
	GatewayARN string `json:"GatewayARN"`
}

// DescribeGatewayInformationOutput is undocumented.
type DescribeGatewayInformationOutput struct {
	GatewayARN                 string             `json:"GatewayARN,omitempty"`
	GatewayID                  string             `json:"GatewayId,omitempty"`
	GatewayNetworkInterfaces   []NetworkInterface `json:"GatewayNetworkInterfaces,omitempty"`
	GatewayState               string             `json:"GatewayState,omitempty"`
	GatewayTimezone            string             `json:"GatewayTimezone,omitempty"`
	GatewayType                string             `json:"GatewayType,omitempty"`
	NextUpdateAvailabilityDate string             `json:"NextUpdateAvailabilityDate,omitempty"`
}

// DescribeMaintenanceStartTimeInput is undocumented.
type DescribeMaintenanceStartTimeInput struct {
	GatewayARN string `json:"GatewayARN"`
}

// DescribeMaintenanceStartTimeOutput is undocumented.
type DescribeMaintenanceStartTimeOutput struct {
	DayOfWeek    int    `json:"DayOfWeek,omitempty"`
	GatewayARN   string `json:"GatewayARN,omitempty"`
	HourOfDay    int    `json:"HourOfDay,omitempty"`
	MinuteOfHour int    `json:"MinuteOfHour,omitempty"`
	Timezone     string `json:"Timezone,omitempty"`
}

// DescribeSnapshotScheduleInput is undocumented.
type DescribeSnapshotScheduleInput struct {
	VolumeARN string `json:"VolumeARN"`
}

// DescribeSnapshotScheduleOutput is undocumented.
type DescribeSnapshotScheduleOutput struct {
	Description       string `json:"Description,omitempty"`
	RecurrenceInHours int    `json:"RecurrenceInHours,omitempty"`
	StartAt           int    `json:"StartAt,omitempty"`
	Timezone          string `json:"Timezone,omitempty"`
	VolumeARN         string `json:"VolumeARN,omitempty"`
}

// DescribeStorediSCSIVolumesInput is undocumented.
type DescribeStorediSCSIVolumesInput struct {
	VolumeARNs []string `json:"VolumeARNs"`
}

// DescribeStorediSCSIVolumesOutput is undocumented.
type DescribeStorediSCSIVolumesOutput struct {
	StorediSCSIVolumes []StorediSCSIVolume `json:"StorediSCSIVolumes,omitempty"`
}

// DescribeTapeArchivesInput is undocumented.
type DescribeTapeArchivesInput struct {
	Limit    int      `json:"Limit,omitempty"`
	Marker   string   `json:"Marker,omitempty"`
	TapeARNs []string `json:"TapeARNs,omitempty"`
}

// DescribeTapeArchivesOutput is undocumented.
type DescribeTapeArchivesOutput struct {
	Marker       string        `json:"Marker,omitempty"`
	TapeArchives []TapeArchive `json:"TapeArchives,omitempty"`
}

// DescribeTapeRecoveryPointsInput is undocumented.
type DescribeTapeRecoveryPointsInput struct {
	GatewayARN string `json:"GatewayARN"`
	Limit      int    `json:"Limit,omitempty"`
	Marker     string `json:"Marker,omitempty"`
}

// DescribeTapeRecoveryPointsOutput is undocumented.
type DescribeTapeRecoveryPointsOutput struct {
	GatewayARN             string                  `json:"GatewayARN,omitempty"`
	Marker                 string                  `json:"Marker,omitempty"`
	TapeRecoveryPointInfos []TapeRecoveryPointInfo `json:"TapeRecoveryPointInfos,omitempty"`
}

// DescribeTapesInput is undocumented.
type DescribeTapesInput struct {
	GatewayARN string   `json:"GatewayARN"`
	Limit      int      `json:"Limit,omitempty"`
	Marker     string   `json:"Marker,omitempty"`
	TapeARNs   []string `json:"TapeARNs,omitempty"`
}

// DescribeTapesOutput is undocumented.
type DescribeTapesOutput struct {
	Marker string `json:"Marker,omitempty"`
	Tapes  []Tape `json:"Tapes,omitempty"`
}

// DescribeUploadBufferInput is undocumented.
type DescribeUploadBufferInput struct {
	GatewayARN string `json:"GatewayARN"`
}

// DescribeUploadBufferOutput is undocumented.
type DescribeUploadBufferOutput struct {
	DiskIds                      []string `json:"DiskIds,omitempty"`
	GatewayARN                   string   `json:"GatewayARN,omitempty"`
	UploadBufferAllocatedInBytes int64    `json:"UploadBufferAllocatedInBytes,omitempty"`
	UploadBufferUsedInBytes      int64    `json:"UploadBufferUsedInBytes,omitempty"`
}

// DescribeVTLDevicesInput is undocumented.
type DescribeVTLDevicesInput struct {
	GatewayARN    string   `json:"GatewayARN"`
	Limit         int      `json:"Limit,omitempty"`
	Marker        string   `json:"Marker,omitempty"`
	VTLDeviceARNs []string `json:"VTLDeviceARNs,omitempty"`
}

// DescribeVTLDevicesOutput is undocumented.
type DescribeVTLDevicesOutput struct {
	GatewayARN string      `json:"GatewayARN,omitempty"`
	Marker     string      `json:"Marker,omitempty"`
	VTLDevices []VTLDevice `json:"VTLDevices,omitempty"`
}

// DescribeWorkingStorageInput is undocumented.
type DescribeWorkingStorageInput struct {
	GatewayARN string `json:"GatewayARN"`
}

// DescribeWorkingStorageOutput is undocumented.
type DescribeWorkingStorageOutput struct {
	DiskIds                        []string `json:"DiskIds,omitempty"`
	GatewayARN                     string   `json:"GatewayARN,omitempty"`
	WorkingStorageAllocatedInBytes int64    `json:"WorkingStorageAllocatedInBytes,omitempty"`
	WorkingStorageUsedInBytes      int64    `json:"WorkingStorageUsedInBytes,omitempty"`
}

// DeviceiSCSIAttributes is undocumented.
type DeviceiSCSIAttributes struct {
	ChapEnabled          bool   `json:"ChapEnabled,omitempty"`
	NetworkInterfaceID   string `json:"NetworkInterfaceId,omitempty"`
	NetworkInterfacePort int    `json:"NetworkInterfacePort,omitempty"`
	TargetARN            string `json:"TargetARN,omitempty"`
}

// DisableGatewayInput is undocumented.
type DisableGatewayInput struct {
	GatewayARN string `json:"GatewayARN"`
}

// DisableGatewayOutput is undocumented.
type DisableGatewayOutput struct {
	GatewayARN string `json:"GatewayARN,omitempty"`
}

// Disk is undocumented.
type Disk struct {
	DiskAllocationResource string `json:"DiskAllocationResource,omitempty"`
	DiskAllocationType     string `json:"DiskAllocationType,omitempty"`
	DiskID                 string `json:"DiskId,omitempty"`
	DiskNode               string `json:"DiskNode,omitempty"`
	DiskPath               string `json:"DiskPath,omitempty"`
	DiskSizeInBytes        int64  `json:"DiskSizeInBytes,omitempty"`
}

// GatewayInfo is undocumented.
type GatewayInfo struct {
	GatewayARN              string `json:"GatewayARN,omitempty"`
	GatewayOperationalState string `json:"GatewayOperationalState,omitempty"`
	GatewayType             string `json:"GatewayType,omitempty"`
}

// ListGatewaysInput is undocumented.
type ListGatewaysInput struct {
	Limit  int    `json:"Limit,omitempty"`
	Marker string `json:"Marker,omitempty"`
}

// ListGatewaysOutput is undocumented.
type ListGatewaysOutput struct {
	Gateways []GatewayInfo `json:"Gateways,omitempty"`
	Marker   string        `json:"Marker,omitempty"`
}

// ListLocalDisksInput is undocumented.
type ListLocalDisksInput struct {
	GatewayARN string `json:"GatewayARN"`
}

// ListLocalDisksOutput is undocumented.
type ListLocalDisksOutput struct {
	Disks      []Disk `json:"Disks,omitempty"`
	GatewayARN string `json:"GatewayARN,omitempty"`
}

// ListVolumeRecoveryPointsInput is undocumented.
type ListVolumeRecoveryPointsInput struct {
	GatewayARN string `json:"GatewayARN"`
}

// ListVolumeRecoveryPointsOutput is undocumented.
type ListVolumeRecoveryPointsOutput struct {
	GatewayARN               string                    `json:"GatewayARN,omitempty"`
	VolumeRecoveryPointInfos []VolumeRecoveryPointInfo `json:"VolumeRecoveryPointInfos,omitempty"`
}

// ListVolumesInput is undocumented.
type ListVolumesInput struct {
	GatewayARN string `json:"GatewayARN"`
	Limit      int    `json:"Limit,omitempty"`
	Marker     string `json:"Marker,omitempty"`
}

// ListVolumesOutput is undocumented.
type ListVolumesOutput struct {
	GatewayARN  string       `json:"GatewayARN,omitempty"`
	Marker      string       `json:"Marker,omitempty"`
	VolumeInfos []VolumeInfo `json:"VolumeInfos,omitempty"`
}

// NetworkInterface is undocumented.
type NetworkInterface struct {
	IPv4Address string `json:"Ipv4Address,omitempty"`
	IPv6Address string `json:"Ipv6Address,omitempty"`
	MacAddress  string `json:"MacAddress,omitempty"`
}

// RetrieveTapeArchiveInput is undocumented.
type RetrieveTapeArchiveInput struct {
	GatewayARN string `json:"GatewayARN"`
	TapeARN    string `json:"TapeARN"`
}

// RetrieveTapeArchiveOutput is undocumented.
type RetrieveTapeArchiveOutput struct {
	TapeARN string `json:"TapeARN,omitempty"`
}

// RetrieveTapeRecoveryPointInput is undocumented.
type RetrieveTapeRecoveryPointInput struct {
	GatewayARN string `json:"GatewayARN"`
	TapeARN    string `json:"TapeARN"`
}

// RetrieveTapeRecoveryPointOutput is undocumented.
type RetrieveTapeRecoveryPointOutput struct {
	TapeARN string `json:"TapeARN,omitempty"`
}

// ShutdownGatewayInput is undocumented.
type ShutdownGatewayInput struct {
	GatewayARN string `json:"GatewayARN"`
}

// ShutdownGatewayOutput is undocumented.
type ShutdownGatewayOutput struct {
	GatewayARN string `json:"GatewayARN,omitempty"`
}

// StartGatewayInput is undocumented.
type StartGatewayInput struct {
	GatewayARN string `json:"GatewayARN"`
}

// StartGatewayOutput is undocumented.
type StartGatewayOutput struct {
	GatewayARN string `json:"GatewayARN,omitempty"`
}

// StorageGatewayError is undocumented.
type StorageGatewayError struct {
	ErrorCode    string            `json:"errorCode,omitempty"`
	ErrorDetails map[string]string `json:"errorDetails,omitempty"`
}

// StorediSCSIVolume is undocumented.
type StorediSCSIVolume struct {
	PreservedExistingData bool                  `json:"PreservedExistingData,omitempty"`
	SourceSnapshotID      string                `json:"SourceSnapshotId,omitempty"`
	VolumeARN             string                `json:"VolumeARN,omitempty"`
	VolumeDiskID          string                `json:"VolumeDiskId,omitempty"`
	VolumeID              string                `json:"VolumeId,omitempty"`
	VolumeProgress        float64               `json:"VolumeProgress,omitempty"`
	VolumeSizeInBytes     int64                 `json:"VolumeSizeInBytes,omitempty"`
	VolumeStatus          string                `json:"VolumeStatus,omitempty"`
	VolumeType            string                `json:"VolumeType,omitempty"`
	VolumeiSCSIAttributes VolumeiSCSIAttributes `json:"VolumeiSCSIAttributes,omitempty"`
}

// Tape is undocumented.
type Tape struct {
	Progress        float64 `json:"Progress,omitempty"`
	TapeARN         string  `json:"TapeARN,omitempty"`
	TapeBarcode     string  `json:"TapeBarcode,omitempty"`
	TapeSizeInBytes int64   `json:"TapeSizeInBytes,omitempty"`
	TapeStatus      string  `json:"TapeStatus,omitempty"`
	VTLDevice       string  `json:"VTLDevice,omitempty"`
}

// TapeArchive is undocumented.
type TapeArchive struct {
	CompletionTime  time.Time `json:"CompletionTime,omitempty"`
	RetrievedTo     string    `json:"RetrievedTo,omitempty"`
	TapeARN         string    `json:"TapeARN,omitempty"`
	TapeBarcode     string    `json:"TapeBarcode,omitempty"`
	TapeSizeInBytes int64     `json:"TapeSizeInBytes,omitempty"`
	TapeStatus      string    `json:"TapeStatus,omitempty"`
}

// TapeRecoveryPointInfo is undocumented.
type TapeRecoveryPointInfo struct {
	TapeARN               string    `json:"TapeARN,omitempty"`
	TapeRecoveryPointTime time.Time `json:"TapeRecoveryPointTime,omitempty"`
	TapeSizeInBytes       int64     `json:"TapeSizeInBytes,omitempty"`
	TapeStatus            string    `json:"TapeStatus,omitempty"`
}

// UpdateBandwidthRateLimitInput is undocumented.
type UpdateBandwidthRateLimitInput struct {
	AverageDownloadRateLimitInBitsPerSec int64  `json:"AverageDownloadRateLimitInBitsPerSec,omitempty"`
	AverageUploadRateLimitInBitsPerSec   int64  `json:"AverageUploadRateLimitInBitsPerSec,omitempty"`
	GatewayARN                           string `json:"GatewayARN"`
}

// UpdateBandwidthRateLimitOutput is undocumented.
type UpdateBandwidthRateLimitOutput struct {
	GatewayARN string `json:"GatewayARN,omitempty"`
}

// UpdateChapCredentialsInput is undocumented.
type UpdateChapCredentialsInput struct {
	InitiatorName                 string `json:"InitiatorName"`
	SecretToAuthenticateInitiator string `json:"SecretToAuthenticateInitiator"`
	SecretToAuthenticateTarget    string `json:"SecretToAuthenticateTarget,omitempty"`
	TargetARN                     string `json:"TargetARN"`
}

// UpdateChapCredentialsOutput is undocumented.
type UpdateChapCredentialsOutput struct {
	InitiatorName string `json:"InitiatorName,omitempty"`
	TargetARN     string `json:"TargetARN,omitempty"`
}

// UpdateGatewayInformationInput is undocumented.
type UpdateGatewayInformationInput struct {
	GatewayARN      string `json:"GatewayARN"`
	GatewayName     string `json:"GatewayName,omitempty"`
	GatewayTimezone string `json:"GatewayTimezone,omitempty"`
}

// UpdateGatewayInformationOutput is undocumented.
type UpdateGatewayInformationOutput struct {
	GatewayARN string `json:"GatewayARN,omitempty"`
}

// UpdateGatewaySoftwareNowInput is undocumented.
type UpdateGatewaySoftwareNowInput struct {
	GatewayARN string `json:"GatewayARN"`
}

// UpdateGatewaySoftwareNowOutput is undocumented.
type UpdateGatewaySoftwareNowOutput struct {
	GatewayARN string `json:"GatewayARN,omitempty"`
}

// UpdateMaintenanceStartTimeInput is undocumented.
type UpdateMaintenanceStartTimeInput struct {
	DayOfWeek    int    `json:"DayOfWeek"`
	GatewayARN   string `json:"GatewayARN"`
	HourOfDay    int    `json:"HourOfDay"`
	MinuteOfHour int    `json:"MinuteOfHour"`
}

// UpdateMaintenanceStartTimeOutput is undocumented.
type UpdateMaintenanceStartTimeOutput struct {
	GatewayARN string `json:"GatewayARN,omitempty"`
}

// UpdateSnapshotScheduleInput is undocumented.
type UpdateSnapshotScheduleInput struct {
	Description       string `json:"Description,omitempty"`
	RecurrenceInHours int    `json:"RecurrenceInHours"`
	StartAt           int    `json:"StartAt"`
	VolumeARN         string `json:"VolumeARN"`
}

// UpdateSnapshotScheduleOutput is undocumented.
type UpdateSnapshotScheduleOutput struct {
	VolumeARN string `json:"VolumeARN,omitempty"`
}

// VTLDevice is undocumented.
type VTLDevice struct {
	DeviceiSCSIAttributes      DeviceiSCSIAttributes `json:"DeviceiSCSIAttributes,omitempty"`
	VTLDeviceARN               string                `json:"VTLDeviceARN,omitempty"`
	VTLDeviceProductIdentifier string                `json:"VTLDeviceProductIdentifier,omitempty"`
	VTLDeviceType              string                `json:"VTLDeviceType,omitempty"`
	VTLDeviceVendor            string                `json:"VTLDeviceVendor,omitempty"`
}

// VolumeInfo is undocumented.
type VolumeInfo struct {
	VolumeARN  string `json:"VolumeARN,omitempty"`
	VolumeType string `json:"VolumeType,omitempty"`
}

// VolumeRecoveryPointInfo is undocumented.
type VolumeRecoveryPointInfo struct {
	VolumeARN               string `json:"VolumeARN,omitempty"`
	VolumeRecoveryPointTime string `json:"VolumeRecoveryPointTime,omitempty"`
	VolumeSizeInBytes       int64  `json:"VolumeSizeInBytes,omitempty"`
	VolumeUsageInBytes      int64  `json:"VolumeUsageInBytes,omitempty"`
}

// VolumeiSCSIAttributes is undocumented.
type VolumeiSCSIAttributes struct {
	ChapEnabled          bool   `json:"ChapEnabled,omitempty"`
	LunNumber            int    `json:"LunNumber,omitempty"`
	NetworkInterfaceID   string `json:"NetworkInterfaceId,omitempty"`
	NetworkInterfacePort int    `json:"NetworkInterfacePort,omitempty"`
	TargetARN            string `json:"TargetARN,omitempty"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
