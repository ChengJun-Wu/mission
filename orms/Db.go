package orms

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("mission"), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
	}
	return db
}