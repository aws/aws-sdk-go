// Package rds provides a client for Amazon Relational Database Service.
package rds

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// RDS is a client for Amazon Relational Database Service.
type RDS struct {
	client *aws.QueryClient
}

// New returns a new RDS client.
func New(key, secret, region string, client *http.Client) *RDS {
	if client == nil {
		client = http.DefaultClient
	}

	return &RDS{
		client: &aws.QueryClient{
			Signer: &aws.V4Signer{
				Key:     key,
				Secret:  secret,
				Service: "rds",
				Region:  region,
				IncludeXAmzContentSha256: true,
			},
			Client:     client,
			Endpoint:   endpoints.Lookup("rds", region),
			APIVersion: "2014-09-01",
		},
	}
}

// AddSourceIdentifierToSubscription adds a source identifier to an
// existing RDS event notification subscription.
func (c *RDS) AddSourceIdentifierToSubscription(req AddSourceIdentifierToSubscriptionMessage) (resp *AddSourceIdentifierToSubscriptionResult, err error) {
	resp = &AddSourceIdentifierToSubscriptionResult{}
	err = c.client.Do("AddSourceIdentifierToSubscription", "POST", "/", req, resp)
	return
}

// AddTagsToResource adds metadata tags to an Amazon RDS resource. These
// tags can also be used with cost allocation reporting to track cost
// associated with Amazon RDS resources, or used in Condition statement in
// IAM policy for Amazon For an overview on tagging Amazon RDS resources,
// see Tagging Amazon RDS Resources
func (c *RDS) AddTagsToResource(req AddTagsToResourceMessage) (err error) {
	// NRE
	err = c.client.Do("AddTagsToResource", "POST", "/", req, nil)
	return
}

// AuthorizeDBSecurityGroupIngress enables ingress to a DBSecurityGroup
// using one of two forms of authorization. First, EC2 or VPC security
// groups can be added to the DBSecurityGroup if the application using the
// database is running on EC2 or VPC instances. Second, IP ranges are
// available if the application accessing your database is running on the
// Internet. Required parameters for this API are one of range,
// EC2SecurityGroupId for or (EC2SecurityGroupOwnerId and either
// EC2SecurityGroupName or EC2SecurityGroupId for non-VPC). For an overview
// of ranges, go to the Wikipedia Tutorial .
func (c *RDS) AuthorizeDBSecurityGroupIngress(req AuthorizeDBSecurityGroupIngressMessage) (resp *AuthorizeDBSecurityGroupIngressResult, err error) {
	resp = &AuthorizeDBSecurityGroupIngressResult{}
	err = c.client.Do("AuthorizeDBSecurityGroupIngress", "POST", "/", req, resp)
	return
}

// CopyDBParameterGroup is undocumented.
func (c *RDS) CopyDBParameterGroup(req CopyDBParameterGroupMessage) (resp *CopyDBParameterGroupResult, err error) {
	resp = &CopyDBParameterGroupResult{}
	err = c.client.Do("CopyDBParameterGroup", "POST", "/", req, resp)
	return
}

// CopyDBSnapshot copies the specified DBSnapshot. The source DBSnapshot
// must be in the "available" state.
func (c *RDS) CopyDBSnapshot(req CopyDBSnapshotMessage) (resp *CopyDBSnapshotResult, err error) {
	resp = &CopyDBSnapshotResult{}
	err = c.client.Do("CopyDBSnapshot", "POST", "/", req, resp)
	return
}

// CopyOptionGroup is undocumented.
func (c *RDS) CopyOptionGroup(req CopyOptionGroupMessage) (resp *CopyOptionGroupResult, err error) {
	resp = &CopyOptionGroupResult{}
	err = c.client.Do("CopyOptionGroup", "POST", "/", req, resp)
	return
}

// CreateDBInstance is undocumented.
func (c *RDS) CreateDBInstance(req CreateDBInstanceMessage) (resp *CreateDBInstanceResult, err error) {
	resp = &CreateDBInstanceResult{}
	err = c.client.Do("CreateDBInstance", "POST", "/", req, resp)
	return
}

// CreateDBInstanceReadReplica creates a DB instance that acts as a read
// replica of a source DB instance. All read replica DB instances are
// created as Single-AZ deployments with backups disabled. All other DB
// instance attributes (including DB security groups and DB parameter
// groups) are inherited from the source DB instance, except as specified
// below. The source DB instance must have backup retention enabled.
func (c *RDS) CreateDBInstanceReadReplica(req CreateDBInstanceReadReplicaMessage) (resp *CreateDBInstanceReadReplicaResult, err error) {
	resp = &CreateDBInstanceReadReplicaResult{}
	err = c.client.Do("CreateDBInstanceReadReplica", "POST", "/", req, resp)
	return
}

// CreateDBParameterGroup creates a new DB parameter group. A DB parameter
// group is initially created with the default parameters for the database
// engine used by the DB instance. To provide custom values for any of the
// parameters, you must modify the group after creating it using
// ModifyDBParameterGroup . Once you've created a DB parameter group, you
// need to associate it with your DB instance using ModifyDBInstance . When
// you associate a new DB parameter group with a running DB instance, you
// need to reboot the DB instance without failover for the new DB parameter
// group and associated settings to take effect. After you create a DB
// parameter group, you should wait at least 5 minutes before creating your
// first DB instance that uses that DB parameter group as the default
// parameter group. This allows Amazon RDS to fully complete the create
// action before the parameter group is used as the default for a new DB
// instance. This is especially important for parameters that are critical
// when creating the default database for a DB instance, such as the
// character set for the default database defined by the
// character_set_database parameter. You can use the Parameter Groups
// option of the Amazon RDS console or the DescribeDBParameters command to
// verify that your DB parameter group has been created or modified.
func (c *RDS) CreateDBParameterGroup(req CreateDBParameterGroupMessage) (resp *CreateDBParameterGroupResult, err error) {
	resp = &CreateDBParameterGroupResult{}
	err = c.client.Do("CreateDBParameterGroup", "POST", "/", req, resp)
	return
}

// CreateDBSecurityGroup creates a new DB security group. DB security
// groups control access to a DB instance.
func (c *RDS) CreateDBSecurityGroup(req CreateDBSecurityGroupMessage) (resp *CreateDBSecurityGroupResult, err error) {
	resp = &CreateDBSecurityGroupResult{}
	err = c.client.Do("CreateDBSecurityGroup", "POST", "/", req, resp)
	return
}

// CreateDBSnapshot creates a DBSnapshot. The source DBInstance must be in
// "available" state.
func (c *RDS) CreateDBSnapshot(req CreateDBSnapshotMessage) (resp *CreateDBSnapshotResult, err error) {
	resp = &CreateDBSnapshotResult{}
	err = c.client.Do("CreateDBSnapshot", "POST", "/", req, resp)
	return
}

// CreateDBSubnetGroup creates a new DB subnet group. DB subnet groups must
// contain at least one subnet in at least two AZs in the region.
func (c *RDS) CreateDBSubnetGroup(req CreateDBSubnetGroupMessage) (resp *CreateDBSubnetGroupResult, err error) {
	resp = &CreateDBSubnetGroupResult{}
	err = c.client.Do("CreateDBSubnetGroup", "POST", "/", req, resp)
	return
}

// CreateEventSubscription creates an RDS event notification subscription.
// This action requires a topic ARN (Amazon Resource Name) created by
// either the RDS console, the SNS console, or the SNS To obtain an ARN
// with you must create a topic in Amazon SNS and subscribe to the topic.
// The ARN is displayed in the SNS console. You can specify the type of
// source (SourceType) you want to be notified of, provide a list of RDS
// sources (SourceIds) that triggers the events, and provide a list of
// event categories (EventCategories) for events you want to be notified
// of. For example, you can specify SourceType = db-instance, SourceIds =
// mydbinstance1, mydbinstance2 and EventCategories = Availability, Backup.
// If you specify both the SourceType and SourceIds, such as SourceType =
// db-instance and SourceIdentifier = myDBInstance1, you will be notified
// of all the db-instance events for the specified source. If you specify a
// SourceType but do not specify a SourceIdentifier, you will receive
// notice of the events for that source type for all your RDS sources. If
// you do not specify either the SourceType nor the SourceIdentifier, you
// will be notified of events generated from all RDS sources belonging to
// your customer account.
func (c *RDS) CreateEventSubscription(req CreateEventSubscriptionMessage) (resp *CreateEventSubscriptionResult, err error) {
	resp = &CreateEventSubscriptionResult{}
	err = c.client.Do("CreateEventSubscription", "POST", "/", req, resp)
	return
}

// CreateOptionGroup creates a new option group. You can create up to 20
// option groups.
func (c *RDS) CreateOptionGroup(req CreateOptionGroupMessage) (resp *CreateOptionGroupResult, err error) {
	resp = &CreateOptionGroupResult{}
	err = c.client.Do("CreateOptionGroup", "POST", "/", req, resp)
	return
}

// DeleteDBInstance the DeleteDBInstance action deletes a previously
// provisioned DB instance. A successful response from the web service
// indicates the request was received correctly. When you delete a DB
// instance, all automated backups for that instance are deleted and cannot
// be recovered. Manual DB snapshots of the DB instance to be deleted are
// not deleted. If a final DB snapshot is requested the status of the RDS
// instance will be "deleting" until the DB snapshot is created. The API
// action DescribeDBInstance is used to monitor the status of this
// operation. The action cannot be canceled or reverted once submitted.
func (c *RDS) DeleteDBInstance(req DeleteDBInstanceMessage) (resp *DeleteDBInstanceResult, err error) {
	resp = &DeleteDBInstanceResult{}
	err = c.client.Do("DeleteDBInstance", "POST", "/", req, resp)
	return
}

// DeleteDBParameterGroup deletes a specified DBParameterGroup. The
// DBParameterGroup to be deleted cannot be associated with any DB
// instances.
func (c *RDS) DeleteDBParameterGroup(req DeleteDBParameterGroupMessage) (err error) {
	// NRE
	err = c.client.Do("DeleteDBParameterGroup", "POST", "/", req, nil)
	return
}

// DeleteDBSecurityGroup is undocumented.
func (c *RDS) DeleteDBSecurityGroup(req DeleteDBSecurityGroupMessage) (err error) {
	// NRE
	err = c.client.Do("DeleteDBSecurityGroup", "POST", "/", req, nil)
	return
}

// DeleteDBSnapshot deletes a DBSnapshot. If the snapshot is being copied,
// the copy operation is terminated.
func (c *RDS) DeleteDBSnapshot(req DeleteDBSnapshotMessage) (resp *DeleteDBSnapshotResult, err error) {
	resp = &DeleteDBSnapshotResult{}
	err = c.client.Do("DeleteDBSnapshot", "POST", "/", req, resp)
	return
}

// DeleteDBSubnetGroup is undocumented.
func (c *RDS) DeleteDBSubnetGroup(req DeleteDBSubnetGroupMessage) (err error) {
	// NRE
	err = c.client.Do("DeleteDBSubnetGroup", "POST", "/", req, nil)
	return
}

// DeleteEventSubscription is undocumented.
func (c *RDS) DeleteEventSubscription(req DeleteEventSubscriptionMessage) (resp *DeleteEventSubscriptionResult, err error) {
	resp = &DeleteEventSubscriptionResult{}
	err = c.client.Do("DeleteEventSubscription", "POST", "/", req, resp)
	return
}

// DeleteOptionGroup is undocumented.
func (c *RDS) DeleteOptionGroup(req DeleteOptionGroupMessage) (err error) {
	// NRE
	err = c.client.Do("DeleteOptionGroup", "POST", "/", req, nil)
	return
}

