package cmd

import (
	"bufio"
	"dockerfile-inspector/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(outputCmd)
}

var outputCmd = &cobra.Command{
	Use:   "output",
	Short: "Runs dockerfile-inspector to generate ouput in stdout",
	Long:  `Runs dockerfile-inspector to generate ouput in stdout`,
	Run: func(cmd *cobra.Command, args []string) {
		dockerFileContent, err := os.Open(dockerFilePath)
		if err != nil {
			logrus.Errorf("Error parsing config at %s: %v", dockerFilePath, err)
			os.Exit(1)
		}
		content := bufio.NewScanner(dockerFileContent)
		utils.CheckRules(content)
	},
}
