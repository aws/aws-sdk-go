package route53domains_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/route53domains"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleRoute53Domains_CheckDomainAvailability() {
	svc := route53domains.New(nil)

	params := &route53domains.CheckDomainAvailabilityInput{
		DomainName:  aws.String("DomainName"), // Required
		IDNLangCode: aws.String("LangCode"),
	}
	resp, err := svc.CheckDomainAvailability(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleRoute53Domains_DeleteTagsForDomain() {
	svc := route53domains.New(nil)

	params := &route53domains.DeleteTagsForDomainInput{
		DomainName: aws.String("DomainName"), // Required
		TagsToDelete: []*string{ // Required
			aws.String("TagKey"), // Required
			// More values...
		},
	}
	resp, err := svc.DeleteTagsForDomain(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleRoute53Domains_DisableDomainAutoRenew() {
	svc := route53domains.New(nil)

	params := &route53domains.DisableDomainAutoRenewInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.DisableDomainAutoRenew(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleRoute53Domains_DisableDomainTransferLock() {
	svc := route53domains.New(nil)

	params := &route53domains.DisableDomainTransferLockInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.DisableDomainTransferLock(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleRoute53Domains_EnableDomainAutoRenew() {
	svc := route53domains.New(nil)

	params := &route53domains.EnableDomainAutoRenewInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.EnableDomainAutoRenew(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleRoute53Domains_EnableDomainTransferLock() {
	svc := route53domains.New(nil)

	params := &route53domains.EnableDomainTransferLockInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.EnableDomainTransferLock(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleRoute53Domains_GetDomainDetail() {
	svc := route53domains.New(nil)

	params := &route53domains.GetDomainDetailInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.GetDomainDetail(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleRoute53Domains_GetOperationDetail() {
	svc := route53domains.New(nil)

	params := &route53domains.GetOperationDetailInput{
		OperationID: aws.String("OperationId"), // Required
	}
	resp, err := svc.GetOperationDetail(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleRoute53Domains_ListDomains() {
	svc := route53domains.New(nil)

	params := &route53domains.ListDomainsInput{
		Marker:   aws.String("PageMarker"),
		MaxItems: aws.Long(1),
	}
	resp, err := svc.ListDomains(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleRoute53Domains_ListOperations() {
	svc := route53domains.New(nil)

	params := &route53domains.ListOperationsInput{
		Marker:   aws.String("PageMarker"),
		MaxItems: aws.Long(1),
	}
	resp, err := svc.ListOperations(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleRoute53Domains_ListTagsForDomain() {
	svc := route53domains.New(nil)

	params := &route53domains.ListTagsForDomainInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.ListTagsForDomain(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleRoute53Domains_RegisterDomain() {
	svc := route53domains.New(nil)

	params := &route53domains.RegisterDomainInput{
		AdminContact: &route53domains.ContactDetail{ // Required
			AddressLine1: aws.String("AddressLine"),
			AddressLine2: aws.String("AddressLine"),
			City:         aws.String("City"),
			ContactType:  aws.String("ContactType"),
			CountryCode:  aws.String("CountryCode"),
			Email:        aws.String("Email"),
			ExtraParams: []*route53domains.ExtraParam{
				&route53domains.ExtraParam{ // Required
					Name:  aws.String("ExtraParamName"),  // Required
					Value: aws.String("ExtraParamValue"), // Required
				},
				// More values...
			},
			Fax:              aws.String("ContactNumber"),
			FirstName:        aws.String("ContactName"),
			LastName:         aws.String("ContactName"),
			OrganizationName: aws.String("ContactName"),
			PhoneNumber:      aws.String("ContactNumber"),
			State:            aws.String("State"),
			ZipCode:          aws.String("ZipCode"),
		},
		DomainName:      aws.String("DomainName"), // Required
		DurationInYears: aws.Long(1),              // Required
		RegistrantContact: &route53domains.ContactDetail{ // Required
			AddressLine1: aws.String("AddressLine"),
			AddressLine2: aws.String("AddressLine"),
			City:         aws.String("City"),
			ContactType:  aws.String("ContactType"),
			CountryCode:  aws.String("CountryCode"),
			Email:        aws.String("Email"),
			ExtraParams: []*route53domains.ExtraParam{
				&route53domains.ExtraParam{ // Required
					Name:  aws.String("ExtraParamName"),  // Required
					Value: aws.String("ExtraParamValue"), // Required
				},
				// More values...
			},
			Fax:              aws.String("ContactNumber"),
			FirstName:        aws.String("ContactName"),
			LastName:         aws.String("ContactName"),
			OrganizationName: aws.String("ContactName"),
			PhoneNumber:      aws.String("ContactNumber"),
			State:            aws.String("State"),
			ZipCode:          aws.String("ZipCode"),
		},
		TechContact: &route53domains.ContactDetail{ // Required
			AddressLine1: aws.String("AddressLine"),
			AddressLine2: aws.String("AddressLine"),
			City:         aws.String("City"),
			ContactType:  aws.String("ContactType"),
			CountryCode:  aws.String("CountryCode"),
			Email:        aws.String("Email"),
			ExtraParams: []*route53domains.ExtraParam{
				&route53domains.ExtraParam{ // Required
					Name:  aws.String("ExtraParamName"),  // Required
					Value: aws.String("ExtraParamValue"), // Required
				},
				// More values...
			},
			Fax:              aws.String("ContactNumber"),
			FirstName:        aws.String("ContactName"),
			LastName:         aws.String("ContactName"),
			OrganizationName: aws.String("ContactName"),
			PhoneNumber:      aws.String("ContactNumber"),
			State:            aws.String("State"),
			ZipCode:          aws.String("ZipCode"),
		},
		AutoRenew:                       aws.Boolean(true),
		IDNLangCode:                     aws.String("LangCode"),
		PrivacyProtectAdminContact:      aws.Boolean(true),
		PrivacyProtectRegistrantContact: aws.Boolean(true),
		PrivacyProtectTechContact:       aws.Boolean(true),
	}
	resp, err := svc.RegisterDomain(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleRoute53Domains_RetrieveDomainAuthCode() {
	svc := route53domains.New(nil)

	params := &route53domains.RetrieveDomainAuthCodeInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.RetrieveDomainAuthCode(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleRoute53Domains_TransferDomain() {
	svc := route53domains.New(nil)

	params := &route53domains.TransferDomainInput{
		AdminContact: &route53domains.ContactDetail{ // Required
			AddressLine1: aws.String("AddressLine"),
			AddressLine2: aws.String("AddressLine"),
			City:         aws.String("City"),
			ContactType:  aws.String("ContactType"),
			CountryCode:  aws.String("CountryCode"),
			Email:        aws.String("Email"),
			ExtraParams: []*route53domains.ExtraParam{
				&route53domains.ExtraParam{ // Required
					Name:  aws.String("ExtraParamName"),  // Required
					Value: aws.String("ExtraParamValue"), // Required
				},
				// More values...
			},
			Fax:              aws.String("ContactNumber"),
			FirstName:        aws.String("ContactName"),
			LastName:         aws.String("ContactName"),
			OrganizationName: aws.String("ContactName"),
			PhoneNumber:      aws.String("ContactNumber"),
			State:            aws.String("State"),
			ZipCode:          aws.String("ZipCode"),
		},
		DomainName:      aws.String("DomainName"), // Required
		DurationInYears: aws.Long(1),              // Required
		RegistrantContact: &route53domains.ContactDetail{ // Required
			AddressLine1: aws.String("AddressLine"),
			AddressLine2: aws.String("AddressLine"),
			City:         aws.String("City"),
			ContactType:  aws.String("ContactType"),
			CountryCode:  aws.String("CountryCode"),
			Email:        aws.String("Email"),
			ExtraParams: []*route53domains.ExtraParam{
				&route53domains.ExtraParam{ // Required
					Name:  aws.String("ExtraParamName"),  // Required
					Value: aws.String("ExtraParamValue"), // Required
				},
				// More values...
			},
			Fax:              aws.String("ContactNumber"),
			FirstName:        aws.String("ContactName"),
			LastName:         aws.String("ContactName"),
			OrganizationName: aws.String("ContactName"),
			PhoneNumber:      aws.String("ContactNumber"),
			State:            aws.String("State"),
			ZipCode:          aws.String("ZipCode"),
		},
		TechContact: &route53domains.ContactDetail{ // Required
			AddressLine1: aws.String("AddressLine"),
			AddressLine2: aws.String("AddressLine"),
			City:         aws.String("City"),
			ContactType:  aws.String("ContactType"),
			CountryCode:  aws.String("CountryCode"),
			Email:        aws.String("Email"),
			ExtraParams: []*route53domains.ExtraParam{
				&route53domains.ExtraParam{ // Required
					Name:  aws.String("ExtraParamName"),  // Required
					Value: aws.String("ExtraParamValue"), // Required
				},
				// More values...
			},
			Fax:              aws.String("ContactNumber"),
			FirstName:        aws.String("ContactName"),
			LastName:         aws.String("ContactName"),
			OrganizationName: aws.String("ContactName"),
			PhoneNumber:      aws.String("ContactNumber"),
			State:            aws.String("State"),
			ZipCode:          aws.String("ZipCode"),
		},
		AuthCode:    aws.String("DomainAuthCode"),
		AutoRenew:   aws.Boolean(true),
		IDNLangCode: aws.String("LangCode"),
		Nameservers: []*route53domains.Nameserver{
			&route53domains.Nameserver{ // Required
				Name: aws.String("HostName"), // Required
				GlueIPs: []*string{
					aws.String("GlueIp"), // Required
					// More values...
				},
			},
			// More values...
		},
		PrivacyProtectAdminContact:      aws.Boolean(true),
		PrivacyProtectRegistrantContact: aws.Boolean(true),
		PrivacyProtectTechContact:       aws.Boolean(true),
	}
	resp, err := svc.TransferDomain(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleRoute53Domains_UpdateDomainContact() {
	svc := route53domains.New(nil)

	params := &route53domains.UpdateDomainContactInput{
		DomainName: aws.String("DomainName"), // Required
		AdminContact: &route53domains.ContactDetail{
			AddressLine1: aws.String("AddressLine"),
			AddressLine2: aws.String("AddressLine"),
			City:         aws.String("City"),
			ContactType:  aws.String("ContactType"),
			CountryCode:  aws.String("CountryCode"),
			Email:        aws.String("Email"),
			ExtraParams: []*route53domains.ExtraParam{
				&route53domains.ExtraParam{ // Required
					Name:  aws.String("ExtraParamName"),  // Required
					Value: aws.String("ExtraParamValue"), // Required
				},
				// More values...
			},
			Fax:              aws.String("ContactNumber"),
			FirstName:        aws.String("ContactName"),
			LastName:         aws.String("ContactName"),
			OrganizationName: aws.String("ContactName"),
			PhoneNumber:      aws.String("ContactNumber"),
			State:            aws.String("State"),
			ZipCode:          aws.String("ZipCode"),
		},
		RegistrantContact: &route53domains.ContactDetail{
			AddressLine1: aws.String("AddressLine"),
			AddressLine2: aws.String("AddressLine"),
			City:         aws.String("City"),
			ContactType:  aws.String("ContactType"),
			CountryCode:  aws.String("CountryCode"),
			Email:        aws.String("Email"),
			ExtraParams: []*route53domains.ExtraParam{
				&route53domains.ExtraParam{ // Required
					Name:  aws.String("ExtraParamName"),  // Required
					Value: aws.String("ExtraParamValue"), // Required
				},
				// More values...
			},
			Fax:              aws.String("ContactNumber"),
			FirstName:        aws.String("ContactName"),
			LastName:         aws.String("ContactName"),
			OrganizationName: aws.String("ContactName"),
			PhoneNumber:      aws.String("ContactNumber"),
			State:            aws.String("State"),
			ZipCode:          aws.String("ZipCode"),
		},
		TechContact: &route53domains.ContactDetail{
			AddressLine1: aws.String("AddressLine"),
			AddressLine2: aws.String("AddressLine"),
			City:         aws.String("City"),
			ContactType:  aws.String("ContactType"),
			CountryCode:  aws.String("CountryCode"),
			Email:        aws.String("Email"),
			ExtraParams: []*route53domains.ExtraParam{
				&route53domains.ExtraParam{ // Required
					Name:  aws.String("ExtraParamName"),  // Required
					Value: aws.String("ExtraParamValue"), // Required
				},
				// More values...
			},
			Fax:              aws.String("ContactNumber"),
			FirstName:        aws.String("ContactName"),
			LastName:         aws.String("ContactName"),
			OrganizationName: aws.String("ContactName"),
			PhoneNumber:      aws.String("ContactNumber"),
			State:            aws.String("State"),
			ZipCode:          aws.String("ZipCode"),
		},
	}
	resp, err := svc.UpdateDomainContact(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleRoute53Domains_UpdateDomainContactPrivacy() {
	svc := route53domains.New(nil)

	params := &route53domains.UpdateDomainContactPrivacyInput{
		DomainName:        aws.String("DomainName"), // Required
		AdminPrivacy:      aws.Boolean(true),
		RegistrantPrivacy: aws.Boolean(true),
		TechPrivacy:       aws.Boolean(true),
	}
	resp, err := svc.UpdateDomainContactPrivacy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleRoute53Domains_UpdateDomainNameservers() {
	svc := route53domains.New(nil)

	params := &route53domains.UpdateDomainNameserversInput{
		DomainName: aws.String("DomainName"), // Required
		Nameservers: []*route53domains.Nameserver{ // Required
			&route53domains.Nameserver{ // Required
				Name: aws.String("HostName"), // Required
				GlueIPs: []*string{
					aws.String("GlueIp"), // Required
					// More values...
				},
			},
			// More values...
		},
		FIAuthKey: aws.String("FIAuthKey"),
	}
	resp, err := svc.UpdateDomainNameservers(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleRoute53Domains_UpdateTagsForDomain() {
	svc := route53domains.New(nil)

	params := &route53domains.UpdateTagsForDomainInput{
		DomainName: aws.String("DomainName"), // Required
		TagsToUpdate: []*route53domains.Tag{
			&route53domains.Tag{ // Required
				Key:   aws.String("TagKey"),
				Value: aws.String("TagValue"),
			},
			// More values...
		},
	}
	resp, err := svc.UpdateTagsForDomain(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}