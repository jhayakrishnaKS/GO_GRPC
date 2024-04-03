package main

import (
	"context"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/blog/proto"
)

func DeleteBlog(c pb.BlogServiceClient,id string){
	log.Println("----deleteBlog is invoked----")

	_,err:=c.DeleteBlog(context.Background(),&pb.BlogId{Id: id})
	
	if err!=nil{
		log.Fatalf("Error while deleting: %v\n",err)
	}

	log.Println("Blog is deleted")
}