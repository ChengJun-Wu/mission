package models

import "gorm.io/gorm"

type Server struct {
	gorm.Model
	Name string
	Ip string
	RsaKey string
}
