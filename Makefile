# Generate the Go code from our protobuf definitions
.PHONY: proto
proto:
	protoc --proto_path=proto --go_out=proto --go_opt=paths=source_relative --go-grpc_out=proto --go-grpc_opt=paths=source_relative analyser.proto