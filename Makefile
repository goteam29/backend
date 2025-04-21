
generate_api:
#	mkdir -p ./pkg/api/file-service/
#	mkdir -p ./pkg/api/user-service/
	protoc  --proto_path=./pkg/proto/file-service --go_out=./pkg/api/file-service --go_opt=paths=source_relative  --go-grpc_out=./pkg/api/file-service --go-grpc_opt=paths=source_relative ./pkg/proto/file-service/*.proto
	protoc -I . --go_out . --go-grpc_out . --grpc-gateway_out . pkg/proto/text-service/text.proto
	protoc -I . --go_out . --go-grpc_out . --grpc-gateway_out . pkg/proto/user-service/user.proto


build:
	go mod download
	go build -o service ./cmd/main.go

run: generate_api build


