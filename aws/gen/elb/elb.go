// Package elb provides a client for Elastic Load Balancing.
package elb

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// ELB is a client for Elastic Load Balancing.
type ELB struct {
	client *aws.QueryClient
}

// New returns a new ELB client.
func New(key, secret, region string, client *http.Client) *ELB {
	if client == nil {
		client = http.DefaultClient
	}

	service := "elasticloadbalancing"
	endpoint, service, region := endpoints.Lookup("elasticloadbalancing", region)

	return &ELB{
		client: &aws.QueryClient{
			Signer: &aws.V4Signer{
				Key:     key,
				Secret:  secret,
				Service: service,
				Region:  region,
				IncludeXAmzContentSha256: true,
			},
			Client:     client,
			Endpoint:   endpoint,
			APIVersion: "2012-06-01",
		},
	}
}

// AddTags adds one or more tags for the specified load balancer. Each load
// balancer can have a maximum of 10 tags. Each tag consists of a key and
// an optional value. Tag keys must be unique for each load balancer. If a
// tag with the same key is already associated with the load balancer, this
// action will update the value of the key. For more information, see
// Tagging in the Elastic Load Balancing Developer Guide
func (c *ELB) AddTags(req AddTagsInput) (resp *AddTagsResult, err error) {
	resp = &AddTagsResult{}
	err = c.client.Do("AddTags", "POST", "/", req, resp)
	return
}

// ApplySecurityGroupsToLoadBalancer associates one or more security groups
// with your load balancer in Amazon Virtual Private Cloud (Amazon The
// provided security group IDs will override any currently applied security
// groups. For more information, see Manage Security Groups in Amazon in
// the Elastic Load Balancing Developer Guide
func (c *ELB) ApplySecurityGroupsToLoadBalancer(req ApplySecurityGroupsToLoadBalancerInput) (resp *ApplySecurityGroupsToLoadBalancerResult, err error) {
	resp = &ApplySecurityGroupsToLoadBalancerResult{}
	err = c.client.Do("ApplySecurityGroupsToLoadBalancer", "POST", "/", req, resp)
	return
}

// AttachLoadBalancerToSubnets adds one or more subnets to the set of
// configured subnets in the Amazon Virtual Private Cloud (Amazon for the
// load balancer. The load balancers evenly distribute requests across all
// of the registered subnets. For more information, see Deploy Elastic Load
// Balancing in Amazon in the Elastic Load Balancing Developer Guide .
func (c *ELB) AttachLoadBalancerToSubnets(req AttachLoadBalancerToSubnetsInput) (resp *AttachLoadBalancerToSubnetsResult, err error) {
	resp = &AttachLoadBalancerToSubnetsResult{}
	err = c.client.Do("AttachLoadBalancerToSubnets", "POST", "/", req, resp)
	return
}

// ConfigureHealthCheck specifies the health check settings to use for
// evaluating the health state of your back-end instances. For more
// information, see Health Check in the Elastic Load Balancing Developer
// Guide
func (c *ELB) ConfigureHealthCheck(req ConfigureHealthCheckInput) (resp *ConfigureHealthCheckResult, err error) {
	resp = &ConfigureHealthCheckResult{}
	err = c.client.Do("ConfigureHealthCheck", "POST", "/", req, resp)
	return
}

// CreateAppCookieStickinessPolicy generates a stickiness policy with
// sticky session lifetimes that follow that of an application-generated
// cookie. This policy can be associated only with listeners. This policy
// is similar to the policy created by CreateLBCookieStickinessPolicy ,
// except that the lifetime of the special Elastic Load Balancing cookie
// follows the lifetime of the application-generated cookie specified in
// the policy configuration. The load balancer only inserts a new
// stickiness cookie when the application response includes a new
// application cookie. If the application cookie is explicitly removed or
// expires, the session stops being sticky until a new application cookie
// is issued. An application client must receive and send two cookies: the
// application-generated cookie and the special Elastic Load Balancing
// cookie named . This is the default behavior for many common web
// browsers. For more information, see Enabling Application-Controlled
// Session Stickiness in the Elastic Load Balancing Developer Guide
func (c *ELB) CreateAppCookieStickinessPolicy(req CreateAppCookieStickinessPolicyInput) (resp *CreateAppCookieStickinessPolicyResult, err error) {
	resp = &CreateAppCookieStickinessPolicyResult{}
	err = c.client.Do("CreateAppCookieStickinessPolicy", "POST", "/", req, resp)
	return
}

// CreateLBCookieStickinessPolicy generates a stickiness policy with sticky
// session lifetimes controlled by the lifetime of the browser (user-agent)
// or a specified expiration period. This policy can be associated only
// with listeners. When a load balancer implements this policy, the load
// balancer uses a special cookie to track the backend server instance for
// each request. When the load balancer receives a request, it first checks
// to see if this cookie is present in the request. If so, the load
// balancer sends the request to the application server specified in the
// cookie. If not, the load balancer sends the request to a server that is
// chosen based on the existing load balancing algorithm. A cookie is
// inserted into the response for binding subsequent requests from the same
// user to that server. The validity of the cookie is based on the cookie
// expiration time, which is specified in the policy configuration. For
// more information, see Enabling Duration-Based Session Stickiness in the
// Elastic Load Balancing Developer Guide
func (c *ELB) CreateLBCookieStickinessPolicy(req CreateLBCookieStickinessPolicyInput) (resp *CreateLBCookieStickinessPolicyResult, err error) {
	resp = &CreateLBCookieStickinessPolicyResult{}
	err = c.client.Do("CreateLBCookieStickinessPolicy", "POST", "/", req, resp)
	return
}

// CreateLoadBalancer creates a new load balancer. After the call has
// completed successfully, a new load balancer is created with a unique
// Domain Name Service name. The DNS name includes the name of the AWS
// region in which the load balance was created. For example, if your load
// balancer was created in the United States, the DNS name might end with
// either of the following: us-east-1.elb.amazonaws.com (for the Northern
// Virginia region) us-west-1.elb.amazonaws.com (for the Northern
// California region) For information about the AWS regions supported by
// Elastic Load Balancing, see Regions and Endpoints You can create up to
// 20 load balancers per region per account. Elastic Load Balancing
// supports load balancing your Amazon EC2 instances launched within any
// one of the following platforms: For information on creating and managing
// your load balancers in EC2-Classic, see Deploy Elastic Load Balancing in
// Amazon EC2-Classic For information on creating and managing your load
// balancers in EC2-VPC, see Deploy Elastic Load Balancing in Amazon
func (c *ELB) CreateLoadBalancer(req CreateAccessPointInput) (resp *CreateLoadBalancerResult, err error) {
	resp = &CreateLoadBalancerResult{}
	err = c.client.Do("CreateLoadBalancer", "POST", "/", req, resp)
	return
}

