package cmd

import (
	"fmt"
	"log"

	"github.com/lorenzo-mignola/jira-tempo-log/forms"
	projectService "github.com/lorenzo-mignola/jira-tempo-log/service/project"
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

	add := &cobra.Command{
		Use:   "add",
		Short: "Add a project to favorites",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("✨ Add new project ✨")
			forms.AddProject()
		},
	}
	projectCommand.AddCommand(add)

	list := &cobra.Command{
		Use:   "list",
		Short: "Add a project to favorites",
		Run: func(cmd *cobra.Command, args []string) {
			projects := projectService.GetAllProjects()
			for i := 0; i < len(projects); i++ {
				projects[i].Print()
			}
		},
	}
	projectCommand.AddCommand(list)

	return projectCommand
}
