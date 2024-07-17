package persistence

import (
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Instance *gorm.DB
}

func NewDatabase() *Database {
	return &Database{}
}

func (db *Database) Connect(config *config.Config) error {
	conn, err := gorm.Open(postgres.Open(config.DbConfig.Url), &gorm.Config{})
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