// CreateLoadBalancerListeners creates one or more listeners on a load
// balancer for the specified port. If a listener with the given port does
// not already exist, it will be created; otherwise, the properties of the
// new listener must match the properties of the existing listener. For
// more information, see Add a Listener to Your Load Balancer in the
// Elastic Load Balancing Developer Guide
func (c *ELB) CreateLoadBalancerListeners(req CreateLoadBalancerListenerInput) (resp *CreateLoadBalancerListenersResult, err error) {
	resp = &CreateLoadBalancerListenersResult{}
	err = c.client.Do("CreateLoadBalancerListeners", "POST", "/", req, resp)
	return
}

// CreateLoadBalancerPolicy creates a new policy that contains the
// necessary attributes depending on the policy type. Policies are settings
// that are saved for your load balancer and that can be applied to the
// front-end listener, or the back-end application server, depending on
// your policy type.
func (c *ELB) CreateLoadBalancerPolicy(req CreateLoadBalancerPolicyInput) (resp *CreateLoadBalancerPolicyResult, err error) {
	resp = &CreateLoadBalancerPolicyResult{}
	err = c.client.Do("CreateLoadBalancerPolicy", "POST", "/", req, resp)
	return
}

// DeleteLoadBalancer deletes the specified load balancer. If attempting to
// recreate the load balancer, you must reconfigure all the settings. The
// DNS name associated with a deleted load balancer will no longer be
// usable. Once deleted, the name and associated DNS record of the load
// balancer no longer exist and traffic sent to any of its IP addresses
// will no longer be delivered to back-end instances. To successfully call
// this you must provide the same account credentials as were used to
// create the load balancer. By design, if the load balancer does not exist
// or has already been deleted, a call to DeleteLoadBalancer action still
// succeeds.
func (c *ELB) DeleteLoadBalancer(req DeleteAccessPointInput) (resp *DeleteLoadBalancerResult, err error) {
	resp = &DeleteLoadBalancerResult{}
	err = c.client.Do("DeleteLoadBalancer", "POST", "/", req, resp)
	return
}

// DeleteLoadBalancerListeners deletes listeners from the load balancer for
// the specified port.
func (c *ELB) DeleteLoadBalancerListeners(req DeleteLoadBalancerListenerInput) (resp *DeleteLoadBalancerListenersResult, err error) {
	resp = &DeleteLoadBalancerListenersResult{}
	err = c.client.Do("DeleteLoadBalancerListeners", "POST", "/", req, resp)
	return
}

// DeleteLoadBalancerPolicy deletes a policy from the load balancer. The
// specified policy must not be enabled for any listeners.
func (c *ELB) DeleteLoadBalancerPolicy(req DeleteLoadBalancerPolicyInput) (resp *DeleteLoadBalancerPolicyResult, err error) {
	resp = &DeleteLoadBalancerPolicyResult{}
	err = c.client.Do("DeleteLoadBalancerPolicy", "POST", "/", req, resp)
	return
}

// DeregisterInstancesFromLoadBalancer deregisters instances from the load
// balancer. Once the instance is deregistered, it will stop receiving
// traffic from the load balancer. In order to successfully call this the
// same account credentials as those used to create the load balancer must
// be provided. For more information, see De-register and Register Amazon
// EC2 Instances in the Elastic Load Balancing Developer Guide You can use
// DescribeLoadBalancers to verify if the instance is deregistered from the
// load balancer.
func (c *ELB) DeregisterInstancesFromLoadBalancer(req DeregisterEndPointsInput) (resp *DeregisterInstancesFromLoadBalancerResult, err error) {
	resp = &DeregisterInstancesFromLoadBalancerResult{}
	err = c.client.Do("DeregisterInstancesFromLoadBalancer", "POST", "/", req, resp)
	return
}

// DescribeInstanceHealth returns the current state of the specified
// instances registered with the specified load balancer. If no instances
// are specified, the state of all the instances registered with the load
// balancer is returned. You must provide the same account credentials as
// those that were used to create the load balancer.
func (c *ELB) DescribeInstanceHealth(req DescribeEndPointStateInput) (resp *DescribeInstanceHealthResult, err error) {
	resp = &DescribeInstanceHealthResult{}
	err = c.client.Do("DescribeInstanceHealth", "POST", "/", req, resp)
	return
}

// DescribeLoadBalancerAttributes returns detailed information about all of
// the attributes associated with the specified load balancer.
func (c *ELB) DescribeLoadBalancerAttributes(req DescribeLoadBalancerAttributesInput) (resp *DescribeLoadBalancerAttributesResult, err error) {
	resp = &DescribeLoadBalancerAttributesResult{}
	err = c.client.Do("DescribeLoadBalancerAttributes", "POST", "/", req, resp)
	return
}

// DescribeLoadBalancerPolicies returns detailed descriptions of the
// policies. If you specify a load balancer name, the action returns the
// descriptions of all the policies created for the load balancer. If you
// specify a policy name associated with your load balancer, the action
// returns the description of that policy. If you don't specify a load
// balancer name, the action returns descriptions of the specified sample
// policies, or descriptions of all the sample policies. The names of the
// sample policies have the ELBSample- prefix.
func (c *ELB) DescribeLoadBalancerPolicies(req DescribeLoadBalancerPoliciesInput) (resp *DescribeLoadBalancerPoliciesResult, err error) {
	resp = &DescribeLoadBalancerPoliciesResult{}
	err = c.client.Do("DescribeLoadBalancerPolicies", "POST", "/", req, resp)
	return
}

// DescribeLoadBalancerPolicyTypes returns meta-information on the
// specified load balancer policies defined by the Elastic Load Balancing
// service. The policy types that are returned from this action can be used
// in a CreateLoadBalancerPolicy action to instantiate specific policy
// configurations that will be applied to a load balancer.
func (c *ELB) DescribeLoadBalancerPolicyTypes(req DescribeLoadBalancerPolicyTypesInput) (resp *DescribeLoadBalancerPolicyTypesResult, err error) {
	resp = &DescribeLoadBalancerPolicyTypesResult{}
	err = c.client.Do("DescribeLoadBalancerPolicyTypes", "POST", "/", req, resp)
	return
}

