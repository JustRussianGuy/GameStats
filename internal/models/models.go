package models

import "time"

// Событие убийства игрока

type GameEvent struct {
	KillerID   uint64
	VictimID   uint64
	OccurredAt time.Time
}

// Агрегированная статистика игрока
type PlayerStats struct {
	PlayerID string
	Kills    uint64
	Deaths   uint64
	Score    int64 // kills - deaths
}
