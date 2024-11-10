package database

import (
	"context"
	"database/sql"
	"project/internal/config"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func createModels(db *bun.DB) error {
	ctx := context.Background()
	for _, model := range models {
		if _, err := db.NewCreateTable().Model(model).IfNotExists().Exec(ctx); err != nil {
			return err
		}
	}

	return nil
}

func Open(settings *config.Settings) (*bun.DB, error) {
	connector := pgdriver.NewConnector(
		pgdriver.WithAddr("database:5432"),
		pgdriver.WithUser(settings.PostgresUser),
		pgdriver.WithPassword(settings.PostgresPassword),
		pgdriver.WithDatabase(settings.PostgresDb),
		pgdriver.WithInsecure(true),
	)

	sqlDB := sql.OpenDB(connector)
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	bunDB := bun.NewDB(sqlDB, pgdialect.New())
	if err := createModels(bunDB); err != nil {
		return nil, err
	}

	return bunDB, nil
}
