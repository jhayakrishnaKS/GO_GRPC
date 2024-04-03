package main

import (
	"context"
	"io"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func ListBlog(c pb.BlogServiceClient){
	log.Println("----ListBlog is invoked----")

	stream,err:=c.ListBlog(context.Background(),&emptypb.Empty{})

	if err!=nil{
		log.Fatalf("Error while calling ListBlogs: %v\n",err)
	}
	for{
		res,err:=stream.Recv()

		if err==io.EOF{
			break
		}

		if err !=nil{
			log.Fatalf("something happened: %v\n",err)
		}
		log.Println(res)
	}
}