// DescribeLoadBalancers returns detailed configuration information for all
// the load balancers created for the account. If you specify load balancer
// names, the action returns configuration information of the specified
// load balancers. In order to retrieve this information, you must provide
// the same account credentials that was used to create the load balancer.
func (c *ELB) DescribeLoadBalancers(req DescribeAccessPointsInput) (resp *DescribeLoadBalancersResult, err error) {
	resp = &DescribeLoadBalancersResult{}
	err = c.client.Do("DescribeLoadBalancers", "POST", "/", req, resp)
	return
}

// DescribeTags describes the tags associated with one or more load
// balancers.
func (c *ELB) DescribeTags(req DescribeTagsInput) (resp *DescribeTagsResult, err error) {
	resp = &DescribeTagsResult{}
	err = c.client.Do("DescribeTags", "POST", "/", req, resp)
	return
}

// DetachLoadBalancerFromSubnets removes subnets from the set of configured
// subnets in the Amazon Virtual Private Cloud (Amazon for the load
// balancer. After a subnet is removed all of the EC2 instances registered
// with the load balancer that are in the removed subnet will go into the
// OutOfService state. When a subnet is removed, the load balancer will
// balance the traffic among the remaining routable subnets for the load
// balancer.
func (c *ELB) DetachLoadBalancerFromSubnets(req DetachLoadBalancerFromSubnetsInput) (resp *DetachLoadBalancerFromSubnetsResult, err error) {
	resp = &DetachLoadBalancerFromSubnetsResult{}
	err = c.client.Do("DetachLoadBalancerFromSubnets", "POST", "/", req, resp)
	return
}

// DisableAvailabilityZonesForLoadBalancer removes the specified EC2
// Availability Zones from the set of configured Availability Zones for the
// load balancer. There must be at least one Availability Zone registered
// with a load balancer at all times. Once an Availability Zone is removed,
// all the instances registered with the load balancer that are in the
// removed Availability Zone go into the OutOfService state. Upon
// Availability Zone removal, the load balancer attempts to equally balance
// the traffic among its remaining usable Availability Zones. Trying to
// remove an Availability Zone that was not associated with the load
// balancer does nothing. For more information, see Disable an Availability
// Zone from a Load-Balanced Application in the Elastic Load Balancing
// Developer Guide
func (c *ELB) DisableAvailabilityZonesForLoadBalancer(req RemoveAvailabilityZonesInput) (resp *DisableAvailabilityZonesForLoadBalancerResult, err error) {
	resp = &DisableAvailabilityZonesForLoadBalancerResult{}
	err = c.client.Do("DisableAvailabilityZonesForLoadBalancer", "POST", "/", req, resp)
	return
}

// EnableAvailabilityZonesForLoadBalancer adds one or more EC2 Availability
// Zones to the load balancer. The load balancer evenly distributes
// requests across all its registered Availability Zones that contain
// instances. The new EC2 Availability Zones to be added must be in the
// same EC2 Region as the Availability Zones for which the load balancer
// was created. For more information, see Expand a Load Balanced
// Application to an Additional Availability Zone in the Elastic Load
// Balancing Developer Guide
func (c *ELB) EnableAvailabilityZonesForLoadBalancer(req AddAvailabilityZonesInput) (resp *EnableAvailabilityZonesForLoadBalancerResult, err error) {
	resp = &EnableAvailabilityZonesForLoadBalancerResult{}
	err = c.client.Do("EnableAvailabilityZonesForLoadBalancer", "POST", "/", req, resp)
	return
}

// ModifyLoadBalancerAttributes modifies the attributes of a specified load
// balancer. You can modify the load balancer attributes, such as
// AccessLogs , ConnectionDraining , and CrossZoneLoadBalancing by either
// enabling or disabling them. Or, you can modify the load balancer
// attribute ConnectionSettings by specifying an idle connection timeout
// value for your load balancer. For more information, see the following:
func (c *ELB) ModifyLoadBalancerAttributes(req ModifyLoadBalancerAttributesInput) (resp *ModifyLoadBalancerAttributesResult, err error) {
	resp = &ModifyLoadBalancerAttributesResult{}
	err = c.client.Do("ModifyLoadBalancerAttributes", "POST", "/", req, resp)
	return
}

// RegisterInstancesWithLoadBalancer adds new instances to the load
// balancer. Once the instance is registered, it starts receiving traffic
// and requests from the load balancer. Any instance that is not in any of
// the Availability Zones registered for the load balancer will be moved to
// the OutOfService state. It will move to the InService state when the
// Availability Zone is added to the load balancer. When an instance
// registered with a load balancer is stopped and then restarted, the IP
// addresses associated with the instance changes. Elastic Load Balancing
// cannot recognize the new IP address, which prevents it from routing
// traffic to the instances. We recommend that you de-register your Amazon
// EC2 instances from your load balancer after you stop your instance, and
// then register the load balancer with your instance after you've
// restarted. To de-register your instances from load balancer, use
// DeregisterInstancesFromLoadBalancer action. For more information, see
// De-register and Register Amazon EC2 Instances in the Elastic Load
// Balancing Developer Guide In order for this call to be successful, you
// must provide the same account credentials as those that were used to
// create the load balancer. Completion of this API does not guarantee that
// operation has completed. Rather, it means that the request has been
// registered and the changes will happen shortly. You can use
// DescribeLoadBalancers or DescribeInstanceHealth action to check the
// state of the newly registered instances.
func (c *ELB) RegisterInstancesWithLoadBalancer(req RegisterEndPointsInput) (resp *RegisterInstancesWithLoadBalancerResult, err error) {
	resp = &RegisterInstancesWithLoadBalancerResult{}
	err = c.client.Do("RegisterInstancesWithLoadBalancer", "POST", "/", req, resp)
	return
}

// RemoveTags removes one or more tags from the specified load balancer.
func (c *ELB) RemoveTags(req RemoveTagsInput) (resp *RemoveTagsResult, err error) {
	resp = &RemoveTagsResult{}
	err = c.client.Do("RemoveTags", "POST", "/", req, resp)
	return
}

