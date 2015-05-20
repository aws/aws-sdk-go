package kms_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/kms"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleKMS_CreateAlias() {
	svc := kms.New(nil)

	params := &kms.CreateAliasInput{
		AliasName:   aws.String("AliasNameType"), // Required
		TargetKeyID: aws.String("KeyIdType"),     // Required
	}
	resp, err := svc.CreateAlias(params)

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

func ExampleKMS_CreateGrant() {
	svc := kms.New(nil)

	params := &kms.CreateGrantInput{
		GranteePrincipal: aws.String("PrincipalIdType"), // Required
		KeyID:            aws.String("KeyIdType"),       // Required
		Constraints: &kms.GrantConstraints{
			EncryptionContextEquals: &map[string]*string{
				"Key": aws.String("EncryptionContextValue"), // Required
				// More values...
			},
			EncryptionContextSubset: &map[string]*string{
				"Key": aws.String("EncryptionContextValue"), // Required
				// More values...
			},
		},
		GrantTokens: []*string{
			aws.String("GrantTokenType"), // Required
			// More values...
		},
		Operations: []*string{
			aws.String("GrantOperation"), // Required
			// More values...
		},
		RetiringPrincipal: aws.String("PrincipalIdType"),
	}
	resp, err := svc.CreateGrant(params)

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

func ExampleKMS_CreateKey() {
	svc := kms.New(nil)

	params := &kms.CreateKeyInput{
		Description: aws.String("DescriptionType"),
		KeyUsage:    aws.String("KeyUsageType"),
		Policy:      aws.String("PolicyType"),
	}
	resp, err := svc.CreateKey(params)

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

func ExampleKMS_Decrypt() {
	svc := kms.New(nil)

	params := &kms.DecryptInput{
		CiphertextBlob: []byte("PAYLOAD"), // Required
		EncryptionContext: &map[string]*string{
			"Key": aws.String("EncryptionContextValue"), // Required
			// More values...
		},
		GrantTokens: []*string{
			aws.String("GrantTokenType"), // Required
			// More values...
		},
	}
	resp, err := svc.Decrypt(params)

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

func ExampleKMS_DeleteAlias() {
	svc := kms.New(nil)

	params := &kms.DeleteAliasInput{
		AliasName: aws.String("AliasNameType"), // Required
	}
	resp, err := svc.DeleteAlias(params)

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

func ExampleKMS_DescribeKey() {
	svc := kms.New(nil)

	params := &kms.DescribeKeyInput{
		KeyID: aws.String("KeyIdType"), // Required
	}
	resp, err := svc.DescribeKey(params)

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

func ExampleKMS_DisableKey() {
	svc := kms.New(nil)

	params := &kms.DisableKeyInput{
		KeyID: aws.String("KeyIdType"), // Required
	}
	resp, err := svc.DisableKey(params)

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

func ExampleKMS_DisableKeyRotation() {
	svc := kms.New(nil)

	params := &kms.DisableKeyRotationInput{
		KeyID: aws.String("KeyIdType"), // Required
	}
	resp, err := svc.DisableKeyRotation(params)

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

func ExampleKMS_EnableKey() {
	svc := kms.New(nil)

	params := &kms.EnableKeyInput{
		KeyID: aws.String("KeyIdType"), // Required
	}
	resp, err := svc.EnableKey(params)

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

func ExampleKMS_EnableKeyRotation() {
	svc := kms.New(nil)

	params := &kms.EnableKeyRotationInput{
		KeyID: aws.String("KeyIdType"), // Required
	}
	resp, err := svc.EnableKeyRotation(params)

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

func ExampleKMS_Encrypt() {
	svc := kms.New(nil)

	params := &kms.EncryptInput{
		KeyID:     aws.String("KeyIdType"), // Required
		Plaintext: []byte("PAYLOAD"),       // Required
		EncryptionContext: &map[string]*string{
			"Key": aws.String("EncryptionContextValue"), // Required
			// More values...
		},
		GrantTokens: []*string{
			aws.String("GrantTokenType"), // Required
			// More values...
		},
	}
	resp, err := svc.Encrypt(params)

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

func ExampleKMS_GenerateDataKey() {
	svc := kms.New(nil)

	params := &kms.GenerateDataKeyInput{
		KeyID: aws.String("KeyIdType"), // Required
		EncryptionContext: &map[string]*string{
			"Key": aws.String("EncryptionContextValue"), // Required
			// More values...
		},
		GrantTokens: []*string{
			aws.String("GrantTokenType"), // Required
			// More values...
		},
		KeySpec:       aws.String("DataKeySpec"),
		NumberOfBytes: aws.Long(1),
	}
	resp, err := svc.GenerateDataKey(params)

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

func ExampleKMS_GenerateDataKeyWithoutPlaintext() {
	svc := kms.New(nil)

	params := &kms.GenerateDataKeyWithoutPlaintextInput{
		KeyID: aws.String("KeyIdType"), // Required
		EncryptionContext: &map[string]*string{
			"Key": aws.String("EncryptionContextValue"), // Required
			// More values...
		},
		GrantTokens: []*string{
			aws.String("GrantTokenType"), // Required
			// More values...
		},
		KeySpec:       aws.String("DataKeySpec"),
		NumberOfBytes: aws.Long(1),
	}
	resp, err := svc.GenerateDataKeyWithoutPlaintext(params)

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

func ExampleKMS_GenerateRandom() {
	svc := kms.New(nil)

	params := &kms.GenerateRandomInput{
		NumberOfBytes: aws.Long(1),
	}
	resp, err := svc.GenerateRandom(params)

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

func ExampleKMS_GetKeyPolicy() {
	svc := kms.New(nil)

	params := &kms.GetKeyPolicyInput{
		KeyID:      aws.String("KeyIdType"),      // Required
		PolicyName: aws.String("PolicyNameType"), // Required
	}
	resp, err := svc.GetKeyPolicy(params)

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

func ExampleKMS_GetKeyRotationStatus() {
	svc := kms.New(nil)

	params := &kms.GetKeyRotationStatusInput{
		KeyID: aws.String("KeyIdType"), // Required
	}
	resp, err := svc.GetKeyRotationStatus(params)

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

func ExampleKMS_ListAliases() {
	svc := kms.New(nil)

	params := &kms.ListAliasesInput{
		Limit:  aws.Long(1),
		Marker: aws.String("MarkerType"),
	}
	resp, err := svc.ListAliases(params)

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

func ExampleKMS_ListGrants() {
	svc := kms.New(nil)

	params := &kms.ListGrantsInput{
		KeyID:  aws.String("KeyIdType"), // Required
		Limit:  aws.Long(1),
		Marker: aws.String("MarkerType"),
	}
	resp, err := svc.ListGrants(params)

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

func ExampleKMS_ListKeyPolicies() {
	svc := kms.New(nil)

	params := &kms.ListKeyPoliciesInput{
		KeyID:  aws.String("KeyIdType"), // Required
		Limit:  aws.Long(1),
		Marker: aws.String("MarkerType"),
	}
	resp, err := svc.ListKeyPolicies(params)

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

func ExampleKMS_ListKeys() {
	svc := kms.New(nil)

	params := &kms.ListKeysInput{
		Limit:  aws.Long(1),
		Marker: aws.String("MarkerType"),
	}
	resp, err := svc.ListKeys(params)

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

func ExampleKMS_PutKeyPolicy() {
	svc := kms.New(nil)

	params := &kms.PutKeyPolicyInput{
		KeyID:      aws.String("KeyIdType"),      // Required
		Policy:     aws.String("PolicyType"),     // Required
		PolicyName: aws.String("PolicyNameType"), // Required
	}
	resp, err := svc.PutKeyPolicy(params)

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

func ExampleKMS_ReEncrypt() {
	svc := kms.New(nil)

	params := &kms.ReEncryptInput{
		CiphertextBlob:   []byte("PAYLOAD"),       // Required
		DestinationKeyID: aws.String("KeyIdType"), // Required
		DestinationEncryptionContext: &map[string]*string{
			"Key": aws.String("EncryptionContextValue"), // Required
			// More values...
		},
		GrantTokens: []*string{
			aws.String("GrantTokenType"), // Required
			// More values...
		},
		SourceEncryptionContext: &map[string]*string{
			"Key": aws.String("EncryptionContextValue"), // Required
			// More values...
		},
	}
	resp, err := svc.ReEncrypt(params)

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

func ExampleKMS_RetireGrant() {
	svc := kms.New(nil)

	params := &kms.RetireGrantInput{
		GrantToken: aws.String("GrantTokenType"), // Required
	}
	resp, err := svc.RetireGrant(params)

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

func ExampleKMS_RevokeGrant() {
	svc := kms.New(nil)

	params := &kms.RevokeGrantInput{
		GrantID: aws.String("GrantIdType"), // Required
		KeyID:   aws.String("KeyIdType"),   // Required
	}
	resp, err := svc.RevokeGrant(params)

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

func ExampleKMS_UpdateKeyDescription() {
	svc := kms.New(nil)

	params := &kms.UpdateKeyDescriptionInput{
		Description: aws.String("DescriptionType"), // Required
		KeyID:       aws.String("KeyIdType"),       // Required
	}
	resp, err := svc.UpdateKeyDescription(params)

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