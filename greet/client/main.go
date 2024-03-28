package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/jhayakrishnaKS/GO_GRPC/greet/proto"
)

var addr string = "localhost:5051"

func main() {
	conn, err := grpc.Dial(addr,grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err!=nil{
		log.Fatalf("failed to connect:%v\n",err)
	}
	defer conn.Close()

	c:=pb.NewGreetServiceClient(conn)

	// doGreet(c)
	// doGreetManyTimes(c)
	// doMul(c)
	// doLongGreet(c)
	doGreetEveryone(c)

}