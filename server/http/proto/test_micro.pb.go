// Code generated by protoc-gen-micro
// source: test.proto
package pb

import (
	"context"

	micro_api "github.com/unistack-org/micro/v3/api"
	micro_client "github.com/unistack-org/micro/v3/client"
)

// NewTestEndpoints provides api endpoints metdata for Test service
func NewTestEndpoints() []*micro_api.Endpoint {
	endpoints := make([]*micro_api.Endpoint, 0, 2)
	var endpoint *micro_api.Endpoint
	endpoint = &micro_api.Endpoint{
		Name:    "Test.Call",
		Path:    []string{"/v1/test/call/{name}"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}
	endpoints = append(endpoints, endpoint)
	endpoint = &micro_api.Endpoint{
		Name:    "Test.CallError",
		Path:    []string{"/v1/test/callerror/{name}"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}
	endpoints = append(endpoints, endpoint)
	return endpoints
}

// TestService interface
type TestService interface {
	Call(context.Context, *CallReq, ...micro_client.CallOption) (*CallRsp, error)
	CallError(context.Context, *CallReq1, ...micro_client.CallOption) (*CallRsp1, error)
}

// Micro server stuff

// TestHandler server handler
type TestHandler interface {
	Call(context.Context, *CallReq, *CallRsp) error
	CallError(context.Context, *CallReq1, *CallRsp1) error
}
