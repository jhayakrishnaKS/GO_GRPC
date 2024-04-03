package main

import (
	"context"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/blog/proto"
)

func UpdateBlog(c pb.BlogServiceClient,id string){
	log.Println("----UpdateBlog is invoked----")

	newBlog:=&pb.Blog{
		Id:id,
		AuthorId: "Not krish",
		Title: "A new title",
		Content: "content of thefirst blog with some inprovements",
	}

	_,err:=c.UpdateBlog(context.Background(),newBlog)

	if err!=nil{
		log.Fatalf("Error happened while updating: %v\n",err)
	}
	
	log.Println("Blog is updated!!")
}