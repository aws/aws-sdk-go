// Package ec2 provides a client for Amazon Elastic Compute Cloud.
package ec2

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// EC2 is a client for Amazon Elastic Compute Cloud.
type EC2 struct {
	client *aws.QueryClient
}

// New returns a new EC2 client.
func New(key, secret, region string, client *http.Client) *EC2 {
	if client == nil {
		client = http.DefaultClient
	}

	service := "ec2"
	endpoint, service, region := endpoints.Lookup("ec2", region)

	return &EC2{
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
			APIVersion: "2014-10-01",
		},
	}
}

// AcceptVpcPeeringConnection accept a VPC peering connection request. To
// accept a request, the VPC peering connection must be in the
// pending-acceptance state, and you must be the owner of the peer Use the
// DescribeVpcPeeringConnections request to view your outstanding VPC
// peering connection requests.
func (c *EC2) AcceptVpcPeeringConnection(req AcceptVpcPeeringConnectionRequest) (resp *AcceptVpcPeeringConnectionResult, err error) {
	resp = &AcceptVpcPeeringConnectionResult{}
	err = c.client.Do("AcceptVpcPeeringConnection", "POST", "/", req, resp)
	return
}

// AllocateAddress acquires an Elastic IP address. An Elastic IP address is
// for use either in the EC2-Classic platform or in a For more information,
// see Elastic IP Addresses in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) AllocateAddress(req AllocateAddressRequest) (resp *AllocateAddressResult, err error) {
	resp = &AllocateAddressResult{}
	err = c.client.Do("AllocateAddress", "POST", "/", req, resp)
	return
}

// AssignPrivateIPAddresses assigns one or more secondary private IP
// addresses to the specified network interface. You can specify one or
// more specific secondary IP addresses, or you can specify the number of
// secondary IP addresses to be automatically assigned within the subnet's
// block range. The number of secondary IP addresses that you can assign to
// an instance varies by instance type. For information about instance
// types, see Instance Types in the Amazon Elastic Compute Cloud User Guide
// . For more information about Elastic IP addresses, see Elastic IP
// Addresses in the Amazon Elastic Compute Cloud User Guide
// AssignPrivateIpAddresses is available only in EC2-VPC.
func (c *EC2) AssignPrivateIPAddresses(req AssignPrivateIPAddressesRequest) (err error) {
	// NRE
	err = c.client.Do("AssignPrivateIpAddresses", "POST", "/", req, nil)
	return
}

// AssociateAddress associates an Elastic IP address with an instance or a
// network interface. An Elastic IP address is for use in either the
// EC2-Classic platform or in a For more information, see Elastic IP
// Addresses in the Amazon Elastic Compute Cloud User Guide [EC2-Classic,
// VPC in an EC2-VPC-only account] If the Elastic IP address is already
// associated with a different instance, it is disassociated from that
// instance and associated with the specified instance. in an EC2-Classic
// account] If you don't specify a private IP address, the Elastic IP
// address is associated with the primary IP address. If the Elastic IP
// address is already associated with a different instance or a network
// interface, you get an error unless you allow reassociation. This is an
// idempotent operation. If you perform the operation more than once,
// Amazon EC2 doesn't return an error.
func (c *EC2) AssociateAddress(req AssociateAddressRequest) (resp *AssociateAddressResult, err error) {
	resp = &AssociateAddressResult{}
	err = c.client.Do("AssociateAddress", "POST", "/", req, resp)
	return
}

// AssociateDhcpOptions associates a set of options (that you've previously
// created) with the specified or associates no options with the After you
// associate the options with the any existing instances and all new
// instances that you launch in that VPC use the options. You don't need to
// restart or relaunch the instances. They automatically pick up the
// changes within a few hours, depending on how frequently the instance
// renews its lease. You can explicitly renew the lease using the operating
// system on the instance. For more information, see Options Sets in the
// Amazon Virtual Private Cloud User Guide
func (c *EC2) AssociateDhcpOptions(req AssociateDhcpOptionsRequest) (err error) {
	// NRE
	err = c.client.Do("AssociateDhcpOptions", "POST", "/", req, nil)
	return
}

// AssociateRouteTable associates a subnet with a route table. The subnet
// and route table must be in the same This association causes traffic
// originating from the subnet to be routed according to the routes in the
// route table. The action returns an association ID, which you need in
// order to disassociate the route table from the subnet later. A route
// table can be associated with multiple subnets. For more information
// about route tables, see Route Tables in the Amazon Virtual Private Cloud
// User Guide
func (c *EC2) AssociateRouteTable(req AssociateRouteTableRequest) (resp *AssociateRouteTableResult, err error) {
	resp = &AssociateRouteTableResult{}
	err = c.client.Do("AssociateRouteTable", "POST", "/", req, resp)
	return
}

// AttachInternetGateway attaches an Internet gateway to a enabling
// connectivity between the Internet and the For more information about
// your VPC and Internet gateway, see the Amazon Virtual Private Cloud User
// Guide
func (c *EC2) AttachInternetGateway(req AttachInternetGatewayRequest) (err error) {
	// NRE
	err = c.client.Do("AttachInternetGateway", "POST", "/", req, nil)
	return
}

// AttachNetworkInterface is undocumented.
func (c *EC2) AttachNetworkInterface(req AttachNetworkInterfaceRequest) (resp *AttachNetworkInterfaceResult, err error) {
	resp = &AttachNetworkInterfaceResult{}
	err = c.client.Do("AttachNetworkInterface", "POST", "/", req, resp)
	return
}

// AttachVolume attaches an Amazon EBS volume to a running or stopped
// instance and exposes it to the instance with the specified device name.
// Encrypted Amazon EBS volumes may only be attached to instances that
// support Amazon EBS encryption. For more information, see Amazon EBS
// Encryption in the Amazon Elastic Compute Cloud User Guide For a list of
// supported device names, see Attaching an Amazon EBS Volume to an
// Instance . Any device names that aren't reserved for instance store
// volumes can be used for Amazon EBS volumes. For more information, see
// Amazon EC2 Instance Store in the Amazon Elastic Compute Cloud User Guide
// If a volume has an AWS Marketplace product code: The volume can only be
// attached as the root device of a stopped instance. You must be
// subscribed to the AWS Marketplace code that is on the volume. The
// configuration (instance type, operating system) of the instance must
// support that specific AWS Marketplace code. For example, you cannot take
// a volume from a Windows instance and attach it to a Linux instance. AWS
// Marketplace product codes are copied from the volume to the instance.
// For an overview of the AWS Marketplace, see
// https://aws.amazon.com/marketplace/help/200900000 . For more information
// about how to use the AWS Marketplace, see AWS Marketplace For more
// information about Amazon EBS volumes, see Attaching Amazon EBS Volumes
// in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) AttachVolume(req AttachVolumeRequest) (resp *VolumeAttachment, err error) {
	resp = &VolumeAttachment{}
	err = c.client.Do("AttachVolume", "POST", "/", req, resp)
	return
}

// AttachVpnGateway attaches a virtual private gateway to a For more
// information, see Adding a Hardware Virtual Private Gateway to Your in
// the Amazon Virtual Private Cloud User Guide
func (c *EC2) AttachVpnGateway(req AttachVpnGatewayRequest) (resp *AttachVpnGatewayResult, err error) {
	resp = &AttachVpnGatewayResult{}
	err = c.client.Do("AttachVpnGateway", "POST", "/", req, resp)
	return
}

// AuthorizeSecurityGroupEgress adds one or more egress rules to a security
// group for use with a Specifically, this action permits instances to send
// traffic to one or more destination IP address ranges, or to one or more
// destination security groups for the same You can have up to 50 rules per
// security group (covering both ingress and egress rules). A security
// group is for use with instances either in the EC2-Classic platform or in
// a specific This action doesn't apply to security groups for use in
// EC2-Classic. For more information, see Security Groups for Your in the
// Amazon Virtual Private Cloud User Guide Each rule consists of the
// protocol (for example, plus either a range or a source group. For the
// TCP and UDP protocols, you must also specify the destination port or
// port range. For the protocol, you must also specify the type and code.
// You can use -1 for the type or code to mean all types or all codes. Rule
// changes are propagated to affected instances as quickly as possible.
// However, a small delay might occur.
func (c *EC2) AuthorizeSecurityGroupEgress(req AuthorizeSecurityGroupEgressRequest) (err error) {
	// NRE
	err = c.client.Do("AuthorizeSecurityGroupEgress", "POST", "/", req, nil)
	return
}

// AuthorizeSecurityGroupIngress adds one or more ingress rules to a
// security group. EC2-Classic: You can have up to 100 rules per group.
// EC2-VPC: You can have up to 50 rules per group (covering both ingress
// and egress rules). Rule changes are propagated to instances within the
// security group as quickly as possible. However, a small delay might
// occur. [EC2-Classic] This action gives one or more IP address ranges
// permission to access a security group in your account, or gives one or
// more security groups (called the source groups ) permission to access a
// security group for your account. A source group can be for your own AWS
// account, or another. [EC2-VPC] This action gives one or more IP address
// ranges permission to access a security group in your or gives one or
// more other security groups (called the source groups ) permission to
// access a security group for your The security groups must all be for the
// same
func (c *EC2) AuthorizeSecurityGroupIngress(req AuthorizeSecurityGroupIngressRequest) (err error) {
	// NRE
	err = c.client.Do("AuthorizeSecurityGroupIngress", "POST", "/", req, nil)
	return
}

// BundleInstance bundles an Amazon instance store-backed Windows instance.
// During bundling, only the root device volume is bundled. Data on other
// instance store volumes is not preserved. This procedure is not
// applicable for Linux/Unix instances or Windows instances that are backed
// by Amazon For more information, see Creating an Instance Store-Backed
// Windows
func (c *EC2) BundleInstance(req BundleInstanceRequest) (resp *BundleInstanceResult, err error) {
	resp = &BundleInstanceResult{}
	err = c.client.Do("BundleInstance", "POST", "/", req, resp)
	return
}

// CancelBundleTask cancels a bundling operation for an instance
// store-backed Windows instance.
func (c *EC2) CancelBundleTask(req CancelBundleTaskRequest) (resp *CancelBundleTaskResult, err error) {
	resp = &CancelBundleTaskResult{}
	err = c.client.Do("CancelBundleTask", "POST", "/", req, resp)
	return
}

// CancelConversionTask cancels an active conversion task. The task can be
// the import of an instance or volume. The action removes all artifacts of
// the conversion, including a partially uploaded volume or instance. If
// the conversion is complete or is in the process of transferring the
// final disk image, the command fails and returns an exception. For more
// information, see Using the Command Line Tools to Import Your Virtual
// Machine to Amazon EC2 in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) CancelConversionTask(req CancelConversionRequest) (err error) {
	// NRE
	err = c.client.Do("CancelConversionTask", "POST", "/", req, nil)
	return
}

// CancelExportTask cancels an active export task. The request removes all
// artifacts of the export, including any partially-created Amazon S3
// objects. If the export task is complete or is in the process of
// transferring the final disk image, the command fails and returns an
// error.
func (c *EC2) CancelExportTask(req CancelExportTaskRequest) (err error) {
	// NRE
	err = c.client.Do("CancelExportTask", "POST", "/", req, nil)
	return
}

// CancelReservedInstancesListing cancels the specified Reserved Instance
// listing in the Reserved Instance Marketplace. For more information, see
// Reserved Instance Marketplace in the Amazon Elastic Compute Cloud User
// Guide
func (c *EC2) CancelReservedInstancesListing(req CancelReservedInstancesListingRequest) (resp *CancelReservedInstancesListingResult, err error) {
	resp = &CancelReservedInstancesListingResult{}
	err = c.client.Do("CancelReservedInstancesListing", "POST", "/", req, resp)
	return
}

// CancelSpotInstanceRequests cancels one or more Spot Instance requests.
// Spot Instances are instances that Amazon EC2 starts on your behalf when
// the maximum price that you specify exceeds the current Spot Price.
// Amazon EC2 periodically sets the Spot Price based on available Spot
// Instance capacity and current Spot Instance requests. For more
// information about Spot Instances, see Spot Instances in the Amazon
// Elastic Compute Cloud User Guide Canceling a Spot Instance request does
// not terminate running Spot Instances associated with the request.
func (c *EC2) CancelSpotInstanceRequests(req CancelSpotInstanceRequestsRequest) (resp *CancelSpotInstanceRequestsResult, err error) {
	resp = &CancelSpotInstanceRequestsResult{}
	err = c.client.Do("CancelSpotInstanceRequests", "POST", "/", req, resp)
	return
}

// ConfirmProductInstance determines whether a product code is associated
// with an instance. This action can only be used by the owner of the
// product code. It is useful when a product code owner needs to verify
// whether another user's instance is eligible for support.
func (c *EC2) ConfirmProductInstance(req ConfirmProductInstanceRequest) (resp *ConfirmProductInstanceResult, err error) {
	resp = &ConfirmProductInstanceResult{}
	err = c.client.Do("ConfirmProductInstance", "POST", "/", req, resp)
	return
}

// CopyImage initiates the copy of an AMI from the specified source region
// to the region in which the request was made. You specify the destination
// region by using its endpoint when making the request. AMIs that use
// encrypted Amazon EBS snapshots cannot be copied with this method. For
// more information, see Copying AMIs in the Amazon Elastic Compute Cloud
// User Guide
func (c *EC2) CopyImage(req CopyImageRequest) (resp *CopyImageResult, err error) {
	resp = &CopyImageResult{}
	err = c.client.Do("CopyImage", "POST", "/", req, resp)
	return
}

// CopySnapshot copies a point-in-time snapshot of an Amazon EBS volume and
// stores it in Amazon S3. You can copy the snapshot within the same region
// or from one region to another. You can use the snapshot to create Amazon
// EBS volumes or Amazon Machine Images (AMIs). The snapshot is copied to
// the regional endpoint that you send the request to. Copies of encrypted
// Amazon EBS snapshots remain encrypted. Copies of unencrypted snapshots
// remain unencrypted. For more information, see Copying an Amazon EBS
// Snapshot in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) CopySnapshot(req CopySnapshotRequest) (resp *CopySnapshotResult, err error) {
	resp = &CopySnapshotResult{}
	err = c.client.Do("CopySnapshot", "POST", "/", req, resp)
	return
}

// CreateCustomerGateway provides information to AWS about your VPN
// customer gateway device. The customer gateway is the appliance at your
// end of the VPN connection. (The device on the AWS side of the VPN
// connection is the virtual private gateway.) You must provide the
// Internet-routable IP address of the customer gateway's external
// interface. The IP address must be static and can't be behind a device
// performing network address translation For devices that use Border
// Gateway Protocol you can also provide the device's BGP Autonomous System
// Number You can use an existing ASN assigned to your network. If you
// don't have an ASN already, you can use a private ASN (in the 64512 -
// 65534 range). Amazon EC2 supports all 2-byte ASN numbers in the range of
// 1 - 65534, with the exception of 7224, which is reserved in the
// us-east-1 region, and 9059, which is reserved in the eu-west-1 region.
// For more information about VPN customer gateways, see Adding a Hardware
// Virtual Private Gateway to Your in the Amazon Virtual Private Cloud User
// Guide
func (c *EC2) CreateCustomerGateway(req CreateCustomerGatewayRequest) (resp *CreateCustomerGatewayResult, err error) {
	resp = &CreateCustomerGatewayResult{}
	err = c.client.Do("CreateCustomerGateway", "POST", "/", req, resp)
	return
}

// CreateDhcpOptions creates a set of options for your After creating the
// set, you must associate it with the causing all existing and new
// instances that you launch in the VPC to use this set of options. The
// following are the individual options you can specify. For more
// information about the options, see RFC 2132 domain-name-servers - The IP
// addresses of up to four domain name servers, or AmazonProvidedDNS . The
// default option set specifies AmazonProvidedDNS . If specifying more than
// one domain name server, specify the IP addresses in a single parameter,
// separated by commas. domain-name - If you're using AmazonProvidedDNS in
// us-east-1 , specify ec2.internal . If you're using AmazonProvidedDNS in
// another region, specify region.compute.internal (for example,
// ap-northeast-1.compute.internal ). Otherwise, specify a domain name (for
// example, MyCompany.com ). If specifying more than one domain name,
// separate them with spaces. ntp-servers - The IP addresses of up to four
// Network Time Protocol servers. netbios-name-servers - The IP addresses
// of up to four NetBIOS name servers. netbios-node-type - The NetBIOS node
// type (1, 2, 4, or 8). We recommend that you specify 2 (broadcast and
// multicast are not currently supported). For more information about these
// node types, see RFC 2132 . Your VPC automatically starts out with a set
// of options that includes only a DNS server that we provide
// (AmazonProvidedDNS). If you create a set of options, and if your VPC has
// an Internet gateway, make sure to set the domain-name-servers option
// either to AmazonProvidedDNS or to a domain name server of your choice.
// For more information about options, see Options Sets in the Amazon
// Virtual Private Cloud User Guide
func (c *EC2) CreateDhcpOptions(req CreateDhcpOptionsRequest) (resp *CreateDhcpOptionsResult, err error) {
	resp = &CreateDhcpOptionsResult{}
	err = c.client.Do("CreateDhcpOptions", "POST", "/", req, resp)
	return
}

// CreateImage creates an Amazon EBS-backed AMI from an Amazon EBS-backed
// instance that is either running or stopped. If you customized your
// instance with instance store volumes or EBS volumes in addition to the
// root device volume, the new AMI contains block device mapping
// information for those volumes. When you launch an instance from this new
// the instance automatically launches with those additional volumes. For
// more information, see Creating Amazon EBS-Backed Linux AMIs in the
// Amazon Elastic Compute Cloud User Guide
func (c *EC2) CreateImage(req CreateImageRequest) (resp *CreateImageResult, err error) {
	resp = &CreateImageResult{}
	err = c.client.Do("CreateImage", "POST", "/", req, resp)
	return
}

// CreateInstanceExportTask exports a running or stopped instance to an
// Amazon S3 bucket. For information about the supported operating systems,
// image formats, and known limitations for the types of instances you can
// export, see Exporting EC2 Instances in the Amazon Elastic Compute Cloud
// User Guide
func (c *EC2) CreateInstanceExportTask(req CreateInstanceExportTaskRequest) (resp *CreateInstanceExportTaskResult, err error) {
	resp = &CreateInstanceExportTaskResult{}
	err = c.client.Do("CreateInstanceExportTask", "POST", "/", req, resp)
	return
}

// CreateInternetGateway creates an Internet gateway for use with a After
// creating the Internet gateway, you attach it to a VPC using
// AttachInternetGateway For more information about your VPC and Internet
// gateway, see the Amazon Virtual Private Cloud User Guide
func (c *EC2) CreateInternetGateway(req CreateInternetGatewayRequest) (resp *CreateInternetGatewayResult, err error) {
	resp = &CreateInternetGatewayResult{}
	err = c.client.Do("CreateInternetGateway", "POST", "/", req, resp)
	return
}

// CreateKeyPair creates a 2048-bit RSA key pair with the specified name.
// Amazon EC2 stores the public key and displays the private key for you to
// save to a file. The private key is returned as an unencrypted PEM
// encoded PKCS#8 private key. If a key with the specified name already
// exists, Amazon EC2 returns an error. You can have up to five thousand
// key pairs per region. The key pair returned to you is available only in
// the region in which you create it. To create a key pair that is
// available in all regions, use ImportKeyPair For more information about
// key pairs, see Key Pairs in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) CreateKeyPair(req CreateKeyPairRequest) (resp *KeyPair, err error) {
	resp = &KeyPair{}
	err = c.client.Do("CreateKeyPair", "POST", "/", req, resp)
	return
}

// CreateNetworkAcl creates a network ACL in a Network ACLs provide an
// optional layer of security (in addition to security groups) for the
// instances in your For more information about network ACLs, see Network
// ACLs in the Amazon Virtual Private Cloud User Guide
func (c *EC2) CreateNetworkAcl(req CreateNetworkAclRequest) (resp *CreateNetworkAclResult, err error) {
	resp = &CreateNetworkAclResult{}
	err = c.client.Do("CreateNetworkAcl", "POST", "/", req, resp)
	return
}

// CreateNetworkAclEntry creates an entry (a rule) in a network ACL with
// the specified rule number. Each network ACL has a set of numbered
// ingress rules and a separate set of numbered egress rules. When
// determining whether a packet should be allowed in or out of a subnet
// associated with the we process the entries in the ACL according to the
// rule numbers, in ascending order. Each network ACL has a set of ingress
// rules and a separate set of egress rules. We recommend that you leave
// room between the rule numbers (for example, 100, 110, 120, and not
// number them one right after the other (for example, 101, 102, 103, This
// makes it easier to add a rule between existing ones without having to
// renumber the rules. After you add an entry, you can't modify it; you
// must either replace it, or create an entry and delete the old one. For
// more information about network ACLs, see Network ACLs in the Amazon
// Virtual Private Cloud User Guide
func (c *EC2) CreateNetworkAclEntry(req CreateNetworkAclEntryRequest) (err error) {
	// NRE
	err = c.client.Do("CreateNetworkAclEntry", "POST", "/", req, nil)
	return
}

// CreateNetworkInterface creates a network interface in the specified
// subnet. For more information about network interfaces, see Elastic
// Network Interfaces in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) CreateNetworkInterface(req CreateNetworkInterfaceRequest) (resp *CreateNetworkInterfaceResult, err error) {
	resp = &CreateNetworkInterfaceResult{}
	err = c.client.Do("CreateNetworkInterface", "POST", "/", req, resp)
	return
}

// CreatePlacementGroup creates a placement group that you launch cluster
// instances into. You must give the group a name that's unique within the
// scope of your account. For more information about placement groups and
// cluster instances, see Cluster Instances in the Amazon Elastic Compute
// Cloud User Guide
func (c *EC2) CreatePlacementGroup(req CreatePlacementGroupRequest) (err error) {
	// NRE
	err = c.client.Do("CreatePlacementGroup", "POST", "/", req, nil)
	return
}

// CreateReservedInstancesListing creates a listing for Amazon EC2 Reserved
// Instances to be sold in the Reserved Instance Marketplace. You can
// submit one Reserved Instance listing at a time. To get a list of your
// Reserved Instances, you can use the DescribeReservedInstances operation.
// The Reserved Instance Marketplace matches sellers who want to resell
// Reserved Instance capacity that they no longer need with buyers who want
// to purchase additional capacity. Reserved Instances bought and sold
// through the Reserved Instance Marketplace work like any other Reserved
// Instances. To sell your Reserved Instances, you must first register as a
// Seller in the Reserved Instance Marketplace. After completing the
// registration process, you can create a Reserved Instance Marketplace
// listing of some or all of your Reserved Instances, and specify the
// upfront price to receive for them. Your Reserved Instance listings then
// become available for purchase. To view the details of your Reserved
// Instance listing, you can use the DescribeReservedInstancesListings
// operation. For more information, see Reserved Instance Marketplace in
// the Amazon Elastic Compute Cloud User Guide
func (c *EC2) CreateReservedInstancesListing(req CreateReservedInstancesListingRequest) (resp *CreateReservedInstancesListingResult, err error) {
	resp = &CreateReservedInstancesListingResult{}
	err = c.client.Do("CreateReservedInstancesListing", "POST", "/", req, resp)
	return
}

// CreateRoute creates a route in a route table within a You must specify
// one of the following targets: Internet gateway or virtual private
// gateway, NAT instance, VPC peering connection, or network interface.
// When determining how to route traffic, we use the route with the most
// specific match. For example, let's say the traffic is destined for
// 192.0.2.3 , and the route table includes the following two routes:
// 192.0.2.0/24 (goes to some target 192.0.2.0/28 (goes to some target Both
// routes apply to the traffic destined for 192.0.2.3 . However, the second
// route in the list covers a smaller number of IP addresses and is
// therefore more specific, so we use that route to determine where to
// target the traffic. For more information about route tables, see Route
// Tables in the Amazon Virtual Private Cloud User Guide
func (c *EC2) CreateRoute(req CreateRouteRequest) (err error) {
	// NRE
	err = c.client.Do("CreateRoute", "POST", "/", req, nil)
	return
}

// CreateRouteTable creates a route table for the specified After you
// create a route table, you can add routes and associate the table with a
// subnet. For more information about route tables, see Route Tables in the
// Amazon Virtual Private Cloud User Guide
func (c *EC2) CreateRouteTable(req CreateRouteTableRequest) (resp *CreateRouteTableResult, err error) {
	resp = &CreateRouteTableResult{}
	err = c.client.Do("CreateRouteTable", "POST", "/", req, resp)
	return
}

// CreateSecurityGroup creates a security group. A security group is for
// use with instances either in the EC2-Classic platform or in a specific
// For more information, see Amazon EC2 Security Groups in the Amazon
// Elastic Compute Cloud User Guide and Security Groups for Your in the
// Amazon Virtual Private Cloud User Guide EC2-Classic: You can have up to
// 500 security groups. EC2-VPC: You can create up to 100 security groups
// per When you create a security group, you specify a friendly name of
// your choice. You can have a security group for use in EC2-Classic with
// the same name as a security group for use in a However, you can't have
// two security groups for use in EC2-Classic with the same name or two
// security groups for use in a VPC with the same name. You have a default
// security group for use in EC2-Classic and a default security group for
// use in your If you don't specify a security group when you launch an
// instance, the instance is launched into the appropriate default security
// group. A default security group includes a default rule that grants
// instances unrestricted network access to each other. You can add or
// remove rules from your security groups using
// AuthorizeSecurityGroupIngress , AuthorizeSecurityGroupEgress ,
// RevokeSecurityGroupIngress , and RevokeSecurityGroupEgress
func (c *EC2) CreateSecurityGroup(req CreateSecurityGroupRequest) (resp *CreateSecurityGroupResult, err error) {
	resp = &CreateSecurityGroupResult{}
	err = c.client.Do("CreateSecurityGroup", "POST", "/", req, resp)
	return
}

