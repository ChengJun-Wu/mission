package models

import "gorm.io/gorm"

type Config struct {
	gorm.Model
	Key string
	Value string
}
