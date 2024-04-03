package main

import (
	pb "github.com/jhayakrishnaKS/GO_GRPC/office/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type officeItem struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	EmpId   string             `bson:"emp_id"`
	EmpName string             `bson:"emp_name"`
	EmpAge  string             `bson:"emp_age"`
}

func documentToOffice(data *officeItem) *pb.Office {
	return &pb.Office{
		Id:      data.ID.Hex(),
		EmpId:   data.EmpId,
		EmpName: data.EmpName,
		EmpAge:  data.EmpAge,
	}
}
