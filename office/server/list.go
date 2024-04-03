package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/office/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func(s *Server) ListEmp(in *emptypb.Empty,stream pb.OfficeService_ListEmpServer) error{
	log.Println("ListEmp is invoked")

	cur,err:=collection.Find(context.Background(),primitive.D{{}})

	if err!=nil{
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("unknown internal error: %v\n",err),
		)
	}
	defer cur.Close(context.Background())

	for cur.Next((context.Background())){
		data:=&officeItem{}
		err:=cur.Decode(data)

		if err!=nil{
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from mongo db: %v\n",err),
			)
		}
		stream.Send(documentToOffice(data))
	}
	if err=cur.Err();err!=nil{
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("unknow internal error: %v\n",err),
		)
	}

	return nil
}