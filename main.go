package main

import (
	"dockerfile-inspector/cmd"
	"fmt"
)

var (
	GitCommit string
	Version   string
)

func main() {
	version := fmt.Sprintf("%s: %s", Version, GitCommit)
	cmd.Execute(version)
}
