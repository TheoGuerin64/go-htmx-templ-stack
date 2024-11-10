package database

import "github.com/uptrace/bun"

var models = []interface{}{
	(*User)(nil),
}

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID   int64 `bun:",pk,autoincrement"`
	Name string
}
