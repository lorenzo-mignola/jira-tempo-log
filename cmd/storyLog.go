package cmd

import (
	"log"
	"time"

	"github.com/lorenzo-mignola/jira-tempo-log/forms"
	storyLogService "github.com/lorenzo-mignola/jira-tempo-log/service/storyLog"
	"github.com/spf13/cobra"
)

func LogCommand() *cobra.Command {
	logCommand := &cobra.Command{
		Use:   "log",
		Short: "Log on a project",
		Run:   logStory,
	}

	listCommand := &cobra.Command{
		Use:   "list",
		Short: "List all logs",
		Run:   printAllLogs,
	}

	logCommand.Flags().BoolP("now", "n", false, "set today as date")
	logCommand.AddCommand(listCommand)

	return logCommand
}

func printAllLogs(cmd *cobra.Command, args []string) {
	logs := storyLogService.GetAll()
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
