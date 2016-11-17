package rds

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/endpoints"
)

func init() {
	ops := []string{
		opCopyDBSnapshot,
	}
	initRequest = func(r *request.Request) {
		for _, operation := range ops {
			if r.Operation.Name == operation {
				r.Handlers.Build.PushFront(fillPresignedURL)
			}
		}
	}
}

func fillPresignedURL(r *request.Request) {
	fns := map[string]func(r *request.Request){
		opCopyDBSnapshot: copyDBSnapshotPresign,
	}
	if !r.ParamsFilled() {
		return
	}
	if f, ok := fns[r.Operation.Name]; ok {
		f(r)
	}
}

func copyDBSnapshotPresign(r *request.Request) {
	originParams := r.Params.(*CopyDBSnapshotInput)

	if originParams.PreSignedUrl != nil || originParams.DestinationRegion != nil {
		return
	}

	originParams.DestinationRegion = r.Config.Region
	newParams := awsutil.CopyOf(r.Params).(*CopyDBSnapshotInput)
	originParams.PreSignedUrl = presignURL(r, originParams.SourceRegion, newParams)
}

// presignURL will presign the request by using SoureRegion to sign with. SourceRegion is not
// sent to the service, and is only used to not have the SDKs parsing ARNs.
func presignURL(r *request.Request, sourceRegion *string, newParams interface{}) *string {
	cfg := r.Config.Copy(aws.NewConfig().
		WithEndpoint("").
		WithRegion(aws.StringValue(sourceRegion)))

	clientInfo := r.ClientInfo
	clientInfo.Endpoint, clientInfo.SigningRegion = endpoints.EndpointForRegion(
		clientInfo.ServiceName,
		aws.StringValue(cfg.Region),
		aws.BoolValue(cfg.DisableSSL),
		aws.BoolValue(cfg.UseDualStack),
	)

	// Presign a request with modified params
	req := request.New(*cfg, clientInfo, r.Handlers, r.Retryer, r.Operation, newParams, r.Data)
	req.Operation.HTTPMethod = "GET"
	uri, err := req.Presign(5 * time.Minute) // 5 minutes should be enough.
	if err != nil {                          // bubble error back up to original request
		r.Error = err
		return nil
	}

	// We have our URL, set it on params
	return &uri
}
