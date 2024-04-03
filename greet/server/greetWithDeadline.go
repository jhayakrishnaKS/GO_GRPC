package main

import (
	"context"
	"log"
	"time"

	pb "github.com/jhayakrishnaKS/GO_GRPC/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func(s *Server)GreetWithDeadline(ctx context.Context,in *pb.GreetRequest) (*pb.GreetResponse, error){
	 log.Printf("GreetWithDeadline is invoked with: %v\n",in)
	 
	 for i:=0;i<3;i++{
		if ctx.Err()==context.DeadlineExceeded{
			log.Println("The client canceled the request")
			return nil,status.Error(codes.Canceled,"The client canceled")
		}

		time.Sleep(1*time.Second)
	 }

	 return &pb.GreetResponse{
		Result: "Hello"+in.FirstName,
	 },nil

}