package cmd

import (
	"context"
	"fmt"
	"github.com/Norsk-Tipping/gha-cli/pkg/client"
	"github.com/spf13/cobra"
	"os"
)

var (
	repositoryCmd = &cobra.Command{
		Use:   "repository",
		Short: "Display and manipulate GitHub repositories",
		Long:  "Display and manipulate GitHub repositories",

		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
			return initializeConfig(cmd)
		},
	}

	repositoryListCmd = &cobra.Command{
		Use:   "list",
		Short: "List GitHub repositories for organization",
		Long:  "List GitHub repositories for organization",
		Run: func(cmd *cobra.Command, args []string) {
			listRepositories()
		},
	}
)

func listRepositories() {
	// initialize github client
	ghClient = client.CreateGitHubClientWithPrivateKeyFile(gitHubAppId, gitHubInstallationId, gitHubPrivateKeyFile)
	ctx := context.Background()

	repos, _, err := ghClient.Repositories.List(ctx, gitHubOrganizationName, nil)
	if err != nil {
		fmt.Println("Got error while listing repositories:", err)
		os.Exit(1)
	}
	fmt.Println("Repositories for:", gitHubOrganizationName)
	for _, item := range repos {
		fmt.Printf("- %s\n", *item.Name)
	}

}

func init() {
	addRequiredFlags(repositoryCmd)
	repositoryCmd.AddCommand(repositoryListCmd)
	rootCmd.AddCommand(repositoryCmd)
}
