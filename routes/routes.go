// routes/routes.go
package routes

import (
	"goLangAssignment/controllers"
	"goLangAssignment/models"
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/process-single", sequentialSortingHandler).Methods("POST")
	r.HandleFunc("/process-concurrent", concurrentSortingHandler).Methods("POST")
}

func SetupOtherRoutes(r *mux.Router) {
	// Add other routes here if needed
}

func sequentialSortingHandler(w http.ResponseWriter, r *http.Request) {
	handleSortingRequest(w, r, controllers.SequentialSorting)
}

func concurrentSortingHandler(w http.ResponseWriter, r *http.Request) {
	handleSortingRequest(w, r, controllers.ConcurrentSorting)
}

func handleSortingRequest(w http.ResponseWriter, r *http.Request, sortingFunc func(*models.Payload) *models.Response) {
	var payload models.Payload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	response := sortingFunc(&payload)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
	// Log the response sent sucessfully
	fmt.Println("Response sent successfully. \n");
}
