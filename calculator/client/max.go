package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/markraiter/grpc/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was invoked...")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []int32{1, 3, 7, 2, 17, 42, 34, 13, 9, 17, 58}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending: %v\n", req)
			stream.Send(&pb.MaxRequest{Num: req})
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
