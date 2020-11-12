package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the current version.",
	Long:  `Prints the current version of ot-docker-linter.`,
	Run: func(cmd *cobra.Command, args []string) {
		versionOutput := fmt.Sprintf("Dockerfile Inspector %s", version)
		fmt.Println(versionOutput)
	},
}
