package main

import (
	"fmt"

	"github.com/TobiasBales/shell-prompt/git"
	"github.com/TobiasBales/shell-prompt/path"
)

func main() {
	fmt.Printf("λ %v%v ", path.Indicator(), git.Indicator())
}
