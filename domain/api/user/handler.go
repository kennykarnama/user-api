package user

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"user-api/domain/api/shared"
	"user-api/domain/service/user"
)

type registerUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type registerUserResponse struct {
	ID                   int    `json:"id"`
	UUID                 string `json:"uuid"`
	shared.ErrorResponse `json:"errorResponse,omitempty"`
}

type Handler struct {
	userService user.Service
	validate    *validator.Validate
}

func NewHandler(userService user.Service, validate *validator.Validate) *Handler {
	return &Handler{
		userService: userService,
		validate:    validate,
	}
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	return
}
