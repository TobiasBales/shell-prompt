package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/TobiasBales/shell-prompt/log"
	toml "github.com/pelletier/go-toml"
)

// Config contains the configuration for the shell-prompt
type Config struct {
	Line *string
	VCS  *bool
}

func (c *Config) populateWithDefaults(dc Config) {
	if c.Line == nil {
		c.Line = dc.Line
	}

	if c.VCS == nil {
		c.VCS = dc.VCS
	}
}

func getDefaultConfig() Config {
	defaultLine := "λ |path||git| "
	defaultVCS := true

	return Config{Line: &defaultLine, VCS: &defaultVCS}
}

// ReadConfig reads the configuration from the hard drive
func ReadConfig() chan Config {
	ch := make(chan Config)

	go func() {

		var c Config
		dc := getDefaultConfig()

		cp := strings.Replace("$HOME/.shellpromptrc", "$HOME", os.Getenv("HOME"), 1)

		raw, err := ioutil.ReadFile(cp)
		if err != nil {
			log.Println(fmt.Sprintf("Error reading rc file: %v", err.Error()))

			ch <- dc
			return
		}

		err = toml.Unmarshal(raw, &c)
		if err != nil {
			log.Println(fmt.Sprintf("Error parsing rc file: %v", err.Error()))

			ch <- dc
			return
		}

		c.populateWithDefaults(dc)

		ch <- c
	}()

	return ch
}
