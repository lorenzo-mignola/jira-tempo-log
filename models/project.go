package models

import "gorm.io/gorm"

type ProjectDTO struct {
	Name        string
	Description string
}

type Project struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Description string
}
