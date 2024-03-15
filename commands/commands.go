package commands

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func Cli() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func (c *cliCommand) Run() error {
	return c.callback()
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	for _, v := range Cli() {
		fmt.Printf("%v: %v\n", v.name, v.description)
	}

	fmt.Println()

	return nil
}

func commandExit() error {
	os.Exit(0)
	// unreachable
	return nil
}
