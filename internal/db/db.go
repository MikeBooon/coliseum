// Package db contains database interactions
package db

import (
	"github.com/MikeBooon/coliseum/internal/config"
	"github.com/MikeBooon/coliseum/internal/db/dao"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate(c *config.Config) error {
	db, err := gorm.Open(postgres.Open(c.DBConn), &gorm.Config{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(dao.Models...)
	if err != nil {
		return err
	}
	return nil
}
