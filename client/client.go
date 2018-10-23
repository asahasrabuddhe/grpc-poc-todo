package client

import (
	"context"
	"go.ajitem.com/pb/todo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

func Run() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := todo.NewTodoServiceClient(conn)

	// header
	header := metadata.New(map[string]string{
		//"authorization": "aabbccdd",
		"space":  "spspspsp",
		"org": "ajitem",
		"limit": "10",
		"offset": "0",
	})
	// this is the critical step that includes your headers
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	r, err := c.CreateTodo(ctx, &todo.CreateTodoRequest{
		Todo: &todo.Todo{
			Id: 1,
			Title:"Buy Vegetables",
			Description: "Tomato, Potato, Onion",
		},
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.GetTodoId())
}
