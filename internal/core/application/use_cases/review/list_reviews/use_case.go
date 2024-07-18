package list_reviews_uc

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	list_reviews_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/review/list_reviews"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/app_error"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/persistence"
	"gorm.io/gorm"
)

type useCase struct {
	dbService *persistence.DbService
}

func NewUseCase(dbService *persistence.DbService) list_reviews_contract.UseCase {
	return &useCase{
		dbService: dbService,
	}
}

func (uc *useCase) Execute(ctx context.Context, userID uint) ([]entities.Review, error) {
	tx := uc.dbService.Instance.WithContext(ctx)

	user, err := uc.getUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	query := "patient_id = ?"

	if user.IsDoctor() {
		query = "doctor_id = ?"
	}

	reviews := new([]entities.Review)
	if err := tx.Where(query, userID).Order("created_at desc").Find(reviews).Error; err != nil {
		return nil, err
	}

	return *reviews, nil
}

func (uc *useCase) getUser(ctx context.Context, userID uint) (*entities.User, error) {
	tx := uc.dbService.Instance.WithContext(ctx)

	user := new(entities.User)
	if err := tx.Where("id = ?", userID).Preload("Doctor").First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, app_error.New(http.StatusNotFound, fmt.Sprintf("user with id %d not found", userID))
		}
	}

	return user, nil
}
