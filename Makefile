generate_api:
	protoc \
	  --proto_path=./pkg/proto \
	  --go_out=./pkg/api/file-service \
	  --go_opt=paths=source_relative \
	  --go-grpc_out=./pkg/api/file-service \
	  --go-grpc_opt=paths=source_relative \
	  --grpc-gateway_out=./pkg/api/file-service \
	  --grpc-gateway_opt=paths=source_relative \
	  --grpc-gateway_opt=generate_unbound_methods=true \
	  pkg/proto/file-service/*.proto

	protoc \
	  --proto_path=./pkg/proto \
	  --go_out=. \
	  --go-grpc_out=. \
	  --grpc-gateway_out=. \
	  pkg/proto/text-service/text.proto

	protoc \
	  --proto_path=./pkg/proto \
	  --go_out=. \
	  --go-grpc_out=. \
	  --grpc-gateway_out=. \
	  pkg/proto/user-service/user.proto


build:
	go mod download
	go build -o service ./cmd/main.go

run: generate_api build


