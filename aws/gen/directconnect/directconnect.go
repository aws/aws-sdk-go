// Package directconnect provides a client for AWS Direct Connect.
package directconnect

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// DirectConnect is a client for AWS Direct Connect.
type DirectConnect struct {
	client *aws.JSONClient
}

// New returns a new DirectConnect client.
func New(creds aws.Credentials, region string, client *http.Client) *DirectConnect {
	if client == nil {
		client = http.DefaultClient
	}

	service := "directconnect"
	endpoint, service, region := endpoints.Lookup("directconnect", region)

	return &DirectConnect{
		client: &aws.JSONClient{
			Credentials:  creds,
			Service:      service,
			Region:       region,
			Client:       client,
			Endpoint:     endpoint,
			JSONVersion:  "1.1",
			TargetPrefix: "OvertureService",
		},
	}
}

// AllocateConnectionOnInterconnect creates a hosted connection on an
// interconnect. Allocates a number and a specified amount of bandwidth for
// use by a hosted connection on the given interconnect.
func (c *DirectConnect) AllocateConnectionOnInterconnect(req AllocateConnectionOnInterconnectRequest) (resp *Connection, err error) {
	resp = &Connection{}
	err = c.client.Do("AllocateConnectionOnInterconnect", "POST", "/", req, resp)
	return
}

// AllocatePrivateVirtualInterface provisions a private virtual interface
// to be owned by a different customer. The owner of a connection calls
// this function to provision a private virtual interface which will be
// owned by another AWS customer. Virtual interfaces created using this
// function must be confirmed by the virtual interface owner by calling
// ConfirmPrivateVirtualInterface. Until this step has been completed, the
// virtual interface will be in 'Confirming' state, and will not be
// available for handling traffic.
func (c *DirectConnect) AllocatePrivateVirtualInterface(req AllocatePrivateVirtualInterfaceRequest) (resp *VirtualInterface, err error) {
	resp = &VirtualInterface{}
	err = c.client.Do("AllocatePrivateVirtualInterface", "POST", "/", req, resp)
	return
}

// AllocatePublicVirtualInterface provisions a public virtual interface to
// be owned by a different customer. The owner of a connection calls this
// function to provision a public virtual interface which will be owned by
// another AWS customer. Virtual interfaces created using this function
// must be confirmed by the virtual interface owner by calling
// ConfirmPublicVirtualInterface. Until this step has been completed, the
// virtual interface will be in 'Confirming' state, and will not be
// available for handling traffic.
func (c *DirectConnect) AllocatePublicVirtualInterface(req AllocatePublicVirtualInterfaceRequest) (resp *VirtualInterface, err error) {
	resp = &VirtualInterface{}
	err = c.client.Do("AllocatePublicVirtualInterface", "POST", "/", req, resp)
	return
}

// ConfirmConnection confirm the creation of a hosted connection on an
// interconnect. Upon creation, the hosted connection is initially in the
// 'Ordering' state, and will remain in this state until the owner calls
// ConfirmConnection to confirm creation of the hosted connection.
func (c *DirectConnect) ConfirmConnection(req ConfirmConnectionRequest) (resp *ConfirmConnectionResponse, err error) {
	resp = &ConfirmConnectionResponse{}
	err = c.client.Do("ConfirmConnection", "POST", "/", req, resp)
	return
}

// ConfirmPrivateVirtualInterface accept ownership of a private virtual
// interface created by another customer. After the virtual interface owner
// calls this function, the virtual interface will be created and attached
// to the given virtual private gateway, and will be available for handling
// traffic.
func (c *DirectConnect) ConfirmPrivateVirtualInterface(req ConfirmPrivateVirtualInterfaceRequest) (resp *ConfirmPrivateVirtualInterfaceResponse, err error) {
	resp = &ConfirmPrivateVirtualInterfaceResponse{}
	err = c.client.Do("ConfirmPrivateVirtualInterface", "POST", "/", req, resp)
	return
}

// ConfirmPublicVirtualInterface accept ownership of a public virtual
// interface created by another customer. After the virtual interface owner
// calls this function, the specified virtual interface will be created and
// made available for handling traffic.
func (c *DirectConnect) ConfirmPublicVirtualInterface(req ConfirmPublicVirtualInterfaceRequest) (resp *ConfirmPublicVirtualInterfaceResponse, err error) {
	resp = &ConfirmPublicVirtualInterfaceResponse{}
	err = c.client.Do("ConfirmPublicVirtualInterface", "POST", "/", req, resp)
	return
}

