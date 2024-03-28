package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/jhayakrishnaKS/GO_GRPC/calculator/proto"
)

type Server struct {
	pb.CalculatorServiceServer
}

var addr string = "0.0.0.0:5051"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Listening on %s", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}