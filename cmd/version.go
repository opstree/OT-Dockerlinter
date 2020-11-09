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
	Long:  `Prints the current version of dockerfile-inspector.`,
	Run: func(cmd *cobra.Command, args []string) {
		versionOutput := fmt.Sprintf("dockerfile-inspector %s", version)
		fmt.Println(versionOutput)
	},
}
