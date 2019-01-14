package http

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"testing"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry/mock"
	"github.com/micro/go-micro/server"
)

type testHandler struct{}

func (t *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"hello": "world"}`))
}

func TestHTTPRouter(t *testing.T) {
	c, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	addr := c.Addr().String()

	url := fmt.Sprintf("http://%s", addr)

	testCases := []struct {
		// http endpoint to call e.g /foo/bar
		httpEp string
		// rpc endpoint called e.g Foo.Bar
		rpcEp string
		// should be an error
		err bool
	}{
		{"/", "Foo.Bar", false},
		{"/", "Foo.Baz", false},
		{"/helloworld", "Hello.World", true},
	}

	// handler
	http.Handle("/", new(testHandler))

	// new proxy
	p := NewSingleHostRouter(url)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	// new micro service
	service := micro.NewService(
		micro.Context(ctx),
		micro.Name("foobar"),
		micro.Registry(mock.NewRegistry()),
		micro.AfterStart(func() error {
			wg.Done()
			return nil
		}),
	)

	// set router
	service.Server().Init(
		server.WithRouter(p),
	)

	// run service
	// server
	go http.Serve(c, nil)
	go service.Run()

	// wait till service is started
	wg.Wait()

	for _, test := range testCases {
		req := service.Client().NewRequest("foobar", test.rpcEp, map[string]string{"foo": "bar"}, client.WithContentType("application/json"))
		var rsp map[string]string
		err := service.Client().Call(ctx, req, &rsp)
		if err != nil && test.err == false {
			t.Fatal(err)
		}
		if v := rsp["hello"]; v != "world" {
			t.Fatalf("Expected hello world got %s from %s", v, test.rpcEp)
		}
	}
}

func TestHTTPRouterOptions(t *testing.T) {
	// test endpoint
	service := NewService(
		WithBackend("http://foo.bar"),
	)

	r := service.Server().Options().Router
	httpRouter, ok := r.(*Router)
	if !ok {
		t.Fatal("Expected http router to be installed")
	}
	if httpRouter.Backend != "http://foo.bar" {
		t.Fatalf("Expected endpoint http://foo.bar got %v", httpRouter.Backend)
	}

	// test router
	service = NewService(
		WithRouter(&Router{Backend: "http://foo2.bar"}),
	)
	r = service.Server().Options().Router
	httpRouter, ok = r.(*Router)
	if !ok {
		t.Fatal("Expected http router to be installed")
	}
	if httpRouter.Backend != "http://foo2.bar" {
		t.Fatalf("Expected endpoint http://foo2.bar got %v", httpRouter.Backend)
	}
}
