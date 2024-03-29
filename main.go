package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	// Define CORS options
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*"}, // You can customize this based on your needs
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"}, // You can customize this based on your needs
		AllowCredentials: true,
		MaxAge:           300, // Maximum age for cache, in seconds
	}
	router := chi.NewRouter()
	v1 := chi.NewRouter()
	router.Use(cors.Handler(corsOptions))
	v1.Get("/hello", func(w http.ResponseWriter, r *http.Request) { respondWithJson(w, 200, "hello") })
	router.Mount("/v1", v1)
	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadHeaderTimeout: time.Minute,
	}
	log.Printf("serving server on port %v", port)

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())

}
