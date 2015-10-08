package simpledb

import "github.com/aws/aws-sdk-go/aws/service"

func init() {
	initService = func(s *service.Service) {
		// SimpleDB uses custom error unmarshaling logic
		s.Handlers.UnmarshalError.Clear()
		s.Handlers.UnmarshalError.PushBack(unmarshalError)
	}
}
