package main

import (
	"context"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient){
	log.Println("doAvg is invoked")

	stream,err:=c.Avg(context.Background())

	if err!=nil{
		log.Fatalf("Error while opening the stream: %v\n",err)
	}

	numbers:=[]int32{3,5,9,54,23}
	for _,number:=range numbers{
		log.Printf("Sending number: %d\n",number)
		stream.Send(&pb.AvgRequest{
			FirstNumber: number,
		})
	}
	res,err:=stream.CloseAndRecv()
	if err!=nil{
		log.Fatalf("Error while receiving the stream: %v\n",err)
	}
	log.Printf("Avg: %v\n",res)
}