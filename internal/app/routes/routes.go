package routes

import (
	"../controllers"
	"github.com/gorilla/mux"
)

// GetRouter for the app
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/universe", controllers.SendUniverse).Methods("GET")
	r.HandleFunc("/player", controllers.PostNewPlayer).Methods("POST")
	r.HandleFunc("/player/{id}", controllers.PutPlayer).Methods("PUT")
	r.HandleFunc("/stats", controllers.GetStats).Methods("GET")

	return r
}
