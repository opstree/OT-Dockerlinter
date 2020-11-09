package utils

import (
	"bufio"
	"dockerfile-inspector/pkg/rules"
	"fmt"
	log "github.com/sirupsen/logrus"
	"regexp"
)

type CheckOutput struct {
	LineNumber     int    `json:"line_number"`
	Line           string `json:"line"`
	Description    string `json:"description"`
	Level          string `json:"level"`
	Recommendation string `json:"recommendation"`
}

// DockerFileCheckStatus will return the status of check
func DockerFileCheckStatus(content *bufio.Scanner) []CheckOutput {
	var checkOutput []CheckOutput

	line := 1
	for content.Scan() {
		for _, rule := range rules.DockerFileRules {
			r, err := regexp.Compile(rule.RegularExpression)
			if err != nil {
				log.Errorf("Unable to compile the regular expression: %v", err)
			}
			if r.MatchString(content.Text()) {
				checkOutput = append(checkOutput, CheckOutput{
					LineNumber:     line,
					Line:           content.Text(),
					Description:    rule.Description,
					Level:          rule.Level,
					Recommendation: rule.Recommendation,
				})
			}
		}
		line++
	}
	return checkOutput
}

// CheckRules will check all the rules for Dockerfile
func CheckRules(content *bufio.Scanner) {
	fmt.Println(DockerFileCheckStatus(content))
}
