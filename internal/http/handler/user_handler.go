package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/2group/2rent.core-service/internal/dto"
	"github.com/2group/2rent.core-service/internal/grpc"
	auth "github.com/2group/2rent.core-service/internal/http/middleware"
	userv1 "github.com/2group/2rent.core-service/pkg/gen/go/user"
	"github.com/2group/2rent.core-service/pkg/json"
)

type UserHandler struct {
	client *grpc.UserClient
}

func NewUserHandler(client *grpc.UserClient) *UserHandler {
	return &UserHandler{client: client}
}

func (h *UserHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var req userv1.LoginRequest
	json.ParseJSON(r, &req)
	response, err := h.client.Api.Login(r.Context(), &req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	userDTO := dto.MapUserToDTO(response.User)

	loginResponse := map[string]interface{}{
		"token": response.Token,
		"user":  userDTO,
	}

	json.WriteJSON(w, http.StatusOK, loginResponse)
	return
}

func (h *UserHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON request
	var req userv1.RegisterRequest
	if err := json.ParseJSON(r, &req); err != nil {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request format: %w", err))
		return
	}

	// Call the gRPC Register service
	response, err := h.client.Api.Register(r.Context(), &req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to register user: %w", err))
		return
	}

	// Map the gRPC User object to a UserDTO
	userDTO := dto.MapUserToDTO(response.User)

	// Construct the final response with the token and the UserDTO
	registerResponse := map[string]interface{}{
		"token": response.Token,
		"user":  userDTO,
	}

	// Write the JSON response
	json.WriteJSON(w, http.StatusOK, registerResponse)
}

func (h *UserHandler) HandleGetProfile(w http.ResponseWriter, r *http.Request) {
	var req userv1.GetUserRequest
	user_id, ok := auth.GetUserID(r)
	if !ok {
		json.WriteError(w, http.StatusInternalServerError, fmt.Errorf("invalid token"))
		return
	}
	req.UserId = user_id
	log.Printf("user_id: %s", user_id)

	response, err := h.client.Api.GetUser(r.Context(), &req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	userDTO := dto.MapUserToDTO(response.User)

	json.WriteJSON(w, http.StatusOK, map[string]*dto.UserDTO{"user": userDTO})
	return
}

func (h *UserHandler) HandleUpdateProfile(w http.ResponseWriter, r *http.Request) {
	var req userv1.UpdateUserRequest
	json.ParseJSON(r, &req)

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
	return
}
