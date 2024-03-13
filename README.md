# GRPC CONNECT

## tool install

```zsh
$ go install github.com/bufbuild/buf/cmd/buf@latest
$ go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
```

## code gen

```zsh
$ buf generate
```

## server start

```zsh
$ go run ./cmd/server/main.go
```

## request

```zsh
$ curl \
    --header "Content-Type: application/json" \
    --data '{"name": "nakatsu"}' \
    http://localhost:8080/greet.v1.GreetService/Greet
```
