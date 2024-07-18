package create_review_contract

import (
	"context"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/dtos/review_dto"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
)

type UseCase interface {
	Execute(ctx context.Context, patientID uint, request *review_dto.CreateReviewRequest) (*entities.Review, error)
}
