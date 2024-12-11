package handler

import (
	"net/http"

	"github.com/2group/2rent.core-service/internal/grpc"
	userv1 "github.com/2group/2rent.core-service/pkg/gen/go/user"
	"github.com/2group/2rent.core-service/pkg/json"
)

type UserHandler struct {
	client *grpc.Client
}

func NewUserHandler(client *grpc.Client) *UserHandler {
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
