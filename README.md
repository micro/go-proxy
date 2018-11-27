# Go Proxy [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/github.com/micro/go-proxy?status.svg)](https://godoc.org/github.com/micro/go-proxy)

The go-proxy library provides proxy enabled micro services.

## Overview

There are situations where we're running dozens if not hundreds of unique micro service applications. 
While go-micro gives us a useful framework for writing those applications, certain aspects we may 
want to run and manage separately e.g auth, tracing, rate limiting, etc.

The go-proxy library let's us create proxied versions of go-micro services.

## Supported

- **Consul Connect** native Go Micro services

Contributions welcome for:

- Istio
- Linkerd
- ?

## Usage

### Server

Create a Consul Connect native micro service.

```go
import (
	"github.com/micro/go-micro"
	"github.com/micro/go-proxy/proxy/connect"
)

// Create a Consul Connect service
service := connect.NewService(
	micro.Service("greeter"),
)
```

### Client

```go
import (
	"github.com/micro/go-proxy/proxy/connect"
)

// create a new connect enabled service
service := connect.NewService()

// now use the client
greeter := proto.NewGreeterService("greeter", service.Client())
```

