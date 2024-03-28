package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/calculator/proto"
)

func doMultiple(c pb.CalculatorServiceClient) {
	log.Println("DoMultiple is invoked")

	var number int64

	fmt.Println("Enter the number for multiples: ")
	if _, err := fmt.Scanln(&number); err != nil {
		log.Fatalf("Error")
	}

	req := &pb.MultipleRequest{
		Number: number,
	}
	stream, err := c.Multiples(context.Background(), req)

	if err!=nil{
		log.Fatalf("error while calling Multiples: %v\n",err)
	}
	for{
		res,err:=stream.Recv()

		if err==io.EOF{
			break
		}

		if err!=nil{
			log.Fatalf("error while reading Multiples: %v\n",err)
		}

		log.Printf("Multiples: %d\n",res.Result)
	}
}
