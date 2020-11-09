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
func DockerFileCheckStatus(check string, content *bufio.Scanner, description string, level string, recommendation string) []CheckOutput {
	var checkOuput []CheckOutput
	r, err := regexp.Compile(check)
	if err != nil {
		log.Errorf("Unable to compile the regular expression: %v", err)
	}
	line := 1
	for content.Scan() {
		if r.MatchString(content.Text()) {
			checkOuput = append(checkOuput, CheckOutput{
				LineNumber:     line,
				Line:           content.Text(),
				Description:    description,
				Level:          level,
				Recommendation: recommendation,
			})
		}
		line++
	}
	return checkOuput
}

func CheckRules(content *bufio.Scanner) {
	for _, rule := range rules.DockerFileRules {
		fmt.Println(DockerFileCheckStatus(rule.RegularExpression, content, rule.Description, rule.Level, rule.Recommendation))
	}
}
