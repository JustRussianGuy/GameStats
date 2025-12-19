package pgstorage

import (
	"context"

	"github.com/JustRussianGuy/GameStats/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (s *PGstorage) GetPlayerStats(
	ctx context.Context,
	playerID string,
) (*models.PlayerStats, error) {

	query := squirrel.
		Select(PlayerIDColumn, KillsColumn, DeathsColumn, ScoreColumn).
		From(tableName).
		Where(squirrel.Eq{PlayerIDColumn: playerID}).
		PlaceholderFormat(squirrel.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "build get stats query")
	}

	row := s.db.QueryRow(ctx, sql, args...)

	var stats models.PlayerStats
	err = row.Scan(
		&stats.PlayerID,
		&stats.Kills,
		&stats.Deaths,
		&stats.Score,
	)
	if err != nil {
		return nil, errors.Wrap(err, "scan stats")
	}

	return &stats, nil
}

func (s *PGstorage) GetLeaderboard(
	ctx context.Context,
	limit int,
) ([]*models.PlayerStats, error) {

	query := squirrel.
		Select(PlayerIDColumn, KillsColumn, DeathsColumn, ScoreColumn).
		From(tableName).
		OrderBy("score DESC").
		Limit(uint64(limit)).
		PlaceholderFormat(squirrel.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "build leaderboard query")
	}

	rows, err := s.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "query leaderboard")
	}
	defer rows.Close()

	var res []*models.PlayerStats
	for rows.Next() {
		var ps models.PlayerStats
		if err := rows.Scan(
			&ps.PlayerID,
			&ps.Kills,
			&ps.Deaths,
			&ps.Score,
		); err != nil {
			return nil, errors.Wrap(err, "scan leaderboard")
		}
		res = append(res, &ps)
	}

	return res, nil
}
