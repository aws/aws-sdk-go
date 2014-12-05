package aws

type Client interface {
	Do(op, method, uri string, req, resp interface{}) error
}
