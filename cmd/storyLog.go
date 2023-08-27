package cmd

import (
	"log"
	"time"

	"github.com/lorenzo-mignola/jira-tempo-log/forms"
	"github.com/spf13/cobra"
)

func LogCommand() *cobra.Command {

	logCommand := &cobra.Command{
		Use:   "log",
		Short: "Log on a project",
		Run:   logStory,
	}

	logCommand.Flags().BoolP("now", "n", false, "set today as date")

	return logCommand
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
