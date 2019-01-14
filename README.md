# Go Proxy [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/github.com/micro/go-proxy?status.svg)](https://godoc.org/github.com/micro/go-proxy)

Go Proxy is a library for creating micro proxies.

## Overview

Go Micro is a distributed systems framework for client/server communication. It handles the details 
around discovery, fault tolerance, etc. We may want to leverage this in broader ecosystems that use 
standard http or offload a number of requirements.

## Features

- **Single Backend Router** - Enable the single backend router to proxy directly to your local app. The proxy 
allows you to set a router which serves your protocol e.g. http, grpc.

- **Protocol Aware Handler** - Set a request handler which speaks your protocol to make outbound RPC requests.

## Supported

- [x] [Consul](https://www.consul.io/docs/connect/native.html) - Using Connect-Native to provide secure mTLS.
- [x] [NATS](https://nats.io/) - Fully leveraging NATS as the control plane and data plane.

Contributions welcome!

## Usage

### Server

Create a Consul Connect-Native micro service.

```go
import (
	"github.com/micro/go-micro"
	"github.com/micro/go-proxy/service/consul"
)

// Create a Consul Connect service
service := consul.NewService(
	micro.Service("greeter"),
)
```

### Client

```go
import (
	"github.com/micro/go-proxy/service/consul"
)

// create a new consul enabled service
service := consul.NewService()

// now use the client
greeter := proto.NewGreeterService("greeter", service.Client())
```
