package repo

import (
	"github.com/AurChatOrg/aurchat-server/internal/config"
	"github.com/AurChatOrg/aurchat-server/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Postgres *gorm.DB

// InitPostgres Init Postgres
func InitPostgres(cfg *config.Config) error {
	var err error
	Postgres, err = gorm.Open(postgres.Open(cfg.DSN.Postgres), &gorm.Config{})

	if err != nil {
		return err
	}

	err = Postgres.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}

	return nil
}
