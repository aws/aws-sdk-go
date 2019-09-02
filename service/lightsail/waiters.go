package lightsail

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// WaitUntilInstanceRunning uses the Amazon Lightsail API operation
// GetInstanceState to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *Lightsail) WaitUntilInstanceRunning(input *GetInstanceStateInput) error {
	return c.WaitUntilInstanceRunningWithContext(aws.BackgroundContext(), input)
}

// WaitUntilInstanceRunningWithContext is an extended version of WaitUntilInstanceRunning.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Lightsail) WaitUntilInstanceRunningWithContext(ctx aws.Context, input *GetInstanceStateInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilInstanceRunning",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "State.Name",
				Expected: "running",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "State.Name",
				Expected: "shutting-down",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "State.Name",
				Expected: "stopped",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "State.Name",
				Expected: "stopping",
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "State.Name",
				Expected: "pending",
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "NotFoundException",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *GetInstanceStateInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetInstanceStateRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
