package server

import (
	doctor_repository_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/repositories/doctor"
	user_repository_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/repositories/user"
	create_user_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/create_user"
	get_user_by_id_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/get_user_by_id"
	list_users_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/list_users"
	login_user_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/login_user"
	update_user_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/update_user"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/hasher"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/token"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/cache"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/persistence"
)

type Dependencies struct {
	Cache     cache.Cache
	DbService *persistence.DbService

	Hasher       hasher.Hasher
	TokenService token.TokenService

	UserRepository   user_repository_contract.Repository
	DoctorRepository doctor_repository_contract.Repository

	CreateUserUseCase  create_user_contract.UseCase
	GetUserByIdUseCase get_user_by_id_contract.UseCase
	UpdateUserUseCase  update_user_contract.UseCase
	ListUsersUseCase   list_users_contract.UseCase
	LoginUserUseCase   login_user_contract.UseCase
}
