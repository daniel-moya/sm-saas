package main

import (
	"log"
	"user-service/internal/config"
	"user-service/internal/database"
	"user-service/internal/repositories"
	"user-service/internal/services"
	"user-service/internal/transport/grpc"
)

func main() {

    // Load configuration
    cfg := config.LoadConfig()

    // Initialize database connection
    db, err := database.NewDatabase(cfg)
    if err != nil {
        log.Fatalf("could not connect to the database: %v", err)
    }
    defer db.Close()

    // Initialize repository and service
    userRepo := repositories.NewUserRepository(db)
    userService := services.NewUserService(&userRepo)

    // Start gRPC server
    grpc.StartGRPCServer(userService, cfg.APPPort)
}
