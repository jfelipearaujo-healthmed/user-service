package list_reviews

import (
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/dtos/review_dto"
	list_reviews_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/review/list_reviews"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/http_response"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase list_reviews_contract.UseCase
}

func NewHandler(useCase list_reviews_contract.UseCase) *handler {
	return &handler{
		useCase: useCase,
	}
}

func (h *handler) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	userId := c.Get("userId").(uint)

	reviews, err := h.useCase.Execute(ctx, userId)
	if err != nil {
		return http_response.HandleErr(c, err)
	}

	if len(reviews) == 0 {
		return http_response.NotFound(c)
	}

	return http_response.OK(c, review_dto.MapFromDomainSlice(reviews))
}
