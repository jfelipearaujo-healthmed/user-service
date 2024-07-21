package user_repository_contract

import (
	"context"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/http/middlewares/role"
)

type ListFilter struct {
	DocumentID *string
	Email      *string
	FullName   *string
	Phone      *string

	MedicalID *string
	Specialty *string
	AvgRating *float64

	Role role.Role
}

type Repository interface {
	GetByID(ctx context.Context, userID uint, roleFilter role.Role) (*entities.User, error)
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
	GetByDocumentID(ctx context.Context, documentID string) (*entities.User, error)
	GetByDocumentIDOrEmail(ctx context.Context, documentID string, email string) (*entities.User, error)
	List(ctx context.Context, filter *ListFilter) ([]entities.User, error)
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) (*entities.User, error)
	Delete(ctx context.Context, id uint) error
}
