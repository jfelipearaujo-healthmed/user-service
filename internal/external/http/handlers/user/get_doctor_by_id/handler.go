package get_doctor_by_id

import (
	"strconv"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/dtos/user_dto"
	get_user_by_id_uc "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/get_user_by_id"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/http_response"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/http/middlewares/role"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase get_user_by_id_uc.UseCase
}

func NewHandler(useCase get_user_by_id_uc.UseCase) *handler {
	return &handler{
		useCase: useCase,
	}
}

func (h *handler) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	doctorId := c.Param("doctorId")
	parsedDoctorId, err := strconv.ParseUint(doctorId, 10, 64)
	if err != nil {
		return http_response.BadRequest(c, "invalid doctor id", err)
	}

	roleFilter := c.Get("roleFilter").(role.Role)

	user, err := h.useCase.Execute(ctx, uint(parsedDoctorId), roleFilter)
	if err != nil {
		return http_response.HandleErr(c, err)
	}

	if user == nil {
		return http_response.NotFound(c)
	}

	return http_response.OK(c, user_dto.MapFromDomain(user))
}
