package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// RunGitPassthrough passes process args through to git command
func RunGitPassthrough() {
	args := append([]string{"-q", "/dev/null", "git"}, strings.Join(os.Args[1:], " "))
	cmd := exec.Command("script", args...)
	output, err := cmd.CombinedOutput()
	printError(err)
	printOutput(output)
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("%s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("%s", string(outs))
	}
}
