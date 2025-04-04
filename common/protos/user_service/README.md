# User service

### To regenerate the user_service/user.pb.go and user_service/user_grpc.pb.go run:

`cd common/protos/`

`protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative user_service/user.proto`