// CreateSnapshot creates a snapshot of an Amazon EBS volume and stores it
// in Amazon S3. You can use snapshots for backups, to make copies of
// Amazon EBS volumes, and to save data before shutting down an instance.
// When a snapshot is created, any AWS Marketplace product codes that are
// associated with the source volume are propagated to the snapshot. You
// can take a snapshot of an attached volume that is in use. However,
// snapshots only capture data that has been written to your Amazon EBS
// volume at the time the snapshot command is issued; this may exclude any
// data that has been cached by any applications or the operating system.
// If you can pause any file systems on the volume long enough to take a
// snapshot, your snapshot should be complete. However, if you cannot pause
// all file writes to the volume, you should unmount the volume from within
// the instance, issue the snapshot command, and then remount the volume to
// ensure a consistent and complete snapshot. You may remount and use your
// volume while the snapshot status is pending To create a snapshot for
// Amazon EBS volumes that serve as root devices, you should stop the
// instance before taking the snapshot. Snapshots that are taken from
// encrypted volumes are automatically encrypted. Volumes that are created
// from encrypted snapshots are also automatically encrypted. Your
// encrypted volumes and any associated snapshots always remain protected.
// For more information, see Amazon Elastic Block Store and Amazon EBS
// Encryption in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) CreateSnapshot(req CreateSnapshotRequest) (resp *Snapshot, err error) {
	resp = &Snapshot{}
	err = c.client.Do("CreateSnapshot", "POST", "/", req, resp)
	return
}

// CreateSpotDatafeedSubscription creates a datafeed for Spot Instances,
// enabling you to view Spot Instance usage logs. You can create one data
// feed per AWS account. For more information, see Spot Instances in the
// Amazon Elastic Compute Cloud User Guide
func (c *EC2) CreateSpotDatafeedSubscription(req CreateSpotDatafeedSubscriptionRequest) (resp *CreateSpotDatafeedSubscriptionResult, err error) {
	resp = &CreateSpotDatafeedSubscriptionResult{}
	err = c.client.Do("CreateSpotDatafeedSubscription", "POST", "/", req, resp)
	return
}

// CreateSubnet creates a subnet in an existing When you create each
// subnet, you provide the VPC ID and the block you want for the subnet.
// After you create a subnet, you can't change its block. The subnet's
// block can be the same as the VPC's block (assuming you want only a
// single subnet in the or a subset of the VPC's block. If you create more
// than one subnet in a the subnets' blocks must not overlap. The smallest
// subnet (and you can create uses a /28 netmask (16 IP addresses), and the
// largest uses a /16 netmask (65,536 IP addresses). AWS reserves both the
// first four and the last IP address in each subnet's block. They're not
// available for use. If you add more than one subnet to a they're set up
// in a star topology with a logical router in the middle. If you launch an
// instance in a VPC using an Amazon EBS-backed the IP address doesn't
// change if you stop and restart the instance (unlike a similar instance
// launched outside a which gets a new IP address when restarted). It's
// therefore possible to have a subnet with no running instances (they're
// all stopped), but no remaining IP addresses available. For more
// information about subnets, see Your VPC and Subnets in the Amazon
// Virtual Private Cloud User Guide
func (c *EC2) CreateSubnet(req CreateSubnetRequest) (resp *CreateSubnetResult, err error) {
	resp = &CreateSubnetResult{}
	err = c.client.Do("CreateSubnet", "POST", "/", req, resp)
	return
}

// CreateTags adds or overwrites one or more tags for the specified EC2
// resource or resources. Each resource can have a maximum of 10 tags. Each
// tag consists of a key and optional value. Tag keys must be unique per
// resource. For more information about tags, see Tagging Your Resources in
// the Amazon Elastic Compute Cloud User Guide
func (c *EC2) CreateTags(req CreateTagsRequest) (err error) {
	// NRE
	err = c.client.Do("CreateTags", "POST", "/", req, nil)
	return
}

// CreateVolume creates an Amazon EBS volume that can be attached to an
// instance in the same Availability Zone. The volume is created in the
// specified region. You can create a new empty volume or restore a volume
// from an Amazon EBS snapshot. Any AWS Marketplace product codes from the
// snapshot are propagated to the volume. You can create encrypted volumes
// with the Encrypted parameter. Encrypted volumes may only be attached to
// instances that support Amazon EBS encryption. Volumes that are created
// from encrypted snapshots are also automatically encrypted. For more
// information, see Amazon EBS Encryption in the Amazon Elastic Compute
// Cloud User Guide For more information, see Creating or Restoring an
// Amazon EBS Volume in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) CreateVolume(req CreateVolumeRequest) (resp *Volume, err error) {
	resp = &Volume{}
	err = c.client.Do("CreateVolume", "POST", "/", req, resp)
	return
}

// CreateVpc creates a VPC with the specified block. The smallest VPC you
// can create uses a /28 netmask (16 IP addresses), and the largest uses a
// /16 netmask (65,536 IP addresses). To help you decide how big to make
// your see Your VPC and Subnets in the Amazon Virtual Private Cloud User
// Guide By default, each instance you launch in the VPC has the default
// options, which includes only a default DNS server that we provide
// (AmazonProvidedDNS). For more information about options, see Options
// Sets in the Amazon Virtual Private Cloud User Guide
func (c *EC2) CreateVpc(req CreateVpcRequest) (resp *CreateVpcResult, err error) {
	resp = &CreateVpcResult{}
	err = c.client.Do("CreateVpc", "POST", "/", req, resp)
	return
}

// CreateVpcPeeringConnection requests a VPC peering connection between two
// VPCs: a requester VPC that you own and a peer VPC with which to create
// the connection. The peer VPC can belong to another AWS account. The
// requester VPC and peer VPC cannot have overlapping blocks. The owner of
// the peer VPC must accept the peering request to activate the peering
// connection. The VPC peering connection request expires after 7 days,
// after which it cannot be accepted or rejected. A
// CreateVpcPeeringConnection request between VPCs with overlapping blocks
// results in the VPC peering connection having a status of failed
func (c *EC2) CreateVpcPeeringConnection(req CreateVpcPeeringConnectionRequest) (resp *CreateVpcPeeringConnectionResult, err error) {
	resp = &CreateVpcPeeringConnectionResult{}
	err = c.client.Do("CreateVpcPeeringConnection", "POST", "/", req, resp)
	return
}

// CreateVpnConnection creates a VPN connection between an existing virtual
// private gateway and a VPN customer gateway. The only supported
// connection type is ipsec.1 The response includes information that you
// need to give to your network administrator to configure your customer
// gateway. We strongly recommend that you use when calling this operation
// because the response contains sensitive cryptographic information for
// configuring your customer gateway. If you decide to shut down your VPN
// connection for any reason and later create a new VPN connection, you
// must reconfigure your customer gateway with the new information returned
// from this call. For more information about VPN connections, see Adding a
// Hardware Virtual Private Gateway to Your in the Amazon Virtual Private
// Cloud User Guide
func (c *EC2) CreateVpnConnection(req CreateVpnConnectionRequest) (resp *CreateVpnConnectionResult, err error) {
	resp = &CreateVpnConnectionResult{}
	err = c.client.Do("CreateVpnConnection", "POST", "/", req, resp)
	return
}

// CreateVpnConnectionRoute creates a static route associated with a VPN
// connection between an existing virtual private gateway and a VPN
// customer gateway. The static route allows traffic to be routed from the
// virtual private gateway to the VPN customer gateway. For more
// information about VPN connections, see Adding a Hardware Virtual Private
// Gateway to Your in the Amazon Virtual Private Cloud User Guide
func (c *EC2) CreateVpnConnectionRoute(req CreateVpnConnectionRouteRequest) (err error) {
	// NRE
	err = c.client.Do("CreateVpnConnectionRoute", "POST", "/", req, nil)
	return
}

// CreateVpnGateway creates a virtual private gateway. A virtual private
// gateway is the endpoint on the VPC side of your VPN connection. You can
// create a virtual private gateway before creating the VPC itself. For
// more information about virtual private gateways, see Adding a Hardware
// Virtual Private Gateway to Your in the Amazon Virtual Private Cloud User
// Guide
func (c *EC2) CreateVpnGateway(req CreateVpnGatewayRequest) (resp *CreateVpnGatewayResult, err error) {
	resp = &CreateVpnGatewayResult{}
	err = c.client.Do("CreateVpnGateway", "POST", "/", req, resp)
	return
}

// DeleteCustomerGateway deletes the specified customer gateway. You must
// delete the VPN connection before you can delete the customer gateway.
func (c *EC2) DeleteCustomerGateway(req DeleteCustomerGatewayRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteCustomerGateway", "POST", "/", req, nil)
	return
}

// DeleteDhcpOptions deletes the specified set of options. You must
// disassociate the set of options before you can delete it. You can
// disassociate the set of options by associating either a new set of
// options or the default set of options with the
func (c *EC2) DeleteDhcpOptions(req DeleteDhcpOptionsRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteDhcpOptions", "POST", "/", req, nil)
	return
}

// DeleteInternetGateway deletes the specified Internet gateway. You must
// detach the Internet gateway from the VPC before you can delete it.
func (c *EC2) DeleteInternetGateway(req DeleteInternetGatewayRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteInternetGateway", "POST", "/", req, nil)
	return
}

// DeleteKeyPair deletes the specified key pair, by removing the public key
// from Amazon EC2.
func (c *EC2) DeleteKeyPair(req DeleteKeyPairRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteKeyPair", "POST", "/", req, nil)
	return
}

// DeleteNetworkAcl deletes the specified network You can't delete the ACL
// if it's associated with any subnets. You can't delete the default
// network
func (c *EC2) DeleteNetworkAcl(req DeleteNetworkAclRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteNetworkAcl", "POST", "/", req, nil)
	return
}

// DeleteNetworkAclEntry deletes the specified ingress or egress entry
// (rule) from the specified network
func (c *EC2) DeleteNetworkAclEntry(req DeleteNetworkAclEntryRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteNetworkAclEntry", "POST", "/", req, nil)
	return
}

// DeleteNetworkInterface deletes the specified network interface. You must
// detach the network interface before you can delete it.
func (c *EC2) DeleteNetworkInterface(req DeleteNetworkInterfaceRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteNetworkInterface", "POST", "/", req, nil)
	return
}

// DeletePlacementGroup deletes the specified placement group. You must
// terminate all instances in the placement group before you can delete the
// placement group. For more information about placement groups and cluster
// instances, see Cluster Instances in the Amazon Elastic Compute Cloud
// User Guide
func (c *EC2) DeletePlacementGroup(req DeletePlacementGroupRequest) (err error) {
	// NRE
	err = c.client.Do("DeletePlacementGroup", "POST", "/", req, nil)
	return
}

// DeleteRoute deletes the specified route from the specified route table.
func (c *EC2) DeleteRoute(req DeleteRouteRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteRoute", "POST", "/", req, nil)
	return
}

// DeleteRouteTable deletes the specified route table. You must
// disassociate the route table from any subnets before you can delete it.
// You can't delete the main route table.
func (c *EC2) DeleteRouteTable(req DeleteRouteTableRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteRouteTable", "POST", "/", req, nil)
	return
}

// DeleteSecurityGroup deletes a security group. If you attempt to delete a
// security group that is associated with an instance, or is referenced by
// another security group, the operation fails with InvalidGroup.InUse in
// EC2-Classic or DependencyViolation in EC2-VPC.
func (c *EC2) DeleteSecurityGroup(req DeleteSecurityGroupRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteSecurityGroup", "POST", "/", req, nil)
	return
}

// DeleteSnapshot deletes the specified snapshot. When you make periodic
// snapshots of a volume, the snapshots are incremental, and only the
// blocks on the device that have changed since your last snapshot are
// saved in the new snapshot. When you delete a snapshot, only the data not
// needed for any other snapshot is removed. So regardless of which prior
// snapshots have been deleted, all active snapshots will have access to
// all the information needed to restore the volume. You cannot delete a
// snapshot of the root device of an Amazon EBS volume used by a registered
// You must first de-register the AMI before you can delete the snapshot.
// For more information, see Deleting an Amazon EBS Snapshot in the Amazon
// Elastic Compute Cloud User Guide
func (c *EC2) DeleteSnapshot(req DeleteSnapshotRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteSnapshot", "POST", "/", req, nil)
	return
}

// DeleteSpotDatafeedSubscription deletes the datafeed for Spot Instances.
// For more information, see Spot Instances in the Amazon Elastic Compute
// Cloud User Guide
func (c *EC2) DeleteSpotDatafeedSubscription(req DeleteSpotDatafeedSubscriptionRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteSpotDatafeedSubscription", "POST", "/", req, nil)
	return
}

// DeleteSubnet deletes the specified subnet. You must terminate all
// running instances in the subnet before you can delete the subnet.
func (c *EC2) DeleteSubnet(req DeleteSubnetRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteSubnet", "POST", "/", req, nil)
	return
}

// DeleteTags deletes the specified set of tags from the specified set of
// resources. This call is designed to follow a DescribeTags request. For
// more information about tags, see Tagging Your Resources in the Amazon
// Elastic Compute Cloud User Guide
func (c *EC2) DeleteTags(req DeleteTagsRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteTags", "POST", "/", req, nil)
	return
}

// DeleteVolume deletes the specified Amazon EBS volume. The volume must be
// in the available state (not attached to an instance). The volume may
// remain in the deleting state for several minutes. For more information,
// see Deleting an Amazon EBS Volume in the Amazon Elastic Compute Cloud
// User Guide
func (c *EC2) DeleteVolume(req DeleteVolumeRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteVolume", "POST", "/", req, nil)
	return
}

// DeleteVpc deletes the specified You must detach or delete all gateways
// and resources that are associated with the VPC before you can delete it.
// For example, you must terminate all instances running in the delete all
// security groups associated with the VPC (except the default one), delete
// all route tables associated with the VPC (except the default one), and
// so on.
func (c *EC2) DeleteVpc(req DeleteVpcRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteVpc", "POST", "/", req, nil)
	return
}

// DeleteVpcPeeringConnection deletes a VPC peering connection. Either the
// owner of the requester VPC or the owner of the peer VPC can delete the
// VPC peering connection if it's in the active state. The owner of the
// requester VPC can delete a VPC peering connection in the
// pending-acceptance state.
func (c *EC2) DeleteVpcPeeringConnection(req DeleteVpcPeeringConnectionRequest) (resp *DeleteVpcPeeringConnectionResult, err error) {
	resp = &DeleteVpcPeeringConnectionResult{}
	err = c.client.Do("DeleteVpcPeeringConnection", "POST", "/", req, resp)
	return
}

// DeleteVpnConnection deletes the specified VPN connection. If you're
// deleting the VPC and its associated components, we recommend that you
// detach the virtual private gateway from the VPC and delete the VPC
// before deleting the VPN connection. If you believe that the tunnel
// credentials for your VPN connection have been compromised, you can
// delete the VPN connection and create a new one that has new keys,
// without needing to delete the VPC or virtual private gateway. If you
// create a new VPN connection, you must reconfigure the customer gateway
// using the new configuration information returned with the new VPN
// connection
func (c *EC2) DeleteVpnConnection(req DeleteVpnConnectionRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteVpnConnection", "POST", "/", req, nil)
	return
}

// DeleteVpnConnectionRoute deletes the specified static route associated
// with a VPN connection between an existing virtual private gateway and a
// VPN customer gateway. The static route allows traffic to be routed from
// the virtual private gateway to the VPN customer gateway.
func (c *EC2) DeleteVpnConnectionRoute(req DeleteVpnConnectionRouteRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteVpnConnectionRoute", "POST", "/", req, nil)
	return
}

// DeleteVpnGateway deletes the specified virtual private gateway. We
// recommend that before you delete a virtual private gateway, you detach
// it from the VPC and delete the VPN connection. Note that you don't need
// to delete the virtual private gateway if you plan to delete and recreate
// the VPN connection between your VPC and your network.
func (c *EC2) DeleteVpnGateway(req DeleteVpnGatewayRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteVpnGateway", "POST", "/", req, nil)
	return
}

// DeregisterImage deregisters the specified After you deregister an it
// can't be used to launch new instances. This command does not delete the
func (c *EC2) DeregisterImage(req DeregisterImageRequest) (err error) {
	// NRE
	err = c.client.Do("DeregisterImage", "POST", "/", req, nil)
	return
}

// DescribeAccountAttributes describes the specified attribute of your AWS
// account.
func (c *EC2) DescribeAccountAttributes(req DescribeAccountAttributesRequest) (resp *DescribeAccountAttributesResult, err error) {
	resp = &DescribeAccountAttributesResult{}
	err = c.client.Do("DescribeAccountAttributes", "POST", "/", req, resp)
	return
}

// DescribeAddresses describes one or more of your Elastic IP addresses. An
// Elastic IP address is for use in either the EC2-Classic platform or in a
// For more information, see Elastic IP Addresses in the Amazon Elastic
// Compute Cloud User Guide
func (c *EC2) DescribeAddresses(req DescribeAddressesRequest) (resp *DescribeAddressesResult, err error) {
	resp = &DescribeAddressesResult{}
	err = c.client.Do("DescribeAddresses", "POST", "/", req, resp)
	return
}

// DescribeAvailabilityZones describes one or more of the Availability
// Zones that are available to you. The results include zones only for the
// region you're currently using. If there is an event impacting an
// Availability Zone, you can use this request to view the state and any
// provided message for that Availability Zone. For more information, see
// Regions and Availability Zones in the Amazon Elastic Compute Cloud User
// Guide
func (c *EC2) DescribeAvailabilityZones(req DescribeAvailabilityZonesRequest) (resp *DescribeAvailabilityZonesResult, err error) {
	resp = &DescribeAvailabilityZonesResult{}
	err = c.client.Do("DescribeAvailabilityZones", "POST", "/", req, resp)
	return
}

// DescribeBundleTasks describes one or more of your bundling tasks.
// Completed bundle tasks are listed for only a limited time. If your
// bundle task is no longer in the list, you can still register an AMI from
// it. Just use RegisterImage with the Amazon S3 bucket name and image
// manifest name you provided to the bundle task.
func (c *EC2) DescribeBundleTasks(req DescribeBundleTasksRequest) (resp *DescribeBundleTasksResult, err error) {
	resp = &DescribeBundleTasksResult{}
	err = c.client.Do("DescribeBundleTasks", "POST", "/", req, resp)
	return
}

// DescribeConversionTasks describes one or more of your conversion tasks.
// For more information, see Using the Command Line Tools to Import Your
// Virtual Machine to Amazon EC2 in the Amazon Elastic Compute Cloud User
// Guide
func (c *EC2) DescribeConversionTasks(req DescribeConversionTasksRequest) (resp *DescribeConversionTasksResult, err error) {
	resp = &DescribeConversionTasksResult{}
	err = c.client.Do("DescribeConversionTasks", "POST", "/", req, resp)
	return
}

// DescribeCustomerGateways describes one or more of your VPN customer
// gateways. For more information about VPN customer gateways, see Adding a
// Hardware Virtual Private Gateway to Your in the Amazon Virtual Private
// Cloud User Guide
func (c *EC2) DescribeCustomerGateways(req DescribeCustomerGatewaysRequest) (resp *DescribeCustomerGatewaysResult, err error) {
	resp = &DescribeCustomerGatewaysResult{}
	err = c.client.Do("DescribeCustomerGateways", "POST", "/", req, resp)
	return
}

// DescribeDhcpOptions describes one or more of your options sets. For more
// information about options sets, see Options Sets in the Amazon Virtual
// Private Cloud User Guide
func (c *EC2) DescribeDhcpOptions(req DescribeDhcpOptionsRequest) (resp *DescribeDhcpOptionsResult, err error) {
	resp = &DescribeDhcpOptionsResult{}
	err = c.client.Do("DescribeDhcpOptions", "POST", "/", req, resp)
	return
}

// DescribeExportTasks is undocumented.
func (c *EC2) DescribeExportTasks(req DescribeExportTasksRequest) (resp *DescribeExportTasksResult, err error) {
	resp = &DescribeExportTasksResult{}
	err = c.client.Do("DescribeExportTasks", "POST", "/", req, resp)
	return
}

// DescribeImageAttribute describes the specified attribute of the
// specified You can specify only one attribute at a time.
func (c *EC2) DescribeImageAttribute(req DescribeImageAttributeRequest) (resp *ImageAttribute, err error) {
	resp = &ImageAttribute{}
	err = c.client.Do("DescribeImageAttribute", "POST", "/", req, resp)
	return
}

// DescribeImages describes one or more of the images (AMIs, AKIs, and
// ARIs) available to you. Images available to you include public images,
// private images that you own, and private images owned by other AWS
// accounts but for which you have explicit launch permissions.
// Deregistered images are included in the returned results for an
// unspecified interval after deregistration.
func (c *EC2) DescribeImages(req DescribeImagesRequest) (resp *DescribeImagesResult, err error) {
	resp = &DescribeImagesResult{}
	err = c.client.Do("DescribeImages", "POST", "/", req, resp)
	return
}

// DescribeInstanceAttribute describes the specified attribute of the
// specified instance. You can specify only one attribute at a time. Valid
// attribute values are: instanceType | kernel | ramdisk | userData |
// disableApiTermination | instanceInitiatedShutdownBehavior |
// rootDeviceName | blockDeviceMapping | productCodes | sourceDestCheck |
// groupSet | ebsOptimized | sriovNetSupport
func (c *EC2) DescribeInstanceAttribute(req DescribeInstanceAttributeRequest) (resp *InstanceAttribute, err error) {
	resp = &InstanceAttribute{}
	err = c.client.Do("DescribeInstanceAttribute", "POST", "/", req, resp)
	return
}

// DescribeInstanceStatus describes the status of one or more instances,
// including any scheduled events. Instance status has two main components:
// System Status reports impaired functionality that stems from issues
// related to the systems that support an instance, such as such as
// hardware failures and network connectivity problems. This call reports
// such problems as impaired reachability. Instance Status reports impaired
// functionality that arises from problems internal to the instance. This
// call reports such problems as impaired reachability. Instance status
// provides information about four types of scheduled events for an
// instance that may require your attention: Scheduled Reboot: When Amazon
// EC2 determines that an instance must be rebooted, the instances status
// returns one of two event codes: system-reboot or instance-reboot .
// System reboot commonly occurs if certain maintenance or upgrade
// operations require a reboot of the underlying host that supports an
// instance. Instance reboot commonly occurs if the instance must be
// rebooted, rather than the underlying host. Rebooting events include a
// scheduled start and end time. System Maintenance: When Amazon EC2
// determines that an instance requires maintenance that requires power or
// network impact, the instance status is the event code system-maintenance
// . System maintenance is either power maintenance or network maintenance.
// For power maintenance, your instance will be unavailable for a brief
// period of time and then rebooted. For network maintenance, your instance
// will experience a brief loss of network connectivity. System maintenance
// events include a scheduled start and end time. You will also be notified
// by email if one of your instances is set for system maintenance. The
// email message indicates when your instance is scheduled for maintenance.
// Scheduled Retirement: When Amazon EC2 determines that an instance must
// be shut down, the instance status is the event code instance-retirement
// . Retirement commonly occurs when the underlying host is degraded and
// must be replaced. Retirement events include a scheduled start and end
// time. You will also be notified by email if one of your instances is set
// to retiring. The email message indicates when your instance will be
// permanently retired. Scheduled Stop: When Amazon EC2 determines that an
// instance must be shut down, the instances status returns an event code
// called instance-stop . Stop events include a scheduled start and end
// time. You will also be notified by email if one of your instances is set
// to stop. The email message indicates when your instance will be stopped.
// When your instance is retired, it will either be terminated (if its root
// device type is the instance-store) or stopped (if its root device type
// is an EBS volume). Instances stopped due to retirement will not be
// restarted, but you can do so manually. You can also avoid retirement of
// EBS-backed instances by manually restarting your instance when its event
// code is instance-retirement . This ensures that your instance is started
// on a different underlying host. For more information about failed status
// checks, see Troubleshooting Instances with Failed Status Checks in the
// Amazon Elastic Compute Cloud User Guide . For more information about
// working with scheduled events, see Working with an Instance That Has a
// Scheduled Event in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) DescribeInstanceStatus(req DescribeInstanceStatusRequest) (resp *DescribeInstanceStatusResult, err error) {
	resp = &DescribeInstanceStatusResult{}
	err = c.client.Do("DescribeInstanceStatus", "POST", "/", req, resp)
	return
}

// DescribeInstances describes one or more of your instances. If you
// specify one or more instance IDs, Amazon EC2 returns information for
// those instances. If you do not specify instance IDs, Amazon EC2 returns
// information for all relevant instances. If you specify an instance ID
// that is not valid, an error is returned. If you specify an instance that
// you do not own, it is not included in the returned results. Recently
// terminated instances might appear in the returned results. This interval
// is usually less than one hour.
func (c *EC2) DescribeInstances(req DescribeInstancesRequest) (resp *DescribeInstancesResult, err error) {
	resp = &DescribeInstancesResult{}
	err = c.client.Do("DescribeInstances", "POST", "/", req, resp)
	return
}

// DescribeInternetGateways is undocumented.
func (c *EC2) DescribeInternetGateways(req DescribeInternetGatewaysRequest) (resp *DescribeInternetGatewaysResult, err error) {
	resp = &DescribeInternetGatewaysResult{}
	err = c.client.Do("DescribeInternetGateways", "POST", "/", req, resp)
	return
}

// DescribeKeyPairs describes one or more of your key pairs. For more
// information about key pairs, see Key Pairs in the Amazon Elastic Compute
// Cloud User Guide
func (c *EC2) DescribeKeyPairs(req DescribeKeyPairsRequest) (resp *DescribeKeyPairsResult, err error) {
	resp = &DescribeKeyPairsResult{}
	err = c.client.Do("DescribeKeyPairs", "POST", "/", req, resp)
	return
}

// DescribeNetworkAcls describes one or more of your network ACLs. For more
// information about network ACLs, see Network ACLs in the Amazon Virtual
// Private Cloud User Guide
func (c *EC2) DescribeNetworkAcls(req DescribeNetworkAclsRequest) (resp *DescribeNetworkAclsResult, err error) {
	resp = &DescribeNetworkAclsResult{}
	err = c.client.Do("DescribeNetworkAcls", "POST", "/", req, resp)
	return
}

// DescribeNetworkInterfaceAttribute describes a network interface
// attribute. You can specify only one attribute at a time.
func (c *EC2) DescribeNetworkInterfaceAttribute(req DescribeNetworkInterfaceAttributeRequest) (resp *DescribeNetworkInterfaceAttributeResult, err error) {
	resp = &DescribeNetworkInterfaceAttributeResult{}
	err = c.client.Do("DescribeNetworkInterfaceAttribute", "POST", "/", req, resp)
	return
}

// DescribeNetworkInterfaces is undocumented.
func (c *EC2) DescribeNetworkInterfaces(req DescribeNetworkInterfacesRequest) (resp *DescribeNetworkInterfacesResult, err error) {
	resp = &DescribeNetworkInterfacesResult{}
	err = c.client.Do("DescribeNetworkInterfaces", "POST", "/", req, resp)
	return
}

