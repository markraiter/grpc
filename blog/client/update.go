package main

import (
	"context"
	"log"

	pb "github.com/markraiter/grpc/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---updateBlog was invoked---")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Mark",
		Title:    "A new title",
		Content:  "Brand new content has been updated",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Printf("Error happen while updating: %v\n", err)
	}

	log.Println("Blog was updated")
}