// DescribeDBEngineVersions is undocumented.
func (c *RDS) DescribeDBEngineVersions(req DescribeDBEngineVersionsMessage) (resp *DescribeDBEngineVersionsResult, err error) {
	resp = &DescribeDBEngineVersionsResult{}
	err = c.client.Do("DescribeDBEngineVersions", "POST", "/", req, resp)
	return
}

// DescribeDBInstances returns information about provisioned RDS instances.
// This API supports pagination.
func (c *RDS) DescribeDBInstances(req DescribeDBInstancesMessage) (resp *DescribeDBInstancesResult, err error) {
	resp = &DescribeDBInstancesResult{}
	err = c.client.Do("DescribeDBInstances", "POST", "/", req, resp)
	return
}

// DescribeDBLogFiles returns a list of DB log files for the DB instance.
func (c *RDS) DescribeDBLogFiles(req DescribeDBLogFilesMessage) (resp *DescribeDBLogFilesResult, err error) {
	resp = &DescribeDBLogFilesResult{}
	err = c.client.Do("DescribeDBLogFiles", "POST", "/", req, resp)
	return
}

// DescribeDBParameterGroups returns a list of DBParameterGroup
// descriptions. If a DBParameterGroupName is specified, the list will
// contain only the description of the specified DB parameter group.
func (c *RDS) DescribeDBParameterGroups(req DescribeDBParameterGroupsMessage) (resp *DescribeDBParameterGroupsResult, err error) {
	resp = &DescribeDBParameterGroupsResult{}
	err = c.client.Do("DescribeDBParameterGroups", "POST", "/", req, resp)
	return
}

// DescribeDBParameters returns the detailed parameter list for a
// particular DB parameter group.
func (c *RDS) DescribeDBParameters(req DescribeDBParametersMessage) (resp *DescribeDBParametersResult, err error) {
	resp = &DescribeDBParametersResult{}
	err = c.client.Do("DescribeDBParameters", "POST", "/", req, resp)
	return
}

// DescribeDBSecurityGroups returns a list of DBSecurityGroup descriptions.
// If a DBSecurityGroupName is specified, the list will contain only the
// descriptions of the specified DB security group.
func (c *RDS) DescribeDBSecurityGroups(req DescribeDBSecurityGroupsMessage) (resp *DescribeDBSecurityGroupsResult, err error) {
	resp = &DescribeDBSecurityGroupsResult{}
	err = c.client.Do("DescribeDBSecurityGroups", "POST", "/", req, resp)
	return
}

// DescribeDBSnapshots returns information about DB snapshots. This API
// supports pagination.
func (c *RDS) DescribeDBSnapshots(req DescribeDBSnapshotsMessage) (resp *DescribeDBSnapshotsResult, err error) {
	resp = &DescribeDBSnapshotsResult{}
	err = c.client.Do("DescribeDBSnapshots", "POST", "/", req, resp)
	return
}

// DescribeDBSubnetGroups returns a list of DBSubnetGroup descriptions. If
// a DBSubnetGroupName is specified, the list will contain only the
// descriptions of the specified DBSubnetGroup. For an overview of ranges,
// go to the Wikipedia Tutorial .
func (c *RDS) DescribeDBSubnetGroups(req DescribeDBSubnetGroupsMessage) (resp *DescribeDBSubnetGroupsResult, err error) {
	resp = &DescribeDBSubnetGroupsResult{}
	err = c.client.Do("DescribeDBSubnetGroups", "POST", "/", req, resp)
	return
}

// DescribeEngineDefaultParameters returns the default engine and system
// parameter information for the specified database engine.
func (c *RDS) DescribeEngineDefaultParameters(req DescribeEngineDefaultParametersMessage) (resp *DescribeEngineDefaultParametersResult, err error) {
	resp = &DescribeEngineDefaultParametersResult{}
	err = c.client.Do("DescribeEngineDefaultParameters", "POST", "/", req, resp)
	return
}

// DescribeEventCategories displays a list of categories for all event
// source types, or, if specified, for a specified source type. You can see
// a list of the event categories and source types in the Events topic in
// the Amazon RDS User Guide.
func (c *RDS) DescribeEventCategories(req DescribeEventCategoriesMessage) (resp *DescribeEventCategoriesResult, err error) {
	resp = &DescribeEventCategoriesResult{}
	err = c.client.Do("DescribeEventCategories", "POST", "/", req, resp)
	return
}

// DescribeEventSubscriptions lists all the subscription descriptions for a
// customer account. The description for a subscription includes
// SubscriptionName, SNSTopicARN, CustomerID, SourceType, SourceID,
// CreationTime, and Status. If you specify a SubscriptionName, lists the
// description for that subscription.
func (c *RDS) DescribeEventSubscriptions(req DescribeEventSubscriptionsMessage) (resp *DescribeEventSubscriptionsResult, err error) {
	resp = &DescribeEventSubscriptionsResult{}
	err = c.client.Do("DescribeEventSubscriptions", "POST", "/", req, resp)
	return
}

// DescribeEvents returns events related to DB instances, DB security
// groups, DB snapshots, and DB parameter groups for the past 14 days.
// Events specific to a particular DB instance, DB security group, database
// snapshot, or DB parameter group can be obtained by providing the name as
// a parameter. By default, the past hour of events are returned.
func (c *RDS) DescribeEvents(req DescribeEventsMessage) (resp *DescribeEventsResult, err error) {
	resp = &DescribeEventsResult{}
	err = c.client.Do("DescribeEvents", "POST", "/", req, resp)
	return
}

// DescribeOptionGroupOptions is undocumented.
func (c *RDS) DescribeOptionGroupOptions(req DescribeOptionGroupOptionsMessage) (resp *DescribeOptionGroupOptionsResult, err error) {
	resp = &DescribeOptionGroupOptionsResult{}
	err = c.client.Do("DescribeOptionGroupOptions", "POST", "/", req, resp)
	return
}

// DescribeOptionGroups is undocumented.
func (c *RDS) DescribeOptionGroups(req DescribeOptionGroupsMessage) (resp *DescribeOptionGroupsResult, err error) {
	resp = &DescribeOptionGroupsResult{}
	err = c.client.Do("DescribeOptionGroups", "POST", "/", req, resp)
	return
}

// DescribeOrderableDBInstanceOptions returns a list of orderable DB
// instance options for the specified engine.
func (c *RDS) DescribeOrderableDBInstanceOptions(req DescribeOrderableDBInstanceOptionsMessage) (resp *DescribeOrderableDBInstanceOptionsResult, err error) {
	resp = &DescribeOrderableDBInstanceOptionsResult{}
	err = c.client.Do("DescribeOrderableDBInstanceOptions", "POST", "/", req, resp)
	return
}

// DescribeReservedDBInstances returns information about reserved DB
// instances for this account, or about a specified reserved DB instance.
func (c *RDS) DescribeReservedDBInstances(req DescribeReservedDBInstancesMessage) (resp *DescribeReservedDBInstancesResult, err error) {
	resp = &DescribeReservedDBInstancesResult{}
	err = c.client.Do("DescribeReservedDBInstances", "POST", "/", req, resp)
	return
}

// DescribeReservedDBInstancesOfferings is undocumented.
func (c *RDS) DescribeReservedDBInstancesOfferings(req DescribeReservedDBInstancesOfferingsMessage) (resp *DescribeReservedDBInstancesOfferingsResult, err error) {
	resp = &DescribeReservedDBInstancesOfferingsResult{}
	err = c.client.Do("DescribeReservedDBInstancesOfferings", "POST", "/", req, resp)
	return
}

// DownloadDBLogFilePortion downloads all or a portion of the specified log
// file.
func (c *RDS) DownloadDBLogFilePortion(req DownloadDBLogFilePortionMessage) (resp *DownloadDBLogFilePortionResult, err error) {
	resp = &DownloadDBLogFilePortionResult{}
	err = c.client.Do("DownloadDBLogFilePortion", "POST", "/", req, resp)
	return
}

// ListTagsForResource lists all tags on an Amazon RDS resource. For an
// overview on tagging an Amazon RDS resource, see Tagging Amazon RDS
// Resources
func (c *RDS) ListTagsForResource(req ListTagsForResourceMessage) (resp *ListTagsForResourceResult, err error) {
	resp = &ListTagsForResourceResult{}
	err = c.client.Do("ListTagsForResource", "POST", "/", req, resp)
	return
}

// ModifyDBInstance modify settings for a DB instance. You can change one
// or more database configuration parameters by specifying these parameters
// and the new values in the request.
func (c *RDS) ModifyDBInstance(req ModifyDBInstanceMessage) (resp *ModifyDBInstanceResult, err error) {
	resp = &ModifyDBInstanceResult{}
	err = c.client.Do("ModifyDBInstance", "POST", "/", req, resp)
	return
}

// ModifyDBParameterGroup modifies the parameters of a DB parameter group.
// To modify more than one parameter, submit a list of the following:
// ParameterName , ParameterValue , and ApplyMethod . A maximum of 20
// parameters can be modified in a single request. After you modify a DB
// parameter group, you should wait at least 5 minutes before creating your
// first DB instance that uses that DB parameter group as the default
// parameter group. This allows Amazon RDS to fully complete the modify
// action before the parameter group is used as the default for a new DB
// instance. This is especially important for parameters that are critical
// when creating the default database for a DB instance, such as the
// character set for the default database defined by the
// character_set_database parameter. You can use the Parameter Groups
// option of the Amazon RDS console or the DescribeDBParameters command to
// verify that your DB parameter group has been created or modified.
func (c *RDS) ModifyDBParameterGroup(req ModifyDBParameterGroupMessage) (resp *ModifyDBParameterGroupResult, err error) {
	resp = &ModifyDBParameterGroupResult{}
	err = c.client.Do("ModifyDBParameterGroup", "POST", "/", req, resp)
	return
}

// ModifyDBSubnetGroup modifies an existing DB subnet group. DB subnet
// groups must contain at least one subnet in at least two AZs in the
// region.
func (c *RDS) ModifyDBSubnetGroup(req ModifyDBSubnetGroupMessage) (resp *ModifyDBSubnetGroupResult, err error) {
	resp = &ModifyDBSubnetGroupResult{}
	err = c.client.Do("ModifyDBSubnetGroup", "POST", "/", req, resp)
	return
}

// ModifyEventSubscription modifies an existing RDS event notification
// subscription. Note that you cannot modify the source identifiers using
// this call; to change source identifiers for a subscription, use the
// AddSourceIdentifierToSubscription and
// RemoveSourceIdentifierFromSubscription calls. You can see a list of the
// event categories for a given SourceType in the Events topic in the
// Amazon RDS User Guide or by using the DescribeEventCategories action.
func (c *RDS) ModifyEventSubscription(req ModifyEventSubscriptionMessage) (resp *ModifyEventSubscriptionResult, err error) {
	resp = &ModifyEventSubscriptionResult{}
	err = c.client.Do("ModifyEventSubscription", "POST", "/", req, resp)
	return
}

// ModifyOptionGroup is undocumented.
func (c *RDS) ModifyOptionGroup(req ModifyOptionGroupMessage) (resp *ModifyOptionGroupResult, err error) {
	resp = &ModifyOptionGroupResult{}
	err = c.client.Do("ModifyOptionGroup", "POST", "/", req, resp)
	return
}

// PromoteReadReplica promotes a read replica DB instance to a standalone
// DB instance.
func (c *RDS) PromoteReadReplica(req PromoteReadReplicaMessage) (resp *PromoteReadReplicaResult, err error) {
	resp = &PromoteReadReplicaResult{}
	err = c.client.Do("PromoteReadReplica", "POST", "/", req, resp)
	return
}

