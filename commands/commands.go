package commands

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
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
		"map": {
			name:        "map",
			description: "Display next 20 locations in the world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 locations in the world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Display pokemon for a region",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a named Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display all caught Pokemon",
			callback:    commandPokedex,
		},
	}
}

func (c *cliCommand) Run(conf *Config, args ...string) error {
	return c.callback(conf, args...)
}