// SetLoadBalancerListenerSSLCertificate sets the certificate that
// terminates the specified listener's SSL connections. The specified
// certificate replaces any prior certificate that was used on the same
// load balancer and port. For more information on updating your SSL
// certificate, see Updating an SSL Certificate for a Load Balancer in the
// Elastic Load Balancing Developer Guide
func (c *ELB) SetLoadBalancerListenerSSLCertificate(req SetLoadBalancerListenerSSLCertificateInput) (resp *SetLoadBalancerListenerSSLCertificateResult, err error) {
	resp = &SetLoadBalancerListenerSSLCertificateResult{}
	err = c.client.Do("SetLoadBalancerListenerSSLCertificate", "POST", "/", req, resp)
	return
}

// SetLoadBalancerPoliciesForBackendServer replaces the current set of
// policies associated with a port on which the back-end server is
// listening with a new set of policies. After the policies have been
// created using CreateLoadBalancerPolicy , they can be applied here as a
// list. At this time, only the back-end server authentication policy type
// can be applied to the back-end ports; this policy type is composed of
// multiple public key policies. The
// SetLoadBalancerPoliciesForBackendServer replaces the current set of
// policies associated with the specified instance port. Every time you use
// this action to enable the policies, use the PolicyNames parameter to
// list all the policies you want to enable. You can use
// DescribeLoadBalancers or DescribeLoadBalancerPolicies action to verify
// that the policy has been associated with the back-end server.
func (c *ELB) SetLoadBalancerPoliciesForBackendServer(req SetLoadBalancerPoliciesForBackendServerInput) (resp *SetLoadBalancerPoliciesForBackendServerResult, err error) {
	resp = &SetLoadBalancerPoliciesForBackendServerResult{}
	err = c.client.Do("SetLoadBalancerPoliciesForBackendServer", "POST", "/", req, resp)
	return
}

// SetLoadBalancerPoliciesOfListener associates, updates, or disables a
// policy with a listener on the load balancer. You can associate multiple
// policies with a listener.
func (c *ELB) SetLoadBalancerPoliciesOfListener(req SetLoadBalancerPoliciesOfListenerInput) (resp *SetLoadBalancerPoliciesOfListenerResult, err error) {
	resp = &SetLoadBalancerPoliciesOfListenerResult{}
	err = c.client.Do("SetLoadBalancerPoliciesOfListener", "POST", "/", req, resp)
	return
}

// AccessLog is undocumented.
type AccessLog struct {
	EmitInterval   int    `xml:"EmitInterval"`
	Enabled        bool   `xml:"Enabled"`
	S3BucketName   string `xml:"S3BucketName"`
	S3BucketPrefix string `xml:"S3BucketPrefix"`
}

// AddAvailabilityZonesInput is undocumented.
type AddAvailabilityZonesInput struct {
	AvailabilityZones []string `xml:"AvailabilityZones>member"`
	LoadBalancerName  string   `xml:"LoadBalancerName"`
}

// AddAvailabilityZonesOutput is undocumented.
type AddAvailabilityZonesOutput struct {
	AvailabilityZones []string `xml:"EnableAvailabilityZonesForLoadBalancerResult>AvailabilityZones>member"`
}

// AddTagsInput is undocumented.
type AddTagsInput struct {
	LoadBalancerNames []string `xml:"LoadBalancerNames>member"`
	Tags              []Tag    `xml:"Tags>member"`
}

// AddTagsOutput is undocumented.
type AddTagsOutput struct {
}

// AdditionalAttribute is undocumented.
type AdditionalAttribute struct {
	Key   string `xml:"Key"`
	Value string `xml:"Value"`
}

// AppCookieStickinessPolicy is undocumented.
type AppCookieStickinessPolicy struct {
	CookieName string `xml:"CookieName"`
	PolicyName string `xml:"PolicyName"`
}

// ApplySecurityGroupsToLoadBalancerInput is undocumented.
type ApplySecurityGroupsToLoadBalancerInput struct {
	LoadBalancerName string   `xml:"LoadBalancerName"`
	SecurityGroups   []string `xml:"SecurityGroups>member"`
}

// ApplySecurityGroupsToLoadBalancerOutput is undocumented.
type ApplySecurityGroupsToLoadBalancerOutput struct {
	SecurityGroups []string `xml:"ApplySecurityGroupsToLoadBalancerResult>SecurityGroups>member"`
}

// AttachLoadBalancerToSubnetsInput is undocumented.
type AttachLoadBalancerToSubnetsInput struct {
	LoadBalancerName string   `xml:"LoadBalancerName"`
	Subnets          []string `xml:"Subnets>member"`
}

// AttachLoadBalancerToSubnetsOutput is undocumented.
type AttachLoadBalancerToSubnetsOutput struct {
	Subnets []string `xml:"AttachLoadBalancerToSubnetsResult>Subnets>member"`
}

// BackendServerDescription is undocumented.
type BackendServerDescription struct {
	InstancePort int      `xml:"InstancePort"`
	PolicyNames  []string `xml:"PolicyNames>member"`
}

// ConfigureHealthCheckInput is undocumented.
type ConfigureHealthCheckInput struct {
	HealthCheck      HealthCheck `xml:"HealthCheck"`
	LoadBalancerName string      `xml:"LoadBalancerName"`
}

// ConfigureHealthCheckOutput is undocumented.
type ConfigureHealthCheckOutput struct {
	HealthCheck HealthCheck `xml:"ConfigureHealthCheckResult>HealthCheck"`
}

// ConnectionDraining is undocumented.
type ConnectionDraining struct {
	Enabled bool `xml:"Enabled"`
	Timeout int  `xml:"Timeout"`
}

// ConnectionSettings is undocumented.
type ConnectionSettings struct {
	IdleTimeout int `xml:"IdleTimeout"`
}

// CreateAccessPointInput is undocumented.
type CreateAccessPointInput struct {
	AvailabilityZones []string   `xml:"AvailabilityZones>member"`
	Listeners         []Listener `xml:"Listeners>member"`
	LoadBalancerName  string     `xml:"LoadBalancerName"`
	Scheme            string     `xml:"Scheme"`
	SecurityGroups    []string   `xml:"SecurityGroups>member"`
	Subnets           []string   `xml:"Subnets>member"`
	Tags              []Tag      `xml:"Tags>member"`
}

