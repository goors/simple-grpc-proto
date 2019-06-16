# rpc stuff sample with handler and one method

How to run?

1. `protoc --go_out=plugins=grpc:. api/*.proto`
2. Run server `go run server/main.go`
3. Test client with `go run client/main.go`

