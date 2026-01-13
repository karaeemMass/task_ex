package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"task_ex/internal/database"
	"task_ex/internal/handler"

	"task_ex/internal/repository"
	"task_ex/internal/service"
	pb "task_ex/service/pb"
)

func main() {
	db, _ := database.NewMySQLDB()

	repo := repository.NewTaskRepository(db)
	svc := service.NewTaskService(repo)
	handler := handler.NewTaskHandler(svc)

	port := ":9080"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTaskServiceServer(grpcServer, handler)

	log.Println("gRPC server running on port", port)
	grpcServer.Serve(lis)
}
