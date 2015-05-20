package cognitoidentity_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/cognitoidentity"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCognitoIdentity_CreateIdentityPool() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.CreateIdentityPoolInput{
		AllowUnauthenticatedIdentities: aws.Boolean(true),              // Required
		IdentityPoolName:               aws.String("IdentityPoolName"), // Required
		DeveloperProviderName:          aws.String("DeveloperProviderName"),
		OpenIDConnectProviderARNs: []*string{
			aws.String("ARNString"), // Required
			// More values...
		},
		SupportedLoginProviders: &map[string]*string{
			"Key": aws.String("IdentityProviderId"), // Required
			// More values...
		},
	}
	resp, err := svc.CreateIdentityPool(params)

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

func ExampleCognitoIdentity_DeleteIdentityPool() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.DeleteIdentityPoolInput{
		IdentityPoolID: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.DeleteIdentityPool(params)

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

func ExampleCognitoIdentity_DescribeIdentity() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.DescribeIdentityInput{
		IdentityID: aws.String("IdentityId"), // Required
	}
	resp, err := svc.DescribeIdentity(params)

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

func ExampleCognitoIdentity_DescribeIdentityPool() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.DescribeIdentityPoolInput{
		IdentityPoolID: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.DescribeIdentityPool(params)

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

func ExampleCognitoIdentity_GetCredentialsForIdentity() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.GetCredentialsForIdentityInput{
		IdentityID: aws.String("IdentityId"), // Required
		Logins: &map[string]*string{
			"Key": aws.String("IdentityProviderToken"), // Required
			// More values...
		},
	}
	resp, err := svc.GetCredentialsForIdentity(params)

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

func ExampleCognitoIdentity_GetID() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.GetIDInput{
		IdentityPoolID: aws.String("IdentityPoolId"), // Required
		AccountID:      aws.String("AccountId"),
		Logins: &map[string]*string{
			"Key": aws.String("IdentityProviderToken"), // Required
			// More values...
		},
	}
	resp, err := svc.GetID(params)

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

func ExampleCognitoIdentity_GetIdentityPoolRoles() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.GetIdentityPoolRolesInput{
		IdentityPoolID: aws.String("IdentityPoolId"),
	}
	resp, err := svc.GetIdentityPoolRoles(params)

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

func ExampleCognitoIdentity_GetOpenIDToken() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.GetOpenIDTokenInput{
		IdentityID: aws.String("IdentityId"), // Required
		Logins: &map[string]*string{
			"Key": aws.String("IdentityProviderToken"), // Required
			// More values...
		},
	}
	resp, err := svc.GetOpenIDToken(params)

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

func ExampleCognitoIdentity_GetOpenIDTokenForDeveloperIdentity() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.GetOpenIDTokenForDeveloperIdentityInput{
		IdentityPoolID: aws.String("IdentityPoolId"), // Required
		Logins: &map[string]*string{ // Required
			"Key": aws.String("IdentityProviderToken"), // Required
			// More values...
		},
		IdentityID:    aws.String("IdentityId"),
		TokenDuration: aws.Long(1),
	}
	resp, err := svc.GetOpenIDTokenForDeveloperIdentity(params)

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

func ExampleCognitoIdentity_ListIdentities() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.ListIdentitiesInput{
		IdentityPoolID: aws.String("IdentityPoolId"), // Required
		MaxResults:     aws.Long(1),                  // Required
		NextToken:      aws.String("PaginationKey"),
	}
	resp, err := svc.ListIdentities(params)

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

func ExampleCognitoIdentity_ListIdentityPools() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.ListIdentityPoolsInput{
		MaxResults: aws.Long(1), // Required
		NextToken:  aws.String("PaginationKey"),
	}
	resp, err := svc.ListIdentityPools(params)

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

func ExampleCognitoIdentity_LookupDeveloperIdentity() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.LookupDeveloperIdentityInput{
		IdentityPoolID:          aws.String("IdentityPoolId"), // Required
		DeveloperUserIdentifier: aws.String("DeveloperUserIdentifier"),
		IdentityID:              aws.String("IdentityId"),
		MaxResults:              aws.Long(1),
		NextToken:               aws.String("PaginationKey"),
	}
	resp, err := svc.LookupDeveloperIdentity(params)

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

func ExampleCognitoIdentity_MergeDeveloperIdentities() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.MergeDeveloperIdentitiesInput{
		DestinationUserIdentifier: aws.String("DeveloperUserIdentifier"), // Required
		DeveloperProviderName:     aws.String("DeveloperProviderName"),   // Required
		IdentityPoolID:            aws.String("IdentityPoolId"),          // Required
		SourceUserIdentifier:      aws.String("DeveloperUserIdentifier"), // Required
	}
	resp, err := svc.MergeDeveloperIdentities(params)

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

func ExampleCognitoIdentity_SetIdentityPoolRoles() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.SetIdentityPoolRolesInput{
		IdentityPoolID: aws.String("IdentityPoolId"), // Required
		Roles: &map[string]*string{ // Required
			"Key": aws.String("ARNString"), // Required
			// More values...
		},
	}
	resp, err := svc.SetIdentityPoolRoles(params)

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

func ExampleCognitoIdentity_UnlinkDeveloperIdentity() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.UnlinkDeveloperIdentityInput{
		DeveloperProviderName:   aws.String("DeveloperProviderName"),   // Required
		DeveloperUserIdentifier: aws.String("DeveloperUserIdentifier"), // Required
		IdentityID:              aws.String("IdentityId"),              // Required
		IdentityPoolID:          aws.String("IdentityPoolId"),          // Required
	}
	resp, err := svc.UnlinkDeveloperIdentity(params)

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

func ExampleCognitoIdentity_UnlinkIdentity() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.UnlinkIdentityInput{
		IdentityID: aws.String("IdentityId"), // Required
		Logins: &map[string]*string{ // Required
			"Key": aws.String("IdentityProviderToken"), // Required
			// More values...
		},
		LoginsToRemove: []*string{ // Required
			aws.String("IdentityProviderName"), // Required
			// More values...
		},
	}
	resp, err := svc.UnlinkIdentity(params)

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

func ExampleCognitoIdentity_UpdateIdentityPool() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.IdentityPool{
		AllowUnauthenticatedIdentities: aws.Boolean(true),              // Required
		IdentityPoolID:                 aws.String("IdentityPoolId"),   // Required
		IdentityPoolName:               aws.String("IdentityPoolName"), // Required
		DeveloperProviderName:          aws.String("DeveloperProviderName"),
		OpenIDConnectProviderARNs: []*string{
			aws.String("ARNString"), // Required
			// More values...
		},
		SupportedLoginProviders: &map[string]*string{
			"Key": aws.String("IdentityProviderId"), // Required
			// More values...
		},
	}
	resp, err := svc.UpdateIdentityPool(params)

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