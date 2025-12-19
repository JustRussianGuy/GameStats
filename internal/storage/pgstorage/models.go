package pgstorage

type PlayerStats struct {
	PlayerID string `db:"player_id"`
	Kills    int64  `db:"kills"`
	Deaths   int64  `db:"deaths"`
	Score    int64  `db:"score"`
}

const (
	tableName = "player_stats"

	PlayerIDColumn = "player_id"
	KillsColumn    = "kills"
	DeathsColumn   = "deaths"
	ScoreColumn    = "score"
)

