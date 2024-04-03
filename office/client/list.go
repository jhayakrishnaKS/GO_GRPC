package main

import (
	"context"
	"io"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/office/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func ListEmp(c pb.OfficeServiceClient) {
    log.Println("----ListEmp is invoked----")

    stream, err := c.ListEmp(context.Background(), &emptypb.Empty{})
    if err != nil {
        log.Fatalf("Error while calling ListEmp: %v\n", err)
    }

    for {
        res, err := stream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatalf("Error while receiving stream: %v\n", err)
        }
        log.Println(res)
    }
}
