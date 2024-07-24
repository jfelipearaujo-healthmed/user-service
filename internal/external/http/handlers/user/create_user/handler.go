package create_user

import (
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/dtos/user_dto"
	create_user_uc "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/create_user"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/hasher"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/http_response"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/validator"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase create_user_uc.UseCase
	hasher  hasher.Hasher
}

func NewHandler(useCase create_user_uc.UseCase, hasher hasher.Hasher) *handler {
	return &handler{
		useCase: useCase,
		hasher:  hasher,
	}
}

func (h *handler) Handle(c echo.Context) error {
	req := new(user_dto.CreateUserRequest)

	if err := c.Bind(req); err != nil {
		return http_response.BadRequest(c, "unable to read request body", err)
	}

	if err := validator.Validate(req); err != nil {
		return http_response.UnprocessableEntity(c, "invalid request body", err)
	}

	ctx := c.Request().Context()

	passHashed, err := h.hasher.HashPassword(ctx, req.Password)
	if err != nil {
		return http_response.BadRequest(c, "unable to hash the password", err)
	}

	req.Password = passHashed

	user, err := h.useCase.Execute(ctx, req)
	if err != nil {
		return http_response.HandleErr(c, err)
	}

	return http_response.Created(c, user_dto.MapFromDomain(user))
}
