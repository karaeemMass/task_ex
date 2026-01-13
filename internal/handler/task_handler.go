package handler

import (
	"context"
	"time"

	"task_ex/internal/model"
	"task_ex/internal/service"
	pb "task_ex/service/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TaskHandler struct {
	pb.UnimplementedTaskServiceServer
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) CreateTask(ctx context.Context, in *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	// set a timeout for the request context
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	task := &model.Task{
		Title:       in.GetTitle(),
		Description: in.GetDescription(),
	}

	if err := h.service.Create(ctx, task); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.CreateTaskResponse{Id: int32(task.ID)}, nil
}

func (h *TaskHandler) ListTasks(ctx context.Context, _ *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	// set a timeout for the request context
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	tasks, err := h.service.List(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var pbTasks []*pb.Task
	for _, t := range tasks {
		pbTasks = append(pbTasks, &pb.Task{
			Id:          int32(t.ID),
			Title:       t.Title,
			Description: t.Description,
			Completed:   t.Completed,
		})
	}

	return &pb.ListTasksResponse{Tasks: pbTasks}, nil
}

func (h *TaskHandler) FindTasks(ctx context.Context, in *pb.FindTasksRequest) (*pb.FindTasksResponse, error) {
	// set a timeout for the request context
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	task, err := h.service.FindByID(ctx, uint(in.GetId()))
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &pb.FindTasksResponse{
		Tasks: []*pb.Task{
			{
				Id:          int32(task.ID),
				Title:       task.Title,
				Description: task.Description,
				Completed:   task.Completed,
			},
		},
	}, nil
}
