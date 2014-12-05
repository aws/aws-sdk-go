// Package opsworks provides a client for AWS OpsWorks.
package opsworks

import (
	"fmt"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
)

// OpsWorks is a client for AWS OpsWorks.
type OpsWorks struct {
	client aws.Client
}

// New returns a new OpsWorks client.
func New(key, secret, region string, client *http.Client) *OpsWorks {
	if client == nil {
		client = http.DefaultClient
	}

	return &OpsWorks{
		client: &aws.JSONClient{
			Client:       client,
			Region:       region,
			Endpoint:     fmt.Sprintf("https://opsworks.%s.amazonaws.com", region),
			Prefix:       "opsworks",
			Key:          key,
			Secret:       secret,
			JSONVersion:  "1.1",
			TargetPrefix: "OpsWorks_20130218",
		},
	}
}

// AssignVolume assigns one of the stack's registered Amazon EBS volumes to
// a specified instance. The volume must first be registered with the stack
// by calling RegisterVolume . For more information, see Resource
// Management Required Permissions : To use this action, an IAM user must
// have a Manage permissions level for the stack, or an attached policy
// that explicitly grants permissions. For more information on user
// permissions, see Managing User Permissions
func (c *OpsWorks) AssignVolume(req AssignVolumeRequest) (err error) {
	// NRE
	err = c.client.Do("AssignVolume", "POST", "/", req, nil)
	return
}

// AssociateElasticIP associates one of the stack's registered Elastic IP
// addresses with a specified instance. The address must first be
// registered with the stack by calling RegisterElasticIp . For more
// information, see Resource Management Required Permissions : To use this
// action, an IAM user must have a Manage permissions level for the stack,
// or an attached policy that explicitly grants permissions. For more
// information on user permissions, see Managing User Permissions
func (c *OpsWorks) AssociateElasticIP(req AssociateElasticIPRequest) (err error) {
	// NRE
	err = c.client.Do("AssociateElasticIp", "POST", "/", req, nil)
	return
}

// AttachElasticLoadBalancer attaches an Elastic Load Balancing load
// balancer to a specified layer. For more information, see Elastic Load
// Balancing Required Permissions : To use this action, an IAM user must
// have a Manage permissions level for the stack, or an attached policy
// that explicitly grants permissions. For more information on user
// permissions, see Managing User Permissions
func (c *OpsWorks) AttachElasticLoadBalancer(req AttachElasticLoadBalancerRequest) (err error) {
	// NRE
	err = c.client.Do("AttachElasticLoadBalancer", "POST", "/", req, nil)
	return
}

// CloneStack creates a clone of a specified stack. For more information,
// see Clone a Stack Required Permissions : To use this action, an IAM user
// must have an attached policy that explicitly grants permissions. For
// more information on user permissions, see Managing User Permissions
func (c *OpsWorks) CloneStack(req CloneStackRequest) (resp *CloneStackResult, err error) {
	resp = &CloneStackResult{}
	err = c.client.Do("CloneStack", "POST", "/", req, resp)
	return
}

// CreateApp creates an app for a specified stack. For more information,
// see Creating Apps Required Permissions : To use this action, an IAM user
// must have a Manage permissions level for the stack, or an attached
// policy that explicitly grants permissions. For more information on user
// permissions, see Managing User Permissions
func (c *OpsWorks) CreateApp(req CreateAppRequest) (resp *CreateAppResult, err error) {
	resp = &CreateAppResult{}
	err = c.client.Do("CreateApp", "POST", "/", req, resp)
	return
}

// CreateDeployment deploys a stack or app. App deployment generates a
// deploy event, which runs the associated recipes and passes them a stack
// configuration object that includes information about the app. Stack
// deployment runs the deploy recipes but does not raise an event. For more
// information, see Deploying Apps and Run Stack Commands Required
// Permissions : To use this action, an IAM user must have a Deploy or
// Manage permissions level for the stack, or an attached policy that
// explicitly grants permissions. For more information on user permissions,
// see Managing User Permissions
func (c *OpsWorks) CreateDeployment(req CreateDeploymentRequest) (resp *CreateDeploymentResult, err error) {
	resp = &CreateDeploymentResult{}
	err = c.client.Do("CreateDeployment", "POST", "/", req, resp)
	return
}

// CreateInstance creates an instance in a specified stack. For more
// information, see Adding an Instance to a Layer Required Permissions : To
// use this action, an IAM user must have a Manage permissions level for
// the stack, or an attached policy that explicitly grants permissions. For
// more information on user permissions, see Managing User Permissions
func (c *OpsWorks) CreateInstance(req CreateInstanceRequest) (resp *CreateInstanceResult, err error) {
	resp = &CreateInstanceResult{}
	err = c.client.Do("CreateInstance", "POST", "/", req, resp)
	return
}

// CreateLayer creates a layer. For more information, see How to Create a
// Layer Required Permissions : To use this action, an IAM user must have a
// Manage permissions level for the stack, or an attached policy that
// explicitly grants permissions. For more information on user permissions,
// see Managing User Permissions
func (c *OpsWorks) CreateLayer(req CreateLayerRequest) (resp *CreateLayerResult, err error) {
	resp = &CreateLayerResult{}
	err = c.client.Do("CreateLayer", "POST", "/", req, resp)
	return
}

// CreateStack creates a new stack. For more information, see Create a New
// Stack Required Permissions : To use this action, an IAM user must have
// an attached policy that explicitly grants permissions. For more
// information on user permissions, see Managing User Permissions
func (c *OpsWorks) CreateStack(req CreateStackRequest) (resp *CreateStackResult, err error) {
	resp = &CreateStackResult{}
	err = c.client.Do("CreateStack", "POST", "/", req, resp)
	return
}

// CreateUserProfile creates a new user profile. Required Permissions : To
// use this action, an IAM user must have an attached policy that
// explicitly grants permissions. For more information on user permissions,
// see Managing User Permissions
func (c *OpsWorks) CreateUserProfile(req CreateUserProfileRequest) (resp *CreateUserProfileResult, err error) {
	resp = &CreateUserProfileResult{}
	err = c.client.Do("CreateUserProfile", "POST", "/", req, resp)
	return
}

// DeleteApp deletes a specified app. Required Permissions : To use this
// action, an IAM user must have a Manage permissions level for the stack,
// or an attached policy that explicitly grants permissions. For more
// information on user permissions, see Managing User Permissions
func (c *OpsWorks) DeleteApp(req DeleteAppRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteApp", "POST", "/", req, nil)
	return
}

// DeleteInstance deletes a specified instance. You must stop an instance
// before you can delete it. For more information, see Deleting Instances
// Required Permissions : To use this action, an IAM user must have a
// Manage permissions level for the stack, or an attached policy that
// explicitly grants permissions. For more information on user permissions,
// see Managing User Permissions
func (c *OpsWorks) DeleteInstance(req DeleteInstanceRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteInstance", "POST", "/", req, nil)
	return
}

// DeleteLayer deletes a specified layer. You must first stop and then
// delete all associated instances. For more information, see How to Delete
// a Layer Required Permissions : To use this action, an IAM user must have
// a Manage permissions level for the stack, or an attached policy that
// explicitly grants permissions. For more information on user permissions,
// see Managing User Permissions
func (c *OpsWorks) DeleteLayer(req DeleteLayerRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteLayer", "POST", "/", req, nil)
	return
}

// DeleteStack deletes a specified stack. You must first delete all
// instances, layers, and apps. For more information, see Shut Down a Stack
// Required Permissions : To use this action, an IAM user must have a
// Manage permissions level for the stack, or an attached policy that
// explicitly grants permissions. For more information on user permissions,
// see Managing User Permissions
func (c *OpsWorks) DeleteStack(req DeleteStackRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteStack", "POST", "/", req, nil)
	return
}

// DeleteUserProfile deletes a user profile. Required Permissions : To use
// this action, an IAM user must have an attached policy that explicitly
// grants permissions. For more information on user permissions, see
// Managing User Permissions
func (c *OpsWorks) DeleteUserProfile(req DeleteUserProfileRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteUserProfile", "POST", "/", req, nil)
	return
}

// DeregisterElasticIP deregisters a specified Elastic IP address. The
// address can then be registered by another stack. For more information,
// see Resource Management Required Permissions : To use this action, an
// IAM user must have a Manage permissions level for the stack, or an
// attached policy that explicitly grants permissions. For more information
// on user permissions, see Managing User Permissions
func (c *OpsWorks) DeregisterElasticIP(req DeregisterElasticIPRequest) (err error) {
	// NRE
	err = c.client.Do("DeregisterElasticIp", "POST", "/", req, nil)
	return
}

