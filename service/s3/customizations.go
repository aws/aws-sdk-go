package s3

import "github.com/awslabs/aws-sdk-go/aws"

func init() {
	initService = func(s *aws.Service) {
		// S3 uses custom error unmarshaling logic
		s.Handlers.UnmarshalError.Init()
		s.Handlers.UnmarshalError.PushBack(unmarshalError)
	}
}
