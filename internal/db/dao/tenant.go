package dao

import "github.com/uptrace/bun"

type Tenant struct {
	Base
	bun.BaseModel
	Name string `bun:"name,notnull"`
}
