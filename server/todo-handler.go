package server

import (
	"context"
	"go.ajitem.com/pb/todo"
	"google.golang.org/grpc/metadata"
)

type TodoHandler struct {
}

func (th *TodoHandler) CreateTodo(ctx context.Context, req *todo.CreateTodoRequest) (res *todo.CreateTodoResponse, e error) {
	headers, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		// Error. Could not parse headers
	}

	auth := headers.Get("authorization")

	if len(auth) == 0 {
		// Unauthorized
	}

	return &todo.CreateTodoResponse{}, nil
}

func (th *TodoHandler) ReadTodo(context.Context, *todo.ReadTodoRequest) (*todo.ReadTodoResponse, error) {

}
func (th *TodoHandler) ReadTodos(context.Context, *todo.ReadTodosRequest) (*todo.ReadTodosResponse, error) {

}
func (th *TodoHandler) UpdateTodo(context.Context, *todo.UpdateTodoRequest) (*todo.UpdateTodoResponse, error) {

}
func (th *TodoHandler) DeleteTodo(context.Context, *todo.DeleteTodoRequest) (*todo.DeleteTodoResponse, error) {

}
