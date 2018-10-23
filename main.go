package main

import (
	"go.ajitem.com/grpc-poc-todo/client"
	"go.ajitem.com/grpc-poc-todo/server"
)

func main() {
	go func() {
		server.Run()
	}()
	client.Run()
}
