
generate api:
#	mkdir -p ./pkg/api/file-service/ кому нужно - делайте сами пж)) у меня иначе мейки не работают
#	mkdir -p ./pkg/api/user-service/
	protoc  --proto_path=./pkg/proto/file-service --go_out=./pkg/api/file-service --go_opt=paths=source_relative  --go-grpc_out=./pkg/api/file-service --go-grpc_opt=paths=source_relative ./pkg/proto/file-service/*.proto
	protoc  --proto_path=./pkg/proto/user-service --go_out=./pkg/api/user-service --go_opt=paths=source_relative  --go-grpc_out=./pkg/api/user-service --go-grpc_opt=paths=source_relative ./pkg/proto/user-service/*.proto
	protoc -I. -Igoogle/api --go_out=. --go-grpc_out=. --grpc-gateway_out=. pkg/proto/user-service/user.proto


build:
	go mod download
	go build -o service ./cmd/main.go

run: generate api build


