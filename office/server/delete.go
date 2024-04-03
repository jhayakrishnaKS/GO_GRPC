package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/office/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func(s *Server)DeleteEmp(ctx context.Context,in *pb.EmpId) (*emptypb.Empty, error){ 
	log.Printf("DeleteEmp is invoked: %v\n",in)

	oid,err:=primitive.ObjectIDFromHex(in.Id)

	if err!=nil{
		return nil,status.Errorf(
			codes.Internal,
			"Cannot parse Id",
		)
	}

	res,err:=collection.DeleteOne(ctx,bson.M{"_id":oid})

	if err!=nil{
		return nil,status.Errorf(
			codes.Internal,
			fmt.Sprintf("cannot delete object in mongoDb: %v\n",err),
		)
	}

	if res.DeletedCount==0{
		return nil,status.Errorf(
			codes.NotFound,
			"emp is not found",
		)
	}

	return &emptypb.Empty{},nil
}