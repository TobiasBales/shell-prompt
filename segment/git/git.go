package git

import (
	"fmt"

	"github.com/TobiasBales/shell-prompt/config"
	"github.com/TobiasBales/shell-prompt/exec"
	"github.com/TobiasBales/shell-prompt/segment"
	"github.com/TobiasBales/shell-prompt/utils"
	"github.com/logrusorgru/aurora"
)

func branchIndicator() chan string {
	c := make(chan string)

	go func() {
		stdout, _, err := exec.Execute("git", "rev-parse", "--abbrev-ref", "HEAD")

		select {
		case branch := <-stdout:
			c <- aurora.Green(utils.Trim(branch)).String()
		case <-err:
			c <- ""
		}

	}()

	return c

}

func dirtyIndicator() chan string {
	c := make(chan string)

	go func() {
		stdout, _, err := exec.Execute("git", "status", "--porcelain")

		select {
		case status := <-stdout:
			if len(utils.Trim(status)) == 0 {
				c <- aurora.Green("✓").String()
			}

			c <- aurora.Red("×").String()

		case <-err:
			c <- ""
		}

	}()

	return c
}

type git struct {
	c config.Config
}

func (g *git) Value() chan string {
	c := make(chan string)

	go func() {
		if g.c.VCS == nil || *g.c.VCS == false {
			c <- ""
			return
		}

		branch := <-branchIndicator()
		dirty := <-dirtyIndicator()

		if len(utils.Trim(branch)) == 0 {
			c <- ""
		}

		c <- fmt.Sprintf("%v %v ", branch, dirty)
	}()

	return c
}

func (g *git) Placeholder() string {
	return "|git|"
}

// Indicator returns a segment representing the git status of the cwd
func Indicator(c config.Config) segment.Segment {
	return &git{c: c}
}
