package dto

import (
	userv1 "github.com/2group/2rent.core-service/pkg/gen/go/user"
)

type UserDTO struct {
	ID           string  `json:"id"`
	UIN          *string `json:"uin"`
	Email        *string `json:"email"`
	PhoneNumber  *string `json:"phone_number"`
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
	AvatarURL    *string `json:"avatar_url"`
	Verified     bool    `json:"verified"`
	Role         *string `json:"role"`
	Organization *string `json:"organization_id"`
	CreatedAt    *string `json:"created_at"`
	UpdatedAt    *string `json:"updated_at"`
}

func MapUserToDTO(user *userv1.User) *UserDTO {
	return &UserDTO{
		ID:           user.Id,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		AvatarURL:    user.AvatarUrl,
		Verified:     user.Verified,
		Role:         user.Role,
		Organization: user.OrganizationId,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
