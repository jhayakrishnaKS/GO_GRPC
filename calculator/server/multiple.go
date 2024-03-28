package main

import (
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/calculator/proto"
)

func(s *Server)Multiples(in *pb.MultipleRequest,stream pb.CalculatorService_MultiplesServer) error{
	log.Printf("Multiples function was invoked %v\n",in)
	
	var number int64=in.Number
	var i int64
	for i=1;i<=number;i++{
		var multiple int64 = number*i
		stream.Send(&pb.MultipleResponse{
			Result: int64(multiple),
		})
	}
	return nil
}