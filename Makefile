
build:
	protoc  --proto_path=./pkg/proto/file-service --go_out=./pkg/api/file-service --go_opt=paths=source_relative  --go-grpc_out=./pkg/api/file-service --go-grpc_opt=paths=source_relative ./pkg/proto/file-service/*.proto
	protoc  --proto_path=./pkg/proto/user-service --go_out=./pkg/api/user-service --go_opt=paths=source_relative  --go-grpc_out=./pkg/api/user-service --go-grpc_opt=paths=source_relative ./pkg/proto/user-service/*.proto