// CreateConnection creates a new connection between the customer network
// and a specific AWS Direct Connect location. A connection links your
// internal network to an AWS Direct Connect location over a standard 1
// gigabit or 10 gigabit Ethernet fiber-optic cable. One end of the cable
// is connected to your router, the other to an AWS Direct Connect router.
// An AWS Direct Connect location provides access to Amazon Web Services in
// the region it is associated with. You can establish connections with AWS
// Direct Connect locations in multiple regions, but a connection in one
// region does not provide connectivity to other regions.
func (c *DirectConnect) CreateConnection(req CreateConnectionRequest) (resp *Connection, err error) {
	resp = &Connection{}
	err = c.client.Do("CreateConnection", "POST", "/", req, resp)
	return
}

// CreateInterconnect creates a new interconnect between a AWS Direct
// Connect partner's network and a specific AWS Direct Connect location. An
// interconnect is a connection which is capable of hosting other
// connections. The AWS Direct Connect partner can use an interconnect to
// provide sub-1Gbps AWS Direct Connect service to tier 2 customers who do
// not have their own connections. Like a standard connection, an
// interconnect links the AWS Direct Connect partner's network to an AWS
// Direct Connect location over a standard 1 Gbps or 10 Gbps Ethernet
// fiber-optic cable. One end is connected to the partner's router, the
// other to an AWS Direct Connect router. For each end customer, the AWS
// Direct Connect partner provisions a connection on their interconnect by
// calling AllocateConnectionOnInterconnect. The end customer can then
// connect to AWS resources by creating a virtual interface on their
// connection, using the assigned to them by the AWS Direct Connect
// partner.
func (c *DirectConnect) CreateInterconnect(req CreateInterconnectRequest) (resp *Interconnect, err error) {
	resp = &Interconnect{}
	err = c.client.Do("CreateInterconnect", "POST", "/", req, resp)
	return
}

// CreatePrivateVirtualInterface creates a new private virtual interface. A
// virtual interface is the that transports AWS Direct Connect traffic. A
// private virtual interface supports sending traffic to a single virtual
// private cloud
func (c *DirectConnect) CreatePrivateVirtualInterface(req CreatePrivateVirtualInterfaceRequest) (resp *VirtualInterface, err error) {
	resp = &VirtualInterface{}
	err = c.client.Do("CreatePrivateVirtualInterface", "POST", "/", req, resp)
	return
}

// CreatePublicVirtualInterface creates a new public virtual interface. A
// virtual interface is the that transports AWS Direct Connect traffic. A
// public virtual interface supports sending traffic to public services of
// AWS such as Amazon Simple Storage Service (Amazon S3).
func (c *DirectConnect) CreatePublicVirtualInterface(req CreatePublicVirtualInterfaceRequest) (resp *VirtualInterface, err error) {
	resp = &VirtualInterface{}
	err = c.client.Do("CreatePublicVirtualInterface", "POST", "/", req, resp)
	return
}

// DeleteConnection deletes the connection. Deleting a connection only
// stops the AWS Direct Connect port hour and data transfer charges. You
// need to cancel separately with the providers any services or charges for
// cross-connects or network circuits that connect you to the AWS Direct
// Connect location.
func (c *DirectConnect) DeleteConnection(req DeleteConnectionRequest) (resp *Connection, err error) {
	resp = &Connection{}
	err = c.client.Do("DeleteConnection", "POST", "/", req, resp)
	return
}

// DeleteInterconnect is undocumented.
func (c *DirectConnect) DeleteInterconnect(req DeleteInterconnectRequest) (resp *DeleteInterconnectResponse, err error) {
	resp = &DeleteInterconnectResponse{}
	err = c.client.Do("DeleteInterconnect", "POST", "/", req, resp)
	return
}

// DeleteVirtualInterface is undocumented.
func (c *DirectConnect) DeleteVirtualInterface(req DeleteVirtualInterfaceRequest) (resp *DeleteVirtualInterfaceResponse, err error) {
	resp = &DeleteVirtualInterfaceResponse{}
	err = c.client.Do("DeleteVirtualInterface", "POST", "/", req, resp)
	return
}

// DescribeConnections displays all connections in this region. If a
// connection ID is provided, the call returns only that particular
// connection.
func (c *DirectConnect) DescribeConnections(req DescribeConnectionsRequest) (resp *Connections, err error) {
	resp = &Connections{}
	err = c.client.Do("DescribeConnections", "POST", "/", req, resp)
	return
}

