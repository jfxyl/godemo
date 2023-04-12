go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

protoc  --go_out=. --go-grpc_out=.  grpc/proto/*.proto