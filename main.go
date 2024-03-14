package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	fmt.Println("Hello Pikachu")
	repl()
}

func repl() {
	cli := cli()
	const prompt = "Pokedex > "
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s ", prompt)
		if ok := scanner.Scan(); !ok {
			fmt.Printf("Error: %v", scanner.Err())
			return
		}
		command := scanner.Text()
		command = strings.TrimSpace(command)
		if cmd, ok := cli[command]; ok {
			cmd.callback()
		}
	}
}

func cli() map[string]cliCommand {
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

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	for _, v := range cli() {
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
