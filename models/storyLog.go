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

func (storyLog StoryLog) Print() {
	storyName := chalk.Green.Color(storyLog.getFullStoryName())
	date := chalk.Blue.Color(storyLog.getDateString())
	durationStr := fmt.Sprintf("%.0fs", storyLog.Duration)
	duration := chalk.Red.Color(durationStr)

	fmt.Printf("🔨 Story: %s - Date: %s - Duration: %.0fs Description: %s 🔨\n", storyName, date, duration, storyLog.Description)
}
