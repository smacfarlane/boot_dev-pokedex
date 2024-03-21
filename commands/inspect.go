package commands

import (
	"errors"
	"fmt"
)

func commandInspect(c *Config, args ...string) error {
	if len(args) == 0 {
		return errors.New("expected name of pokemon")
	}

	name := args[0]
	pokemon, ok := c.Pokedex.Pokemon[name]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("%s\n", pokemon.String())

	return nil
}
