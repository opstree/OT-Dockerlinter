package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	auditCmd.PersistentFlags().StringVarP(&dockerFilePath, "docker.file", "d", "Dockerfile", "Location of Dockerfile.")
	auditCmd.PersistentFlags().StringVarP(&outputFormat, "ouput.format", "o", "table", "Output format of report. available options - json, table, xml")
	rootCmd.AddCommand(auditCmd)
}

var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Runs dockerfile-inspector audit",
	Long:  `Runs dockerfile-inspector audit`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Open(dockerFilePath)
		if err != nil {
			logrus.Errorf("Error parsing config at %s: %v", dockerFilePath, err)
			os.Exit(1)
		}
	},
}
