package login_user

import (
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/dtos/user_dto"
	login_user_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/login_user"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/http_response"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase login_user_contract.UseCase
}

func NewHandler(useCase login_user_contract.UseCase) *handler {
	return &handler{
		useCase: useCase,
	}
}

func (h *handler) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(user_dto.LoginUserRequest)
	if err := c.Bind(req); err != nil {
		return http_response.BadRequest(c, "invalid request body", err)
	}

	token, err := h.useCase.Execute(ctx, req)
	if err != nil {
		return http_response.HandleErr(c, err)
	}

	return http_response.Created(c, token)
}
