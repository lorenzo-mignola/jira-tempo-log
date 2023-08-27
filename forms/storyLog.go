package forms

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/lorenzo-mignola/jira-tempo-log/models"
	projectService "github.com/lorenzo-mignola/jira-tempo-log/service/project"
	storyLogService "github.com/lorenzo-mignola/jira-tempo-log/service/storyLog"
	"github.com/manifoldco/promptui"
)

func AddLog(defaultDate time.Time) {
	project, err := selectProject()

	if err != nil {
		log.Fatal(err)
	}

	storyNumber := storyNumber()
	fmt.Printf("Story: %s-%d", project.Name, storyNumber)
	description := description()
	date := date(defaultDate)

	storyLog := models.StoryLog{
		StoryNumber: storyNumber,
		Description: description,
		Date:        date,
		ProjectID:   project.ID,
	}

	storyLogService.Insert(&storyLog)
}

func selectProject() (*models.Project, error) {
	projectsNames := projectService.GetProjectNames()

	prompt := promptui.Select{
		Label: "Select Project to log",
		Items: projectsNames,
	}

	_, projectSelected, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil, err
	}

	fmt.Printf("You choose %s\n", projectSelected)
	project := projectService.GetByName(projectSelected)
	return &project, nil
}

func storyNumber() int {
	prompt := promptui.Prompt{
		Label: "Story number",
		Validate: func(input string) error {
			if len(input) == 0 {
				return errors.New("story number can't be empty")
			}
			return nil
		},
	}

	storyNumberStr, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	storyNumber, err := strconv.Atoi(storyNumberStr)

	if err != nil {
		panic(err)
	}

	return int(storyNumber)
}

func description() string {
	prompt := promptui.Prompt{
		Label: "Description",
		Templates: &promptui.PromptTemplates{
			Valid: "📝 {{ . }}: ",
		},
	}

	description, err := prompt.Run()

	if err != nil {
		log.Fatal(err)
	}

	return description
}

// TODO ask the user to insert the date
func date(defaultDate time.Time) time.Time {
	if defaultDate.IsZero() {
		return defaultDate
	}

	return time.Now()
}
