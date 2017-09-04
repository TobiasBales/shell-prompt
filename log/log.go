package log

import (
	"log"
	"os"
	"strings"
)

// Println logs s to the default logfile
func Println(s string) {
	filename := strings.Replace("$HOME/.shellprompt.log", "$HOME", os.Getenv("HOME"), 1)

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Error logging to %v: %v\n", filename, err.Error())
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	logger.Println(s)
}
