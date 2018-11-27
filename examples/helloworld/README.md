# Hello World

This is a Consul Connect enabled go-micro service example.

## Overview

- The `main.go` program demonstrates a simple helloworld example using consul connect for mtls. 
- The `client/main.go` program demonstrates how to create a client to call the service.

## Usage

The two programs depend on Consul.

```bash
# run consul
consul agent -dev

# run the service
go run main.go

# run the client
go run client/main.go
```