// DeregisterRdsDbInstance is undocumented.
func (c *OpsWorks) DeregisterRdsDbInstance(req DeregisterRdsDbInstanceRequest) (err error) {
	// NRE
	err = c.client.Do("DeregisterRdsDbInstance", "POST", "/", req, nil)
	return
}

// DeregisterVolume deregisters an Amazon EBS volume. The volume can then
// be registered by another stack. For more information, see Resource
// Management Required Permissions : To use this action, an IAM user must
// have a Manage permissions level for the stack, or an attached policy
// that explicitly grants permissions. For more information on user
// permissions, see Managing User Permissions
func (c *OpsWorks) DeregisterVolume(req DeregisterVolumeRequest) (err error) {
	// NRE
	err = c.client.Do("DeregisterVolume", "POST", "/", req, nil)
	return
}

// DescribeApps requests a description of a specified set of apps. Required
// Permissions : To use this action, an IAM user must have a Show, Deploy,
// or Manage permissions level for the stack, or an attached policy that
// explicitly grants permissions. For more information on user permissions,
// see Managing User Permissions
func (c *OpsWorks) DescribeApps(req DescribeAppsRequest) (resp *DescribeAppsResult, err error) {
	resp = &DescribeAppsResult{}
	err = c.client.Do("DescribeApps", "POST", "/", req, resp)
	return
}

// DescribeCommands describes the results of specified commands. Required
// Permissions : To use this action, an IAM user must have a Show, Deploy,
// or Manage permissions level for the stack, or an attached policy that
// explicitly grants permissions. For more information on user permissions,
// see Managing User Permissions
func (c *OpsWorks) DescribeCommands(req DescribeCommandsRequest) (resp *DescribeCommandsResult, err error) {
	resp = &DescribeCommandsResult{}
	err = c.client.Do("DescribeCommands", "POST", "/", req, resp)
	return
}

// DescribeDeployments requests a description of a specified set of
// deployments. Required Permissions : To use this action, an IAM user must
// have a Show, Deploy, or Manage permissions level for the stack, or an
// attached policy that explicitly grants permissions. For more information
// on user permissions, see Managing User Permissions
func (c *OpsWorks) DescribeDeployments(req DescribeDeploymentsRequest) (resp *DescribeDeploymentsResult, err error) {
	resp = &DescribeDeploymentsResult{}
	err = c.client.Do("DescribeDeployments", "POST", "/", req, resp)
	return
}

// DescribeElasticIPs describes Elastic IP addresses Required Permissions :
// To use this action, an IAM user must have a Show, Deploy, or Manage
// permissions level for the stack, or an attached policy that explicitly
// grants permissions. For more information on user permissions, see
// Managing User Permissions
func (c *OpsWorks) DescribeElasticIPs(req DescribeElasticIPsRequest) (resp *DescribeElasticIPsResult, err error) {
	resp = &DescribeElasticIPsResult{}
	err = c.client.Do("DescribeElasticIps", "POST", "/", req, resp)
	return
}

// DescribeElasticLoadBalancers describes a stack's Elastic Load Balancing
// instances. Required Permissions : To use this action, an IAM user must
// have a Show, Deploy, or Manage permissions level for the stack, or an
// attached policy that explicitly grants permissions. For more information
// on user permissions, see Managing User Permissions
func (c *OpsWorks) DescribeElasticLoadBalancers(req DescribeElasticLoadBalancersRequest) (resp *DescribeElasticLoadBalancersResult, err error) {
	resp = &DescribeElasticLoadBalancersResult{}
	err = c.client.Do("DescribeElasticLoadBalancers", "POST", "/", req, resp)
	return
}

// DescribeInstances requests a description of a set of instances. Required
// Permissions : To use this action, an IAM user must have a Show, Deploy,
// or Manage permissions level for the stack, or an attached policy that
// explicitly grants permissions. For more information on user permissions,
// see Managing User Permissions
func (c *OpsWorks) DescribeInstances(req DescribeInstancesRequest) (resp *DescribeInstancesResult, err error) {
	resp = &DescribeInstancesResult{}
	err = c.client.Do("DescribeInstances", "POST", "/", req, resp)
	return
}

// DescribeLayers requests a description of one or more layers in a
// specified stack. Required Permissions : To use this action, an IAM user
// must have a Show, Deploy, or Manage permissions level for the stack, or
// an attached policy that explicitly grants permissions. For more
// information on user permissions, see Managing User Permissions
func (c *OpsWorks) DescribeLayers(req DescribeLayersRequest) (resp *DescribeLayersResult, err error) {
	resp = &DescribeLayersResult{}
	err = c.client.Do("DescribeLayers", "POST", "/", req, resp)
	return
}

// DescribeLoadBasedAutoScaling describes load-based auto scaling
// configurations for specified layers. Required Permissions : To use this
// action, an IAM user must have a Show, Deploy, or Manage permissions
// level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions
func (c *OpsWorks) DescribeLoadBasedAutoScaling(req DescribeLoadBasedAutoScalingRequest) (resp *DescribeLoadBasedAutoScalingResult, err error) {
	resp = &DescribeLoadBasedAutoScalingResult{}
	err = c.client.Do("DescribeLoadBasedAutoScaling", "POST", "/", req, resp)
	return
}

// DescribeMyUserProfile describes a user's SSH information. Required
// Permissions : To use this action, an IAM user must have self-management
// enabled or an attached policy that explicitly grants permissions. For
// more information on user permissions, see Managing User Permissions
func (c *OpsWorks) DescribeMyUserProfile() (resp *DescribeMyUserProfileResult, err error) {
	resp = &DescribeMyUserProfileResult{}
	err = c.client.Do("DescribeMyUserProfile", "POST", "/", nil, resp)
	return
}

// DescribePermissions describes the permissions for a specified stack.
// Required Permissions : To use this action, an IAM user must have a
// Manage permissions level for the stack, or an attached policy that
// explicitly grants permissions. For more information on user permissions,
// see Managing User Permissions
func (c *OpsWorks) DescribePermissions(req DescribePermissionsRequest) (resp *DescribePermissionsResult, err error) {
	resp = &DescribePermissionsResult{}
	err = c.client.Do("DescribePermissions", "POST", "/", req, resp)
	return
}

// DescribeRaidArrays describe an instance's arrays. Required Permissions :
// To use this action, an IAM user must have a Show, Deploy, or Manage
// permissions level for the stack, or an attached policy that explicitly
// grants permissions. For more information on user permissions, see
// Managing User Permissions
func (c *OpsWorks) DescribeRaidArrays(req DescribeRaidArraysRequest) (resp *DescribeRaidArraysResult, err error) {
	resp = &DescribeRaidArraysResult{}
	err = c.client.Do("DescribeRaidArrays", "POST", "/", req, resp)
	return
}

// DescribeRdsDbInstances is undocumented.
func (c *OpsWorks) DescribeRdsDbInstances(req DescribeRdsDbInstancesRequest) (resp *DescribeRdsDbInstancesResult, err error) {
	resp = &DescribeRdsDbInstancesResult{}
	err = c.client.Do("DescribeRdsDbInstances", "POST", "/", req, resp)
	return
}

// DescribeServiceErrors describes AWS OpsWorks service errors. Required
// Permissions : To use this action, an IAM user must have a Show, Deploy,
// or Manage permissions level for the stack, or an attached policy that
// explicitly grants permissions. For more information on user permissions,
// see Managing User Permissions
func (c *OpsWorks) DescribeServiceErrors(req DescribeServiceErrorsRequest) (resp *DescribeServiceErrorsResult, err error) {
	resp = &DescribeServiceErrorsResult{}
	err = c.client.Do("DescribeServiceErrors", "POST", "/", req, resp)
	return
}

// DescribeStackSummary describes the number of layers and apps in a
// specified stack, and the number of instances in each state, such as
// running_setup or online Required Permissions : To use this action, an
// IAM user must have a Show, Deploy, or Manage permissions level for the
// stack, or an attached policy that explicitly grants permissions. For
// more information on user permissions, see Managing User Permissions
func (c *OpsWorks) DescribeStackSummary(req DescribeStackSummaryRequest) (resp *DescribeStackSummaryResult, err error) {
	resp = &DescribeStackSummaryResult{}
	err = c.client.Do("DescribeStackSummary", "POST", "/", req, resp)
	return
}

// DescribeStacks requests a description of one or more stacks. Required
// Permissions : To use this action, an IAM user must have a Show, Deploy,
// or Manage permissions level for the stack, or an attached policy that
// explicitly grants permissions. For more information on user permissions,
// see Managing User Permissions
func (c *OpsWorks) DescribeStacks(req DescribeStacksRequest) (resp *DescribeStacksResult, err error) {
	resp = &DescribeStacksResult{}
	err = c.client.Do("DescribeStacks", "POST", "/", req, resp)
	return
}

