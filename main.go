package main

import (
	"fmt"
	"strings"

	"github.com/TobiasBales/shell-prompt/config"
	"github.com/TobiasBales/shell-prompt/segment"
	"github.com/TobiasBales/shell-prompt/segment/git"
	"github.com/TobiasBales/shell-prompt/segment/k8n"
	"github.com/TobiasBales/shell-prompt/segment/path"
	"github.com/TobiasBales/shell-prompt/segment/time"
)

func main() {
	segments := []segment.Segment{
		git.Indicator(),
		path.Indicator(),
		time.Indicator(),
		k8n.Indicator(),
	}

	c := <-config.ReadConfig()

	prompt := *c.Line
	for _, s := range segments {
		prompt = strings.Replace(prompt, s.Placeholder(), <-s.Value(), 1)
	}

	fmt.Print(prompt)
}
