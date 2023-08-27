package main

import (
	"fmt"
	"os"

	"github.com/lorenzo-mignola/jira-tempo-log/cmd"
	"github.com/lorenzo-mignola/jira-tempo-log/db"
	"github.com/spf13/cobra"
)

func main() {
	db.Migrate()

	app := &cobra.Command{
		Use:   "jtl",
		Short: "Jira Tempo Log",
		Long:  "Log the time spent working on Jira issue",
	}

	app.AddCommand(cmd.ProjectCommand())
	app.AddCommand(cmd.LogCommand())

	err := app.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
