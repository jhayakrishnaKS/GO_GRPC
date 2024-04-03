package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/office/proto"
)

func readEmp(c pb.OfficeServiceClient) *pb.Office {
	log.Println("----readEmp is invoked----")

	log.Println("Enter Employee ID to read: ")
	var id string
	fmt.Scanln(&id)
	req := &pb.EmpId{Id: id}
	res, err := c.ReadEmp(context.Background(), req)

	if err != nil {
		log.Printf("Error happened while reading: %v\n", err)
	}

	log.Printf("Emp was read: %v\n", res)
	return res
}
