package glacier

import (
	"encoding/hex"
	"reflect"

	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/service"
)

var (
	defaultAccountID = "-"
)

func init() {
	initRequest = func(r *service.Request) {
		r.Handlers.Validate.PushFront(addAccountID)
		r.Handlers.Validate.PushFront(copyParams) // this happens first
		r.Handlers.Build.PushBack(addChecksum)
		r.Handlers.Build.PushBack(addAPIVersion)
	}
}

func copyParams(r *service.Request) {
	r.Params = awsutil.CopyOf(r.Params)
}

func addAccountID(r *service.Request) {
	if !r.ParamsFilled() {
		return
	}

	v := reflect.Indirect(reflect.ValueOf(r.Params))
	if f := v.FieldByName("AccountId"); f.IsNil() {
		f.Set(reflect.ValueOf(&defaultAccountID))
	}
}

func addChecksum(r *service.Request) {
	if r.Body == nil {
		return
	}

	h := ComputeHashes(r.Body)

	if r.HTTPRequest.Header.Get("X-Amz-Content-Sha256") == "" {
		hstr := hex.EncodeToString(h.LinearHash)
		r.HTTPRequest.Header.Set("X-Amz-Content-Sha256", hstr)
	}
	if r.HTTPRequest.Header.Get("X-Amz-Sha256-Tree-Hash") == "" {
		hstr := hex.EncodeToString(h.TreeHash)
		r.HTTPRequest.Header.Set("X-Amz-Sha256-Tree-Hash", hstr)
	}
}

func addAPIVersion(r *service.Request) {
	r.HTTPRequest.Header.Set("X-Amz-Glacier-Version", r.Service.APIVersion)
}