// PurchaseReservedDBInstancesOffering is undocumented.
func (c *RDS) PurchaseReservedDBInstancesOffering(req PurchaseReservedDBInstancesOfferingMessage) (resp *PurchaseReservedDBInstancesOfferingResult, err error) {
	resp = &PurchaseReservedDBInstancesOfferingResult{}
	err = c.client.Do("PurchaseReservedDBInstancesOffering", "POST", "/", req, resp)
	return
}

// RebootDBInstance rebooting a DB instance restarts the database engine
// service. A reboot also applies to the DB instance any modifications to
// the associated DB parameter group that were pending. Rebooting a DB
// instance results in a momentary outage of the instance, during which the
// DB instance status is set to rebooting. If the RDS instance is
// configured for MultiAZ, it is possible that the reboot will be conducted
// through a failover. An Amazon RDS event is created when the reboot is
// completed. If your DB instance is deployed in multiple Availability
// Zones, you can force a failover from one AZ to the other during the
// reboot. You might force a failover to test the availability of your DB
// instance deployment or to restore operations to the original AZ after a
// failover occurs. The time required to reboot is a function of the
// specific database engine's crash recovery process. To improve the reboot
// time, we recommend that you reduce database activities as much as
// possible during the reboot process to reduce rollback activity for
// in-transit transactions.
func (c *RDS) RebootDBInstance(req RebootDBInstanceMessage) (resp *RebootDBInstanceResult, err error) {
	resp = &RebootDBInstanceResult{}
	err = c.client.Do("RebootDBInstance", "POST", "/", req, resp)
	return
}

// RemoveSourceIdentifierFromSubscription removes a source identifier from
// an existing RDS event notification subscription.
func (c *RDS) RemoveSourceIdentifierFromSubscription(req RemoveSourceIdentifierFromSubscriptionMessage) (resp *RemoveSourceIdentifierFromSubscriptionResult, err error) {
	resp = &RemoveSourceIdentifierFromSubscriptionResult{}
	err = c.client.Do("RemoveSourceIdentifierFromSubscription", "POST", "/", req, resp)
	return
}

// RemoveTagsFromResource removes metadata tags from an Amazon RDS
// resource. For an overview on tagging an Amazon RDS resource, see Tagging
// Amazon RDS Resources
func (c *RDS) RemoveTagsFromResource(req RemoveTagsFromResourceMessage) (err error) {
	// NRE
	err = c.client.Do("RemoveTagsFromResource", "POST", "/", req, nil)
	return
}

// ResetDBParameterGroup modifies the parameters of a DB parameter group to
// the engine/system default value. To reset specific parameters submit a
// list of the following: ParameterName and ApplyMethod . To reset the
// entire DB parameter group, specify the DBParameterGroup name and
// ResetAllParameters parameters. When resetting the entire group, dynamic
// parameters are updated immediately and static parameters are set to
// pending-reboot to take effect on the next DB instance restart or
// RebootDBInstance request.
func (c *RDS) ResetDBParameterGroup(req ResetDBParameterGroupMessage) (resp *ResetDBParameterGroupResult, err error) {
	resp = &ResetDBParameterGroupResult{}
	err = c.client.Do("ResetDBParameterGroup", "POST", "/", req, resp)
	return
}

// RestoreDBInstanceFromDBSnapshot creates a new DB instance from a DB
// snapshot. The target database is created from the source database
// restore point with the same configuration as the original source
// database, except that the new RDS instance is created with the default
// security group.
func (c *RDS) RestoreDBInstanceFromDBSnapshot(req RestoreDBInstanceFromDBSnapshotMessage) (resp *RestoreDBInstanceFromDBSnapshotResult, err error) {
	resp = &RestoreDBInstanceFromDBSnapshotResult{}
	err = c.client.Do("RestoreDBInstanceFromDBSnapshot", "POST", "/", req, resp)
	return
}

// RestoreDBInstanceToPointInTime restores a DB instance to an arbitrary
// point-in-time. Users can restore to any point in time before the
// latestRestorableTime for up to backupRetentionPeriod days. The target
// database is created from the source database with the same configuration
// as the original database except that the DB instance is created with the
// default DB security group.
func (c *RDS) RestoreDBInstanceToPointInTime(req RestoreDBInstanceToPointInTimeMessage) (resp *RestoreDBInstanceToPointInTimeResult, err error) {
	resp = &RestoreDBInstanceToPointInTimeResult{}
	err = c.client.Do("RestoreDBInstanceToPointInTime", "POST", "/", req, resp)
	return
}

// RevokeDBSecurityGroupIngress revokes ingress from a DBSecurityGroup for
// previously authorized IP ranges or EC2 or VPC Security Groups. Required
// parameters for this API are one of EC2SecurityGroupId for or
// (EC2SecurityGroupOwnerId and either EC2SecurityGroupName or
// EC2SecurityGroupId).
func (c *RDS) RevokeDBSecurityGroupIngress(req RevokeDBSecurityGroupIngressMessage) (resp *RevokeDBSecurityGroupIngressResult, err error) {
	resp = &RevokeDBSecurityGroupIngressResult{}
	err = c.client.Do("RevokeDBSecurityGroupIngress", "POST", "/", req, resp)
	return
}

// AddSourceIdentifierToSubscriptionMessage is undocumented.
type AddSourceIdentifierToSubscriptionMessage struct {
	SourceIdentifier string `xml:"SourceIdentifier"`
	SubscriptionName string `xml:"SubscriptionName"`
}

// AddSourceIdentifierToSubscriptionResult is undocumented.
type AddSourceIdentifierToSubscriptionResult struct {
	EventSubscription EventSubscription `xml:"AddSourceIdentifierToSubscriptionResult>EventSubscription"`
}

// AddTagsToResourceMessage is undocumented.
type AddTagsToResourceMessage struct {
	ResourceName string `xml:"ResourceName"`
	Tags         []Tag  `xml:"Tags>Tag"`
}

// AuthorizeDBSecurityGroupIngressMessage is undocumented.
type AuthorizeDBSecurityGroupIngressMessage struct {
	CIDRIP                  string `xml:"CIDRIP"`
	DBSecurityGroupName     string `xml:"DBSecurityGroupName"`
	EC2SecurityGroupID      string `xml:"EC2SecurityGroupId"`
	EC2SecurityGroupName    string `xml:"EC2SecurityGroupName"`
	EC2SecurityGroupOwnerID string `xml:"EC2SecurityGroupOwnerId"`
}

// AuthorizeDBSecurityGroupIngressResult is undocumented.
type AuthorizeDBSecurityGroupIngressResult struct {
	DBSecurityGroup DBSecurityGroup `xml:"AuthorizeDBSecurityGroupIngressResult>DBSecurityGroup"`
}

// AvailabilityZone is undocumented.
type AvailabilityZone struct {
	Name string `xml:"Name"`
}

// CharacterSet is undocumented.
type CharacterSet struct {
	CharacterSetDescription string `xml:"CharacterSetDescription"`
	CharacterSetName        string `xml:"CharacterSetName"`
}

// CopyDBParameterGroupMessage is undocumented.
type CopyDBParameterGroupMessage struct {
	SourceDBParameterGroupIdentifier  string `xml:"SourceDBParameterGroupIdentifier"`
	Tags                              []Tag  `xml:"Tags>Tag"`
	TargetDBParameterGroupDescription string `xml:"TargetDBParameterGroupDescription"`
	TargetDBParameterGroupIdentifier  string `xml:"TargetDBParameterGroupIdentifier"`
}

// CopyDBParameterGroupResult is undocumented.
type CopyDBParameterGroupResult struct {
	DBParameterGroup DBParameterGroup `xml:"CopyDBParameterGroupResult>DBParameterGroup"`
}

// CopyDBSnapshotMessage is undocumented.
type CopyDBSnapshotMessage struct {
	SourceDBSnapshotIdentifier string `xml:"SourceDBSnapshotIdentifier"`
	Tags                       []Tag  `xml:"Tags>Tag"`
	TargetDBSnapshotIdentifier string `xml:"TargetDBSnapshotIdentifier"`
}

// CopyDBSnapshotResult is undocumented.
type CopyDBSnapshotResult struct {
	DBSnapshot DBSnapshot `xml:"CopyDBSnapshotResult>DBSnapshot"`
}

// CopyOptionGroupMessage is undocumented.
type CopyOptionGroupMessage struct {
	SourceOptionGroupIdentifier  string `xml:"SourceOptionGroupIdentifier"`
	Tags                         []Tag  `xml:"Tags>Tag"`
	TargetOptionGroupDescription string `xml:"TargetOptionGroupDescription"`
	TargetOptionGroupIdentifier  string `xml:"TargetOptionGroupIdentifier"`
}

// CopyOptionGroupResult is undocumented.
type CopyOptionGroupResult struct {
	OptionGroup OptionGroup `xml:"CopyOptionGroupResult>OptionGroup"`
}

// CreateDBInstanceMessage is undocumented.
type CreateDBInstanceMessage struct {
	AllocatedStorage           int      `xml:"AllocatedStorage"`
	AutoMinorVersionUpgrade    bool     `xml:"AutoMinorVersionUpgrade"`
	AvailabilityZone           string   `xml:"AvailabilityZone"`
	BackupRetentionPeriod      int      `xml:"BackupRetentionPeriod"`
	CharacterSetName           string   `xml:"CharacterSetName"`
	DBInstanceClass            string   `xml:"DBInstanceClass"`
	DBInstanceIdentifier       string   `xml:"DBInstanceIdentifier"`
	DBName                     string   `xml:"DBName"`
	DBParameterGroupName       string   `xml:"DBParameterGroupName"`
	DBSecurityGroups           []string `xml:"DBSecurityGroups>DBSecurityGroupName"`
	DBSubnetGroupName          string   `xml:"DBSubnetGroupName"`
	Engine                     string   `xml:"Engine"`
	EngineVersion              string   `xml:"EngineVersion"`
	Iops                       int      `xml:"Iops"`
	LicenseModel               string   `xml:"LicenseModel"`
	MasterUserPassword         string   `xml:"MasterUserPassword"`
	MasterUsername             string   `xml:"MasterUsername"`
	MultiAZ                    bool     `xml:"MultiAZ"`
	OptionGroupName            string   `xml:"OptionGroupName"`
	Port                       int      `xml:"Port"`
	PreferredBackupWindow      string   `xml:"PreferredBackupWindow"`
	PreferredMaintenanceWindow string   `xml:"PreferredMaintenanceWindow"`
	PubliclyAccessible         bool     `xml:"PubliclyAccessible"`
	StorageType                string   `xml:"StorageType"`
	Tags                       []Tag    `xml:"Tags>Tag"`
	TdeCredentialARN           string   `xml:"TdeCredentialArn"`
	TdeCredentialPassword      string   `xml:"TdeCredentialPassword"`
	VpcSecurityGroupIds        []string `xml:"VpcSecurityGroupIds>VpcSecurityGroupId"`
}

// CreateDBInstanceReadReplicaMessage is undocumented.
type CreateDBInstanceReadReplicaMessage struct {
	AutoMinorVersionUpgrade    bool   `xml:"AutoMinorVersionUpgrade"`
	AvailabilityZone           string `xml:"AvailabilityZone"`
	DBInstanceClass            string `xml:"DBInstanceClass"`
	DBInstanceIdentifier       string `xml:"DBInstanceIdentifier"`
	DBSubnetGroupName          string `xml:"DBSubnetGroupName"`
	Iops                       int    `xml:"Iops"`
	OptionGroupName            string `xml:"OptionGroupName"`
	Port                       int    `xml:"Port"`
	PubliclyAccessible         bool   `xml:"PubliclyAccessible"`
	SourceDBInstanceIdentifier string `xml:"SourceDBInstanceIdentifier"`
	StorageType                string `xml:"StorageType"`
	Tags                       []Tag  `xml:"Tags>Tag"`
}

