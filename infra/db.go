package infra

import (
	"database/sql"
	"os"

	"github.com/blezzio/triniti/utils"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPostgresConn() (*sql.DB, error) {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, utils.Trace(err, "failed to connect database")
	}
	return db, nil
}
