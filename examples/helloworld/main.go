package main

import (
	"context"
	"fmt"

	proto "github.com/micro/examples/helloworld/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-proxy/micro/connect"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	// create a new connect enabled service
	service := connect.NewService(
		micro.Name("greeter"),
	)

	service.Init()

	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
