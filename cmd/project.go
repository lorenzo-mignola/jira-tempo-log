package cmd

import (
	"fmt"
	"log"

	"github.com/lorenzo-mignola/jira-tempo-log/forms"
	"github.com/spf13/cobra"
)

func ProjectCommand() *cobra.Command {
	projectCommand := &cobra.Command{
		Use:   "project",
		Short: "Manage project",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("✨ Add new project ✨")
			addFlag, err := cmd.Flags().GetBool("add")
			if err != nil {
				log.Fatal(err)
			}

			if addFlag {
				forms.AddProject()
			}
		},
	}
	projectCommand.Flags().BoolP("add", "a", false, "Add project to the list of favorites projects")
	return projectCommand
}
