package main

import (
	"context"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/office/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func(s *Server)UpdateEmp(ctx context.Context,in *pb.Office) (*emptypb.Empty, error){
	log.Printf("UpdateEmp is invoked: %v\n",in)

	oid,err:=primitive.ObjectIDFromHex(in.Id)

	if err!=nil{
		return nil,status.Errorf(
			codes.InvalidArgument,
			"cannot parse ID",
		)
	}

	data:=&officeItem{
		EmpId: in.EmpId,
		EmpName: in.EmpName,
		EmpAge: in.EmpAge,
	}

	res,err:=collection.UpdateOne(
		ctx,
		bson.M{"_id":oid},
		bson.M{"$set":data},
	)
	if err!=nil{
		return nil,status.Errorf(
			codes.Internal,
			"could not update",
		)
	}
	if res.MatchedCount==0{
		return nil,status.Errorf(
			codes.NotFound,
			"Cannot find emp with Id",
		)
	}
	return &emptypb.Empty{},nil
}