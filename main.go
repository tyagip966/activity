//go:generate protoc ./aishW.proto --go_out=plugins=grpc:./pb
package main

import (
	"aishW/models/activity"
	"aishW/models/user"
	"aishW/pb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))

	if err != nil {
		log.Panic(err)
	}

	serv := grpc.NewServer()
	pb.RegisterActivityServiceServer(serv, &activity.Service{})
	pb.RegisterUserServiceServer(serv,&user.ServiceUser{})
	reflection.Register(serv)
	log.Println("Server Started at ",8080)
	if err := serv.Serve(lis); err != nil {
		log.Fatalf("could not start the server: %v", err)
	}
}
