package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/office/proto"
)

func createEmp(c pb.OfficeServiceClient) string {
	log.Println("----createEmp is invoked----")

	var empId, empName, empAge string

	fmt.Print("Enter Employee ID: ")
	fmt.Scanln(&empId)

	fmt.Print("Enter Employee Name: ")
	fmt.Scanln(&empName)

	fmt.Print("Enter Employee Age: ")
	fmt.Scanln(&empAge)

	office := &pb.Office{
		EmpId:   empId,
		EmpName: empName,
		EmpAge:  empAge,
	}

	res, err := c.CreateEmp(context.Background(), office)

	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}

	log.Printf("Emp has been created: %s\n", res.Id)
	return res.Id
}