// DescribePlacementGroups describes one or more of your placement groups.
// For more information about placement groups and cluster instances, see
// Cluster Instances in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) DescribePlacementGroups(req DescribePlacementGroupsRequest) (resp *DescribePlacementGroupsResult, err error) {
	resp = &DescribePlacementGroupsResult{}
	err = c.client.Do("DescribePlacementGroups", "POST", "/", req, resp)
	return
}

// DescribeRegions describes one or more regions that are currently
// available to you. For a list of the regions supported by Amazon EC2, see
// Regions and Endpoints
func (c *EC2) DescribeRegions(req DescribeRegionsRequest) (resp *DescribeRegionsResult, err error) {
	resp = &DescribeRegionsResult{}
	err = c.client.Do("DescribeRegions", "POST", "/", req, resp)
	return
}

// DescribeReservedInstances describes one or more of the Reserved
// Instances that you purchased. For more information about Reserved
// Instances, see Reserved Instances in the Amazon Elastic Compute Cloud
// User Guide
func (c *EC2) DescribeReservedInstances(req DescribeReservedInstancesRequest) (resp *DescribeReservedInstancesResult, err error) {
	resp = &DescribeReservedInstancesResult{}
	err = c.client.Do("DescribeReservedInstances", "POST", "/", req, resp)
	return
}

// DescribeReservedInstancesListings describes your account's Reserved
// Instance listings in the Reserved Instance Marketplace. The Reserved
// Instance Marketplace matches sellers who want to resell Reserved
// Instance capacity that they no longer need with buyers who want to
// purchase additional capacity. Reserved Instances bought and sold through
// the Reserved Instance Marketplace work like any other Reserved
// Instances. As a seller, you choose to list some or all of your Reserved
// Instances, and you specify the upfront price to receive for them. Your
// Reserved Instances are then listed in the Reserved Instance Marketplace
// and are available for purchase. As a buyer, you specify the
// configuration of the Reserved Instance to purchase, and the Marketplace
// matches what you're searching for with what's available. The Marketplace
// first sells the lowest priced Reserved Instances to you, and continues
// to sell available Reserved Instance listings to you until your demand is
// met. You are charged based on the total price of all of the listings
// that you purchase. For more information, see Reserved Instance
// Marketplace in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) DescribeReservedInstancesListings(req DescribeReservedInstancesListingsRequest) (resp *DescribeReservedInstancesListingsResult, err error) {
	resp = &DescribeReservedInstancesListingsResult{}
	err = c.client.Do("DescribeReservedInstancesListings", "POST", "/", req, resp)
	return
}

// DescribeReservedInstancesModifications describes the modifications made
// to your Reserved Instances. If no parameter is specified, information
// about all your Reserved Instances modification requests is returned. If
// a modification ID is specified, only information about the specific
// modification is returned. For more information, see Modifying Reserved
// Instances in the Amazon Elastic Compute Cloud User Guide.
func (c *EC2) DescribeReservedInstancesModifications(req DescribeReservedInstancesModificationsRequest) (resp *DescribeReservedInstancesModificationsResult, err error) {
	resp = &DescribeReservedInstancesModificationsResult{}
	err = c.client.Do("DescribeReservedInstancesModifications", "POST", "/", req, resp)
	return
}

// DescribeReservedInstancesOfferings describes Reserved Instance offerings
// that are available for purchase. With Reserved Instances, you purchase
// the right to launch instances for a period of time. During that time
// period, you do not receive insufficient capacity errors, and you pay a
// lower usage rate than the rate charged for On-Demand instances for the
// actual time used. For more information, see Reserved Instance
// Marketplace in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) DescribeReservedInstancesOfferings(req DescribeReservedInstancesOfferingsRequest) (resp *DescribeReservedInstancesOfferingsResult, err error) {
	resp = &DescribeReservedInstancesOfferingsResult{}
	err = c.client.Do("DescribeReservedInstancesOfferings", "POST", "/", req, resp)
	return
}

// DescribeRouteTables describes one or more of your route tables. For more
// information about route tables, see Route Tables in the Amazon Virtual
// Private Cloud User Guide
func (c *EC2) DescribeRouteTables(req DescribeRouteTablesRequest) (resp *DescribeRouteTablesResult, err error) {
	resp = &DescribeRouteTablesResult{}
	err = c.client.Do("DescribeRouteTables", "POST", "/", req, resp)
	return
}

// DescribeSecurityGroups describes one or more of your security groups. A
// security group is for use with instances either in the EC2-Classic
// platform or in a specific For more information, see Amazon EC2 Security
// Groups in the Amazon Elastic Compute Cloud User Guide and Security
// Groups for Your in the Amazon Virtual Private Cloud User Guide
func (c *EC2) DescribeSecurityGroups(req DescribeSecurityGroupsRequest) (resp *DescribeSecurityGroupsResult, err error) {
	resp = &DescribeSecurityGroupsResult{}
	err = c.client.Do("DescribeSecurityGroups", "POST", "/", req, resp)
	return
}

// DescribeSnapshotAttribute describes the specified attribute of the
// specified snapshot. You can specify only one attribute at a time. For
// more information about Amazon EBS snapshots, see Amazon EBS Snapshots in
// the Amazon Elastic Compute Cloud User Guide
func (c *EC2) DescribeSnapshotAttribute(req DescribeSnapshotAttributeRequest) (resp *DescribeSnapshotAttributeResult, err error) {
	resp = &DescribeSnapshotAttributeResult{}
	err = c.client.Do("DescribeSnapshotAttribute", "POST", "/", req, resp)
	return
}

// DescribeSnapshots describes one or more of the Amazon EBS snapshots
// available to you. Available snapshots include public snapshots available
// for any AWS account to launch, private snapshots that you own, and
// private snapshots owned by another AWS account but for which you've been
// given explicit create volume permissions. The create volume permissions
// fall into the following categories: public : The owner of the snapshot
// granted create volume permissions for the snapshot to the all group. All
// AWS accounts have create volume permissions for these snapshots.
// explicit : The owner of the snapshot granted create volume permissions
// to a specific AWS account. implicit : An AWS account has implicit create
// volume permissions for all snapshots it owns. The list of snapshots
// returned can be modified by specifying snapshot IDs, snapshot owners, or
// AWS accounts with create volume permissions. If no options are
// specified, Amazon EC2 returns all snapshots for which you have create
// volume permissions. If you specify one or more snapshot IDs, only
// snapshots that have the specified IDs are returned. If you specify an
// invalid snapshot ID, an error is returned. If you specify a snapshot ID
// for which you do not have access, it is not included in the returned
// results. If you specify one or more snapshot owners, only snapshots from
// the specified owners and for which you have access are returned. The
// results can include the AWS account IDs of the specified owners, amazon
// for snapshots owned by Amazon, or self for snapshots that you own. If
// you specify a list of restorable users, only snapshots with create
// snapshot permissions for those users are returned. You can specify AWS
// account IDs (if you own the snapshots), self for snapshots for which you
// own or have explicit permissions, or all for public snapshots. For more
// information about Amazon EBS snapshots, see Amazon EBS Snapshots in the
// Amazon Elastic Compute Cloud User Guide
func (c *EC2) DescribeSnapshots(req DescribeSnapshotsRequest) (resp *DescribeSnapshotsResult, err error) {
	resp = &DescribeSnapshotsResult{}
	err = c.client.Do("DescribeSnapshots", "POST", "/", req, resp)
	return
}

// DescribeSpotDatafeedSubscription describes the datafeed for Spot
// Instances. For more information, see Spot Instances in the Amazon
// Elastic Compute Cloud User Guide
func (c *EC2) DescribeSpotDatafeedSubscription(req DescribeSpotDatafeedSubscriptionRequest) (resp *DescribeSpotDatafeedSubscriptionResult, err error) {
	resp = &DescribeSpotDatafeedSubscriptionResult{}
	err = c.client.Do("DescribeSpotDatafeedSubscription", "POST", "/", req, resp)
	return
}

// DescribeSpotInstanceRequests describes the Spot Instance requests that
// belong to your account. Spot Instances are instances that Amazon EC2
// starts on your behalf when the maximum price that you specify exceeds
// the current Spot Price. Amazon EC2 periodically sets the Spot Price
// based on available Spot Instance capacity and current Spot Instance
// requests. For more information about Spot Instances, see Spot Instances
// in the Amazon Elastic Compute Cloud User Guide You can use
// DescribeSpotInstanceRequests to find a running Spot Instance by
// examining the response. If the status of the Spot Instance is fulfilled
// , the instance ID appears in the response and contains the identifier of
// the instance. Alternatively, you can use DescribeInstances with a filter
// to look for instances where the instance lifecycle is spot
func (c *EC2) DescribeSpotInstanceRequests(req DescribeSpotInstanceRequestsRequest) (resp *DescribeSpotInstanceRequestsResult, err error) {
	resp = &DescribeSpotInstanceRequestsResult{}
	err = c.client.Do("DescribeSpotInstanceRequests", "POST", "/", req, resp)
	return
}

// DescribeSpotPriceHistory describes the Spot Price history. Spot
// Instances are instances that Amazon EC2 starts on your behalf when the
// maximum price that you specify exceeds the current Spot Price. Amazon
// EC2 periodically sets the Spot Price based on available Spot Instance
// capacity and current Spot Instance requests. For more information about
// Spot Instances, see Spot Instances in the Amazon Elastic Compute Cloud
// User Guide When you specify an Availability Zone, this operation
// describes the price history for the specified Availability Zone with the
// most recent set of prices listed first. If you don't specify an
// Availability Zone, you get the prices across all Availability Zones,
// starting with the most recent set. However, if you're using an API
// version earlier than 2011-05-15, you get the lowest price across the
// region for the specified time period. The prices returned are listed in
// chronological order, from the oldest to the most recent. When you
// specify the start and end time options, this operation returns two
// pieces of data: the prices of the instance types within the time range
// that you specified and the time when the price changed. The price is
// valid within the time period that you specified; the response merely
// indicates the last time that the price changed.
func (c *EC2) DescribeSpotPriceHistory(req DescribeSpotPriceHistoryRequest) (resp *DescribeSpotPriceHistoryResult, err error) {
	resp = &DescribeSpotPriceHistoryResult{}
	err = c.client.Do("DescribeSpotPriceHistory", "POST", "/", req, resp)
	return
}

// DescribeSubnets describes one or more of your subnets. For more
// information about subnets, see Your VPC and Subnets in the Amazon
// Virtual Private Cloud User Guide
func (c *EC2) DescribeSubnets(req DescribeSubnetsRequest) (resp *DescribeSubnetsResult, err error) {
	resp = &DescribeSubnetsResult{}
	err = c.client.Do("DescribeSubnets", "POST", "/", req, resp)
	return
}

// DescribeTags describes one or more of the tags for your EC2 resources.
// For more information about tags, see Tagging Your Resources in the
// Amazon Elastic Compute Cloud User Guide
func (c *EC2) DescribeTags(req DescribeTagsRequest) (resp *DescribeTagsResult, err error) {
	resp = &DescribeTagsResult{}
	err = c.client.Do("DescribeTags", "POST", "/", req, resp)
	return
}

// DescribeVolumeAttribute describes the specified attribute of the
// specified volume. You can specify only one attribute at a time. For more
// information about Amazon EBS volumes, see Amazon EBS Volumes in the
// Amazon Elastic Compute Cloud User Guide
func (c *EC2) DescribeVolumeAttribute(req DescribeVolumeAttributeRequest) (resp *DescribeVolumeAttributeResult, err error) {
	resp = &DescribeVolumeAttributeResult{}
	err = c.client.Do("DescribeVolumeAttribute", "POST", "/", req, resp)
	return
}

// DescribeVolumeStatus describes the status of the specified volumes.
// Volume status provides the result of the checks performed on your
// volumes to determine events that can impair the performance of your
// volumes. The performance of a volume can be affected if an issue occurs
// on the volume's underlying host. If the volume's underlying host
// experiences a power outage or system issue, after the system is
// restored, there could be data inconsistencies on the volume. Volume
// events notify you if this occurs. Volume actions notify you if any
// action needs to be taken in response to the event. The
// DescribeVolumeStatus operation provides the following information about
// the specified volumes: Status : Reflects the current status of the
// volume. The possible values are ok , impaired , warning , or
// insufficient-data . If all checks pass, the overall status of the volume
// is ok . If the check fails, the overall status is impaired . If the
// status is insufficient-data , then the checks may still be taking place
// on your volume at the time. We recommend that you retry the request. For
// more information on volume status, see Monitoring the Status of Your
// Volumes Events : Reflect the cause of a volume status and may require
// you to take action. For example, if your volume returns an impaired
// status, then the volume event might be potential-data-inconsistency .
// This means that your volume has been affected by an issue with the
// underlying host, has all I/O operations disabled, and may have
// inconsistent data. Actions : Reflect the actions you may have to take in
// response to an event. For example, if the status of the volume is
// impaired and the volume event shows potential-data-inconsistency , then
// the action shows enable-volume-io . This means that you may want to
// enable the I/O operations for the volume by calling the EnableVolumeIO
// action and then check the volume for data consistency. Volume status is
// based on the volume status checks, and does not reflect the volume
// state. Therefore, volume status does not indicate volumes in the error
// state (for example, when a volume is incapable of accepting
func (c *EC2) DescribeVolumeStatus(req DescribeVolumeStatusRequest) (resp *DescribeVolumeStatusResult, err error) {
	resp = &DescribeVolumeStatusResult{}
	err = c.client.Do("DescribeVolumeStatus", "POST", "/", req, resp)
	return
}

// DescribeVolumes describes the specified Amazon EBS volumes. If you are
// describing a long list of volumes, you can paginate the output to make
// the list more manageable. The MaxResults parameter sets the maximum
// number of results returned in a single page. If the list of results
// exceeds your MaxResults value, then that number of results is returned
// along with a NextToken value that can be passed to a subsequent
// DescribeVolumes request to retrieve the remaining results. For more
// information about Amazon EBS volumes, see Amazon EBS Volumes in the
// Amazon Elastic Compute Cloud User Guide
func (c *EC2) DescribeVolumes(req DescribeVolumesRequest) (resp *DescribeVolumesResult, err error) {
	resp = &DescribeVolumesResult{}
	err = c.client.Do("DescribeVolumes", "POST", "/", req, resp)
	return
}

// DescribeVpcAttribute describes the specified attribute of the specified
// You can specify only one attribute at a time.
func (c *EC2) DescribeVpcAttribute(req DescribeVpcAttributeRequest) (resp *DescribeVpcAttributeResult, err error) {
	resp = &DescribeVpcAttributeResult{}
	err = c.client.Do("DescribeVpcAttribute", "POST", "/", req, resp)
	return
}

// DescribeVpcPeeringConnections describes one or more of your VPC peering
// connections.
func (c *EC2) DescribeVpcPeeringConnections(req DescribeVpcPeeringConnectionsRequest) (resp *DescribeVpcPeeringConnectionsResult, err error) {
	resp = &DescribeVpcPeeringConnectionsResult{}
	err = c.client.Do("DescribeVpcPeeringConnections", "POST", "/", req, resp)
	return
}

// DescribeVpcs is undocumented.
func (c *EC2) DescribeVpcs(req DescribeVpcsRequest) (resp *DescribeVpcsResult, err error) {
	resp = &DescribeVpcsResult{}
	err = c.client.Do("DescribeVpcs", "POST", "/", req, resp)
	return
}

// DescribeVpnConnections describes one or more of your VPN connections.
// For more information about VPN connections, see Adding a Hardware
// Virtual Private Gateway to Your in the Amazon Virtual Private Cloud User
// Guide
func (c *EC2) DescribeVpnConnections(req DescribeVpnConnectionsRequest) (resp *DescribeVpnConnectionsResult, err error) {
	resp = &DescribeVpnConnectionsResult{}
	err = c.client.Do("DescribeVpnConnections", "POST", "/", req, resp)
	return
}

// DescribeVpnGateways describes one or more of your virtual private
// gateways. For more information about virtual private gateways, see
// Adding an IPsec Hardware VPN to Your in the Amazon Virtual Private Cloud
// User Guide
func (c *EC2) DescribeVpnGateways(req DescribeVpnGatewaysRequest) (resp *DescribeVpnGatewaysResult, err error) {
	resp = &DescribeVpnGatewaysResult{}
	err = c.client.Do("DescribeVpnGateways", "POST", "/", req, resp)
	return
}

// DetachInternetGateway detaches an Internet gateway from a disabling
// connectivity between the Internet and the The VPC must not contain any
// running instances with Elastic IP addresses.
func (c *EC2) DetachInternetGateway(req DetachInternetGatewayRequest) (err error) {
	// NRE
	err = c.client.Do("DetachInternetGateway", "POST", "/", req, nil)
	return
}

// DetachNetworkInterface is undocumented.
func (c *EC2) DetachNetworkInterface(req DetachNetworkInterfaceRequest) (err error) {
	// NRE
	err = c.client.Do("DetachNetworkInterface", "POST", "/", req, nil)
	return
}

// DetachVolume detaches an Amazon EBS volume from an instance. Make sure
// to unmount any file systems on the device within your operating system
// before detaching the volume. Failure to do so results in the volume
// being stuck in a busy state while detaching. If an Amazon EBS volume is
// the root device of an instance, it can't be detached while the instance
// is running. To detach the root volume, stop the instance first. If the
// root volume is detached from an instance with an AWS Marketplace product
// code, then the AWS Marketplace product codes from that volume are no
// longer associated with the instance. For more information, see Detaching
// an Amazon EBS Volume in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) DetachVolume(req DetachVolumeRequest) (resp *VolumeAttachment, err error) {
	resp = &VolumeAttachment{}
	err = c.client.Do("DetachVolume", "POST", "/", req, resp)
	return
}

// DetachVpnGateway detaches a virtual private gateway from a You do this
// if you're planning to turn off the VPC and not use it anymore. You can
// confirm a virtual private gateway has been completely detached from a
// VPC by describing the virtual private gateway (any attachments to the
// virtual private gateway are also described). You must wait for the
// attachment's state to switch to detached before you can delete the VPC
// or attach a different VPC to the virtual private gateway.
func (c *EC2) DetachVpnGateway(req DetachVpnGatewayRequest) (err error) {
	// NRE
	err = c.client.Do("DetachVpnGateway", "POST", "/", req, nil)
	return
}

// DisableVgwRoutePropagation disables a virtual private gateway from
// propagating routes to a specified route table of a
func (c *EC2) DisableVgwRoutePropagation(req DisableVgwRoutePropagationRequest) (err error) {
	// NRE
	err = c.client.Do("DisableVgwRoutePropagation", "POST", "/", req, nil)
	return
}

// DisassociateAddress disassociates an Elastic IP address from the
// instance or network interface it's associated with. An Elastic IP
// address is for use in either the EC2-Classic platform or in a For more
// information, see Elastic IP Addresses in the Amazon Elastic Compute
// Cloud User Guide This is an idempotent operation. If you perform the
// operation more than once, Amazon EC2 doesn't return an error.
func (c *EC2) DisassociateAddress(req DisassociateAddressRequest) (err error) {
	// NRE
	err = c.client.Do("DisassociateAddress", "POST", "/", req, nil)
	return
}

// DisassociateRouteTable disassociates a subnet from a route table. After
// you perform this action, the subnet no longer uses the routes in the
// route table. Instead, it uses the routes in the VPC's main route table.
// For more information about route tables, see Route Tables in the Amazon
// Virtual Private Cloud User Guide
func (c *EC2) DisassociateRouteTable(req DisassociateRouteTableRequest) (err error) {
	// NRE
	err = c.client.Do("DisassociateRouteTable", "POST", "/", req, nil)
	return
}

// EnableVgwRoutePropagation enables a virtual private gateway to propagate
// routes to the specified route table of a
func (c *EC2) EnableVgwRoutePropagation(req EnableVgwRoutePropagationRequest) (err error) {
	// NRE
	err = c.client.Do("EnableVgwRoutePropagation", "POST", "/", req, nil)
	return
}

// EnableVolumeIO enables I/O operations for a volume that had I/O
// operations disabled because the data on the volume was potentially
// inconsistent.
func (c *EC2) EnableVolumeIO(req EnableVolumeIORequest) (err error) {
	// NRE
	err = c.client.Do("EnableVolumeIO", "POST", "/", req, nil)
	return
}

// GetConsoleOutput gets the console output for the specified instance.
// Instances do not have a physical monitor through which you can view
// their console output. They also lack physical controls that allow you to
// power up, reboot, or shut them down. To allow these actions, we provide
// them through the Amazon EC2 API and command line interface. Instance
// console output is buffered and posted shortly after instance boot,
// reboot, and termination. Amazon EC2 preserves the most recent 64 KB
// output which is available for at least one hour after the most recent
// post. For Linux/Unix instances, the instance console output displays the
// exact console output that would normally be displayed on a physical
// monitor attached to a machine. This output is buffered because the
// instance produces it and then posts it to a store where the instance's
// owner can retrieve it. For Windows instances, the instance console
// output displays the last three system event log errors.
func (c *EC2) GetConsoleOutput(req GetConsoleOutputRequest) (resp *GetConsoleOutputResult, err error) {
	resp = &GetConsoleOutputResult{}
	err = c.client.Do("GetConsoleOutput", "POST", "/", req, resp)
	return
}

// GetPasswordData retrieves the encrypted administrator password for an
// instance running Windows. The Windows password is generated at boot if
// the EC2Config service plugin, Ec2SetPassword , is enabled. This usually
// only happens the first time an AMI is launched, and then Ec2SetPassword
// is automatically disabled. The password is not generated for rebundled
// AMIs unless Ec2SetPassword is enabled before bundling. The password is
// encrypted using the key pair that you specified when you launched the
// instance. You must provide the corresponding key pair file. Password
// generation and encryption takes a few moments. We recommend that you
// wait up to 15 minutes after launching an instance before trying to
// retrieve the generated password.
func (c *EC2) GetPasswordData(req GetPasswordDataRequest) (resp *GetPasswordDataResult, err error) {
	resp = &GetPasswordDataResult{}
	err = c.client.Do("GetPasswordData", "POST", "/", req, resp)
	return
}

// ImportInstance creates an import instance task using metadata from the
// specified disk image. After importing the image, you then upload it
// using the command in the EC2 command line tools. For more information,
// see Using the Command Line Tools to Import Your Virtual Machine to
// Amazon EC2
func (c *EC2) ImportInstance(req ImportInstanceRequest) (resp *ImportInstanceResult, err error) {
	resp = &ImportInstanceResult{}
	err = c.client.Do("ImportInstance", "POST", "/", req, resp)
	return
}

// ImportKeyPair imports the public key from an RSA key pair that you
// created with a third-party tool. Compare this with CreateKeyPair , in
// which AWS creates the key pair and gives the keys to you keeps a copy of
// the public key). With ImportKeyPair, you create the key pair and give
// AWS just the public key. The private key is never transferred between
// you and For more information about key pairs, see Key Pairs in the
// Amazon Elastic Compute Cloud User Guide
func (c *EC2) ImportKeyPair(req ImportKeyPairRequest) (resp *ImportKeyPairResult, err error) {
	resp = &ImportKeyPairResult{}
	err = c.client.Do("ImportKeyPair", "POST", "/", req, resp)
	return
}

// ImportVolume creates an import volume task using metadata from the
// specified disk image. After importing the image, you then upload it
// using the command in the Amazon EC2 command-line interface tools. For
// more information, see Using the Command Line Tools to Import Your
// Virtual Machine to Amazon EC2 in the Amazon Elastic Compute Cloud User
// Guide
func (c *EC2) ImportVolume(req ImportVolumeRequest) (resp *ImportVolumeResult, err error) {
	resp = &ImportVolumeResult{}
	err = c.client.Do("ImportVolume", "POST", "/", req, resp)
	return
}

// ModifyImageAttribute modifies the specified attribute of the specified
// You can specify only one attribute at a time. AWS Marketplace product
// codes cannot be modified. Images with an AWS Marketplace product code
// cannot be made public.
func (c *EC2) ModifyImageAttribute(req ModifyImageAttributeRequest) (err error) {
	// NRE
	err = c.client.Do("ModifyImageAttribute", "POST", "/", req, nil)
	return
}

// ModifyInstanceAttribute modifies the specified attribute of the
// specified instance. You can specify only one attribute at a time. To
// modify some attributes, the instance must be stopped. For more
// information, see Modifying Attributes of a Stopped Instance in the
// Amazon Elastic Compute Cloud User Guide
func (c *EC2) ModifyInstanceAttribute(req ModifyInstanceAttributeRequest) (err error) {
	// NRE
	err = c.client.Do("ModifyInstanceAttribute", "POST", "/", req, nil)
	return
}

// ModifyNetworkInterfaceAttribute modifies the specified network interface
// attribute. You can specify only one attribute at a time.
func (c *EC2) ModifyNetworkInterfaceAttribute(req ModifyNetworkInterfaceAttributeRequest) (err error) {
	// NRE
	err = c.client.Do("ModifyNetworkInterfaceAttribute", "POST", "/", req, nil)
	return
}

// ModifyReservedInstances modifies the Availability Zone, instance count,
// instance type, or network platform (EC2-Classic or EC2-VPC) of your
// Reserved Instances. The Reserved Instances to be modified must be
// identical, except for Availability Zone, network platform, and instance
// type. For more information, see Modifying Reserved Instances in the
// Amazon Elastic Compute Cloud User Guide.
func (c *EC2) ModifyReservedInstances(req ModifyReservedInstancesRequest) (resp *ModifyReservedInstancesResult, err error) {
	resp = &ModifyReservedInstancesResult{}
	err = c.client.Do("ModifyReservedInstances", "POST", "/", req, resp)
	return
}

// ModifySnapshotAttribute adds or removes permission settings for the
// specified snapshot. You may add or remove specified AWS account IDs from
// a snapshot's list of create volume permissions, but you cannot do both
// in a single API call. If you need to both add and remove account IDs for
// a snapshot, you must use multiple API calls. For more information on
// modifying snapshot permissions, see Sharing Snapshots in the Amazon
// Elastic Compute Cloud User Guide Snapshots with AWS Marketplace product
// codes cannot be made public.
func (c *EC2) ModifySnapshotAttribute(req ModifySnapshotAttributeRequest) (err error) {
	// NRE
	err = c.client.Do("ModifySnapshotAttribute", "POST", "/", req, nil)
	return
}

// ModifySubnetAttribute is undocumented.
func (c *EC2) ModifySubnetAttribute(req ModifySubnetAttributeRequest) (err error) {
	// NRE
	err = c.client.Do("ModifySubnetAttribute", "POST", "/", req, nil)
	return
}

