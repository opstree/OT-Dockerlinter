package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"fmt"
	"dockerfile-inspector/pkg/analyzer"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK = iota + 1
	ExitCodeParseFlagsError
	ExitCodeNoExistError
	ExitCodeFileError
	ExitCodeAstParseError
	ExitCodeLintCheckError
)

type sliceString []string

func (ss *sliceString) String() string {
	return fmt.Sprintf("%s", *ss)
}

func (ss *sliceString) Set(value string) error {
	*ss = append(*ss, value)
	return nil
}

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
		runAudit()
	},
}

func runAudit() {
	var ignoreRules sliceString
	file, err := os.Open(dockerFilePath)
	if err != nil {
		logrus.Errorf("Error opening Dockerfile %s: %v", dockerFilePath, err)
		os.Exit(1)
	}
	r, err := parser.Parse(file)
	if err != nil {
		logrus.Errorf("Error parsing Dockefile %s: %v", dockerFilePath, err)
		os.Exit(1)
	}
	analyzer := linter.NewAnalyzer(ignoreRules)
	rst, err := analyzer.Run(r.AST)
	if err != nil {
		logrus.Errorf("Unable to run analyzer %s: %v", dockerFilePath, err)
	}
	fmt.Println(rst)
}
