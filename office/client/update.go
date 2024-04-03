package main

import (
	"context"
	"log"
	"fmt"

	pb "github.com/jhayakrishnaKS/GO_GRPC/office/proto"
)

func UpdateEmp(c pb.OfficeServiceClient) {
	log.Println("----updateEmp is invoked----")

	log.Println("Enter Employee ID to update: ")
	var id string
	fmt.Scanln(&id)

	log.Println("Enter new Employee ID: ")
	var empId string
	fmt.Scanln(&empId)

	log.Println("Enter new Employee Name: ")
	var empName string
	fmt.Scanln(&empName)

	log.Println("Enter new Employee Age: ")
	var empAge string
	fmt.Scanln(&empAge)

	newEmp := &pb.Office{
		Id:     id,
		EmpId:  empId,
		EmpName: empName,
		EmpAge:  empAge,
	}
	_, err := c.UpdateEmp(context.Background(), newEmp)

	if err != nil {
		log.Fatalf("Error happened while updating : %v\n", err)
	}

	log.Fatalln("Emp is updated!")
}
