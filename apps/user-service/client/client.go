package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "google.golang.org/grpc"
    "user-service/internal/proto"
)

func main() {
    // Set up a connection to the server.
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    client := proto.NewUserServiceClient(conn)

    // Prepare the login request
    loginRequest := &proto.LoginRequest{
        Email:    "test@example.com",
        Password: "password123",
    }

    // Set a timeout context
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
    defer cancel()

    // Call the Login method
    response, err := client.Login(ctx, loginRequest)
    if err != nil {
        log.Fatalf("could not log in: %v", err)
    }

    fmt.Printf("Login successful, received token: %s\n", response.Token)
}
