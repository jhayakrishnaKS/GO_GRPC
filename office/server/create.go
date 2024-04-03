package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/jhayakrishnaKS/GO_GRPC/office/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateEmp(ctx context.Context, in *pb.Office) (*pb.EmpId, error) {
	log.Printf("createEmp is invoked %v\n", in)

	data := officeItem{
		EmpId:   in.EmpId,
		EmpName: in.EmpName,
		EmpAge:  in.EmpAge,
	}

	res, err := collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Intenal Error: %v\n", err),
		)
	}

	oid,ok:=res.InsertedID.(primitive.ObjectID)

	
	if !ok{
		return nil,status.Errorf(
			codes.Internal,
			"Cannot convert to OID",
		)
	}

	return &pb.EmpId{
		Id: oid.Hex(),
	},nil

}
