package get_review_by_id_contract

import (
	"context"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
)

type UseCase interface {
	Execute(ctx context.Context, userID uint, reviewID uint) (*entities.Review, error)
}