// DescribeConnectionsOnInterconnect return a list of connections that have
// been provisioned on the given interconnect.
func (c *DirectConnect) DescribeConnectionsOnInterconnect(req DescribeConnectionsOnInterconnectRequest) (resp *Connections, err error) {
	resp = &Connections{}
	err = c.client.Do("DescribeConnectionsOnInterconnect", "POST", "/", req, resp)
	return
}

// DescribeInterconnects returns a list of interconnects owned by the AWS
// account. If an interconnect ID is provided, it will only return this
// particular interconnect.
func (c *DirectConnect) DescribeInterconnects(req DescribeInterconnectsRequest) (resp *Interconnects, err error) {
	resp = &Interconnects{}
	err = c.client.Do("DescribeInterconnects", "POST", "/", req, resp)
	return
}

// DescribeLocations returns the list of AWS Direct Connect locations in
// the current AWS region. These are the locations that may be selected
// when calling CreateConnection or CreateInterconnect.
func (c *DirectConnect) DescribeLocations() (resp *Locations, err error) {
	resp = &Locations{}
	err = c.client.Do("DescribeLocations", "POST", "/", nil, resp)
	return
}

// DescribeVirtualGateways returns a list of virtual private gateways owned
// by the AWS account. You can create one or more AWS Direct Connect
// private virtual interfaces linking to a virtual private gateway. A
// virtual private gateway can be managed via Amazon Virtual Private Cloud
// console or the EC2 CreateVpnGateway action.
func (c *DirectConnect) DescribeVirtualGateways() (resp *VirtualGateways, err error) {
	resp = &VirtualGateways{}
	err = c.client.Do("DescribeVirtualGateways", "POST", "/", nil, resp)
	return
}

// DescribeVirtualInterfaces displays all virtual interfaces for an AWS
// account. Virtual interfaces deleted fewer than 15 minutes before
// DescribeVirtualInterfaces is called are also returned. If a connection
// ID is included then only virtual interfaces associated with this
// connection will be returned. If a virtual interface ID is included then
// only a single virtual interface will be returned. A virtual interface
// transmits the traffic between the AWS Direct Connect location and the
// customer. If a connection ID is provided, only virtual interfaces
// provisioned on the specified connection will be returned. If a virtual
// interface ID is provided, only this particular virtual interface will be
// returned.
func (c *DirectConnect) DescribeVirtualInterfaces(req DescribeVirtualInterfacesRequest) (resp *VirtualInterfaces, err error) {
	resp = &VirtualInterfaces{}
	err = c.client.Do("DescribeVirtualInterfaces", "POST", "/", req, resp)
	return
}

// AllocateConnectionOnInterconnectRequest is undocumented.
type AllocateConnectionOnInterconnectRequest struct {
	Bandwidth      string `json:"bandwidth"`
	ConnectionName string `json:"connectionName"`
	InterconnectID string `json:"interconnectId"`
	OwnerAccount   string `json:"ownerAccount"`
	Vlan           int    `json:"vlan"`
}

// AllocatePrivateVirtualInterfaceRequest is undocumented.
type AllocatePrivateVirtualInterfaceRequest struct {
	ConnectionID                         string                               `json:"connectionId"`
	NewPrivateVirtualInterfaceAllocation NewPrivateVirtualInterfaceAllocation `json:"newPrivateVirtualInterfaceAllocation"`
	OwnerAccount                         string                               `json:"ownerAccount"`
}

// AllocatePublicVirtualInterfaceRequest is undocumented.
type AllocatePublicVirtualInterfaceRequest struct {
	ConnectionID                        string                              `json:"connectionId"`
	NewPublicVirtualInterfaceAllocation NewPublicVirtualInterfaceAllocation `json:"newPublicVirtualInterfaceAllocation"`
	OwnerAccount                        string                              `json:"ownerAccount"`
}

// ConfirmConnectionRequest is undocumented.
type ConfirmConnectionRequest struct {
	ConnectionID string `json:"connectionId"`
}

// ConfirmConnectionResponse is undocumented.
type ConfirmConnectionResponse struct {
	ConnectionState string `json:"connectionState,omitempty"`
}

// ConfirmPrivateVirtualInterfaceRequest is undocumented.
type ConfirmPrivateVirtualInterfaceRequest struct {
	VirtualGatewayID   string `json:"virtualGatewayId"`
	VirtualInterfaceID string `json:"virtualInterfaceId"`
}

