package restjson

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/protocol/jsonrpc"
	"github.com/awslabs/aws-sdk-go/aws/protocol/rest"
)

func Build(r *aws.Request) {
	rest.Build(r)

	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
}

func Unmarshal(r *aws.Request) {
	m := rest.PayloadMember(r.Data)
	if m != nil {
		jsonrpc.Unmarshal(r)
	}
}

func UnmarshalMeta(r *aws.Request) {
	rest.Unmarshal(r)
}

func UnmarshalError(r *aws.Request) {
	// TODO unmarshal error responses
}
