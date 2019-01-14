# Control Plane

This is a service managed by consul connect.

## Overview

- The `server/main.go` program demonstrates a helloworld service using consul connect as the control plane.
- The `client/main.go` program demonstrates how to create a client to call the service.

## Usage

The two programs depend on Consul.

```bash
# run consul
consul agent -dev

# run the service
go run server/main.go

# run the client
go run client/main.go
```
