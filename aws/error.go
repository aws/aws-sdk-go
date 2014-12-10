package aws

// An APIError is an error returned by an AWS API.
type APIError struct {
	Type      string
	Code      string
	Message   string
	RequestID string
	HostID    string
	Specifics map[string]string
}

func (e APIError) Error() string {
	return e.Message
}