// DescribeTimeBasedAutoScaling describes time-based auto scaling
// configurations for specified instances. Required Permissions : To use
// this action, an IAM user must have a Show, Deploy, or Manage permissions
// level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions
func (c *OpsWorks) DescribeTimeBasedAutoScaling(req DescribeTimeBasedAutoScalingRequest) (resp *DescribeTimeBasedAutoScalingResult, err error) {
	resp = &DescribeTimeBasedAutoScalingResult{}
	err = c.client.Do("DescribeTimeBasedAutoScaling", "POST", "/", req, resp)
	return
}

// DescribeUserProfiles describe specified users. Required Permissions : To
// use this action, an IAM user must have an attached policy that
// explicitly grants permissions. For more information on user permissions,
// see Managing User Permissions
func (c *OpsWorks) DescribeUserProfiles(req DescribeUserProfilesRequest) (resp *DescribeUserProfilesResult, err error) {
	resp = &DescribeUserProfilesResult{}
	err = c.client.Do("DescribeUserProfiles", "POST", "/", req, resp)
	return
}

// DescribeVolumes describes an instance's Amazon EBS volumes. Required
// Permissions : To use this action, an IAM user must have a Show, Deploy,
// or Manage permissions level for the stack, or an attached policy that
// explicitly grants permissions. For more information on user permissions,
// see Managing User Permissions
func (c *OpsWorks) DescribeVolumes(req DescribeVolumesRequest) (resp *DescribeVolumesResult, err error) {
	resp = &DescribeVolumesResult{}
	err = c.client.Do("DescribeVolumes", "POST", "/", req, resp)
	return
}

// DetachElasticLoadBalancer detaches a specified Elastic Load Balancing
// instance from its layer. Required Permissions : To use this action, an
// IAM user must have a Manage permissions level for the stack, or an
// attached policy that explicitly grants permissions. For more information
// on user permissions, see Managing User Permissions
func (c *OpsWorks) DetachElasticLoadBalancer(req DetachElasticLoadBalancerRequest) (err error) {
	// NRE
	err = c.client.Do("DetachElasticLoadBalancer", "POST", "/", req, nil)
	return
}

// DisassociateElasticIP disassociates an Elastic IP address from its
// instance. The address remains registered with the stack. For more
// information, see Resource Management Required Permissions : To use this
// action, an IAM user must have a Manage permissions level for the stack,
// or an attached policy that explicitly grants permissions. For more
// information on user permissions, see Managing User Permissions
func (c *OpsWorks) DisassociateElasticIP(req DisassociateElasticIPRequest) (err error) {
	// NRE
	err = c.client.Do("DisassociateElasticIp", "POST", "/", req, nil)
	return
}

// GetHostnameSuggestion gets a generated host name for the specified
// layer, based on the current host name theme. Required Permissions : To
// use this action, an IAM user must have a Manage permissions level for
// the stack, or an attached policy that explicitly grants permissions. For
// more information on user permissions, see Managing User Permissions
func (c *OpsWorks) GetHostnameSuggestion(req GetHostnameSuggestionRequest) (resp *GetHostnameSuggestionResult, err error) {
	resp = &GetHostnameSuggestionResult{}
	err = c.client.Do("GetHostnameSuggestion", "POST", "/", req, resp)
	return
}

// RebootInstance reboots a specified instance. For more information, see
// Starting, Stopping, and Rebooting Instances Required Permissions : To
// use this action, an IAM user must have a Manage permissions level for
// the stack, or an attached policy that explicitly grants permissions. For
// more information on user permissions, see Managing User Permissions
func (c *OpsWorks) RebootInstance(req RebootInstanceRequest) (err error) {
	// NRE
	err = c.client.Do("RebootInstance", "POST", "/", req, nil)
	return
}

// RegisterElasticIP registers an Elastic IP address with a specified
// stack. An address can be registered with only one stack at a time. If
// the address is already registered, you must first deregister it by
// calling DeregisterElasticIp . For more information, see Resource
// Management Required Permissions : To use this action, an IAM user must
// have a Manage permissions level for the stack, or an attached policy
// that explicitly grants permissions. For more information on user
// permissions, see Managing User Permissions
func (c *OpsWorks) RegisterElasticIP(req RegisterElasticIPRequest) (resp *RegisterElasticIPResult, err error) {
	resp = &RegisterElasticIPResult{}
	err = c.client.Do("RegisterElasticIp", "POST", "/", req, resp)
	return
}

// RegisterRdsDbInstance is undocumented.
func (c *OpsWorks) RegisterRdsDbInstance(req RegisterRdsDbInstanceRequest) (err error) {
	// NRE
	err = c.client.Do("RegisterRdsDbInstance", "POST", "/", req, nil)
	return
}

// RegisterVolume registers an Amazon EBS volume with a specified stack. A
// volume can be registered with only one stack at a time. If the volume is
// already registered, you must first deregister it by calling
// DeregisterVolume . For more information, see Resource Management
// Required Permissions : To use this action, an IAM user must have a
// Manage permissions level for the stack, or an attached policy that
// explicitly grants permissions. For more information on user permissions,
// see Managing User Permissions
func (c *OpsWorks) RegisterVolume(req RegisterVolumeRequest) (resp *RegisterVolumeResult, err error) {
	resp = &RegisterVolumeResult{}
	err = c.client.Do("RegisterVolume", "POST", "/", req, resp)
	return
}

// SetLoadBasedAutoScaling specify the load-based auto scaling
// configuration for a specified layer. For more information, see Managing
// Load with Time-based and Load-based Instances Required Permissions : To
// use this action, an IAM user must have a Manage permissions level for
// the stack, or an attached policy that explicitly grants permissions. For
// more information on user permissions, see Managing User Permissions
func (c *OpsWorks) SetLoadBasedAutoScaling(req SetLoadBasedAutoScalingRequest) (err error) {
	// NRE
	err = c.client.Do("SetLoadBasedAutoScaling", "POST", "/", req, nil)
	return
}

// SetPermission specifies a user's permissions. For more information, see
// Security and Permissions Required Permissions : To use this action, an
// IAM user must have a Manage permissions level for the stack, or an
// attached policy that explicitly grants permissions. For more information
// on user permissions, see Managing User Permissions
func (c *OpsWorks) SetPermission(req SetPermissionRequest) (err error) {
	// NRE
	err = c.client.Do("SetPermission", "POST", "/", req, nil)
	return
}

// SetTimeBasedAutoScaling specify the time-based auto scaling
// configuration for a specified instance. For more information, see
// Managing Load with Time-based and Load-based Instances Required
// Permissions : To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly
// grants permissions. For more information on user permissions, see
// Managing User Permissions
func (c *OpsWorks) SetTimeBasedAutoScaling(req SetTimeBasedAutoScalingRequest) (err error) {
	// NRE
	err = c.client.Do("SetTimeBasedAutoScaling", "POST", "/", req, nil)
	return
}

// StartInstance starts a specified instance. For more information, see
// Starting, Stopping, and Rebooting Instances Required Permissions : To
// use this action, an IAM user must have a Manage permissions level for
// the stack, or an attached policy that explicitly grants permissions. For
// more information on user permissions, see Managing User Permissions
func (c *OpsWorks) StartInstance(req StartInstanceRequest) (err error) {
	// NRE
	err = c.client.Do("StartInstance", "POST", "/", req, nil)
	return
}

// StartStack starts a stack's instances. Required Permissions : To use
// this action, an IAM user must have a Manage permissions level for the
// stack, or an attached policy that explicitly grants permissions. For
// more information on user permissions, see Managing User Permissions
func (c *OpsWorks) StartStack(req StartStackRequest) (err error) {
	// NRE
	err = c.client.Do("StartStack", "POST", "/", req, nil)
	return
}

// StopInstance stops a specified instance. When you stop a standard
// instance, the data disappears and must be reinstalled when you restart
// the instance. You can stop an Amazon EBS-backed instance without losing
// data. For more information, see Starting, Stopping, and Rebooting
// Instances Required Permissions : To use this action, an IAM user must
// have a Manage permissions level for the stack, or an attached policy
// that explicitly grants permissions. For more information on user
// permissions, see Managing User Permissions
func (c *OpsWorks) StopInstance(req StopInstanceRequest) (err error) {
	// NRE
	err = c.client.Do("StopInstance", "POST", "/", req, nil)
	return
}

// StopStack stops a specified stack. Required Permissions : To use this
// action, an IAM user must have a Manage permissions level for the stack,
// or an attached policy that explicitly grants permissions. For more
// information on user permissions, see Managing User Permissions
func (c *OpsWorks) StopStack(req StopStackRequest) (err error) {
	// NRE
	err = c.client.Do("StopStack", "POST", "/", req, nil)
	return
}

