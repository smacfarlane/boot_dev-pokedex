package commands

import (
	"errors"
	"fmt"
)

func commandExplore(c *Config, args ...string) error {
	if len(args) == 0 {
		return errors.New("no region specified")
	}
	fmt.Printf("Exploring %s\n", args[0])

	location, err := c.Client.GetLocationArea(args[0])
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range location.Encounters {
		fmt.Println(" - " + encounter.Pokemon.Name)
	}
	return nil
}