// ModifyVolumeAttribute modifies a volume attribute. By default, all I/O
// operations for the volume are suspended when the data on the volume is
// determined to be potentially inconsistent, to prevent undetectable,
// latent data corruption. The I/O access to the volume can be resumed by
// first enabling I/O access and then checking the data consistency on your
// volume. You can change the default behavior to resume I/O operations. We
// recommend that you change this only for boot volumes or for volumes that
// are stateless or disposable.
func (c *EC2) ModifyVolumeAttribute(req ModifyVolumeAttributeRequest) (err error) {
	// NRE
	err = c.client.Do("ModifyVolumeAttribute", "POST", "/", req, nil)
	return
}

// ModifyVpcAttribute is undocumented.
func (c *EC2) ModifyVpcAttribute(req ModifyVpcAttributeRequest) (err error) {
	// NRE
	err = c.client.Do("ModifyVpcAttribute", "POST", "/", req, nil)
	return
}

// MonitorInstances enables monitoring for a running instance. For more
// information about monitoring instances, see Monitoring Your Instances
// and Volumes in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) MonitorInstances(req MonitorInstancesRequest) (resp *MonitorInstancesResult, err error) {
	resp = &MonitorInstancesResult{}
	err = c.client.Do("MonitorInstances", "POST", "/", req, resp)
	return
}

// PurchaseReservedInstancesOffering purchases a Reserved Instance for use
// with your account. With Amazon EC2 Reserved Instances, you obtain a
// capacity reservation for a certain instance configuration over a
// specified period of time. You pay a lower usage rate than with On-Demand
// instances for the time that you actually use the capacity reservation.
// Use DescribeReservedInstancesOfferings to get a list of Reserved
// Instance offerings that match your specifications. After you've
// purchased a Reserved Instance, you can check for your new Reserved
// Instance with DescribeReservedInstances For more information, see
// Reserved Instances and Reserved Instance Marketplace in the Amazon
// Elastic Compute Cloud User Guide
func (c *EC2) PurchaseReservedInstancesOffering(req PurchaseReservedInstancesOfferingRequest) (resp *PurchaseReservedInstancesOfferingResult, err error) {
	resp = &PurchaseReservedInstancesOfferingResult{}
	err = c.client.Do("PurchaseReservedInstancesOffering", "POST", "/", req, resp)
	return
}

// RebootInstances requests a reboot of one or more instances. This
// operation is asynchronous; it only queues a request to reboot the
// specified instances. The operation succeeds if the instances are valid
// and belong to you. Requests to reboot terminated instances are ignored.
// If a Linux/Unix instance does not cleanly shut down within four minutes,
// Amazon EC2 performs a hard reboot. For more information about
// troubleshooting, see Getting Console Output and Rebooting Instances in
// the Amazon Elastic Compute Cloud User Guide
func (c *EC2) RebootInstances(req RebootInstancesRequest) (err error) {
	// NRE
	err = c.client.Do("RebootInstances", "POST", "/", req, nil)
	return
}

// RegisterImage registers an When you're creating an this is the final
// step you must complete before you can launch an instance from the For
// more information about creating AMIs, see Creating Your Own AMIs in the
// Amazon Elastic Compute Cloud User Guide For Amazon EBS-backed instances,
// CreateImage creates and registers the AMI in a single request, so you
// don't have to register the AMI yourself. You can also use RegisterImage
// to create an Amazon EBS-backed AMI from a snapshot of a root device
// volume. For more information, see Launching an Instance from a Snapshot
// in the Amazon Elastic Compute Cloud User Guide If needed, you can
// deregister an AMI at any time. Any modifications you make to an AMI
// backed by an instance store volume invalidates its registration. If you
// make changes to an image, deregister the previous image and register the
// new image. You can't register an image where a secondary (non-root)
// snapshot has AWS Marketplace product codes.
func (c *EC2) RegisterImage(req RegisterImageRequest) (resp *RegisterImageResult, err error) {
	resp = &RegisterImageResult{}
	err = c.client.Do("RegisterImage", "POST", "/", req, resp)
	return
}

// RejectVpcPeeringConnection rejects a VPC peering connection request. The
// VPC peering connection must be in the pending-acceptance state. Use the
// DescribeVpcPeeringConnections request to view your outstanding VPC
// peering connection requests. To delete an active VPC peering connection,
// or to delete a VPC peering connection request that you initiated, use
// DeleteVpcPeeringConnection
func (c *EC2) RejectVpcPeeringConnection(req RejectVpcPeeringConnectionRequest) (resp *RejectVpcPeeringConnectionResult, err error) {
	resp = &RejectVpcPeeringConnectionResult{}
	err = c.client.Do("RejectVpcPeeringConnection", "POST", "/", req, resp)
	return
}

// ReleaseAddress releases the specified Elastic IP address. After
// releasing an Elastic IP address, it is released to the IP address pool
// and might be unavailable to you. Be sure to update your DNS records and
// any servers or devices that communicate with the address. If you attempt
// to release an Elastic IP address that you already released, you'll get
// an AuthFailure error if the address is already allocated to another AWS
// account. [EC2-Classic, default Releasing an Elastic IP address
// automatically disassociates it from any instance that it's associated
// with. To disassociate an Elastic IP address without releasing it, use
// DisassociateAddress [Nondefault You must use DisassociateAddress to
// disassociate the Elastic IP address before you try to release it.
// Otherwise, Amazon EC2 returns an error InvalidIPAddress.InUse
func (c *EC2) ReleaseAddress(req ReleaseAddressRequest) (err error) {
	// NRE
	err = c.client.Do("ReleaseAddress", "POST", "/", req, nil)
	return
}

// ReplaceNetworkAclAssociation changes which network ACL a subnet is
// associated with. By default when you create a subnet, it's automatically
// associated with the default network For more information about network
// ACLs, see Network ACLs in the Amazon Virtual Private Cloud User Guide
func (c *EC2) ReplaceNetworkAclAssociation(req ReplaceNetworkAclAssociationRequest) (resp *ReplaceNetworkAclAssociationResult, err error) {
	resp = &ReplaceNetworkAclAssociationResult{}
	err = c.client.Do("ReplaceNetworkAclAssociation", "POST", "/", req, resp)
	return
}

// ReplaceNetworkAclEntry replaces an entry (rule) in a network For more
// information about network ACLs, see Network ACLs in the Amazon Virtual
// Private Cloud User Guide
func (c *EC2) ReplaceNetworkAclEntry(req ReplaceNetworkAclEntryRequest) (err error) {
	// NRE
	err = c.client.Do("ReplaceNetworkAclEntry", "POST", "/", req, nil)
	return
}

// ReplaceRoute replaces an existing route within a route table in a You
// must provide only one of the following: Internet gateway or virtual
// private gateway, NAT instance, VPC peering connection, or network
// interface. For more information about route tables, see Route Tables in
// the Amazon Virtual Private Cloud User Guide
func (c *EC2) ReplaceRoute(req ReplaceRouteRequest) (err error) {
	// NRE
	err = c.client.Do("ReplaceRoute", "POST", "/", req, nil)
	return
}

// ReplaceRouteTableAssociation changes the route table associated with a
// given subnet in a After the operation completes, the subnet uses the
// routes in the new route table it's associated with. For more information
// about route tables, see Route Tables in the Amazon Virtual Private Cloud
// User Guide You can also use ReplaceRouteTableAssociation to change which
// table is the main route table in the You just specify the main route
// table's association ID and the route table to be the new main route
// table.
func (c *EC2) ReplaceRouteTableAssociation(req ReplaceRouteTableAssociationRequest) (resp *ReplaceRouteTableAssociationResult, err error) {
	resp = &ReplaceRouteTableAssociationResult{}
	err = c.client.Do("ReplaceRouteTableAssociation", "POST", "/", req, resp)
	return
}

// ReportInstanceStatus submits feedback about the status of an instance.
// The instance must be in the running state. If your experience with the
// instance differs from the instance status returned by
// DescribeInstanceStatus , use ReportInstanceStatus to report your
// experience with the instance. Amazon EC2 collects this information to
// improve the accuracy of status checks. Use of this action does not
// change the value returned by DescribeInstanceStatus
func (c *EC2) ReportInstanceStatus(req ReportInstanceStatusRequest) (err error) {
	// NRE
	err = c.client.Do("ReportInstanceStatus", "POST", "/", req, nil)
	return
}

// RequestSpotInstances creates a Spot Instance request. Spot Instances are
// instances that Amazon EC2 starts on your behalf when the maximum price
// that you specify exceeds the current Spot Price. Amazon EC2 periodically
// sets the Spot Price based on available Spot Instance capacity and
// current Spot Instance requests. For more information about Spot
// Instances, see Spot Instances in the Amazon Elastic Compute Cloud User
// Guide Users must be subscribed to the required product to run an
// instance with AWS Marketplace product codes.
func (c *EC2) RequestSpotInstances(req RequestSpotInstancesRequest) (resp *RequestSpotInstancesResult, err error) {
	resp = &RequestSpotInstancesResult{}
	err = c.client.Do("RequestSpotInstances", "POST", "/", req, resp)
	return
}

// ResetImageAttribute resets an attribute of an AMI to its default value.
func (c *EC2) ResetImageAttribute(req ResetImageAttributeRequest) (err error) {
	// NRE
	err = c.client.Do("ResetImageAttribute", "POST", "/", req, nil)
	return
}

// ResetInstanceAttribute resets an attribute of an instance to its default
// value. To reset the kernel or ramdisk , the instance must be in a
// stopped state. To reset the SourceDestCheck , the instance can be either
// running or stopped. The SourceDestCheck attribute controls whether
// source/destination checking is enabled. The default value is true ,
// which means checking is enabled. This value must be false for a NAT
// instance to perform For more information, see NAT Instances in the
// Amazon Virtual Private Cloud User Guide
func (c *EC2) ResetInstanceAttribute(req ResetInstanceAttributeRequest) (err error) {
	// NRE
	err = c.client.Do("ResetInstanceAttribute", "POST", "/", req, nil)
	return
}

// ResetNetworkInterfaceAttribute resets a network interface attribute. You
// can specify only one attribute at a time.
func (c *EC2) ResetNetworkInterfaceAttribute(req ResetNetworkInterfaceAttributeRequest) (err error) {
	// NRE
	err = c.client.Do("ResetNetworkInterfaceAttribute", "POST", "/", req, nil)
	return
}

// ResetSnapshotAttribute resets permission settings for the specified
// snapshot. For more information on modifying snapshot permissions, see
// Sharing Snapshots in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) ResetSnapshotAttribute(req ResetSnapshotAttributeRequest) (err error) {
	// NRE
	err = c.client.Do("ResetSnapshotAttribute", "POST", "/", req, nil)
	return
}

// RevokeSecurityGroupEgress removes one or more egress rules from a
// security group for EC2-VPC. The values that you specify in the revoke
// request (for example, ports) must match the existing rule's values for
// the rule to be revoked. Each rule consists of the protocol and the range
// or source security group. For the TCP and UDP protocols, you must also
// specify the destination port or range of ports. For the protocol, you
// must also specify the type and code. Rule changes are propagated to
// instances within the security group as quickly as possible. However, a
// small delay might occur.
func (c *EC2) RevokeSecurityGroupEgress(req RevokeSecurityGroupEgressRequest) (err error) {
	// NRE
	err = c.client.Do("RevokeSecurityGroupEgress", "POST", "/", req, nil)
	return
}

// RevokeSecurityGroupIngress removes one or more ingress rules from a
// security group. The values that you specify in the revoke request (for
// example, ports) must match the existing rule's values for the rule to be
// removed. Each rule consists of the protocol and the range or source
// security group. For the TCP and UDP protocols, you must also specify the
// destination port or range of ports. For the protocol, you must also
// specify the type and code. Rule changes are propagated to instances
// within the security group as quickly as possible. However, a small delay
// might occur.
func (c *EC2) RevokeSecurityGroupIngress(req RevokeSecurityGroupIngressRequest) (err error) {
	// NRE
	err = c.client.Do("RevokeSecurityGroupIngress", "POST", "/", req, nil)
	return
}

// RunInstances launches the specified number of instances using an AMI for
// which you have permissions. When you launch an instance, it enters the
// pending state. After the instance is ready for you, it enters the
// running state. To check the state of your instance, call
// DescribeInstances If you don't specify a security group when launching
// an instance, Amazon EC2 uses the default security group. For more
// information, see Security Groups in the Amazon Elastic Compute Cloud
// User Guide Linux instances have access to the public key of the key pair
// at boot. You can use this key to provide secure access to the instance.
// Amazon EC2 public images use this feature to provide secure access
// without passwords. For more information, see Key Pairs in the Amazon
// Elastic Compute Cloud User Guide You can provide optional user data when
// launching an instance. For more information, see Instance Metadata in
// the Amazon Elastic Compute Cloud User Guide If any of the AMIs have a
// product code attached for which the user has not subscribed,
// RunInstances fails. T2 instance types can only be launched into a If you
// do not have a default or if you do not specify a subnet ID in the
// request, RunInstances fails. For more information about troubleshooting,
// see What To Do If An Instance Immediately Terminates , and
// Troubleshooting Connecting to Your Instance in the Amazon Elastic
// Compute Cloud User Guide
func (c *EC2) RunInstances(req RunInstancesRequest) (resp *Reservation, err error) {
	resp = &Reservation{}
	err = c.client.Do("RunInstances", "POST", "/", req, resp)
	return
}

// StartInstances starts an Amazon EBS-backed AMI that you've previously
// stopped. Instances that use Amazon EBS volumes as their root devices can
// be quickly stopped and started. When an instance is stopped, the compute
// resources are released and you are not billed for hourly instance usage.
// However, your root partition Amazon EBS volume remains, continues to
// persist your data, and you are charged for Amazon EBS volume usage. You
// can restart your instance at any time. Each time you transition an
// instance from stopped to started, Amazon EC2 charges a full instance
// hour, even if transitions happen multiple times within a single hour.
// Before stopping an instance, make sure it is in a state from which it
// can be restarted. Stopping an instance does not preserve data stored in
// Performing this operation on an instance that uses an instance store as
// its root device returns an error. For more information, see Stopping
// Instances in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) StartInstances(req StartInstancesRequest) (resp *StartInstancesResult, err error) {
	resp = &StartInstancesResult{}
	err = c.client.Do("StartInstances", "POST", "/", req, resp)
	return
}

// StopInstances stops an Amazon EBS-backed instance. Each time you
// transition an instance from stopped to started, Amazon EC2 charges a
// full instance hour, even if transitions happen multiple times within a
// single hour. You can't start or stop Spot Instances. Instances that use
// Amazon EBS volumes as their root devices can be quickly stopped and
// started. When an instance is stopped, the compute resources are released
// and you are not billed for hourly instance usage. However, your root
// partition Amazon EBS volume remains, continues to persist your data, and
// you are charged for Amazon EBS volume usage. You can restart your
// instance at any time. Before stopping an instance, make sure it is in a
// state from which it can be restarted. Stopping an instance does not
// preserve data stored in Performing this operation on an instance that
// uses an instance store as its root device returns an error. You can
// stop, start, and terminate EBS-backed instances. You can only terminate
// instance store-backed instances. What happens to an instance differs if
// you stop it or terminate it. For example, when you stop an instance, the
// root device and any other devices attached to the instance persist. When
// you terminate an instance, the root device and any other devices
// attached during the instance launch are automatically deleted. For more
// information about the differences between stopping and terminating
// instances, see Instance Lifecycle in the Amazon Elastic Compute Cloud
// User Guide For more information about troubleshooting, see
// Troubleshooting Stopping Your Instance in the Amazon Elastic Compute
// Cloud User Guide
func (c *EC2) StopInstances(req StopInstancesRequest) (resp *StopInstancesResult, err error) {
	resp = &StopInstancesResult{}
	err = c.client.Do("StopInstances", "POST", "/", req, resp)
	return
}

// TerminateInstances shuts down one or more instances. This operation is
// idempotent; if you terminate an instance more than once, each call
// succeeds. Terminated instances remain visible after termination (for
// approximately one hour). By default, Amazon EC2 deletes all Amazon EBS
// volumes that were attached when the instance launched. Volumes attached
// after instance launch continue running. You can stop, start, and
// terminate EBS-backed instances. You can only terminate instance
// store-backed instances. What happens to an instance differs if you stop
// it or terminate it. For example, when you stop an instance, the root
// device and any other devices attached to the instance persist. When you
// terminate an instance, the root device and any other devices attached
// during the instance launch are automatically deleted. For more
// information about the differences between stopping and terminating
// instances, see Instance Lifecycle in the Amazon Elastic Compute Cloud
// User Guide For more information about troubleshooting, see
// Troubleshooting Terminating Your Instance in the Amazon Elastic Compute
// Cloud User Guide
func (c *EC2) TerminateInstances(req TerminateInstancesRequest) (resp *TerminateInstancesResult, err error) {
	resp = &TerminateInstancesResult{}
	err = c.client.Do("TerminateInstances", "POST", "/", req, resp)
	return
}

// UnassignPrivateIPAddresses unassigns one or more secondary private IP
// addresses from a network interface.
func (c *EC2) UnassignPrivateIPAddresses(req UnassignPrivateIPAddressesRequest) (err error) {
	// NRE
	err = c.client.Do("UnassignPrivateIpAddresses", "POST", "/", req, nil)
	return
}

// UnmonitorInstances disables monitoring for a running instance. For more
// information about monitoring instances, see Monitoring Your Instances
// and Volumes in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) UnmonitorInstances(req UnmonitorInstancesRequest) (resp *UnmonitorInstancesResult, err error) {
	resp = &UnmonitorInstancesResult{}
	err = c.client.Do("UnmonitorInstances", "POST", "/", req, resp)
	return
}

// AcceptVpcPeeringConnectionRequest is undocumented.
type AcceptVpcPeeringConnectionRequest struct {
	DryRun                 bool   `xml:"dryRun"`
	VpcPeeringConnectionID string `xml:"vpcPeeringConnectionId"`
}

// AcceptVpcPeeringConnectionResult is undocumented.
type AcceptVpcPeeringConnectionResult struct {
	VpcPeeringConnection VpcPeeringConnection `xml:"vpcPeeringConnection"`
}

// AccountAttribute is undocumented.
type AccountAttribute struct {
	AttributeName   string                  `xml:"attributeName"`
	AttributeValues []AccountAttributeValue `xml:"attributeValueSet>item"`
}

// AccountAttributeValue is undocumented.
type AccountAttributeValue struct {
	AttributeValue string `xml:"attributeValue"`
}

// Address is undocumented.
type Address struct {
	AllocationID            string `xml:"allocationId"`
	AssociationID           string `xml:"associationId"`
	Domain                  string `xml:"domain"`
	InstanceID              string `xml:"instanceId"`
	NetworkInterfaceID      string `xml:"networkInterfaceId"`
	NetworkInterfaceOwnerID string `xml:"networkInterfaceOwnerId"`
	PrivateIPAddress        string `xml:"privateIpAddress"`
	PublicIP                string `xml:"publicIp"`
}

// AllocateAddressRequest is undocumented.
type AllocateAddressRequest struct {
	Domain string `xml:"Domain"`
	DryRun bool   `xml:"dryRun"`
}

// AllocateAddressResult is undocumented.
type AllocateAddressResult struct {
	AllocationID string `xml:"allocationId"`
	Domain       string `xml:"domain"`
	PublicIP     string `xml:"publicIp"`
}

// AssignPrivateIPAddressesRequest is undocumented.
type AssignPrivateIPAddressesRequest struct {
	AllowReassignment              bool     `xml:"allowReassignment"`
	NetworkInterfaceID             string   `xml:"networkInterfaceId"`
	PrivateIPAddresses             []string `xml:"privateIpAddress>PrivateIpAddress"`
	SecondaryPrivateIPAddressCount int      `xml:"secondaryPrivateIpAddressCount"`
}

// AssociateAddressRequest is undocumented.
type AssociateAddressRequest struct {
	AllocationID       string `xml:"AllocationId"`
	AllowReassociation bool   `xml:"allowReassociation"`
	DryRun             bool   `xml:"dryRun"`
	InstanceID         string `xml:"InstanceId"`
	NetworkInterfaceID string `xml:"networkInterfaceId"`
	PrivateIPAddress   string `xml:"privateIpAddress"`
	PublicIP           string `xml:"PublicIp"`
}

// AssociateAddressResult is undocumented.
type AssociateAddressResult struct {
	AssociationID string `xml:"associationId"`
}

// AssociateDhcpOptionsRequest is undocumented.
type AssociateDhcpOptionsRequest struct {
	DhcpOptionsID string `xml:"DhcpOptionsId"`
	DryRun        bool   `xml:"dryRun"`
	VpcID         string `xml:"VpcId"`
}

// AssociateRouteTableRequest is undocumented.
type AssociateRouteTableRequest struct {
	DryRun       bool   `xml:"dryRun"`
	RouteTableID string `xml:"routeTableId"`
	SubnetID     string `xml:"subnetId"`
}

// AssociateRouteTableResult is undocumented.
type AssociateRouteTableResult struct {
	AssociationID string `xml:"associationId"`
}

// AttachInternetGatewayRequest is undocumented.
type AttachInternetGatewayRequest struct {
	DryRun            bool   `xml:"dryRun"`
	InternetGatewayID string `xml:"internetGatewayId"`
	VpcID             string `xml:"vpcId"`
}

// AttachNetworkInterfaceRequest is undocumented.
type AttachNetworkInterfaceRequest struct {
	DeviceIndex        int    `xml:"deviceIndex"`
	DryRun             bool   `xml:"dryRun"`
	InstanceID         string `xml:"instanceId"`
	NetworkInterfaceID string `xml:"networkInterfaceId"`
}

// AttachNetworkInterfaceResult is undocumented.
type AttachNetworkInterfaceResult struct {
	AttachmentID string `xml:"attachmentId"`
}

// AttachVolumeRequest is undocumented.
type AttachVolumeRequest struct {
	Device     string `xml:"Device"`
	DryRun     bool   `xml:"dryRun"`
	InstanceID string `xml:"InstanceId"`
	VolumeID   string `xml:"VolumeId"`
}

// AttachVpnGatewayRequest is undocumented.
type AttachVpnGatewayRequest struct {
	DryRun       bool   `xml:"dryRun"`
	VpcID        string `xml:"VpcId"`
	VpnGatewayID string `xml:"VpnGatewayId"`
}

// AttachVpnGatewayResult is undocumented.
type AttachVpnGatewayResult struct {
	VpcAttachment VpcAttachment `xml:"attachment"`
}

// AttributeBooleanValue is undocumented.
type AttributeBooleanValue struct {
	Value bool `xml:"value"`
}

// AttributeValue is undocumented.
type AttributeValue struct {
	Value string `xml:"value"`
}

// AuthorizeSecurityGroupEgressRequest is undocumented.
type AuthorizeSecurityGroupEgressRequest struct {
	CidrIP                     string         `xml:"cidrIp"`
	DryRun                     bool           `xml:"dryRun"`
	FromPort                   int            `xml:"fromPort"`
	GroupID                    string         `xml:"groupId"`
	IPPermissions              []IPPermission `xml:"ipPermissions>item"`
	IPProtocol                 string         `xml:"ipProtocol"`
	SourceSecurityGroupName    string         `xml:"sourceSecurityGroupName"`
	SourceSecurityGroupOwnerID string         `xml:"sourceSecurityGroupOwnerId"`
	ToPort                     int            `xml:"toPort"`
}

// AuthorizeSecurityGroupIngressRequest is undocumented.
type AuthorizeSecurityGroupIngressRequest struct {
	CidrIP                     string         `xml:"CidrIp"`
	DryRun                     bool           `xml:"dryRun"`
	FromPort                   int            `xml:"FromPort"`
	GroupID                    string         `xml:"GroupId"`
	GroupName                  string         `xml:"GroupName"`
	IPPermissions              []IPPermission `xml:"IpPermissions>item"`
	IPProtocol                 string         `xml:"IpProtocol"`
	SourceSecurityGroupName    string         `xml:"SourceSecurityGroupName"`
	SourceSecurityGroupOwnerID string         `xml:"SourceSecurityGroupOwnerId"`
	ToPort                     int            `xml:"ToPort"`
}

// AvailabilityZone is undocumented.
type AvailabilityZone struct {
	Messages   []AvailabilityZoneMessage `xml:"messageSet>item"`
	RegionName string                    `xml:"regionName"`
	State      string                    `xml:"zoneState"`
	ZoneName   string                    `xml:"zoneName"`
}

// AvailabilityZoneMessage is undocumented.
type AvailabilityZoneMessage struct {
	Message string `xml:"message"`
}

// BlobAttributeValue is undocumented.
type BlobAttributeValue struct {
	Value []byte `xml:"value"`
}

// BlockDeviceMapping is undocumented.
type BlockDeviceMapping struct {
	DeviceName  string         `xml:"deviceName"`
	Ebs         EbsBlockDevice `xml:"ebs"`
	NoDevice    string         `xml:"noDevice"`
	VirtualName string         `xml:"virtualName"`
}

// BundleInstanceRequest is undocumented.
type BundleInstanceRequest struct {
	DryRun     bool    `xml:"dryRun"`
	InstanceID string  `xml:"InstanceId"`
	Storage    Storage `xml:"Storage"`
}

// BundleInstanceResult is undocumented.
type BundleInstanceResult struct {
	BundleTask BundleTask `xml:"bundleInstanceTask"`
}

// BundleTask is undocumented.
type BundleTask struct {
	BundleID        string          `xml:"bundleId"`
	BundleTaskError BundleTaskError `xml:"error"`
	InstanceID      string          `xml:"instanceId"`
	Progress        string          `xml:"progress"`
	StartTime       time.Time       `xml:"startTime"`
	State           string          `xml:"state"`
	Storage         Storage         `xml:"storage"`
	UpdateTime      time.Time       `xml:"updateTime"`
}

// BundleTaskError is undocumented.
type BundleTaskError struct {
	Code    string `xml:"code"`
	Message string `xml:"message"`
}

// CancelBundleTaskRequest is undocumented.
type CancelBundleTaskRequest struct {
	BundleID string `xml:"BundleId"`
	DryRun   bool   `xml:"dryRun"`
}

// CancelBundleTaskResult is undocumented.
type CancelBundleTaskResult struct {
	BundleTask BundleTask `xml:"bundleInstanceTask"`
}

// CancelConversionRequest is undocumented.
type CancelConversionRequest struct {
	ConversionTaskID string `xml:"conversionTaskId"`
	DryRun           bool   `xml:"dryRun"`
	ReasonMessage    string `xml:"reasonMessage"`
}

// CancelExportTaskRequest is undocumented.
type CancelExportTaskRequest struct {
	ExportTaskID string `xml:"exportTaskId"`
}

