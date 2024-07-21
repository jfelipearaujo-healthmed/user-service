package list_users

import (
	"errors"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/dtos/user_dto"
	list_users_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/list_users"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/http_response"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/validator"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/http/middlewares/role"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase list_users_contract.UseCase
}

func NewHandler(useCase list_users_contract.UseCase) *handler {
	return &handler{
		useCase: useCase,
	}
}

func (h *handler) Handle(c echo.Context) error {
	req := new(list_users_contract.Filter)

	if err := c.Bind(req); err != nil {
		return http_response.BadRequest(c, "unable to read request body", err)
	}

	if req.IsEmpty() {
		return http_response.BadRequest(c, "invalid request body", errors.New("empty queries, please provide at least one query"))
	}

	if err := validator.Validate(req); err != nil {
		return http_response.UnprocessableEntity(c, "invalid request body", err)
	}

	ctx := c.Request().Context()
	req.RoleFilter = c.Get("roleFilter").(role.Role)

	users, err := h.useCase.Execute(ctx, req)
	if err != nil {
		return http_response.HandleErr(c, err)
	}

	return http_response.Created(c, user_dto.MapFromSlice(users))
}
