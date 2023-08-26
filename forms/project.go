package forms

import (
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
)

func AddProject() {
	project := getProject()
	description := getDescription()

	fmt.Printf("Project: %s\nDescription:%s", project, description)
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
