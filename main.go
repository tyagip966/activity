//go:generate protoc ./aishW.proto --go_out=plugins=grpc:./pb
package main

import (
	"aishW/models/activity"
	"aishW/models/activity/mysql"
	"aishW/models/user"
	mysql2 "aishW/models/user/mysql"
	"aishW/pb"
	"aishW/utils"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	injectDependencies()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8088))

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

func injectDependencies() {
	db,err := utils.InitDBConnection("root:root@tcp(127.0.0.1:3306)/activity")
	if err != nil {
		log.Fatalf("could not make connection with the database: %v", err)
	}

	log.Println("database connected !!")
	activityRepo := mysql.NewActivityRepo(db)
	activity.NewService(activityRepo)

	userRepo := mysql2.NewUserRepo(db)
	user.NewServiceUser(userRepo)
}
