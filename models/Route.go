package models

import "gorm.io/gorm"

type Route struct {
	gorm.Model
	Method string
	Name string
	Desc string
}