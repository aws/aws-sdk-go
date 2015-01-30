package restjson

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/protocol/jsonrpc"
	"github.com/awslabs/aws-sdk-go/aws/protocol/rest"
)

func Build(r *aws.Request) {
	rest.Build(r)

	m := rest.PayloadMember(r.Params, r.Operation.InPayload)
	if m != nil || r.Operation.InPayload == "" {
		jsonrpc.Build(r)
	}
}

func Unmarshal(r *aws.Request) {
	rest.Unmarshal(r)

	m := rest.PayloadMember(r.Data, r.Operation.OutPayload)
	if m != nil || r.Operation.OutPayload == "" {
		jsonrpc.Unmarshal(r)
	}
}
