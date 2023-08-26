package service

import (
	"github.com/lorenzo-mignola/jira-tempo-log/db"
	"github.com/lorenzo-mignola/jira-tempo-log/models"
)

func Insert(project *models.Project) models.Project {
	db := db.GetDb()

	db.Create(&project)

	return GetByName(project.Name)
}

func GetByName(name string) models.Project {
	var createdProject models.Project
	db.GetDb().Where("name = ?", name).First(&createdProject)
	return createdProject
}
