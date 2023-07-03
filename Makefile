# Generate the Go code from our protobuf definitions
.PHONY: proto
proto:
	protoc --proto_path=proto --go_out=proto --go_opt=paths=source_relative --go-grpc_out=proto --go-grpc_opt=paths=source_relative analyser.proto

# Run tests
test:
	go test -v ./...

# Build the server
server/build:
	cd service && go build -o ../bin/server

# Run the server
server/run:
	cd service && go run .