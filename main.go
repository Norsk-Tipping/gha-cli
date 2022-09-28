package main

import (
	"github.com/Norsk-Tipping/gha-cli/cmd"
)

func main() {
	if len(cmd.GitTag) == 0 {
		cmd.Version = cmd.GitBranch + " (" + cmd.GitHash + ")"
	} else {
		cmd.Version = cmd.GitTag + " (" + cmd.GitHash + ")"
	}
	cmd.Execute()
}
