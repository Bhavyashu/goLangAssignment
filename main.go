package main

import (
	"app/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	// Update the home route to return an HTML response
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	    w.Header().Set("Content-Type", "text/html")
	    w.WriteHeader(http.StatusOK)
	    w.Write([]byte("<h1>Hello, World!</h1>"))
	}).Methods("GET")




	// Use CORS middleware to enable cross-origin requests
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},    // Allow any origin
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// Use http.Handle("/", corsHandler.Handler(router)) to handle all routes with the router and enable CORS
	http.Handle("/", corsHandler.Handler(router))

	// If there is no error, the server has started successfully
	log.Printf("Server is starting on port: 8080...\n")

	// Start the server on the specified port
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
	    log.Fatal("Error starting server:", err)
	}

}
