package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/commands"
	"pokedex/internal/pokeapi"
	"pokedex/internal/pokedex"
	"strings"
)

func repl() {
	cli := commands.Cli()
	const prompt = "Pokedex > "
	scanner := bufio.NewScanner(os.Stdin)
	config := commands.Config{
		Client:  pokeapi.NewClient(),
		Pokedex: pokedex.NewPokedex(),
	}

	for {
		fmt.Printf("%s ", prompt)
		if ok := scanner.Scan(); !ok {
			fmt.Printf("Error: %v", scanner.Err())
			return
		}
		input := scanner.Text()
		input = strings.TrimSpace(input)
		command := strings.Split(input, " ")[0]
		args := strings.Split(input, " ")[1:]
		if cmd, ok := cli[command]; ok {
			err := cmd.Run(&config, args...)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
			}
		} else {
			fmt.Fprintf(os.Stderr, "error: unknown command '%v'\n", input)
		}
	}
}
