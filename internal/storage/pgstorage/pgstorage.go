package pgstorage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type PGstorage struct {
	db *pgxpool.Pool
}

func NewPGStorage(connString string) (*PGstorage, error) {

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, errors.Wrap(err, "config parcing error")
	}

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, errors.Wrap(err, "connection error")
	}
	storage := &PGstorage{
		db: db,
	}
	err = storage.initTables()
	if err != nil {
		return nil, err
	}

	return storage, nil
}

func (s *PGstorage) initTables() error {
	sql := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %v (
		%v VARCHAR(64) PRIMARY KEY,
		%v BIGINT NOT NULL DEFAULT 0,
		%v BIGINT NOT NULL DEFAULT 0,
		%v BIGINT NOT NULL DEFAULT 0
	)`,
		tableName,
		PlayerIDColumn,
		KillsColumn,
		DeathsColumn,
		ScoreColumn,
	)

	_, err := s.db.Exec(context.Background(), sql)
	if err != nil {
		return errors.Wrap(err, "init tables")
	}
	return nil
}
