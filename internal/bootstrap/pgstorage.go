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
		"postgres://%s:%s@%s:%d/%s",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName, // берём имя базы из config.yaml
	)

	storage, err := pgstorage.NewPGStorage(connectionString)
	if err != nil {
		log.Panicf("ошибка инициализации БД: %v", err)
	}

	return storage
}
