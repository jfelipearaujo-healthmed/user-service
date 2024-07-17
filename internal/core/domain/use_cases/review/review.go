package review

import (
	"context"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
)

type CreateReview interface {
	Execute(ctx context.Context, doctorID string, avgRating float64, totalPatients int) (entities.Review, error)
}

type GetReviewByID interface {
	Execute(ctx context.Context, id string) (entities.Review, error)
}

type ListReviews interface {
	Execute(ctx context.Context, userID string) ([]entities.Review, error)
}
