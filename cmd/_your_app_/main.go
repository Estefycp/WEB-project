package main

import (
	"net/http"
	"os"

	"github.com/Estefycp/WEB-project/internal/app/routes"
	"github.com/Estefycp/WEB-project/internal/app/routines"
	"github.com/Estefycp/WEB-project/internal/app/storage"
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

	http.ListenAndServe(":"+os.Getenv("PORT"), handler)
}
