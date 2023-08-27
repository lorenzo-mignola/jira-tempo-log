package models

import (
	"fmt"
	"time"

	"github.com/ttacon/chalk"
	"gorm.io/gorm"
)

type StoryLog struct {
	gorm.Model
	StoryNumber int `gorm:"not null"`
	Date        time.Time
	Description string
	ProjectID   uint
	Project     Project
}

func (storyLog StoryLog) getFullStoryName() string {
	return fmt.Sprintf("%s-%d", storyLog.Project.Name, storyLog.StoryNumber)
}

func (storyLog StoryLog) getStringDate() string {
	return fmt.Sprintf("%d.%d.%d", storyLog.Date.Day(), storyLog.Date.Month(), storyLog.Date.Year())
}

func (storyLog StoryLog) Print() {
	storyName := chalk.Green.Color(storyLog.getFullStoryName())
	date := chalk.Blue.Color(storyLog.getStringDate())

	fmt.Printf("🔨 Story: %s - Date: %s - Description: %s 🔨\n", storyName, date, storyLog.Description)
}
