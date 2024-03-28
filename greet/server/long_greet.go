package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("Long Greet function is invoked")

	res := ""

	for {
		req, err := stream.Recv()
		if err==io.EOF{
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err!=nil{
			log.Fatalf("Error while reading client stream: %v\n",err)
		}
		log.Printf("Receving :%v\n",req)
		res+=fmt.Sprintf("Hello %s!\n",req.FirstName)
	}
}