// CreateDBInstanceReadReplicaResult is undocumented.
type CreateDBInstanceReadReplicaResult struct {
	DBInstance DBInstance `xml:"CreateDBInstanceReadReplicaResult>DBInstance"`
}

// CreateDBInstanceResult is undocumented.
type CreateDBInstanceResult struct {
	DBInstance DBInstance `xml:"CreateDBInstanceResult>DBInstance"`
}

// CreateDBParameterGroupMessage is undocumented.
type CreateDBParameterGroupMessage struct {
	DBParameterGroupFamily string `xml:"DBParameterGroupFamily"`
	DBParameterGroupName   string `xml:"DBParameterGroupName"`
	Description            string `xml:"Description"`
	Tags                   []Tag  `xml:"Tags>Tag"`
}

// CreateDBParameterGroupResult is undocumented.
type CreateDBParameterGroupResult struct {
	DBParameterGroup DBParameterGroup `xml:"CreateDBParameterGroupResult>DBParameterGroup"`
}

// CreateDBSecurityGroupMessage is undocumented.
type CreateDBSecurityGroupMessage struct {
	DBSecurityGroupDescription string `xml:"DBSecurityGroupDescription"`
	DBSecurityGroupName        string `xml:"DBSecurityGroupName"`
	Tags                       []Tag  `xml:"Tags>Tag"`
}

// CreateDBSecurityGroupResult is undocumented.
type CreateDBSecurityGroupResult struct {
	DBSecurityGroup DBSecurityGroup `xml:"CreateDBSecurityGroupResult>DBSecurityGroup"`
}

// CreateDBSnapshotMessage is undocumented.
type CreateDBSnapshotMessage struct {
	DBInstanceIdentifier string `xml:"DBInstanceIdentifier"`
	DBSnapshotIdentifier string `xml:"DBSnapshotIdentifier"`
	Tags                 []Tag  `xml:"Tags>Tag"`
}

// CreateDBSnapshotResult is undocumented.
type CreateDBSnapshotResult struct {
	DBSnapshot DBSnapshot `xml:"CreateDBSnapshotResult>DBSnapshot"`
}

// CreateDBSubnetGroupMessage is undocumented.
type CreateDBSubnetGroupMessage struct {
	DBSubnetGroupDescription string   `xml:"DBSubnetGroupDescription"`
	DBSubnetGroupName        string   `xml:"DBSubnetGroupName"`
	SubnetIds                []string `xml:"SubnetIds>SubnetIdentifier"`
	Tags                     []Tag    `xml:"Tags>Tag"`
}

// CreateDBSubnetGroupResult is undocumented.
type CreateDBSubnetGroupResult struct {
	DBSubnetGroup DBSubnetGroup `xml:"CreateDBSubnetGroupResult>DBSubnetGroup"`
}

// CreateEventSubscriptionMessage is undocumented.
type CreateEventSubscriptionMessage struct {
	Enabled          bool     `xml:"Enabled"`
	EventCategories  []string `xml:"EventCategories>EventCategory"`
	SnsTopicARN      string   `xml:"SnsTopicArn"`
	SourceIds        []string `xml:"SourceIds>SourceId"`
	SourceType       string   `xml:"SourceType"`
	SubscriptionName string   `xml:"SubscriptionName"`
	Tags             []Tag    `xml:"Tags>Tag"`
}

// CreateEventSubscriptionResult is undocumented.
type CreateEventSubscriptionResult struct {
	EventSubscription EventSubscription `xml:"CreateEventSubscriptionResult>EventSubscription"`
}

// CreateOptionGroupMessage is undocumented.
type CreateOptionGroupMessage struct {
	EngineName             string `xml:"EngineName"`
	MajorEngineVersion     string `xml:"MajorEngineVersion"`
	OptionGroupDescription string `xml:"OptionGroupDescription"`
	OptionGroupName        string `xml:"OptionGroupName"`
	Tags                   []Tag  `xml:"Tags>Tag"`
}

// CreateOptionGroupResult is undocumented.
type CreateOptionGroupResult struct {
	OptionGroup OptionGroup `xml:"CreateOptionGroupResult>OptionGroup"`
}

// DBEngineVersion is undocumented.
type DBEngineVersion struct {
	DBEngineDescription        string         `xml:"DBEngineDescription"`
	DBEngineVersionDescription string         `xml:"DBEngineVersionDescription"`
	DBParameterGroupFamily     string         `xml:"DBParameterGroupFamily"`
	DefaultCharacterSet        CharacterSet   `xml:"DefaultCharacterSet"`
	Engine                     string         `xml:"Engine"`
	EngineVersion              string         `xml:"EngineVersion"`
	SupportedCharacterSets     []CharacterSet `xml:"SupportedCharacterSets>CharacterSet"`
}

// DBEngineVersionMessage is undocumented.
type DBEngineVersionMessage struct {
	DBEngineVersions []DBEngineVersion `xml:"DescribeDBEngineVersionsResult>DBEngineVersions>DBEngineVersion"`
	Marker           string            `xml:"DescribeDBEngineVersionsResult>Marker"`
}

// DBInstance is undocumented.
type DBInstance struct {
	AllocatedStorage                      int                          `xml:"AllocatedStorage"`
	AutoMinorVersionUpgrade               bool                         `xml:"AutoMinorVersionUpgrade"`
	AvailabilityZone                      string                       `xml:"AvailabilityZone"`
	BackupRetentionPeriod                 int                          `xml:"BackupRetentionPeriod"`
	CharacterSetName                      string                       `xml:"CharacterSetName"`
	DBInstanceClass                       string                       `xml:"DBInstanceClass"`
	DBInstanceIdentifier                  string                       `xml:"DBInstanceIdentifier"`
	DBInstanceStatus                      string                       `xml:"DBInstanceStatus"`
	DBName                                string                       `xml:"DBName"`
	DBParameterGroups                     []DBParameterGroupStatus     `xml:"DBParameterGroups>DBParameterGroup"`
	DBSecurityGroups                      []DBSecurityGroupMembership  `xml:"DBSecurityGroups>DBSecurityGroup"`
	DBSubnetGroup                         DBSubnetGroup                `xml:"DBSubnetGroup"`
	Endpoint                              Endpoint                     `xml:"Endpoint"`
	Engine                                string                       `xml:"Engine"`
	EngineVersion                         string                       `xml:"EngineVersion"`
	InstanceCreateTime                    time.Time                    `xml:"InstanceCreateTime"`
	Iops                                  int                          `xml:"Iops"`
	LatestRestorableTime                  time.Time                    `xml:"LatestRestorableTime"`
	LicenseModel                          string                       `xml:"LicenseModel"`
	MasterUsername                        string                       `xml:"MasterUsername"`
	MultiAZ                               bool                         `xml:"MultiAZ"`
	OptionGroupMemberships                []OptionGroupMembership      `xml:"OptionGroupMemberships>OptionGroupMembership"`
	PendingModifiedValues                 PendingModifiedValues        `xml:"PendingModifiedValues"`
	PreferredBackupWindow                 string                       `xml:"PreferredBackupWindow"`
	PreferredMaintenanceWindow            string                       `xml:"PreferredMaintenanceWindow"`
	PubliclyAccessible                    bool                         `xml:"PubliclyAccessible"`
	ReadReplicaDBInstanceIdentifiers      []string                     `xml:"ReadReplicaDBInstanceIdentifiers>ReadReplicaDBInstanceIdentifier"`
	ReadReplicaSourceDBInstanceIdentifier string                       `xml:"ReadReplicaSourceDBInstanceIdentifier"`
	SecondaryAvailabilityZone             string                       `xml:"SecondaryAvailabilityZone"`
	StatusInfos                           []DBInstanceStatusInfo       `xml:"StatusInfos>DBInstanceStatusInfo"`
	StorageType                           string                       `xml:"StorageType"`
	TdeCredentialARN                      string                       `xml:"TdeCredentialArn"`
	VpcSecurityGroups                     []VpcSecurityGroupMembership `xml:"VpcSecurityGroups>VpcSecurityGroupMembership"`
}

// DBInstanceMessage is undocumented.
type DBInstanceMessage struct {
	DBInstances []DBInstance `xml:"DescribeDBInstancesResult>DBInstances>DBInstance"`
	Marker      string       `xml:"DescribeDBInstancesResult>Marker"`
}

// DBInstanceStatusInfo is undocumented.
type DBInstanceStatusInfo struct {
	Message    string `xml:"Message"`
	Normal     bool   `xml:"Normal"`
	Status     string `xml:"Status"`
	StatusType string `xml:"StatusType"`
}

// DBParameterGroup is undocumented.
type DBParameterGroup struct {
	DBParameterGroupFamily string `xml:"DBParameterGroupFamily"`
	DBParameterGroupName   string `xml:"DBParameterGroupName"`
	Description            string `xml:"Description"`
}

// DBParameterGroupDetails is undocumented.
type DBParameterGroupDetails struct {
	Marker     string      `xml:"DescribeDBParametersResult>Marker"`
	Parameters []Parameter `xml:"DescribeDBParametersResult>Parameters>Parameter"`
}

// DBParameterGroupNameMessage is undocumented.
type DBParameterGroupNameMessage struct {
	DBParameterGroupName string `xml:"DBParameterGroupName"`
}

// DBParameterGroupStatus is undocumented.
type DBParameterGroupStatus struct {
	DBParameterGroupName string `xml:"DBParameterGroupName"`
	ParameterApplyStatus string `xml:"ParameterApplyStatus"`
}

// DBParameterGroupsMessage is undocumented.
type DBParameterGroupsMessage struct {
	DBParameterGroups []DBParameterGroup `xml:"DescribeDBParameterGroupsResult>DBParameterGroups>DBParameterGroup"`
	Marker            string             `xml:"DescribeDBParameterGroupsResult>Marker"`
}

// DBSecurityGroup is undocumented.
type DBSecurityGroup struct {
	DBSecurityGroupDescription string             `xml:"DBSecurityGroupDescription"`
	DBSecurityGroupName        string             `xml:"DBSecurityGroupName"`
	EC2SecurityGroups          []EC2SecurityGroup `xml:"EC2SecurityGroups>EC2SecurityGroup"`
	IPRanges                   []IPRange          `xml:"IPRanges>IPRange"`
	OwnerID                    string             `xml:"OwnerId"`
	VpcID                      string             `xml:"VpcId"`
}

// DBSecurityGroupMembership is undocumented.
type DBSecurityGroupMembership struct {
	DBSecurityGroupName string `xml:"DBSecurityGroupName"`
	Status              string `xml:"Status"`
}

// DBSecurityGroupMessage is undocumented.
type DBSecurityGroupMessage struct {
	DBSecurityGroups []DBSecurityGroup `xml:"DescribeDBSecurityGroupsResult>DBSecurityGroups>DBSecurityGroup"`
	Marker           string            `xml:"DescribeDBSecurityGroupsResult>Marker"`
}

