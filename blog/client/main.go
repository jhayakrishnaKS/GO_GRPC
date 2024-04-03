package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/jhayakrishnaKS/GO_GRPC/blog/proto"
)

var addr string = "localhost:5051"

func main() {
	conn, err := grpc.Dial(addr,grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err!=nil{
		log.Fatalf("failed to connect:%v\n",err)
	}
	defer conn.Close()

	c:=pb.NewBlogServiceClient(conn)

	id:=createBlog(c)
	readBlog(c,id)//valid
	readBlog(c,"blah")//Not Valid
	UpdateBlog(c,id)
	ListBlog(c)
	DeleteBlog(c,id)
}