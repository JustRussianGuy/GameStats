package models

import "time"

// Событие убийства игрока
type GameEvent struct {
	KillerID   string
	VictimID  string
	OccurredAt time.Time
}

// Агрегированная статистика игрока
type PlayerStats struct {
	PlayerID string
	Kills    int64
	Deaths   int64
	Score    int64 // kills - deaths
}