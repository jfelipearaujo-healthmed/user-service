package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	rating_doctor_uc "github.com/jfelipearaujo-healthmed/user-service/internal/core/application/use_cases/doctor/rating_doctor"
	create_user_uc "github.com/jfelipearaujo-healthmed/user-service/internal/core/application/use_cases/user/create_user"
	get_user_by_id_uc "github.com/jfelipearaujo-healthmed/user-service/internal/core/application/use_cases/user/get_user_by_id"
	list_users_uc "github.com/jfelipearaujo-healthmed/user-service/internal/core/application/use_cases/user/list_users"
	login_user_uc "github.com/jfelipearaujo-healthmed/user-service/internal/core/application/use_cases/user/login_user"
	update_user_uc "github.com/jfelipearaujo-healthmed/user-service/internal/core/application/use_cases/user/update_user"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/config"
	doctor_repository "github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/repositories/doctor"
	user_repository "github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/repositories/user"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/hasher"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/cache"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/http/handlers/doctor/rating_doctor"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/http/handlers/health"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/http/handlers/user/create_user"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/http/handlers/user/get_doctor_by_id"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/http/handlers/user/get_user_by_id"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/http/handlers/user/list_users"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/http/handlers/user/login_user"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/http/handlers/user/update_user"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/http/middlewares/logger"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/http/middlewares/role"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/http/middlewares/token"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/persistence"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/secret"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	tokenService "github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/token"
)

type Server struct {
	Config *config.Config

	Dependencies
}

func NewServer(ctx context.Context, config *config.Config) (*Server, error) {
	cloudConfig, err := awsConfig.LoadDefaultConfig(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "error getting aws config", "error", err)
		return nil, err
	}

	if config.CloudConfig.IsBaseEndpointSet() {
		cloudConfig.BaseEndpoint = aws.String(config.CloudConfig.BaseEndpoint)
	}

	secretService := secret.NewService(cloudConfig)

	dbUrl, err := secretService.GetSecret(ctx, config.DbConfig.UrlSecretName)
	if err != nil {
		slog.ErrorContext(ctx, "error getting secret", "secret_name", config.DbConfig.UrlSecretName, "error", err)
		return nil, err
	}

	cacheUrl, err := secretService.GetSecret(ctx, config.CacheConfig.HostSecretName)
	if err != nil {
		slog.ErrorContext(ctx, "error getting secret", "secret_name", config.CacheConfig.HostSecretName, "error", err)
		return nil, err
	}

	config.DbConfig.Url = dbUrl
	config.CacheConfig.Host = cacheUrl

	dbService := persistence.NewDbService()

	if err := dbService.Connect(config); err != nil {
		slog.ErrorContext(ctx, "error connecting to database", "error", err)
		return nil, err
	}

	cache := cache.NewRedisCache(ctx, config)

	userRepository := user_repository.NewRepository(cache, dbService)
	doctorRepository := doctor_repository.NewRepository(dbService)

	hasher := hasher.NewHasher()
	tokenService := tokenService.NewService(config)

	return &Server{
		Config: config,
		Dependencies: Dependencies{
			Cache:     cache,
			DbService: dbService,

			Hasher:       hasher,
			TokenService: tokenService,

			CreateUserUseCase:  create_user_uc.NewUseCase(userRepository),
			GetUserByIdUseCase: get_user_by_id_uc.NewUseCase(userRepository),
			UpdateUserUseCase:  update_user_uc.NewUseCase(cache, userRepository, doctorRepository),
			ListUsersUseCase:   list_users_uc.NewUseCase(userRepository),
			LoginUserUseCase:   login_user_uc.NewUseCase(userRepository, tokenService, hasher),

			RatingDoctorUseCase: rating_doctor_uc.NewUseCase(doctorRepository),
		},
	}, nil
}

func (s *Server) GetServer() *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf(":%d", s.Config.ApiConfig.Port),
		Handler:      s.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(logger.Middleware())
	e.Use(middleware.Recover())

	s.addHealthCheckRoutes(e)

	api := e.Group(fmt.Sprintf("/api/%s", s.Config.ApiConfig.ApiVersion))

	s.addUserRoutes(api)

	return e
}

func (server *Server) addHealthCheckRoutes(e *echo.Echo) {
	healthHandler := health.NewHandler(server.DbService)

	e.GET("/health", healthHandler.Handle)
}

func (server *Server) addUserRoutes(g *echo.Group) {
	userHandler := create_user.NewHandler(server.CreateUserUseCase, server.Hasher)
	getUserByIdHandler := get_user_by_id.NewHandler(server.GetUserByIdUseCase)
	getDoctorByIdHandler := get_doctor_by_id.NewHandler(server.GetUserByIdUseCase)
	updateUserHandler := update_user.NewHandler(server.UpdateUserUseCase)
	listUsersHandler := list_users.NewHandler(server.ListUsersUseCase)
	loginUserHandler := login_user.NewHandler(server.LoginUserUseCase)
	ratingDoctorHandler := rating_doctor.NewHandler(server.RatingDoctorUseCase)

	g.POST("/users", userHandler.Handle)
	g.POST("/users/login", loginUserHandler.Handle)

	g.GET("/users/me", getUserByIdHandler.Handle,
		token.Middleware(),
		role.MiddlewareAllowRole(role.Any))
	g.PUT("/users/me", updateUserHandler.Handle,
		token.Middleware(),
		role.MiddlewareAllowRole(role.Any))
	g.GET("/users/doctors", listUsersHandler.Handle,
		token.Middleware(),
		role.MiddlewareAllowRole(role.Patient),
		role.MiddlewareFilterRole(role.Doctor))
	g.GET("/users/doctors/:doctorId", getDoctorByIdHandler.Handle,
		token.Middleware(),
		role.MiddlewareAllowRole(role.Patient),
		role.MiddlewareFilterRole(role.Doctor))
	g.POST("/users/doctors/:doctorId/ratings", ratingDoctorHandler.Handle,
		token.Middleware(),
		role.MiddlewareAllowRole(role.Patient))
}
