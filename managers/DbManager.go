package managers

import (
	"mission/models"
	"mission/orms"
)

type DbManager struct {
}

func (m *DbManager) Boot() {
	orms.DB().AutoMigrate(
		&models.Config{},
		&models.User{},
		&models.Route{},
		&models.Server{},
	)
}

func NewDbManager() *DbManager {
	return new(DbManager)
}