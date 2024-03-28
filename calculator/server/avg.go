package main

import (
	"io"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/calculator/proto"
)

func(s *Server) Avg(stream pb.CalculatorService_AvgServer) error{
	log.Println("Avg fuction is invoked")
	sum:=0
	count:=0

	for{
		req,err:=stream.Recv()

		if err==io.EOF{
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum)/float64(count),
			})
		}

		if err!=nil{
			log.Fatalf("Error while reading stream: %v\n",err)
		}
		log.Printf("receiving number : %d\n",req.FirstNumber)
		sum+=int(req.FirstNumber)
		count++

	}
}

