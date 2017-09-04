package path

import (
	"os"
	"regexp"
	"strings"

	"github.com/TobiasBales/shell-prompt/segment"
	"github.com/logrusorgru/aurora"
)

type path struct {
}

func (p *path) Value() chan string {
	gh := regexp.MustCompile("(.*github.com/)([a-zA-Z0-9-_]*)/([a-zA-Z0-9-_]*)(/?.*)")
	c := make(chan string)

	go func() {
		p := strings.Replace(os.Getenv("PWD"), os.Getenv("HOME"), "~", 1)
		if !gh.MatchString(p) {
			c <- p + " "
			return
		}

		m := gh.FindStringSubmatch(p)
		c <- m[1] + aurora.Cyan(m[2]).String() + "/" + aurora.Red(m[3]).String() + m[4] + " "
	}()

	return c
}

func (p *path) Placeholder() string {
	return "|path|"
}

// Indicator returns a segment representing the current path
func Indicator() segment.Segment {
	return &path{}
}