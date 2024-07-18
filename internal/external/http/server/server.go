package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	create_user_uc "github.com/jfelipearaujo-healthmed/user-service/internal/core/application/use_cases/user/create_user"
	get_user_by_id_uc "github.com/jfelipearaujo-healthmed/user-service/internal/core/application/use_cases/user/get_user_by_id"
	update_user_uc "github.com/jfelipearaujo-healthmed/user-service/internal/core/application/use_cases/user/update_user"
	create_user_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/create_user"
	get_user_by_id_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/get_user_by_id"
	update_user_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/update_user"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/config"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/hasher"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/http/handlers/health"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/http/handlers/user/create_user"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/http/handlers/user/get_user_by_id"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/http/handlers/user/update_user"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/http/middlewares/logger"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/http/middlewares/token"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/persistence"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/secret"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Dependencies struct {
	DbService *persistence.DbService

	Hasher hasher.Hasher

	CreateUserUseCase  create_user_contract.UseCase
	GetUserByIdUseCase get_user_by_id_contract.UseCase
	UpdateUserUseCase  update_user_contract.UseCase
}

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

	config.DbConfig.Url = dbUrl

	dbService := persistence.NewDbService()

	if err := dbService.Connect(config); err != nil {
		slog.ErrorContext(ctx, "error connecting to database", "error", err)
		return nil, err
	}

	return &Server{
		Config: config,
		Dependencies: Dependencies{
			DbService: dbService,

			Hasher: hasher.NewHasher(),

			CreateUserUseCase:  create_user_uc.NewUseCase(dbService),
			GetUserByIdUseCase: get_user_by_id_uc.NewUseCase(dbService),
			UpdateUserUseCase:  update_user_uc.NewUseCase(dbService),
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

	authenticated := e.Group(fmt.Sprintf("/api/%s", s.Config.ApiConfig.ApiVersion))
	nonAuthenticated := e.Group(fmt.Sprintf("/api/%s", s.Config.ApiConfig.ApiVersion))

	authenticated.Use(token.Middleware())

	s.addUserAuthRoutes(nonAuthenticated)
	s.addUserRoutes(authenticated)

	return e
}

func (server *Server) addHealthCheckRoutes(e *echo.Echo) {
	healthHandler := health.NewHandler(server.DbService)

	e.GET("/health", healthHandler.Handle)
}

func (server *Server) addUserAuthRoutes(g *echo.Group) {
	userHandler := create_user.NewHandler(server.CreateUserUseCase, server.Hasher)

	g.POST("/users", userHandler.Handle)
}

func (server *Server) addUserRoutes(g *echo.Group) {
	getUserByIdHandler := get_user_by_id.NewHandler(server.GetUserByIdUseCase)
	updateUserHandler := update_user.NewHandler(server.UpdateUserUseCase)

	g.GET("/users/me", getUserByIdHandler.Handle)
	g.PUT("/users/me", updateUserHandler.Handle)
}
