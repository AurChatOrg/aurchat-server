package repo

import (
	"sync"

	"github.com/AurChatOrg/aurchat-server/internal/config"
	"github.com/AurChatOrg/aurchat-server/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Postgres *gorm.DB
	initOnce sync.Once
	initErr  error
)

// InitPostgres Init Postgres
func InitPostgres(cfg *config.Config) error {
	initOnce.Do(func() {
		db, err := gorm.Open(postgres.Open(cfg.DSN.Postgres), &gorm.Config{})
		if err != nil {
			initErr = err
			return
		}
		if err = db.AutoMigrate(&model.User{}); err != nil {
			initErr = err
			return
		}

		postgresDB, _ := db.DB()
		postgresDB.SetMaxOpenConns(100)
		postgresDB.SetMaxIdleConns(10)
		Postgres = db
	})
	return initErr
}
