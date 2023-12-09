package main

import (
	"goLangAssignment/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Create a new Gorilla Mux router
	router := mux.NewRouter()

	// Setup application routes
	routes.SetupRoutes(router)

	// Add a test endpoint directly under the main router
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "API working, hello from arraySorting repository"}`))
	}).Methods("GET")

	// Use http.Handle("/", router) to handle all routes with the router
	http.Handle("/", router)


	// If there is no error, the server has started successfully
	log.Printf("Server is starting on port: 8080...\n")

	// Start the server on the specified port
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
	    log.Fatal("Error starting server:", err)
	}

}
