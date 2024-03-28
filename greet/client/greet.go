package main

import (
	"context"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/greet/proto"
)
func doGreet(c pb.GreetServiceClient){
	log.Println("doGreet is invoked")
	res,err:=c.Greet(context.Background(),&pb.GreetRequest{
		FirstName: "krish",
	})

	if err!=nil{
		log.Fatalf("Could not greet:%v\n",err)
	}
	log.Printf("Greeting:%s\n",res.Result)
}