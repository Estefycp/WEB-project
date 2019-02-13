package main

import (
	"net/http"

	"../../internal/app/routes"
	"github.com/rs/cors"
)

func main() {
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

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT"},
	}).Handler(r)

	http.ListenAndServe(":8081", handler)
}
