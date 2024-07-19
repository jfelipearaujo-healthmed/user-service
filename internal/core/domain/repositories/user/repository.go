package user_repository_contract

import (
	"context"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
)

type ListUsersFilter struct {
	DocumentID string
	Email      string
	FullName   string
	Phone      string
	Role       string
}

type Repository interface {
	GetByID(ctx context.Context, id uint) (*entities.User, error)
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
	GetByDocumentID(ctx context.Context, documentID string) (*entities.User, error)
	GetByDocumentIDOrEmail(ctx context.Context, documentID string, email string) (*entities.User, error)
	List(ctx context.Context, filter *ListUsersFilter) ([]entities.User, error)
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) (*entities.User, error)
	Delete(ctx context.Context, id uint) error
}
