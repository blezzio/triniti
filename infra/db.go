package infra

import (
	"context"
	"database/sql"
	"os"
	"time"

	"github.com/blezzio/triniti/utils"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPostgresConn() (*sql.DB, error) {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, utils.Trace(err, "failed to open db")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, utils.Trace(err, "failed to ping db")
	}

	return db, nil
}
