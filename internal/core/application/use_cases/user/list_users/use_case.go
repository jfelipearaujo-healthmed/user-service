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
	repoFilter := ToRepositoryFilter(filter)

	return uc.repository.List(ctx, repoFilter)
}

func ToRepositoryFilter(filter *list_users_contract.Filter) *user_repository_contract.ListFilter {
	return &user_repository_contract.ListFilter{
		DocumentID: filter.DocumentID,
		Email:      filter.Email,
		FullName:   filter.FullName,
		Phone:      filter.Phone,
		MedicalID:  filter.MedicalID,
		Specialty:  filter.Specialty,
		AvgRating:  filter.AvgRating,
		City:       filter.City,
		State:      filter.State,
		Zip:        filter.Zip,
		Role:       filter.RoleFilter,
	}
}