// UnassignVolume unassigns an assigned Amazon EBS volume. The volume
// remains registered with the stack. For more information, see Resource
// Management Required Permissions : To use this action, an IAM user must
// have a Manage permissions level for the stack, or an attached policy
// that explicitly grants permissions. For more information on user
// permissions, see Managing User Permissions
func (c *OpsWorks) UnassignVolume(req UnassignVolumeRequest) (err error) {
	// NRE
	err = c.client.Do("UnassignVolume", "POST", "/", req, nil)
	return
}

// UpdateApp updates a specified app. Required Permissions : To use this
// action, an IAM user must have a Deploy or Manage permissions level for
// the stack, or an attached policy that explicitly grants permissions. For
// more information on user permissions, see Managing User Permissions
func (c *OpsWorks) UpdateApp(req UpdateAppRequest) (err error) {
	// NRE
	err = c.client.Do("UpdateApp", "POST", "/", req, nil)
	return
}

// UpdateElasticIP updates a registered Elastic IP address's name. For more
// information, see Resource Management Required Permissions : To use this
// action, an IAM user must have a Manage permissions level for the stack,
// or an attached policy that explicitly grants permissions. For more
// information on user permissions, see Managing User Permissions
func (c *OpsWorks) UpdateElasticIP(req UpdateElasticIPRequest) (err error) {
	// NRE
	err = c.client.Do("UpdateElasticIp", "POST", "/", req, nil)
	return
}

// UpdateInstance updates a specified instance. Required Permissions : To
// use this action, an IAM user must have a Manage permissions level for
// the stack, or an attached policy that explicitly grants permissions. For
// more information on user permissions, see Managing User Permissions
func (c *OpsWorks) UpdateInstance(req UpdateInstanceRequest) (err error) {
	// NRE
	err = c.client.Do("UpdateInstance", "POST", "/", req, nil)
	return
}

// UpdateLayer updates a specified layer. Required Permissions : To use
// this action, an IAM user must have a Manage permissions level for the
// stack, or an attached policy that explicitly grants permissions. For
// more information on user permissions, see Managing User Permissions
func (c *OpsWorks) UpdateLayer(req UpdateLayerRequest) (err error) {
	// NRE
	err = c.client.Do("UpdateLayer", "POST", "/", req, nil)
	return
}

// UpdateMyUserProfile updates a user's SSH public key. Required
// Permissions : To use this action, an IAM user must have self-management
// enabled or an attached policy that explicitly grants permissions. For
// more information on user permissions, see Managing User Permissions
func (c *OpsWorks) UpdateMyUserProfile(req UpdateMyUserProfileRequest) (err error) {
	// NRE
	err = c.client.Do("UpdateMyUserProfile", "POST", "/", req, nil)
	return
}

// UpdateRdsDbInstance is undocumented.
func (c *OpsWorks) UpdateRdsDbInstance(req UpdateRdsDbInstanceRequest) (err error) {
	// NRE
	err = c.client.Do("UpdateRdsDbInstance", "POST", "/", req, nil)
	return
}

// UpdateStack updates a specified stack. Required Permissions : To use
// this action, an IAM user must have a Manage permissions level for the
// stack, or an attached policy that explicitly grants permissions. For
// more information on user permissions, see Managing User Permissions
func (c *OpsWorks) UpdateStack(req UpdateStackRequest) (err error) {
	// NRE
	err = c.client.Do("UpdateStack", "POST", "/", req, nil)
	return
}

// UpdateUserProfile updates a specified user profile. Required Permissions
// : To use this action, an IAM user must have an attached policy that
// explicitly grants permissions. For more information on user permissions,
// see Managing User Permissions
func (c *OpsWorks) UpdateUserProfile(req UpdateUserProfileRequest) (err error) {
	// NRE
	err = c.client.Do("UpdateUserProfile", "POST", "/", req, nil)
	return
}

// UpdateVolume updates an Amazon EBS volume's name or mount point. For
// more information, see Resource Management Required Permissions : To use
// this action, an IAM user must have a Manage permissions level for the
// stack, or an attached policy that explicitly grants permissions. For
// more information on user permissions, see Managing User Permissions
func (c *OpsWorks) UpdateVolume(req UpdateVolumeRequest) (err error) {
	// NRE
	err = c.client.Do("UpdateVolume", "POST", "/", req, nil)
	return
}

type App struct {
	AppID            string            `json:"AppId,omitempty"`
	AppSource        Source            `json:"AppSource,omitempty"`
	Attributes       map[string]string `json:"Attributes,omitempty"`
	CreatedAt        string            `json:"CreatedAt,omitempty"`
	DataSources      []DataSource      `json:"DataSources,omitempty"`
	Description      string            `json:"Description,omitempty"`
	Domains          []string          `json:"Domains,omitempty"`
	EnableSsl        bool              `json:"EnableSsl,omitempty"`
	Name             string            `json:"Name,omitempty"`
	Shortname        string            `json:"Shortname,omitempty"`
	SslConfiguration SslConfiguration  `json:"SslConfiguration,omitempty"`
	StackID          string            `json:"StackId,omitempty"`
	Type             string            `json:"Type,omitempty"`
}

type AssignVolumeRequest struct {
	InstanceID string `json:"InstanceId,omitempty"`
	VolumeID   string `json:"VolumeId"`
}

type AssociateElasticIPRequest struct {
	ElasticIP  string `json:"ElasticIp"`
	InstanceID string `json:"InstanceId,omitempty"`
}

type AttachElasticLoadBalancerRequest struct {
	ElasticLoadBalancerName string `json:"ElasticLoadBalancerName"`
	LayerID                 string `json:"LayerId"`
}

type AutoScalingThresholds struct {
	CPUThreshold       float64 `json:"CpuThreshold,omitempty"`
	IgnoreMetricsTime  int     `json:"IgnoreMetricsTime,omitempty"`
	InstanceCount      int     `json:"InstanceCount,omitempty"`
	LoadThreshold      float64 `json:"LoadThreshold,omitempty"`
	MemoryThreshold    float64 `json:"MemoryThreshold,omitempty"`
	ThresholdsWaitTime int     `json:"ThresholdsWaitTime,omitempty"`
}

type ChefConfiguration struct {
	BerkshelfVersion string `json:"BerkshelfVersion,omitempty"`
	ManageBerkshelf  bool   `json:"ManageBerkshelf,omitempty"`
}

type CloneStackRequest struct {
	Attributes                map[string]string         `json:"Attributes,omitempty"`
	ChefConfiguration         ChefConfiguration         `json:"ChefConfiguration,omitempty"`
	CloneAppIds               []string                  `json:"CloneAppIds,omitempty"`
	ClonePermissions          bool                      `json:"ClonePermissions,omitempty"`
	ConfigurationManager      StackConfigurationManager `json:"ConfigurationManager,omitempty"`
	CustomCookbooksSource     Source                    `json:"CustomCookbooksSource,omitempty"`
	CustomJSON                string                    `json:"CustomJson,omitempty"`
	DefaultAvailabilityZone   string                    `json:"DefaultAvailabilityZone,omitempty"`
	DefaultInstanceProfileARN string                    `json:"DefaultInstanceProfileArn,omitempty"`
	DefaultOs                 string                    `json:"DefaultOs,omitempty"`
	DefaultRootDeviceType     string                    `json:"DefaultRootDeviceType,omitempty"`
	DefaultSSHKeyName         string                    `json:"DefaultSshKeyName,omitempty"`
	DefaultSubnetID           string                    `json:"DefaultSubnetId,omitempty"`
	HostnameTheme             string                    `json:"HostnameTheme,omitempty"`
	Name                      string                    `json:"Name,omitempty"`
	Region                    string                    `json:"Region,omitempty"`
	ServiceRoleARN            string                    `json:"ServiceRoleArn"`
	SourceStackID             string                    `json:"SourceStackId"`
	UseCustomCookbooks        bool                      `json:"UseCustomCookbooks,omitempty"`
	UseOpsworksSecurityGroups bool                      `json:"UseOpsworksSecurityGroups,omitempty"`
	VpcID                     string                    `json:"VpcId,omitempty"`
}

type CloneStackResult struct {
	StackID string `json:"StackId,omitempty"`
}

type Command struct {
	AcknowledgedAt string `json:"AcknowledgedAt,omitempty"`
	CommandID      string `json:"CommandId,omitempty"`
	CompletedAt    string `json:"CompletedAt,omitempty"`
	CreatedAt      string `json:"CreatedAt,omitempty"`
	DeploymentID   string `json:"DeploymentId,omitempty"`
	ExitCode       int    `json:"ExitCode,omitempty"`
	InstanceID     string `json:"InstanceId,omitempty"`
	LogURL         string `json:"LogUrl,omitempty"`
	Status         string `json:"Status,omitempty"`
	Type           string `json:"Type,omitempty"`
}

