package path

import (
	"os"
	"strings"
)

func getPwd() string {
	return os.Getenv("PWD")
}

func getHome() string {
	return os.Getenv("HOME")
}

// Indicator returns the path string of the current working dir
func Indicator() string {
	return strings.Replace(getPwd(), getHome(), "~", 1)
}
