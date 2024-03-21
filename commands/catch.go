package commands

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(c *Config, args ...string) error {
	if len(args) == 0 {
		return errors.New("expected name of pokemon")
	}
	pokemon, err := c.Client.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a ball at %s...", pokemon.Name)

	// 0 base xp is a 50/50 chance to catch
	throw := rand.Intn(100 + pokemon.BaseExperience)
	if throw < 50 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		c.Pokedex.Add(&pokemon)
	} else {
		fmt.Printf("%s escaped\n", pokemon.Name)
	}

	return nil
}
