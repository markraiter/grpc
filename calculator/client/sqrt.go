package main

import (
	"context"
	"log"

	pb "github.com/markraiter/grpc/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("doSqrt function was invoked...")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Num: n})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error message from server: %v\n", e.Message())
			log.Printf("Error code from server: %v\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Println("We probably sent a negative number!")
				return
			}
		} else {
			log.Printf("A non gRPC error: %v\n", err)
		}
	}

	log.Printf("Sqrt: %v\n", res.Result)
}