type CreateAppRequest struct {
	AppSource        Source            `json:"AppSource,omitempty"`
	Attributes       map[string]string `json:"Attributes,omitempty"`
	DataSources      []DataSource      `json:"DataSources,omitempty"`
	Description      string            `json:"Description,omitempty"`
	Domains          []string          `json:"Domains,omitempty"`
	EnableSsl        bool              `json:"EnableSsl,omitempty"`
	Name             string            `json:"Name"`
	Shortname        string            `json:"Shortname,omitempty"`
	SslConfiguration SslConfiguration  `json:"SslConfiguration,omitempty"`
	StackID          string            `json:"StackId"`
	Type             string            `json:"Type"`
}

type CreateAppResult struct {
	AppID string `json:"AppId,omitempty"`
}

type CreateDeploymentRequest struct {
	AppID       string            `json:"AppId,omitempty"`
	Command     DeploymentCommand `json:"Command"`
	Comment     string            `json:"Comment,omitempty"`
	CustomJSON  string            `json:"CustomJson,omitempty"`
	InstanceIds []string          `json:"InstanceIds,omitempty"`
	StackID     string            `json:"StackId"`
}

type CreateDeploymentResult struct {
	DeploymentID string `json:"DeploymentId,omitempty"`
}

type CreateInstanceRequest struct {
	AmiID                string   `json:"AmiId,omitempty"`
	Architecture         string   `json:"Architecture,omitempty"`
	AutoScalingType      string   `json:"AutoScalingType,omitempty"`
	AvailabilityZone     string   `json:"AvailabilityZone,omitempty"`
	EbsOptimized         bool     `json:"EbsOptimized,omitempty"`
	Hostname             string   `json:"Hostname,omitempty"`
	InstallUpdatesOnBoot bool     `json:"InstallUpdatesOnBoot,omitempty"`
	InstanceType         string   `json:"InstanceType"`
	LayerIds             []string `json:"LayerIds"`
	Os                   string   `json:"Os,omitempty"`
	RootDeviceType       string   `json:"RootDeviceType,omitempty"`
	SSHKeyName           string   `json:"SshKeyName,omitempty"`
	StackID              string   `json:"StackId"`
	SubnetID             string   `json:"SubnetId,omitempty"`
	VirtualizationType   string   `json:"VirtualizationType,omitempty"`
}

type CreateInstanceResult struct {
	InstanceID string `json:"InstanceId,omitempty"`
}

type CreateLayerRequest struct {
	Attributes               map[string]string     `json:"Attributes,omitempty"`
	AutoAssignElasticIPs     bool                  `json:"AutoAssignElasticIps,omitempty"`
	AutoAssignPublicIPs      bool                  `json:"AutoAssignPublicIps,omitempty"`
	CustomInstanceProfileARN string                `json:"CustomInstanceProfileArn,omitempty"`
	CustomRecipes            Recipes               `json:"CustomRecipes,omitempty"`
	CustomSecurityGroupIds   []string              `json:"CustomSecurityGroupIds,omitempty"`
	EnableAutoHealing        bool                  `json:"EnableAutoHealing,omitempty"`
	InstallUpdatesOnBoot     bool                  `json:"InstallUpdatesOnBoot,omitempty"`
	Name                     string                `json:"Name"`
	Packages                 []string              `json:"Packages,omitempty"`
	Shortname                string                `json:"Shortname"`
	StackID                  string                `json:"StackId"`
	Type                     string                `json:"Type"`
	UseEbsOptimizedInstances bool                  `json:"UseEbsOptimizedInstances,omitempty"`
	VolumeConfigurations     []VolumeConfiguration `json:"VolumeConfigurations,omitempty"`
}

type CreateLayerResult struct {
	LayerID string `json:"LayerId,omitempty"`
}

type CreateStackRequest struct {
	Attributes                map[string]string         `json:"Attributes,omitempty"`
	ChefConfiguration         ChefConfiguration         `json:"ChefConfiguration,omitempty"`
	ConfigurationManager      StackConfigurationManager `json:"ConfigurationManager,omitempty"`
	CustomCookbooksSource     Source                    `json:"CustomCookbooksSource,omitempty"`
	CustomJSON                string                    `json:"CustomJson,omitempty"`
	DefaultAvailabilityZone   string                    `json:"DefaultAvailabilityZone,omitempty"`
	DefaultInstanceProfileARN string                    `json:"DefaultInstanceProfileArn"`
	DefaultOs                 string                    `json:"DefaultOs,omitempty"`
	DefaultRootDeviceType     string                    `json:"DefaultRootDeviceType,omitempty"`
	DefaultSSHKeyName         string                    `json:"DefaultSshKeyName,omitempty"`
	DefaultSubnetID           string                    `json:"DefaultSubnetId,omitempty"`
	HostnameTheme             string                    `json:"HostnameTheme,omitempty"`
	Name                      string                    `json:"Name"`
	Region                    string                    `json:"Region"`
	ServiceRoleARN            string                    `json:"ServiceRoleArn"`
	UseCustomCookbooks        bool                      `json:"UseCustomCookbooks,omitempty"`
	UseOpsworksSecurityGroups bool                      `json:"UseOpsworksSecurityGroups,omitempty"`
	VpcID                     string                    `json:"VpcId,omitempty"`
}

type CreateStackResult struct {
	StackID string `json:"StackId,omitempty"`
}

type CreateUserProfileRequest struct {
	AllowSelfManagement bool   `json:"AllowSelfManagement,omitempty"`
	IamUserARN          string `json:"IamUserArn"`
	SSHPublicKey        string `json:"SshPublicKey,omitempty"`
	SSHUsername         string `json:"SshUsername,omitempty"`
}

type CreateUserProfileResult struct {
	IamUserARN string `json:"IamUserArn,omitempty"`
}

type DataSource struct {
	ARN          string `json:"Arn,omitempty"`
	DatabaseName string `json:"DatabaseName,omitempty"`
	Type         string `json:"Type,omitempty"`
}

type DeleteAppRequest struct {
	AppID string `json:"AppId"`
}

type DeleteInstanceRequest struct {
	DeleteElasticIP bool   `json:"DeleteElasticIp,omitempty"`
	DeleteVolumes   bool   `json:"DeleteVolumes,omitempty"`
	InstanceID      string `json:"InstanceId"`
}

type DeleteLayerRequest struct {
	LayerID string `json:"LayerId"`
}

type DeleteStackRequest struct {
	StackID string `json:"StackId"`
}

type DeleteUserProfileRequest struct {
	IamUserARN string `json:"IamUserArn"`
}

type Deployment struct {
	AppID        string            `json:"AppId,omitempty"`
	Command      DeploymentCommand `json:"Command,omitempty"`
	Comment      string            `json:"Comment,omitempty"`
	CompletedAt  string            `json:"CompletedAt,omitempty"`
	CreatedAt    string            `json:"CreatedAt,omitempty"`
	CustomJSON   string            `json:"CustomJson,omitempty"`
	DeploymentID string            `json:"DeploymentId,omitempty"`
	Duration     int               `json:"Duration,omitempty"`
	IamUserARN   string            `json:"IamUserArn,omitempty"`
	InstanceIds  []string          `json:"InstanceIds,omitempty"`
	StackID      string            `json:"StackId,omitempty"`
	Status       string            `json:"Status,omitempty"`
}

type DeploymentCommand struct {
	Args map[string][]string `json:"Args,omitempty"`
	Name string              `json:"Name"`
}

type DeregisterElasticIPRequest struct {
	ElasticIP string `json:"ElasticIp"`
}

type DeregisterRdsDbInstanceRequest struct {
	RdsDbInstanceARN string `json:"RdsDbInstanceArn"`
}

type DeregisterVolumeRequest struct {
	VolumeID string `json:"VolumeId"`
}

type DescribeAppsRequest struct {
	AppIds  []string `json:"AppIds,omitempty"`
	StackID string   `json:"StackId,omitempty"`
}

type DescribeAppsResult struct {
	Apps []App `json:"Apps,omitempty"`
}

type DescribeCommandsRequest struct {
	CommandIds   []string `json:"CommandIds,omitempty"`
	DeploymentID string   `json:"DeploymentId,omitempty"`
	InstanceID   string   `json:"InstanceId,omitempty"`
}

type DescribeCommandsResult struct {
	Commands []Command `json:"Commands,omitempty"`
}

type DescribeDeploymentsRequest struct {
	AppID         string   `json:"AppId,omitempty"`
	DeploymentIds []string `json:"DeploymentIds,omitempty"`
	StackID       string   `json:"StackId,omitempty"`
}

type DescribeDeploymentsResult struct {
	Deployments []Deployment `json:"Deployments,omitempty"`
}

