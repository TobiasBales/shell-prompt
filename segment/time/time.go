package time

import (
	"time"

	"github.com/TobiasBales/shell-prompt/config"
	"github.com/TobiasBales/shell-prompt/segment"
	"github.com/logrusorgru/aurora"
)

type t struct {
	c config.Config
}

func (t *t) Value() chan string {
	c := make(chan string)
	now := time.Now()

	go func() {
		if t.c.Time == nil || *t.c.Time == false {
			c <- ""
			return
		}
		c <- aurora.Gray(now.Format("15:04")).String() + " "
	}()

	return c
}

func (t *t) Placeholder() string {
	return "|time|"
}

// Indicator returns a segment representing the git status of the cwd
func Indicator(c config.Config) segment.Segment {
	return &t{c: c}
}
