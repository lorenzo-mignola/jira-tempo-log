package service

import (
	"github.com/lorenzo-mignola/jira-tempo-log/db"
	"github.com/lorenzo-mignola/jira-tempo-log/models"
)

func Insert(project *models.Project) models.Project {
	db := db.GetDb()

	db.Create(&project)

	return *project
}

func GetByName(name string) models.Project {
	var createdProject models.Project
	db.GetDb().Where("name = ?", name).First(&createdProject)
	return createdProject
}

func GetAllProjects() []models.Project {
	var projects []models.Project
	db.GetDb().Find(&projects)
	return projects
}

func GetProjectNames() []string {
	projects := GetAllProjects()
	var names []string
	for i := 0; i < len(projects); i++ {
		names = append(names, projects[i].Name)
	}

	return names
}
