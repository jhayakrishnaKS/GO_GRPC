package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/calculator/proto"
)

func doSumOfN(c pb.CalculatorServiceClient) {
	log.Println("doSumOfN is invoked")

	stream, err := c.SumOfN(context.Background())

	if err != nil {
		log.Fatalf("Error while opening the stream: %v\n", err)
	}
	var number int32
	var n int32

	fmt.Println("Enter the number of values for sum: ")
	if _, err := fmt.Scanln(&number); err != nil {
		log.Fatalf("Error %v", err)
	}

	for i := int32(1); i <= number; i++ {
		fmt.Printf("Enter number %d: ", i)
		if _, err := fmt.Scanln(&n); err != nil {
			log.Fatalf("Error %v", err)
		}

		stream.Send(&pb.SumOfNRequest{
			FirstNumber: n,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response: %v\n", err)
	}

	log.Printf("Sum of %d numbers is: %d\n", number, res.Result)
}
