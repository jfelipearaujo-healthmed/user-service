package rating_doctor_contract

import "context"

type UseCase interface {
	Execute(ctx context.Context, doctorID uint, rating float64) error
}
