package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"rest-wsgo/handlers"
	"rest-wsgo/server"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env files")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:       PORT,
		JWTSecret:  JWT_SECRET,
		DatbaseURL: DATABASE_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRouters)
}

func BindRouters(
	s server.Server,
	r *mux.Router,
) {
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
}