// CreateAccessPointOutput is undocumented.
type CreateAccessPointOutput struct {
	DNSName string `xml:"CreateLoadBalancerResult>DNSName"`
}

// CreateAppCookieStickinessPolicyInput is undocumented.
type CreateAppCookieStickinessPolicyInput struct {
	CookieName       string `xml:"CookieName"`
	LoadBalancerName string `xml:"LoadBalancerName"`
	PolicyName       string `xml:"PolicyName"`
}

// CreateAppCookieStickinessPolicyOutput is undocumented.
type CreateAppCookieStickinessPolicyOutput struct {
}

// CreateLBCookieStickinessPolicyInput is undocumented.
type CreateLBCookieStickinessPolicyInput struct {
	CookieExpirationPeriod int64  `xml:"CookieExpirationPeriod"`
	LoadBalancerName       string `xml:"LoadBalancerName"`
	PolicyName             string `xml:"PolicyName"`
}

// CreateLBCookieStickinessPolicyOutput is undocumented.
type CreateLBCookieStickinessPolicyOutput struct {
}

// CreateLoadBalancerListenerInput is undocumented.
type CreateLoadBalancerListenerInput struct {
	Listeners        []Listener `xml:"Listeners>member"`
	LoadBalancerName string     `xml:"LoadBalancerName"`
}

// CreateLoadBalancerListenerOutput is undocumented.
type CreateLoadBalancerListenerOutput struct {
}

// CreateLoadBalancerPolicyInput is undocumented.
type CreateLoadBalancerPolicyInput struct {
	LoadBalancerName string            `xml:"LoadBalancerName"`
	PolicyAttributes []PolicyAttribute `xml:"PolicyAttributes>member"`
	PolicyName       string            `xml:"PolicyName"`
	PolicyTypeName   string            `xml:"PolicyTypeName"`
}

// CreateLoadBalancerPolicyOutput is undocumented.
type CreateLoadBalancerPolicyOutput struct {
}

// CrossZoneLoadBalancing is undocumented.
type CrossZoneLoadBalancing struct {
	Enabled bool `xml:"Enabled"`
}

// DeleteAccessPointInput is undocumented.
type DeleteAccessPointInput struct {
	LoadBalancerName string `xml:"LoadBalancerName"`
}

// DeleteAccessPointOutput is undocumented.
type DeleteAccessPointOutput struct {
}

// DeleteLoadBalancerListenerInput is undocumented.
type DeleteLoadBalancerListenerInput struct {
	LoadBalancerName  string `xml:"LoadBalancerName"`
	LoadBalancerPorts []int  `xml:"LoadBalancerPorts>member"`
}

// DeleteLoadBalancerListenerOutput is undocumented.
type DeleteLoadBalancerListenerOutput struct {
}

// DeleteLoadBalancerPolicyInput is undocumented.
type DeleteLoadBalancerPolicyInput struct {
	LoadBalancerName string `xml:"LoadBalancerName"`
	PolicyName       string `xml:"PolicyName"`
}

// DeleteLoadBalancerPolicyOutput is undocumented.
type DeleteLoadBalancerPolicyOutput struct {
}

// DeregisterEndPointsInput is undocumented.
type DeregisterEndPointsInput struct {
	Instances        []Instance `xml:"Instances>member"`
	LoadBalancerName string     `xml:"LoadBalancerName"`
}

// DeregisterEndPointsOutput is undocumented.
type DeregisterEndPointsOutput struct {
	Instances []Instance `xml:"DeregisterInstancesFromLoadBalancerResult>Instances>member"`
}

// DescribeAccessPointsInput is undocumented.
type DescribeAccessPointsInput struct {
	LoadBalancerNames []string `xml:"LoadBalancerNames>member"`
	Marker            string   `xml:"Marker"`
	PageSize          int      `xml:"PageSize"`
}

// DescribeAccessPointsOutput is undocumented.
type DescribeAccessPointsOutput struct {
	LoadBalancerDescriptions []LoadBalancerDescription `xml:"DescribeLoadBalancersResult>LoadBalancerDescriptions>member"`
	NextMarker               string                    `xml:"DescribeLoadBalancersResult>NextMarker"`
}

// DescribeEndPointStateInput is undocumented.
type DescribeEndPointStateInput struct {
	Instances        []Instance `xml:"Instances>member"`
	LoadBalancerName string     `xml:"LoadBalancerName"`
}

// DescribeEndPointStateOutput is undocumented.
type DescribeEndPointStateOutput struct {
	InstanceStates []InstanceState `xml:"DescribeInstanceHealthResult>InstanceStates>member"`
}

// DescribeLoadBalancerAttributesInput is undocumented.
type DescribeLoadBalancerAttributesInput struct {
	LoadBalancerName string `xml:"LoadBalancerName"`
}

// DescribeLoadBalancerAttributesOutput is undocumented.
type DescribeLoadBalancerAttributesOutput struct {
	LoadBalancerAttributes LoadBalancerAttributes `xml:"DescribeLoadBalancerAttributesResult>LoadBalancerAttributes"`
}

// DescribeLoadBalancerPoliciesInput is undocumented.
type DescribeLoadBalancerPoliciesInput struct {
	LoadBalancerName string   `xml:"LoadBalancerName"`
	PolicyNames      []string `xml:"PolicyNames>member"`
}

// DescribeLoadBalancerPoliciesOutput is undocumented.
type DescribeLoadBalancerPoliciesOutput struct {
	PolicyDescriptions []PolicyDescription `xml:"DescribeLoadBalancerPoliciesResult>PolicyDescriptions>member"`
}

// DescribeLoadBalancerPolicyTypesInput is undocumented.
type DescribeLoadBalancerPolicyTypesInput struct {
	PolicyTypeNames []string `xml:"PolicyTypeNames>member"`
}

// DescribeLoadBalancerPolicyTypesOutput is undocumented.
type DescribeLoadBalancerPolicyTypesOutput struct {
	PolicyTypeDescriptions []PolicyTypeDescription `xml:"DescribeLoadBalancerPolicyTypesResult>PolicyTypeDescriptions>member"`
}

// DescribeTagsInput is undocumented.
type DescribeTagsInput struct {
	LoadBalancerNames []string `xml:"LoadBalancerNames>member"`
}

