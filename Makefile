generate_api:
	protoc \
	  --proto_path=./pkg/proto \
	  --go_out=. \
	  --go-grpc_out=. \
	  --grpc-gateway_out=. \
	  pkg/proto/file-service/file-service.proto

	protoc \
  	  --proto_path=./pkg/proto \
  	  --go_out=. \
  	  --go-grpc_out=. \
  	  --grpc-gateway_out=. \
  	  pkg/proto/video-service/video.proto

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

