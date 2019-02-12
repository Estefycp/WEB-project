package main

import (
	"net/http"

	"../../internal/app/routes"
	"github.com/rs/cors"
)

func main() {
	r := routes.GetRouter()
	http.Handle("/", r)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(r)
	http.ListenAndServe(":8081", handler)
}
