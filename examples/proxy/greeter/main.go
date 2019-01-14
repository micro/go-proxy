package main

import (
	"log"

	"github.com/micro/go-micro"
	proto "github.com/micro/go-proxy/examples/proxy/greeter/proto"

	"context"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	log.Print("Received Greeter.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
	)

	service.Init()

	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