// DBSnapshot is undocumented.
type DBSnapshot struct {
	AllocatedStorage     int       `xml:"AllocatedStorage"`
	AvailabilityZone     string    `xml:"AvailabilityZone"`
	DBInstanceIdentifier string    `xml:"DBInstanceIdentifier"`
	DBSnapshotIdentifier string    `xml:"DBSnapshotIdentifier"`
	Engine               string    `xml:"Engine"`
	EngineVersion        string    `xml:"EngineVersion"`
	InstanceCreateTime   time.Time `xml:"InstanceCreateTime"`
	Iops                 int       `xml:"Iops"`
	LicenseModel         string    `xml:"LicenseModel"`
	MasterUsername       string    `xml:"MasterUsername"`
	OptionGroupName      string    `xml:"OptionGroupName"`
	PercentProgress      int       `xml:"PercentProgress"`
	Port                 int       `xml:"Port"`
	SnapshotCreateTime   time.Time `xml:"SnapshotCreateTime"`
	SnapshotType         string    `xml:"SnapshotType"`
	SourceRegion         string    `xml:"SourceRegion"`
	Status               string    `xml:"Status"`
	StorageType          string    `xml:"StorageType"`
	TdeCredentialARN     string    `xml:"TdeCredentialArn"`
	VpcID                string    `xml:"VpcId"`
}

// DBSnapshotMessage is undocumented.
type DBSnapshotMessage struct {
	DBSnapshots []DBSnapshot `xml:"DescribeDBSnapshotsResult>DBSnapshots>DBSnapshot"`
	Marker      string       `xml:"DescribeDBSnapshotsResult>Marker"`
}

// DBSubnetGroup is undocumented.
type DBSubnetGroup struct {
	DBSubnetGroupDescription string   `xml:"DBSubnetGroupDescription"`
	DBSubnetGroupName        string   `xml:"DBSubnetGroupName"`
	SubnetGroupStatus        string   `xml:"SubnetGroupStatus"`
	Subnets                  []Subnet `xml:"Subnets>Subnet"`
	VpcID                    string   `xml:"VpcId"`
}

// DBSubnetGroupMessage is undocumented.
type DBSubnetGroupMessage struct {
	DBSubnetGroups []DBSubnetGroup `xml:"DescribeDBSubnetGroupsResult>DBSubnetGroups>DBSubnetGroup"`
	Marker         string          `xml:"DescribeDBSubnetGroupsResult>Marker"`
}

// DeleteDBInstanceMessage is undocumented.
type DeleteDBInstanceMessage struct {
	DBInstanceIdentifier      string `xml:"DBInstanceIdentifier"`
	FinalDBSnapshotIdentifier string `xml:"FinalDBSnapshotIdentifier"`
	SkipFinalSnapshot         bool   `xml:"SkipFinalSnapshot"`
}

// DeleteDBInstanceResult is undocumented.
type DeleteDBInstanceResult struct {
	DBInstance DBInstance `xml:"DeleteDBInstanceResult>DBInstance"`
}

// DeleteDBParameterGroupMessage is undocumented.
type DeleteDBParameterGroupMessage struct {
	DBParameterGroupName string `xml:"DBParameterGroupName"`
}

// DeleteDBSecurityGroupMessage is undocumented.
type DeleteDBSecurityGroupMessage struct {
	DBSecurityGroupName string `xml:"DBSecurityGroupName"`
}

// DeleteDBSnapshotMessage is undocumented.
type DeleteDBSnapshotMessage struct {
	DBSnapshotIdentifier string `xml:"DBSnapshotIdentifier"`
}

// DeleteDBSnapshotResult is undocumented.
type DeleteDBSnapshotResult struct {
	DBSnapshot DBSnapshot `xml:"DeleteDBSnapshotResult>DBSnapshot"`
}

// DeleteDBSubnetGroupMessage is undocumented.
type DeleteDBSubnetGroupMessage struct {
	DBSubnetGroupName string `xml:"DBSubnetGroupName"`
}

// DeleteEventSubscriptionMessage is undocumented.
type DeleteEventSubscriptionMessage struct {
	SubscriptionName string `xml:"SubscriptionName"`
}

// DeleteEventSubscriptionResult is undocumented.
type DeleteEventSubscriptionResult struct {
	EventSubscription EventSubscription `xml:"DeleteEventSubscriptionResult>EventSubscription"`
}

// DeleteOptionGroupMessage is undocumented.
type DeleteOptionGroupMessage struct {
	OptionGroupName string `xml:"OptionGroupName"`
}

// DescribeDBEngineVersionsMessage is undocumented.
type DescribeDBEngineVersionsMessage struct {
	DBParameterGroupFamily     string   `xml:"DBParameterGroupFamily"`
	DefaultOnly                bool     `xml:"DefaultOnly"`
	Engine                     string   `xml:"Engine"`
	EngineVersion              string   `xml:"EngineVersion"`
	Filters                    []Filter `xml:"Filters>Filter"`
	ListSupportedCharacterSets bool     `xml:"ListSupportedCharacterSets"`
	Marker                     string   `xml:"Marker"`
	MaxRecords                 int      `xml:"MaxRecords"`
}

// DescribeDBInstancesMessage is undocumented.
type DescribeDBInstancesMessage struct {
	DBInstanceIdentifier string   `xml:"DBInstanceIdentifier"`
	Filters              []Filter `xml:"Filters>Filter"`
	Marker               string   `xml:"Marker"`
	MaxRecords           int      `xml:"MaxRecords"`
}

// DescribeDBLogFilesDetails is undocumented.
type DescribeDBLogFilesDetails struct {
	LastWritten int64  `xml:"LastWritten"`
	LogFileName string `xml:"LogFileName"`
	Size        int64  `xml:"Size"`
}

// DescribeDBLogFilesMessage is undocumented.
type DescribeDBLogFilesMessage struct {
	DBInstanceIdentifier string   `xml:"DBInstanceIdentifier"`
	FileLastWritten      int64    `xml:"FileLastWritten"`
	FileSize             int64    `xml:"FileSize"`
	FilenameContains     string   `xml:"FilenameContains"`
	Filters              []Filter `xml:"Filters>Filter"`
	Marker               string   `xml:"Marker"`
	MaxRecords           int      `xml:"MaxRecords"`
}

// DescribeDBLogFilesResponse is undocumented.
type DescribeDBLogFilesResponse struct {
	DescribeDBLogFiles []DescribeDBLogFilesDetails `xml:"DescribeDBLogFilesResult>DescribeDBLogFiles>DescribeDBLogFilesDetails"`
	Marker             string                      `xml:"DescribeDBLogFilesResult>Marker"`
}

// DescribeDBParameterGroupsMessage is undocumented.
type DescribeDBParameterGroupsMessage struct {
	DBParameterGroupName string   `xml:"DBParameterGroupName"`
	Filters              []Filter `xml:"Filters>Filter"`
	Marker               string   `xml:"Marker"`
	MaxRecords           int      `xml:"MaxRecords"`
}

// DescribeDBParametersMessage is undocumented.
type DescribeDBParametersMessage struct {
	DBParameterGroupName string   `xml:"DBParameterGroupName"`
	Filters              []Filter `xml:"Filters>Filter"`
	Marker               string   `xml:"Marker"`
	MaxRecords           int      `xml:"MaxRecords"`
	Source               string   `xml:"Source"`
}

// DescribeDBSecurityGroupsMessage is undocumented.
type DescribeDBSecurityGroupsMessage struct {
	DBSecurityGroupName string   `xml:"DBSecurityGroupName"`
	Filters             []Filter `xml:"Filters>Filter"`
	Marker              string   `xml:"Marker"`
	MaxRecords          int      `xml:"MaxRecords"`
}

// DescribeDBSnapshotsMessage is undocumented.
type DescribeDBSnapshotsMessage struct {
	DBInstanceIdentifier string   `xml:"DBInstanceIdentifier"`
	DBSnapshotIdentifier string   `xml:"DBSnapshotIdentifier"`
	Filters              []Filter `xml:"Filters>Filter"`
	Marker               string   `xml:"Marker"`
	MaxRecords           int      `xml:"MaxRecords"`
	SnapshotType         string   `xml:"SnapshotType"`
}

// DescribeDBSubnetGroupsMessage is undocumented.
type DescribeDBSubnetGroupsMessage struct {
	DBSubnetGroupName string   `xml:"DBSubnetGroupName"`
	Filters           []Filter `xml:"Filters>Filter"`
	Marker            string   `xml:"Marker"`
	MaxRecords        int      `xml:"MaxRecords"`
}

// DescribeEngineDefaultParametersMessage is undocumented.
type DescribeEngineDefaultParametersMessage struct {
	DBParameterGroupFamily string   `xml:"DBParameterGroupFamily"`
	Filters                []Filter `xml:"Filters>Filter"`
	Marker                 string   `xml:"Marker"`
	MaxRecords             int      `xml:"MaxRecords"`
}

// DescribeEngineDefaultParametersResult is undocumented.
type DescribeEngineDefaultParametersResult struct {
	EngineDefaults EngineDefaults `xml:"DescribeEngineDefaultParametersResult>EngineDefaults"`
}

// DescribeEventCategoriesMessage is undocumented.
type DescribeEventCategoriesMessage struct {
	Filters    []Filter `xml:"Filters>Filter"`
	SourceType string   `xml:"SourceType"`
}

// DescribeEventSubscriptionsMessage is undocumented.
type DescribeEventSubscriptionsMessage struct {
	Filters          []Filter `xml:"Filters>Filter"`
	Marker           string   `xml:"Marker"`
	MaxRecords       int      `xml:"MaxRecords"`
	SubscriptionName string   `xml:"SubscriptionName"`
}

// DescribeEventsMessage is undocumented.
type DescribeEventsMessage struct {
	Duration         int       `xml:"Duration"`
	EndTime          time.Time `xml:"EndTime"`
	EventCategories  []string  `xml:"EventCategories>EventCategory"`
	Filters          []Filter  `xml:"Filters>Filter"`
	Marker           string    `xml:"Marker"`
	MaxRecords       int       `xml:"MaxRecords"`
	SourceIdentifier string    `xml:"SourceIdentifier"`
	SourceType       string    `xml:"SourceType"`
	StartTime        time.Time `xml:"StartTime"`
}

// DescribeOptionGroupOptionsMessage is undocumented.
type DescribeOptionGroupOptionsMessage struct {
	EngineName         string   `xml:"EngineName"`
	Filters            []Filter `xml:"Filters>Filter"`
	MajorEngineVersion string   `xml:"MajorEngineVersion"`
	Marker             string   `xml:"Marker"`
	MaxRecords         int      `xml:"MaxRecords"`
}

// DescribeOptionGroupsMessage is undocumented.
type DescribeOptionGroupsMessage struct {
	EngineName         string   `xml:"EngineName"`
	Filters            []Filter `xml:"Filters>Filter"`
	MajorEngineVersion string   `xml:"MajorEngineVersion"`
	Marker             string   `xml:"Marker"`
	MaxRecords         int      `xml:"MaxRecords"`
	OptionGroupName    string   `xml:"OptionGroupName"`
}

// DescribeOrderableDBInstanceOptionsMessage is undocumented.
type DescribeOrderableDBInstanceOptionsMessage struct {
	DBInstanceClass string   `xml:"DBInstanceClass"`
	Engine          string   `xml:"Engine"`
	EngineVersion   string   `xml:"EngineVersion"`
	Filters         []Filter `xml:"Filters>Filter"`
	LicenseModel    string   `xml:"LicenseModel"`
	Marker          string   `xml:"Marker"`
	MaxRecords      int      `xml:"MaxRecords"`
	Vpc             bool     `xml:"Vpc"`
}

// DescribeReservedDBInstancesMessage is undocumented.
type DescribeReservedDBInstancesMessage struct {
	DBInstanceClass               string   `xml:"DBInstanceClass"`
	Duration                      string   `xml:"Duration"`
	Filters                       []Filter `xml:"Filters>Filter"`
	Marker                        string   `xml:"Marker"`
	MaxRecords                    int      `xml:"MaxRecords"`
	MultiAZ                       bool     `xml:"MultiAZ"`
	OfferingType                  string   `xml:"OfferingType"`
	ProductDescription            string   `xml:"ProductDescription"`
	ReservedDBInstanceID          string   `xml:"ReservedDBInstanceId"`
	ReservedDBInstancesOfferingID string   `xml:"ReservedDBInstancesOfferingId"`
}

