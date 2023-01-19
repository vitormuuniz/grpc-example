package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-stream-live/pb"
	"io"
	"log"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewMathServiceClient(connection)

	defer connection.Close()

	request := &pb.FibonacciRequest{
		Number: 10,
	}

	responseStream, err := client.Fibonacci(context.Background(), request)
	if err != nil {
		panic(err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		log.Printf("Fibonacci: %v", stream.GetResult())
	}
}

func Sum(a float32, b float32, client pb.MathServiceClient) {
	request := &pb.NewSumRequest{
		Sum: &pb.Sum{
			A: a,
			B: b,
		},
	}

	response, err := client.Sum(context.Background(), request)
	if err != nil {
		panic(err)
	}

	log.Println(response)
}