// CancelReservedInstancesListingRequest is undocumented.
type CancelReservedInstancesListingRequest struct {
	ReservedInstancesListingID string `xml:"reservedInstancesListingId"`
}

// CancelReservedInstancesListingResult is undocumented.
type CancelReservedInstancesListingResult struct {
	ReservedInstancesListings []ReservedInstancesListing `xml:"reservedInstancesListingsSet>item"`
}

// CancelSpotInstanceRequestsRequest is undocumented.
type CancelSpotInstanceRequestsRequest struct {
	DryRun                 bool     `xml:"dryRun"`
	SpotInstanceRequestIds []string `xml:"SpotInstanceRequestId>SpotInstanceRequestId"`
}

// CancelSpotInstanceRequestsResult is undocumented.
type CancelSpotInstanceRequestsResult struct {
	CancelledSpotInstanceRequests []CancelledSpotInstanceRequest `xml:"spotInstanceRequestSet>item"`
}

// CancelledSpotInstanceRequest is undocumented.
type CancelledSpotInstanceRequest struct {
	SpotInstanceRequestID string `xml:"spotInstanceRequestId"`
	State                 string `xml:"state"`
}

// ConfirmProductInstanceRequest is undocumented.
type ConfirmProductInstanceRequest struct {
	DryRun      bool   `xml:"dryRun"`
	InstanceID  string `xml:"InstanceId"`
	ProductCode string `xml:"ProductCode"`
}

// ConfirmProductInstanceResult is undocumented.
type ConfirmProductInstanceResult struct {
	OwnerID string `xml:"ownerId"`
}

// ConversionTask is undocumented.
type ConversionTask struct {
	ConversionTaskID string                    `xml:"conversionTaskId"`
	ExpirationTime   string                    `xml:"expirationTime"`
	ImportInstance   ImportInstanceTaskDetails `xml:"importInstance"`
	ImportVolume     ImportVolumeTaskDetails   `xml:"importVolume"`
	State            string                    `xml:"state"`
	StatusMessage    string                    `xml:"statusMessage"`
	Tags             []Tag                     `xml:"tagSet>item"`
}

// CopyImageRequest is undocumented.
type CopyImageRequest struct {
	ClientToken   string `xml:"ClientToken"`
	Description   string `xml:"Description"`
	DryRun        bool   `xml:"dryRun"`
	Name          string `xml:"Name"`
	SourceImageID string `xml:"SourceImageId"`
	SourceRegion  string `xml:"SourceRegion"`
}

// CopyImageResult is undocumented.
type CopyImageResult struct {
	ImageID string `xml:"imageId"`
}

// CopySnapshotRequest is undocumented.
type CopySnapshotRequest struct {
	Description       string `xml:"Description"`
	DestinationRegion string `xml:"destinationRegion"`
	DryRun            bool   `xml:"dryRun"`
	PresignedURL      string `xml:"presignedUrl"`
	SourceRegion      string `xml:"SourceRegion"`
	SourceSnapshotID  string `xml:"SourceSnapshotId"`
}

// CopySnapshotResult is undocumented.
type CopySnapshotResult struct {
	SnapshotID string `xml:"snapshotId"`
}

// CreateCustomerGatewayRequest is undocumented.
type CreateCustomerGatewayRequest struct {
	BgpAsn   int    `xml:"BgpAsn"`
	DryRun   bool   `xml:"dryRun"`
	PublicIP string `xml:"IpAddress"`
	Type     string `xml:"Type"`
}

// CreateCustomerGatewayResult is undocumented.
type CreateCustomerGatewayResult struct {
	CustomerGateway CustomerGateway `xml:"customerGateway"`
}

// CreateDhcpOptionsRequest is undocumented.
type CreateDhcpOptionsRequest struct {
	DhcpConfigurations []NewDhcpConfiguration `xml:"dhcpConfiguration>item"`
	DryRun             bool                   `xml:"dryRun"`
}

// CreateDhcpOptionsResult is undocumented.
type CreateDhcpOptionsResult struct {
	DhcpOptions DhcpOptions `xml:"dhcpOptions"`
}

// CreateImageRequest is undocumented.
type CreateImageRequest struct {
	BlockDeviceMappings []BlockDeviceMapping `xml:"blockDeviceMapping>BlockDeviceMapping"`
	Description         string               `xml:"description"`
	DryRun              bool                 `xml:"dryRun"`
	InstanceID          string               `xml:"instanceId"`
	Name                string               `xml:"name"`
	NoReboot            bool                 `xml:"noReboot"`
}

// CreateImageResult is undocumented.
type CreateImageResult struct {
	ImageID string `xml:"imageId"`
}

// CreateInstanceExportTaskRequest is undocumented.
type CreateInstanceExportTaskRequest struct {
	Description       string                      `xml:"description"`
	ExportToS3Task    ExportToS3TaskSpecification `xml:"exportToS3"`
	InstanceID        string                      `xml:"instanceId"`
	TargetEnvironment string                      `xml:"targetEnvironment"`
}

// CreateInstanceExportTaskResult is undocumented.
type CreateInstanceExportTaskResult struct {
	ExportTask ExportTask `xml:"exportTask"`
}

// CreateInternetGatewayRequest is undocumented.
type CreateInternetGatewayRequest struct {
	DryRun bool `xml:"dryRun"`
}

// CreateInternetGatewayResult is undocumented.
type CreateInternetGatewayResult struct {
	InternetGateway InternetGateway `xml:"internetGateway"`
}

// CreateKeyPairRequest is undocumented.
type CreateKeyPairRequest struct {
	DryRun  bool   `xml:"dryRun"`
	KeyName string `xml:"KeyName"`
}

// CreateNetworkAclEntryRequest is undocumented.
type CreateNetworkAclEntryRequest struct {
	CidrBlock    string       `xml:"cidrBlock"`
	DryRun       bool         `xml:"dryRun"`
	Egress       bool         `xml:"egress"`
	IcmpTypeCode IcmpTypeCode `xml:"Icmp"`
	NetworkAclID string       `xml:"networkAclId"`
	PortRange    PortRange    `xml:"portRange"`
	Protocol     string       `xml:"protocol"`
	RuleAction   string       `xml:"ruleAction"`
	RuleNumber   int          `xml:"ruleNumber"`
}

// CreateNetworkAclRequest is undocumented.
type CreateNetworkAclRequest struct {
	DryRun bool   `xml:"dryRun"`
	VpcID  string `xml:"vpcId"`
}

// CreateNetworkAclResult is undocumented.
type CreateNetworkAclResult struct {
	NetworkAcl NetworkAcl `xml:"networkAcl"`
}

// CreateNetworkInterfaceRequest is undocumented.
type CreateNetworkInterfaceRequest struct {
	Description                    string                          `xml:"description"`
	DryRun                         bool                            `xml:"dryRun"`
	Groups                         []string                        `xml:"SecurityGroupId>SecurityGroupId"`
	PrivateIPAddress               string                          `xml:"privateIpAddress"`
	PrivateIPAddresses             []PrivateIPAddressSpecification `xml:"privateIpAddresses>item"`
	SecondaryPrivateIPAddressCount int                             `xml:"secondaryPrivateIpAddressCount"`
	SubnetID                       string                          `xml:"subnetId"`
}

// CreateNetworkInterfaceResult is undocumented.
type CreateNetworkInterfaceResult struct {
	NetworkInterface NetworkInterface `xml:"networkInterface"`
}

// CreatePlacementGroupRequest is undocumented.
type CreatePlacementGroupRequest struct {
	DryRun    bool   `xml:"dryRun"`
	GroupName string `xml:"groupName"`
	Strategy  string `xml:"strategy"`
}

// CreateReservedInstancesListingRequest is undocumented.
type CreateReservedInstancesListingRequest struct {
	ClientToken         string                       `xml:"clientToken"`
	InstanceCount       int                          `xml:"instanceCount"`
	PriceSchedules      []PriceScheduleSpecification `xml:"priceSchedules>item"`
	ReservedInstancesID string                       `xml:"reservedInstancesId"`
}

// CreateReservedInstancesListingResult is undocumented.
type CreateReservedInstancesListingResult struct {
	ReservedInstancesListings []ReservedInstancesListing `xml:"reservedInstancesListingsSet>item"`
}

// CreateRouteRequest is undocumented.
type CreateRouteRequest struct {
	DestinationCidrBlock   string `xml:"destinationCidrBlock"`
	DryRun                 bool   `xml:"dryRun"`
	GatewayID              string `xml:"gatewayId"`
	InstanceID             string `xml:"instanceId"`
	NetworkInterfaceID     string `xml:"networkInterfaceId"`
	RouteTableID           string `xml:"routeTableId"`
	VpcPeeringConnectionID string `xml:"vpcPeeringConnectionId"`
}

// CreateRouteTableRequest is undocumented.
type CreateRouteTableRequest struct {
	DryRun bool   `xml:"dryRun"`
	VpcID  string `xml:"vpcId"`
}

// CreateRouteTableResult is undocumented.
type CreateRouteTableResult struct {
	RouteTable RouteTable `xml:"routeTable"`
}

// CreateSecurityGroupRequest is undocumented.
type CreateSecurityGroupRequest struct {
	Description string `xml:"GroupDescription"`
	DryRun      bool   `xml:"dryRun"`
	GroupName   string `xml:"GroupName"`
	VpcID       string `xml:"VpcId"`
}

// CreateSecurityGroupResult is undocumented.
type CreateSecurityGroupResult struct {
	GroupID string `xml:"groupId"`
}

// CreateSnapshotRequest is undocumented.
type CreateSnapshotRequest struct {
	Description string `xml:"Description"`
	DryRun      bool   `xml:"dryRun"`
	VolumeID    string `xml:"VolumeId"`
}

// CreateSpotDatafeedSubscriptionRequest is undocumented.
type CreateSpotDatafeedSubscriptionRequest struct {
	Bucket string `xml:"bucket"`
	DryRun bool   `xml:"dryRun"`
	Prefix string `xml:"prefix"`
}

// CreateSpotDatafeedSubscriptionResult is undocumented.
type CreateSpotDatafeedSubscriptionResult struct {
	SpotDatafeedSubscription SpotDatafeedSubscription `xml:"spotDatafeedSubscription"`
}

// CreateSubnetRequest is undocumented.
type CreateSubnetRequest struct {
	AvailabilityZone string `xml:"AvailabilityZone"`
	CidrBlock        string `xml:"CidrBlock"`
	DryRun           bool   `xml:"dryRun"`
	VpcID            string `xml:"VpcId"`
}

// CreateSubnetResult is undocumented.
type CreateSubnetResult struct {
	Subnet Subnet `xml:"subnet"`
}

// CreateTagsRequest is undocumented.
type CreateTagsRequest struct {
	DryRun    bool     `xml:"dryRun"`
	Resources []string `xml:"ResourceId>member"`
	Tags      []Tag    `xml:"Tag>item"`
}

// CreateVolumePermission is undocumented.
type CreateVolumePermission struct {
	Group  string `xml:"group"`
	UserID string `xml:"userId"`
}

// CreateVolumePermissionModifications is undocumented.
type CreateVolumePermissionModifications struct {
	Add    []CreateVolumePermission `xml:"Add>item"`
	Remove []CreateVolumePermission `xml:"Remove>item"`
}

// CreateVolumeRequest is undocumented.
type CreateVolumeRequest struct {
	AvailabilityZone string `xml:"AvailabilityZone"`
	DryRun           bool   `xml:"dryRun"`
	Encrypted        bool   `xml:"encrypted"`
	Iops             int    `xml:"Iops"`
	KmsKeyID         string `xml:"KmsKeyId"`
	Size             int    `xml:"Size"`
	SnapshotID       string `xml:"SnapshotId"`
	VolumeType       string `xml:"VolumeType"`
}

// CreateVpcPeeringConnectionRequest is undocumented.
type CreateVpcPeeringConnectionRequest struct {
	DryRun      bool   `xml:"dryRun"`
	PeerOwnerID string `xml:"peerOwnerId"`
	PeerVpcID   string `xml:"peerVpcId"`
	VpcID       string `xml:"vpcId"`
}

// CreateVpcPeeringConnectionResult is undocumented.
type CreateVpcPeeringConnectionResult struct {
	VpcPeeringConnection VpcPeeringConnection `xml:"vpcPeeringConnection"`
}

// CreateVpcRequest is undocumented.
type CreateVpcRequest struct {
	CidrBlock       string `xml:"CidrBlock"`
	DryRun          bool   `xml:"dryRun"`
	InstanceTenancy string `xml:"instanceTenancy"`
}

// CreateVpcResult is undocumented.
type CreateVpcResult struct {
	Vpc Vpc `xml:"vpc"`
}

// CreateVpnConnectionRequest is undocumented.
type CreateVpnConnectionRequest struct {
	CustomerGatewayID string                            `xml:"CustomerGatewayId"`
	DryRun            bool                              `xml:"dryRun"`
	Options           VpnConnectionOptionsSpecification `xml:"options"`
	Type              string                            `xml:"Type"`
	VpnGatewayID      string                            `xml:"VpnGatewayId"`
}

// CreateVpnConnectionResult is undocumented.
type CreateVpnConnectionResult struct {
	VpnConnection VpnConnection `xml:"vpnConnection"`
}

// CreateVpnConnectionRouteRequest is undocumented.
type CreateVpnConnectionRouteRequest struct {
	DestinationCidrBlock string `xml:"DestinationCidrBlock"`
	VpnConnectionID      string `xml:"VpnConnectionId"`
}

// CreateVpnGatewayRequest is undocumented.
type CreateVpnGatewayRequest struct {
	AvailabilityZone string `xml:"AvailabilityZone"`
	DryRun           bool   `xml:"dryRun"`
	Type             string `xml:"Type"`
}

// CreateVpnGatewayResult is undocumented.
type CreateVpnGatewayResult struct {
	VpnGateway VpnGateway `xml:"vpnGateway"`
}

// CustomerGateway is undocumented.
type CustomerGateway struct {
	BgpAsn            string `xml:"bgpAsn"`
	CustomerGatewayID string `xml:"customerGatewayId"`
	IPAddress         string `xml:"ipAddress"`
	State             string `xml:"state"`
	Tags              []Tag  `xml:"tagSet>item"`
	Type              string `xml:"type"`
}

// DeleteCustomerGatewayRequest is undocumented.
type DeleteCustomerGatewayRequest struct {
	CustomerGatewayID string `xml:"CustomerGatewayId"`
	DryRun            bool   `xml:"dryRun"`
}

// DeleteDhcpOptionsRequest is undocumented.
type DeleteDhcpOptionsRequest struct {
	DhcpOptionsID string `xml:"DhcpOptionsId"`
	DryRun        bool   `xml:"dryRun"`
}

// DeleteInternetGatewayRequest is undocumented.
type DeleteInternetGatewayRequest struct {
	DryRun            bool   `xml:"dryRun"`
	InternetGatewayID string `xml:"internetGatewayId"`
}

// DeleteKeyPairRequest is undocumented.
type DeleteKeyPairRequest struct {
	DryRun  bool   `xml:"dryRun"`
	KeyName string `xml:"KeyName"`
}

// DeleteNetworkAclEntryRequest is undocumented.
type DeleteNetworkAclEntryRequest struct {
	DryRun       bool   `xml:"dryRun"`
	Egress       bool   `xml:"egress"`
	NetworkAclID string `xml:"networkAclId"`
	RuleNumber   int    `xml:"ruleNumber"`
}

// DeleteNetworkAclRequest is undocumented.
type DeleteNetworkAclRequest struct {
	DryRun       bool   `xml:"dryRun"`
	NetworkAclID string `xml:"networkAclId"`
}

// DeleteNetworkInterfaceRequest is undocumented.
type DeleteNetworkInterfaceRequest struct {
	DryRun             bool   `xml:"dryRun"`
	NetworkInterfaceID string `xml:"networkInterfaceId"`
}

// DeletePlacementGroupRequest is undocumented.
type DeletePlacementGroupRequest struct {
	DryRun    bool   `xml:"dryRun"`
	GroupName string `xml:"groupName"`
}

// DeleteRouteRequest is undocumented.
type DeleteRouteRequest struct {
	DestinationCidrBlock string `xml:"destinationCidrBlock"`
	DryRun               bool   `xml:"dryRun"`
	RouteTableID         string `xml:"routeTableId"`
}

// DeleteRouteTableRequest is undocumented.
type DeleteRouteTableRequest struct {
	DryRun       bool   `xml:"dryRun"`
	RouteTableID string `xml:"routeTableId"`
}

// DeleteSecurityGroupRequest is undocumented.
type DeleteSecurityGroupRequest struct {
	DryRun    bool   `xml:"dryRun"`
	GroupID   string `xml:"GroupId"`
	GroupName string `xml:"GroupName"`
}

// DeleteSnapshotRequest is undocumented.
type DeleteSnapshotRequest struct {
	DryRun     bool   `xml:"dryRun"`
	SnapshotID string `xml:"SnapshotId"`
}

// DeleteSpotDatafeedSubscriptionRequest is undocumented.
type DeleteSpotDatafeedSubscriptionRequest struct {
	DryRun bool `xml:"dryRun"`
}

// DeleteSubnetRequest is undocumented.
type DeleteSubnetRequest struct {
	DryRun   bool   `xml:"dryRun"`
	SubnetID string `xml:"SubnetId"`
}

// DeleteTagsRequest is undocumented.
type DeleteTagsRequest struct {
	DryRun    bool     `xml:"dryRun"`
	Resources []string `xml:"resourceId>member"`
	Tags      []Tag    `xml:"tag>item"`
}

// DeleteVolumeRequest is undocumented.
type DeleteVolumeRequest struct {
	DryRun   bool   `xml:"dryRun"`
	VolumeID string `xml:"VolumeId"`
}

// DeleteVpcPeeringConnectionRequest is undocumented.
type DeleteVpcPeeringConnectionRequest struct {
	DryRun                 bool   `xml:"dryRun"`
	VpcPeeringConnectionID string `xml:"vpcPeeringConnectionId"`
}

// DeleteVpcPeeringConnectionResult is undocumented.
type DeleteVpcPeeringConnectionResult struct {
	Return bool `xml:"return"`
}

// DeleteVpcRequest is undocumented.
type DeleteVpcRequest struct {
	DryRun bool   `xml:"dryRun"`
	VpcID  string `xml:"VpcId"`
}

// DeleteVpnConnectionRequest is undocumented.
type DeleteVpnConnectionRequest struct {
	DryRun          bool   `xml:"dryRun"`
	VpnConnectionID string `xml:"VpnConnectionId"`
}

// DeleteVpnConnectionRouteRequest is undocumented.
type DeleteVpnConnectionRouteRequest struct {
	DestinationCidrBlock string `xml:"DestinationCidrBlock"`
	VpnConnectionID      string `xml:"VpnConnectionId"`
}

// DeleteVpnGatewayRequest is undocumented.
type DeleteVpnGatewayRequest struct {
	DryRun       bool   `xml:"dryRun"`
	VpnGatewayID string `xml:"VpnGatewayId"`
}

// DeregisterImageRequest is undocumented.
type DeregisterImageRequest struct {
	DryRun  bool   `xml:"dryRun"`
	ImageID string `xml:"ImageId"`
}

// DescribeAccountAttributesRequest is undocumented.
type DescribeAccountAttributesRequest struct {
	AttributeNames []string `xml:"attributeName>attributeName"`
	DryRun         bool     `xml:"dryRun"`
}

// DescribeAccountAttributesResult is undocumented.
type DescribeAccountAttributesResult struct {
	AccountAttributes []AccountAttribute `xml:"accountAttributeSet>item"`
}

// DescribeAddressesRequest is undocumented.
type DescribeAddressesRequest struct {
	AllocationIds []string `xml:"AllocationId>AllocationId"`
	DryRun        bool     `xml:"dryRun"`
	Filters       []Filter `xml:"Filter>Filter"`
	PublicIPs     []string `xml:"PublicIp>PublicIp"`
}

// DescribeAddressesResult is undocumented.
type DescribeAddressesResult struct {
	Addresses []Address `xml:"addressesSet>item"`
}

// DescribeAvailabilityZonesRequest is undocumented.
type DescribeAvailabilityZonesRequest struct {
	DryRun    bool     `xml:"dryRun"`
	Filters   []Filter `xml:"Filter>Filter"`
	ZoneNames []string `xml:"ZoneName>ZoneName"`
}

// DescribeAvailabilityZonesResult is undocumented.
type DescribeAvailabilityZonesResult struct {
	AvailabilityZones []AvailabilityZone `xml:"availabilityZoneInfo>item"`
}

// DescribeBundleTasksRequest is undocumented.
type DescribeBundleTasksRequest struct {
	BundleIds []string `xml:"BundleId>BundleId"`
	DryRun    bool     `xml:"dryRun"`
	Filters   []Filter `xml:"Filter>Filter"`
}

// DescribeBundleTasksResult is undocumented.
type DescribeBundleTasksResult struct {
	BundleTasks []BundleTask `xml:"bundleInstanceTasksSet>item"`
}

// DescribeConversionTasksRequest is undocumented.
type DescribeConversionTasksRequest struct {
	ConversionTaskIds []string `xml:"conversionTaskId>item"`
	DryRun            bool     `xml:"dryRun"`
	Filters           []Filter `xml:"filter>Filter"`
}

// DescribeConversionTasksResult is undocumented.
type DescribeConversionTasksResult struct {
	ConversionTasks []ConversionTask `xml:"conversionTasks>item"`
}

// DescribeCustomerGatewaysRequest is undocumented.
type DescribeCustomerGatewaysRequest struct {
	CustomerGatewayIds []string `xml:"CustomerGatewayId>CustomerGatewayId"`
	DryRun             bool     `xml:"dryRun"`
	Filters            []Filter `xml:"Filter>Filter"`
}

// DescribeCustomerGatewaysResult is undocumented.
type DescribeCustomerGatewaysResult struct {
	CustomerGateways []CustomerGateway `xml:"customerGatewaySet>item"`
}

// DescribeDhcpOptionsRequest is undocumented.
type DescribeDhcpOptionsRequest struct {
	DhcpOptionsIds []string `xml:"DhcpOptionsId>DhcpOptionsId"`
	DryRun         bool     `xml:"dryRun"`
	Filters        []Filter `xml:"Filter>Filter"`
}

// DescribeDhcpOptionsResult is undocumented.
type DescribeDhcpOptionsResult struct {
	DhcpOptions []DhcpOptions `xml:"dhcpOptionsSet>item"`
}

// DescribeExportTasksRequest is undocumented.
type DescribeExportTasksRequest struct {
	ExportTaskIds []string `xml:"exportTaskId>ExportTaskId"`
}

// DescribeExportTasksResult is undocumented.
type DescribeExportTasksResult struct {
	ExportTasks []ExportTask `xml:"exportTaskSet>item"`
}

// DescribeImageAttributeRequest is undocumented.
type DescribeImageAttributeRequest struct {
	Attribute string `xml:"Attribute"`
	DryRun    bool   `xml:"dryRun"`
	ImageID   string `xml:"ImageId"`
}

// DescribeImagesRequest is undocumented.
type DescribeImagesRequest struct {
	DryRun          bool     `xml:"dryRun"`
	ExecutableUsers []string `xml:"ExecutableBy>ExecutableBy"`
	Filters         []Filter `xml:"Filter>Filter"`
	ImageIds        []string `xml:"ImageId>ImageId"`
	Owners          []string `xml:"Owner>Owner"`
}

// DescribeImagesResult is undocumented.
type DescribeImagesResult struct {
	Images []Image `xml:"imagesSet>item"`
}

// DescribeInstanceAttributeRequest is undocumented.
type DescribeInstanceAttributeRequest struct {
	Attribute  string `xml:"attribute"`
	DryRun     bool   `xml:"dryRun"`
	InstanceID string `xml:"instanceId"`
}

// DescribeInstanceStatusRequest is undocumented.
type DescribeInstanceStatusRequest struct {
	DryRun              bool     `xml:"dryRun"`
	Filters             []Filter `xml:"Filter>Filter"`
	IncludeAllInstances bool     `xml:"includeAllInstances"`
	InstanceIds         []string `xml:"InstanceId>InstanceId"`
	MaxResults          int      `xml:"MaxResults"`
	NextToken           string   `xml:"NextToken"`
}

// DescribeInstanceStatusResult is undocumented.
type DescribeInstanceStatusResult struct {
	InstanceStatuses []InstanceStatus `xml:"instanceStatusSet>item"`
	NextToken        string           `xml:"nextToken"`
}

// DescribeInstancesRequest is undocumented.
type DescribeInstancesRequest struct {
	DryRun      bool     `xml:"dryRun"`
	Filters     []Filter `xml:"Filter>Filter"`
	InstanceIds []string `xml:"InstanceId>InstanceId"`
	MaxResults  int      `xml:"maxResults"`
	NextToken   string   `xml:"nextToken"`
}

// DescribeInstancesResult is undocumented.
type DescribeInstancesResult struct {
	NextToken    string        `xml:"nextToken"`
	Reservations []Reservation `xml:"reservationSet>item"`
}

// DescribeInternetGatewaysRequest is undocumented.
type DescribeInternetGatewaysRequest struct {
	DryRun             bool     `xml:"dryRun"`
	Filters            []Filter `xml:"Filter>Filter"`
	InternetGatewayIds []string `xml:"internetGatewayId>item"`
}

// DescribeInternetGatewaysResult is undocumented.
type DescribeInternetGatewaysResult struct {
	InternetGateways []InternetGateway `xml:"internetGatewaySet>item"`
}

// DescribeKeyPairsRequest is undocumented.
type DescribeKeyPairsRequest struct {
	DryRun   bool     `xml:"dryRun"`
	Filters  []Filter `xml:"Filter>Filter"`
	KeyNames []string `xml:"KeyName>KeyName"`
}

// DescribeKeyPairsResult is undocumented.
type DescribeKeyPairsResult struct {
	KeyPairs []KeyPairInfo `xml:"keySet>item"`
}

// DescribeNetworkAclsRequest is undocumented.
type DescribeNetworkAclsRequest struct {
	DryRun        bool     `xml:"dryRun"`
	Filters       []Filter `xml:"Filter>Filter"`
	NetworkAclIds []string `xml:"NetworkAclId>item"`
}

// DescribeNetworkAclsResult is undocumented.
type DescribeNetworkAclsResult struct {
	NetworkAcls []NetworkAcl `xml:"networkAclSet>item"`
}

// DescribeNetworkInterfaceAttributeRequest is undocumented.
type DescribeNetworkInterfaceAttributeRequest struct {
	Attribute          string `xml:"attribute"`
	DryRun             bool   `xml:"dryRun"`
	NetworkInterfaceID string `xml:"networkInterfaceId"`
}

