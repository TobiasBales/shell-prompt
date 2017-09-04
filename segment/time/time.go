package time

import (
	"time"

	"github.com/TobiasBales/shell-prompt/segment"
	"github.com/logrusorgru/aurora"
)

type t struct {
}

func (t *t) Value() chan string {
	c := make(chan string)
	now := time.Now()

	go func() {
		c <- aurora.Gray(now.Format("15:04")).String() + " "
	}()

	return c
}

func (t *t) Placeholder() string {
	return "|time|"
}

// Indicator returns a segment representing the git status of the cwd
func Indicator() segment.Segment {
	return &t{}
}
