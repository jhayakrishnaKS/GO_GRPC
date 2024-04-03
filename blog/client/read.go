package main

import (
	"context"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/blog/proto"
)

func readBlog(c pb.BlogServiceClient,id string) *pb.Blog{
	log.Println("----readBlog is invoked----")

	req:=&pb.BlogId{Id:id}
	res,err:=c.ReadBlog(context.Background(),req)

	if err!=nil{
		log.Printf("Error happend while reading: %v\n",err)
	}

	log.Printf("Blog is read: %v\n",res)
	return res
}