package main

import (
	"net/http"

	"../../internal/app/routes"
	"../../internal/app/routines"
	"../../internal/app/storage"
	"github.com/rs/cors"
)

func main() {
	storage.GetInstance()
	routines.StartRoutines()

	r := routes.GetRouter()
	http.Handle("/", r)

	// handler := handlers.CORS(
	// 	handlers.AllowedHeaders(
	// 		[]string{
	// 			"X-Requested-With",
	// 			"Content-Type",
	// 			"Authorization",
	// 		},
	// 	),
	// 	handlers.AllowedMethods(
	// 		[]string{
	// 			"GET",
	// 			"POST",
	// 			"PUT",
	// 			"HEAD",
	// 			"OPTIONS",
	// 		},
	// 	),
	// 	handlers.AllowedOrigins(
	// 		[]string{"*"},
	// 	),
	// )(r)

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT"},
	}).Handler(r)

	http.ListenAndServe(":8081", handler)
}