// ConfirmPrivateVirtualInterfaceResponse is undocumented.
type ConfirmPrivateVirtualInterfaceResponse struct {
	VirtualInterfaceState string `json:"virtualInterfaceState,omitempty"`
}

// ConfirmPublicVirtualInterfaceRequest is undocumented.
type ConfirmPublicVirtualInterfaceRequest struct {
	VirtualInterfaceID string `json:"virtualInterfaceId"`
}

// ConfirmPublicVirtualInterfaceResponse is undocumented.
type ConfirmPublicVirtualInterfaceResponse struct {
	VirtualInterfaceState string `json:"virtualInterfaceState,omitempty"`
}

// Connection is undocumented.
type Connection struct {
	Bandwidth       string `json:"bandwidth,omitempty"`
	ConnectionID    string `json:"connectionId,omitempty"`
	ConnectionName  string `json:"connectionName,omitempty"`
	ConnectionState string `json:"connectionState,omitempty"`
	Location        string `json:"location,omitempty"`
	OwnerAccount    string `json:"ownerAccount,omitempty"`
	PartnerName     string `json:"partnerName,omitempty"`
	Region          string `json:"region,omitempty"`
	Vlan            int    `json:"vlan,omitempty"`
}

// Connections is undocumented.
type Connections struct {
	Connections []Connection `json:"connections,omitempty"`
}

// CreateConnectionRequest is undocumented.
type CreateConnectionRequest struct {
	Bandwidth      string `json:"bandwidth"`
	ConnectionName string `json:"connectionName"`
	Location       string `json:"location"`
}

// CreateInterconnectRequest is undocumented.
type CreateInterconnectRequest struct {
	Bandwidth        string `json:"bandwidth"`
	InterconnectName string `json:"interconnectName"`
	Location         string `json:"location"`
}

// CreatePrivateVirtualInterfaceRequest is undocumented.
type CreatePrivateVirtualInterfaceRequest struct {
	ConnectionID               string                     `json:"connectionId"`
	NewPrivateVirtualInterface NewPrivateVirtualInterface `json:"newPrivateVirtualInterface"`
}

// CreatePublicVirtualInterfaceRequest is undocumented.
type CreatePublicVirtualInterfaceRequest struct {
	ConnectionID              string                    `json:"connectionId"`
	NewPublicVirtualInterface NewPublicVirtualInterface `json:"newPublicVirtualInterface"`
}

// DeleteConnectionRequest is undocumented.
type DeleteConnectionRequest struct {
	ConnectionID string `json:"connectionId"`
}

// DeleteInterconnectRequest is undocumented.
type DeleteInterconnectRequest struct {
	InterconnectID string `json:"interconnectId"`
}

// DeleteInterconnectResponse is undocumented.
type DeleteInterconnectResponse struct {
	InterconnectState string `json:"interconnectState,omitempty"`
}

// DeleteVirtualInterfaceRequest is undocumented.
type DeleteVirtualInterfaceRequest struct {
	VirtualInterfaceID string `json:"virtualInterfaceId"`
}

// DeleteVirtualInterfaceResponse is undocumented.
type DeleteVirtualInterfaceResponse struct {
	VirtualInterfaceState string `json:"virtualInterfaceState,omitempty"`
}

// DescribeConnectionsOnInterconnectRequest is undocumented.
type DescribeConnectionsOnInterconnectRequest struct {
	InterconnectID string `json:"interconnectId"`
}

// DescribeConnectionsRequest is undocumented.
type DescribeConnectionsRequest struct {
	ConnectionID string `json:"connectionId,omitempty"`
}

// DescribeInterconnectsRequest is undocumented.
type DescribeInterconnectsRequest struct {
	InterconnectID string `json:"interconnectId,omitempty"`
}

// DescribeVirtualInterfacesRequest is undocumented.
type DescribeVirtualInterfacesRequest struct {
	ConnectionID       string `json:"connectionId,omitempty"`
	VirtualInterfaceID string `json:"virtualInterfaceId,omitempty"`
}

// Interconnect is undocumented.
type Interconnect struct {
	Bandwidth         string `json:"bandwidth,omitempty"`
	InterconnectID    string `json:"interconnectId,omitempty"`
	InterconnectName  string `json:"interconnectName,omitempty"`
	InterconnectState string `json:"interconnectState,omitempty"`
	Location          string `json:"location,omitempty"`
	Region            string `json:"region,omitempty"`
}

