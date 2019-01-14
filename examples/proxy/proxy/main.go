package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-proxy/router/proxy"
)

func main() {
	// create a new proxy
	p := proxy.NewService(
		micro.Name("go.micro.proxy"),
	)
	// initialise
	p.Init()
	// run the proxy
	p.Run()
}
