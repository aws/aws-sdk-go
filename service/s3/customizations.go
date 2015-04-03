package s3

import "github.com/awslabs/aws-sdk-go/aws"

func init() {
	initService = func(s *aws.Service) {
		// S3 uses custom error unmarshaling logic
		s.Handlers.UnmarshalError.Clear()
		s.Handlers.UnmarshalError.PushBack(unmarshalError)
	}

	initRequest = func(r *aws.Request) {
		switch r.Operation {
		case opPutBucketCORS, opPutBucketLifecycle, opPutBucketTagging, opDeleteObjects:
			// These S3 operations require Content-MD5 to be set
			r.Handlers.Build.PushBack(contentMD5)
		}
	}
}
