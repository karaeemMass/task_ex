package handler

import (
	_ "context"
	"task_ex/service/pb"
)

type UsersHandler struct {
	pb.UnimplementedUserServiceServer
}
