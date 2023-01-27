package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "client/bidirection"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTestGRPCClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	a, err := c.Send(ctx, &pb.Req{Req: "11"})

	fmt.Println(a)
	// if err != nil {
	// 	log.Fatalf("Could not get product: %v", err)
	// }
	// log.Printf(a.Res)
	// log.Printf("Product: %v", product.String())
	// product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	// if err != nil {
	// 	log.Fatalf("Could not get product: %v", err)
	// }
	// log.Printf("Product: %v", product.String())
}
