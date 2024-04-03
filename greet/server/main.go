package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/jhayakrishnaKS/GO_GRPC/greet/proto"
)

type Server struct {
	pb.GreetServiceServer
}

var addr string = "0.0.0.0:5051"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Listening on %s", addr)

	opts:=[]grpc.ServerOption{}
	tls:=true


	if tls{
		certFile:="ssl/server.crt"
		keyFile:="ssl/server.pem"
		creds,err:=credentials.NewServerTLSFromFile(certFile,keyFile)
		if err!=nil{
			log.Fatalf("Failed to load certificates: %v\n",err)
		}
		opts=append(opts, grpc.Creds(creds))

	}
	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}