type DescribeElasticIPsRequest struct {
	InstanceID string   `json:"InstanceId,omitempty"`
	IPs        []string `json:"Ips,omitempty"`
	StackID    string   `json:"StackId,omitempty"`
}

type DescribeElasticIPsResult struct {
	ElasticIPs []ElasticIP `json:"ElasticIps,omitempty"`
}

type DescribeElasticLoadBalancersRequest struct {
	LayerIds []string `json:"LayerIds,omitempty"`
	StackID  string   `json:"StackId,omitempty"`
}

type DescribeElasticLoadBalancersResult struct {
	ElasticLoadBalancers []ElasticLoadBalancer `json:"ElasticLoadBalancers,omitempty"`
}

type DescribeInstancesRequest struct {
	InstanceIds []string `json:"InstanceIds,omitempty"`
	LayerID     string   `json:"LayerId,omitempty"`
	StackID     string   `json:"StackId,omitempty"`
}

type DescribeInstancesResult struct {
	Instances []Instance `json:"Instances,omitempty"`
}

type DescribeLayersRequest struct {
	LayerIds []string `json:"LayerIds,omitempty"`
	StackID  string   `json:"StackId,omitempty"`
}

type DescribeLayersResult struct {
	Layers []Layer `json:"Layers,omitempty"`
}

type DescribeLoadBasedAutoScalingRequest struct {
	LayerIds []string `json:"LayerIds"`
}

type DescribeLoadBasedAutoScalingResult struct {
	LoadBasedAutoScalingConfigurations []LoadBasedAutoScalingConfiguration `json:"LoadBasedAutoScalingConfigurations,omitempty"`
}

type DescribeMyUserProfileResult struct {
	UserProfile SelfUserProfile `json:"UserProfile,omitempty"`
}

type DescribePermissionsRequest struct {
	IamUserARN string `json:"IamUserArn,omitempty"`
	StackID    string `json:"StackId,omitempty"`
}

type DescribePermissionsResult struct {
	Permissions []Permission `json:"Permissions,omitempty"`
}

type DescribeRaidArraysRequest struct {
	InstanceID   string   `json:"InstanceId,omitempty"`
	RaidArrayIds []string `json:"RaidArrayIds,omitempty"`
}

type DescribeRaidArraysResult struct {
	RaidArrays []RaidArray `json:"RaidArrays,omitempty"`
}

type DescribeRdsDbInstancesRequest struct {
	RdsDbInstanceARNs []string `json:"RdsDbInstanceArns,omitempty"`
	StackID           string   `json:"StackId"`
}

type DescribeRdsDbInstancesResult struct {
	RdsDbInstances []RdsDbInstance `json:"RdsDbInstances,omitempty"`
}

type DescribeServiceErrorsRequest struct {
	InstanceID      string   `json:"InstanceId,omitempty"`
	ServiceErrorIds []string `json:"ServiceErrorIds,omitempty"`
	StackID         string   `json:"StackId,omitempty"`
}

type DescribeServiceErrorsResult struct {
	ServiceErrors []ServiceError `json:"ServiceErrors,omitempty"`
}

type DescribeStackSummaryRequest struct {
	StackID string `json:"StackId"`
}

type DescribeStackSummaryResult struct {
	StackSummary StackSummary `json:"StackSummary,omitempty"`
}

type DescribeStacksRequest struct {
	StackIds []string `json:"StackIds,omitempty"`
}

type DescribeStacksResult struct {
	Stacks []Stack `json:"Stacks,omitempty"`
}

type DescribeTimeBasedAutoScalingRequest struct {
	InstanceIds []string `json:"InstanceIds"`
}

type DescribeTimeBasedAutoScalingResult struct {
	TimeBasedAutoScalingConfigurations []TimeBasedAutoScalingConfiguration `json:"TimeBasedAutoScalingConfigurations,omitempty"`
}

type DescribeUserProfilesRequest struct {
	IamUserARNs []string `json:"IamUserArns,omitempty"`
}

type DescribeUserProfilesResult struct {
	UserProfiles []UserProfile `json:"UserProfiles,omitempty"`
}

type DescribeVolumesRequest struct {
	InstanceID  string   `json:"InstanceId,omitempty"`
	RaidArrayID string   `json:"RaidArrayId,omitempty"`
	StackID     string   `json:"StackId,omitempty"`
	VolumeIds   []string `json:"VolumeIds,omitempty"`
}

type DescribeVolumesResult struct {
	Volumes []Volume `json:"Volumes,omitempty"`
}

type DetachElasticLoadBalancerRequest struct {
	ElasticLoadBalancerName string `json:"ElasticLoadBalancerName"`
	LayerID                 string `json:"LayerId"`
}

type DisassociateElasticIPRequest struct {
	ElasticIP string `json:"ElasticIp"`
}

type ElasticIP struct {
	Domain     string `json:"Domain,omitempty"`
	InstanceID string `json:"InstanceId,omitempty"`
	IP         string `json:"Ip,omitempty"`
	Name       string `json:"Name,omitempty"`
	Region     string `json:"Region,omitempty"`
}

type ElasticLoadBalancer struct {
	AvailabilityZones       []string `json:"AvailabilityZones,omitempty"`
	DNSName                 string   `json:"DnsName,omitempty"`
	Ec2InstanceIds          []string `json:"Ec2InstanceIds,omitempty"`
	ElasticLoadBalancerName string   `json:"ElasticLoadBalancerName,omitempty"`
	LayerID                 string   `json:"LayerId,omitempty"`
	Region                  string   `json:"Region,omitempty"`
	StackID                 string   `json:"StackId,omitempty"`
	SubnetIds               []string `json:"SubnetIds,omitempty"`
	VpcID                   string   `json:"VpcId,omitempty"`
}

type GetHostnameSuggestionRequest struct {
	LayerID string `json:"LayerId"`
}

type GetHostnameSuggestionResult struct {
	Hostname string `json:"Hostname,omitempty"`
	LayerID  string `json:"LayerId,omitempty"`
}

type Instance struct {
	AmiID                    string   `json:"AmiId,omitempty"`
	Architecture             string   `json:"Architecture,omitempty"`
	AutoScalingType          string   `json:"AutoScalingType,omitempty"`
	AvailabilityZone         string   `json:"AvailabilityZone,omitempty"`
	CreatedAt                string   `json:"CreatedAt,omitempty"`
	EbsOptimized             bool     `json:"EbsOptimized,omitempty"`
	Ec2InstanceID            string   `json:"Ec2InstanceId,omitempty"`
	ElasticIP                string   `json:"ElasticIp,omitempty"`
	Hostname                 string   `json:"Hostname,omitempty"`
	InstallUpdatesOnBoot     bool     `json:"InstallUpdatesOnBoot,omitempty"`
	InstanceID               string   `json:"InstanceId,omitempty"`
	InstanceProfileARN       string   `json:"InstanceProfileArn,omitempty"`
	InstanceType             string   `json:"InstanceType,omitempty"`
	LastServiceErrorID       string   `json:"LastServiceErrorId,omitempty"`
	LayerIds                 []string `json:"LayerIds,omitempty"`
	Os                       string   `json:"Os,omitempty"`
	PrivateDNS               string   `json:"PrivateDns,omitempty"`
	PrivateIP                string   `json:"PrivateIp,omitempty"`
	PublicDNS                string   `json:"PublicDns,omitempty"`
	PublicIP                 string   `json:"PublicIp,omitempty"`
	RootDeviceType           string   `json:"RootDeviceType,omitempty"`
	RootDeviceVolumeID       string   `json:"RootDeviceVolumeId,omitempty"`
	SecurityGroupIds         []string `json:"SecurityGroupIds,omitempty"`
	SSHHostDsaKeyFingerprint string   `json:"SshHostDsaKeyFingerprint,omitempty"`
	SSHHostRsaKeyFingerprint string   `json:"SshHostRsaKeyFingerprint,omitempty"`
	SSHKeyName               string   `json:"SshKeyName,omitempty"`
	StackID                  string   `json:"StackId,omitempty"`
	Status                   string   `json:"Status,omitempty"`
	SubnetID                 string   `json:"SubnetId,omitempty"`
	VirtualizationType       string   `json:"VirtualizationType,omitempty"`
}

type InstancesCount struct {
	Booting        int `json:"Booting,omitempty"`
	ConnectionLost int `json:"ConnectionLost,omitempty"`
	Online         int `json:"Online,omitempty"`
	Pending        int `json:"Pending,omitempty"`
	Rebooting      int `json:"Rebooting,omitempty"`
	Requested      int `json:"Requested,omitempty"`
	RunningSetup   int `json:"RunningSetup,omitempty"`
	SetupFailed    int `json:"SetupFailed,omitempty"`
	ShuttingDown   int `json:"ShuttingDown,omitempty"`
	StartFailed    int `json:"StartFailed,omitempty"`
	Stopped        int `json:"Stopped,omitempty"`
	Stopping       int `json:"Stopping,omitempty"`
	Terminated     int `json:"Terminated,omitempty"`
	Terminating    int `json:"Terminating,omitempty"`
}

