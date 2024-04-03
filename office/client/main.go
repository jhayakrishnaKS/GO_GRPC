package main

import (
	"fmt"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/office/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:5051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewOfficeServiceClient(conn)

	// id := createEmp(c)
	// readEmp(c, id)
	// UpdateEmp(c,id)
	// ListEmp(c)
	// DeleteEmp(c,id)
	var number int32
	for {
		
list()
		fmt.Print("Enter the operation number: ")
		fmt.Scanln(&number)

		switch number {
		case 1:
			createEmp(c)
		case 2:
			readEmp(c)
		case 3:
			UpdateEmp(c)
		case 4:
			ListEmp(c)
		case 5:
			DeleteEmp(c)
		case 0:
			log.Fatal("Exiting the program. Bye Bye!")
		default:
			fmt.Println("Invalid operation number. Please try again.")
		}
	}
}
func list(){
	fmt.Println("Select an operation:")
	fmt.Println("1. Create Employee")
	fmt.Println("2. Read Employee")
	fmt.Println("3. Update Employee")
	fmt.Println("4. List Employees")
	fmt.Println("5. Delete Employee")
	fmt.Println("0. Exit")
}