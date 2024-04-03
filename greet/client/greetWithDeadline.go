package main

import (
	"context"
	"log"
	"time"

	pb "github.com/jhayakrishnaKS/GO_GRPC/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient,timeout time.Duration){
	log.Println("DoGreetWithDeadline is invoked")

	ctx,cancel:=context.WithTimeout(context.Background(),timeout)
	defer cancel()
	
	req:=&pb.GreetRequest{
		FirstName: "krish",
	}
	res,err:=c.GreetWithDeadline(ctx,req)

	if err!=nil{
		e,ok:=status.FromError(err)
		if ok{
			if e.Code()==codes.DeadlineExceeded{
				log.Fatalf("DeadLine Exceeded !")
				return
			}else{
				log.Fatalf("unExpected gRPC error: %v\n",err)
			}
		}else{
			log.Fatalf("A Non gRPC error: %v\n",e)
		}
	}
	log.Printf("GreetWithDeadline: %s\n",res.Result)
}