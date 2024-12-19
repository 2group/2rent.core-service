package handler

import (
	"fmt"
	"net/http"

	"github.com/2group/2rent.core-service/internal/grpc"
	auth "github.com/2group/2rent.core-service/internal/http/middleware"
	organizationv1 "github.com/2group/2rent.core-service/pkg/gen/go/organization"
	"github.com/2group/2rent.core-service/pkg/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type OrganizationHandler struct {
	client *grpc.OrganizationClient
}

func NewOrganizationHandler(client *grpc.OrganizationClient) *OrganizationHandler {
	return &OrganizationHandler{client: client}
}

func (h *OrganizationHandler) HandleCreateOrganization(w http.ResponseWriter, r *http.Request) {
	var req *organizationv1.CreateOrganizationRequest
	user_id, ok := auth.GetUserID(r)
	if !ok {
		json.WriteError(w, http.StatusInternalServerError, fmt.Errorf("invalid token"))
		return
	}
	json.ParseJSON(r, &req)
	req.UserId = user_id

	response, err := h.client.Api.CreateOrganization(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	_ = response

	json.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Organization created succesfully"})
	return
}

func (h *OrganizationHandler) HandleGetOrganizationByID(w http.ResponseWriter, r *http.Request) {
	// Extract organization_id from URL
	organizationID := chi.URLParam(r, "organization_id")
	if organizationID == "" {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("organization_id is required"))
		return
	}

	// Create the gRPC request
	req := &organizationv1.GetOrganizationByIdRequest{Id: organizationID}

	// Call the gRPC service
	response, err := h.client.Api.GetOrganizationById(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to fetch organization: %w", err))
		return
	}

	// Return the response as JSON
	json.WriteJSON(w, http.StatusOK, response)
}

func (h *OrganizationHandler) HandleUpdateOrganizationByID(w http.ResponseWriter, r *http.Request) {
	orgID := chi.URLParam(r, "id")
	if orgID == "" {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("organization ID is required"))
		return
	}

	// Validate UUID format
	if _, err := uuid.Parse(orgID); err != nil {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid organization ID format"))
		return
	}

	var req organizationv1.UpdateOrganizationRequest
	if err := json.ParseJSON(r, &req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	req.Id = orgID // Set the ID from URL

	response, err := h.client.Api.UpdateOrganization(r.Context(), &req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *OrganizationHandler) HandlePatchOrganizationByID(w http.ResponseWriter, r *http.Request) {
	// Extract organization_id from URL
	organizationID := chi.URLParam(r, "organization_id")
	if organizationID == "" {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("organization_id is required"))
		return
	}

	// Parse the request body
	var req organizationv1.PatchOrganizationRequest
	if err := json.ParseJSON(r, &req); err != nil {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request format: %w", err))
		return
	}
	req.Id = organizationID

	// Call the gRPC service
	response, err := h.client.Api.PatchOrganization(r.Context(), &req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to patch organization: %w", err))
		return
	}

	// Return the response as JSON
	json.WriteJSON(w, http.StatusOK, response)
}

func (h *OrganizationHandler) HandleDeleteOrganizationByID(w http.ResponseWriter, r *http.Request) {
	// Extract organization_id from URL
	organizationID := chi.URLParam(r, "organization_id")
	if organizationID == "" {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("organization_id is required"))
		return
	}

	// Create the gRPC request
	req := &organizationv1.DeleteOrganizationRequest{Id: organizationID}

	// Call the gRPC service
	response, err := h.client.Api.DeleteOrganization(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to delete organization: %w", err))
		return
	}

	// Return the response as JSON
	json.WriteJSON(w, http.StatusOK, response)
}
