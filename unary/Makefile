test:
	echo "Hello"


protoc:
	mkdir -p go client server
	protoc product_info.proto -I=./proto --go_out=go  
	protoc product_info.proto -I=./proto --go-grpc_out=require_unimplemented_servers=false:go
	cp -r ./go/unary ./client
	cp -r ./go/unary ./server

	echo "protoc finished"