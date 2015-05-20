package cloudsearch_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/cloudsearch"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCloudSearch_BuildSuggesters() {
	svc := cloudsearch.New(nil)

	params := &cloudsearch.BuildSuggestersInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.BuildSuggesters(params)

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

func ExampleCloudSearch_CreateDomain() {
	svc := cloudsearch.New(nil)

	params := &cloudsearch.CreateDomainInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.CreateDomain(params)

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

func ExampleCloudSearch_DefineAnalysisScheme() {
	svc := cloudsearch.New(nil)

	params := &cloudsearch.DefineAnalysisSchemeInput{
		AnalysisScheme: &cloudsearch.AnalysisScheme{ // Required
			AnalysisSchemeLanguage: aws.String("AnalysisSchemeLanguage"), // Required
			AnalysisSchemeName:     aws.String("StandardName"),           // Required
			AnalysisOptions: &cloudsearch.AnalysisOptions{
				AlgorithmicStemming:            aws.String("AlgorithmicStemming"),
				JapaneseTokenizationDictionary: aws.String("String"),
				StemmingDictionary:             aws.String("String"),
				Stopwords:                      aws.String("String"),
				Synonyms:                       aws.String("String"),
			},
		},
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.DefineAnalysisScheme(params)

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

func ExampleCloudSearch_DefineExpression() {
	svc := cloudsearch.New(nil)

	params := &cloudsearch.DefineExpressionInput{
		DomainName: aws.String("DomainName"), // Required
		Expression: &cloudsearch.Expression{ // Required
			ExpressionName:  aws.String("StandardName"),    // Required
			ExpressionValue: aws.String("ExpressionValue"), // Required
		},
	}
	resp, err := svc.DefineExpression(params)

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

func ExampleCloudSearch_DefineIndexField() {
	svc := cloudsearch.New(nil)

	params := &cloudsearch.DefineIndexFieldInput{
		DomainName: aws.String("DomainName"), // Required
		IndexField: &cloudsearch.IndexField{ // Required
			IndexFieldName: aws.String("DynamicFieldName"), // Required
			IndexFieldType: aws.String("IndexFieldType"),   // Required
			DateArrayOptions: &cloudsearch.DateArrayOptions{
				DefaultValue:  aws.String("FieldValue"),
				FacetEnabled:  aws.Boolean(true),
				ReturnEnabled: aws.Boolean(true),
				SearchEnabled: aws.Boolean(true),
				SourceFields:  aws.String("FieldNameCommaList"),
			},
			DateOptions: &cloudsearch.DateOptions{
				DefaultValue:  aws.String("FieldValue"),
				FacetEnabled:  aws.Boolean(true),
				ReturnEnabled: aws.Boolean(true),
				SearchEnabled: aws.Boolean(true),
				SortEnabled:   aws.Boolean(true),
				SourceField:   aws.String("FieldName"),
			},
			DoubleArrayOptions: &cloudsearch.DoubleArrayOptions{
				DefaultValue:  aws.Double(1.0),
				FacetEnabled:  aws.Boolean(true),
				ReturnEnabled: aws.Boolean(true),
				SearchEnabled: aws.Boolean(true),
				SourceFields:  aws.String("FieldNameCommaList"),
			},
			DoubleOptions: &cloudsearch.DoubleOptions{
				DefaultValue:  aws.Double(1.0),
				FacetEnabled:  aws.Boolean(true),
				ReturnEnabled: aws.Boolean(true),
				SearchEnabled: aws.Boolean(true),
				SortEnabled:   aws.Boolean(true),
				SourceField:   aws.String("FieldName"),
			},
			IntArrayOptions: &cloudsearch.IntArrayOptions{
				DefaultValue:  aws.Long(1),
				FacetEnabled:  aws.Boolean(true),
				ReturnEnabled: aws.Boolean(true),
				SearchEnabled: aws.Boolean(true),
				SourceFields:  aws.String("FieldNameCommaList"),
			},
			IntOptions: &cloudsearch.IntOptions{
				DefaultValue:  aws.Long(1),
				FacetEnabled:  aws.Boolean(true),
				ReturnEnabled: aws.Boolean(true),
				SearchEnabled: aws.Boolean(true),
				SortEnabled:   aws.Boolean(true),
				SourceField:   aws.String("FieldName"),
			},
			LatLonOptions: &cloudsearch.LatLonOptions{
				DefaultValue:  aws.String("FieldValue"),
				FacetEnabled:  aws.Boolean(true),
				ReturnEnabled: aws.Boolean(true),
				SearchEnabled: aws.Boolean(true),
				SortEnabled:   aws.Boolean(true),
				SourceField:   aws.String("FieldName"),
			},
			LiteralArrayOptions: &cloudsearch.LiteralArrayOptions{
				DefaultValue:  aws.String("FieldValue"),
				FacetEnabled:  aws.Boolean(true),
				ReturnEnabled: aws.Boolean(true),
				SearchEnabled: aws.Boolean(true),
				SourceFields:  aws.String("FieldNameCommaList"),
			},
			LiteralOptions: &cloudsearch.LiteralOptions{
				DefaultValue:  aws.String("FieldValue"),
				FacetEnabled:  aws.Boolean(true),
				ReturnEnabled: aws.Boolean(true),
				SearchEnabled: aws.Boolean(true),
				SortEnabled:   aws.Boolean(true),
				SourceField:   aws.String("FieldName"),
			},
			TextArrayOptions: &cloudsearch.TextArrayOptions{
				AnalysisScheme:   aws.String("Word"),
				DefaultValue:     aws.String("FieldValue"),
				HighlightEnabled: aws.Boolean(true),
				ReturnEnabled:    aws.Boolean(true),
				SourceFields:     aws.String("FieldNameCommaList"),
			},
			TextOptions: &cloudsearch.TextOptions{
				AnalysisScheme:   aws.String("Word"),
				DefaultValue:     aws.String("FieldValue"),
				HighlightEnabled: aws.Boolean(true),
				ReturnEnabled:    aws.Boolean(true),
				SortEnabled:      aws.Boolean(true),
				SourceField:      aws.String("FieldName"),
			},
		},
	}
	resp, err := svc.DefineIndexField(params)

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

func ExampleCloudSearch_DefineSuggester() {
	svc := cloudsearch.New(nil)

	params := &cloudsearch.DefineSuggesterInput{
		DomainName: aws.String("DomainName"), // Required
		Suggester: &cloudsearch.Suggester{ // Required
			DocumentSuggesterOptions: &cloudsearch.DocumentSuggesterOptions{ // Required
				SourceField:    aws.String("FieldName"), // Required
				FuzzyMatching:  aws.String("SuggesterFuzzyMatching"),
				SortExpression: aws.String("String"),
			},
			SuggesterName: aws.String("StandardName"), // Required
		},
	}
	resp, err := svc.DefineSuggester(params)

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

func ExampleCloudSearch_DeleteAnalysisScheme() {
	svc := cloudsearch.New(nil)

	params := &cloudsearch.DeleteAnalysisSchemeInput{
		AnalysisSchemeName: aws.String("StandardName"), // Required
		DomainName:         aws.String("DomainName"),   // Required
	}
	resp, err := svc.DeleteAnalysisScheme(params)

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

func ExampleCloudSearch_DeleteDomain() {
	svc := cloudsearch.New(nil)

	params := &cloudsearch.DeleteDomainInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.DeleteDomain(params)

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

func ExampleCloudSearch_DeleteExpression() {
	svc := cloudsearch.New(nil)

	params := &cloudsearch.DeleteExpressionInput{
		DomainName:     aws.String("DomainName"),   // Required
		ExpressionName: aws.String("StandardName"), // Required
	}
	resp, err := svc.DeleteExpression(params)

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

func ExampleCloudSearch_DeleteIndexField() {
	svc := cloudsearch.New(nil)

	params := &cloudsearch.DeleteIndexFieldInput{
		DomainName:     aws.String("DomainName"),       // Required
		IndexFieldName: aws.String("DynamicFieldName"), // Required
	}
	resp, err := svc.DeleteIndexField(params)

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

func ExampleCloudSearch_DeleteSuggester() {
	svc := cloudsearch.New(nil)

	params := &cloudsearch.DeleteSuggesterInput{
		DomainName:    aws.String("DomainName"),   // Required
		SuggesterName: aws.String("StandardName"), // Required
	}
	resp, err := svc.DeleteSuggester(params)

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

func ExampleCloudSearch_DescribeAnalysisSchemes() {
	svc := cloudsearch.New(nil)

	params := &cloudsearch.DescribeAnalysisSchemesInput{
		DomainName: aws.String("DomainName"), // Required
		AnalysisSchemeNames: []*string{
			aws.String("StandardName"), // Required
			// More values...
		},
		Deployed: aws.Boolean(true),
	}
	resp, err := svc.DescribeAnalysisSchemes(params)

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

func ExampleCloudSearch_DescribeAvailabilityOptions() {
	svc := cloudsearch.New(nil)

	params := &cloudsearch.DescribeAvailabilityOptionsInput{
		DomainName: aws.String("DomainName"), // Required
		Deployed:   aws.Boolean(true),
	}
	resp, err := svc.DescribeAvailabilityOptions(params)

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

func ExampleCloudSearch_DescribeDomains() {
	svc := cloudsearch.New(nil)

	params := &cloudsearch.DescribeDomainsInput{
		DomainNames: []*string{
			aws.String("DomainName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeDomains(params)

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

func ExampleCloudSearch_DescribeExpressions() {
	svc := cloudsearch.New(nil)

	params := &cloudsearch.DescribeExpressionsInput{
		DomainName: aws.String("DomainName"), // Required
		Deployed:   aws.Boolean(true),
		ExpressionNames: []*string{
			aws.String("StandardName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeExpressions(params)

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

func ExampleCloudSearch_DescribeIndexFields() {
	svc := cloudsearch.New(nil)

	params := &cloudsearch.DescribeIndexFieldsInput{
		DomainName: aws.String("DomainName"), // Required
		Deployed:   aws.Boolean(true),
		FieldNames: []*string{
			aws.String("DynamicFieldName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeIndexFields(params)

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

func ExampleCloudSearch_DescribeScalingParameters() {
	svc := cloudsearch.New(nil)

	params := &cloudsearch.DescribeScalingParametersInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.DescribeScalingParameters(params)

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

func ExampleCloudSearch_DescribeServiceAccessPolicies() {
	svc := cloudsearch.New(nil)

	params := &cloudsearch.DescribeServiceAccessPoliciesInput{
		DomainName: aws.String("DomainName"), // Required
		Deployed:   aws.Boolean(true),
	}
	resp, err := svc.DescribeServiceAccessPolicies(params)

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

func ExampleCloudSearch_DescribeSuggesters() {
	svc := cloudsearch.New(nil)

	params := &cloudsearch.DescribeSuggestersInput{
		DomainName: aws.String("DomainName"), // Required
		Deployed:   aws.Boolean(true),
		SuggesterNames: []*string{
			aws.String("StandardName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeSuggesters(params)

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

func ExampleCloudSearch_IndexDocuments() {
	svc := cloudsearch.New(nil)

	params := &cloudsearch.IndexDocumentsInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.IndexDocuments(params)

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

func ExampleCloudSearch_ListDomainNames() {
	svc := cloudsearch.New(nil)

	var params *cloudsearch.ListDomainNamesInput
	resp, err := svc.ListDomainNames(params)

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

func ExampleCloudSearch_UpdateAvailabilityOptions() {
	svc := cloudsearch.New(nil)

	params := &cloudsearch.UpdateAvailabilityOptionsInput{
		DomainName: aws.String("DomainName"), // Required
		MultiAZ:    aws.Boolean(true),        // Required
	}
	resp, err := svc.UpdateAvailabilityOptions(params)

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

func ExampleCloudSearch_UpdateScalingParameters() {
	svc := cloudsearch.New(nil)

	params := &cloudsearch.UpdateScalingParametersInput{
		DomainName: aws.String("DomainName"), // Required
		ScalingParameters: &cloudsearch.ScalingParameters{ // Required
			DesiredInstanceType:     aws.String("PartitionInstanceType"),
			DesiredPartitionCount:   aws.Long(1),
			DesiredReplicationCount: aws.Long(1),
		},
	}
	resp, err := svc.UpdateScalingParameters(params)

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

func ExampleCloudSearch_UpdateServiceAccessPolicies() {
	svc := cloudsearch.New(nil)

	params := &cloudsearch.UpdateServiceAccessPoliciesInput{
		AccessPolicies: aws.String("PolicyDocument"), // Required
		DomainName:     aws.String("DomainName"),     // Required
	}
	resp, err := svc.UpdateServiceAccessPolicies(params)

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