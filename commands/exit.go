package commands

import (
	"os"
)

func commandExit(_ *Config, _ ...string) error {
	os.Exit(0)
	// unreachable
	return nil
}