// DescribeReservedDBInstancesOfferingsMessage is undocumented.
type DescribeReservedDBInstancesOfferingsMessage struct {
	DBInstanceClass               string   `xml:"DBInstanceClass"`
	Duration                      string   `xml:"Duration"`
	Filters                       []Filter `xml:"Filters>Filter"`
	Marker                        string   `xml:"Marker"`
	MaxRecords                    int      `xml:"MaxRecords"`
	MultiAZ                       bool     `xml:"MultiAZ"`
	OfferingType                  string   `xml:"OfferingType"`
	ProductDescription            string   `xml:"ProductDescription"`
	ReservedDBInstancesOfferingID string   `xml:"ReservedDBInstancesOfferingId"`
}

// DownloadDBLogFilePortionDetails is undocumented.
type DownloadDBLogFilePortionDetails struct {
	AdditionalDataPending bool   `xml:"DownloadDBLogFilePortionResult>AdditionalDataPending"`
	LogFileData           string `xml:"DownloadDBLogFilePortionResult>LogFileData"`
	Marker                string `xml:"DownloadDBLogFilePortionResult>Marker"`
}

// DownloadDBLogFilePortionMessage is undocumented.
type DownloadDBLogFilePortionMessage struct {
	DBInstanceIdentifier string `xml:"DBInstanceIdentifier"`
	LogFileName          string `xml:"LogFileName"`
	Marker               string `xml:"Marker"`
	NumberOfLines        int    `xml:"NumberOfLines"`
}

// EC2SecurityGroup is undocumented.
type EC2SecurityGroup struct {
	EC2SecurityGroupID      string `xml:"EC2SecurityGroupId"`
	EC2SecurityGroupName    string `xml:"EC2SecurityGroupName"`
	EC2SecurityGroupOwnerID string `xml:"EC2SecurityGroupOwnerId"`
	Status                  string `xml:"Status"`
}

// Endpoint is undocumented.
type Endpoint struct {
	Address string `xml:"Address"`
	Port    int    `xml:"Port"`
}

// EngineDefaults is undocumented.
type EngineDefaults struct {
	DBParameterGroupFamily string      `xml:"DBParameterGroupFamily"`
	Marker                 string      `xml:"Marker"`
	Parameters             []Parameter `xml:"Parameters>Parameter"`
}

// Event is undocumented.
type Event struct {
	Date             time.Time `xml:"Date"`
	EventCategories  []string  `xml:"EventCategories>EventCategory"`
	Message          string    `xml:"Message"`
	SourceIdentifier string    `xml:"SourceIdentifier"`
	SourceType       string    `xml:"SourceType"`
}

// EventCategoriesMap is undocumented.
type EventCategoriesMap struct {
	EventCategories []string `xml:"EventCategories>EventCategory"`
	SourceType      string   `xml:"SourceType"`
}

// EventCategoriesMessage is undocumented.
type EventCategoriesMessage struct {
	EventCategoriesMapList []EventCategoriesMap `xml:"DescribeEventCategoriesResult>EventCategoriesMapList>EventCategoriesMap"`
}

// EventSubscription is undocumented.
type EventSubscription struct {
	CustSubscriptionID       string   `xml:"CustSubscriptionId"`
	CustomerAwsID            string   `xml:"CustomerAwsId"`
	Enabled                  bool     `xml:"Enabled"`
	EventCategoriesList      []string `xml:"EventCategoriesList>EventCategory"`
	SnsTopicARN              string   `xml:"SnsTopicArn"`
	SourceIdsList            []string `xml:"SourceIdsList>SourceId"`
	SourceType               string   `xml:"SourceType"`
	Status                   string   `xml:"Status"`
	SubscriptionCreationTime string   `xml:"SubscriptionCreationTime"`
}

// EventSubscriptionsMessage is undocumented.
type EventSubscriptionsMessage struct {
	EventSubscriptionsList []EventSubscription `xml:"DescribeEventSubscriptionsResult>EventSubscriptionsList>EventSubscription"`
	Marker                 string              `xml:"DescribeEventSubscriptionsResult>Marker"`
}

// EventsMessage is undocumented.
type EventsMessage struct {
	Events []Event `xml:"DescribeEventsResult>Events>Event"`
	Marker string  `xml:"DescribeEventsResult>Marker"`
}

// Filter is undocumented.
type Filter struct {
	Name   string   `xml:"Name"`
	Values []string `xml:"Values>Value"`
}

// IPRange is undocumented.
type IPRange struct {
	CIDRIP string `xml:"CIDRIP"`
	Status string `xml:"Status"`
}

// ListTagsForResourceMessage is undocumented.
type ListTagsForResourceMessage struct {
	Filters      []Filter `xml:"Filters>Filter"`
	ResourceName string   `xml:"ResourceName"`
}

// ModifyDBInstanceMessage is undocumented.
type ModifyDBInstanceMessage struct {
	AllocatedStorage           int      `xml:"AllocatedStorage"`
	AllowMajorVersionUpgrade   bool     `xml:"AllowMajorVersionUpgrade"`
	ApplyImmediately           bool     `xml:"ApplyImmediately"`
	AutoMinorVersionUpgrade    bool     `xml:"AutoMinorVersionUpgrade"`
	BackupRetentionPeriod      int      `xml:"BackupRetentionPeriod"`
	DBInstanceClass            string   `xml:"DBInstanceClass"`
	DBInstanceIdentifier       string   `xml:"DBInstanceIdentifier"`
	DBParameterGroupName       string   `xml:"DBParameterGroupName"`
	DBSecurityGroups           []string `xml:"DBSecurityGroups>DBSecurityGroupName"`
	EngineVersion              string   `xml:"EngineVersion"`
	Iops                       int      `xml:"Iops"`
	MasterUserPassword         string   `xml:"MasterUserPassword"`
	MultiAZ                    bool     `xml:"MultiAZ"`
	NewDBInstanceIdentifier    string   `xml:"NewDBInstanceIdentifier"`
	OptionGroupName            string   `xml:"OptionGroupName"`
	PreferredBackupWindow      string   `xml:"PreferredBackupWindow"`
	PreferredMaintenanceWindow string   `xml:"PreferredMaintenanceWindow"`
	StorageType                string   `xml:"StorageType"`
	TdeCredentialARN           string   `xml:"TdeCredentialArn"`
	TdeCredentialPassword      string   `xml:"TdeCredentialPassword"`
	VpcSecurityGroupIds        []string `xml:"VpcSecurityGroupIds>VpcSecurityGroupId"`
}

// ModifyDBInstanceResult is undocumented.
type ModifyDBInstanceResult struct {
	DBInstance DBInstance `xml:"ModifyDBInstanceResult>DBInstance"`
}

// ModifyDBParameterGroupMessage is undocumented.
type ModifyDBParameterGroupMessage struct {
	DBParameterGroupName string      `xml:"DBParameterGroupName"`
	Parameters           []Parameter `xml:"Parameters>Parameter"`
}

// ModifyDBSubnetGroupMessage is undocumented.
type ModifyDBSubnetGroupMessage struct {
	DBSubnetGroupDescription string   `xml:"DBSubnetGroupDescription"`
	DBSubnetGroupName        string   `xml:"DBSubnetGroupName"`
	SubnetIds                []string `xml:"SubnetIds>SubnetIdentifier"`
}

// ModifyDBSubnetGroupResult is undocumented.
type ModifyDBSubnetGroupResult struct {
	DBSubnetGroup DBSubnetGroup `xml:"ModifyDBSubnetGroupResult>DBSubnetGroup"`
}

// ModifyEventSubscriptionMessage is undocumented.
type ModifyEventSubscriptionMessage struct {
	Enabled          bool     `xml:"Enabled"`
	EventCategories  []string `xml:"EventCategories>EventCategory"`
	SnsTopicARN      string   `xml:"SnsTopicArn"`
	SourceType       string   `xml:"SourceType"`
	SubscriptionName string   `xml:"SubscriptionName"`
}

// ModifyEventSubscriptionResult is undocumented.
type ModifyEventSubscriptionResult struct {
	EventSubscription EventSubscription `xml:"ModifyEventSubscriptionResult>EventSubscription"`
}

// ModifyOptionGroupMessage is undocumented.
type ModifyOptionGroupMessage struct {
	ApplyImmediately bool                  `xml:"ApplyImmediately"`
	OptionGroupName  string                `xml:"OptionGroupName"`
	OptionsToInclude []OptionConfiguration `xml:"OptionsToInclude>OptionConfiguration"`
	OptionsToRemove  []string              `xml:"OptionsToRemove>member"`
}

// ModifyOptionGroupResult is undocumented.
type ModifyOptionGroupResult struct {
	OptionGroup OptionGroup `xml:"ModifyOptionGroupResult>OptionGroup"`
}

// Option is undocumented.
type Option struct {
	DBSecurityGroupMemberships  []DBSecurityGroupMembership  `xml:"DBSecurityGroupMemberships>DBSecurityGroup"`
	OptionDescription           string                       `xml:"OptionDescription"`
	OptionName                  string                       `xml:"OptionName"`
	OptionSettings              []OptionSetting              `xml:"OptionSettings>OptionSetting"`
	Permanent                   bool                         `xml:"Permanent"`
	Persistent                  bool                         `xml:"Persistent"`
	Port                        int                          `xml:"Port"`
	VpcSecurityGroupMemberships []VpcSecurityGroupMembership `xml:"VpcSecurityGroupMemberships>VpcSecurityGroupMembership"`
}

// OptionConfiguration is undocumented.
type OptionConfiguration struct {
	DBSecurityGroupMemberships  []string        `xml:"DBSecurityGroupMemberships>DBSecurityGroupName"`
	OptionName                  string          `xml:"OptionName"`
	OptionSettings              []OptionSetting `xml:"OptionSettings>OptionSetting"`
	Port                        int             `xml:"Port"`
	VpcSecurityGroupMemberships []string        `xml:"VpcSecurityGroupMemberships>VpcSecurityGroupId"`
}

// OptionGroup is undocumented.
type OptionGroup struct {
	AllowsVpcAndNonVpcInstanceMemberships bool     `xml:"AllowsVpcAndNonVpcInstanceMemberships"`
	EngineName                            string   `xml:"EngineName"`
	MajorEngineVersion                    string   `xml:"MajorEngineVersion"`
	OptionGroupDescription                string   `xml:"OptionGroupDescription"`
	OptionGroupName                       string   `xml:"OptionGroupName"`
	Options                               []Option `xml:"Options>Option"`
	VpcID                                 string   `xml:"VpcId"`
}

// OptionGroupMembership is undocumented.
type OptionGroupMembership struct {
	OptionGroupName string `xml:"OptionGroupName"`
	Status          string `xml:"Status"`
}

// OptionGroupOption is undocumented.
type OptionGroupOption struct {
	DefaultPort                       int                        `xml:"DefaultPort"`
	Description                       string                     `xml:"Description"`
	EngineName                        string                     `xml:"EngineName"`
	MajorEngineVersion                string                     `xml:"MajorEngineVersion"`
	MinimumRequiredMinorEngineVersion string                     `xml:"MinimumRequiredMinorEngineVersion"`
	Name                              string                     `xml:"Name"`
	OptionGroupOptionSettings         []OptionGroupOptionSetting `xml:"OptionGroupOptionSettings>OptionGroupOptionSetting"`
	OptionsDependedOn                 []string                   `xml:"OptionsDependedOn>OptionName"`
	Permanent                         bool                       `xml:"Permanent"`
	Persistent                        bool                       `xml:"Persistent"`
	PortRequired                      bool                       `xml:"PortRequired"`
}

