USER_PROTO_FILE=./common/protos/user_service/user.proto

regenerate user:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative "${USER_PROTO_FILE}"