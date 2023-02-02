test:
	echo "Hello"


protoc:
	mkdir -p go client server
	protoc order_management.proto -I=./proto --go_out=go  
	protoc order_management.proto -I=./proto --go-grpc_out=require_unimplemented_servers=false:go
	cp -r ./go/clientstreaming ./client
	cp -r ./go/clientstreaming ./server

	echo "protoc finished"