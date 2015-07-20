package simpledb

import "github.com/aws/aws-sdk-go/aws"

func init() {
	initService = func(s *aws.Service) {
		// SimpleDB uses custom error unmarshaling logic
		s.Handlers.UnmarshalError.Clear()
		s.Handlers.UnmarshalError.PushBack(unmarshalError)
	}
}
