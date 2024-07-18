package create_review

import (
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/dtos/review_dto"
	create_review_uc "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/review/create_review"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/http_response"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/validator"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase create_review_uc.UseCase
}

func NewHandler(useCase create_review_uc.UseCase) *handler {
	return &handler{
		useCase: useCase,
	}
}

func (h *handler) Handle(c echo.Context) error {
	req := new(review_dto.CreateReviewRequest)

	if err := c.Bind(req); err != nil {
		return http_response.BadRequest(c, "unable to read request body", err)
	}

	if err := validator.Validate(req); err != nil {
		return http_response.UnprocessableEntity(c, "invalid request body", err)
	}

	ctx := c.Request().Context()

	userId := c.Get("userId").(uint)

	review, err := h.useCase.Execute(ctx, userId, req)
	if err != nil {
		return http_response.HandleErr(c, err)
	}

	return http_response.Created(c, review_dto.MapFromDomain(review))
}
