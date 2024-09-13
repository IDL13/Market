path_reset:
	export PATH=$PATH:$(go env GOPATH)/bin
	export PATH="$GOPATH/bin:$PATH"

run:
	go run main.go

build:
	go build -o market ./cmd/main.go

grpc_gen:
	protoc -I=api/proto/ --go_out=pkg/ api/proto/market.proto

	protoc --go_out=pkg/ --go_opt=paths=source_relative \
    --go-grpc_out=pkg/ --go-grpc_opt=paths=source_relative \
    api/proto/market.proto