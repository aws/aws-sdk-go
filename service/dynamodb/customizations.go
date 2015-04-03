package dynamodb

import (
	"math"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
)

func init() {
	initService = func(s *aws.Service) {
		s.DefaultMaxRetries = 10
		s.RetryRules = func(r *aws.Request) time.Duration {
			delay := time.Duration(math.Pow(2, float64(r.RetryCount))) * 50
			return delay * time.Millisecond
		}
	}
}
