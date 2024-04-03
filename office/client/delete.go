package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/office/proto"
)

func DeleteEmp(c pb.OfficeServiceClient) {
    log.Println("----deleteEmp is invoked----")

    fmt.Print("Enter Employee ID to delete: ")
	var id string
	fmt.Scanln(&id)

    resp, err := c.DeleteEmp(context.Background(), &pb.EmpId{Id: id})
    if err != nil {
        log.Fatalf("Error while deleting: %v\n", err)
    }

    log.Println("DeleteEmp response:", resp)
    log.Println("Emp is deleted successfully")
}
