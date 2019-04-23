package rds

import (
	"fmt"
	"io"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
)

const (
	opDownloadCompleteDBLogFile = "DownloadCompleteDBLogFile"
)

// DownloadCompleteDBLogFileRequest generates a "aws/request.Request" representing the
// client's request for the DownloadCompleteDBLogFile operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Cannot use "Send" method on the returned Request to send the API call to the service.
// Use http.DefaultClient.Do(req.HTTPRequest) instead of "Send" method.
//
// See DownloadCompleteDBLogFile for more information on using the DownloadCompleteDBLogFile
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DownloadCompleteDBLogFileRequest method.
//    req, out := client.DownloadCompleteDBLogFileRequest(params)
//
//    resp, err := http.DefaultClient.Do(req.HTTPRequest)
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//    out.SetBody(resp.Body)
//
// See also, https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#DownloadCompleteDBLogFile
func (c *RDS) DownloadCompleteDBLogFileRequest(input *DownloadCompleteDBLogFileInput) (req *request.Request, output *DownloadCompleteDBLogFileOutput) {
	op := &request.Operation{
		Name:       opDownloadCompleteDBLogFile,
		HTTPMethod: "GET",
		HTTPPath: fmt.Sprintf(
			"/v13/downloadCompleteLogFile/%s/%s",
			*input.DBInstanceIdentifier,
			*input.LogFileName,
		),
	}

	if input == nil {
		input = &DownloadCompleteDBLogFileInput{}
	}

	output = &DownloadCompleteDBLogFileOutput{}
	req = c.newRequest(op, input, output)
	return
}

// DownloadCompleteDBLogFile API operation for Amazon Relational Database Service.
//
// Downloads the contents of the specified database log file.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Relational Database Service's
// API operation DownloadCompleteDBLogFile for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeDBInstanceNotFoundFault "DBInstanceNotFound"
//   DBInstanceIdentifier doesn't refer to an existing DB instance.
//
//   * ErrCodeDBLogFileNotFoundFault "DBLogFileNotFoundFault"
//   LogFileName doesn't refer to an existing DB log file.
//
// See also, https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#DownloadCompleteDBLogFile
func (c *RDS) DownloadCompleteDBLogFile(input *DownloadCompleteDBLogFileInput) (*DownloadCompleteDBLogFileOutput, error) {
	req, out := c.DownloadCompleteDBLogFileRequest(input)
	v4.SignSDKRequest(req)
	resp, err := http.DefaultClient.Do(req.HTTPRequest)
	if err != nil {
		return nil, err
	}
	out.SetBody(resp.Body)
	return out, nil
}

// DownloadCompleteDBLogFileWithContext is the same as DownloadCompleteDBLogFile with the addition of
// the ability to pass a context and additional request options.
//
// See DownloadCompleteDBLogFile for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDS) DownloadCompleteDBLogFileWithContext(ctx aws.Context, input *DownloadCompleteDBLogFileInput, opts ...request.Option) (*DownloadCompleteDBLogFileOutput, error) {
	req, out := c.DownloadCompleteDBLogFileRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	v4.SignSDKRequest(req)
	resp, err := http.DefaultClient.Do(req.HTTPRequest)
	if err != nil {
		return nil, err
	}
	out.SetBody(resp.Body)
	return out, nil
}

type DownloadCompleteDBLogFileInput struct {
	_ struct{} `type:"structure"`

	// The customer-assigned name of the DB instance that contains the log files
	// you want to list.
	//
	// Constraints:
	//
	//    * Must match the identifier of an existing DBInstance.
	//
	// DBInstanceIdentifier is a required field
	DBInstanceIdentifier *string `type:"string" required:"true"`

	// The name of the log file to be downloaded.
	//
	// LogFileName is a required field
	LogFileName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DownloadCompleteDBLogFileInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DownloadCompleteDBLogFileInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DownloadCompleteDBLogFileInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DownloadCompleteDBLogFileInput"}
	if s.DBInstanceIdentifier == nil {
		invalidParams.Add(request.NewErrParamRequired("DBInstanceIdentifier"))
	}
	if s.LogFileName == nil {
		invalidParams.Add(request.NewErrParamRequired("LogFileName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDBInstanceIdentifier sets the DBInstanceIdentifier field's value.
func (s *DownloadCompleteDBLogFileInput) SetDBInstanceIdentifier(v string) *DownloadCompleteDBLogFileInput {
	s.DBInstanceIdentifier = &v
	return s
}

// SetLogFileName sets the LogFileName field's value.
func (s *DownloadCompleteDBLogFileInput) SetLogFileName(v string) *DownloadCompleteDBLogFileInput {
	s.LogFileName = &v
	return s
}

// This data type is used as a response element to DownloadCompleteDBLogFile.
type DownloadCompleteDBLogFileOutput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// Object data.
	Body io.ReadCloser `type:"blob"`
}

// String returns the string representation
func (s DownloadCompleteDBLogFileOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DownloadCompleteDBLogFileOutput) GoString() string {
	return s.String()
}

// SetBody sets the Body field's value.
func (s *DownloadCompleteDBLogFileOutput) SetBody(v io.ReadCloser) *DownloadCompleteDBLogFileOutput {
	s.Body = v
	return s
}
