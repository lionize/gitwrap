package commands

import (
	"fmt"
	"os"

	"github.com/lionize/gitwrap/git"
)

// Command represents a program command
type Command struct {
	Name        string
	Description string
	Run         func(args []string)
}

var commands = make([]Command, 0)

func init() {
	AddCommand(Init)
}

// AddCommand adds a command to the list of runnable commands
func AddCommand(command Command) {
	commands = append(commands, command)
}

// Execute selects the command to execute based on process args and runs it
func Execute() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(0)
	}

	script := os.Args[1]

	for _, cmd := range commands {
		if cmd.Name == script {
			cmd.Run(os.Args[1:])
			return
		}
	}

	git.RunGitPassthrough()
}

func usage() {
	usageString := `usage: %s [script]

A tiny wrapper around git. If the provided script doesn't exist, all arguments will be passed through to git.

Available scripts:
%s`
	usageString = fmt.Sprintf(usageString, os.Args[0], commandDescriptions())
	fmt.Println(usageString)
}

func commandDescriptions() string {
	descriptions := ""

	for _, cmd := range commands {
		descriptions += fmt.Sprintf("\t%s - %s\n", cmd.Name, cmd.Description)
	}

	return descriptions
}
