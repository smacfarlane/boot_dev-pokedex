package commands

import (
	"errors"
	"fmt"
	"pokedex/internal/pokeapi"
)

func commandMap(c *Config) error {
	resp, err := pokeapi.GetLocationAreas(c.Next)

	if err != nil {
		return err
	}

	if resp.Next != nil {
		c.Next = resp.Next
	}

	for _, result := range resp.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapb(c *Config) error {
	if c.Previous == nil {
		return errors.New("no previous page")
	}

	resp, err := pokeapi.GetLocationAreas(c.Previous)

	if err != nil {
		return err
	}

	if resp.Previous != nil {
		c.Previous = resp.Previous
	}

	for _, result := range resp.Results {
		fmt.Println(result.Name)
	}

	return nil
}
