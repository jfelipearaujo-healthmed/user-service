package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/application/use_cases/user/create_user"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/application/use_cases/user/delete_user"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/application/use_cases/user/get_user_by_id"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/application/use_cases/user/list_users"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/application/use_cases/user/update_user"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	list_users_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/list_users"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/config"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/persistence"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/secret"
	"github.com/joho/godotenv"
)

func init() {
	var err error
	time.Local, err = time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()

	if err := godotenv.Load("../../.env"); err != nil {
		slog.ErrorContext(ctx, "error loading .env file", "error", err)
		panic(err)
	}

	config, err := config.LoadFromEnv(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "error loading config from env", "error", err)
		panic(err)
	}

	cloudConfig, err := awsConfig.LoadDefaultConfig(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "error getting aws config", "error", err)
		panic(err)
	}

	if config.CloudConfig.IsBaseEndpointSet() {
		cloudConfig.BaseEndpoint = aws.String(config.CloudConfig.BaseEndpoint)
	}

	secretService := secret.NewService(cloudConfig)

	dbUrl, err := secretService.GetSecret(ctx, config.DbConfig.UrlSecretName)
	if err != nil {
		slog.ErrorContext(ctx, "error getting secret", "secret_name", config.DbConfig.UrlSecretName, "error", err)
		panic(err)
	}

	config.DbConfig.Url = dbUrl

	database := persistence.NewDatabase()

	if err := database.Connect(config); err != nil {
		slog.ErrorContext(ctx, "error connecting to database", "error", err)
		panic(err)
	}

	slog.InfoContext(ctx, "database connected")

	createUser := create_user.NewUseCase(database)
	updateUser := update_user.NewUseCase(database)
	getUser := get_user_by_id.NewUseCase(database)
	listUsers := list_users.NewUseCase(database)
	deleteUser := delete_user.NewUseCase(database)

	data := &entities.User{
		FullName:   "João Felipe Araujo",
		Email:      "joao.araujo@healthmed.com",
		Password:   "123456",
		DocumentID: "123456",
		Phone:      "1234567890",
		Role:       "doctor",
	}

	if err := createUser.Execute(ctx, data); err != nil {
		slog.ErrorContext(ctx, "error creating user", "error", err)
		panic(err)
	}

	slog.InfoContext(ctx, "user created")

	updateData := &entities.User{
		FullName: "João Araujo",
		Phone:    "00000000",
	}

	if err := updateUser.Execute(ctx, data.ID, updateData); err != nil {
		slog.ErrorContext(ctx, "error updating user", "error", err)
		panic(err)
	}

	slog.InfoContext(ctx, "user updated")

	user, err := getUser.Execute(ctx, data.ID)
	if err != nil {
		slog.ErrorContext(ctx, "error getting user", "error", err)
		panic(err)
	}

	slog.InfoContext(ctx, "user retrieved", "user", user)

	users, err := listUsers.Execute(ctx, &list_users_contract.Filter{
		FullName: "Araujo",
	})
	if err != nil {
		slog.ErrorContext(ctx, "error listing users", "error", err)
		panic(err)
	}

	slog.InfoContext(ctx, "users listed", "users", users)

	if err := deleteUser.Execute(ctx, data.ID); err != nil {
		slog.ErrorContext(ctx, "error deleting user", "error", err)
		panic(err)
	}

	slog.InfoContext(ctx, "user deleted")

	user, err = getUser.Execute(ctx, data.ID)
	if err != nil {
		slog.ErrorContext(ctx, "error getting deleted user", "error", err)
		panic(err)
	}

	slog.InfoContext(ctx, "user found", "user", user)
}
