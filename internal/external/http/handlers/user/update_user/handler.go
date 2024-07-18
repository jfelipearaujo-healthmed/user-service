package update_user

import (
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/dtos/user_dto"
	update_user_uc "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/update_user"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/http_response"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/validator"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase update_user_uc.UseCase
}

func NewHandler(useCase update_user_uc.UseCase) *handler {
	return &handler{
		useCase: useCase,
	}
}

func (h *handler) Handle(c echo.Context) error {
	req := new(user_dto.UpdateUserRequest)

	if err := c.Bind(req); err != nil {
		return http_response.BadRequest(c, "unable to read request body", err)
	}

	if err := validator.Validate(req); err != nil {
		return http_response.UnprocessableEntity(c, "invalid request body", err)
	}

	ctx := c.Request().Context()

	userId := c.Get("userId").(uint)

	user, err := h.useCase.Execute(ctx, userId, req)
	if err != nil {
		return http_response.HandleErr(c, err)
	}

	return http_response.OK(c, user_dto.MapFromDomain(user))
}
