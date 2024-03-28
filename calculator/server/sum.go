package main

import (
	"context"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function is invoked with %v\n",in)
	return &pb.SumResponse{
		Result: in.FirstNumber+in.SecondNumber,
	},nil
}
