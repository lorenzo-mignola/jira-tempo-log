package models

import (
	"fmt"

	"github.com/ttacon/chalk"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Description string
}

func (project Project) Print() {
	projectName := chalk.Green.Color(project.Name)
	description := chalk.Blue.Color(project.Description)

	fmt.Printf("🚀 Project: %s - Description: %s 🚀\n", projectName, description)
}
