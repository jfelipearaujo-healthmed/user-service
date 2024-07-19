package get_review_by_id_uc

import (
	"context"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	review_repository_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/repositories/review"
	user_repository_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/repositories/user"
	get_review_by_id_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/review/get_review_by_id"
)

type useCase struct {
	reviewRepository review_repository_contract.Repository
	userRepository   user_repository_contract.Repository
}

func NewUseCase(
	reviewRepository review_repository_contract.Repository,
	userRepository user_repository_contract.Repository,
) get_review_by_id_contract.UseCase {
	return &useCase{
		reviewRepository: reviewRepository,
		userRepository:   userRepository,
	}
}

func (uc *useCase) Execute(ctx context.Context, userID uint, reviewID uint) (*entities.Review, error) {
	user, err := uc.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return uc.reviewRepository.GetByID(ctx, reviewID, userID, user.IsDoctor())
}
