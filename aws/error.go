package aws

// An APIError is an error returned by an AWS API.
type APIError struct {
	StatusCode int // HTTP status code e.g. 200
	Code       string
	Message    string
	RequestID  string
}

func (e APIError) Error() string {
	return e.Code + ": " + e.Message
}

func Error(e error) *APIError {
	if err, ok := e.(*APIError); ok {
		return err
	} else if err, ok := e.(APIError); ok {
		return &err
	} else {
		return nil
	}
}
