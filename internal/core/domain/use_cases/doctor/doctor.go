package doctor

import (
	"context"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
)

type CreateDoctor interface {
	Execute(ctx context.Context, doctor *entities.Doctor) error
}

type UpdateDoctor interface {
	Execute(ctx context.Context, doctor *entities.Doctor) error
}

type GetDoctorByID interface {
	Execute(ctx context.Context, id string) (entities.Doctor, error)
}

type ListDoctors interface {
	Execute(ctx context.Context) ([]entities.Doctor, error)
}

type DeleteDoctor interface {
	Execute(ctx context.Context, id string) error
}
