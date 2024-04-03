package main

import (
	"context"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/blog/proto"
)

func createBlog(c pb.BlogServiceClient)string{
	log.Println("----createBlog was invoked-----")

	blog:=&pb.Blog{
		AuthorId: "krish",
		Title: "My first Blog",
		Content: "content of the first blog",
	}

	res,err:=c.CreateBlog(context.Background(),blog)

	if err!=nil {
		log.Fatalf("unexpected error: %v",err)
	}

	log.Printf("Blog has been created: %s\n",res.Id)
	return res.Id
}
