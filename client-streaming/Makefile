test:
	echo "Hello"


protoc:
	mkdir -p go
	protoc test.proto -I=. --go_out=go  
	protoc test.proto -I=. --go-grpc_out=require_unimplemented_servers=false:go
	cp -r ./go/bidirection ./client
	cp -r ./go/bidirection ./server

	echo "protoc finished"



