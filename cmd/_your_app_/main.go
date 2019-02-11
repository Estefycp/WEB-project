package main

import (
	"log"
	"net/http"

	"../../internal/app/routes"
)

func main() {
	r := routes.GetRouter()
	http.Handle("/", r)

	err := http.ListenAndServe("127.0.0.1"+":"+"8080", nil)
	if err != nil {
		log.Fatal("error en el servidor : ", err)
		return
	}
}
