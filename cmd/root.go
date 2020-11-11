package cmd

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	dockerFilePath    string
	outputFormat      string
	logLevel          string
	logFmt            string
	version           string
	dockerFileContent *os.File
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&logLevel, "log.level", "", logrus.InfoLevel.String(), "dockerfile-inspector logging level.")
	rootCmd.PersistentFlags().StringVarP(&logFmt, "log.format", "", "text", "dockerfile-inspector log format.")
	flag.Parse()
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
}

var rootCmd = &cobra.Command{
	Use:   "dockerfile-inspector",
	Short: "dockerfile-inspector",
	Long:  `A tool for checking Dockerfile best practices.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		parsedLevel, err := logrus.ParseLevel(logLevel)
		if err != nil {
			logrus.Errorf("log-level flag has invalid value %s", logLevel)
		} else {
			logrus.SetLevel(parsedLevel)
		}
		if logFmt == "json" {
			logrus.SetFormatter(&logrus.JSONFormatter{})
		} else {
			logrus.SetFormatter(&logrus.TextFormatter{
				DisableColors: true,
				FullTimestamp: true,
			})
		}

	},
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			logrus.Error(err)
		}
		os.Exit(1)
	},
}

// Execute the stuff
func Execute(VERSION string) {
	version = VERSION
	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