type Layer struct {
	Attributes                map[string]string     `json:"Attributes,omitempty"`
	AutoAssignElasticIPs      bool                  `json:"AutoAssignElasticIps,omitempty"`
	AutoAssignPublicIPs       bool                  `json:"AutoAssignPublicIps,omitempty"`
	CreatedAt                 string                `json:"CreatedAt,omitempty"`
	CustomInstanceProfileARN  string                `json:"CustomInstanceProfileArn,omitempty"`
	CustomRecipes             Recipes               `json:"CustomRecipes,omitempty"`
	CustomSecurityGroupIds    []string              `json:"CustomSecurityGroupIds,omitempty"`
	DefaultRecipes            Recipes               `json:"DefaultRecipes,omitempty"`
	DefaultSecurityGroupNames []string              `json:"DefaultSecurityGroupNames,omitempty"`
	EnableAutoHealing         bool                  `json:"EnableAutoHealing,omitempty"`
	InstallUpdatesOnBoot      bool                  `json:"InstallUpdatesOnBoot,omitempty"`
	LayerID                   string                `json:"LayerId,omitempty"`
	Name                      string                `json:"Name,omitempty"`
	Packages                  []string              `json:"Packages,omitempty"`
	Shortname                 string                `json:"Shortname,omitempty"`
	StackID                   string                `json:"StackId,omitempty"`
	Type                      string                `json:"Type,omitempty"`
	UseEbsOptimizedInstances  bool                  `json:"UseEbsOptimizedInstances,omitempty"`
	VolumeConfigurations      []VolumeConfiguration `json:"VolumeConfigurations,omitempty"`
}

type LoadBasedAutoScalingConfiguration struct {
	DownScaling AutoScalingThresholds `json:"DownScaling,omitempty"`
	Enable      bool                  `json:"Enable,omitempty"`
	LayerID     string                `json:"LayerId,omitempty"`
	UpScaling   AutoScalingThresholds `json:"UpScaling,omitempty"`
}

type Permission struct {
	AllowSSH   bool   `json:"AllowSsh,omitempty"`
	AllowSudo  bool   `json:"AllowSudo,omitempty"`
	IamUserARN string `json:"IamUserArn,omitempty"`
	Level      string `json:"Level,omitempty"`
	StackID    string `json:"StackId,omitempty"`
}

type RaidArray struct {
	AvailabilityZone string `json:"AvailabilityZone,omitempty"`
	CreatedAt        string `json:"CreatedAt,omitempty"`
	Device           string `json:"Device,omitempty"`
	InstanceID       string `json:"InstanceId,omitempty"`
	Iops             int    `json:"Iops,omitempty"`
	MountPoint       string `json:"MountPoint,omitempty"`
	Name             string `json:"Name,omitempty"`
	NumberOfDisks    int    `json:"NumberOfDisks,omitempty"`
	RaidArrayID      string `json:"RaidArrayId,omitempty"`
	RaidLevel        int    `json:"RaidLevel,omitempty"`
	Size             int    `json:"Size,omitempty"`
	VolumeType       string `json:"VolumeType,omitempty"`
}

type RdsDbInstance struct {
	Address              string `json:"Address,omitempty"`
	DbInstanceIdentifier string `json:"DbInstanceIdentifier,omitempty"`
	DbPassword           string `json:"DbPassword,omitempty"`
	DbUser               string `json:"DbUser,omitempty"`
	Engine               string `json:"Engine,omitempty"`
	MissingOnRds         bool   `json:"MissingOnRds,omitempty"`
	RdsDbInstanceARN     string `json:"RdsDbInstanceArn,omitempty"`
	Region               string `json:"Region,omitempty"`
	StackID              string `json:"StackId,omitempty"`
}

type RebootInstanceRequest struct {
	InstanceID string `json:"InstanceId"`
}

type Recipes struct {
	Configure []string `json:"Configure,omitempty"`
	Deploy    []string `json:"Deploy,omitempty"`
	Setup     []string `json:"Setup,omitempty"`
	Shutdown  []string `json:"Shutdown,omitempty"`
	Undeploy  []string `json:"Undeploy,omitempty"`
}

type RegisterElasticIPRequest struct {
	ElasticIP string `json:"ElasticIp"`
	StackID   string `json:"StackId"`
}

type RegisterElasticIPResult struct {
	ElasticIP string `json:"ElasticIp,omitempty"`
}

type RegisterRdsDbInstanceRequest struct {
	DbPassword       string `json:"DbPassword"`
	DbUser           string `json:"DbUser"`
	RdsDbInstanceARN string `json:"RdsDbInstanceArn"`
	StackID          string `json:"StackId"`
}

type RegisterVolumeRequest struct {
	Ec2VolumeID string `json:"Ec2VolumeId,omitempty"`
	StackID     string `json:"StackId"`
}

type RegisterVolumeResult struct {
	VolumeID string `json:"VolumeId,omitempty"`
}

type SelfUserProfile struct {
	IamUserARN   string `json:"IamUserArn,omitempty"`
	Name         string `json:"Name,omitempty"`
	SSHPublicKey string `json:"SshPublicKey,omitempty"`
	SSHUsername  string `json:"SshUsername,omitempty"`
}

type ServiceError struct {
	CreatedAt      string `json:"CreatedAt,omitempty"`
	InstanceID     string `json:"InstanceId,omitempty"`
	Message        string `json:"Message,omitempty"`
	ServiceErrorID string `json:"ServiceErrorId,omitempty"`
	StackID        string `json:"StackId,omitempty"`
	Type           string `json:"Type,omitempty"`
}

type SetLoadBasedAutoScalingRequest struct {
	DownScaling AutoScalingThresholds `json:"DownScaling,omitempty"`
	Enable      bool                  `json:"Enable,omitempty"`
	LayerID     string                `json:"LayerId"`
	UpScaling   AutoScalingThresholds `json:"UpScaling,omitempty"`
}

type SetPermissionRequest struct {
	AllowSSH   bool   `json:"AllowSsh,omitempty"`
	AllowSudo  bool   `json:"AllowSudo,omitempty"`
	IamUserARN string `json:"IamUserArn"`
	Level      string `json:"Level,omitempty"`
	StackID    string `json:"StackId"`
}

type SetTimeBasedAutoScalingRequest struct {
	AutoScalingSchedule WeeklyAutoScalingSchedule `json:"AutoScalingSchedule,omitempty"`
	InstanceID          string                    `json:"InstanceId"`
}

type Source struct {
	Password string `json:"Password,omitempty"`
	Revision string `json:"Revision,omitempty"`
	SSHKey   string `json:"SshKey,omitempty"`
	Type     string `json:"Type,omitempty"`
	URL      string `json:"Url,omitempty"`
	Username string `json:"Username,omitempty"`
}

type SslConfiguration struct {
	Certificate string `json:"Certificate"`
	Chain       string `json:"Chain,omitempty"`
	PrivateKey  string `json:"PrivateKey"`
}

type Stack struct {
	ARN                       string                    `json:"Arn,omitempty"`
	Attributes                map[string]string         `json:"Attributes,omitempty"`
	ChefConfiguration         ChefConfiguration         `json:"ChefConfiguration,omitempty"`
	ConfigurationManager      StackConfigurationManager `json:"ConfigurationManager,omitempty"`
	CreatedAt                 string                    `json:"CreatedAt,omitempty"`
	CustomCookbooksSource     Source                    `json:"CustomCookbooksSource,omitempty"`
	CustomJSON                string                    `json:"CustomJson,omitempty"`
	DefaultAvailabilityZone   string                    `json:"DefaultAvailabilityZone,omitempty"`
	DefaultInstanceProfileARN string                    `json:"DefaultInstanceProfileArn,omitempty"`
	DefaultOs                 string                    `json:"DefaultOs,omitempty"`
	DefaultRootDeviceType     string                    `json:"DefaultRootDeviceType,omitempty"`
	DefaultSSHKeyName         string                    `json:"DefaultSshKeyName,omitempty"`
	DefaultSubnetID           string                    `json:"DefaultSubnetId,omitempty"`
	HostnameTheme             string                    `json:"HostnameTheme,omitempty"`
	Name                      string                    `json:"Name,omitempty"`
	Region                    string                    `json:"Region,omitempty"`
	ServiceRoleARN            string                    `json:"ServiceRoleArn,omitempty"`
	StackID                   string                    `json:"StackId,omitempty"`
	UseCustomCookbooks        bool                      `json:"UseCustomCookbooks,omitempty"`
	UseOpsworksSecurityGroups bool                      `json:"UseOpsworksSecurityGroups,omitempty"`
	VpcID                     string                    `json:"VpcId,omitempty"`
}

