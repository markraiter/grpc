package main

import (
	"io"
	"log"

	pb "github.com/markraiter/grpc/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max function was invoked...")
	var max int32 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		if number := req.Num; number > max {
			max = number
			err := stream.Send(&pb.MaxResponse{Result: max})
			if err != nil {
				log.Fatalf("Error while sending data to cliend: %v\n", err)
			}
		}
	}
}
