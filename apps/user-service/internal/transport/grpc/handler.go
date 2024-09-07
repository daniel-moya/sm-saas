package grpc

import (
	"context"
	pb "user-service/internal/proto"
	"user-service/internal/services"
)

type UserHandler struct {
    service services.UserService
    pb.UnimplementedUserServiceServer
}

func NewUserHandler(service services.UserService) *UserHandler {
    return &UserHandler{service: service}
}

func (h *UserHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
    user, err := h.service.RegisterUser(
        ctx,
        req.Email,
        req.Password,
    )

    if err != nil {
        return nil, err
    }

    return &pb.RegisterResponse{
        Id:       user.ID,
        Email:    user.Email,
    }, nil
}

func (h *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
    token, err := h.service.LoginUser(
        ctx,
        req.Email,
        req.Password,
    )

    if err != nil {
        return nil, err
    }

    return &pb.LoginResponse{Token: token}, nil
}

func (h *UserHandler) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
    err := h.service.Delete(
        ctx,
        req.Id,
    )

    if err != nil {
        return nil, err
    }

    return &pb.DeleteResponse{Message: "deleted" }, nil
}