// DescribeTagsOutput is undocumented.
type DescribeTagsOutput struct {
	TagDescriptions []TagDescription `xml:"DescribeTagsResult>TagDescriptions>member"`
}

// DetachLoadBalancerFromSubnetsInput is undocumented.
type DetachLoadBalancerFromSubnetsInput struct {
	LoadBalancerName string   `xml:"LoadBalancerName"`
	Subnets          []string `xml:"Subnets>member"`
}

// DetachLoadBalancerFromSubnetsOutput is undocumented.
type DetachLoadBalancerFromSubnetsOutput struct {
	Subnets []string `xml:"DetachLoadBalancerFromSubnetsResult>Subnets>member"`
}

// HealthCheck is undocumented.
type HealthCheck struct {
	HealthyThreshold   int    `xml:"HealthyThreshold"`
	Interval           int    `xml:"Interval"`
	Target             string `xml:"Target"`
	Timeout            int    `xml:"Timeout"`
	UnhealthyThreshold int    `xml:"UnhealthyThreshold"`
}

// Instance is undocumented.
type Instance struct {
	InstanceID string `xml:"InstanceId"`
}

// InstanceState is undocumented.
type InstanceState struct {
	Description string `xml:"Description"`
	InstanceID  string `xml:"InstanceId"`
	ReasonCode  string `xml:"ReasonCode"`
	State       string `xml:"State"`
}

// LBCookieStickinessPolicy is undocumented.
type LBCookieStickinessPolicy struct {
	CookieExpirationPeriod int64  `xml:"CookieExpirationPeriod"`
	PolicyName             string `xml:"PolicyName"`
}

// Listener is undocumented.
type Listener struct {
	InstancePort     int    `xml:"InstancePort"`
	InstanceProtocol string `xml:"InstanceProtocol"`
	LoadBalancerPort int    `xml:"LoadBalancerPort"`
	Protocol         string `xml:"Protocol"`
	SSLCertificateID string `xml:"SSLCertificateId"`
}

// ListenerDescription is undocumented.
type ListenerDescription struct {
	Listener    Listener `xml:"Listener"`
	PolicyNames []string `xml:"PolicyNames>member"`
}

// LoadBalancerAttributes is undocumented.
type LoadBalancerAttributes struct {
	AccessLog              AccessLog              `xml:"AccessLog"`
	AdditionalAttributes   []AdditionalAttribute  `xml:"AdditionalAttributes>member"`
	ConnectionDraining     ConnectionDraining     `xml:"ConnectionDraining"`
	ConnectionSettings     ConnectionSettings     `xml:"ConnectionSettings"`
	CrossZoneLoadBalancing CrossZoneLoadBalancing `xml:"CrossZoneLoadBalancing"`
}

// LoadBalancerDescription is undocumented.
type LoadBalancerDescription struct {
	AvailabilityZones         []string                   `xml:"AvailabilityZones>member"`
	BackendServerDescriptions []BackendServerDescription `xml:"BackendServerDescriptions>member"`
	CanonicalHostedZoneName   string                     `xml:"CanonicalHostedZoneName"`
	CanonicalHostedZoneNameID string                     `xml:"CanonicalHostedZoneNameID"`
	CreatedTime               time.Time                  `xml:"CreatedTime"`
	DNSName                   string                     `xml:"DNSName"`
	HealthCheck               HealthCheck                `xml:"HealthCheck"`
	Instances                 []Instance                 `xml:"Instances>member"`
	ListenerDescriptions      []ListenerDescription      `xml:"ListenerDescriptions>member"`
	LoadBalancerName          string                     `xml:"LoadBalancerName"`
	Policies                  Policies                   `xml:"Policies"`
	Scheme                    string                     `xml:"Scheme"`
	SecurityGroups            []string                   `xml:"SecurityGroups>member"`
	SourceSecurityGroup       SourceSecurityGroup        `xml:"SourceSecurityGroup"`
	Subnets                   []string                   `xml:"Subnets>member"`
	VPCID                     string                     `xml:"VPCId"`
}

// ModifyLoadBalancerAttributesInput is undocumented.
type ModifyLoadBalancerAttributesInput struct {
	LoadBalancerAttributes LoadBalancerAttributes `xml:"LoadBalancerAttributes"`
	LoadBalancerName       string                 `xml:"LoadBalancerName"`
}

// ModifyLoadBalancerAttributesOutput is undocumented.
type ModifyLoadBalancerAttributesOutput struct {
	LoadBalancerAttributes LoadBalancerAttributes `xml:"ModifyLoadBalancerAttributesResult>LoadBalancerAttributes"`
	LoadBalancerName       string                 `xml:"ModifyLoadBalancerAttributesResult>LoadBalancerName"`
}

// Policies is undocumented.
type Policies struct {
	AppCookieStickinessPolicies []AppCookieStickinessPolicy `xml:"AppCookieStickinessPolicies>member"`
	LBCookieStickinessPolicies  []LBCookieStickinessPolicy  `xml:"LBCookieStickinessPolicies>member"`
	OtherPolicies               []string                    `xml:"OtherPolicies>member"`
}

// PolicyAttribute is undocumented.
type PolicyAttribute struct {
	AttributeName  string `xml:"AttributeName"`
	AttributeValue string `xml:"AttributeValue"`
}

// PolicyAttributeDescription is undocumented.
type PolicyAttributeDescription struct {
	AttributeName  string `xml:"AttributeName"`
	AttributeValue string `xml:"AttributeValue"`
}

// PolicyAttributeTypeDescription is undocumented.
type PolicyAttributeTypeDescription struct {
	AttributeName string `xml:"AttributeName"`
	AttributeType string `xml:"AttributeType"`
	Cardinality   string `xml:"Cardinality"`
	DefaultValue  string `xml:"DefaultValue"`
	Description   string `xml:"Description"`
}

// PolicyDescription is undocumented.
type PolicyDescription struct {
	PolicyAttributeDescriptions []PolicyAttributeDescription `xml:"PolicyAttributeDescriptions>member"`
	PolicyName                  string                       `xml:"PolicyName"`
	PolicyTypeName              string                       `xml:"PolicyTypeName"`
}

// PolicyTypeDescription is undocumented.
type PolicyTypeDescription struct {
	Description                     string                           `xml:"Description"`
	PolicyAttributeTypeDescriptions []PolicyAttributeTypeDescription `xml:"PolicyAttributeTypeDescriptions>member"`
	PolicyTypeName                  string                           `xml:"PolicyTypeName"`
}