// DescribeNetworkInterfaceAttributeResult is undocumented.
type DescribeNetworkInterfaceAttributeResult struct {
	Attachment         NetworkInterfaceAttachment `xml:"attachment"`
	Description        AttributeValue             `xml:"description"`
	Groups             []GroupIdentifier          `xml:"groupSet>item"`
	NetworkInterfaceID string                     `xml:"networkInterfaceId"`
	SourceDestCheck    AttributeBooleanValue      `xml:"sourceDestCheck"`
}

// DescribeNetworkInterfacesRequest is undocumented.
type DescribeNetworkInterfacesRequest struct {
	DryRun              bool     `xml:"dryRun"`
	Filters             []Filter `xml:"filter>Filter"`
	NetworkInterfaceIds []string `xml:"NetworkInterfaceId>item"`
}

// DescribeNetworkInterfacesResult is undocumented.
type DescribeNetworkInterfacesResult struct {
	NetworkInterfaces []NetworkInterface `xml:"networkInterfaceSet>item"`
}

// DescribePlacementGroupsRequest is undocumented.
type DescribePlacementGroupsRequest struct {
	DryRun     bool     `xml:"dryRun"`
	Filters    []Filter `xml:"Filter>Filter"`
	GroupNames []string `xml:"groupName>member"`
}

// DescribePlacementGroupsResult is undocumented.
type DescribePlacementGroupsResult struct {
	PlacementGroups []PlacementGroup `xml:"placementGroupSet>item"`
}

// DescribeRegionsRequest is undocumented.
type DescribeRegionsRequest struct {
	DryRun      bool     `xml:"dryRun"`
	Filters     []Filter `xml:"Filter>Filter"`
	RegionNames []string `xml:"RegionName>RegionName"`
}

// DescribeRegionsResult is undocumented.
type DescribeRegionsResult struct {
	Regions []Region `xml:"regionInfo>item"`
}

// DescribeReservedInstancesListingsRequest is undocumented.
type DescribeReservedInstancesListingsRequest struct {
	Filters                    []Filter `xml:"filters>Filter"`
	ReservedInstancesID        string   `xml:"reservedInstancesId"`
	ReservedInstancesListingID string   `xml:"reservedInstancesListingId"`
}

// DescribeReservedInstancesListingsResult is undocumented.
type DescribeReservedInstancesListingsResult struct {
	ReservedInstancesListings []ReservedInstancesListing `xml:"reservedInstancesListingsSet>item"`
}

// DescribeReservedInstancesModificationsRequest is undocumented.
type DescribeReservedInstancesModificationsRequest struct {
	Filters                          []Filter `xml:"Filter>Filter"`
	NextToken                        string   `xml:"nextToken"`
	ReservedInstancesModificationIds []string `xml:"ReservedInstancesModificationId>ReservedInstancesModificationId"`
}

// DescribeReservedInstancesModificationsResult is undocumented.
type DescribeReservedInstancesModificationsResult struct {
	NextToken                      string                          `xml:"nextToken"`
	ReservedInstancesModifications []ReservedInstancesModification `xml:"reservedInstancesModificationsSet>item"`
}

// DescribeReservedInstancesOfferingsRequest is undocumented.
type DescribeReservedInstancesOfferingsRequest struct {
	AvailabilityZone             string   `xml:"AvailabilityZone"`
	DryRun                       bool     `xml:"dryRun"`
	Filters                      []Filter `xml:"Filter>Filter"`
	IncludeMarketplace           bool     `xml:"IncludeMarketplace"`
	InstanceTenancy              string   `xml:"instanceTenancy"`
	InstanceType                 string   `xml:"InstanceType"`
	MaxDuration                  int64    `xml:"MaxDuration"`
	MaxInstanceCount             int      `xml:"MaxInstanceCount"`
	MaxResults                   int      `xml:"maxResults"`
	MinDuration                  int64    `xml:"MinDuration"`
	NextToken                    string   `xml:"nextToken"`
	OfferingType                 string   `xml:"offeringType"`
	ProductDescription           string   `xml:"ProductDescription"`
	ReservedInstancesOfferingIds []string `xml:"ReservedInstancesOfferingId>member"`
}

// DescribeReservedInstancesOfferingsResult is undocumented.
type DescribeReservedInstancesOfferingsResult struct {
	NextToken                  string                      `xml:"nextToken"`
	ReservedInstancesOfferings []ReservedInstancesOffering `xml:"reservedInstancesOfferingsSet>item"`
}

// DescribeReservedInstancesRequest is undocumented.
type DescribeReservedInstancesRequest struct {
	DryRun               bool     `xml:"dryRun"`
	Filters              []Filter `xml:"Filter>Filter"`
	OfferingType         string   `xml:"offeringType"`
	ReservedInstancesIds []string `xml:"ReservedInstancesId>ReservedInstancesId"`
}

// DescribeReservedInstancesResult is undocumented.
type DescribeReservedInstancesResult struct {
	ReservedInstances []ReservedInstances `xml:"reservedInstancesSet>item"`
}

// DescribeRouteTablesRequest is undocumented.
type DescribeRouteTablesRequest struct {
	DryRun        bool     `xml:"dryRun"`
	Filters       []Filter `xml:"Filter>Filter"`
	RouteTableIds []string `xml:"RouteTableId>item"`
}

// DescribeRouteTablesResult is undocumented.
type DescribeRouteTablesResult struct {
	RouteTables []RouteTable `xml:"routeTableSet>item"`
}

// DescribeSecurityGroupsRequest is undocumented.
type DescribeSecurityGroupsRequest struct {
	DryRun     bool     `xml:"dryRun"`
	Filters    []Filter `xml:"Filter>Filter"`
	GroupIds   []string `xml:"GroupId>groupId"`
	GroupNames []string `xml:"GroupName>GroupName"`
}

// DescribeSecurityGroupsResult is undocumented.
type DescribeSecurityGroupsResult struct {
	SecurityGroups []SecurityGroup `xml:"securityGroupInfo>item"`
}

// DescribeSnapshotAttributeRequest is undocumented.
type DescribeSnapshotAttributeRequest struct {
	Attribute  string `xml:"Attribute"`
	DryRun     bool   `xml:"dryRun"`
	SnapshotID string `xml:"SnapshotId"`
}

// DescribeSnapshotAttributeResult is undocumented.
type DescribeSnapshotAttributeResult struct {
	CreateVolumePermissions []CreateVolumePermission `xml:"createVolumePermission>item"`
	ProductCodes            []ProductCode            `xml:"productCodes>item"`
	SnapshotID              string                   `xml:"snapshotId"`
}

// DescribeSnapshotsRequest is undocumented.
type DescribeSnapshotsRequest struct {
	DryRun              bool     `xml:"dryRun"`
	Filters             []Filter `xml:"Filter>Filter"`
	OwnerIds            []string `xml:"Owner>Owner"`
	RestorableByUserIds []string `xml:"RestorableBy>member"`
	SnapshotIds         []string `xml:"SnapshotId>SnapshotId"`
}

// DescribeSnapshotsResult is undocumented.
type DescribeSnapshotsResult struct {
	Snapshots []Snapshot `xml:"snapshotSet>item"`
}

// DescribeSpotDatafeedSubscriptionRequest is undocumented.
type DescribeSpotDatafeedSubscriptionRequest struct {
	DryRun bool `xml:"dryRun"`
}

// DescribeSpotDatafeedSubscriptionResult is undocumented.
type DescribeSpotDatafeedSubscriptionResult struct {
	SpotDatafeedSubscription SpotDatafeedSubscription `xml:"spotDatafeedSubscription"`
}

// DescribeSpotInstanceRequestsRequest is undocumented.
type DescribeSpotInstanceRequestsRequest struct {
	DryRun                 bool     `xml:"dryRun"`
	Filters                []Filter `xml:"Filter>Filter"`
	SpotInstanceRequestIds []string `xml:"SpotInstanceRequestId>SpotInstanceRequestId"`
}

// DescribeSpotInstanceRequestsResult is undocumented.
type DescribeSpotInstanceRequestsResult struct {
	SpotInstanceRequests []SpotInstanceRequest `xml:"spotInstanceRequestSet>item"`
}

// DescribeSpotPriceHistoryRequest is undocumented.
type DescribeSpotPriceHistoryRequest struct {
	AvailabilityZone    string    `xml:"availabilityZone"`
	DryRun              bool      `xml:"dryRun"`
	EndTime             time.Time `xml:"endTime"`
	Filters             []Filter  `xml:"Filter>Filter"`
	InstanceTypes       []string  `xml:"InstanceType>member"`
	MaxResults          int       `xml:"maxResults"`
	NextToken           string    `xml:"nextToken"`
	ProductDescriptions []string  `xml:"ProductDescription>member"`
	StartTime           time.Time `xml:"startTime"`
}

// DescribeSpotPriceHistoryResult is undocumented.
type DescribeSpotPriceHistoryResult struct {
	NextToken        string      `xml:"nextToken"`
	SpotPriceHistory []SpotPrice `xml:"spotPriceHistorySet>item"`
}

// DescribeSubnetsRequest is undocumented.
type DescribeSubnetsRequest struct {
	DryRun    bool     `xml:"dryRun"`
	Filters   []Filter `xml:"Filter>Filter"`
	SubnetIds []string `xml:"SubnetId>SubnetId"`
}

// DescribeSubnetsResult is undocumented.
type DescribeSubnetsResult struct {
	Subnets []Subnet `xml:"subnetSet>item"`
}

// DescribeTagsRequest is undocumented.
type DescribeTagsRequest struct {
	DryRun     bool     `xml:"dryRun"`
	Filters    []Filter `xml:"Filter>Filter"`
	MaxResults int      `xml:"maxResults"`
	NextToken  string   `xml:"nextToken"`
}

// DescribeTagsResult is undocumented.
type DescribeTagsResult struct {
	NextToken string           `xml:"nextToken"`
	Tags      []TagDescription `xml:"tagSet>item"`
}

// DescribeVolumeAttributeRequest is undocumented.
type DescribeVolumeAttributeRequest struct {
	Attribute string `xml:"Attribute"`
	DryRun    bool   `xml:"dryRun"`
	VolumeID  string `xml:"VolumeId"`
}

// DescribeVolumeAttributeResult is undocumented.
type DescribeVolumeAttributeResult struct {
	AutoEnableIO AttributeBooleanValue `xml:"autoEnableIO"`
	ProductCodes []ProductCode         `xml:"productCodes>item"`
	VolumeID     string                `xml:"volumeId"`
}

// DescribeVolumeStatusRequest is undocumented.
type DescribeVolumeStatusRequest struct {
	DryRun     bool     `xml:"dryRun"`
	Filters    []Filter `xml:"Filter>Filter"`
	MaxResults int      `xml:"MaxResults"`
	NextToken  string   `xml:"NextToken"`
	VolumeIds  []string `xml:"VolumeId>VolumeId"`
}

// DescribeVolumeStatusResult is undocumented.
type DescribeVolumeStatusResult struct {
	NextToken      string             `xml:"nextToken"`
	VolumeStatuses []VolumeStatusItem `xml:"volumeStatusSet>item"`
}

// DescribeVolumesRequest is undocumented.
type DescribeVolumesRequest struct {
	DryRun     bool     `xml:"dryRun"`
	Filters    []Filter `xml:"Filter>Filter"`
	MaxResults int      `xml:"maxResults"`
	NextToken  string   `xml:"nextToken"`
	VolumeIds  []string `xml:"VolumeId>VolumeId"`
}

// DescribeVolumesResult is undocumented.
type DescribeVolumesResult struct {
	NextToken string   `xml:"nextToken"`
	Volumes   []Volume `xml:"volumeSet>item"`
}

// DescribeVpcAttributeRequest is undocumented.
type DescribeVpcAttributeRequest struct {
	Attribute string `xml:"Attribute"`
	DryRun    bool   `xml:"dryRun"`
	VpcID     string `xml:"VpcId"`
}

// DescribeVpcAttributeResult is undocumented.
type DescribeVpcAttributeResult struct {
	EnableDNSHostnames AttributeBooleanValue `xml:"enableDnsHostnames"`
	EnableDNSSupport   AttributeBooleanValue `xml:"enableDnsSupport"`
	VpcID              string                `xml:"vpcId"`
}

// DescribeVpcPeeringConnectionsRequest is undocumented.
type DescribeVpcPeeringConnectionsRequest struct {
	DryRun                  bool     `xml:"dryRun"`
	Filters                 []Filter `xml:"Filter>Filter"`
	VpcPeeringConnectionIds []string `xml:"VpcPeeringConnectionId>item"`
}

// DescribeVpcPeeringConnectionsResult is undocumented.
type DescribeVpcPeeringConnectionsResult struct {
	VpcPeeringConnections []VpcPeeringConnection `xml:"vpcPeeringConnectionSet>item"`
}

// DescribeVpcsRequest is undocumented.
type DescribeVpcsRequest struct {
	DryRun  bool     `xml:"dryRun"`
	Filters []Filter `xml:"Filter>Filter"`
	VpcIds  []string `xml:"VpcId>VpcId"`
}

// DescribeVpcsResult is undocumented.
type DescribeVpcsResult struct {
	Vpcs []Vpc `xml:"vpcSet>item"`
}

// DescribeVpnConnectionsRequest is undocumented.
type DescribeVpnConnectionsRequest struct {
	DryRun           bool     `xml:"dryRun"`
	Filters          []Filter `xml:"Filter>Filter"`
	VpnConnectionIds []string `xml:"VpnConnectionId>VpnConnectionId"`
}

// DescribeVpnConnectionsResult is undocumented.
type DescribeVpnConnectionsResult struct {
	VpnConnections []VpnConnection `xml:"vpnConnectionSet>item"`
}

// DescribeVpnGatewaysRequest is undocumented.
type DescribeVpnGatewaysRequest struct {
	DryRun        bool     `xml:"dryRun"`
	Filters       []Filter `xml:"Filter>Filter"`
	VpnGatewayIds []string `xml:"VpnGatewayId>VpnGatewayId"`
}

// DescribeVpnGatewaysResult is undocumented.
type DescribeVpnGatewaysResult struct {
	VpnGateways []VpnGateway `xml:"vpnGatewaySet>item"`
}

// DetachInternetGatewayRequest is undocumented.
type DetachInternetGatewayRequest struct {
	DryRun            bool   `xml:"dryRun"`
	InternetGatewayID string `xml:"internetGatewayId"`
	VpcID             string `xml:"vpcId"`
}

// DetachNetworkInterfaceRequest is undocumented.
type DetachNetworkInterfaceRequest struct {
	AttachmentID string `xml:"attachmentId"`
	DryRun       bool   `xml:"dryRun"`
	Force        bool   `xml:"force"`
}

// DetachVolumeRequest is undocumented.
type DetachVolumeRequest struct {
	Device     string `xml:"Device"`
	DryRun     bool   `xml:"dryRun"`
	Force      bool   `xml:"Force"`
	InstanceID string `xml:"InstanceId"`
	VolumeID   string `xml:"VolumeId"`
}

// DetachVpnGatewayRequest is undocumented.
type DetachVpnGatewayRequest struct {
	DryRun       bool   `xml:"dryRun"`
	VpcID        string `xml:"VpcId"`
	VpnGatewayID string `xml:"VpnGatewayId"`
}

// DhcpConfiguration is undocumented.
type DhcpConfiguration struct {
	Key    string           `xml:"key"`
	Values []AttributeValue `xml:"valueSet>item"`
}

// DhcpOptions is undocumented.
type DhcpOptions struct {
	DhcpConfigurations []DhcpConfiguration `xml:"dhcpConfigurationSet>item"`
	DhcpOptionsID      string              `xml:"dhcpOptionsId"`
	Tags               []Tag               `xml:"tagSet>item"`
}

// DisableVgwRoutePropagationRequest is undocumented.
type DisableVgwRoutePropagationRequest struct {
	GatewayID    string `xml:"GatewayId"`
	RouteTableID string `xml:"RouteTableId"`
}

// DisassociateAddressRequest is undocumented.
type DisassociateAddressRequest struct {
	AssociationID string `xml:"AssociationId"`
	DryRun        bool   `xml:"dryRun"`
	PublicIP      string `xml:"PublicIp"`
}

// DisassociateRouteTableRequest is undocumented.
type DisassociateRouteTableRequest struct {
	AssociationID string `xml:"associationId"`
	DryRun        bool   `xml:"dryRun"`
}

// DiskImage is undocumented.
type DiskImage struct {
	Description string          `xml:"Description"`
	Image       DiskImageDetail `xml:"Image"`
	Volume      VolumeDetail    `xml:"Volume"`
}

// DiskImageDescription is undocumented.
type DiskImageDescription struct {
	Checksum          string `xml:"checksum"`
	Format            string `xml:"format"`
	ImportManifestURL string `xml:"importManifestUrl"`
	Size              int64  `xml:"size"`
}

// DiskImageDetail is undocumented.
type DiskImageDetail struct {
	Bytes             int64  `xml:"bytes"`
	Format            string `xml:"format"`
	ImportManifestURL string `xml:"importManifestUrl"`
}

// DiskImageVolumeDescription is undocumented.
type DiskImageVolumeDescription struct {
	ID   string `xml:"id"`
	Size int64  `xml:"size"`
}

// EbsBlockDevice is undocumented.
type EbsBlockDevice struct {
	DeleteOnTermination bool   `xml:"deleteOnTermination"`
	Encrypted           bool   `xml:"encrypted"`
	Iops                int    `xml:"iops"`
	SnapshotID          string `xml:"snapshotId"`
	VolumeSize          int    `xml:"volumeSize"`
	VolumeType          string `xml:"volumeType"`
}

// EbsInstanceBlockDevice is undocumented.
type EbsInstanceBlockDevice struct {
	AttachTime          time.Time `xml:"attachTime"`
	DeleteOnTermination bool      `xml:"deleteOnTermination"`
	Status              string    `xml:"status"`
	VolumeID            string    `xml:"volumeId"`
}

// EbsInstanceBlockDeviceSpecification is undocumented.
type EbsInstanceBlockDeviceSpecification struct {
	DeleteOnTermination bool   `xml:"deleteOnTermination"`
	VolumeID            string `xml:"volumeId"`
}

// EnableVgwRoutePropagationRequest is undocumented.
type EnableVgwRoutePropagationRequest struct {
	GatewayID    string `xml:"GatewayId"`
	RouteTableID string `xml:"RouteTableId"`
}

// EnableVolumeIORequest is undocumented.
type EnableVolumeIORequest struct {
	DryRun   bool   `xml:"dryRun"`
	VolumeID string `xml:"volumeId"`
}

// ExportTask is undocumented.
type ExportTask struct {
	Description           string                `xml:"description"`
	ExportTaskID          string                `xml:"exportTaskId"`
	ExportToS3Task        ExportToS3Task        `xml:"exportToS3"`
	InstanceExportDetails InstanceExportDetails `xml:"instanceExport"`
	State                 string                `xml:"state"`
	StatusMessage         string                `xml:"statusMessage"`
}

// ExportToS3Task is undocumented.
type ExportToS3Task struct {
	ContainerFormat string `xml:"containerFormat"`
	DiskImageFormat string `xml:"diskImageFormat"`
	S3Bucket        string `xml:"s3Bucket"`
	S3Key           string `xml:"s3Key"`
}

// ExportToS3TaskSpecification is undocumented.
type ExportToS3TaskSpecification struct {
	ContainerFormat string `xml:"containerFormat"`
	DiskImageFormat string `xml:"diskImageFormat"`
	S3Bucket        string `xml:"s3Bucket"`
	S3Prefix        string `xml:"s3Prefix"`
}

// Filter is undocumented.
type Filter struct {
	Name   string   `xml:"Name"`
	Values []string `xml:"Value>item"`
}

// GetConsoleOutputRequest is undocumented.
type GetConsoleOutputRequest struct {
	DryRun     bool   `xml:"dryRun"`
	InstanceID string `xml:"InstanceId"`
}

// GetConsoleOutputResult is undocumented.
type GetConsoleOutputResult struct {
	InstanceID string    `xml:"instanceId"`
	Output     string    `xml:"output"`
	Timestamp  time.Time `xml:"timestamp"`
}

// GetPasswordDataRequest is undocumented.
type GetPasswordDataRequest struct {
	DryRun     bool   `xml:"dryRun"`
	InstanceID string `xml:"InstanceId"`
}

// GetPasswordDataResult is undocumented.
type GetPasswordDataResult struct {
	InstanceID   string    `xml:"instanceId"`
	PasswordData string    `xml:"passwordData"`
	Timestamp    time.Time `xml:"timestamp"`
}

// GroupIdentifier is undocumented.
type GroupIdentifier struct {
	GroupID   string `xml:"groupId"`
	GroupName string `xml:"groupName"`
}

// IamInstanceProfile is undocumented.
type IamInstanceProfile struct {
	ARN string `xml:"arn"`
	ID  string `xml:"id"`
}

// IamInstanceProfileSpecification is undocumented.
type IamInstanceProfileSpecification struct {
	ARN  string `xml:"arn"`
	Name string `xml:"name"`
}

// IcmpTypeCode is undocumented.
type IcmpTypeCode struct {
	Code int `xml:"code"`
	Type int `xml:"type"`
}

// Image is undocumented.
type Image struct {
	Architecture        string               `xml:"architecture"`
	BlockDeviceMappings []BlockDeviceMapping `xml:"blockDeviceMapping>item"`
	Description         string               `xml:"description"`
	Hypervisor          string               `xml:"hypervisor"`
	ImageID             string               `xml:"imageId"`
	ImageLocation       string               `xml:"imageLocation"`
	ImageOwnerAlias     string               `xml:"imageOwnerAlias"`
	ImageType           string               `xml:"imageType"`
	KernelID            string               `xml:"kernelId"`
	Name                string               `xml:"name"`
	OwnerID             string               `xml:"imageOwnerId"`
	Platform            string               `xml:"platform"`
	ProductCodes        []ProductCode        `xml:"productCodes>item"`
	Public              bool                 `xml:"isPublic"`
	RamdiskID           string               `xml:"ramdiskId"`
	RootDeviceName      string               `xml:"rootDeviceName"`
	RootDeviceType      string               `xml:"rootDeviceType"`
	SriovNetSupport     string               `xml:"sriovNetSupport"`
	State               string               `xml:"imageState"`
	StateReason         StateReason          `xml:"stateReason"`
	Tags                []Tag                `xml:"tagSet>item"`
	VirtualizationType  string               `xml:"virtualizationType"`
}

// ImageAttribute is undocumented.
type ImageAttribute struct {
	BlockDeviceMappings []BlockDeviceMapping `xml:"blockDeviceMapping>item"`
	Description         AttributeValue       `xml:"description"`
	ImageID             string               `xml:"imageId"`
	KernelID            AttributeValue       `xml:"kernel"`
	LaunchPermissions   []LaunchPermission   `xml:"launchPermission>item"`
	ProductCodes        []ProductCode        `xml:"productCodes>item"`
	RamdiskID           AttributeValue       `xml:"ramdisk"`
	SriovNetSupport     AttributeValue       `xml:"sriovNetSupport"`
}

// ImportInstanceLaunchSpecification is undocumented.
type ImportInstanceLaunchSpecification struct {
	AdditionalInfo                    string    `xml:"additionalInfo"`
	Architecture                      string    `xml:"architecture"`
	GroupIds                          []string  `xml:"GroupId>SecurityGroupId"`
	GroupNames                        []string  `xml:"GroupName>SecurityGroup"`
	InstanceInitiatedShutdownBehavior string    `xml:"instanceInitiatedShutdownBehavior"`
	InstanceType                      string    `xml:"instanceType"`
	Monitoring                        bool      `xml:"monitoring"`
	Placement                         Placement `xml:"placement"`
	PrivateIPAddress                  string    `xml:"privateIpAddress"`
	SubnetID                          string    `xml:"subnetId"`
	UserData                          string    `xml:"userData"`
}

// ImportInstanceRequest is undocumented.
type ImportInstanceRequest struct {
	Description         string                            `xml:"description"`
	DiskImages          []DiskImage                       `xml:"diskImage>member"`
	DryRun              bool                              `xml:"dryRun"`
	LaunchSpecification ImportInstanceLaunchSpecification `xml:"launchSpecification"`
	Platform            string                            `xml:"platform"`
}

// ImportInstanceResult is undocumented.
type ImportInstanceResult struct {
	ConversionTask ConversionTask `xml:"conversionTask"`
}

// ImportInstanceTaskDetails is undocumented.
type ImportInstanceTaskDetails struct {
	Description string                           `xml:"description"`
	InstanceID  string                           `xml:"instanceId"`
	Platform    string                           `xml:"platform"`
	Volumes     []ImportInstanceVolumeDetailItem `xml:"volumes>item"`
}

// ImportInstanceVolumeDetailItem is undocumented.
type ImportInstanceVolumeDetailItem struct {
	AvailabilityZone string                     `xml:"availabilityZone"`
	BytesConverted   int64                      `xml:"bytesConverted"`
	Description      string                     `xml:"description"`
	Image            DiskImageDescription       `xml:"image"`
	Status           string                     `xml:"status"`
	StatusMessage    string                     `xml:"statusMessage"`
	Volume           DiskImageVolumeDescription `xml:"volume"`
}

// ImportKeyPairRequest is undocumented.
type ImportKeyPairRequest struct {
	DryRun            bool   `xml:"dryRun"`
	KeyName           string `xml:"keyName"`
	PublicKeyMaterial []byte `xml:"publicKeyMaterial"`
}

// ImportKeyPairResult is undocumented.
type ImportKeyPairResult struct {
	KeyFingerprint string `xml:"keyFingerprint"`
	KeyName        string `xml:"keyName"`
}

// ImportVolumeRequest is undocumented.
type ImportVolumeRequest struct {
	AvailabilityZone string          `xml:"availabilityZone"`
	Description      string          `xml:"description"`
	DryRun           bool            `xml:"dryRun"`
	Image            DiskImageDetail `xml:"image"`
	Volume           VolumeDetail    `xml:"volume"`
}

// ImportVolumeResult is undocumented.
type ImportVolumeResult struct {
	ConversionTask ConversionTask `xml:"conversionTask"`
}

// ImportVolumeTaskDetails is undocumented.
type ImportVolumeTaskDetails struct {
	AvailabilityZone string                     `xml:"availabilityZone"`
	BytesConverted   int64                      `xml:"bytesConverted"`
	Description      string                     `xml:"description"`
	Image            DiskImageDescription       `xml:"image"`
	Volume           DiskImageVolumeDescription `xml:"volume"`
}

