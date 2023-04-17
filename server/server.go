package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"rest-wsgo/databases"
	"rest-wsgo/repository"

	"github.com/gorilla/mux"
)

type Config struct {
	Port       string
	JWTSecret  string
	DatbaseURL string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("Port is required")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("JWTSecret is required")
	}
	if config.DatbaseURL == "" {
		return nil, errors.New("DatbaseURL is required")
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}

	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	// Connecting to Database
	repo, err := databases.NewPostgresRepository(b.config.DatbaseURL)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	repository.SetRepository(repo) // inicializar la variable implementation
	log.Println("Starting server on port: ", b.Config().Port)
	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		log.Fatalf("server stopped")
	}
}
