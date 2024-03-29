package repositories

import (
	"context"
	"database/sql"

	"github.com/blezzio/triniti/data/interfaces"
	"github.com/blezzio/triniti/services/dtos"
	"github.com/blezzio/triniti/utils"
)

type URL struct {
	db interfaces.DB
}

func NewURL(db interfaces.DB) *URL {
	return &URL{
		db: db,
	}
}

func (r *URL) WithTx(tx *sql.Tx) *URL {
	return &URL{db: tx}
}

const createQuery string = `INSERT INTO "url_lookup"("hash", "url") VALUES ($1, $2)`

func (r *URL) Create(ctx context.Context, param *dtos.CreateHash) error {
	_, err := r.db.ExecContext(ctx, createQuery, param.Hash, param.FullURL)
	return utils.Trace(err, "failed to insert %+v", param)
}

const getFullURLQuery string = `SELECT "url" FROM "url_lookup" WHERE "hash" = $1`

func (r *URL) GetFullURL(ctx context.Context, hash string) (string, error) {
	row := r.db.QueryRowContext(ctx, getFullURLQuery, hash)
	if row.Err() != nil {
		return "", utils.Trace(row.Err(), "failed to query for url of hash %s", hash)
	}

	var res string
	err := row.Scan(&res)
	return res, utils.Trace(err, "failed to scan row")
}

const deleteQuery string = `DELETE FROM "url_lookup" WHERE "hash" = $1 OR "url" = $2`

func (r *URL) Delete(ctx context.Context, hashOrFullURL string) error {
	_, err := r.db.ExecContext(ctx, deleteQuery, hashOrFullURL, hashOrFullURL)
	return utils.Trace(err, "failed to delete %s", hashOrFullURL)
}
