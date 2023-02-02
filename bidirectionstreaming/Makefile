test:
	echo "Hello"


protoc:
	mkdir -p go client server
	protoc order_management.proto -I=./proto --go_out=go  
	protoc order_management.proto -I=./proto --go-grpc_out=require_unimplemented_servers=false:go
	cp -r ./go/bidirectionstreaming ./client
	cp -r ./go/bidirectionstreaming ./server

	echo "protoc finished"