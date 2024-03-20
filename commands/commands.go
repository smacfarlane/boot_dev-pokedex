package commands

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
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
	}
}

func (c *cliCommand) Run(conf *Config) error {
	return c.callback(conf)
}
