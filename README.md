# Go Proxy [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/github.com/micro/go-proxy?status.svg)](https://godoc.org/github.com/micro/go-proxy)

Go Proxy provides the ability to create proxy aware Go Micro services.

## Overview

Go Micro is a distributed systems framework for client/server communication. It handles the details 
around discovery, fault tolerance, etc as a library but this may not make sense at scale. The Go Proxy 
library enables handing off these concerns to a proxy or "service mesh".

Go Proxy let's us create proxied versions of Go Micro services.

## Supported

- [x] [Consul Connect-Native](https://www.consul.io/docs/connect/native.html)
- [x] [NATS](https://nats.io/)
- [ ] Istio
- [ ] Linkerd
- [ ] ?

Contributions welcome!

## Usage

### Server

Create a Consul Connect-Native micro service.

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
