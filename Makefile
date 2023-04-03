tidy:
	go mod tidy

path_proto:=./src/pb
proto:
	protoc --go_out=${path_proto} --go-grpc_out=${path_proto} ${path_proto}/*.proto
