package delete_user_uc

import (
	"context"

	user_repository_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/repositories/user"
	delete_user_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/delete_user"
)

type useCase struct {
	repository user_repository_contract.Repository
}

func NewUseCase(repository user_repository_contract.Repository) delete_user_contract.UseCase {
	return &useCase{
		repository: repository,
	}
}

func (uc *useCase) Execute(ctx context.Context, id uint) error {
	return uc.repository.Delete(ctx, id)
}
