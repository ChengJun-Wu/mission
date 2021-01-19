package managers

import (
	"errors"
	"gorm.io/gorm"
	"mission/helpers"
	"mission/models"
	"mission/orms"
)

type InitManager struct {
}

func (m *InitManager) Boot() {
	db := orms.DB()
	var config models.Config
	rs := db.Where("key = ?", "AesKey").First(&config)
	if rs.Error != nil && errors.Is(rs.Error, gorm.ErrRecordNotFound) {
		config = models.Config{
			Key: "AesKey",
			Value: helpers.RandomString(16),
		}
		db.Create(&config)
	}
	var (
		count int64
		user models.User
	)
	db.Model(&user).Count(&count)
	if count == 0 {
		user = models.User{
			Username: "admin",
			Password: helpers.PasswordHash("admin"),
		}
		db.Create(&user)
	}
}

func NewInitManager() *InitManager {
	return new(InitManager)
}