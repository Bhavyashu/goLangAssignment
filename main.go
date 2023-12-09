package main

import (
	"goLangAssignment/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	routes.SetupRoutes(router)

	// Add a new route for the '/test' endpoint directly under the main router
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "API working, hello from arraySorting repository"}`))
	}).Methods("GET")

	// Get the port from the environment variable, defaulting to 8080 if not set
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Print a message to indicate that the server is starting
	log.Printf("Server is starting at http://localhost:%s...\n", port)
	
	// Use http.Handle("/", router) to handle all routes with the router
	http.Handle("/", router)

	// Start the server on  specified port
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