// RegisterEndPointsInput is undocumented.
type RegisterEndPointsInput struct {
	Instances        []Instance `xml:"Instances>member"`
	LoadBalancerName string     `xml:"LoadBalancerName"`
}

// RegisterEndPointsOutput is undocumented.
type RegisterEndPointsOutput struct {
	Instances []Instance `xml:"RegisterInstancesWithLoadBalancerResult>Instances>member"`
}

// RemoveAvailabilityZonesInput is undocumented.
type RemoveAvailabilityZonesInput struct {
	AvailabilityZones []string `xml:"AvailabilityZones>member"`
	LoadBalancerName  string   `xml:"LoadBalancerName"`
}

// RemoveAvailabilityZonesOutput is undocumented.
type RemoveAvailabilityZonesOutput struct {
	AvailabilityZones []string `xml:"DisableAvailabilityZonesForLoadBalancerResult>AvailabilityZones>member"`
}

// RemoveTagsInput is undocumented.
type RemoveTagsInput struct {
	LoadBalancerNames []string     `xml:"LoadBalancerNames>member"`
	Tags              []TagKeyOnly `xml:"Tags>member"`
}

// RemoveTagsOutput is undocumented.
type RemoveTagsOutput struct {
}

// SetLoadBalancerListenerSSLCertificateInput is undocumented.
type SetLoadBalancerListenerSSLCertificateInput struct {
	LoadBalancerName string `xml:"LoadBalancerName"`
	LoadBalancerPort int    `xml:"LoadBalancerPort"`
	SSLCertificateID string `xml:"SSLCertificateId"`
}

// SetLoadBalancerListenerSSLCertificateOutput is undocumented.
type SetLoadBalancerListenerSSLCertificateOutput struct {
}

// SetLoadBalancerPoliciesForBackendServerInput is undocumented.
type SetLoadBalancerPoliciesForBackendServerInput struct {
	InstancePort     int      `xml:"InstancePort"`
	LoadBalancerName string   `xml:"LoadBalancerName"`
	PolicyNames      []string `xml:"PolicyNames>member"`
}

// SetLoadBalancerPoliciesForBackendServerOutput is undocumented.
type SetLoadBalancerPoliciesForBackendServerOutput struct {
}

// SetLoadBalancerPoliciesOfListenerInput is undocumented.
type SetLoadBalancerPoliciesOfListenerInput struct {
	LoadBalancerName string   `xml:"LoadBalancerName"`
	LoadBalancerPort int      `xml:"LoadBalancerPort"`
	PolicyNames      []string `xml:"PolicyNames>member"`
}

// SetLoadBalancerPoliciesOfListenerOutput is undocumented.
type SetLoadBalancerPoliciesOfListenerOutput struct {
}

// SourceSecurityGroup is undocumented.
type SourceSecurityGroup struct {
	GroupName  string `xml:"GroupName"`
	OwnerAlias string `xml:"OwnerAlias"`
}

// Tag is undocumented.
type Tag struct {
	Key   string `xml:"Key"`
	Value string `xml:"Value"`
}

// TagDescription is undocumented.
type TagDescription struct {
	LoadBalancerName string `xml:"LoadBalancerName"`
	Tags             []Tag  `xml:"Tags>member"`
}

// TagKeyOnly is undocumented.
type TagKeyOnly struct {
	Key string `xml:"Key"`
}

// AddTagsResult is a wrapper for AddTagsOutput.
type AddTagsResult struct {
	XMLName xml.Name `xml:"AddTagsResponse"`
}

// ApplySecurityGroupsToLoadBalancerResult is a wrapper for ApplySecurityGroupsToLoadBalancerOutput.
type ApplySecurityGroupsToLoadBalancerResult struct {
	XMLName xml.Name `xml:"ApplySecurityGroupsToLoadBalancerResponse"`

	SecurityGroups []string `xml:"ApplySecurityGroupsToLoadBalancerResult>SecurityGroups>member"`
}

// AttachLoadBalancerToSubnetsResult is a wrapper for AttachLoadBalancerToSubnetsOutput.
type AttachLoadBalancerToSubnetsResult struct {
	XMLName xml.Name `xml:"AttachLoadBalancerToSubnetsResponse"`

	Subnets []string `xml:"AttachLoadBalancerToSubnetsResult>Subnets>member"`
}

// ConfigureHealthCheckResult is a wrapper for ConfigureHealthCheckOutput.
type ConfigureHealthCheckResult struct {
	XMLName xml.Name `xml:"ConfigureHealthCheckResponse"`

	HealthCheck HealthCheck `xml:"ConfigureHealthCheckResult>HealthCheck"`
}

// CreateAppCookieStickinessPolicyResult is a wrapper for CreateAppCookieStickinessPolicyOutput.
type CreateAppCookieStickinessPolicyResult struct {
	XMLName xml.Name `xml:"CreateAppCookieStickinessPolicyResponse"`
}

// CreateLBCookieStickinessPolicyResult is a wrapper for CreateLBCookieStickinessPolicyOutput.
type CreateLBCookieStickinessPolicyResult struct {
	XMLName xml.Name `xml:"CreateLBCookieStickinessPolicyResponse"`
}

// CreateLoadBalancerListenersResult is a wrapper for CreateLoadBalancerListenerOutput.
type CreateLoadBalancerListenersResult struct {
	XMLName xml.Name `xml:"CreateLoadBalancerListenersResponse"`
}

// CreateLoadBalancerPolicyResult is a wrapper for CreateLoadBalancerPolicyOutput.
type CreateLoadBalancerPolicyResult struct {
	XMLName xml.Name `xml:"CreateLoadBalancerPolicyResponse"`
}

// CreateLoadBalancerResult is a wrapper for CreateAccessPointOutput.
type CreateLoadBalancerResult struct {
	XMLName xml.Name `xml:"CreateLoadBalancerResponse"`

	DNSName string `xml:"CreateLoadBalancerResult>DNSName"`
}

// DeleteLoadBalancerListenersResult is a wrapper for DeleteLoadBalancerListenerOutput.
type DeleteLoadBalancerListenersResult struct {
	XMLName xml.Name `xml:"DeleteLoadBalancerListenersResponse"`
}

