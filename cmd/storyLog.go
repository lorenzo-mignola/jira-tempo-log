package cmd

import (
	"log"
	"time"

	"github.com/lorenzo-mignola/jira-tempo-log/forms"
	"github.com/lorenzo-mignola/jira-tempo-log/models"
	storyLogService "github.com/lorenzo-mignola/jira-tempo-log/service/storyLog"
	"github.com/lorenzo-mignola/jira-tempo-log/view"
	"github.com/spf13/cobra"
)

func LogCommand() *cobra.Command {
	logCommand := &cobra.Command{
		Use:   "log",
		Short: "Log on a project",
		Run:   logStory,
	}
	logCommand.Flags().BoolP("now", "n", false, "set today as date")

	// TODO flag week
	listCommand := &cobra.Command{
		Use:   "list",
		Short: "List all logs",
		Run:   printAllLogs,
	}
	listCommand.Flags().BoolP("plain", "p", false, "print plain text (non interactive mode)")

	logCommand.AddCommand(listCommand)

	return logCommand
}

func printAllLogs(cmd *cobra.Command, args []string) {
	plain, err := cmd.Flags().GetBool("plain")

	if err != nil {
		log.Fatal(err)
	}

	logs := storyLogService.GetAll()

	if plain {
		printPlain(logs)
		return
	}

	view.ViewStoryLog(logs)
}

func printPlain(logs []models.StoryLog) {
	for i := 0; i < len(logs); i++ {
		logs[i].Print()
	}
}

func logStory(cmd *cobra.Command, args []string) {
	var date time.Time
	nowFlag, err := cmd.Flags().GetBool("now")

	if err != nil {
		log.Fatal(err)
	}

	if nowFlag {
		date = time.Now()
	}

	forms.AddLog(date)
}
