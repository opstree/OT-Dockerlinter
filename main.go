package main

import (
	"fmt"
	"dockerfile-inspector/cmd"
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
