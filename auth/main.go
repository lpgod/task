package auth

import (
	"fmt"
	"log"
	"net"

	"github.com/lpgod/task/proto"

	"github.com/lpgod/task/auth"

	"google.golang.org/grpc"
)

func main() {
	port := "50053" // Change this port to your desired port number

	// Create a TCP listener to listen for incoming connections
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	server := grpc.NewServer()

	// Initialize the auth service
	authService := &auth.AuthService{} // Assuming you have implemented this service

	// Register the auth service with the gRPC server
	proto.RegisterAuthServer(server, authService)

	fmt.Println("Auth service is now running on port:", port)

	// Start serving incoming connections
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
