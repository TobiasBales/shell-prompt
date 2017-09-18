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
	c := <-config.ReadConfig()

	segments := []segment.Segment{
		git.Indicator(c),
		path.Indicator(c),
		time.Indicator(c),
		k8n.Indicator(c),
	}

	prompt := *c.Line
	for _, s := range segments {
		prompt = strings.Replace(prompt, s.Placeholder(), <-s.Value(), 1)
	}

	fmt.Print(prompt)
}
