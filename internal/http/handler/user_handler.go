package handler

import (
	"fmt"
	"log"
	"net/http"

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

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *UserHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var req userv1.RegisterRequest
	json.ParseJSON(r, &req)

	response, err := h.client.Api.Register(r.Context(), &req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
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

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *UserHandler) HandleUpdateProfile(w http.ResponseWriter, r *http.Request) {
	var req userv1.UpdateUserRequest
	json.ParseJSON(r, &req)
	user_id, ok := auth.GetUserID(r)
	if !ok {
		json.WriteError(w, http.StatusInternalServerError, fmt.Errorf("invalid token"))
		return
	}
	req.UserId = user_id

	response, err := h.client.Api.UpdateUser(r.Context(), &req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}
