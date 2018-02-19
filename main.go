package main

import (
	"fmt"
	"os"

	"github.com/lionize/gitwrap/git"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(0)
	}

	script := os.Args[1]

	if script == "init" {
		fmt.Println("Init")
	} else {
		git.RunGitPassthrough()
	}
}

func usage() {
	usageString := `usage: %s [script]

A tiny wrapper around git. If the provided script doesn't exist, all arguments will be passed through to git.

Available scripts:
	init`
	usageString = fmt.Sprintf(usageString, os.Args[0])
	fmt.Println(usageString)
}