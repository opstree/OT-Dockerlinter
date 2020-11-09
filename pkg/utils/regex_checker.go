package utils

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"regexp"
)

// DockerFileCheckStatus will return the status of check
func DockerFileCheckStatus(check string, content *bufio.Scanner) {
	r, err := regexp.Compile(check)
	if err != nil {
		log.Errorf("Unable to compile the regular expression: %v", err)
	}
	line := 1
	for content.Scan() {
		if r.MatchString(content.Text()) {
			fmt.Println(content.Text(), line)
		}
		line++
	}
}
