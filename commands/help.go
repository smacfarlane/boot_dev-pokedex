package commands

import (
	"fmt"
)

func commandHelp(_ *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	for _, v := range Cli() {
		fmt.Printf("%v: %v\n", v.name, v.description)
	}

	fmt.Println()

	return nil
}
