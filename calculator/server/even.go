package main

import (
	"io"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/calculator/proto"
	
)

func(s *Server)Even(stream pb.CalculatorService_EvenServer) error{
	log.Println("Even function is invoked")

	var even int32

	for{
		req,err:=stream.Recv()

		if err==io.EOF{
			return nil
		}

		if err!=nil{
			log.Fatalf("Error while reading client streaming: %v\n",err)
		}

		if req.Number%2==0{
			even=req.Number
			err:=stream.Send(&pb.EvenResponse{
				Result: even,
			})
			if err!=nil{
				log.Fatalf("Error while sending data to client: %v",err)
			}
		}
		
		
	}
}