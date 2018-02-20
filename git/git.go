package git

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/lionize/gitwrap/config"
)

// RunGitInit initializes the repo and sets the user config to match the
// provided profile
func RunGitInit(profile config.Profile) {
	runGitCmd([]string{"init"})
	gitConfig("user.name", profile.Name)
	gitConfig("user.email", profile.Email)
}

// RunGitPassthrough passes process args through to git command
func RunGitPassthrough() {
	runGitCmd(os.Args[1:])
}

func gitConfig(key string, value string) {
	runGitCmd([]string{"config", key, value})
}

func runGitCmd(args []string) {
	input := append([]string{"-q", "/dev/null", "git"}, args...)
	cmd := exec.Command("script", input...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Start()
	cmd.Wait()
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
