package k8n

import (
	"github.com/TobiasBales/shell-prompt/config"
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

type k8n struct {
	c config.Config
}

func (k *k8n) Value() chan string {
	c := make(chan string)
	go func() {
		if k.c.K8n == nil || *k.c.K8n == false {
			c <- ""
			return
		}

		ctx := <-contextIndicator()
		c <- ctx
	}()

	return c
}

func (k *k8n) Placeholder() string {
	return "|k8n|"
}

// Indicator returns a segment representing the current k8n context
func Indicator(c config.Config) segment.Segment {
	return &k8n{c: c}
}
