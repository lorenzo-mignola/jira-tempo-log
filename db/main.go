package db

import (
	"github.com/lorenzo-mignola/jira-tempo-log/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func Migrate() {
	db := GetDb()
	db.AutoMigrate(&models.Project{})
	db.AutoMigrate(&models.StoryLog{})
}
