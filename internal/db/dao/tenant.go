package dao

import "github.com/uptrace/bun"

type Tenant struct {
	BaseColumns
	bun.BaseModel
	Name string `bun:"name,notnull"`
}
