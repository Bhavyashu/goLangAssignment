// routes.go
package routes

import (
	"fmt"
	"app/controllers"
	"app/models"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

// SetupRoutes configures the main routes for the application.
func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/process-single", sequentialSortingHandler).Methods("POST")
	r.HandleFunc("/process-concurrent", concurrentSortingHandler).Methods("POST")
}

// SetupOtherRoutes can be used to add additional routes if needed.
func SetupOtherRoutes(r *mux.Router) {
	// Add other routes here if needed
}

// sequentialSortingHandler handles the request for sequential sorting.
func sequentialSortingHandler(w http.ResponseWriter, r *http.Request) {
	handleSortingRequest(w, r, controllers.SequentialSorting)
}

// concurrentSortingHandler handles the request for concurrent sorting.
func concurrentSortingHandler(w http.ResponseWriter, r *http.Request) {
	handleSortingRequest(w, r, controllers.ConcurrentSorting)
}

// handleSortingRequest decodes the JSON payload, invokes the sorting function,
// encodes the response, and sends it to the client.
func handleSortingRequest(w http.ResponseWriter, r *http.Request, sortingFunc func(*models.Payload) *models.Response) {
	var payload models.Payload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Invoke the sorting function with the decoded payload
	response := sortingFunc(&payload)

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")

	// Encode the response and send it to the client
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	// Log that the response was sent successfully
	fmt.Println("Response sent successfully.\n")
}
