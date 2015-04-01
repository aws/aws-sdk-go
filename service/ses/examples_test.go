package ses_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/ses"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleSES_DeleteIdentity() {
	svc := ses.New(nil)

	params := &ses.DeleteIdentityInput{
		Identity: aws.String("Identity"), // Required
	}
	resp, err := svc.DeleteIdentity(params)

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

func ExampleSES_DeleteVerifiedEmailAddress() {
	svc := ses.New(nil)

	params := &ses.DeleteVerifiedEmailAddressInput{
		EmailAddress: aws.String("Address"), // Required
	}
	resp, err := svc.DeleteVerifiedEmailAddress(params)

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

func ExampleSES_GetIdentityDKIMAttributes() {
	svc := ses.New(nil)

	params := &ses.GetIdentityDKIMAttributesInput{
		Identities: []*string{ // Required
			aws.String("Identity"), // Required
			// More values...
		},
	}
	resp, err := svc.GetIdentityDKIMAttributes(params)

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

func ExampleSES_GetIdentityNotificationAttributes() {
	svc := ses.New(nil)

	params := &ses.GetIdentityNotificationAttributesInput{
		Identities: []*string{ // Required
			aws.String("Identity"), // Required
			// More values...
		},
	}
	resp, err := svc.GetIdentityNotificationAttributes(params)

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

func ExampleSES_GetIdentityVerificationAttributes() {
	svc := ses.New(nil)

	params := &ses.GetIdentityVerificationAttributesInput{
		Identities: []*string{ // Required
			aws.String("Identity"), // Required
			// More values...
		},
	}
	resp, err := svc.GetIdentityVerificationAttributes(params)

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

func ExampleSES_GetSendQuota() {
	svc := ses.New(nil)

	var params *ses.GetSendQuotaInput
	resp, err := svc.GetSendQuota(params)

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

func ExampleSES_GetSendStatistics() {
	svc := ses.New(nil)

	var params *ses.GetSendStatisticsInput
	resp, err := svc.GetSendStatistics(params)

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

func ExampleSES_ListIdentities() {
	svc := ses.New(nil)

	params := &ses.ListIdentitiesInput{
		IdentityType: aws.String("IdentityType"),
		MaxItems:     aws.Long(1),
		NextToken:    aws.String("NextToken"),
	}
	resp, err := svc.ListIdentities(params)

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

func ExampleSES_ListVerifiedEmailAddresses() {
	svc := ses.New(nil)

	var params *ses.ListVerifiedEmailAddressesInput
	resp, err := svc.ListVerifiedEmailAddresses(params)

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

func ExampleSES_SendEmail() {
	svc := ses.New(nil)

	params := &ses.SendEmailInput{
		Destination: &ses.Destination{ // Required
			BCCAddresses: []*string{
				aws.String("Address"), // Required
				// More values...
			},
			CCAddresses: []*string{
				aws.String("Address"), // Required
				// More values...
			},
			ToAddresses: []*string{
				aws.String("Address"), // Required
				// More values...
			},
		},
		Message: &ses.Message{ // Required
			Body: &ses.Body{ // Required
				HTML: &ses.Content{
					Data:    aws.String("MessageData"), // Required
					Charset: aws.String("Charset"),
				},
				Text: &ses.Content{
					Data:    aws.String("MessageData"), // Required
					Charset: aws.String("Charset"),
				},
			},
			Subject: &ses.Content{ // Required
				Data:    aws.String("MessageData"), // Required
				Charset: aws.String("Charset"),
			},
		},
		Source: aws.String("Address"), // Required
		ReplyToAddresses: []*string{
			aws.String("Address"), // Required
			// More values...
		},
		ReturnPath: aws.String("Address"),
	}
	resp, err := svc.SendEmail(params)

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

func ExampleSES_SendRawEmail() {
	svc := ses.New(nil)

	params := &ses.SendRawEmailInput{
		RawMessage: &ses.RawMessage{ // Required
			Data: []byte("PAYLOAD"), // Required
		},
		Destinations: []*string{
			aws.String("Address"), // Required
			// More values...
		},
		Source: aws.String("Address"),
	}
	resp, err := svc.SendRawEmail(params)

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

func ExampleSES_SetIdentityDKIMEnabled() {
	svc := ses.New(nil)

	params := &ses.SetIdentityDKIMEnabledInput{
		DKIMEnabled: aws.Boolean(true),      // Required
		Identity:    aws.String("Identity"), // Required
	}
	resp, err := svc.SetIdentityDKIMEnabled(params)

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

func ExampleSES_SetIdentityFeedbackForwardingEnabled() {
	svc := ses.New(nil)

	params := &ses.SetIdentityFeedbackForwardingEnabledInput{
		ForwardingEnabled: aws.Boolean(true),      // Required
		Identity:          aws.String("Identity"), // Required
	}
	resp, err := svc.SetIdentityFeedbackForwardingEnabled(params)

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

func ExampleSES_SetIdentityNotificationTopic() {
	svc := ses.New(nil)

	params := &ses.SetIdentityNotificationTopicInput{
		Identity:         aws.String("Identity"),         // Required
		NotificationType: aws.String("NotificationType"), // Required
		SNSTopic:         aws.String("NotificationTopic"),
	}
	resp, err := svc.SetIdentityNotificationTopic(params)

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

func ExampleSES_VerifyDomainDKIM() {
	svc := ses.New(nil)

	params := &ses.VerifyDomainDKIMInput{
		Domain: aws.String("Domain"), // Required
	}
	resp, err := svc.VerifyDomainDKIM(params)

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

func ExampleSES_VerifyDomainIdentity() {
	svc := ses.New(nil)

	params := &ses.VerifyDomainIdentityInput{
		Domain: aws.String("Domain"), // Required
	}
	resp, err := svc.VerifyDomainIdentity(params)

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

func ExampleSES_VerifyEmailAddress() {
	svc := ses.New(nil)

	params := &ses.VerifyEmailAddressInput{
		EmailAddress: aws.String("Address"), // Required
	}
	resp, err := svc.VerifyEmailAddress(params)

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

func ExampleSES_VerifyEmailIdentity() {
	svc := ses.New(nil)

	params := &ses.VerifyEmailIdentityInput{
		EmailAddress: aws.String("Address"), // Required
	}
	resp, err := svc.VerifyEmailIdentity(params)

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