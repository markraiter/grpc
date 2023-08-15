package main

import (
	"context"
	"log"

	pb "github.com/markraiter/grpc/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoket with %v\n", in)

	return &pb.SumResponse{Result: in.FirstDigit + in.SecondDigit}, nil
}
