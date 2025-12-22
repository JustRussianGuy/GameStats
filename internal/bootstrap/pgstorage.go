package bootstrap

import (
	"fmt"
	"log"

	"github.com/JustRussianGuy/GameStats/config"
	"github.com/JustRussianGuy/GameStats/internal/storage/pgstorage"
)

func InitPGStorage(cfg *config.Config) *pgstorage.PGstorage {

	// Формируем строку подключения в формате URI
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.DBName, // <- вот здесь
		cfg.Database.SSLMode,
	)

	storage, err := pgstorage.NewPGStorage(connectionString)
	if err != nil {
		log.Panicf("ошибка инициализации БД: %v", err)
	}

	return storage
}
