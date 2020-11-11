package main

import (
	"dockerfile-inspector/cmd"
	"fmt"
)

var GitCommit string

const (
	// Version represents the current release version of dockerfile-inspector
	Version = "v0.1"
)

func main() {
	version := fmt.Sprintf("%s: %s", Version, GitCommit)
	cmd.Execute(version)
}
