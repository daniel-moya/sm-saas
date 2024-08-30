package grpc

import (
	"fmt"
	"log"
	"net"
	pb "user-service/internal/proto"
	"user-service/internal/services"

	"google.golang.org/grpc"
)

func StartGRPCServer(service services.UserService, port string) {
    lis, err := net.Listen("tcp", fmt.Sprint(":%v", port))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterUserServiceServer(s, NewUserHandler(service))

    log.Printf("gRPC server is running on port %v...\n", port)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

