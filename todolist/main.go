package todolist

import (
	"fmt"
	"log"
	"net"

	"github.com/lpgod/task/proto"
	"github.com/lpgod/task/todolist"

	"google.golang.org/grpc"
)

func main() {
	port := "50052" // Change this port to your desired port number

	// Create a TCP listener to listen for incoming connections
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	server := grpc.NewServer()

	// Initialize the todo list service
	todoListService := &todolist.TodoListService{} // Assuming you have implemented this service

	// Register the todo list service with the gRPC server
	proto.RegisterTodoListServer(server, todoListService)

	fmt.Println("Todo list service is now running on port:", port)

	// Start serving incoming connections
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
