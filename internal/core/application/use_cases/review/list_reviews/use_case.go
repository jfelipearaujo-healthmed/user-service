package list_reviews_uc

import (
	"context"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	review_repository_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/repositories/review"
	user_repository_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/repositories/user"
	list_reviews_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/review/list_reviews"
)

type useCase struct {
	reviewRepository review_repository_contract.Repository
	userRepository   user_repository_contract.Repository
}

func NewUseCase(
	reviewRepository review_repository_contract.Repository,
	userRepository user_repository_contract.Repository,
) list_reviews_contract.UseCase {
	return &useCase{
		reviewRepository: reviewRepository,
		userRepository:   userRepository,
	}
}

func (uc *useCase) Execute(ctx context.Context, userID uint) ([]entities.Review, error) {
	user, err := uc.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return uc.reviewRepository.List(ctx, userID, user.IsDoctor())
}
