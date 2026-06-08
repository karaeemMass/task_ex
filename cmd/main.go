package main

import (
	"log"
	"net"
	"os"
	"strings"

	//"task_ex/internal/interceptor"

	"google.golang.org/grpc"

	"task_ex/internal/database"
	"task_ex/internal/handler"

	"task_ex/internal/auth"
	redisconfig "task_ex/internal/infrastructure/redis"
	"task_ex/internal/repository"
	"task_ex/internal/service"
	pb "task_ex/service/pb"
)

func main() {
	db, _ := database.NewMySQLDB()

	repo := repository.NewTaskRepository(db)
	svc := service.NewTaskService(repo)
	taskHandler := handler.NewTaskHandler(svc)

	userRepo := repository.NewUserRepository(db)
	redisClient, err := redisconfig.NewRedisClient()
	if err != nil {
		log.Printf("redis unavailable, continuing without cache/session store: %v", err)
	}

	authSvc := service.NewAuthService(userRepo, &auth.JWT{}, redisClient)
	userSvc := service.NewUserService(userRepo, &auth.JWT{})
	authHandler := handler.NewAuthHandler(authSvc, userSvc)

	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = ":9333"
	}
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf("unable to listen on %s: %v", port, err)
		log.Print("falling back to an auto-selected free TCP port")

		lis, err = net.Listen("tcp", ":0")
		if err != nil {
			log.Fatal(err)
		}
	}

	// auth jwt token
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(auth.AuthInterceptor),
	)
	pb.RegisterTaskServiceServer(grpcServer, taskHandler)
	pb.RegisterUserServiceServer(grpcServer, authHandler)

	// service gold price
	service.GetGoldPrice()

	// server port
	log.Println("gRPC server running on", lis.Addr().String())

	grpcServer.Serve(lis)

}
