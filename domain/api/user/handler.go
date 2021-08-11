package user

import (
	"context"
	"github.com/go-playground/validator/v10"
	"net/http"
	"user-api/domain/api/shared"
	userEntity "user-api/domain/models/user"
	"user-api/domain/service/user"
	"user-api/util"
)

type Handler struct {
	userService user.Service
	validate    *validator.Validate
	ctx         context.Context
}

func NewHandler(ctx context.Context, userService user.Service, validate *validator.Validate) *Handler {
	return &Handler{
		userService: userService,
		validate:    validate,
		ctx:         ctx,
	}
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req registerUserRequest

	err := util.DecodeToStruct(r.Body, &req)
	if err != nil {
		shared.ResponseJson(w, shared.ErrorResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	if err := h.validate.Struct(&req); err != nil {
		shared.ResponseJson(w, shared.ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	newUser := &userEntity.User{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}

	err = h.userService.RegisterUser(h.ctx, newUser)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		if err == user.ErrUserAlreadyExist {
			// ref: https://tools.ietf.org/html/rfc7231#section-4.3.3
			httpStatus = http.StatusSeeOther
		}
		shared.ResponseJson(w, shared.ErrorResponse{
			CustomErrorCode: "-",
			Message:         err.Error(),
		}, httpStatus)
		return
	}

	shared.ResponseJson(w, registerUserResponse{
		ID:   newUser.ID,
		UUID: newUser.UUID,
	}, http.StatusCreated)
	return
}
