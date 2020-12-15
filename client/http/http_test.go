package http

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"testing"

	mhttp "github.com/unistack-org/micro-client-http"
	jsoncodec "github.com/unistack-org/micro-codec-json"
	rmemory "github.com/unistack-org/micro-registry-memory"
	rrouter "github.com/unistack-org/micro-router-registry"
	"github.com/unistack-org/micro/v3/client"
	"github.com/unistack-org/micro/v3/codec"
	"github.com/unistack-org/micro/v3/registry"
	"github.com/unistack-org/micro/v3/router"
)

var (
	defaultHTTPCodecs = map[string]codec.Codec{
		"application/json": jsoncodec.NewCodec(),
	}
)

type Message struct {
	Seq  int64
	Data string
}

type GithubRsp struct {
	Name string `json:"name,omitempty"`
}

func TestNative(t *testing.T) {
	c := client.NewClientCallOptions(mhttp.NewClient(client.ContentType("application/json"), client.Codec("application/json", jsoncodec.NewCodec())), client.WithAddress("https://api.github.com"))
	req := c.NewRequest("github", "/users/vtolstov", nil)
	rsp := &GithubRsp{}
	err := c.Call(context.TODO(), req, rsp, mhttp.Method(http.MethodGet))
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Name != "Vasiliy Tolstov" {
		t.Fatalf("invlid rsp received: %#+v\n", rsp)
	}

}

func TestHTTPClient(t *testing.T) {
	reg := rmemory.NewRegistry()
	rtr := rrouter.NewRouter(router.Registry(reg))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/foo/bar", func(w http.ResponseWriter, r *http.Request) {
		// only accept post
		if r.Method != "POST" {
			http.Error(w, "expect post method", 500)
			return
		}

		// get codec
		ct := r.Header.Get("Content-Type")
		codec, ok := defaultHTTPCodecs[ct]
		if !ok {
			http.Error(w, "codec not found", 500)
			return
		}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// extract message
		msg := &Message{}
		if err := codec.Unmarshal(b, msg); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// marshal response
		b, err = codec.Marshal(msg)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// write response
		w.Write(b)
	})
	go http.Serve(l, mux)

	if err := reg.Register(ctx, &registry.Service{
		Name: "test.service",
		Nodes: []*registry.Node{
			{
				Id:      "test.service.1",
				Address: l.Addr().String(),
				Metadata: map[string]string{
					"protocol": "http",
				},
			},
		},
	}); err != nil {
		t.Fatal(err)
	}

	c := mhttp.NewClient(client.ContentType("application/json"), client.Codec("application/json", jsoncodec.NewCodec()), client.Router(rtr))

	for i := 0; i < 10; i++ {
		msg := &Message{
			Seq:  int64(i),
			Data: fmt.Sprintf("message %d", i),
		}
		req := c.NewRequest("test.service", "/foo/bar", msg)
		rsp := &Message{}
		err := c.Call(context.TODO(), req, rsp)
		if err != nil {
			t.Fatal(err)
		}
		if rsp.Seq != msg.Seq {
			t.Fatalf("invalid seq %d for %d", rsp.Seq, msg.Seq)
		}
	}
}

func TestHTTPClientStream(t *testing.T) {
	reg := rmemory.NewRegistry()
	rtr := rrouter.NewRouter(router.Registry(reg))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/foo/bar", func(w http.ResponseWriter, r *http.Request) {
		// only accept post
		if r.Method != "POST" {
			http.Error(w, "expect post method", 500)
			return
		}

		// hijack the connection
		hj, ok := w.(http.Hijacker)
		if !ok {
			http.Error(w, "could not hijack conn", 500)
			return

		}

		// hijacked
		conn, bufrw, err := hj.Hijack()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer conn.Close()

		// read off the first request
		// get codec
		ct := r.Header.Get("Content-Type")
		codec, ok := defaultHTTPCodecs[ct]
		if !ok {
			http.Error(w, "codec not found", 500)
			return
		}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// extract message
		msg := &Message{}
		if err := codec.Unmarshal(b, msg); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// marshal response
		b, err = codec.Marshal(msg)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// write response
		rsp := &http.Response{
			Header:        r.Header,
			Body:          ioutil.NopCloser(bytes.NewBuffer(b)),
			Status:        "200 OK",
			StatusCode:    200,
			Proto:         "HTTP/1.1",
			ProtoMajor:    1,
			ProtoMinor:    1,
			ContentLength: int64(len(b)),
		}

		// write response
		rsp.Write(bufrw)
		bufrw.Flush()

		reader := bufio.NewReader(conn)

		for {
			r, err := http.ReadRequest(reader)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			b, err = ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			// extract message
			msg := &Message{}
			if err := codec.Unmarshal(b, msg); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			// marshal response
			b, err = codec.Marshal(msg)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			rsp := &http.Response{
				Header:        r.Header,
				Body:          ioutil.NopCloser(bytes.NewBuffer(b)),
				Status:        "200 OK",
				StatusCode:    200,
				Proto:         "HTTP/1.1",
				ProtoMajor:    1,
				ProtoMinor:    1,
				ContentLength: int64(len(b)),
			}

			// write response
			rsp.Write(bufrw)
			bufrw.Flush()
		}
	})
	go http.Serve(l, mux)

	if err := reg.Register(ctx, &registry.Service{
		Name: "test.service",
		Nodes: []*registry.Node{
			{
				Id:      "test.service.1",
				Address: l.Addr().String(),
				Metadata: map[string]string{
					"protocol": "http",
				},
			},
		},
	}); err != nil {
		t.Fatal(err)
	}

	c := mhttp.NewClient(client.ContentType("application/json"), client.Codec("application/json", jsoncodec.NewCodec()), client.Router(rtr))
	req := c.NewRequest("test.service", "/foo/bar", &Message{})
	stream, err := c.Stream(context.TODO(), req)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	for i := 0; i < 10; i++ {
		msg := &Message{
			Seq:  int64(i),
			Data: fmt.Sprintf("message %d", i),
		}
		err := stream.Send(msg)
		if err != nil {
			t.Fatal(err)
		}
		rsp := &Message{}
		err = stream.Recv(rsp)
		if err != nil {
			t.Fatal(err)
		}
		if rsp.Seq != msg.Seq {
			t.Fatalf("invalid seq %d for %d", rsp.Seq, msg.Seq)
		}
	}
}
