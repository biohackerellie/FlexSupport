package db

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"time"
)

type DB struct {
	*sqlx.DB
}

func NewDB(ctx context.Context, connectionString string) *DB {
	sqlDB := sqlx.MustOpen("pgx", connectionString)
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		panic(fmt.Sprintf("error connecting to database: %v", err))
	}
	return &DB{DB: sqlDB}
}
