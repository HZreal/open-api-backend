package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "open-api-backend/rpc/proto"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 测试 Add 方法
	r, err := client.Add(ctx, &pb.AddRequest{A: 10, B: 20})
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	log.Printf("Add: %d", r.Result)

	// 测试 Subtract 方法
	r2, err := client.Subtract(ctx, &pb.SubtractRequest{A: 10, B: 5})
	if err != nil {
		log.Fatalf("could not subtract: %v", err)
	}
	log.Printf("Subtract: %d", r2.Result)
}
