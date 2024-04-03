package main

import (
	"log"
	// "time"

	pb "github.com/jhayakrishnaKS/GO_GRPC/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	// "google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:5051"

func main() {
	tls:=true
	opts:=[]grpc.DialOption{}
	if tls{
		certFile:="ssl/ca.crt"
		creds,err:=credentials.NewClientTLSFromFile(certFile,"")
		if err!=nil{
			log.Fatalf("Error while loading CA trust Certificates:%v\n",err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	conn, err := grpc.Dial(addr,opts...)

	if err!=nil{
		log.Fatalf("failed to connect:%v\n",err)
	}
	defer conn.Close()

	c:=pb.NewGreetServiceClient(conn)

	doGreet(c)
	// doGreetManyTimes(c)
	// doMul(c)
	// doLongGreet(c)
	// doGreetEveryone(c)
	// doGreetWithDeadline(c, 5*time.Second)

}