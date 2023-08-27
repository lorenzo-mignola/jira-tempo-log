package forms

import (
	"fmt"
	"log"

	"github.com/lorenzo-mignola/jira-tempo-log/models"
	projectService "github.com/lorenzo-mignola/jira-tempo-log/service/project"
	"github.com/manifoldco/promptui"
)

func AddProject() {
	name := getName()
	description := getDescription()

	newProject := models.Project{
		Name:        name,
		Description: description,
	}

	projectCreated := projectService.Insert(&newProject)

	fmt.Printf("ID:%d\nProject: %s\nDescription:%s", projectCreated.ID, projectCreated.Name, projectCreated.Description)
}

func getName() string {
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
