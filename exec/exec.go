package exec

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

// Execute executes the given command with the suppkied args and returns
// stdout, stderr and an error
func Execute(command string, args ...string) (string, string, error) {
	cmd := exec.Command(command, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", "", fmt.Errorf("Error getting stdout while executing \"%v %v\": %v", command, strings.Join(args, " "), err.Error())
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", "", fmt.Errorf("Error getting stderr while executing \"%v %v\": %v", command, strings.Join(args, " "), err.Error())
	}

	err = cmd.Start()
	if err != nil {
		return "", "", fmt.Errorf("Error starting: while executing \"%v %v\" %v", command, strings.Join(args, " "), err.Error())
	}

	stdoutOutput, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", "", fmt.Errorf("Error reading stdout while executing \"%v %v\": %v", command, strings.Join(args, " "), err.Error())
	}

	stderrOutput, err := ioutil.ReadAll(stderr)
	if err != nil {
		return "", "", fmt.Errorf("Error reading stderr while executing \"%v %v\": %v", command, strings.Join(args, " "), err.Error())
	}

	err = cmd.Wait()
	if err != nil {
		return "", "", fmt.Errorf("Error waiting: while executing \"%v %v\" %v", command, strings.Join(args, " "), err.Error())
	}

	return strings.Trim(string(stdoutOutput), "\n "), strings.Trim(string(stderrOutput), "\n "), nil
}
