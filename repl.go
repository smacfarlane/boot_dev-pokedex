package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/commands"
	"strings"
)

func repl() {
	cli := commands.Cli()
	const prompt = "Pokedex > "
	scanner := bufio.NewScanner(os.Stdin)
	config := commands.Config{}

	for {
		fmt.Printf("%s ", prompt)
		if ok := scanner.Scan(); !ok {
			fmt.Printf("Error: %v", scanner.Err())
			return
		}
		command := scanner.Text()
		command = strings.TrimSpace(command)
		if cmd, ok := cli[command]; ok {
			cmd.Run(&config)
		}
	}
}