// Instance is undocumented.
type Instance struct {
	AmiLaunchIndex        int                          `xml:"amiLaunchIndex"`
	Architecture          string                       `xml:"architecture"`
	BlockDeviceMappings   []InstanceBlockDeviceMapping `xml:"blockDeviceMapping>item"`
	ClientToken           string                       `xml:"clientToken"`
	EbsOptimized          bool                         `xml:"ebsOptimized"`
	Hypervisor            string                       `xml:"hypervisor"`
	IamInstanceProfile    IamInstanceProfile           `xml:"iamInstanceProfile"`
	ImageID               string                       `xml:"imageId"`
	InstanceID            string                       `xml:"instanceId"`
	InstanceLifecycle     string                       `xml:"instanceLifecycle"`
	InstanceType          string                       `xml:"instanceType"`
	KernelID              string                       `xml:"kernelId"`
	KeyName               string                       `xml:"keyName"`
	LaunchTime            time.Time                    `xml:"launchTime"`
	Monitoring            Monitoring                   `xml:"monitoring"`
	NetworkInterfaces     []InstanceNetworkInterface   `xml:"networkInterfaceSet>item"`
	Placement             Placement                    `xml:"placement"`
	Platform              string                       `xml:"platform"`
	PrivateDNSName        string                       `xml:"privateDnsName"`
	PrivateIPAddress      string                       `xml:"privateIpAddress"`
	ProductCodes          []ProductCode                `xml:"productCodes>item"`
	PublicDNSName         string                       `xml:"dnsName"`
	PublicIPAddress       string                       `xml:"ipAddress"`
	RamdiskID             string                       `xml:"ramdiskId"`
	RootDeviceName        string                       `xml:"rootDeviceName"`
	RootDeviceType        string                       `xml:"rootDeviceType"`
	SecurityGroups        []GroupIdentifier            `xml:"groupSet>item"`
	SourceDestCheck       bool                         `xml:"sourceDestCheck"`
	SpotInstanceRequestID string                       `xml:"spotInstanceRequestId"`
	SriovNetSupport       string                       `xml:"sriovNetSupport"`
	State                 InstanceState                `xml:"instanceState"`
	StateReason           StateReason                  `xml:"stateReason"`
	StateTransitionReason string                       `xml:"reason"`
	SubnetID              string                       `xml:"subnetId"`
	Tags                  []Tag                        `xml:"tagSet>item"`
	VirtualizationType    string                       `xml:"virtualizationType"`
	VpcID                 string                       `xml:"vpcId"`
}

// InstanceAttribute is undocumented.
type InstanceAttribute struct {
	BlockDeviceMappings               []InstanceBlockDeviceMapping `xml:"blockDeviceMapping>item"`
	DisableApiTermination             AttributeBooleanValue        `xml:"disableApiTermination"`
	EbsOptimized                      AttributeBooleanValue        `xml:"ebsOptimized"`
	Groups                            []GroupIdentifier            `xml:"groupSet>item"`
	InstanceID                        string                       `xml:"instanceId"`
	InstanceInitiatedShutdownBehavior AttributeValue               `xml:"instanceInitiatedShutdownBehavior"`
	InstanceType                      AttributeValue               `xml:"instanceType"`
	KernelID                          AttributeValue               `xml:"kernel"`
	ProductCodes                      []ProductCode                `xml:"productCodes>item"`
	RamdiskID                         AttributeValue               `xml:"ramdisk"`
	RootDeviceName                    AttributeValue               `xml:"rootDeviceName"`
	SourceDestCheck                   AttributeBooleanValue        `xml:"sourceDestCheck"`
	SriovNetSupport                   AttributeValue               `xml:"sriovNetSupport"`
	UserData                          AttributeValue               `xml:"userData"`
}

// InstanceBlockDeviceMapping is undocumented.
type InstanceBlockDeviceMapping struct {
	DeviceName string                 `xml:"deviceName"`
	Ebs        EbsInstanceBlockDevice `xml:"ebs"`
}

// InstanceBlockDeviceMappingSpecification is undocumented.
type InstanceBlockDeviceMappingSpecification struct {
	DeviceName  string                              `xml:"deviceName"`
	Ebs         EbsInstanceBlockDeviceSpecification `xml:"ebs"`
	NoDevice    string                              `xml:"noDevice"`
	VirtualName string                              `xml:"virtualName"`
}

// InstanceCount is undocumented.
type InstanceCount struct {
	InstanceCount int    `xml:"instanceCount"`
	State         string `xml:"state"`
}

// InstanceExportDetails is undocumented.
type InstanceExportDetails struct {
	InstanceID        string `xml:"instanceId"`
	TargetEnvironment string `xml:"targetEnvironment"`
}

// InstanceMonitoring is undocumented.
type InstanceMonitoring struct {
	InstanceID string     `xml:"instanceId"`
	Monitoring Monitoring `xml:"monitoring"`
}

// InstanceNetworkInterface is undocumented.
type InstanceNetworkInterface struct {
	Association        InstanceNetworkInterfaceAssociation `xml:"association"`
	Attachment         InstanceNetworkInterfaceAttachment  `xml:"attachment"`
	Description        string                              `xml:"description"`
	Groups             []GroupIdentifier                   `xml:"groupSet>item"`
	MacAddress         string                              `xml:"macAddress"`
	NetworkInterfaceID string                              `xml:"networkInterfaceId"`
	OwnerID            string                              `xml:"ownerId"`
	PrivateDNSName     string                              `xml:"privateDnsName"`
	PrivateIPAddress   string                              `xml:"privateIpAddress"`
	PrivateIPAddresses []InstancePrivateIPAddress          `xml:"privateIpAddressesSet>item"`
	SourceDestCheck    bool                                `xml:"sourceDestCheck"`
	Status             string                              `xml:"status"`
	SubnetID           string                              `xml:"subnetId"`
	VpcID              string                              `xml:"vpcId"`
}

// InstanceNetworkInterfaceAssociation is undocumented.
type InstanceNetworkInterfaceAssociation struct {
	IPOwnerID     string `xml:"ipOwnerId"`
	PublicDNSName string `xml:"publicDnsName"`
	PublicIP      string `xml:"publicIp"`
}

// InstanceNetworkInterfaceAttachment is undocumented.
type InstanceNetworkInterfaceAttachment struct {
	AttachTime          time.Time `xml:"attachTime"`
	AttachmentID        string    `xml:"attachmentId"`
	DeleteOnTermination bool      `xml:"deleteOnTermination"`
	DeviceIndex         int       `xml:"deviceIndex"`
	Status              string    `xml:"status"`
}

// InstanceNetworkInterfaceSpecification is undocumented.
type InstanceNetworkInterfaceSpecification struct {
	AssociatePublicIPAddress       bool                            `xml:"associatePublicIpAddress"`
	DeleteOnTermination            bool                            `xml:"deleteOnTermination"`
	Description                    string                          `xml:"description"`
	DeviceIndex                    int                             `xml:"deviceIndex"`
	Groups                         []string                        `xml:"SecurityGroupId>SecurityGroupId"`
	NetworkInterfaceID             string                          `xml:"networkInterfaceId"`
	PrivateIPAddress               string                          `xml:"privateIpAddress"`
	PrivateIPAddresses             []PrivateIPAddressSpecification `xml:"privateIpAddressesSet>item"`
	SecondaryPrivateIPAddressCount int                             `xml:"secondaryPrivateIpAddressCount"`
	SubnetID                       string                          `xml:"subnetId"`
}

// InstancePrivateIPAddress is undocumented.
type InstancePrivateIPAddress struct {
	Association      InstanceNetworkInterfaceAssociation `xml:"association"`
	Primary          bool                                `xml:"primary"`
	PrivateDNSName   string                              `xml:"privateDnsName"`
	PrivateIPAddress string                              `xml:"privateIpAddress"`
}

// InstanceState is undocumented.
type InstanceState struct {
	Code int    `xml:"code"`
	Name string `xml:"name"`
}

// InstanceStateChange is undocumented.
type InstanceStateChange struct {
	CurrentState  InstanceState `xml:"currentState"`
	InstanceID    string        `xml:"instanceId"`
	PreviousState InstanceState `xml:"previousState"`
}

// InstanceStatus is undocumented.
type InstanceStatus struct {
	AvailabilityZone string                `xml:"availabilityZone"`
	Events           []InstanceStatusEvent `xml:"eventsSet>item"`
	InstanceID       string                `xml:"instanceId"`
	InstanceState    InstanceState         `xml:"instanceState"`
	InstanceStatus   InstanceStatusSummary `xml:"instanceStatus"`
	SystemStatus     InstanceStatusSummary `xml:"systemStatus"`
}

// InstanceStatusDetails is undocumented.
type InstanceStatusDetails struct {
	ImpairedSince time.Time `xml:"impairedSince"`
	Name          string    `xml:"name"`
	Status        string    `xml:"status"`
}

// InstanceStatusEvent is undocumented.
type InstanceStatusEvent struct {
	Code        string    `xml:"code"`
	Description string    `xml:"description"`
	NotAfter    time.Time `xml:"notAfter"`
	NotBefore   time.Time `xml:"notBefore"`
}

// InstanceStatusSummary is undocumented.
type InstanceStatusSummary struct {
	Details []InstanceStatusDetails `xml:"details>item"`
	Status  string                  `xml:"status"`
}

// InternetGateway is undocumented.
type InternetGateway struct {
	Attachments       []InternetGatewayAttachment `xml:"attachmentSet>item"`
	InternetGatewayID string                      `xml:"internetGatewayId"`
	Tags              []Tag                       `xml:"tagSet>item"`
}

// InternetGatewayAttachment is undocumented.
type InternetGatewayAttachment struct {
	State string `xml:"state"`
	VpcID string `xml:"vpcId"`
}

// IPPermission is undocumented.
type IPPermission struct {
	FromPort         int               `xml:"fromPort"`
	IPProtocol       string            `xml:"ipProtocol"`
	IPRanges         []IPRange         `xml:"ipRanges>item"`
	ToPort           int               `xml:"toPort"`
	UserIDGroupPairs []UserIDGroupPair `xml:"groups>item"`
}

// IPRange is undocumented.
type IPRange struct {
	CidrIP string `xml:"cidrIp"`
}

// KeyPair is undocumented.
type KeyPair struct {
	KeyFingerprint string `xml:"keyFingerprint"`
	KeyMaterial    string `xml:"keyMaterial"`
	KeyName        string `xml:"keyName"`
}

// KeyPairInfo is undocumented.
type KeyPairInfo struct {
	KeyFingerprint string `xml:"keyFingerprint"`
	KeyName        string `xml:"keyName"`
}

// LaunchPermission is undocumented.
type LaunchPermission struct {
	Group  string `xml:"group"`
	UserID string `xml:"userId"`
}

// LaunchPermissionModifications is undocumented.
type LaunchPermissionModifications struct {
	Add    []LaunchPermission `xml:"Add>item"`
	Remove []LaunchPermission `xml:"Remove>item"`
}

// LaunchSpecification is undocumented.
type LaunchSpecification struct {
	AddressingType      string                                  `xml:"addressingType"`
	BlockDeviceMappings []BlockDeviceMapping                    `xml:"blockDeviceMapping>item"`
	EbsOptimized        bool                                    `xml:"ebsOptimized"`
	IamInstanceProfile  IamInstanceProfileSpecification         `xml:"iamInstanceProfile"`
	ImageID             string                                  `xml:"imageId"`
	InstanceType        string                                  `xml:"instanceType"`
	KernelID            string                                  `xml:"kernelId"`
	KeyName             string                                  `xml:"keyName"`
	Monitoring          RunInstancesMonitoringEnabled           `xml:"monitoring"`
	NetworkInterfaces   []InstanceNetworkInterfaceSpecification `xml:"networkInterfaceSet>item"`
	Placement           SpotPlacement                           `xml:"placement"`
	RamdiskID           string                                  `xml:"ramdiskId"`
	SecurityGroups      []GroupIdentifier                       `xml:"groupSet>item"`
	SubnetID            string                                  `xml:"subnetId"`
	UserData            string                                  `xml:"userData"`
}

// ModifyImageAttributeRequest is undocumented.
type ModifyImageAttributeRequest struct {
	Attribute        string                        `xml:"Attribute"`
	Description      AttributeValue                `xml:"Description"`
	DryRun           bool                          `xml:"dryRun"`
	ImageID          string                        `xml:"ImageId"`
	LaunchPermission LaunchPermissionModifications `xml:"LaunchPermission"`
	OperationType    string                        `xml:"OperationType"`
	ProductCodes     []string                      `xml:"ProductCode>ProductCode"`
	UserGroups       []string                      `xml:"UserGroup>UserGroup"`
	UserIds          []string                      `xml:"UserId>UserId"`
	Value            string                        `xml:"Value"`
}

// ModifyInstanceAttributeRequest is undocumented.
type ModifyInstanceAttributeRequest struct {
	Attribute                         string                                    `xml:"attribute"`
	BlockDeviceMappings               []InstanceBlockDeviceMappingSpecification `xml:"blockDeviceMapping>item"`
	DisableApiTermination             AttributeBooleanValue                     `xml:"disableApiTermination"`
	DryRun                            bool                                      `xml:"dryRun"`
	EbsOptimized                      AttributeBooleanValue                     `xml:"ebsOptimized"`
	Groups                            []string                                  `xml:"GroupId>groupId"`
	InstanceID                        string                                    `xml:"instanceId"`
	InstanceInitiatedShutdownBehavior AttributeValue                            `xml:"instanceInitiatedShutdownBehavior"`
	InstanceType                      AttributeValue                            `xml:"instanceType"`
	Kernel                            AttributeValue                            `xml:"kernel"`
	Ramdisk                           AttributeValue                            `xml:"ramdisk"`
	SourceDestCheck                   AttributeBooleanValue                     `xml:"SourceDestCheck"`
	SriovNetSupport                   AttributeValue                            `xml:"sriovNetSupport"`
	UserData                          BlobAttributeValue                        `xml:"userData"`
	Value                             string                                    `xml:"value"`
}

// ModifyNetworkInterfaceAttributeRequest is undocumented.
type ModifyNetworkInterfaceAttributeRequest struct {
	Attachment         NetworkInterfaceAttachmentChanges `xml:"attachment"`
	Description        AttributeValue                    `xml:"description"`
	DryRun             bool                              `xml:"dryRun"`
	Groups             []string                          `xml:"SecurityGroupId>SecurityGroupId"`
	NetworkInterfaceID string                            `xml:"networkInterfaceId"`
	SourceDestCheck    AttributeBooleanValue             `xml:"sourceDestCheck"`
}

// ModifyReservedInstancesRequest is undocumented.
type ModifyReservedInstancesRequest struct {
	ClientToken          string                           `xml:"clientToken"`
	ReservedInstancesIds []string                         `xml:"ReservedInstancesId>ReservedInstancesId"`
	TargetConfigurations []ReservedInstancesConfiguration `xml:"ReservedInstancesConfigurationSetItemType>item"`
}

// ModifyReservedInstancesResult is undocumented.
type ModifyReservedInstancesResult struct {
	ReservedInstancesModificationID string `xml:"reservedInstancesModificationId"`
}

// ModifySnapshotAttributeRequest is undocumented.
type ModifySnapshotAttributeRequest struct {
	Attribute              string                              `xml:"Attribute"`
	CreateVolumePermission CreateVolumePermissionModifications `xml:"CreateVolumePermission"`
	DryRun                 bool                                `xml:"dryRun"`
	GroupNames             []string                            `xml:"UserGroup>GroupName"`
	OperationType          string                              `xml:"OperationType"`
	SnapshotID             string                              `xml:"SnapshotId"`
	UserIds                []string                            `xml:"UserId>UserId"`
}

// ModifySubnetAttributeRequest is undocumented.
type ModifySubnetAttributeRequest struct {
	MapPublicIPOnLaunch AttributeBooleanValue `xml:"MapPublicIpOnLaunch"`
	SubnetID            string                `xml:"subnetId"`
}

// ModifyVolumeAttributeRequest is undocumented.
type ModifyVolumeAttributeRequest struct {
	AutoEnableIO AttributeBooleanValue `xml:"AutoEnableIO"`
	DryRun       bool                  `xml:"dryRun"`
	VolumeID     string                `xml:"VolumeId"`
}

// ModifyVpcAttributeRequest is undocumented.
type ModifyVpcAttributeRequest struct {
	EnableDNSHostnames AttributeBooleanValue `xml:"EnableDnsHostnames"`
	EnableDNSSupport   AttributeBooleanValue `xml:"EnableDnsSupport"`
	VpcID              string                `xml:"vpcId"`
}

// MonitorInstancesRequest is undocumented.
type MonitorInstancesRequest struct {
	DryRun      bool     `xml:"dryRun"`
	InstanceIds []string `xml:"InstanceId>InstanceId"`
}

// MonitorInstancesResult is undocumented.
type MonitorInstancesResult struct {
	InstanceMonitorings []InstanceMonitoring `xml:"instancesSet>item"`
}

// Monitoring is undocumented.
type Monitoring struct {
	State string `xml:"state"`
}

// NetworkAcl is undocumented.
type NetworkAcl struct {
	Associations []NetworkAclAssociation `xml:"associationSet>item"`
	Entries      []NetworkAclEntry       `xml:"entrySet>item"`
	IsDefault    bool                    `xml:"default"`
	NetworkAclID string                  `xml:"networkAclId"`
	Tags         []Tag                   `xml:"tagSet>item"`
	VpcID        string                  `xml:"vpcId"`
}

// NetworkAclAssociation is undocumented.
type NetworkAclAssociation struct {
	NetworkAclAssociationID string `xml:"networkAclAssociationId"`
	NetworkAclID            string `xml:"networkAclId"`
	SubnetID                string `xml:"subnetId"`
}

// NetworkAclEntry is undocumented.
type NetworkAclEntry struct {
	CidrBlock    string       `xml:"cidrBlock"`
	Egress       bool         `xml:"egress"`
	IcmpTypeCode IcmpTypeCode `xml:"icmpTypeCode"`
	PortRange    PortRange    `xml:"portRange"`
	Protocol     string       `xml:"protocol"`
	RuleAction   string       `xml:"ruleAction"`
	RuleNumber   int          `xml:"ruleNumber"`
}

// NetworkInterface is undocumented.
type NetworkInterface struct {
	Association        NetworkInterfaceAssociation        `xml:"association"`
	Attachment         NetworkInterfaceAttachment         `xml:"attachment"`
	AvailabilityZone   string                             `xml:"availabilityZone"`
	Description        string                             `xml:"description"`
	Groups             []GroupIdentifier                  `xml:"groupSet>item"`
	MacAddress         string                             `xml:"macAddress"`
	NetworkInterfaceID string                             `xml:"networkInterfaceId"`
	OwnerID            string                             `xml:"ownerId"`
	PrivateDNSName     string                             `xml:"privateDnsName"`
	PrivateIPAddress   string                             `xml:"privateIpAddress"`
	PrivateIPAddresses []NetworkInterfacePrivateIPAddress `xml:"privateIpAddressesSet>item"`
	RequesterID        string                             `xml:"requesterId"`
	RequesterManaged   bool                               `xml:"requesterManaged"`
	SourceDestCheck    bool                               `xml:"sourceDestCheck"`
	Status             string                             `xml:"status"`
	SubnetID           string                             `xml:"subnetId"`
	TagSet             []Tag                              `xml:"tagSet>item"`
	VpcID              string                             `xml:"vpcId"`
}

// NetworkInterfaceAssociation is undocumented.
type NetworkInterfaceAssociation struct {
	AllocationID  string `xml:"allocationId"`
	AssociationID string `xml:"associationId"`
	IPOwnerID     string `xml:"ipOwnerId"`
	PublicDNSName string `xml:"publicDnsName"`
	PublicIP      string `xml:"publicIp"`
}

// NetworkInterfaceAttachment is undocumented.
type NetworkInterfaceAttachment struct {
	AttachTime          time.Time `xml:"attachTime"`
	AttachmentID        string    `xml:"attachmentId"`
	DeleteOnTermination bool      `xml:"deleteOnTermination"`
	DeviceIndex         int       `xml:"deviceIndex"`
	InstanceID          string    `xml:"instanceId"`
	InstanceOwnerID     string    `xml:"instanceOwnerId"`
	Status              string    `xml:"status"`
}

// NetworkInterfaceAttachmentChanges is undocumented.
type NetworkInterfaceAttachmentChanges struct {
	AttachmentID        string `xml:"attachmentId"`
	DeleteOnTermination bool   `xml:"deleteOnTermination"`
}

// NetworkInterfacePrivateIPAddress is undocumented.
type NetworkInterfacePrivateIPAddress struct {
	Association      NetworkInterfaceAssociation `xml:"association"`
	Primary          bool                        `xml:"primary"`
	PrivateDNSName   string                      `xml:"privateDnsName"`
	PrivateIPAddress string                      `xml:"privateIpAddress"`
}

// NewDhcpConfiguration is undocumented.
type NewDhcpConfiguration struct {
	Key    string   `xml:"key"`
	Values []string `xml:"Value>item"`
}

// Placement is undocumented.
type Placement struct {
	AvailabilityZone string `xml:"availabilityZone"`
	GroupName        string `xml:"groupName"`
	Tenancy          string `xml:"tenancy"`
}

// PlacementGroup is undocumented.
type PlacementGroup struct {
	GroupName string `xml:"groupName"`
	State     string `xml:"state"`
	Strategy  string `xml:"strategy"`
}

// PortRange is undocumented.
type PortRange struct {
	From int `xml:"from"`
	To   int `xml:"to"`
}

// PriceSchedule is undocumented.
type PriceSchedule struct {
	Active       bool    `xml:"active"`
	CurrencyCode string  `xml:"currencyCode"`
	Price        float64 `xml:"price"`
	Term         int64   `xml:"term"`
}

// PriceScheduleSpecification is undocumented.
type PriceScheduleSpecification struct {
	CurrencyCode string  `xml:"currencyCode"`
	Price        float64 `xml:"price"`
	Term         int64   `xml:"term"`
}

// PricingDetail is undocumented.
type PricingDetail struct {
	Count int     `xml:"count"`
	Price float64 `xml:"price"`
}

// PrivateIPAddressSpecification is undocumented.
type PrivateIPAddressSpecification struct {
	Primary          bool   `xml:"primary"`
	PrivateIPAddress string `xml:"privateIpAddress"`
}

// ProductCode is undocumented.
type ProductCode struct {
	ProductCodeID   string `xml:"productCode"`
	ProductCodeType string `xml:"type"`
}

// PropagatingVgw is undocumented.
type PropagatingVgw struct {
	GatewayID string `xml:"gatewayId"`
}

// PurchaseReservedInstancesOfferingRequest is undocumented.
type PurchaseReservedInstancesOfferingRequest struct {
	DryRun                      bool                       `xml:"dryRun"`
	InstanceCount               int                        `xml:"InstanceCount"`
	LimitPrice                  ReservedInstanceLimitPrice `xml:"limitPrice"`
	ReservedInstancesOfferingID string                     `xml:"ReservedInstancesOfferingId"`
}

// PurchaseReservedInstancesOfferingResult is undocumented.
type PurchaseReservedInstancesOfferingResult struct {
	ReservedInstancesID string `xml:"reservedInstancesId"`
}

// RebootInstancesRequest is undocumented.
type RebootInstancesRequest struct {
	DryRun      bool     `xml:"dryRun"`
	InstanceIds []string `xml:"InstanceId>InstanceId"`
}

// RecurringCharge is undocumented.
type RecurringCharge struct {
	Amount    float64 `xml:"amount"`
	Frequency string  `xml:"frequency"`
}

// Region is undocumented.
type Region struct {
	Endpoint   string `xml:"regionEndpoint"`
	RegionName string `xml:"regionName"`
}

// RegisterImageRequest is undocumented.
type RegisterImageRequest struct {
	Architecture        string               `xml:"architecture"`
	BlockDeviceMappings []BlockDeviceMapping `xml:"BlockDeviceMapping>BlockDeviceMapping"`
	Description         string               `xml:"description"`
	DryRun              bool                 `xml:"dryRun"`
	ImageLocation       string               `xml:"ImageLocation"`
	KernelID            string               `xml:"kernelId"`
	Name                string               `xml:"name"`
	RamdiskID           string               `xml:"ramdiskId"`
	RootDeviceName      string               `xml:"rootDeviceName"`
	SriovNetSupport     string               `xml:"sriovNetSupport"`
	VirtualizationType  string               `xml:"virtualizationType"`
}

// RegisterImageResult is undocumented.
type RegisterImageResult struct {
	ImageID string `xml:"imageId"`
}

// RejectVpcPeeringConnectionRequest is undocumented.
type RejectVpcPeeringConnectionRequest struct {
	DryRun                 bool   `xml:"dryRun"`
	VpcPeeringConnectionID string `xml:"vpcPeeringConnectionId"`
}

// RejectVpcPeeringConnectionResult is undocumented.
type RejectVpcPeeringConnectionResult struct {
	Return bool `xml:"return"`
}

// ReleaseAddressRequest is undocumented.
type ReleaseAddressRequest struct {
	AllocationID string `xml:"AllocationId"`
	DryRun       bool   `xml:"dryRun"`
	PublicIP     string `xml:"PublicIp"`
}

// ReplaceNetworkAclAssociationRequest is undocumented.
type ReplaceNetworkAclAssociationRequest struct {
	AssociationID string `xml:"associationId"`
	DryRun        bool   `xml:"dryRun"`
	NetworkAclID  string `xml:"networkAclId"`
}

// ReplaceNetworkAclAssociationResult is undocumented.
type ReplaceNetworkAclAssociationResult struct {
	NewAssociationID string `xml:"newAssociationId"`
}

// ReplaceNetworkAclEntryRequest is undocumented.
type ReplaceNetworkAclEntryRequest struct {
	CidrBlock    string       `xml:"cidrBlock"`
	DryRun       bool         `xml:"dryRun"`
	Egress       bool         `xml:"egress"`
	IcmpTypeCode IcmpTypeCode `xml:"Icmp"`
	NetworkAclID string       `xml:"networkAclId"`
	PortRange    PortRange    `xml:"portRange"`
	Protocol     string       `xml:"protocol"`
	RuleAction   string       `xml:"ruleAction"`
	RuleNumber   int          `xml:"ruleNumber"`
}

// ReplaceRouteRequest is undocumented.
type ReplaceRouteRequest struct {
	DestinationCidrBlock   string `xml:"destinationCidrBlock"`
	DryRun                 bool   `xml:"dryRun"`
	GatewayID              string `xml:"gatewayId"`
	InstanceID             string `xml:"instanceId"`
	NetworkInterfaceID     string `xml:"networkInterfaceId"`
	RouteTableID           string `xml:"routeTableId"`
	VpcPeeringConnectionID string `xml:"vpcPeeringConnectionId"`
}

