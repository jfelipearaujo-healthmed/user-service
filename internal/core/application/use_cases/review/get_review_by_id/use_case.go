package get_review_by_id_uc

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	get_review_by_id_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/review/get_review_by_id"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/app_error"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/persistence"
	"gorm.io/gorm"
)

type useCase struct {
	dbService *persistence.DbService
}

func NewUseCase(dbService *persistence.DbService) get_review_by_id_contract.UseCase {
	return &useCase{
		dbService: dbService,
	}
}

func (uc *useCase) Execute(ctx context.Context, userID uint, reviewID uint) (*entities.Review, error) {
	tx := uc.dbService.Instance.WithContext(ctx)

	user, err := uc.getUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	query := "patient_id = ? AND id = ?"

	if user.IsDoctor() {
		query = "doctor_id = ? AND id = ?"
	}

	review := new(entities.Review)
	if err := tx.Where(query, userID, reviewID).First(review).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, app_error.New(http.StatusNotFound, fmt.Sprintf("review with id %d not found", reviewID))
		}
	}

	return review, nil
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
