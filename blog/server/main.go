package main

import (
	"context"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	pb "github.com/jhayakrishnaKS/GO_GRPC/blog/proto"
)

var collection *mongo.Collection

type Server struct {
	pb.BlogServiceServer
}

var addr string = "0.0.0.0:5051"

func main() {

	client,err:=mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))

	if err!=nil{
		log.Fatal(err)
	}
	err=client.Connect(context.Background())

	if err!=nil{
		log.Fatal(err)
	}

	collection=client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Listening on %s", addr)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}