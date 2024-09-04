package grpc

import (
	"context"
	"fmt"
	pb "user-service/internal/proto"
	"user-service/internal/services"
)

type ProtectedHandler struct {
    service services.ProtectedService
    pb.UnimplementedProtectedServiceServer
}

func NewProtectedHandler(service services.ProtectedService) *ProtectedHandler {
    return &ProtectedHandler{service: service}
}


func (h *ProtectedHandler) Admin(ctx context.Context, req *pb.AdminRequest) (*pb.AdminResponse, error) {
    session, _ := h.service.Admin(ctx)
    fmt.Println("Admin", req.Filter)
    return &pb.AdminResponse{Session: session}, nil
}
