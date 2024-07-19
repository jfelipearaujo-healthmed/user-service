package user_repository

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	user_repository_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/repositories/user"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/app_error"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/fields"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/persistence"
	"gorm.io/gorm"
)

type repository struct {
	dbService *persistence.DbService
}

func NewRepository(dbService *persistence.DbService) user_repository_contract.Repository {
	return &repository{
		dbService: dbService,
	}
}

func (rp *repository) GetByID(ctx context.Context, id uint) (*entities.User, error) {
	tx := rp.dbService.Instance.WithContext(ctx)

	user := new(entities.User)
	result := tx.Preload("Doctor").First(user, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, app_error.New(http.StatusNotFound, fmt.Sprintf("user with id %d not found", id))
		}

		return nil, result.Error
	}

	return user, result.Error
}

func (rp *repository) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	tx := rp.dbService.Instance.WithContext(ctx)

	user := new(entities.User)
	result := tx.Preload("Doctor").Where("email = ?", email).First(user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, app_error.New(http.StatusNotFound, fmt.Sprintf("user with email %s not found", email))
		}

		return nil, result.Error
	}

	return user, nil
}

func (rp *repository) GetByDocumentID(ctx context.Context, documentID string) (*entities.User, error) {
	tx := rp.dbService.Instance.WithContext(ctx)

	user := new(entities.User)
	result := tx.Preload("Doctor").Where("document_id = ?", documentID).First(user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, app_error.New(http.StatusNotFound, fmt.Sprintf("user with document id %s not found", documentID))
		}

		return nil, result.Error
	}

	return user, nil
}

func (rp *repository) GetByDocumentIDOrEmail(ctx context.Context, documentID string, email string) (*entities.User, error) {
	tx := rp.dbService.Instance.WithContext(ctx)

	user := new(entities.User)
	result := tx.Preload("Doctor").Where("document_id = ? OR email = ?", documentID, email).First(user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, app_error.New(http.StatusNotFound, fmt.Sprintf("user with document id %s or email %s not found", documentID, email))
		}

		return nil, result.Error
	}

	return user, nil
}

func (rp *repository) List(ctx context.Context, filter *user_repository_contract.ListUsersFilter) ([]entities.User, error) {
	tx := rp.dbService.Instance.WithContext(ctx)

	users := new([]entities.User)

	fields := fields.GetNonEmptyFields(filter, &fields.ANY_CHAR, &fields.ANY_CHAR)

	query := tx

	for field, value := range fields {
		query = query.Where(fmt.Sprintf("%s LIKE ?", field), value)
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}

	return *users, nil
}

func (rp *repository) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	tx := rp.dbService.Instance.WithContext(ctx)

	if err := tx.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (rp *repository) Update(ctx context.Context, user *entities.User) (*entities.User, error) {
	tx := rp.dbService.Instance.WithContext(ctx)

	result := tx.Model(user).Save(user)

	if err := result.Error; err != nil {
		return nil, err
	}

	if user.IsDoctor() {
		result = tx.Model(user.Doctor).Save(user.Doctor)
		if err := result.Error; err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (rp *repository) Delete(ctx context.Context, id uint) error {
	tx := rp.dbService.Instance.WithContext(ctx)

	result := tx.Delete(&entities.User{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
