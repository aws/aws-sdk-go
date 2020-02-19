package servicecatalog

import (
"time"

"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/aws/request"
)

// WaitUntilProductUpdateComplete uses the AWS ServiceCatalog API operation
// DescribeRecord to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *ServiceCatalog) WaitUntilProductUpdateComplete(input *DescribeRecordInput) error {
	return c.WaitUntilProductUpdateCompleteWithContext(aws.BackgroundContext(), input)
}

// WaitUntilProductUpdateCompleteWithContext is an extended version of WaitUntilProductUpdateComplete.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ServiceCatalog) WaitUntilProductUpdateCompleteWithContext(ctx aws.Context, input *DescribeRecordInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilProductUpdateComplete",
		MaxAttempts: 120,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "RecordDetail.Status",
				Expected: "SUCCEEDED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "RecordDetail.Status",
				Expected: "FAILED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "RecordDetail.Status",
				Expected: "IN_PROGRESS_IN_ERROR",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeRecordInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeRecordRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}


// WaitUntilProductCreateComplete uses the AWS ServiceCatalog API operation
// DescribeRecord to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *ServiceCatalog) WaitUntilProductCreateComplete(input *DescribeRecordInput) error {
	return c.WaitUntilProductCreateCompleteWithContext(aws.BackgroundContext(), input)
}

// WaitUntilProductCreateCompleteWithContext is an extended version of WaitUntilProductCreateComplete.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ServiceCatalog) WaitUntilProductCreateCompleteWithContext(ctx aws.Context, input *DescribeRecordInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilProductCreateComplete",
		MaxAttempts: 120,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "RecordDetail.Status",
				Expected: "SUCCEEDED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "RecordDetail.Status",
				Expected: "FAILED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "RecordDetail.Status",
				Expected: "IN_PROGRESS_IN_ERROR",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeRecordInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeRecordRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
