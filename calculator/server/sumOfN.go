package main

import (
	"io"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/calculator/proto"
)

func(s *Server)SumOfN(stream pb.CalculatorService_SumOfNServer) error{
	log.Println("SumOfN function is invoked")
	var n int32
	var sum int32
	for{
		req,err:=stream.Recv()

		if err==io.EOF{
			return stream.SendAndClose(&pb.SumOfNResponse{
				Result: sum,
			})
		}
	
		if err!=nil{
			log.Fatalf("Error while reading stream: %v\n",err)
		}
		log.Printf("receiving number: %d\n",req.FirstNumber)
		sum+=int32(req.FirstNumber)
		n++

	}
}