package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/jhayakrishnaKS/GO_GRPC/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone is invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "krish"},
		{FirstName: "Test"},
		{FirstName: "laptop"},
	}

	waitc:=make(chan struct{})

	go func() {
		for _,req:=range reqs{
			log.Printf("send request: %v\n",req)
			stream.Send(req)
			time.Sleep(1*time.Second)
		}
		stream.CloseSend()
	}()

	go func ()  {
		for{
			res,err:= stream.Recv()

			if err==io.EOF{
				break
			}
			
			if err!=nil{
				log.Printf("Error while receving :%v\n",err)
				break
			}

			log.Printf("Received: %v\n",res.Result)

		}
		close(waitc)
	}()

	<-waitc
}
