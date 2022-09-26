package cmd

import (
	"fmt"
	"github.com/google/go-github/v45/github"
	"github.com/spf13/cobra"
	"os"
)

const (
	envPrefix = "GWC"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gha-cli",
		Short: "gha-cli (GitHub App Cli) is a single purpose tool used to work with GitHub using GitHub App authentication",
		Long:  "gha-cli (GitHub App Cli) is a single purpose tool used to work with GitHub using GitHub App authentication",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
			return initializeConfig(cmd)
		},
		/* Run: func(cmd *cobra.Command, args []string) {
			//
		}, */
	}

	debug                  bool
	gitHubAppId            int64
	gitHubInstallationId   int64
	gitHubOrganizationName string
	gitHubPrivateKeyFile   string
	ghClient               *github.Client
)

func init() {

}

// Wrapper function that only marks GWC related flags as required for our own custom commands.
// this is to avoid built in cobra commands such as `completion` to fail
// Note: Should only be added to top level commands
func addRequiredFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug")
	cmd.PersistentFlags().Int64Var(&gitHubAppId, "appid", 0, "GitHub Application ID")
	cmd.PersistentFlags().Int64Var(&gitHubInstallationId, "installationId", 0, "GitHub Installation ID")
	cmd.PersistentFlags().StringVar(&gitHubOrganizationName, "org", "", "GitHub Organization Name")
	cmd.PersistentFlags().StringVarP(&gitHubPrivateKeyFile, "private-key-file", "f", "", "GitHub App oauth2 private key file")

	// mark required flags
	_ = cmd.MarkPersistentFlagRequired("appid")
	_ = cmd.MarkPersistentFlagRequired("installationId")
	_ = cmd.MarkPersistentFlagRequired("org")
	_ = cmd.MarkPersistentFlagRequired("private-key-file")
	_ = cmd.MarkPersistentFlagRequired("private-key-file")
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		_, fErr := fmt.Fprintln(os.Stderr, err)
		if fErr != nil {
			fmt.Println(err)
		}
		os.Exit(1)
	}
}
