package commands

import (
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
	script := os.Args[1]

	for _, cmd := range commands {
		if (cmd.Name == script) {
			cmd.Run(os.Args[1:])
			return
		}
	}

	git.RunGitPassthrough()
}
