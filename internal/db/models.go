package db

import "gorm.io/gorm"

var Models = []any{&User{}}

type User struct {
	gorm.Model
	Email string
}
