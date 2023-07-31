package user

import (
	"log"
	"net/http"

	"github.com/lpgod/task/user"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Register user handlers
	user.RegisterHandlers(router)

	// Start HTTP server
	log.Println("Starting User Service...")
	http.ListenAndServe(":8081", router)
}
