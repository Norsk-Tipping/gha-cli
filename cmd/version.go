package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	BuildDate string
	GitTag    string
	GitHash   string
	GitBranch string

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Display version",
		Long:  "Display version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(Version)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
