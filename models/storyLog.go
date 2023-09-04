package models

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/table"
	"github.com/ttacon/chalk"
	"gorm.io/gorm"
)

type StoryLog struct {
	gorm.Model
	StoryNumber int `gorm:"not null"`
	Date        time.Time
	Description string
	Duration    float64 `gorm:"not null"`
	ProjectID   uint
	Project     Project
}

func (storyLog StoryLog) getFullStoryName() string {
	return fmt.Sprintf("%s-%d", storyLog.Project.Name, storyLog.StoryNumber)
}

func (storyLog StoryLog) getDateString() string {
	return storyLog.Date.Format("02.01.2006")
}

// TODO format duration
func (storyLog StoryLog) getDurationInSeconds() string {
	return fmt.Sprintf("%.0fs", storyLog.Duration)
}

func (storyLog StoryLog) Print() {
	storyName := chalk.Green.Color(storyLog.getFullStoryName())
	date := chalk.Blue.Color(storyLog.getDateString())
	duration := chalk.Red.Color(storyLog.getDurationInSeconds())

	fmt.Printf("🔨 Story: %s - Date: %s - Duration: %s Description: %s 🔨\n", storyName, date, duration, storyLog.Description)
}

func (storyLog StoryLog) ToTableRow() table.Row {
	return table.Row{
		storyLog.getFullStoryName(),
		storyLog.getDateString(),
		storyLog.getDurationInSeconds(),
		storyLog.Description,
	}
}
