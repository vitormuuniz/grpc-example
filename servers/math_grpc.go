package servers

import (
	"context"
	"grpc-stream-live/pb"
	"time"
)

type Math struct{}

func (m *Math) Sum(ctx context.Context, in *pb.NewSumRequest) (*pb.NewSumResponse, error) {
	res := in.Sum.GetA() + in.Sum.GetB()

	return &pb.NewSumResponse{
		Result: res,
	}, nil
}

func (m *Math) Fibonacci(in *pb.FibonacciRequest, stream pb.MathService_FibonacciServer) error {
	number := in.GetNumber()

	var i int32
	for i = 1; i <= number; i++ {
		res := &pb.FibonacciResponse{
			Result: FibonacciRecursion(i),
		}
		stream.Send(res)
		time.Sleep(2000 * time.Millisecond)
	}
	return nil
}

func FibonacciRecursion(number int32) int32 {
	if number <= 1 {
		return number
	}
	return FibonacciRecursion(number-1) + FibonacciRecursion(number-2)
}
