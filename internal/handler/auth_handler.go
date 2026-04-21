package handler

import (
	"context"
	"fmt"

	"task_ex/internal/model"
	"task_ex/internal/service"
	pb "task_ex/service/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthHandler struct {
	pb.UnimplementedUserServiceServer
	authService *service.AuthService
	userService *service.UserService
}

func NewAuthHandler(authSvc *service.AuthService, userSvc *service.UserService) *AuthHandler {
	return &AuthHandler{authService: authSvc, userService: userSvc}
}

func (h *AuthHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	access, refresh, err := h.authService.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	return &pb.AuthResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (h *AuthHandler) RefreshToken(ctx context.Context, req *pb.RefreshRequest) (*pb.AuthResponse, error) {
	return h.authService.RefreshToken(ctx, req)
}

func (h *AuthHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := &model.User{
		Username: req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	if err := h.userService.CreateUser(ctx, user); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.CreateUserResponse{Id: int32(user.ID)}, nil
}

func (h *AuthHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := h.userService.GetUser(ctx, uint(req.Id))
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &pb.GetUserResponse{
		User: &pb.User{
			Id:        int32(user.ID),
			Name:      user.Username,
			Email:     user.Email,
			CreatedAt: fmt.Sprintf("%v", user.CreatedAt),
			UpdatedAt: fmt.Sprintf("%v", user.UpdatedAt),
		},
	}, nil
}

func (h *AuthHandler) ListUsers(ctx context.Context, _ *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users, err := h.userService.ListUsers(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var pbUsers []*pb.User
	for _, u := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:        int32(u.ID),
			Name:      u.Username,
			Email:     u.Email,
			CreatedAt: fmt.Sprintf("%v", u.CreatedAt),
			UpdatedAt: fmt.Sprintf("%v", u.UpdatedAt),
		})
	}
	return &pb.ListUsersResponse{Users: pbUsers}, nil
}
