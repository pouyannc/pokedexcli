package main

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Display the names of 20 location areas in the Pokemon world. Each subsequent call to map displays the next 20 locations.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Functions the same as map, except shows the previous 20 locations on each subsequent call.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Takes the name of a location area as an argument and lists all the Pokemon located there.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Takes the name of a Pokemon as an argument and attempts to catch it, adding it to the Pokedex on successful catch.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Takes the name of a caught Pokemon and prints it's stats (name, height, weight, stats, and type(s)).",
			callback:    commandInspect,
		},
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
