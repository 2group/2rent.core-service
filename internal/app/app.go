package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/2group/2rent.core-service/internal/config"
	"github.com/2group/2rent.core-service/internal/grpc"
	"github.com/2group/2rent.core-service/internal/http/handler"
	auth "github.com/2group/2rent.core-service/internal/http/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type APIServer struct {
	cfg *config.Config
	log *slog.Logger
}

func NewAPIServer(cfg *config.Config, log *slog.Logger) *APIServer {
	return &APIServer{cfg: cfg, log: log}
}

func (s *APIServer) Run() error {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.URLFormat)
	context := context.Background()

	userClient, err := grpc.NewUserClient(context, s.cfg.GRPC.User, time.Hour*2, 2)
	if err != nil {
		panic(err)
	}

	organizationClient, err := grpc.NewOrganizationClient(context, s.cfg.GRPC.Organization, time.Hour*2, 2)
	if err != nil {
		panic(err)
	}

	userHandler := handler.NewUserHandler(userClient)
	organizationHandler := handler.NewOrganizationHandler(organizationClient)

	router.Route("/api/v1", func(apiRouter chi.Router) {
		// User Routes
		apiRouter.Route("/user", func(userRouter chi.Router) {
			// Authentication and registration
			userRouter.Post("/login", userHandler.HandleLogin)
			userRouter.Post("/register", userHandler.HandleRegister)

			// Authenticated user operations
			userRouter.Group(func(authRouter chi.Router) {
				authRouter.Use(auth.AuthMiddleware)
				authRouter.Get("/profile", userHandler.HandleGetProfile)       // Get the authenticated user's profile
				authRouter.Put("/profile", userHandler.HandleUpdateProfile)    // Full update for authenticated user's profile
				authRouter.Patch("/profile", userHandler.HandlePatchProfile)   // Partial update for authenticated user's profile
				authRouter.Delete("/profile", userHandler.HandleDeleteProfile) // Delete the authenticated user's account
			})

			// General user management (admin actions)
			userRouter.Group(func(adminRouter chi.Router) {
				adminRouter.Use(auth.AuthMiddleware)                               // Middleware to ensure admin privileges
				adminRouter.Get("/{user_id}", userHandler.HandleGetUserByID)       // Get a specific user by ID
				adminRouter.Put("/{user_id}", userHandler.HandleUpdateUserByID)    // Full update for a specific user
				adminRouter.Patch("/{user_id}", userHandler.HandlePatchUserByID)   // Partial update for a specific user
				adminRouter.Delete("/{user_id}", userHandler.HandleDeleteUserByID) // Delete a specific user
			})
		})

		// Organization Routes
		apiRouter.Route("/organization", func(orgRouter chi.Router) {
			orgRouter.Group(func(authRouter chi.Router) {
				authRouter.Use(auth.AuthMiddleware)
				authRouter.Post("/", organizationHandler.HandleCreateOrganization)           // Create an organization
				authRouter.Get("/{id}", organizationHandler.HandleGetOrganizationByID)       // Get organization by ID
				authRouter.Put("/{id}", organizationHandler.HandleUpdateOrganizationByID)    // Full update for organization
				authRouter.Patch("/{id}", organizationHandler.HandlePatchOrganizationByID)   // Partial update for organization
				authRouter.Delete("/{id}", organizationHandler.HandleDeleteOrganizationByID) // Delete an organization
			})
		})
	})

	return http.ListenAndServe(fmt.Sprintf("localhost:%d", s.cfg.REST.Port), router)
}