type StackConfigurationManager struct {
	Name    string `json:"Name,omitempty"`
	Version string `json:"Version,omitempty"`
}

type StackSummary struct {
	AppsCount      int            `json:"AppsCount,omitempty"`
	ARN            string         `json:"Arn,omitempty"`
	InstancesCount InstancesCount `json:"InstancesCount,omitempty"`
	LayersCount    int            `json:"LayersCount,omitempty"`
	Name           string         `json:"Name,omitempty"`
	StackID        string         `json:"StackId,omitempty"`
}

type StartInstanceRequest struct {
	InstanceID string `json:"InstanceId"`
}

type StartStackRequest struct {
	StackID string `json:"StackId"`
}

type StopInstanceRequest struct {
	InstanceID string `json:"InstanceId"`
}

type StopStackRequest struct {
	StackID string `json:"StackId"`
}

type TimeBasedAutoScalingConfiguration struct {
	AutoScalingSchedule WeeklyAutoScalingSchedule `json:"AutoScalingSchedule,omitempty"`
	InstanceID          string                    `json:"InstanceId,omitempty"`
}

type UnassignVolumeRequest struct {
	VolumeID string `json:"VolumeId"`
}

type UpdateAppRequest struct {
	AppID            string            `json:"AppId"`
	AppSource        Source            `json:"AppSource,omitempty"`
	Attributes       map[string]string `json:"Attributes,omitempty"`
	DataSources      []DataSource      `json:"DataSources,omitempty"`
	Description      string            `json:"Description,omitempty"`
	Domains          []string          `json:"Domains,omitempty"`
	EnableSsl        bool              `json:"EnableSsl,omitempty"`
	Name             string            `json:"Name,omitempty"`
	SslConfiguration SslConfiguration  `json:"SslConfiguration,omitempty"`
	Type             string            `json:"Type,omitempty"`
}

type UpdateElasticIPRequest struct {
	ElasticIP string `json:"ElasticIp"`
	Name      string `json:"Name,omitempty"`
}

type UpdateInstanceRequest struct {
	AmiID                string   `json:"AmiId,omitempty"`
	Architecture         string   `json:"Architecture,omitempty"`
	AutoScalingType      string   `json:"AutoScalingType,omitempty"`
	EbsOptimized         bool     `json:"EbsOptimized,omitempty"`
	Hostname             string   `json:"Hostname,omitempty"`
	InstallUpdatesOnBoot bool     `json:"InstallUpdatesOnBoot,omitempty"`
	InstanceID           string   `json:"InstanceId"`
	InstanceType         string   `json:"InstanceType,omitempty"`
	LayerIds             []string `json:"LayerIds,omitempty"`
	Os                   string   `json:"Os,omitempty"`
	SSHKeyName           string   `json:"SshKeyName,omitempty"`
}

type UpdateLayerRequest struct {
	Attributes               map[string]string     `json:"Attributes,omitempty"`
	AutoAssignElasticIPs     bool                  `json:"AutoAssignElasticIps,omitempty"`
	AutoAssignPublicIPs      bool                  `json:"AutoAssignPublicIps,omitempty"`
	CustomInstanceProfileARN string                `json:"CustomInstanceProfileArn,omitempty"`
	CustomRecipes            Recipes               `json:"CustomRecipes,omitempty"`
	CustomSecurityGroupIds   []string              `json:"CustomSecurityGroupIds,omitempty"`
	EnableAutoHealing        bool                  `json:"EnableAutoHealing,omitempty"`
	InstallUpdatesOnBoot     bool                  `json:"InstallUpdatesOnBoot,omitempty"`
	LayerID                  string                `json:"LayerId"`
	Name                     string                `json:"Name,omitempty"`
	Packages                 []string              `json:"Packages,omitempty"`
	Shortname                string                `json:"Shortname,omitempty"`
	UseEbsOptimizedInstances bool                  `json:"UseEbsOptimizedInstances,omitempty"`
	VolumeConfigurations     []VolumeConfiguration `json:"VolumeConfigurations,omitempty"`
}

type UpdateMyUserProfileRequest struct {
	SSHPublicKey string `json:"SshPublicKey,omitempty"`
}

type UpdateRdsDbInstanceRequest struct {
	DbPassword       string `json:"DbPassword,omitempty"`
	DbUser           string `json:"DbUser,omitempty"`
	RdsDbInstanceARN string `json:"RdsDbInstanceArn"`
}

type UpdateStackRequest struct {
	Attributes                map[string]string         `json:"Attributes,omitempty"`
	ChefConfiguration         ChefConfiguration         `json:"ChefConfiguration,omitempty"`
	ConfigurationManager      StackConfigurationManager `json:"ConfigurationManager,omitempty"`
	CustomCookbooksSource     Source                    `json:"CustomCookbooksSource,omitempty"`
	CustomJSON                string                    `json:"CustomJson,omitempty"`
	DefaultAvailabilityZone   string                    `json:"DefaultAvailabilityZone,omitempty"`
	DefaultInstanceProfileARN string                    `json:"DefaultInstanceProfileArn,omitempty"`
	DefaultOs                 string                    `json:"DefaultOs,omitempty"`
	DefaultRootDeviceType     string                    `json:"DefaultRootDeviceType,omitempty"`
	DefaultSSHKeyName         string                    `json:"DefaultSshKeyName,omitempty"`
	DefaultSubnetID           string                    `json:"DefaultSubnetId,omitempty"`
	HostnameTheme             string                    `json:"HostnameTheme,omitempty"`
	Name                      string                    `json:"Name,omitempty"`
	ServiceRoleARN            string                    `json:"ServiceRoleArn,omitempty"`
	StackID                   string                    `json:"StackId"`
	UseCustomCookbooks        bool                      `json:"UseCustomCookbooks,omitempty"`
	UseOpsworksSecurityGroups bool                      `json:"UseOpsworksSecurityGroups,omitempty"`
}

type UpdateUserProfileRequest struct {
	AllowSelfManagement bool   `json:"AllowSelfManagement,omitempty"`
	IamUserARN          string `json:"IamUserArn"`
	SSHPublicKey        string `json:"SshPublicKey,omitempty"`
	SSHUsername         string `json:"SshUsername,omitempty"`
}

type UpdateVolumeRequest struct {
	MountPoint string `json:"MountPoint,omitempty"`
	Name       string `json:"Name,omitempty"`
	VolumeID   string `json:"VolumeId"`
}

type UserProfile struct {
	AllowSelfManagement bool   `json:"AllowSelfManagement,omitempty"`
	IamUserARN          string `json:"IamUserArn,omitempty"`
	Name                string `json:"Name,omitempty"`
	SSHPublicKey        string `json:"SshPublicKey,omitempty"`
	SSHUsername         string `json:"SshUsername,omitempty"`
}

type Volume struct {
	AvailabilityZone string `json:"AvailabilityZone,omitempty"`
	Device           string `json:"Device,omitempty"`
	Ec2VolumeID      string `json:"Ec2VolumeId,omitempty"`
	InstanceID       string `json:"InstanceId,omitempty"`
	Iops             int    `json:"Iops,omitempty"`
	MountPoint       string `json:"MountPoint,omitempty"`
	Name             string `json:"Name,omitempty"`
	RaidArrayID      string `json:"RaidArrayId,omitempty"`
	Region           string `json:"Region,omitempty"`
	Size             int    `json:"Size,omitempty"`
	Status           string `json:"Status,omitempty"`
	VolumeID         string `json:"VolumeId,omitempty"`
	VolumeType       string `json:"VolumeType,omitempty"`
}

type VolumeConfiguration struct {
	Iops          int    `json:"Iops,omitempty"`
	MountPoint    string `json:"MountPoint"`
	NumberOfDisks int    `json:"NumberOfDisks"`
	RaidLevel     int    `json:"RaidLevel,omitempty"`
	Size          int    `json:"Size"`
	VolumeType    string `json:"VolumeType,omitempty"`
}

type WeeklyAutoScalingSchedule struct {
	Friday    map[string]string `json:"Friday,omitempty"`
	Monday    map[string]string `json:"Monday,omitempty"`
	Saturday  map[string]string `json:"Saturday,omitempty"`
	Sunday    map[string]string `json:"Sunday,omitempty"`
	Thursday  map[string]string `json:"Thursday,omitempty"`
	Tuesday   map[string]string `json:"Tuesday,omitempty"`
	Wednesday map[string]string `json:"Wednesday,omitempty"`
}

var _ time.Time // to avoid errors if the time package isn't referenced
