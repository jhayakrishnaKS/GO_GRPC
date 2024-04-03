package main

import (
	"context"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/office/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func(s *Server)ReadEmp(ctx context.Context,in *pb.EmpId) (*pb.Office, error){
	log.Printf("ReadEmp is invoked: %v\n",in)

	oid,err:=primitive.ObjectIDFromHex(in.Id)
	
	if err!=nil{
		return nil,status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	data:=&officeItem{}
	filter:=bson.M{"_id":oid}

	res:=collection.FindOne(ctx,filter)

	if err:=res.Decode(data);err!=nil{
		return nil,status.Errorf(
			codes.NotFound,
			"Cannot find blog with the ID Provided",
		)
	}

	return documentToOffice(data),nil

}