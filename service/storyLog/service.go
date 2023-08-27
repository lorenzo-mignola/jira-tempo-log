package storylog

import (
	"github.com/lorenzo-mignola/jira-tempo-log/db"
	"github.com/lorenzo-mignola/jira-tempo-log/models"
)

func Insert(storyLog *models.StoryLog) models.StoryLog {
	db := db.GetDb()

	db.Create(&storyLog)

	return *storyLog
}
