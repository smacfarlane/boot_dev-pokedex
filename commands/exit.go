package commands

import (
	"os"
)

func commandExit(_ *Config) error {
	os.Exit(0)
	// unreachable
	return nil
}
