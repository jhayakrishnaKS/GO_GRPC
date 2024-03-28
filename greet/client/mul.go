package main

import (
	"context"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/greet/proto"
	
)

func doMul(c pb.GreetServiceClient){
	log.Println("doMul is invoked")
	res,err:=c.Mul(context.Background(),&pb.MulRequest{
		FirstNumber: 30,
		SecondNumber: 10,
	})
	if err!=nil{
		log.Fatalf("could not sum: %v ",err)
	}
	log.Printf("Sum: %d\n",res.Result)
}