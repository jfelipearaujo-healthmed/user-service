package rating_doctor

import (
	"strconv"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/dtos/doctor_dto"
	rating_doctor_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/doctor/rating_doctor"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/http_response"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase rating_doctor_contract.UseCase
}

func NewHandler(useCase rating_doctor_contract.UseCase) *handler {
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

	req := new(doctor_dto.RatingDoctor)
	if err := c.Bind(req); err != nil {
		return http_response.BadRequest(c, "invalid request body", err)
	}

	if err := h.useCase.Execute(ctx, uint(parsedDoctorId), req.Rating); err != nil {
		return http_response.HandleErr(c, err)
	}

	return http_response.OK(c, map[string]interface{}{
		"message": "doctor rating updated successfully",
	})
}
