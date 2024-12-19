package handler

import (
	"fmt"
	"net/http"

	"github.com/2group/2rent.core-service/internal/dto"
	"github.com/2group/2rent.core-service/internal/grpc"
	auth "github.com/2group/2rent.core-service/internal/http/middleware"
	userv1 "github.com/2group/2rent.core-service/pkg/gen/go/user"
	"github.com/2group/2rent.core-service/pkg/json"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	client *grpc.UserClient
}

func NewUserHandler(client *grpc.UserClient) *UserHandler {
	return &UserHandler{client: client}
}

// HandleLogin handles user login
func (h *UserHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var req userv1.LoginRequest
	if err := json.ParseJSON(r, &req); err != nil {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request format: %w", err))
		return
	}

	response, err := h.client.Api.Login(r.Context(), &req)
	if err != nil {
		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("login failed: %w", err))
		return
	}

	userDTO := dto.MapUserToDTO(response.User)
	loginResponse := map[string]interface{}{
		"token": response.Token,
		"user":  userDTO,
	}

	json.WriteJSON(w, http.StatusOK, loginResponse)
}

// HandleRegister handles user registration
func (h *UserHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var req userv1.RegisterRequest
	if err := json.ParseJSON(r, &req); err != nil {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request format: %w", err))
		return
	}

	response, err := h.client.Api.Register(r.Context(), &req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to register user: %w", err))
		return
	}

	userDTO := dto.MapUserToDTO(response.User)
	registerResponse := map[string]interface{}{
		"token": response.Token,
		"user":  userDTO,
	}

	json.WriteJSON(w, http.StatusOK, registerResponse)
}

// HandleGetProfile fetches the authenticated user's profile
func (h *UserHandler) HandleGetProfile(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.GetUserID(r)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid token"))
		return
	}

	req := &userv1.GetUserRequest{UserId: userID}
	response, err := h.client.Api.GetUser(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to fetch user profile: %w", err))
		return
	}

	userDTO := dto.MapUserToDTO(response.User)
	json.WriteJSON(w, http.StatusOK, map[string]*dto.UserDTO{"user": userDTO})
}

// HandleUpdateProfile performs a full update on the authenticated user's profile
func (h *UserHandler) HandleUpdateProfile(w http.ResponseWriter, r *http.Request) {
	var req userv1.UpdateUserRequest
	if err := json.ParseJSON(r, &req); err != nil {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request format: %w", err))
		return
	}

	userID, ok := auth.GetUserID(r)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid token"))
		return
	}
	req.UserId = userID

	response, err := h.client.Api.UpdateUser(r.Context(), &req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to update user: %w", err))
		return
	}

	userDTO := dto.MapUserToDTO(response.User)
	json.WriteJSON(w, http.StatusOK, map[string]*dto.UserDTO{"user": userDTO})
}

// HandlePatchProfile performs a partial update on the authenticated user's profile
func (h *UserHandler) HandlePatchProfile(w http.ResponseWriter, r *http.Request) {
	var req userv1.PatchUserRequest
	if err := json.ParseJSON(r, &req); err != nil {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request format: %w", err))
		return
	}

	userID, ok := auth.GetUserID(r)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid token"))
		return
	}
	req.UserId = userID

	response, err := h.client.Api.PatchUser(r.Context(), &req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to patch user: %w", err))
		return
	}

	userDTO := dto.MapUserToDTO(response.User)
	json.WriteJSON(w, http.StatusOK, map[string]*dto.UserDTO{"user": userDTO})
}

// HandleDeleteProfile deletes the authenticated user's account
func (h *UserHandler) HandleDeleteProfile(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.GetUserID(r)
	if !ok {
		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid token"))
		return
	}

	req := &userv1.DeleteUserRequest{UserId: userID}
	response, err := h.client.Api.DeleteUser(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to delete user: %w", err))
		return
	}

	json.WriteJSON(w, http.StatusOK, map[string]bool{"success": response.Success})
}

func (h *UserHandler) HandleGetUserByID(w http.ResponseWriter, r *http.Request) {
	// Extract user_id from URL
	userID := chi.URLParam(r, "user_id")
	if userID == "" {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing user_id in URL"))
		return
	}

	// Create the gRPC request
	req := &userv1.GetUserRequest{UserId: userID}

	// Call the gRPC service
	response, err := h.client.Api.GetUser(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to fetch user: %w", err))
		return
	}

	// Map the gRPC User model to a UserDTO
	userDTO := dto.MapUserToDTO(response.User)
	json.WriteJSON(w, http.StatusOK, map[string]*dto.UserDTO{"user": userDTO})
}

func (h *UserHandler) HandleUpdateUserByID(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var req userv1.UpdateUserRequest
	if err := json.ParseJSON(r, &req); err != nil {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request format: %w", err))
		return
	}

	// Extract user_id from URL
	userID := chi.URLParam(r, "user_id")
	if userID == "" {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing user_id in URL"))
		return
	}
	req.UserId = userID

	// Call the gRPC service
	response, err := h.client.Api.UpdateUser(r.Context(), &req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to update user: %w", err))
		return
	}

	// Map the gRPC User model to a UserDTO
	userDTO := dto.MapUserToDTO(response.User)
	json.WriteJSON(w, http.StatusOK, map[string]*dto.UserDTO{"user": userDTO})
}

func (h *UserHandler) HandlePatchUserByID(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var req userv1.PatchUserRequest
	if err := json.ParseJSON(r, &req); err != nil {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request format: %w", err))
		return
	}

	// Extract user_id from URL
	userID := chi.URLParam(r, "user_id")
	if userID == "" {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing user_id in URL"))
		return
	}
	req.UserId = userID

	// Call the gRPC service
	response, err := h.client.Api.PatchUser(r.Context(), &req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to patch user: %w", err))
		return
	}

	// Map the gRPC User model to a UserDTO
	userDTO := dto.MapUserToDTO(response.User)
	json.WriteJSON(w, http.StatusOK, map[string]*dto.UserDTO{"user": userDTO})
}

func (h *UserHandler) HandleDeleteUserByID(w http.ResponseWriter, r *http.Request) {
	// Extract user_id from URL
	userID := chi.URLParam(r, "user_id")
	if userID == "" {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing user_id in URL"))
		return
	}

	// Create the gRPC request
	req := &userv1.DeleteUserRequest{UserId: userID}

	// Call the gRPC service
	response, err := h.client.Api.DeleteUser(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to delete user: %w", err))
		return
	}

	// Return success response
	json.WriteJSON(w, http.StatusOK, map[string]bool{"success": response.Success})
}
