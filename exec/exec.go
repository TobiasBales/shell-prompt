package exec

import (
	"io/ioutil"
	"os/exec"
)

// Execute executes the given command with the suppkied args and returns
// stdout, stderr and an error
func Execute(command string, args ...string) (chan string, chan string, chan error) {
	stdoutC := make(chan string)
	stderrC := make(chan string)
	errC := make(chan error)

	go func() {
		cmd := exec.Command(command, args...)

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			select {
			case errC <- err:
			}
			return
		}

		stderr, err := cmd.StderrPipe()
		if err != nil {
			select {
			case errC <- err:
			}
			return
		}

		err = cmd.Start()
		if err != nil {
			select {
			case errC <- err:
			}
		}

		stdoutOutput, err := ioutil.ReadAll(stdout)
		if err != nil {
			select {
			case errC <- err:
			}
		}

		stderrOutput, err := ioutil.ReadAll(stderr)
		if err != nil {
			select {
			case errC <- err:
			}
		}

		err = cmd.Wait()
		if err != nil {
			select {
			case errC <- err:
			}
		}

		select {
		case stdoutC <- string(stdoutOutput):
		case stderrC <- string(stderrOutput):
		}
	}()

	return stdoutC, stderrC, errC
}