// OptionGroupOptionSetting is undocumented.
type OptionGroupOptionSetting struct {
	AllowedValues      string `xml:"AllowedValues"`
	ApplyType          string `xml:"ApplyType"`
	DefaultValue       string `xml:"DefaultValue"`
	IsModifiable       bool   `xml:"IsModifiable"`
	SettingDescription string `xml:"SettingDescription"`
	SettingName        string `xml:"SettingName"`
}

// OptionGroupOptionsMessage is undocumented.
type OptionGroupOptionsMessage struct {
	Marker             string              `xml:"DescribeOptionGroupOptionsResult>Marker"`
	OptionGroupOptions []OptionGroupOption `xml:"DescribeOptionGroupOptionsResult>OptionGroupOptions>OptionGroupOption"`
}

// OptionGroups is undocumented.
type OptionGroups struct {
	Marker           string        `xml:"DescribeOptionGroupsResult>Marker"`
	OptionGroupsList []OptionGroup `xml:"DescribeOptionGroupsResult>OptionGroupsList>OptionGroup"`
}

// OptionSetting is undocumented.
type OptionSetting struct {
	AllowedValues string `xml:"AllowedValues"`
	ApplyType     string `xml:"ApplyType"`
	DataType      string `xml:"DataType"`
	DefaultValue  string `xml:"DefaultValue"`
	Description   string `xml:"Description"`
	IsCollection  bool   `xml:"IsCollection"`
	IsModifiable  bool   `xml:"IsModifiable"`
	Name          string `xml:"Name"`
	Value         string `xml:"Value"`
}

// OrderableDBInstanceOption is undocumented.
type OrderableDBInstanceOption struct {
	AvailabilityZones  []AvailabilityZone `xml:"AvailabilityZones>AvailabilityZone"`
	DBInstanceClass    string             `xml:"DBInstanceClass"`
	Engine             string             `xml:"Engine"`
	EngineVersion      string             `xml:"EngineVersion"`
	LicenseModel       string             `xml:"LicenseModel"`
	MultiAZCapable     bool               `xml:"MultiAZCapable"`
	ReadReplicaCapable bool               `xml:"ReadReplicaCapable"`
	StorageType        string             `xml:"StorageType"`
	SupportsIops       bool               `xml:"SupportsIops"`
	Vpc                bool               `xml:"Vpc"`
}

// OrderableDBInstanceOptionsMessage is undocumented.
type OrderableDBInstanceOptionsMessage struct {
	Marker                     string                      `xml:"DescribeOrderableDBInstanceOptionsResult>Marker"`
	OrderableDBInstanceOptions []OrderableDBInstanceOption `xml:"DescribeOrderableDBInstanceOptionsResult>OrderableDBInstanceOptions>OrderableDBInstanceOption"`
}

// Parameter is undocumented.
type Parameter struct {
	AllowedValues        string `xml:"AllowedValues"`
	ApplyMethod          string `xml:"ApplyMethod"`
	ApplyType            string `xml:"ApplyType"`
	DataType             string `xml:"DataType"`
	Description          string `xml:"Description"`
	IsModifiable         bool   `xml:"IsModifiable"`
	MinimumEngineVersion string `xml:"MinimumEngineVersion"`
	ParameterName        string `xml:"ParameterName"`
	ParameterValue       string `xml:"ParameterValue"`
	Source               string `xml:"Source"`
}

// PendingModifiedValues is undocumented.
type PendingModifiedValues struct {
	AllocatedStorage      int    `xml:"AllocatedStorage"`
	BackupRetentionPeriod int    `xml:"BackupRetentionPeriod"`
	DBInstanceClass       string `xml:"DBInstanceClass"`
	DBInstanceIdentifier  string `xml:"DBInstanceIdentifier"`
	EngineVersion         string `xml:"EngineVersion"`
	Iops                  int    `xml:"Iops"`
	MasterUserPassword    string `xml:"MasterUserPassword"`
	MultiAZ               bool   `xml:"MultiAZ"`
	Port                  int    `xml:"Port"`
	StorageType           string `xml:"StorageType"`
}

// PromoteReadReplicaMessage is undocumented.
type PromoteReadReplicaMessage struct {
	BackupRetentionPeriod int    `xml:"BackupRetentionPeriod"`
	DBInstanceIdentifier  string `xml:"DBInstanceIdentifier"`
	PreferredBackupWindow string `xml:"PreferredBackupWindow"`
}

// PromoteReadReplicaResult is undocumented.
type PromoteReadReplicaResult struct {
	DBInstance DBInstance `xml:"PromoteReadReplicaResult>DBInstance"`
}

// PurchaseReservedDBInstancesOfferingMessage is undocumented.
type PurchaseReservedDBInstancesOfferingMessage struct {
	DBInstanceCount               int    `xml:"DBInstanceCount"`
	ReservedDBInstanceID          string `xml:"ReservedDBInstanceId"`
	ReservedDBInstancesOfferingID string `xml:"ReservedDBInstancesOfferingId"`
	Tags                          []Tag  `xml:"Tags>Tag"`
}

// PurchaseReservedDBInstancesOfferingResult is undocumented.
type PurchaseReservedDBInstancesOfferingResult struct {
	ReservedDBInstance ReservedDBInstance `xml:"PurchaseReservedDBInstancesOfferingResult>ReservedDBInstance"`
}

// RebootDBInstanceMessage is undocumented.
type RebootDBInstanceMessage struct {
	DBInstanceIdentifier string `xml:"DBInstanceIdentifier"`
	ForceFailover        bool   `xml:"ForceFailover"`
}

// RebootDBInstanceResult is undocumented.
type RebootDBInstanceResult struct {
	DBInstance DBInstance `xml:"RebootDBInstanceResult>DBInstance"`
}

// RecurringCharge is undocumented.
type RecurringCharge struct {
	RecurringChargeAmount    float64 `xml:"RecurringChargeAmount"`
	RecurringChargeFrequency string  `xml:"RecurringChargeFrequency"`
}

// RemoveSourceIdentifierFromSubscriptionMessage is undocumented.
type RemoveSourceIdentifierFromSubscriptionMessage struct {
	SourceIdentifier string `xml:"SourceIdentifier"`
	SubscriptionName string `xml:"SubscriptionName"`
}

// RemoveSourceIdentifierFromSubscriptionResult is undocumented.
type RemoveSourceIdentifierFromSubscriptionResult struct {
	EventSubscription EventSubscription `xml:"RemoveSourceIdentifierFromSubscriptionResult>EventSubscription"`
}

// RemoveTagsFromResourceMessage is undocumented.
type RemoveTagsFromResourceMessage struct {
	ResourceName string   `xml:"ResourceName"`
	TagKeys      []string `xml:"TagKeys>member"`
}

// ReservedDBInstance is undocumented.
type ReservedDBInstance struct {
	CurrencyCode                  string            `xml:"CurrencyCode"`
	DBInstanceClass               string            `xml:"DBInstanceClass"`
	DBInstanceCount               int               `xml:"DBInstanceCount"`
	Duration                      int               `xml:"Duration"`
	FixedPrice                    float64           `xml:"FixedPrice"`
	MultiAZ                       bool              `xml:"MultiAZ"`
	OfferingType                  string            `xml:"OfferingType"`
	ProductDescription            string            `xml:"ProductDescription"`
	RecurringCharges              []RecurringCharge `xml:"RecurringCharges>RecurringCharge"`
	ReservedDBInstanceID          string            `xml:"ReservedDBInstanceId"`
	ReservedDBInstancesOfferingID string            `xml:"ReservedDBInstancesOfferingId"`
	StartTime                     time.Time         `xml:"StartTime"`
	State                         string            `xml:"State"`
	UsagePrice                    float64           `xml:"UsagePrice"`
}

// ReservedDBInstanceMessage is undocumented.
type ReservedDBInstanceMessage struct {
	Marker              string               `xml:"DescribeReservedDBInstancesResult>Marker"`
	ReservedDBInstances []ReservedDBInstance `xml:"DescribeReservedDBInstancesResult>ReservedDBInstances>ReservedDBInstance"`
}

// ReservedDBInstancesOffering is undocumented.
type ReservedDBInstancesOffering struct {
	CurrencyCode                  string            `xml:"CurrencyCode"`
	DBInstanceClass               string            `xml:"DBInstanceClass"`
	Duration                      int               `xml:"Duration"`
	FixedPrice                    float64           `xml:"FixedPrice"`
	MultiAZ                       bool              `xml:"MultiAZ"`
	OfferingType                  string            `xml:"OfferingType"`
	ProductDescription            string            `xml:"ProductDescription"`
	RecurringCharges              []RecurringCharge `xml:"RecurringCharges>RecurringCharge"`
	ReservedDBInstancesOfferingID string            `xml:"ReservedDBInstancesOfferingId"`
	UsagePrice                    float64           `xml:"UsagePrice"`
}

// ReservedDBInstancesOfferingMessage is undocumented.
type ReservedDBInstancesOfferingMessage struct {
	Marker                       string                        `xml:"DescribeReservedDBInstancesOfferingsResult>Marker"`
	ReservedDBInstancesOfferings []ReservedDBInstancesOffering `xml:"DescribeReservedDBInstancesOfferingsResult>ReservedDBInstancesOfferings>ReservedDBInstancesOffering"`
}

// ResetDBParameterGroupMessage is undocumented.
type ResetDBParameterGroupMessage struct {
	DBParameterGroupName string      `xml:"DBParameterGroupName"`
	Parameters           []Parameter `xml:"Parameters>Parameter"`
	ResetAllParameters   bool        `xml:"ResetAllParameters"`
}

// RestoreDBInstanceFromDBSnapshotMessage is undocumented.
type RestoreDBInstanceFromDBSnapshotMessage struct {
	AutoMinorVersionUpgrade bool   `xml:"AutoMinorVersionUpgrade"`
	AvailabilityZone        string `xml:"AvailabilityZone"`
	DBInstanceClass         string `xml:"DBInstanceClass"`
	DBInstanceIdentifier    string `xml:"DBInstanceIdentifier"`
	DBName                  string `xml:"DBName"`
	DBSnapshotIdentifier    string `xml:"DBSnapshotIdentifier"`
	DBSubnetGroupName       string `xml:"DBSubnetGroupName"`
	Engine                  string `xml:"Engine"`
	Iops                    int    `xml:"Iops"`
	LicenseModel            string `xml:"LicenseModel"`
	MultiAZ                 bool   `xml:"MultiAZ"`
	OptionGroupName         string `xml:"OptionGroupName"`
	Port                    int    `xml:"Port"`
	PubliclyAccessible      bool   `xml:"PubliclyAccessible"`
	StorageType             string `xml:"StorageType"`
	Tags                    []Tag  `xml:"Tags>Tag"`
	TdeCredentialARN        string `xml:"TdeCredentialArn"`
	TdeCredentialPassword   string `xml:"TdeCredentialPassword"`
}

// RestoreDBInstanceFromDBSnapshotResult is undocumented.
type RestoreDBInstanceFromDBSnapshotResult struct {
	DBInstance DBInstance `xml:"RestoreDBInstanceFromDBSnapshotResult>DBInstance"`
}