// ReplaceRouteTableAssociationRequest is undocumented.
type ReplaceRouteTableAssociationRequest struct {
	AssociationID string `xml:"associationId"`
	DryRun        bool   `xml:"dryRun"`
	RouteTableID  string `xml:"routeTableId"`
}

// ReplaceRouteTableAssociationResult is undocumented.
type ReplaceRouteTableAssociationResult struct {
	NewAssociationID string `xml:"newAssociationId"`
}

// ReportInstanceStatusRequest is undocumented.
type ReportInstanceStatusRequest struct {
	Description string    `xml:"description"`
	DryRun      bool      `xml:"dryRun"`
	EndTime     time.Time `xml:"endTime"`
	Instances   []string  `xml:"instanceId>InstanceId"`
	ReasonCodes []string  `xml:"reasonCode>item"`
	StartTime   time.Time `xml:"startTime"`
	Status      string    `xml:"status"`
}

// RequestSpotInstancesRequest is undocumented.
type RequestSpotInstancesRequest struct {
	AvailabilityZoneGroup string                         `xml:"availabilityZoneGroup"`
	DryRun                bool                           `xml:"dryRun"`
	InstanceCount         int                            `xml:"instanceCount"`
	LaunchGroup           string                         `xml:"launchGroup"`
	LaunchSpecification   RequestSpotLaunchSpecification `xml:"LaunchSpecification"`
	SpotPrice             string                         `xml:"spotPrice"`
	Type                  string                         `xml:"type"`
	ValidFrom             time.Time                      `xml:"validFrom"`
	ValidUntil            time.Time                      `xml:"validUntil"`
}

// RequestSpotInstancesResult is undocumented.
type RequestSpotInstancesResult struct {
	SpotInstanceRequests []SpotInstanceRequest `xml:"spotInstanceRequestSet>item"`
}

// RequestSpotLaunchSpecification is undocumented.
type RequestSpotLaunchSpecification struct {
	AddressingType      string                                  `xml:"addressingType"`
	BlockDeviceMappings []BlockDeviceMapping                    `xml:"blockDeviceMapping>item"`
	EbsOptimized        bool                                    `xml:"ebsOptimized"`
	IamInstanceProfile  IamInstanceProfileSpecification         `xml:"iamInstanceProfile"`
	ImageID             string                                  `xml:"imageId"`
	InstanceType        string                                  `xml:"instanceType"`
	KernelID            string                                  `xml:"kernelId"`
	KeyName             string                                  `xml:"keyName"`
	Monitoring          RunInstancesMonitoringEnabled           `xml:"monitoring"`
	NetworkInterfaces   []InstanceNetworkInterfaceSpecification `xml:"NetworkInterface>item"`
	Placement           SpotPlacement                           `xml:"placement"`
	RamdiskID           string                                  `xml:"ramdiskId"`
	SecurityGroupIds    []string                                `xml:"SecurityGroupId>item"`
	SecurityGroups      []string                                `xml:"SecurityGroup>item"`
	SubnetID            string                                  `xml:"subnetId"`
	UserData            string                                  `xml:"userData"`
}

// Reservation is undocumented.
type Reservation struct {
	Groups        []GroupIdentifier `xml:"groupSet>item"`
	Instances     []Instance        `xml:"instancesSet>item"`
	OwnerID       string            `xml:"ownerId"`
	RequesterID   string            `xml:"requesterId"`
	ReservationID string            `xml:"reservationId"`
}

// ReservedInstanceLimitPrice is undocumented.
type ReservedInstanceLimitPrice struct {
	Amount       float64 `xml:"amount"`
	CurrencyCode string  `xml:"currencyCode"`
}

// ReservedInstances is undocumented.
type ReservedInstances struct {
	AvailabilityZone    string            `xml:"availabilityZone"`
	CurrencyCode        string            `xml:"currencyCode"`
	Duration            int64             `xml:"duration"`
	End                 time.Time         `xml:"end"`
	FixedPrice          float32           `xml:"fixedPrice"`
	InstanceCount       int               `xml:"instanceCount"`
	InstanceTenancy     string            `xml:"instanceTenancy"`
	InstanceType        string            `xml:"instanceType"`
	OfferingType        string            `xml:"offeringType"`
	ProductDescription  string            `xml:"productDescription"`
	RecurringCharges    []RecurringCharge `xml:"recurringCharges>item"`
	ReservedInstancesID string            `xml:"reservedInstancesId"`
	Start               time.Time         `xml:"start"`
	State               string            `xml:"state"`
	Tags                []Tag             `xml:"tagSet>item"`
	UsagePrice          float32           `xml:"usagePrice"`
}

// ReservedInstancesConfiguration is undocumented.
type ReservedInstancesConfiguration struct {
	AvailabilityZone string `xml:"availabilityZone"`
	InstanceCount    int    `xml:"instanceCount"`
	InstanceType     string `xml:"instanceType"`
	Platform         string `xml:"platform"`
}

// ReservedInstancesID is undocumented.
type ReservedInstancesID struct {
	ReservedInstancesID string `xml:"reservedInstancesId"`
}

// ReservedInstancesListing is undocumented.
type ReservedInstancesListing struct {
	ClientToken                string          `xml:"clientToken"`
	CreateDate                 time.Time       `xml:"createDate"`
	InstanceCounts             []InstanceCount `xml:"instanceCounts>item"`
	PriceSchedules             []PriceSchedule `xml:"priceSchedules>item"`
	ReservedInstancesID        string          `xml:"reservedInstancesId"`
	ReservedInstancesListingID string          `xml:"reservedInstancesListingId"`
	Status                     string          `xml:"status"`
	StatusMessage              string          `xml:"statusMessage"`
	Tags                       []Tag           `xml:"tagSet>item"`
	UpdateDate                 time.Time       `xml:"updateDate"`
}

// ReservedInstancesModification is undocumented.
type ReservedInstancesModification struct {
	ClientToken                     string                                `xml:"clientToken"`
	CreateDate                      time.Time                             `xml:"createDate"`
	EffectiveDate                   time.Time                             `xml:"effectiveDate"`
	ModificationResults             []ReservedInstancesModificationResult `xml:"modificationResultSet>item"`
	ReservedInstancesIds            []ReservedInstancesID                 `xml:"reservedInstancesSet>item"`
	ReservedInstancesModificationID string                                `xml:"reservedInstancesModificationId"`
	Status                          string                                `xml:"status"`
	StatusMessage                   string                                `xml:"statusMessage"`
	UpdateDate                      time.Time                             `xml:"updateDate"`
}

// ReservedInstancesModificationResult is undocumented.
type ReservedInstancesModificationResult struct {
	ReservedInstancesID string                         `xml:"reservedInstancesId"`
	TargetConfiguration ReservedInstancesConfiguration `xml:"targetConfiguration"`
}

// ReservedInstancesOffering is undocumented.
type ReservedInstancesOffering struct {
	AvailabilityZone            string            `xml:"availabilityZone"`
	CurrencyCode                string            `xml:"currencyCode"`
	Duration                    int64             `xml:"duration"`
	FixedPrice                  float32           `xml:"fixedPrice"`
	InstanceTenancy             string            `xml:"instanceTenancy"`
	InstanceType                string            `xml:"instanceType"`
	Marketplace                 bool              `xml:"marketplace"`
	OfferingType                string            `xml:"offeringType"`
	PricingDetails              []PricingDetail   `xml:"pricingDetailsSet>item"`
	ProductDescription          string            `xml:"productDescription"`
	RecurringCharges            []RecurringCharge `xml:"recurringCharges>item"`
	ReservedInstancesOfferingID string            `xml:"reservedInstancesOfferingId"`
	UsagePrice                  float32           `xml:"usagePrice"`
}

// ResetImageAttributeRequest is undocumented.
type ResetImageAttributeRequest struct {
	Attribute string `xml:"Attribute"`
	DryRun    bool   `xml:"dryRun"`
	ImageID   string `xml:"ImageId"`
}

// ResetInstanceAttributeRequest is undocumented.
type ResetInstanceAttributeRequest struct {
	Attribute  string `xml:"attribute"`
	DryRun     bool   `xml:"dryRun"`
	InstanceID string `xml:"instanceId"`
}

// ResetNetworkInterfaceAttributeRequest is undocumented.
type ResetNetworkInterfaceAttributeRequest struct {
	DryRun             bool   `xml:"dryRun"`
	NetworkInterfaceID string `xml:"networkInterfaceId"`
	SourceDestCheck    string `xml:"sourceDestCheck"`
}

// ResetSnapshotAttributeRequest is undocumented.
type ResetSnapshotAttributeRequest struct {
	Attribute  string `xml:"Attribute"`
	DryRun     bool   `xml:"dryRun"`
	SnapshotID string `xml:"SnapshotId"`
}

// RevokeSecurityGroupEgressRequest is undocumented.
type RevokeSecurityGroupEgressRequest struct {
	CidrIP                     string         `xml:"cidrIp"`
	DryRun                     bool           `xml:"dryRun"`
	FromPort                   int            `xml:"fromPort"`
	GroupID                    string         `xml:"groupId"`
	IPPermissions              []IPPermission `xml:"ipPermissions>item"`
	IPProtocol                 string         `xml:"ipProtocol"`
	SourceSecurityGroupName    string         `xml:"sourceSecurityGroupName"`
	SourceSecurityGroupOwnerID string         `xml:"sourceSecurityGroupOwnerId"`
	ToPort                     int            `xml:"toPort"`
}

// RevokeSecurityGroupIngressRequest is undocumented.
type RevokeSecurityGroupIngressRequest struct {
	CidrIP                     string         `xml:"CidrIp"`
	DryRun                     bool           `xml:"dryRun"`
	FromPort                   int            `xml:"FromPort"`
	GroupID                    string         `xml:"GroupId"`
	GroupName                  string         `xml:"GroupName"`
	IPPermissions              []IPPermission `xml:"IpPermissions>item"`
	IPProtocol                 string         `xml:"IpProtocol"`
	SourceSecurityGroupName    string         `xml:"SourceSecurityGroupName"`
	SourceSecurityGroupOwnerID string         `xml:"SourceSecurityGroupOwnerId"`
	ToPort                     int            `xml:"ToPort"`
}

// Route is undocumented.
type Route struct {
	DestinationCidrBlock   string `xml:"destinationCidrBlock"`
	GatewayID              string `xml:"gatewayId"`
	InstanceID             string `xml:"instanceId"`
	InstanceOwnerID        string `xml:"instanceOwnerId"`
	NetworkInterfaceID     string `xml:"networkInterfaceId"`
	Origin                 string `xml:"origin"`
	State                  string `xml:"state"`
	VpcPeeringConnectionID string `xml:"vpcPeeringConnectionId"`
}

// RouteTable is undocumented.
type RouteTable struct {
	Associations    []RouteTableAssociation `xml:"associationSet>item"`
	PropagatingVgws []PropagatingVgw        `xml:"propagatingVgwSet>item"`
	RouteTableID    string                  `xml:"routeTableId"`
	Routes          []Route                 `xml:"routeSet>item"`
	Tags            []Tag                   `xml:"tagSet>item"`
	VpcID           string                  `xml:"vpcId"`
}

// RouteTableAssociation is undocumented.
type RouteTableAssociation struct {
	Main                    bool   `xml:"main"`
	RouteTableAssociationID string `xml:"routeTableAssociationId"`
	RouteTableID            string `xml:"routeTableId"`
	SubnetID                string `xml:"subnetId"`
}

// RunInstancesMonitoringEnabled is undocumented.
type RunInstancesMonitoringEnabled struct {
	Enabled bool `xml:"enabled"`
}

// RunInstancesRequest is undocumented.
type RunInstancesRequest struct {
	AdditionalInfo                    string                                  `xml:"additionalInfo"`
	BlockDeviceMappings               []BlockDeviceMapping                    `xml:"BlockDeviceMapping>BlockDeviceMapping"`
	ClientToken                       string                                  `xml:"clientToken"`
	DisableApiTermination             bool                                    `xml:"disableApiTermination"`
	DryRun                            bool                                    `xml:"dryRun"`
	EbsOptimized                      bool                                    `xml:"ebsOptimized"`
	IamInstanceProfile                IamInstanceProfileSpecification         `xml:"iamInstanceProfile"`
	ImageID                           string                                  `xml:"ImageId"`
	InstanceInitiatedShutdownBehavior string                                  `xml:"instanceInitiatedShutdownBehavior"`
	InstanceType                      string                                  `xml:"InstanceType"`
	KernelID                          string                                  `xml:"KernelId"`
	KeyName                           string                                  `xml:"KeyName"`
	MaxCount                          int                                     `xml:"MaxCount"`
	MinCount                          int                                     `xml:"MinCount"`
	Monitoring                        RunInstancesMonitoringEnabled           `xml:"Monitoring"`
	NetworkInterfaces                 []InstanceNetworkInterfaceSpecification `xml:"networkInterface>item"`
	Placement                         Placement                               `xml:"Placement"`
	PrivateIPAddress                  string                                  `xml:"privateIpAddress"`
	RamdiskID                         string                                  `xml:"RamdiskId"`
	SecurityGroupIds                  []string                                `xml:"SecurityGroupId>SecurityGroupId"`
	SecurityGroups                    []string                                `xml:"SecurityGroup>SecurityGroup"`
	SubnetID                          string                                  `xml:"SubnetId"`
	UserData                          string                                  `xml:"UserData"`
}

// S3Storage is undocumented.
type S3Storage struct {
	AWSAccessKeyID        string `xml:"AWSAccessKeyId"`
	Bucket                string `xml:"bucket"`
	Prefix                string `xml:"prefix"`
	UploadPolicy          []byte `xml:"uploadPolicy"`
	UploadPolicySignature string `xml:"uploadPolicySignature"`
}

// SecurityGroup is undocumented.
type SecurityGroup struct {
	Description         string         `xml:"groupDescription"`
	GroupID             string         `xml:"groupId"`
	GroupName           string         `xml:"groupName"`
	IPPermissions       []IPPermission `xml:"ipPermissions>item"`
	IPPermissionsEgress []IPPermission `xml:"ipPermissionsEgress>item"`
	OwnerID             string         `xml:"ownerId"`
	Tags                []Tag          `xml:"tagSet>item"`
	VpcID               string         `xml:"vpcId"`
}

// Snapshot is undocumented.
type Snapshot struct {
	Description string    `xml:"description"`
	Encrypted   bool      `xml:"encrypted"`
	KmsKeyID    string    `xml:"kmsKeyId"`
	OwnerAlias  string    `xml:"ownerAlias"`
	OwnerID     string    `xml:"ownerId"`
	Progress    string    `xml:"progress"`
	SnapshotID  string    `xml:"snapshotId"`
	StartTime   time.Time `xml:"startTime"`
	State       string    `xml:"status"`
	Tags        []Tag     `xml:"tagSet>item"`
	VolumeID    string    `xml:"volumeId"`
	VolumeSize  int       `xml:"volumeSize"`
}

// SpotDatafeedSubscription is undocumented.
type SpotDatafeedSubscription struct {
	Bucket  string                 `xml:"bucket"`
	Fault   SpotInstanceStateFault `xml:"fault"`
	OwnerID string                 `xml:"ownerId"`
	Prefix  string                 `xml:"prefix"`
	State   string                 `xml:"state"`
}

// SpotInstanceRequest is undocumented.
type SpotInstanceRequest struct {
	AvailabilityZoneGroup    string                 `xml:"availabilityZoneGroup"`
	CreateTime               time.Time              `xml:"createTime"`
	Fault                    SpotInstanceStateFault `xml:"fault"`
	InstanceID               string                 `xml:"instanceId"`
	LaunchGroup              string                 `xml:"launchGroup"`
	LaunchSpecification      LaunchSpecification    `xml:"launchSpecification"`
	LaunchedAvailabilityZone string                 `xml:"launchedAvailabilityZone"`
	ProductDescription       string                 `xml:"productDescription"`
	SpotInstanceRequestID    string                 `xml:"spotInstanceRequestId"`
	SpotPrice                string                 `xml:"spotPrice"`
	State                    string                 `xml:"state"`
	Status                   SpotInstanceStatus     `xml:"status"`
	Tags                     []Tag                  `xml:"tagSet>item"`
	Type                     string                 `xml:"type"`
	ValidFrom                time.Time              `xml:"validFrom"`
	ValidUntil               time.Time              `xml:"validUntil"`
}

// SpotInstanceStateFault is undocumented.
type SpotInstanceStateFault struct {
	Code    string `xml:"code"`
	Message string `xml:"message"`
}

// SpotInstanceStatus is undocumented.
type SpotInstanceStatus struct {
	Code       string    `xml:"code"`
	Message    string    `xml:"message"`
	UpdateTime time.Time `xml:"updateTime"`
}

// SpotPlacement is undocumented.
type SpotPlacement struct {
	AvailabilityZone string `xml:"availabilityZone"`
	GroupName        string `xml:"groupName"`
}

// SpotPrice is undocumented.
type SpotPrice struct {
	AvailabilityZone   string    `xml:"availabilityZone"`
	InstanceType       string    `xml:"instanceType"`
	ProductDescription string    `xml:"productDescription"`
	SpotPrice          string    `xml:"spotPrice"`
	Timestamp          time.Time `xml:"timestamp"`
}

// StartInstancesRequest is undocumented.
type StartInstancesRequest struct {
	AdditionalInfo string   `xml:"additionalInfo"`
	DryRun         bool     `xml:"dryRun"`
	InstanceIds    []string `xml:"InstanceId>InstanceId"`
}

// StartInstancesResult is undocumented.
type StartInstancesResult struct {
	StartingInstances []InstanceStateChange `xml:"instancesSet>item"`
}

// StateReason is undocumented.
type StateReason struct {
	Code    string `xml:"code"`
	Message string `xml:"message"`
}

// StopInstancesRequest is undocumented.
type StopInstancesRequest struct {
	DryRun      bool     `xml:"dryRun"`
	Force       bool     `xml:"force"`
	InstanceIds []string `xml:"InstanceId>InstanceId"`
}

// StopInstancesResult is undocumented.
type StopInstancesResult struct {
	StoppingInstances []InstanceStateChange `xml:"instancesSet>item"`
}

// Storage is undocumented.
type Storage struct {
	S3 S3Storage `xml:"S3"`
}

// Subnet is undocumented.
type Subnet struct {
	AvailabilityZone        string `xml:"availabilityZone"`
	AvailableIPAddressCount int    `xml:"availableIpAddressCount"`
	CidrBlock               string `xml:"cidrBlock"`
	DefaultForAz            bool   `xml:"defaultForAz"`
	MapPublicIPOnLaunch     bool   `xml:"mapPublicIpOnLaunch"`
	State                   string `xml:"state"`
	SubnetID                string `xml:"subnetId"`
	Tags                    []Tag  `xml:"tagSet>item"`
	VpcID                   string `xml:"vpcId"`
}

// Tag is undocumented.
type Tag struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}

// TagDescription is undocumented.
type TagDescription struct {
	Key          string `xml:"key"`
	ResourceID   string `xml:"resourceId"`
	ResourceType string `xml:"resourceType"`
	Value        string `xml:"value"`
}

// TerminateInstancesRequest is undocumented.
type TerminateInstancesRequest struct {
	DryRun      bool     `xml:"dryRun"`
	InstanceIds []string `xml:"InstanceId>InstanceId"`
}

// TerminateInstancesResult is undocumented.
type TerminateInstancesResult struct {
	TerminatingInstances []InstanceStateChange `xml:"instancesSet>item"`
}

// UnassignPrivateIPAddressesRequest is undocumented.
type UnassignPrivateIPAddressesRequest struct {
	NetworkInterfaceID string   `xml:"networkInterfaceId"`
	PrivateIPAddresses []string `xml:"privateIpAddress>PrivateIpAddress"`
}

// UnmonitorInstancesRequest is undocumented.
type UnmonitorInstancesRequest struct {
	DryRun      bool     `xml:"dryRun"`
	InstanceIds []string `xml:"InstanceId>InstanceId"`
}

// UnmonitorInstancesResult is undocumented.
type UnmonitorInstancesResult struct {
	InstanceMonitorings []InstanceMonitoring `xml:"instancesSet>item"`
}

// UserIDGroupPair is undocumented.
type UserIDGroupPair struct {
	GroupID   string `xml:"groupId"`
	GroupName string `xml:"groupName"`
	UserID    string `xml:"userId"`
}

// VgwTelemetry is undocumented.
type VgwTelemetry struct {
	AcceptedRouteCount int       `xml:"acceptedRouteCount"`
	LastStatusChange   time.Time `xml:"lastStatusChange"`
	OutsideIPAddress   string    `xml:"outsideIpAddress"`
	Status             string    `xml:"status"`
	StatusMessage      string    `xml:"statusMessage"`
}

// Volume is undocumented.
type Volume struct {
	Attachments      []VolumeAttachment `xml:"attachmentSet>item"`
	AvailabilityZone string             `xml:"availabilityZone"`
	CreateTime       time.Time          `xml:"createTime"`
	Encrypted        bool               `xml:"encrypted"`
	Iops             int                `xml:"iops"`
	KmsKeyID         string             `xml:"kmsKeyId"`
	Size             int                `xml:"size"`
	SnapshotID       string             `xml:"snapshotId"`
	State            string             `xml:"status"`
	Tags             []Tag              `xml:"tagSet>item"`
	VolumeID         string             `xml:"volumeId"`
	VolumeType       string             `xml:"volumeType"`
}

// VolumeAttachment is undocumented.
type VolumeAttachment struct {
	AttachTime          time.Time `xml:"attachTime"`
	DeleteOnTermination bool      `xml:"deleteOnTermination"`
	Device              string    `xml:"device"`
	InstanceID          string    `xml:"instanceId"`
	State               string    `xml:"status"`
	VolumeID            string    `xml:"volumeId"`
}

// VolumeDetail is undocumented.
type VolumeDetail struct {
	Size int64 `xml:"size"`
}

// VolumeStatusAction is undocumented.
type VolumeStatusAction struct {
	Code        string `xml:"code"`
	Description string `xml:"description"`
	EventID     string `xml:"eventId"`
	EventType   string `xml:"eventType"`
}

// VolumeStatusDetails is undocumented.
type VolumeStatusDetails struct {
	Name   string `xml:"name"`
	Status string `xml:"status"`
}

// VolumeStatusEvent is undocumented.
type VolumeStatusEvent struct {
	Description string    `xml:"description"`
	EventID     string    `xml:"eventId"`
	EventType   string    `xml:"eventType"`
	NotAfter    time.Time `xml:"notAfter"`
	NotBefore   time.Time `xml:"notBefore"`
}

// VolumeStatusInfo is undocumented.
type VolumeStatusInfo struct {
	Details []VolumeStatusDetails `xml:"details>item"`
	Status  string                `xml:"status"`
}

// VolumeStatusItem is undocumented.
type VolumeStatusItem struct {
	Actions          []VolumeStatusAction `xml:"actionsSet>item"`
	AvailabilityZone string               `xml:"availabilityZone"`
	Events           []VolumeStatusEvent  `xml:"eventsSet>item"`
	VolumeID         string               `xml:"volumeId"`
	VolumeStatus     VolumeStatusInfo     `xml:"volumeStatus"`
}

// Vpc is undocumented.
type Vpc struct {
	CidrBlock       string `xml:"cidrBlock"`
	DhcpOptionsID   string `xml:"dhcpOptionsId"`
	InstanceTenancy string `xml:"instanceTenancy"`
	IsDefault       bool   `xml:"isDefault"`
	State           string `xml:"state"`
	Tags            []Tag  `xml:"tagSet>item"`
	VpcID           string `xml:"vpcId"`
}

// VpcAttachment is undocumented.
type VpcAttachment struct {
	State string `xml:"state"`
	VpcID string `xml:"vpcId"`
}

// VpcPeeringConnection is undocumented.
type VpcPeeringConnection struct {
	AccepterVpcInfo        VpcPeeringConnectionVpcInfo     `xml:"accepterVpcInfo"`
	ExpirationTime         time.Time                       `xml:"expirationTime"`
	RequesterVpcInfo       VpcPeeringConnectionVpcInfo     `xml:"requesterVpcInfo"`
	Status                 VpcPeeringConnectionStateReason `xml:"status"`
	Tags                   []Tag                           `xml:"tagSet>item"`
	VpcPeeringConnectionID string                          `xml:"vpcPeeringConnectionId"`
}

// VpcPeeringConnectionStateReason is undocumented.
type VpcPeeringConnectionStateReason struct {
	Code    string `xml:"code"`
	Message string `xml:"message"`
}

// VpcPeeringConnectionVpcInfo is undocumented.
type VpcPeeringConnectionVpcInfo struct {
	CidrBlock string `xml:"cidrBlock"`
	OwnerID   string `xml:"ownerId"`
	VpcID     string `xml:"vpcId"`
}

// VpnConnection is undocumented.
type VpnConnection struct {
	CustomerGatewayConfiguration string               `xml:"customerGatewayConfiguration"`
	CustomerGatewayID            string               `xml:"customerGatewayId"`
	Options                      VpnConnectionOptions `xml:"options"`
	Routes                       []VpnStaticRoute     `xml:"routes>item"`
	State                        string               `xml:"state"`
	Tags                         []Tag                `xml:"tagSet>item"`
	Type                         string               `xml:"type"`
	VgwTelemetry                 []VgwTelemetry       `xml:"vgwTelemetry>item"`
	VpnConnectionID              string               `xml:"vpnConnectionId"`
	VpnGatewayID                 string               `xml:"vpnGatewayId"`
}

// VpnConnectionOptions is undocumented.
type VpnConnectionOptions struct {
	StaticRoutesOnly bool `xml:"staticRoutesOnly"`
}

// VpnConnectionOptionsSpecification is undocumented.
type VpnConnectionOptionsSpecification struct {
	StaticRoutesOnly bool `xml:"staticRoutesOnly"`
}

// VpnGateway is undocumented.
type VpnGateway struct {
	AvailabilityZone string          `xml:"availabilityZone"`
	State            string          `xml:"state"`
	Tags             []Tag           `xml:"tagSet>item"`
	Type             string          `xml:"type"`
	VpcAttachments   []VpcAttachment `xml:"attachments>item"`
	VpnGatewayID     string          `xml:"vpnGatewayId"`
}

// VpnStaticRoute is undocumented.
type VpnStaticRoute struct {
	DestinationCidrBlock string `xml:"destinationCidrBlock"`
	Source               string `xml:"source"`
	State                string `xml:"state"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
