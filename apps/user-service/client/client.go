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

    registerRequest := &proto.RegisterRequest{
        Email: "test@example.com",
        Password: "password123",
    }

    ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
    defer cancel()


    response, err := client.Register(ctx, registerRequest)
    if err != nil {
        log.Fatalf("could not register: %v", err)
    }

    fmt.Printf("Register successful, received token: %s\n", response.Email)

    // Prepare the login request
    loginRequest := &proto.LoginRequest{
        Email:    "test@example.com",
        Password: "password123",
    }

    // Set a timeout context
    ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
    defer cancel()

    // Call the Login method
    resp2, err2 := client.Login(ctx, loginRequest)
    if err2 != nil {
        log.Fatalf("could not log in: %v", err2)
    }

    fmt.Printf("Login successful, received token: %s\n", resp2.Token)
}
