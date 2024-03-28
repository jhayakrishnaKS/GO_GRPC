package main

import (
	"context"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/greet/proto"
)
func(s *Server)Mul(ctx context.Context,in *pb.MulRequest) (*pb.MulResponse, error){
	log.Printf("mul function is invoked with%v\n",in)
	return &pb.MulResponse{
		Result: in.FirstNumber*in.SecondNumber,
	},nil
}