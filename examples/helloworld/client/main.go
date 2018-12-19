package main

import (
	"context"
	"fmt"

	proto "github.com/micro/examples/helloworld/proto"
	"github.com/micro/go-proxy/micro/connect"
)

func main() {
	// create a new connect enabled service
	service := connect.NewService()
	// create the greeter service client
	greeter := proto.NewGreeterService("greeter", service.Client())

	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{
		Name: "Asim",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Greeting)
}