// DeleteLoadBalancerPolicyResult is a wrapper for DeleteLoadBalancerPolicyOutput.
type DeleteLoadBalancerPolicyResult struct {
	XMLName xml.Name `xml:"DeleteLoadBalancerPolicyResponse"`
}

// DeleteLoadBalancerResult is a wrapper for DeleteAccessPointOutput.
type DeleteLoadBalancerResult struct {
	XMLName xml.Name `xml:"DeleteLoadBalancerResponse"`
}

// DeregisterInstancesFromLoadBalancerResult is a wrapper for DeregisterEndPointsOutput.
type DeregisterInstancesFromLoadBalancerResult struct {
	XMLName xml.Name `xml:"DeregisterInstancesFromLoadBalancerResponse"`

	Instances []Instance `xml:"DeregisterInstancesFromLoadBalancerResult>Instances>member"`
}

// DescribeInstanceHealthResult is a wrapper for DescribeEndPointStateOutput.
type DescribeInstanceHealthResult struct {
	XMLName xml.Name `xml:"DescribeInstanceHealthResponse"`

	InstanceStates []InstanceState `xml:"DescribeInstanceHealthResult>InstanceStates>member"`
}

// DescribeLoadBalancerAttributesResult is a wrapper for DescribeLoadBalancerAttributesOutput.
type DescribeLoadBalancerAttributesResult struct {
	XMLName xml.Name `xml:"DescribeLoadBalancerAttributesResponse"`

	LoadBalancerAttributes LoadBalancerAttributes `xml:"DescribeLoadBalancerAttributesResult>LoadBalancerAttributes"`
}

// DescribeLoadBalancerPoliciesResult is a wrapper for DescribeLoadBalancerPoliciesOutput.
type DescribeLoadBalancerPoliciesResult struct {
	XMLName xml.Name `xml:"DescribeLoadBalancerPoliciesResponse"`

	PolicyDescriptions []PolicyDescription `xml:"DescribeLoadBalancerPoliciesResult>PolicyDescriptions>member"`
}

// DescribeLoadBalancerPolicyTypesResult is a wrapper for DescribeLoadBalancerPolicyTypesOutput.
type DescribeLoadBalancerPolicyTypesResult struct {
	XMLName xml.Name `xml:"DescribeLoadBalancerPolicyTypesResponse"`

	PolicyTypeDescriptions []PolicyTypeDescription `xml:"DescribeLoadBalancerPolicyTypesResult>PolicyTypeDescriptions>member"`
}

// DescribeLoadBalancersResult is a wrapper for DescribeAccessPointsOutput.
type DescribeLoadBalancersResult struct {
	XMLName xml.Name `xml:"DescribeLoadBalancersResponse"`

	LoadBalancerDescriptions []LoadBalancerDescription `xml:"DescribeLoadBalancersResult>LoadBalancerDescriptions>member"`
	NextMarker               string                    `xml:"DescribeLoadBalancersResult>NextMarker"`
}

// DescribeTagsResult is a wrapper for DescribeTagsOutput.
type DescribeTagsResult struct {
	XMLName xml.Name `xml:"DescribeTagsResponse"`

	TagDescriptions []TagDescription `xml:"DescribeTagsResult>TagDescriptions>member"`
}

// DetachLoadBalancerFromSubnetsResult is a wrapper for DetachLoadBalancerFromSubnetsOutput.
type DetachLoadBalancerFromSubnetsResult struct {
	XMLName xml.Name `xml:"DetachLoadBalancerFromSubnetsResponse"`

	Subnets []string `xml:"DetachLoadBalancerFromSubnetsResult>Subnets>member"`
}

// DisableAvailabilityZonesForLoadBalancerResult is a wrapper for RemoveAvailabilityZonesOutput.
type DisableAvailabilityZonesForLoadBalancerResult struct {
	XMLName xml.Name `xml:"DisableAvailabilityZonesForLoadBalancerResponse"`

	AvailabilityZones []string `xml:"DisableAvailabilityZonesForLoadBalancerResult>AvailabilityZones>member"`
}

// EnableAvailabilityZonesForLoadBalancerResult is a wrapper for AddAvailabilityZonesOutput.
type EnableAvailabilityZonesForLoadBalancerResult struct {
	XMLName xml.Name `xml:"EnableAvailabilityZonesForLoadBalancerResponse"`

	AvailabilityZones []string `xml:"EnableAvailabilityZonesForLoadBalancerResult>AvailabilityZones>member"`
}

// ModifyLoadBalancerAttributesResult is a wrapper for ModifyLoadBalancerAttributesOutput.
type ModifyLoadBalancerAttributesResult struct {
	XMLName xml.Name `xml:"ModifyLoadBalancerAttributesResponse"`

	LoadBalancerAttributes LoadBalancerAttributes `xml:"ModifyLoadBalancerAttributesResult>LoadBalancerAttributes"`
	LoadBalancerName       string                 `xml:"ModifyLoadBalancerAttributesResult>LoadBalancerName"`
}

// RegisterInstancesWithLoadBalancerResult is a wrapper for RegisterEndPointsOutput.
type RegisterInstancesWithLoadBalancerResult struct {
	XMLName xml.Name `xml:"RegisterInstancesWithLoadBalancerResponse"`

	Instances []Instance `xml:"RegisterInstancesWithLoadBalancerResult>Instances>member"`
}

// RemoveTagsResult is a wrapper for RemoveTagsOutput.
type RemoveTagsResult struct {
	XMLName xml.Name `xml:"RemoveTagsResponse"`
}

// SetLoadBalancerListenerSSLCertificateResult is a wrapper for SetLoadBalancerListenerSSLCertificateOutput.
type SetLoadBalancerListenerSSLCertificateResult struct {
	XMLName xml.Name `xml:"SetLoadBalancerListenerSSLCertificateResponse"`
}

// SetLoadBalancerPoliciesForBackendServerResult is a wrapper for SetLoadBalancerPoliciesForBackendServerOutput.
type SetLoadBalancerPoliciesForBackendServerResult struct {
	XMLName xml.Name `xml:"SetLoadBalancerPoliciesForBackendServerResponse"`
}

// SetLoadBalancerPoliciesOfListenerResult is a wrapper for SetLoadBalancerPoliciesOfListenerOutput.
type SetLoadBalancerPoliciesOfListenerResult struct {
	XMLName xml.Name `xml:"SetLoadBalancerPoliciesOfListenerResponse"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
