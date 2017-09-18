package k8n

import (
	"github.com/TobiasBales/shell-prompt/exec"
	"github.com/TobiasBales/shell-prompt/segment"
	"github.com/TobiasBales/shell-prompt/utils"
	"github.com/logrusorgru/aurora"
)

func contextIndicator() chan string {
	c := make(chan string)

	go func() {
		stdout, _, err := exec.Execute("kubectl", "config", "current-context")

		select {
		case branch := <-stdout:
			c <- aurora.Blue(utils.Trim(branch)).String() + " "
		case <-err:
			c <- ""
		}

	}()

	return c
}

type k8n struct{}

func (k *k8n) Value() chan string {
	c := make(chan string)

	go func() {
		ctx := <-contextIndicator()
		c <- ctx
	}()

	return c
}

func (k *k8n) Placeholder() string {
	return "|k8n|"
}

// Indicator returns a segment representing the current k8n context
func Indicator() segment.Segment {
	return &k8n{}
}
