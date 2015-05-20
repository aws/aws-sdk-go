package cloudfront_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/cloudfront"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCloudFront_CreateCloudFrontOriginAccessIdentity() {
	svc := cloudfront.New(nil)

	params := &cloudfront.CreateCloudFrontOriginAccessIdentityInput{
		CloudFrontOriginAccessIdentityConfig: &cloudfront.CloudFrontOriginAccessIdentityConfig{ // Required
			CallerReference: aws.String("string"), // Required
			Comment:         aws.String("string"), // Required
		},
	}
	resp, err := svc.CreateCloudFrontOriginAccessIdentity(params)

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

func ExampleCloudFront_CreateDistribution() {
	svc := cloudfront.New(nil)

	params := &cloudfront.CreateDistributionInput{
		DistributionConfig: &cloudfront.DistributionConfig{ // Required
			CallerReference: aws.String("string"), // Required
			Comment:         aws.String("string"), // Required
			DefaultCacheBehavior: &cloudfront.DefaultCacheBehavior{ // Required
				ForwardedValues: &cloudfront.ForwardedValues{ // Required
					Cookies: &cloudfront.CookiePreference{ // Required
						Forward: aws.String("ItemSelection"), // Required
						WhitelistedNames: &cloudfront.CookieNames{
							Quantity: aws.Long(1), // Required
							Items: []*string{
								aws.String("string"), // Required
								// More values...
							},
						},
					},
					QueryString: aws.Boolean(true), // Required
					Headers: &cloudfront.Headers{
						Quantity: aws.Long(1), // Required
						Items: []*string{
							aws.String("string"), // Required
							// More values...
						},
					},
				},
				MinTTL:         aws.Long(1),          // Required
				TargetOriginID: aws.String("string"), // Required
				TrustedSigners: &cloudfront.TrustedSigners{ // Required
					Enabled:  aws.Boolean(true), // Required
					Quantity: aws.Long(1),       // Required
					Items: []*string{
						aws.String("string"), // Required
						// More values...
					},
				},
				ViewerProtocolPolicy: aws.String("ViewerProtocolPolicy"), // Required
				AllowedMethods: &cloudfront.AllowedMethods{
					Items: []*string{ // Required
						aws.String("Method"), // Required
						// More values...
					},
					Quantity: aws.Long(1), // Required
					CachedMethods: &cloudfront.CachedMethods{
						Items: []*string{ // Required
							aws.String("Method"), // Required
							// More values...
						},
						Quantity: aws.Long(1), // Required
					},
				},
				SmoothStreaming: aws.Boolean(true),
			},
			Enabled: aws.Boolean(true), // Required
			Origins: &cloudfront.Origins{ // Required
				Quantity: aws.Long(1), // Required
				Items: []*cloudfront.Origin{
					&cloudfront.Origin{ // Required
						DomainName: aws.String("string"), // Required
						ID:         aws.String("string"), // Required
						CustomOriginConfig: &cloudfront.CustomOriginConfig{
							HTTPPort:             aws.Long(1),                        // Required
							HTTPSPort:            aws.Long(1),                        // Required
							OriginProtocolPolicy: aws.String("OriginProtocolPolicy"), // Required
						},
						OriginPath: aws.String("string"),
						S3OriginConfig: &cloudfront.S3OriginConfig{
							OriginAccessIdentity: aws.String("string"), // Required
						},
					},
					// More values...
				},
			},
			Aliases: &cloudfront.Aliases{
				Quantity: aws.Long(1), // Required
				Items: []*string{
					aws.String("string"), // Required
					// More values...
				},
			},
			CacheBehaviors: &cloudfront.CacheBehaviors{
				Quantity: aws.Long(1), // Required
				Items: []*cloudfront.CacheBehavior{
					&cloudfront.CacheBehavior{ // Required
						ForwardedValues: &cloudfront.ForwardedValues{ // Required
							Cookies: &cloudfront.CookiePreference{ // Required
								Forward: aws.String("ItemSelection"), // Required
								WhitelistedNames: &cloudfront.CookieNames{
									Quantity: aws.Long(1), // Required
									Items: []*string{
										aws.String("string"), // Required
										// More values...
									},
								},
							},
							QueryString: aws.Boolean(true), // Required
							Headers: &cloudfront.Headers{
								Quantity: aws.Long(1), // Required
								Items: []*string{
									aws.String("string"), // Required
									// More values...
								},
							},
						},
						MinTTL:         aws.Long(1),          // Required
						PathPattern:    aws.String("string"), // Required
						TargetOriginID: aws.String("string"), // Required
						TrustedSigners: &cloudfront.TrustedSigners{ // Required
							Enabled:  aws.Boolean(true), // Required
							Quantity: aws.Long(1),       // Required
							Items: []*string{
								aws.String("string"), // Required
								// More values...
							},
						},
						ViewerProtocolPolicy: aws.String("ViewerProtocolPolicy"), // Required
						AllowedMethods: &cloudfront.AllowedMethods{
							Items: []*string{ // Required
								aws.String("Method"), // Required
								// More values...
							},
							Quantity: aws.Long(1), // Required
							CachedMethods: &cloudfront.CachedMethods{
								Items: []*string{ // Required
									aws.String("Method"), // Required
									// More values...
								},
								Quantity: aws.Long(1), // Required
							},
						},
						SmoothStreaming: aws.Boolean(true),
					},
					// More values...
				},
			},
			CustomErrorResponses: &cloudfront.CustomErrorResponses{
				Quantity: aws.Long(1), // Required
				Items: []*cloudfront.CustomErrorResponse{
					&cloudfront.CustomErrorResponse{ // Required
						ErrorCode:          aws.Long(1), // Required
						ErrorCachingMinTTL: aws.Long(1),
						ResponseCode:       aws.String("string"),
						ResponsePagePath:   aws.String("string"),
					},
					// More values...
				},
			},
			DefaultRootObject: aws.String("string"),
			Logging: &cloudfront.LoggingConfig{
				Bucket:         aws.String("string"), // Required
				Enabled:        aws.Boolean(true),    // Required
				IncludeCookies: aws.Boolean(true),    // Required
				Prefix:         aws.String("string"), // Required
			},
			PriceClass: aws.String("PriceClass"),
			Restrictions: &cloudfront.Restrictions{
				GeoRestriction: &cloudfront.GeoRestriction{ // Required
					Quantity:        aws.Long(1),                      // Required
					RestrictionType: aws.String("GeoRestrictionType"), // Required
					Items: []*string{
						aws.String("string"), // Required
						// More values...
					},
				},
			},
			ViewerCertificate: &cloudfront.ViewerCertificate{
				CloudFrontDefaultCertificate: aws.Boolean(true),
				IAMCertificateID:             aws.String("string"),
				MinimumProtocolVersion:       aws.String("MinimumProtocolVersion"),
				SSLSupportMethod:             aws.String("SSLSupportMethod"),
			},
		},
	}
	resp, err := svc.CreateDistribution(params)

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

func ExampleCloudFront_CreateInvalidation() {
	svc := cloudfront.New(nil)

	params := &cloudfront.CreateInvalidationInput{
		DistributionID: aws.String("string"), // Required
		InvalidationBatch: &cloudfront.InvalidationBatch{ // Required
			CallerReference: aws.String("string"), // Required
			Paths: &cloudfront.Paths{ // Required
				Quantity: aws.Long(1), // Required
				Items: []*string{
					aws.String("string"), // Required
					// More values...
				},
			},
		},
	}
	resp, err := svc.CreateInvalidation(params)

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

func ExampleCloudFront_CreateStreamingDistribution() {
	svc := cloudfront.New(nil)

	params := &cloudfront.CreateStreamingDistributionInput{
		StreamingDistributionConfig: &cloudfront.StreamingDistributionConfig{ // Required
			CallerReference: aws.String("string"), // Required
			Comment:         aws.String("string"), // Required
			Enabled:         aws.Boolean(true),    // Required
			S3Origin: &cloudfront.S3Origin{ // Required
				DomainName:           aws.String("string"), // Required
				OriginAccessIdentity: aws.String("string"), // Required
			},
			TrustedSigners: &cloudfront.TrustedSigners{ // Required
				Enabled:  aws.Boolean(true), // Required
				Quantity: aws.Long(1),       // Required
				Items: []*string{
					aws.String("string"), // Required
					// More values...
				},
			},
			Aliases: &cloudfront.Aliases{
				Quantity: aws.Long(1), // Required
				Items: []*string{
					aws.String("string"), // Required
					// More values...
				},
			},
			Logging: &cloudfront.StreamingLoggingConfig{
				Bucket:  aws.String("string"), // Required
				Enabled: aws.Boolean(true),    // Required
				Prefix:  aws.String("string"), // Required
			},
			PriceClass: aws.String("PriceClass"),
		},
	}
	resp, err := svc.CreateStreamingDistribution(params)

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

func ExampleCloudFront_DeleteCloudFrontOriginAccessIdentity() {
	svc := cloudfront.New(nil)

	params := &cloudfront.DeleteCloudFrontOriginAccessIdentityInput{
		ID:      aws.String("string"), // Required
		IfMatch: aws.String("string"),
	}
	resp, err := svc.DeleteCloudFrontOriginAccessIdentity(params)

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

func ExampleCloudFront_DeleteDistribution() {
	svc := cloudfront.New(nil)

	params := &cloudfront.DeleteDistributionInput{
		ID:      aws.String("string"), // Required
		IfMatch: aws.String("string"),
	}
	resp, err := svc.DeleteDistribution(params)

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

func ExampleCloudFront_DeleteStreamingDistribution() {
	svc := cloudfront.New(nil)

	params := &cloudfront.DeleteStreamingDistributionInput{
		ID:      aws.String("string"), // Required
		IfMatch: aws.String("string"),
	}
	resp, err := svc.DeleteStreamingDistribution(params)

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

func ExampleCloudFront_GetCloudFrontOriginAccessIdentity() {
	svc := cloudfront.New(nil)

	params := &cloudfront.GetCloudFrontOriginAccessIdentityInput{
		ID: aws.String("string"), // Required
	}
	resp, err := svc.GetCloudFrontOriginAccessIdentity(params)

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

func ExampleCloudFront_GetCloudFrontOriginAccessIdentityConfig() {
	svc := cloudfront.New(nil)

	params := &cloudfront.GetCloudFrontOriginAccessIdentityConfigInput{
		ID: aws.String("string"), // Required
	}
	resp, err := svc.GetCloudFrontOriginAccessIdentityConfig(params)

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

func ExampleCloudFront_GetDistribution() {
	svc := cloudfront.New(nil)

	params := &cloudfront.GetDistributionInput{
		ID: aws.String("string"), // Required
	}
	resp, err := svc.GetDistribution(params)

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

func ExampleCloudFront_GetDistributionConfig() {
	svc := cloudfront.New(nil)

	params := &cloudfront.GetDistributionConfigInput{
		ID: aws.String("string"), // Required
	}
	resp, err := svc.GetDistributionConfig(params)

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

func ExampleCloudFront_GetInvalidation() {
	svc := cloudfront.New(nil)

	params := &cloudfront.GetInvalidationInput{
		DistributionID: aws.String("string"), // Required
		ID:             aws.String("string"), // Required
	}
	resp, err := svc.GetInvalidation(params)

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

func ExampleCloudFront_GetStreamingDistribution() {
	svc := cloudfront.New(nil)

	params := &cloudfront.GetStreamingDistributionInput{
		ID: aws.String("string"), // Required
	}
	resp, err := svc.GetStreamingDistribution(params)

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

func ExampleCloudFront_GetStreamingDistributionConfig() {
	svc := cloudfront.New(nil)

	params := &cloudfront.GetStreamingDistributionConfigInput{
		ID: aws.String("string"), // Required
	}
	resp, err := svc.GetStreamingDistributionConfig(params)

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

func ExampleCloudFront_ListCloudFrontOriginAccessIdentities() {
	svc := cloudfront.New(nil)

	params := &cloudfront.ListCloudFrontOriginAccessIdentitiesInput{
		Marker:   aws.String("string"),
		MaxItems: aws.Long(1),
	}
	resp, err := svc.ListCloudFrontOriginAccessIdentities(params)

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

func ExampleCloudFront_ListDistributions() {
	svc := cloudfront.New(nil)

	params := &cloudfront.ListDistributionsInput{
		Marker:   aws.String("string"),
		MaxItems: aws.Long(1),
	}
	resp, err := svc.ListDistributions(params)

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

func ExampleCloudFront_ListInvalidations() {
	svc := cloudfront.New(nil)

	params := &cloudfront.ListInvalidationsInput{
		DistributionID: aws.String("string"), // Required
		Marker:         aws.String("string"),
		MaxItems:       aws.Long(1),
	}
	resp, err := svc.ListInvalidations(params)

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

func ExampleCloudFront_ListStreamingDistributions() {
	svc := cloudfront.New(nil)

	params := &cloudfront.ListStreamingDistributionsInput{
		Marker:   aws.String("string"),
		MaxItems: aws.Long(1),
	}
	resp, err := svc.ListStreamingDistributions(params)

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

func ExampleCloudFront_UpdateCloudFrontOriginAccessIdentity() {
	svc := cloudfront.New(nil)

	params := &cloudfront.UpdateCloudFrontOriginAccessIdentityInput{
		CloudFrontOriginAccessIdentityConfig: &cloudfront.CloudFrontOriginAccessIdentityConfig{ // Required
			CallerReference: aws.String("string"), // Required
			Comment:         aws.String("string"), // Required
		},
		ID:      aws.String("string"), // Required
		IfMatch: aws.String("string"),
	}
	resp, err := svc.UpdateCloudFrontOriginAccessIdentity(params)

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

func ExampleCloudFront_UpdateDistribution() {
	svc := cloudfront.New(nil)

	params := &cloudfront.UpdateDistributionInput{
		DistributionConfig: &cloudfront.DistributionConfig{ // Required
			CallerReference: aws.String("string"), // Required
			Comment:         aws.String("string"), // Required
			DefaultCacheBehavior: &cloudfront.DefaultCacheBehavior{ // Required
				ForwardedValues: &cloudfront.ForwardedValues{ // Required
					Cookies: &cloudfront.CookiePreference{ // Required
						Forward: aws.String("ItemSelection"), // Required
						WhitelistedNames: &cloudfront.CookieNames{
							Quantity: aws.Long(1), // Required
							Items: []*string{
								aws.String("string"), // Required
								// More values...
							},
						},
					},
					QueryString: aws.Boolean(true), // Required
					Headers: &cloudfront.Headers{
						Quantity: aws.Long(1), // Required
						Items: []*string{
							aws.String("string"), // Required
							// More values...
						},
					},
				},
				MinTTL:         aws.Long(1),          // Required
				TargetOriginID: aws.String("string"), // Required
				TrustedSigners: &cloudfront.TrustedSigners{ // Required
					Enabled:  aws.Boolean(true), // Required
					Quantity: aws.Long(1),       // Required
					Items: []*string{
						aws.String("string"), // Required
						// More values...
					},
				},
				ViewerProtocolPolicy: aws.String("ViewerProtocolPolicy"), // Required
				AllowedMethods: &cloudfront.AllowedMethods{
					Items: []*string{ // Required
						aws.String("Method"), // Required
						// More values...
					},
					Quantity: aws.Long(1), // Required
					CachedMethods: &cloudfront.CachedMethods{
						Items: []*string{ // Required
							aws.String("Method"), // Required
							// More values...
						},
						Quantity: aws.Long(1), // Required
					},
				},
				SmoothStreaming: aws.Boolean(true),
			},
			Enabled: aws.Boolean(true), // Required
			Origins: &cloudfront.Origins{ // Required
				Quantity: aws.Long(1), // Required
				Items: []*cloudfront.Origin{
					&cloudfront.Origin{ // Required
						DomainName: aws.String("string"), // Required
						ID:         aws.String("string"), // Required
						CustomOriginConfig: &cloudfront.CustomOriginConfig{
							HTTPPort:             aws.Long(1),                        // Required
							HTTPSPort:            aws.Long(1),                        // Required
							OriginProtocolPolicy: aws.String("OriginProtocolPolicy"), // Required
						},
						OriginPath: aws.String("string"),
						S3OriginConfig: &cloudfront.S3OriginConfig{
							OriginAccessIdentity: aws.String("string"), // Required
						},
					},
					// More values...
				},
			},
			Aliases: &cloudfront.Aliases{
				Quantity: aws.Long(1), // Required
				Items: []*string{
					aws.String("string"), // Required
					// More values...
				},
			},
			CacheBehaviors: &cloudfront.CacheBehaviors{
				Quantity: aws.Long(1), // Required
				Items: []*cloudfront.CacheBehavior{
					&cloudfront.CacheBehavior{ // Required
						ForwardedValues: &cloudfront.ForwardedValues{ // Required
							Cookies: &cloudfront.CookiePreference{ // Required
								Forward: aws.String("ItemSelection"), // Required
								WhitelistedNames: &cloudfront.CookieNames{
									Quantity: aws.Long(1), // Required
									Items: []*string{
										aws.String("string"), // Required
										// More values...
									},
								},
							},
							QueryString: aws.Boolean(true), // Required
							Headers: &cloudfront.Headers{
								Quantity: aws.Long(1), // Required
								Items: []*string{
									aws.String("string"), // Required
									// More values...
								},
							},
						},
						MinTTL:         aws.Long(1),          // Required
						PathPattern:    aws.String("string"), // Required
						TargetOriginID: aws.String("string"), // Required
						TrustedSigners: &cloudfront.TrustedSigners{ // Required
							Enabled:  aws.Boolean(true), // Required
							Quantity: aws.Long(1),       // Required
							Items: []*string{
								aws.String("string"), // Required
								// More values...
							},
						},
						ViewerProtocolPolicy: aws.String("ViewerProtocolPolicy"), // Required
						AllowedMethods: &cloudfront.AllowedMethods{
							Items: []*string{ // Required
								aws.String("Method"), // Required
								// More values...
							},
							Quantity: aws.Long(1), // Required
							CachedMethods: &cloudfront.CachedMethods{
								Items: []*string{ // Required
									aws.String("Method"), // Required
									// More values...
								},
								Quantity: aws.Long(1), // Required
							},
						},
						SmoothStreaming: aws.Boolean(true),
					},
					// More values...
				},
			},
			CustomErrorResponses: &cloudfront.CustomErrorResponses{
				Quantity: aws.Long(1), // Required
				Items: []*cloudfront.CustomErrorResponse{
					&cloudfront.CustomErrorResponse{ // Required
						ErrorCode:          aws.Long(1), // Required
						ErrorCachingMinTTL: aws.Long(1),
						ResponseCode:       aws.String("string"),
						ResponsePagePath:   aws.String("string"),
					},
					// More values...
				},
			},
			DefaultRootObject: aws.String("string"),
			Logging: &cloudfront.LoggingConfig{
				Bucket:         aws.String("string"), // Required
				Enabled:        aws.Boolean(true),    // Required
				IncludeCookies: aws.Boolean(true),    // Required
				Prefix:         aws.String("string"), // Required
			},
			PriceClass: aws.String("PriceClass"),
			Restrictions: &cloudfront.Restrictions{
				GeoRestriction: &cloudfront.GeoRestriction{ // Required
					Quantity:        aws.Long(1),                      // Required
					RestrictionType: aws.String("GeoRestrictionType"), // Required
					Items: []*string{
						aws.String("string"), // Required
						// More values...
					},
				},
			},
			ViewerCertificate: &cloudfront.ViewerCertificate{
				CloudFrontDefaultCertificate: aws.Boolean(true),
				IAMCertificateID:             aws.String("string"),
				MinimumProtocolVersion:       aws.String("MinimumProtocolVersion"),
				SSLSupportMethod:             aws.String("SSLSupportMethod"),
			},
		},
		ID:      aws.String("string"), // Required
		IfMatch: aws.String("string"),
	}
	resp, err := svc.UpdateDistribution(params)

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

func ExampleCloudFront_UpdateStreamingDistribution() {
	svc := cloudfront.New(nil)

	params := &cloudfront.UpdateStreamingDistributionInput{
		ID: aws.String("string"), // Required
		StreamingDistributionConfig: &cloudfront.StreamingDistributionConfig{ // Required
			CallerReference: aws.String("string"), // Required
			Comment:         aws.String("string"), // Required
			Enabled:         aws.Boolean(true),    // Required
			S3Origin: &cloudfront.S3Origin{ // Required
				DomainName:           aws.String("string"), // Required
				OriginAccessIdentity: aws.String("string"), // Required
			},
			TrustedSigners: &cloudfront.TrustedSigners{ // Required
				Enabled:  aws.Boolean(true), // Required
				Quantity: aws.Long(1),       // Required
				Items: []*string{
					aws.String("string"), // Required
					// More values...
				},
			},
			Aliases: &cloudfront.Aliases{
				Quantity: aws.Long(1), // Required
				Items: []*string{
					aws.String("string"), // Required
					// More values...
				},
			},
			Logging: &cloudfront.StreamingLoggingConfig{
				Bucket:  aws.String("string"), // Required
				Enabled: aws.Boolean(true),    // Required
				Prefix:  aws.String("string"), // Required
			},
			PriceClass: aws.String("PriceClass"),
		},
		IfMatch: aws.String("string"),
	}
	resp, err := svc.UpdateStreamingDistribution(params)

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