package server

import (
	"go.ajitem.com/pb/todo"
	"go.ajitem.com/pb/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Run() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// Register User Service
	user.RegisterUserServiceServer(s, &UserHandler{Allowed: []string{
		"Register", "Login", "ForgotPassword",
	}})
	// Register Todo Service
	todo.RegisterTodoServiceServer(s, &TodoHandler{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
