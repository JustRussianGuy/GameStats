package pgstorage

import (
	"context"

	"github.com/JustRussianGuy/GameStats/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

func (s *PGstorage) IncrementKill(ctx context.Context, playerID string) error {
	query := squirrel.
		Insert(tableName).
		Columns(PlayerIDColumn, KillsColumn, DeathsColumn, ScoreColumn).
		Values(playerID, 1, 0, 1).
		Suffix(`
			ON CONFLICT (player_id)
			DO UPDATE SET
				kills = player_stats.kills + 1,
				score = player_stats.score + 1
		`).
		PlaceholderFormat(squirrel.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "build increment kill query")
	}

	_, err = s.db.Exec(ctx, sql, args...)
	return errors.Wrap(err, "exec increment kill")
}

func (s *PGstorage) IncrementDeath(ctx context.Context, playerID string) error {
	query := squirrel.
		Insert(tableName).
		Columns(PlayerIDColumn, KillsColumn, DeathsColumn, ScoreColumn).
		Values(playerID, 0, 1, -1).
		Suffix(`
			ON CONFLICT (player_id)
			DO UPDATE SET
				deaths = player_stats.deaths + 1,
				score  = player_stats.score - 1
		`).
		PlaceholderFormat(squirrel.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "build increment death query")
	}

	_, err = s.db.Exec(ctx, sql, args...)
	return errors.Wrap(err, "exec increment death")
}
