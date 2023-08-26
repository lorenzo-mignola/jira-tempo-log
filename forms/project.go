package forms

import (
	"fmt"
	"log"

	"github.com/lorenzo-mignola/jira-tempo-log/models"
	projectService "github.com/lorenzo-mignola/jira-tempo-log/service/project"
	"github.com/manifoldco/promptui"
)

func AddProject() {
	project := getProject()
	description := getDescription()

	newProject := models.Project{
		Name:        project,
		Description: description,
	}

	projectCreated := projectService.Insert(&newProject)

	fmt.Printf("ID:%d\nProject: %s\nDescription:%s", projectCreated.ID, projectCreated.Name, projectCreated.Description)
}

func getProject() string {
	prompt := promptui.Prompt{
		Label: "Project",
		Templates: &promptui.PromptTemplates{
			Valid: "⭐ {{ . }}: ",
		},
	}

	project, err := prompt.Run()

	if err != nil {
		log.Fatal(err)
	}

	return project
}

func getDescription() string {
	prompt := promptui.Prompt{
		Label: "Description",
		Templates: &promptui.PromptTemplates{
			Valid: "📝 {{ . }}: ",
		},
	}

	project, err := prompt.Run()

	if err != nil {
		log.Fatal(err)
	}

	return project
}
