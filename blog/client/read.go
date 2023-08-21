package main

import (
	"context"
	"log"

	pb "github.com/markraiter/grpc/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Printf("---readBlog function was invoked---")

	req := &pb.BlogID{Id: id}
	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Printf("Error happen while reading: %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)
	return res

}
