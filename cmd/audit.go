package cmd

import (
	"dockerfile-inspector/pkg/analyzer"
	"dockerfile-inspector/pkg/rule"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"strconv"
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
	auditCmd.PersistentFlags().StringVarP(&outputFormat, "ouput.format", "o", "table", "Output format of report. available options - json, table")
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
	rst, err := analyzer.Run(r.AST, dockerFilePath)
	if err != nil {
		logrus.Errorf("Unable to run analyzer %s: %v", dockerFilePath, err)
	}
	if outputFormat == "json" {
		result, err := json.Marshal(rst)
		if err != nil {
			logrus.Errorf("Unable to convert output to JSON %v", err)
		}
		fmt.Println(string(result))
	} else if outputFormat == "table" {
		printTable(rst)
	}
}

func printTable(result []rules.Result) {
	csvFile, err := os.Create("./data.csv")
	if err != nil {
		logrus.Errorf("Unable to create CSV file %v", err)
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	header := []string{"Line Number", "Line", "Code", "Description", "Severity", "Filename"}
	writer.Write(header)
	for _, data := range result {
		var row []string
		row = append(row, strconv.Itoa(data.LineNumber))
		row = append(row, data.Line)
		row = append(row, data.Code)
		row = append(row, data.Description)
		row = append(row, data.Severity)
		row = append(row, data.FileName)
		writer.Write(row)
	}
	writer.Flush()

	table, err := tablewriter.NewCSV(os.Stdout, "data.csv", true)
	if err != nil {
		logrus.Errorf("Unable to open CSV file %v", err)
	}
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetRowLine(true)
	table.SetHeader([]string{"Line", "Code", "Description", "Severity"})
	table.Render()
	os.Remove("data.csv")
}
