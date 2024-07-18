package persistence

import (
	"log/slog"
	"time"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/config"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	slogGorm "github.com/orandin/slog-gorm"
)

type DbService struct {
	Instance *gorm.DB
}

func NewDbService() *DbService {
	return &DbService{}
}

func (db *DbService) Connect(config *config.Config) error {
	log := slog.Default()

	gormLogger := slogGorm.New(
		slogGorm.WithHandler(log.Handler()),
		slogGorm.WithIgnoreTrace(),
		slogGorm.WithSlowThreshold(500*time.Millisecond),
		slogGorm.SetLogLevel(slogGorm.ErrorLogType, logger.LOG_ERR),
		slogGorm.SetLogLevel(slogGorm.SlowQueryLogType, logger.LOG_NOTICE),
		slogGorm.SetLogLevel(slogGorm.DefaultLogType, logger.LOG_DEBUG),
	)

	conn, err := gorm.Open(postgres.Open(config.DbConfig.Url), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return err
	}

	if err := conn.AutoMigrate(
		&entities.User{},
		&entities.Doctor{},
		&entities.Review{},
	); err != nil {
		return err
	}

	db.Instance = conn

	return nil
}
