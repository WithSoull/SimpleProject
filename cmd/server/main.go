package main

import (
    "log"
    "net"
    
    "google.golang.org/grpc"
    "github.com/malfoit/SimpleProject/internal/handler/user"
    userRepo "github.com/malfoit/SimpleProject/internal/repository/user"
    userService "github.com/malfoit/SimpleProject/internal/service/user"
    desc "github.com/malfoit/SimpleProject/pkg/user/v1"
)

func main() {
    repo := userRepo.NewRepository()
    
    svc := userService.NewService(repo)
    
    hnd := user.NewHandler(svc)
    
    grpcServer := grpc.NewServer()
    desc.RegisterUserV1Server(grpcServer, hnd)
    
    listener, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    
    log.Println("Server started on :50051")
    if err = grpcServer.Serve(listener); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}