package list_users_contract

import (
	"context"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
)

type Filter struct {
	DocumentID *string `json:"document_id"`
	Email      *string `json:"email"`
	FullName   *string `json:"full_name"`
	Phone      *string `json:"phone"`
	Role       *string `json:"role"`
}

type UseCase interface {
	Execute(ctx context.Context, filter *Filter) ([]entities.User, error)
}
