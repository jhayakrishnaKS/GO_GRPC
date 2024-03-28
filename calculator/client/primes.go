package main

import (
	"context"
	"io"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient){
	log.Println("doPrime was invoked")

	req:=&pb.PrimeRequest{
		Number: 12345564,
	}
	stream,err:=c.Primes(context.Background(),req)

	if err!=nil {
		log.Fatalf("error while calling Primes: %v\n",err)
	}

	for{
		res,err:=stream.Recv()

		if err==io.EOF{
			break
		}
		if err!=nil {
			log.Fatalf("error while reading Primes: %v\n",err)
		}

		log.Printf("Primes: %d\n",res.Result)
	}
}