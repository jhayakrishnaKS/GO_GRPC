package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/jhayakrishnaKS/GO_GRPC/calculator/proto"
)

func doEven(c pb.CalculatorServiceClient) {
	log.Println("doEven is invoked")

	stream, err := c.Even(context.Background())

	if err != nil {
		log.Fatalf("Error while passing stream: %v\n", err)
	}

	waitc := make(chan struct{})

	arr:=[5]int{}

	go func() {
		log.Println("Enter the array of numbers to check even: ")

		for i := int32(0); i < int32(len(arr)); i++ {
			fmt.Scanln(&arr[i])
		}
		for _,number:=range arr{
			log.Printf("Sending number: %d\n",number)
			stream.Send(&pb.EvenRequest{
				Number: int32(number),
			})
			time.Sleep(1*time.Second)
		}
		stream.CloseSend()
	}()
	

	go func ()  {
		for{
			res,err:=stream.Recv()

			if err==io.EOF{
				break
			}

			if err!=nil{
				log.Printf("problem while reading server stream: %v\n",err)
				break
			}

			log.Printf("Received a new even: %d\n",res.Result)

		}
		close(waitc)
	}()
	<-waitc
}
