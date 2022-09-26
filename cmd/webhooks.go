package cmd

import (
	"context"
	"fmt"
	"github.com/Norsk-Tipping/gha-cli/pkg/client"
	"github.com/Norsk-Tipping/gha-cli/pkg/webhooks"
	"github.com/spf13/cobra"
	"os"
)

var (
	webhooksCmd = &cobra.Command{
		Use:   "webhooks",
		Short: "Display and manipulate GitHub repository webhooks",
		Long:  "Display and manipulate GitHub repositories",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
			return initializeConfig(cmd)
		},
	}

	webhooksListCmd = &cobra.Command{
		Use:   "list",
		Short: "List GitHub repository webhooks",
		Long:  "List GitHub repository webhooks",
		Run: func(cmd *cobra.Command, args []string) {
			listWebHooks()
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
			return initializeConfig(cmd)
		},
	}
	webhooksAddCmd = &cobra.Command{
		Use:   "add",
		Short: "Add GitHub repository webhook",
		Long:  "Add GitHub repository webhook",
		Run: func(cmd *cobra.Command, args []string) {
			addWebhook()
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
			return initializeConfig(cmd)
		},
	}

	webhooksDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete GitHub repository webhook",
		Long:  "Delete GitHub repository webhook",
		Run: func(cmd *cobra.Command, args []string) {
			deleteWebhook()
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
			return initializeConfig(cmd)
		},
	}

	repositoryName string
	webhookUrl     string
	webhookId      int64
)

func listWebHooks() {
	// Initialize GitHub client and context
	ghClient = client.CreateGitHubClientWithPrivateKeyFile(gitHubAppId, gitHubInstallationId, gitHubPrivateKeyFile)
	ctx := context.Background()

	hooks, err := webhooks.GetWebhooks(ctx, ghClient, gitHubOrganizationName, repositoryName)
	if err != nil {
		fmt.Println("Got error while listing webhooks:", err)
		os.Exit(1)
	}
	fmt.Printf("GitHub webhooks for repository %s:\n", repositoryName)
	for _, item := range hooks {
		fmt.Printf("- [%d] %s\n", *item.ID, item.Config["url"])
	}
}

func addWebhook() {
	// Initialize GitHub client and context
	ghClient = client.CreateGitHubClientWithPrivateKeyFile(gitHubAppId, gitHubInstallationId, gitHubPrivateKeyFile)
	ctx := context.Background()

	fmt.Printf("Adding webhook (%s) to repository: %s/%s\n", webhookUrl, gitHubOrganizationName, repositoryName)

	// FIXME: We need a list of events we want to receive webhooks for. Hardcoded to 'push' for now
	events := []string{"push"}

	// FIXME: We need to specify what kind of content the webhook receives. Hardcoded to 'json' for now
	contentType := "json"

	err := webhooks.AddWebhook(ctx, ghClient, gitHubOrganizationName, repositoryName, webhookUrl, contentType, events)
	if err != nil {
		fmt.Println("Got an error while adding webhook:", err)
		os.Exit(1)
	}
	fmt.Printf("Successfully added webhook (%s) to repository %s/%s\n", webhookUrl, gitHubOrganizationName, repositoryName)
}

func deleteWebhook() {
	// Initialize GitHub client and context
	ghClient = client.CreateGitHubClientWithPrivateKeyFile(gitHubAppId, gitHubInstallationId, gitHubPrivateKeyFile)
	ctx := context.Background()

	fmt.Printf("Deleting webhook with ID: %d from repository %s/%s\n", webhookId, gitHubOrganizationName, repositoryName)
	err := webhooks.DeleteWebhook(ctx, ghClient, gitHubOrganizationName, repositoryName, webhookId)
	if err != nil {
		fmt.Println("Got an error while deleting webhook:", err)
		os.Exit(1)
	}
	fmt.Printf("Successfully deleted webhook with ID: %d from repository %s/%s\n", webhookId, gitHubOrganizationName, repositoryName)
}

func init() {
	// list cmd
	webhooksListCmd.Flags().StringVarP(&repositoryName, "repository", "r", "", "GitHub repository name")
	_ = webhooksListCmd.MarkFlagRequired("repository")

	// add cmd
	webhooksAddCmd.Flags().StringVarP(&repositoryName, "repository", "r", "", "GitHub repository name")
	webhooksAddCmd.Flags().StringVarP(&webhookUrl, "webhook-url", "u", "", "GitHub webhook URL")
	_ = webhooksAddCmd.MarkFlagRequired("repository")
	_ = webhooksAddCmd.MarkFlagRequired("webhook-url")

	// delete cmd
	webhooksDeleteCmd.Flags().StringVarP(&repositoryName, "repository", "r", "", "GitHub repository name")
	webhooksDeleteCmd.Flags().Int64VarP(&webhookId, "webhook-id", "i", 0, "GitHub webhook ID")
	_ = webhooksDeleteCmd.MarkFlagRequired("repository")
	_ = webhooksDeleteCmd.MarkFlagRequired("webhook-id")

	// main cmd
	webhooksCmd.AddCommand(webhooksListCmd)
	webhooksCmd.AddCommand(webhooksAddCmd)
	webhooksCmd.AddCommand(webhooksDeleteCmd)
	addRequiredFlags(webhooksCmd)
	rootCmd.AddCommand(webhooksCmd)

}
