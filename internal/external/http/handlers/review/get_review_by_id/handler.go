package get_review_by_id

import (
	"strconv"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/dtos/review_dto"
	get_review_by_id_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/review/get_review_by_id"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/http_response"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase get_review_by_id_contract.UseCase
}

func NewHandler(useCase get_review_by_id_contract.UseCase) *handler {
	return &handler{
		useCase: useCase,
	}
}

func (h *handler) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	userId := c.Get("userId").(uint)

	reviewId, err := strconv.ParseUint(c.Param("reviewId"), 10, 64)
	if err != nil {
		return http_response.BadRequest(c, "invalid review id", err)
	}

	review, err := h.useCase.Execute(ctx, userId, uint(reviewId))
	if err != nil {
		return http_response.HandleErr(c, err)
	}

	return http_response.OK(c, review_dto.MapFromDomain(review))
}
