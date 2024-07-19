package list_users_uc

import (
	"context"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	user_repository_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/repositories/user"
	list_users_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/list_users"
)

type useCase struct {
	repository user_repository_contract.Repository
}

func NewUseCase(repository user_repository_contract.Repository) list_users_contract.UseCase {
	return &useCase{
		repository: repository,
	}
}

func (uc *useCase) Execute(ctx context.Context, filter *list_users_contract.Filter) ([]entities.User, error) {
	repoFilter := &user_repository_contract.ListUsersFilter{}

	if filter.DocumentID != nil {
		repoFilter.DocumentID = *filter.DocumentID
	}
	if filter.Email != nil {
		repoFilter.Email = *filter.Email
	}
	if filter.FullName != nil {
		repoFilter.FullName = *filter.FullName
	}
	if filter.Phone != nil {
		repoFilter.Phone = *filter.Phone
	}
	if filter.Role != nil {
		repoFilter.Role = *filter.Role
	}

	return uc.repository.List(ctx, repoFilter)
}