// RestoreDBInstanceToPointInTimeMessage is undocumented.
type RestoreDBInstanceToPointInTimeMessage struct {
	AutoMinorVersionUpgrade    bool      `xml:"AutoMinorVersionUpgrade"`
	AvailabilityZone           string    `xml:"AvailabilityZone"`
	DBInstanceClass            string    `xml:"DBInstanceClass"`
	DBName                     string    `xml:"DBName"`
	DBSubnetGroupName          string    `xml:"DBSubnetGroupName"`
	Engine                     string    `xml:"Engine"`
	Iops                       int       `xml:"Iops"`
	LicenseModel               string    `xml:"LicenseModel"`
	MultiAZ                    bool      `xml:"MultiAZ"`
	OptionGroupName            string    `xml:"OptionGroupName"`
	Port                       int       `xml:"Port"`
	PubliclyAccessible         bool      `xml:"PubliclyAccessible"`
	RestoreTime                time.Time `xml:"RestoreTime"`
	SourceDBInstanceIdentifier string    `xml:"SourceDBInstanceIdentifier"`
	StorageType                string    `xml:"StorageType"`
	Tags                       []Tag     `xml:"Tags>Tag"`
	TargetDBInstanceIdentifier string    `xml:"TargetDBInstanceIdentifier"`
	TdeCredentialARN           string    `xml:"TdeCredentialArn"`
	TdeCredentialPassword      string    `xml:"TdeCredentialPassword"`
	UseLatestRestorableTime    bool      `xml:"UseLatestRestorableTime"`
}

// RestoreDBInstanceToPointInTimeResult is undocumented.
type RestoreDBInstanceToPointInTimeResult struct {
	DBInstance DBInstance `xml:"RestoreDBInstanceToPointInTimeResult>DBInstance"`
}

// RevokeDBSecurityGroupIngressMessage is undocumented.
type RevokeDBSecurityGroupIngressMessage struct {
	CIDRIP                  string `xml:"CIDRIP"`
	DBSecurityGroupName     string `xml:"DBSecurityGroupName"`
	EC2SecurityGroupID      string `xml:"EC2SecurityGroupId"`
	EC2SecurityGroupName    string `xml:"EC2SecurityGroupName"`
	EC2SecurityGroupOwnerID string `xml:"EC2SecurityGroupOwnerId"`
}

// RevokeDBSecurityGroupIngressResult is undocumented.
type RevokeDBSecurityGroupIngressResult struct {
	DBSecurityGroup DBSecurityGroup `xml:"RevokeDBSecurityGroupIngressResult>DBSecurityGroup"`
}

// Subnet is undocumented.
type Subnet struct {
	SubnetAvailabilityZone AvailabilityZone `xml:"SubnetAvailabilityZone"`
	SubnetIdentifier       string           `xml:"SubnetIdentifier"`
	SubnetStatus           string           `xml:"SubnetStatus"`
}

// Tag is undocumented.
type Tag struct {
	Key   string `xml:"Key"`
	Value string `xml:"Value"`
}

// TagListMessage is undocumented.
type TagListMessage struct {
	TagList []Tag `xml:"ListTagsForResourceResult>TagList>Tag"`
}

// VpcSecurityGroupMembership is undocumented.
type VpcSecurityGroupMembership struct {
	Status             string `xml:"Status"`
	VpcSecurityGroupID string `xml:"VpcSecurityGroupId"`
}

// DescribeDBEngineVersionsResult is a wrapper for DBEngineVersionMessage.
type DescribeDBEngineVersionsResult struct {
	XMLName xml.Name `xml:"DescribeDBEngineVersionsResponse"`

	DBEngineVersions []DBEngineVersion `xml:"DescribeDBEngineVersionsResult>DBEngineVersions>DBEngineVersion"`
	Marker           string            `xml:"DescribeDBEngineVersionsResult>Marker"`
}

// DescribeDBInstancesResult is a wrapper for DBInstanceMessage.
type DescribeDBInstancesResult struct {
	XMLName xml.Name `xml:"DescribeDBInstancesResponse"`

	DBInstances []DBInstance `xml:"DescribeDBInstancesResult>DBInstances>DBInstance"`
	Marker      string       `xml:"DescribeDBInstancesResult>Marker"`
}

// DescribeDBLogFilesResult is a wrapper for DescribeDBLogFilesResponse.
type DescribeDBLogFilesResult struct {
	XMLName xml.Name `xml:"DescribeDBLogFilesResponse"`

	DescribeDBLogFiles []DescribeDBLogFilesDetails `xml:"DescribeDBLogFilesResult>DescribeDBLogFiles>DescribeDBLogFilesDetails"`
	Marker             string                      `xml:"DescribeDBLogFilesResult>Marker"`
}

// DescribeDBParameterGroupsResult is a wrapper for DBParameterGroupsMessage.
type DescribeDBParameterGroupsResult struct {
	XMLName xml.Name `xml:"DescribeDBParameterGroupsResponse"`

	DBParameterGroups []DBParameterGroup `xml:"DescribeDBParameterGroupsResult>DBParameterGroups>DBParameterGroup"`
	Marker            string             `xml:"DescribeDBParameterGroupsResult>Marker"`
}

// DescribeDBParametersResult is a wrapper for DBParameterGroupDetails.
type DescribeDBParametersResult struct {
	XMLName xml.Name `xml:"DescribeDBParametersResponse"`

	Marker     string      `xml:"DescribeDBParametersResult>Marker"`
	Parameters []Parameter `xml:"DescribeDBParametersResult>Parameters>Parameter"`
}

// DescribeDBSecurityGroupsResult is a wrapper for DBSecurityGroupMessage.
type DescribeDBSecurityGroupsResult struct {
	XMLName xml.Name `xml:"DescribeDBSecurityGroupsResponse"`

	DBSecurityGroups []DBSecurityGroup `xml:"DescribeDBSecurityGroupsResult>DBSecurityGroups>DBSecurityGroup"`
	Marker           string            `xml:"DescribeDBSecurityGroupsResult>Marker"`
}

// DescribeDBSnapshotsResult is a wrapper for DBSnapshotMessage.
type DescribeDBSnapshotsResult struct {
	XMLName xml.Name `xml:"DescribeDBSnapshotsResponse"`

	DBSnapshots []DBSnapshot `xml:"DescribeDBSnapshotsResult>DBSnapshots>DBSnapshot"`
	Marker      string       `xml:"DescribeDBSnapshotsResult>Marker"`
}

// DescribeDBSubnetGroupsResult is a wrapper for DBSubnetGroupMessage.
type DescribeDBSubnetGroupsResult struct {
	XMLName xml.Name `xml:"DescribeDBSubnetGroupsResponse"`

	DBSubnetGroups []DBSubnetGroup `xml:"DescribeDBSubnetGroupsResult>DBSubnetGroups>DBSubnetGroup"`
	Marker         string          `xml:"DescribeDBSubnetGroupsResult>Marker"`
}

// DescribeEventCategoriesResult is a wrapper for EventCategoriesMessage.
type DescribeEventCategoriesResult struct {
	XMLName xml.Name `xml:"DescribeEventCategoriesResponse"`

	EventCategoriesMapList []EventCategoriesMap `xml:"DescribeEventCategoriesResult>EventCategoriesMapList>EventCategoriesMap"`
}

// DescribeEventSubscriptionsResult is a wrapper for EventSubscriptionsMessage.
type DescribeEventSubscriptionsResult struct {
	XMLName xml.Name `xml:"DescribeEventSubscriptionsResponse"`

	EventSubscriptionsList []EventSubscription `xml:"DescribeEventSubscriptionsResult>EventSubscriptionsList>EventSubscription"`
	Marker                 string              `xml:"DescribeEventSubscriptionsResult>Marker"`
}

// DescribeEventsResult is a wrapper for EventsMessage.
type DescribeEventsResult struct {
	XMLName xml.Name `xml:"DescribeEventsResponse"`

	Events []Event `xml:"DescribeEventsResult>Events>Event"`
	Marker string  `xml:"DescribeEventsResult>Marker"`
}

// DescribeOptionGroupOptionsResult is a wrapper for OptionGroupOptionsMessage.
type DescribeOptionGroupOptionsResult struct {
	XMLName xml.Name `xml:"DescribeOptionGroupOptionsResponse"`

	Marker             string              `xml:"DescribeOptionGroupOptionsResult>Marker"`
	OptionGroupOptions []OptionGroupOption `xml:"DescribeOptionGroupOptionsResult>OptionGroupOptions>OptionGroupOption"`
}

// DescribeOptionGroupsResult is a wrapper for OptionGroups.
type DescribeOptionGroupsResult struct {
	XMLName xml.Name `xml:"DescribeOptionGroupsResponse"`

	Marker           string        `xml:"DescribeOptionGroupsResult>Marker"`
	OptionGroupsList []OptionGroup `xml:"DescribeOptionGroupsResult>OptionGroupsList>OptionGroup"`
}

// DescribeOrderableDBInstanceOptionsResult is a wrapper for OrderableDBInstanceOptionsMessage.
type DescribeOrderableDBInstanceOptionsResult struct {
	XMLName xml.Name `xml:"DescribeOrderableDBInstanceOptionsResponse"`

	Marker                     string                      `xml:"DescribeOrderableDBInstanceOptionsResult>Marker"`
	OrderableDBInstanceOptions []OrderableDBInstanceOption `xml:"DescribeOrderableDBInstanceOptionsResult>OrderableDBInstanceOptions>OrderableDBInstanceOption"`
}

// DescribeReservedDBInstancesOfferingsResult is a wrapper for ReservedDBInstancesOfferingMessage.
type DescribeReservedDBInstancesOfferingsResult struct {
	XMLName xml.Name `xml:"DescribeReservedDBInstancesOfferingsResponse"`

	Marker                       string                        `xml:"DescribeReservedDBInstancesOfferingsResult>Marker"`
	ReservedDBInstancesOfferings []ReservedDBInstancesOffering `xml:"DescribeReservedDBInstancesOfferingsResult>ReservedDBInstancesOfferings>ReservedDBInstancesOffering"`
}

// DescribeReservedDBInstancesResult is a wrapper for ReservedDBInstanceMessage.
type DescribeReservedDBInstancesResult struct {
	XMLName xml.Name `xml:"DescribeReservedDBInstancesResponse"`

	Marker              string               `xml:"DescribeReservedDBInstancesResult>Marker"`
	ReservedDBInstances []ReservedDBInstance `xml:"DescribeReservedDBInstancesResult>ReservedDBInstances>ReservedDBInstance"`
}

// DownloadDBLogFilePortionResult is a wrapper for DownloadDBLogFilePortionDetails.
type DownloadDBLogFilePortionResult struct {
	XMLName xml.Name `xml:"DownloadDBLogFilePortionResponse"`

	AdditionalDataPending bool   `xml:"DownloadDBLogFilePortionResult>AdditionalDataPending"`
	LogFileData           string `xml:"DownloadDBLogFilePortionResult>LogFileData"`
	Marker                string `xml:"DownloadDBLogFilePortionResult>Marker"`
}

// ListTagsForResourceResult is a wrapper for TagListMessage.
type ListTagsForResourceResult struct {
	XMLName xml.Name `xml:"ListTagsForResourceResponse"`

	TagList []Tag `xml:"ListTagsForResourceResult>TagList>Tag"`
}

// ModifyDBParameterGroupResult is a wrapper for DBParameterGroupNameMessage.
type ModifyDBParameterGroupResult struct {
	XMLName xml.Name `xml:"Response"`

	DBParameterGroupName string `xml:"ModifyDBParameterGroupResult>DBParameterGroupName"`
}

// ResetDBParameterGroupResult is a wrapper for DBParameterGroupNameMessage.
type ResetDBParameterGroupResult struct {
	XMLName xml.Name `xml:"Response"`

	DBParameterGroupName string `xml:"ResetDBParameterGroupResult>DBParameterGroupName"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
