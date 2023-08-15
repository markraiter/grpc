package main

import (
	"context"
	"log"

	pb "github.com/markraiter/grpc/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked...")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstDigit:  3,
		SecondDigit: 7,
	})
	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("The result of the sum is: %d\n", res.Result)
}
