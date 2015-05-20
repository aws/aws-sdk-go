package sns_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/sns"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleSNS_AddPermission() {
	svc := sns.New(nil)

	params := &sns.AddPermissionInput{
		AWSAccountID: []*string{ // Required
			aws.String("delegate"), // Required
			// More values...
		},
		ActionName: []*string{ // Required
			aws.String("action"), // Required
			// More values...
		},
		Label:    aws.String("label"),    // Required
		TopicARN: aws.String("topicARN"), // Required
	}
	resp, err := svc.AddPermission(params)

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

func ExampleSNS_ConfirmSubscription() {
	svc := sns.New(nil)

	params := &sns.ConfirmSubscriptionInput{
		Token:                     aws.String("token"),    // Required
		TopicARN:                  aws.String("topicARN"), // Required
		AuthenticateOnUnsubscribe: aws.String("authenticateOnUnsubscribe"),
	}
	resp, err := svc.ConfirmSubscription(params)

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

func ExampleSNS_CreatePlatformApplication() {
	svc := sns.New(nil)

	params := &sns.CreatePlatformApplicationInput{
		Attributes: &map[string]*string{ // Required
			"Key": aws.String("String"), // Required
			// More values...
		},
		Name:     aws.String("String"), // Required
		Platform: aws.String("String"), // Required
	}
	resp, err := svc.CreatePlatformApplication(params)

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

func ExampleSNS_CreatePlatformEndpoint() {
	svc := sns.New(nil)

	params := &sns.CreatePlatformEndpointInput{
		PlatformApplicationARN: aws.String("String"), // Required
		Token: aws.String("String"), // Required
		Attributes: &map[string]*string{
			"Key": aws.String("String"), // Required
			// More values...
		},
		CustomUserData: aws.String("String"),
	}
	resp, err := svc.CreatePlatformEndpoint(params)

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

func ExampleSNS_CreateTopic() {
	svc := sns.New(nil)

	params := &sns.CreateTopicInput{
		Name: aws.String("topicName"), // Required
	}
	resp, err := svc.CreateTopic(params)

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

func ExampleSNS_DeleteEndpoint() {
	svc := sns.New(nil)

	params := &sns.DeleteEndpointInput{
		EndpointARN: aws.String("String"), // Required
	}
	resp, err := svc.DeleteEndpoint(params)

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

func ExampleSNS_DeletePlatformApplication() {
	svc := sns.New(nil)

	params := &sns.DeletePlatformApplicationInput{
		PlatformApplicationARN: aws.String("String"), // Required
	}
	resp, err := svc.DeletePlatformApplication(params)

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

func ExampleSNS_DeleteTopic() {
	svc := sns.New(nil)

	params := &sns.DeleteTopicInput{
		TopicARN: aws.String("topicARN"), // Required
	}
	resp, err := svc.DeleteTopic(params)

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

func ExampleSNS_GetEndpointAttributes() {
	svc := sns.New(nil)

	params := &sns.GetEndpointAttributesInput{
		EndpointARN: aws.String("String"), // Required
	}
	resp, err := svc.GetEndpointAttributes(params)

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

func ExampleSNS_GetPlatformApplicationAttributes() {
	svc := sns.New(nil)

	params := &sns.GetPlatformApplicationAttributesInput{
		PlatformApplicationARN: aws.String("String"), // Required
	}
	resp, err := svc.GetPlatformApplicationAttributes(params)

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

func ExampleSNS_GetSubscriptionAttributes() {
	svc := sns.New(nil)

	params := &sns.GetSubscriptionAttributesInput{
		SubscriptionARN: aws.String("subscriptionARN"), // Required
	}
	resp, err := svc.GetSubscriptionAttributes(params)

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

func ExampleSNS_GetTopicAttributes() {
	svc := sns.New(nil)

	params := &sns.GetTopicAttributesInput{
		TopicARN: aws.String("topicARN"), // Required
	}
	resp, err := svc.GetTopicAttributes(params)

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

func ExampleSNS_ListEndpointsByPlatformApplication() {
	svc := sns.New(nil)

	params := &sns.ListEndpointsByPlatformApplicationInput{
		PlatformApplicationARN: aws.String("String"), // Required
		NextToken:              aws.String("String"),
	}
	resp, err := svc.ListEndpointsByPlatformApplication(params)

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

func ExampleSNS_ListPlatformApplications() {
	svc := sns.New(nil)

	params := &sns.ListPlatformApplicationsInput{
		NextToken: aws.String("String"),
	}
	resp, err := svc.ListPlatformApplications(params)

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

func ExampleSNS_ListSubscriptions() {
	svc := sns.New(nil)

	params := &sns.ListSubscriptionsInput{
		NextToken: aws.String("nextToken"),
	}
	resp, err := svc.ListSubscriptions(params)

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

func ExampleSNS_ListSubscriptionsByTopic() {
	svc := sns.New(nil)

	params := &sns.ListSubscriptionsByTopicInput{
		TopicARN:  aws.String("topicARN"), // Required
		NextToken: aws.String("nextToken"),
	}
	resp, err := svc.ListSubscriptionsByTopic(params)

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

func ExampleSNS_ListTopics() {
	svc := sns.New(nil)

	params := &sns.ListTopicsInput{
		NextToken: aws.String("nextToken"),
	}
	resp, err := svc.ListTopics(params)

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

func ExampleSNS_Publish() {
	svc := sns.New(nil)

	params := &sns.PublishInput{
		Message: aws.String("message"), // Required
		MessageAttributes: &map[string]*sns.MessageAttributeValue{
			"Key": &sns.MessageAttributeValue{ // Required
				DataType:    aws.String("String"), // Required
				BinaryValue: []byte("PAYLOAD"),
				StringValue: aws.String("String"),
			},
			// More values...
		},
		MessageStructure: aws.String("messageStructure"),
		Subject:          aws.String("subject"),
		TargetARN:        aws.String("String"),
		TopicARN:         aws.String("topicARN"),
	}
	resp, err := svc.Publish(params)

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

func ExampleSNS_RemovePermission() {
	svc := sns.New(nil)

	params := &sns.RemovePermissionInput{
		Label:    aws.String("label"),    // Required
		TopicARN: aws.String("topicARN"), // Required
	}
	resp, err := svc.RemovePermission(params)

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

func ExampleSNS_SetEndpointAttributes() {
	svc := sns.New(nil)

	params := &sns.SetEndpointAttributesInput{
		Attributes: &map[string]*string{ // Required
			"Key": aws.String("String"), // Required
			// More values...
		},
		EndpointARN: aws.String("String"), // Required
	}
	resp, err := svc.SetEndpointAttributes(params)

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

func ExampleSNS_SetPlatformApplicationAttributes() {
	svc := sns.New(nil)

	params := &sns.SetPlatformApplicationAttributesInput{
		Attributes: &map[string]*string{ // Required
			"Key": aws.String("String"), // Required
			// More values...
		},
		PlatformApplicationARN: aws.String("String"), // Required
	}
	resp, err := svc.SetPlatformApplicationAttributes(params)

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

func ExampleSNS_SetSubscriptionAttributes() {
	svc := sns.New(nil)

	params := &sns.SetSubscriptionAttributesInput{
		AttributeName:   aws.String("attributeName"),   // Required
		SubscriptionARN: aws.String("subscriptionARN"), // Required
		AttributeValue:  aws.String("attributeValue"),
	}
	resp, err := svc.SetSubscriptionAttributes(params)

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

func ExampleSNS_SetTopicAttributes() {
	svc := sns.New(nil)

	params := &sns.SetTopicAttributesInput{
		AttributeName:  aws.String("attributeName"), // Required
		TopicARN:       aws.String("topicARN"),      // Required
		AttributeValue: aws.String("attributeValue"),
	}
	resp, err := svc.SetTopicAttributes(params)

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

func ExampleSNS_Subscribe() {
	svc := sns.New(nil)

	params := &sns.SubscribeInput{
		Protocol: aws.String("protocol"), // Required
		TopicARN: aws.String("topicARN"), // Required
		Endpoint: aws.String("endpoint"),
	}
	resp, err := svc.Subscribe(params)

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

func ExampleSNS_Unsubscribe() {
	svc := sns.New(nil)

	params := &sns.UnsubscribeInput{
		SubscriptionARN: aws.String("subscriptionARN"), // Required
	}
	resp, err := svc.Unsubscribe(params)

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