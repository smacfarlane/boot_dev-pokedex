package commands

import (
	"fmt"
)

func commandPokedex(c *Config, _ ...string) error {
	fmt.Println("Your pokedex:")
	fmt.Printf("%s\n", c.Pokedex.String())
	return nil
}
