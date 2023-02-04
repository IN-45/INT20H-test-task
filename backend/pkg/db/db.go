package db

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func NewDB(cfg Config) (*bun.DB, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithAddr(cfg.Addr),
		pgdriver.WithUser(cfg.User),
		pgdriver.WithPassword(cfg.Password),
		pgdriver.WithDatabase(cfg.Database),
		pgdriver.WithInsecure(true),
	))
	db := bun.NewDB(sqldb, pgdialect.New())

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
