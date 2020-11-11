package main

import (
	"dockerfile-inspector/cmd"
)

const (
	// Version represents the current release version of dockerfile-inspector
	Version = "0.1"
)

func main() {
	cmd.Execute(Version)
}
