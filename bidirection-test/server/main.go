package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "server/bidirection"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	port = ":50051"
)

// server is used to implement ecommerce/product_info.
type server struct {
	productReq map[string]*pb.Req
	// productRes map[string]*pb.Res
}

func (s *server) Send(ctx context.Context, in *pb.Req) (*pb.Res, error) {

	fmt.Println(in)

	// out, err := uuid.NewV4()

	if s.productReq == nil {
		s.productReq = make(map[string]*pb.Req)
	}

	// in.Req = out.String()

	// s.productReq[in.Req] = in

	s.productReq["hi"] = in

	return &pb.Res{Res: "hi"}, status.New(codes.OK, "").Err()

}

// func (s *server) Recv(ctx context.Context, in *pb.Res) (*pb.Req, error) {

// 	// fmt.Println(in)

// 	product, exists := s.productReq[in.Res]
// 	fmt.Println(exists)

// 	return product, status.New(codes.OK, "").Err()

// }

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterTestGRPCServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
