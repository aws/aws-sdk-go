// +build go1.9

package s3manager

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
)

type cancellerFunc = context.CancelFunc

func canceller(ctx aws.Context) (aws.Context, cancellerFunc) {
	return context.WithCancel(ctx)
}
