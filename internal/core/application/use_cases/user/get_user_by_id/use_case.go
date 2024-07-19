package get_user_by_id_uc

import (
	"context"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	user_repository_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/repositories/user"
	get_user_by_id_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/get_user_by_id"
)

type useCase struct {
	repository user_repository_contract.Repository
}

func NewUseCase(repository user_repository_contract.Repository) get_user_by_id_contract.UseCase {
	return &useCase{
		repository: repository,
	}
}

func (uc *useCase) Execute(ctx context.Context, id uint) (*entities.User, error) {
	return uc.repository.GetByID(ctx, id)
}
