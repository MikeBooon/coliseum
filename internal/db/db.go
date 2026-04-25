// Package db contains database interactions
package db

import (
	"github.com/MikeBooon/coliseum/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate(c *config.Config) error {
	db, err := gorm.Open(postgres.Open(c.DBConn), &gorm.Config{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(Models...)
	if err != nil {
		return err
	}
	return nil
}
