package commands

import (
	"errors"
	"fmt"
)

func commandMap(c *Config, _ ...string) error {
	resp, err := c.Client.GetLocationAreas(c.Next)
	if err != nil {
		return err
	}

	c.Next = resp.Next
	c.Previous = resp.Previous
	fmt.Println(resp.Previous)

	for _, result := range resp.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapb(c *Config, _ ...string) error {
	if c.Previous == nil {
		return errors.New("no previous page")
	}

	resp, err := c.Client.GetLocationAreas(c.Previous)
	if err != nil {
		return err
	}

	c.Next = resp.Next
	c.Previous = resp.Previous

	for _, result := range resp.Results {
		fmt.Println(result.Name)
	}

	return nil
}
