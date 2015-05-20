package directconnect_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/directconnect"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleDirectConnect_AllocateConnectionOnInterconnect() {
	svc := directconnect.New(nil)

	params := &directconnect.AllocateConnectionOnInterconnectInput{
		Bandwidth:      aws.String("Bandwidth"),      // Required
		ConnectionName: aws.String("ConnectionName"), // Required
		InterconnectID: aws.String("InterconnectId"), // Required
		OwnerAccount:   aws.String("OwnerAccount"),   // Required
		VLAN:           aws.Long(1),                  // Required
	}
	resp, err := svc.AllocateConnectionOnInterconnect(params)

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

func ExampleDirectConnect_AllocatePrivateVirtualInterface() {
	svc := directconnect.New(nil)

	params := &directconnect.AllocatePrivateVirtualInterfaceInput{
		ConnectionID: aws.String("ConnectionId"), // Required
		NewPrivateVirtualInterfaceAllocation: &directconnect.NewPrivateVirtualInterfaceAllocation{ // Required
			ASN:                  aws.Long(1),                        // Required
			VLAN:                 aws.Long(1),                        // Required
			VirtualInterfaceName: aws.String("VirtualInterfaceName"), // Required
			AmazonAddress:        aws.String("AmazonAddress"),
			AuthKey:              aws.String("BGPAuthKey"),
			CustomerAddress:      aws.String("CustomerAddress"),
		},
		OwnerAccount: aws.String("OwnerAccount"), // Required
	}
	resp, err := svc.AllocatePrivateVirtualInterface(params)

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

func ExampleDirectConnect_AllocatePublicVirtualInterface() {
	svc := directconnect.New(nil)

	params := &directconnect.AllocatePublicVirtualInterfaceInput{
		ConnectionID: aws.String("ConnectionId"), // Required
		NewPublicVirtualInterfaceAllocation: &directconnect.NewPublicVirtualInterfaceAllocation{ // Required
			ASN:             aws.Long(1),                   // Required
			AmazonAddress:   aws.String("AmazonAddress"),   // Required
			CustomerAddress: aws.String("CustomerAddress"), // Required
			RouteFilterPrefixes: []*directconnect.RouteFilterPrefix{ // Required
				&directconnect.RouteFilterPrefix{ // Required
					CIDR: aws.String("CIDR"),
				},
				// More values...
			},
			VLAN:                 aws.Long(1),                        // Required
			VirtualInterfaceName: aws.String("VirtualInterfaceName"), // Required
			AuthKey:              aws.String("BGPAuthKey"),
		},
		OwnerAccount: aws.String("OwnerAccount"), // Required
	}
	resp, err := svc.AllocatePublicVirtualInterface(params)

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

func ExampleDirectConnect_ConfirmConnection() {
	svc := directconnect.New(nil)

	params := &directconnect.ConfirmConnectionInput{
		ConnectionID: aws.String("ConnectionId"), // Required
	}
	resp, err := svc.ConfirmConnection(params)

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

func ExampleDirectConnect_ConfirmPrivateVirtualInterface() {
	svc := directconnect.New(nil)

	params := &directconnect.ConfirmPrivateVirtualInterfaceInput{
		VirtualGatewayID:   aws.String("VirtualGatewayId"),   // Required
		VirtualInterfaceID: aws.String("VirtualInterfaceId"), // Required
	}
	resp, err := svc.ConfirmPrivateVirtualInterface(params)

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

func ExampleDirectConnect_ConfirmPublicVirtualInterface() {
	svc := directconnect.New(nil)

	params := &directconnect.ConfirmPublicVirtualInterfaceInput{
		VirtualInterfaceID: aws.String("VirtualInterfaceId"), // Required
	}
	resp, err := svc.ConfirmPublicVirtualInterface(params)

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

func ExampleDirectConnect_CreateConnection() {
	svc := directconnect.New(nil)

	params := &directconnect.CreateConnectionInput{
		Bandwidth:      aws.String("Bandwidth"),      // Required
		ConnectionName: aws.String("ConnectionName"), // Required
		Location:       aws.String("LocationCode"),   // Required
	}
	resp, err := svc.CreateConnection(params)

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

func ExampleDirectConnect_CreateInterconnect() {
	svc := directconnect.New(nil)

	params := &directconnect.CreateInterconnectInput{
		Bandwidth:        aws.String("Bandwidth"),        // Required
		InterconnectName: aws.String("InterconnectName"), // Required
		Location:         aws.String("LocationCode"),     // Required
	}
	resp, err := svc.CreateInterconnect(params)

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

func ExampleDirectConnect_CreatePrivateVirtualInterface() {
	svc := directconnect.New(nil)

	params := &directconnect.CreatePrivateVirtualInterfaceInput{
		ConnectionID: aws.String("ConnectionId"), // Required
		NewPrivateVirtualInterface: &directconnect.NewPrivateVirtualInterface{ // Required
			ASN:                  aws.Long(1),                        // Required
			VLAN:                 aws.Long(1),                        // Required
			VirtualGatewayID:     aws.String("VirtualGatewayId"),     // Required
			VirtualInterfaceName: aws.String("VirtualInterfaceName"), // Required
			AmazonAddress:        aws.String("AmazonAddress"),
			AuthKey:              aws.String("BGPAuthKey"),
			CustomerAddress:      aws.String("CustomerAddress"),
		},
	}
	resp, err := svc.CreatePrivateVirtualInterface(params)

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

func ExampleDirectConnect_CreatePublicVirtualInterface() {
	svc := directconnect.New(nil)

	params := &directconnect.CreatePublicVirtualInterfaceInput{
		ConnectionID: aws.String("ConnectionId"), // Required
		NewPublicVirtualInterface: &directconnect.NewPublicVirtualInterface{ // Required
			ASN:             aws.Long(1),                   // Required
			AmazonAddress:   aws.String("AmazonAddress"),   // Required
			CustomerAddress: aws.String("CustomerAddress"), // Required
			RouteFilterPrefixes: []*directconnect.RouteFilterPrefix{ // Required
				&directconnect.RouteFilterPrefix{ // Required
					CIDR: aws.String("CIDR"),
				},
				// More values...
			},
			VLAN:                 aws.Long(1),                        // Required
			VirtualInterfaceName: aws.String("VirtualInterfaceName"), // Required
			AuthKey:              aws.String("BGPAuthKey"),
		},
	}
	resp, err := svc.CreatePublicVirtualInterface(params)

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

func ExampleDirectConnect_DeleteConnection() {
	svc := directconnect.New(nil)

	params := &directconnect.DeleteConnectionInput{
		ConnectionID: aws.String("ConnectionId"), // Required
	}
	resp, err := svc.DeleteConnection(params)

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

func ExampleDirectConnect_DeleteInterconnect() {
	svc := directconnect.New(nil)

	params := &directconnect.DeleteInterconnectInput{
		InterconnectID: aws.String("InterconnectId"), // Required
	}
	resp, err := svc.DeleteInterconnect(params)

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

func ExampleDirectConnect_DeleteVirtualInterface() {
	svc := directconnect.New(nil)

	params := &directconnect.DeleteVirtualInterfaceInput{
		VirtualInterfaceID: aws.String("VirtualInterfaceId"), // Required
	}
	resp, err := svc.DeleteVirtualInterface(params)

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

func ExampleDirectConnect_DescribeConnections() {
	svc := directconnect.New(nil)

	params := &directconnect.DescribeConnectionsInput{
		ConnectionID: aws.String("ConnectionId"),
	}
	resp, err := svc.DescribeConnections(params)

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

func ExampleDirectConnect_DescribeConnectionsOnInterconnect() {
	svc := directconnect.New(nil)

	params := &directconnect.DescribeConnectionsOnInterconnectInput{
		InterconnectID: aws.String("InterconnectId"), // Required
	}
	resp, err := svc.DescribeConnectionsOnInterconnect(params)

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

func ExampleDirectConnect_DescribeInterconnects() {
	svc := directconnect.New(nil)

	params := &directconnect.DescribeInterconnectsInput{
		InterconnectID: aws.String("InterconnectId"),
	}
	resp, err := svc.DescribeInterconnects(params)

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

func ExampleDirectConnect_DescribeLocations() {
	svc := directconnect.New(nil)

	var params *directconnect.DescribeLocationsInput
	resp, err := svc.DescribeLocations(params)

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

func ExampleDirectConnect_DescribeVirtualGateways() {
	svc := directconnect.New(nil)

	var params *directconnect.DescribeVirtualGatewaysInput
	resp, err := svc.DescribeVirtualGateways(params)

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

func ExampleDirectConnect_DescribeVirtualInterfaces() {
	svc := directconnect.New(nil)

	params := &directconnect.DescribeVirtualInterfacesInput{
		ConnectionID:       aws.String("ConnectionId"),
		VirtualInterfaceID: aws.String("VirtualInterfaceId"),
	}
	resp, err := svc.DescribeVirtualInterfaces(params)

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