# Proxy

Transparently proxy Go Micro apps

## Run Proxy

```
// run the proxy
go run proxy/main.go --server_address=":8081"

// run the greeter
go run greeter/main.go
```

## Call Service

Proxy the request

```
curl \
-H 'Content-Type: application/json' \
-H 'X-Micro-Service: greeter' \
-H 'X-Micro-Endpoint: Greeter.Hello' \
-d '{"name": "John"}' \
http://localhost:8081
```
