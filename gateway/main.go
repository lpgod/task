package gateway

import (
	"fmt"
	"log"
	"net"

	"task/proto"

	"github.com/lpgod/task/apigateway"

	"google.golang.org/grpc"
)

func main() {
	port := "50054" // Change this port to your desired port number

	// Create a TCP listener to listen for incoming connections
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	server := grpc.NewServer()

	// Initialize the API gateway service
	apiGatewayService := &apigateway.APIService{} // Assuming you have implemented this service

	// Register the API gateway service with the gRPC server
	proto.RegisterApiGatewayServer(server, apiGatewayService)

	fmt.Println("API gateway service is now running on port:", port)

	// Start serving incoming connections
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
