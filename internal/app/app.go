package app

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/2group/2rent.core-service/internal/grpc"
	"github.com/2group/2rent.core-service/internal/http/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type APIServer struct {
	address string
	log     *slog.Logger
}

func NewAPIServer(address string, log *slog.Logger) *APIServer {
	return &APIServer{address: address, log: log}
}

func (s *APIServer) Run() error {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.URLFormat)
	context := context.Background()

	client, err := grpc.NewClient(context, "localhost:50051", time.Hour*2, 2)
	if err != nil {
		panic(err)
	}

	userHandler := handler.NewUserHandler(client)

	router.Route("/api/v1", func(router chi.Router) {
		router.Route("/user", func(router chi.Router) {
			router.Post("/login", userHandler.HandleLogin)
		})
	})

	return http.ListenAndServe(s.address, router)
}