// Interconnects is undocumented.
type Interconnects struct {
	Interconnects []Interconnect `json:"interconnects,omitempty"`
}

// Location is undocumented.
type Location struct {
	LocationCode string `json:"locationCode,omitempty"`
	LocationName string `json:"locationName,omitempty"`
}

// Locations is undocumented.
type Locations struct {
	Locations []Location `json:"locations,omitempty"`
}

// NewPrivateVirtualInterface is undocumented.
type NewPrivateVirtualInterface struct {
	AmazonAddress        string `json:"amazonAddress,omitempty"`
	Asn                  int    `json:"asn"`
	AuthKey              string `json:"authKey,omitempty"`
	CustomerAddress      string `json:"customerAddress,omitempty"`
	VirtualGatewayID     string `json:"virtualGatewayId"`
	VirtualInterfaceName string `json:"virtualInterfaceName"`
	Vlan                 int    `json:"vlan"`
}

// NewPrivateVirtualInterfaceAllocation is undocumented.
type NewPrivateVirtualInterfaceAllocation struct {
	AmazonAddress        string `json:"amazonAddress,omitempty"`
	Asn                  int    `json:"asn"`
	AuthKey              string `json:"authKey,omitempty"`
	CustomerAddress      string `json:"customerAddress,omitempty"`
	VirtualInterfaceName string `json:"virtualInterfaceName"`
	Vlan                 int    `json:"vlan"`
}

// NewPublicVirtualInterface is undocumented.
type NewPublicVirtualInterface struct {
	AmazonAddress        string              `json:"amazonAddress"`
	Asn                  int                 `json:"asn"`
	AuthKey              string              `json:"authKey,omitempty"`
	CustomerAddress      string              `json:"customerAddress"`
	RouteFilterPrefixes  []RouteFilterPrefix `json:"routeFilterPrefixes"`
	VirtualInterfaceName string              `json:"virtualInterfaceName"`
	Vlan                 int                 `json:"vlan"`
}

// NewPublicVirtualInterfaceAllocation is undocumented.
type NewPublicVirtualInterfaceAllocation struct {
	AmazonAddress        string              `json:"amazonAddress"`
	Asn                  int                 `json:"asn"`
	AuthKey              string              `json:"authKey,omitempty"`
	CustomerAddress      string              `json:"customerAddress"`
	RouteFilterPrefixes  []RouteFilterPrefix `json:"routeFilterPrefixes"`
	VirtualInterfaceName string              `json:"virtualInterfaceName"`
	Vlan                 int                 `json:"vlan"`
}

// RouteFilterPrefix is undocumented.
type RouteFilterPrefix struct {
	Cidr string `json:"cidr,omitempty"`
}

// VirtualGateway is undocumented.
type VirtualGateway struct {
	VirtualGatewayID    string `json:"virtualGatewayId,omitempty"`
	VirtualGatewayState string `json:"virtualGatewayState,omitempty"`
}

// VirtualGateways is undocumented.
type VirtualGateways struct {
	VirtualGateways []VirtualGateway `json:"virtualGateways,omitempty"`
}

// VirtualInterface is undocumented.
type VirtualInterface struct {
	AmazonAddress         string              `json:"amazonAddress,omitempty"`
	Asn                   int                 `json:"asn,omitempty"`
	AuthKey               string              `json:"authKey,omitempty"`
	ConnectionID          string              `json:"connectionId,omitempty"`
	CustomerAddress       string              `json:"customerAddress,omitempty"`
	CustomerRouterConfig  string              `json:"customerRouterConfig,omitempty"`
	Location              string              `json:"location,omitempty"`
	OwnerAccount          string              `json:"ownerAccount,omitempty"`
	RouteFilterPrefixes   []RouteFilterPrefix `json:"routeFilterPrefixes,omitempty"`
	VirtualGatewayID      string              `json:"virtualGatewayId,omitempty"`
	VirtualInterfaceID    string              `json:"virtualInterfaceId,omitempty"`
	VirtualInterfaceName  string              `json:"virtualInterfaceName,omitempty"`
	VirtualInterfaceState string              `json:"virtualInterfaceState,omitempty"`
	VirtualInterfaceType  string              `json:"virtualInterfaceType,omitempty"`
	Vlan                  int                 `json:"vlan,omitempty"`
}

// VirtualInterfaces is undocumented.
type VirtualInterfaces struct {
	VirtualInterfaces []VirtualInterface `json:"virtualInterfaces,omitempty"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
