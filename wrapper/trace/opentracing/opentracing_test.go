package opentracing_test

import (
	"context"
	"fmt"
	"testing"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/stretchr/testify/assert"
	rbroker "github.com/unistack-org/micro-broker-memory/v3"
	cli "github.com/unistack-org/micro-client-grpc/v3"
	jsoncodec "github.com/unistack-org/micro-codec-json/v3"
	rmemory "github.com/unistack-org/micro-register-memory/v3"
	rrouter "github.com/unistack-org/micro-router-register/v3"
	srv "github.com/unistack-org/micro-server-grpc/v3"
	otwrapper "github.com/unistack-org/micro-wrapper-trace-opentracing/v3"
	"github.com/unistack-org/micro/v3/broker"
	"github.com/unistack-org/micro/v3/client"
	"github.com/unistack-org/micro/v3/errors"
	"github.com/unistack-org/micro/v3/router"
	"github.com/unistack-org/micro/v3/server"
)

type Test interface {
	Method(ctx context.Context, in *TestRequest, opts ...client.CallOption) (*TestResponse, error)
}

type TestRequest struct {
	IsError bool
}
type TestResponse struct {
	Message string
}

type testHandler struct{}

func (t *testHandler) Method(ctx context.Context, req *TestRequest, rsp *TestResponse) error {
	if req.IsError {
		return errors.BadRequest("bad", "test error")
	}

	rsp.Message = "passed"

	return nil
}

func TestClient(t *testing.T) {
	// setup
	assert := assert.New(t)
	for name, tt := range map[string]struct {
		message     string
		isError     bool
		wantMessage string
		wantStatus  string
	}{
		"OK": {
			message:     "passed",
			isError:     false,
			wantMessage: "passed",
			wantStatus:  "OK",
		},
		"Invalid": {
			message:     "",
			isError:     true,
			wantMessage: "",
			wantStatus:  "InvalidArgument",
		},
	} {
		t.Run(name, func(t *testing.T) {
			tracer := mocktracer.New()

			reg := rmemory.NewRegister()
			brk := rbroker.NewBroker(broker.Register(reg))

			serverName := "micro.server.name"
			serverID := "id-1234567890"
			serverVersion := "1.0.0"

			rt := rrouter.NewRouter(router.Register(reg))

			c := cli.NewClient(
				client.Codec("application/grpc+json", jsoncodec.NewCodec()),
				client.Codec("application/json", jsoncodec.NewCodec()),
				client.Router(rt),
				client.WrapCall(otwrapper.NewClientCallWrapper(otwrapper.WithTracer(tracer))),
			)

			s := srv.NewServer(
				server.Codec("application/grpc+json", jsoncodec.NewCodec()),
				server.Codec("application/json", jsoncodec.NewCodec()),
				server.Name(serverName),
				server.Version(serverVersion),
				server.Id(serverID),
				server.Register(reg),
				server.Broker(brk),
				server.WrapSubscriber(otwrapper.NewServerSubscriberWrapper(otwrapper.WithTracer(tracer))),
				server.WrapHandler(otwrapper.NewServerHandlerWrapper(otwrapper.WithTracer(tracer))),
			)
			if err := s.Init(); err != nil {
				t.Fatal(err)
			}
			defer s.Stop()

			type Test struct {
				*testHandler
			}

			s.Handle(s.NewHandler(&Test{new(testHandler)}))

			if err := s.Start(); err != nil {
				t.Fatalf("Unexpected error starting server: %v", err)
			}

			ctx, span, err := otwrapper.StartSpanFromContext(context.Background(), tracer, "root")
			assert.NoError(err)

			req := c.NewRequest(serverName, "Test.Method", &TestRequest{IsError: tt.isError}, client.WithContentType("application/json"))
			rsp := TestResponse{}
			err = c.Call(ctx, req, &rsp)

			if tt.isError {
				assert.Error(err)
			} else {
				assert.NoError(err)
			}
			assert.Equal(rsp.Message, tt.message)

			span.Finish()

			spans := tracer.FinishedSpans()
			assert.Len(spans, 3)

			var rootSpan opentracing.Span
			for _, s := range spans {
				fmt.Printf("%#+v\n", s)
				// order of traces in buffer is not garanteed
				switch s.OperationName {
				case "root":
					rootSpan = s
				}
			}

			for _, s := range spans {
				assert.Equal(rootSpan.Context().(mocktracer.MockSpanContext).TraceID, s.Context().(mocktracer.MockSpanContext).TraceID)
			}
		})
	}
}
