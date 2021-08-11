package user

import "user-api/domain/api/shared"

type registerUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type registerUserResponse struct {
	ID            int64                 `json:"id"`
	UUID          string                `json:"uuid"`
	ErrorResponse *shared.ErrorResponse `json:"errorResponse,omitempty"`
}
