package handler

import (
	"net/http"

	"github.com/2group/2rent.core-service/internal/grpc"
	organizationv1 "github.com/2group/2rent.core-service/pkg/gen/go/organization"
	"github.com/2group/2rent.core-service/pkg/json"
)

type OrganizationHandler struct {
	client *grpc.OrganizationClient
}

func NewOrganizationHandler(client *grpc.OrganizationClient) *OrganizationHandler{
    return &OrganizationHandler{client: client}
}

func(h *OrganizationHandler) HanleCreateOrganization(w http.ResponseWriter, r *http.Request) {
    var req *organizationv1.CreateOrganizationRequest
    json.ParseJSON(r, &req)

    response, err := h.client.Api.CreateOrganization(r.Context(), req)
    if err != nil {
        json.WriteError(w, http.StatusInternalServerError, err)
        return
    }
    _ = response

    json.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Organization created succesfully"})
    return
}
