package gateway

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/lpgod/task/proto"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
)

// APIService represents the API gateway service
type APIService struct {
	authClient     proto.AuthClient
	userClient     proto.UsersClient
	todoListClient proto.TodoListClient
}

// NewAPIService creates a new instance of the API gateway service
func NewAPIService() (*APIService, error) {
	authConn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to auth service: %v", err)
	}
	authClient := proto.NewAuthClient(authConn)

	userConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to user service: %v", err)
	}
	userClient := proto.NewUsersClient(userConn)

	todoListConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to todo list service: %v", err)
	}
	todoListClient := proto.NewTodoListClient(todoListConn)

	return &APIService{
		authClient:     authClient,
		userClient:     userClient,
		todoListClient: todoListClient,
	}, nil
}

// ServeHTTP handles incoming HTTP requests and directs them to the relevant gRPC service
func (api *APIService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the HTTP request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	// Parse the incoming request into a proto.Message
	var request proto.Message
	// TODO: Implement code to unmarshal the request body into the appropriate proto.Message
	// For example:
	// err = protojson.Unmarshal(body, &request)
	// if err != nil {
	//     http.Error(w, "Failed to parse request", http.StatusBadRequest)
	//     return
	// }

	// Extract the authentication token from the request headers
	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		http.Error(w, "Authentication token missing", http.StatusUnauthorized)
		return
	}

	// Validate the authentication token and get the user claims
	userClaims, err := api.validateAuthToken(authToken)
	if err != nil {
		http.Error(w, "Invalid authentication token", http.StatusUnauthorized)
		return
	}

	// TODO: Based on the request and user claims, decide which gRPC service and method to call
	if strings.HasPrefix(r.URL.Path, "/user") {
		api.handleUserRequest(w, r, request, userClaims)
	} else if strings.HasPrefix(r.URL.Path, "/todo") {
		api.handleTodoListRequest(w, r, request, userClaims)
	} else {
		http.Error(w, "Invalid API path", http.StatusNotFound)
	}
}

// validateAuthToken validates the authentication token with the auth service
func (api *APIService) validateAuthToken(token string) (jwt.MapClaims, error) {
	authResponse, err := api.authClient.ValidateToken(context.Background(), &proto.ValidateTokenRequest{Token: token})
	if err != nil {
		return nil, err
	}
	if !authResponse.IsValid {
		return nil, errors.New("invalid authentication token")
	}
	return authResponse.Claims, nil
}

// handleUserRequest handles user-related requests by forwarding them to the user service
func (api *APIService) handleUserRequest(w http.ResponseWriter, r *http.Request, request proto.Message, userClaims jwt.MapClaims) {
	// Get the user ID from the user claims
	userID, ok := userClaims["id"].(float64)
	if !ok {
		http.Error(w, "Invalid user ID in claims", http.StatusUnauthorized)
		return
	}

	switch {
	case r.URL.Path == "/user/create":
		// TODO: Implement code to create a new user using api.userClient
		// Example:
		// newUser := &user.User{
		//     Name:     "John Doe",
		//     Email:    "john@example.com",
		//     Password: "password123",
		// }
		// createdUser, err := api.userClient.CreateUser(context.Background(), newUser)
		// if err != nil {
		//     http.Error(w, "Failed to create user", http.StatusInternalServerError)
		//     return
		// }
		// jsonResponse(w, createdUser)

	case r.URL.Path == "/user/get":
		// TODO: Implement code to get user information using api.userClient
		// Example:
		// userID := int64(userID)
		// getUserResponse, err := api.userClient.GetUser(context.Background(), &proto.GetUserRequest{UserID: userID})
		// if err != nil {
		//     http.Error(w, "Failed to get user", http.StatusInternalServerError)
		//     return
		// }
		// jsonResponse(w, getUserResponse)

	// Add more cases for other user-related operations

	default:
		http.Error(w, "Invalid API path for user service", http.StatusNotFound)
	}
}

// handleTodoListRequest handles todo list-related requests by forwarding them to the todo list service
func (api *APIService) handleTodoListRequest(w http.ResponseWriter, r *http.Request, request proto.Message, userClaims jwt.MapClaims) {
	// Get the user ID from the user claims
	userID, ok := userClaims["id"].(float64)
	if !ok {
		http.Error(w, "Invalid user ID in claims", http.StatusUnauthorized)
		return
	}

	switch {
	case r.URL.Path == "/todo/create":
		// TODO: Implement code to create a new todo list using api.todoListClient
		newTodoList := &proto.TodoList{
			Name:        "Shopping List",
			Description: "Buy groceries and household items",
			UserId:      int64(userID),
		}
		createdTodoList, err := api.todoListClient.CreateTodoList(context.Background(), newTodoList)
		if err != nil {
			http.Error(w, "Failed to create todo list", http.StatusInternalServerError)
			return
		}
		jsonResponse(w, createdTodoList)

	case r.URL.Path == "/todo/get":
		// TODO: Implement code to get a todo list using api.todoListClient
		todoListID := int64(1) // Assuming the ID of the requested todo list
		getTodoListResponse, err := api.todoListClient.GetTodoList(context.Background(), &proto.GetTodoListRequest{TodoListId: todoListID})
		if err != nil {
			http.Error(w, "Failed to get todo list", http.StatusInternalServerError)
			return
		}
		jsonResponse(w, getTodoListResponse)

	// Add more cases for other todo list-related operations

	default:
		http.Error(w, "Invalid API path for todo list service", http.StatusNotFound)
	}
}

// jsonResponse writes the given data as a JSON response to the HTTP writer
func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	err := enc.Encode(data)
	if err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}
