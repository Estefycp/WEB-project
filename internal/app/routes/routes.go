package routes

import (
	"log"
	"net/http"
	"../models"

	"github.com/gorilla/mux"
)

var datos []Item

func getUniverse(w http.ResponseWriter, r *http.Request) {
	universe = {}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(universe)
	
}

func createPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func updatePlayerw(w,http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datos[id])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/universe", getUniverse).Methods("GET")
	r.HandleFunc("/player", createPlayer).Methods("POST")
	r.HandleFunc("/player/{id}", updatePlayer).Methods("PUT")
	http.Handle("/", r)

	err := http.ListenAndServe("127.0.0.1"+":"+"8080", nil)
	if err != nil {
		log.Fatal("error en el servidor : ", err)
		return
	}
}
