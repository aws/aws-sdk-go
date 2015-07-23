package machinelearning

import (
	"net/url"

	"github.com/aws/aws-sdk-go/aws/service"
)

func init() {
	initRequest = func(r *service.Request) {
		switch r.Operation.Name {
		case opPredict:
			r.Handlers.Build.PushBack(updatePredictEndpoint)
		}
	}
}

// updatePredictEndpoint rewrites the request endpoint to use the
// "PredictEndpoint" parameter of the Predict operation.
func updatePredictEndpoint(r *service.Request) {
	if !r.ParamsFilled() {
		return
	}

	r.Endpoint = *r.Params.(*PredictInput).PredictEndpoint

	uri, err := url.Parse(r.Endpoint)
	if err != nil {
		r.Error = err
		return
	}
	r.HTTPRequest.URL = uri
}
