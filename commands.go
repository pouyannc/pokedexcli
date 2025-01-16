package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(Config) error
	cf          Config
}

type Config struct {
	locationNext *string
	locationPrev *string
}

var locationConfig Config

func commandExit(config Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config Config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")

	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	fmt.Println()
	return nil
}

func commandMap(config Config) error {
	if locationConfig.locationNext == nil {
		fmt.Println("you're on the final page")
		return nil
	}

	locations, err := fetchLocations(locationConfig.locationNext)
	if err != nil {
		return fmt.Errorf("could not fetch locations: %v", err)
	}

	for _, loc := range locations {
		fmt.Println(loc)
	}
	return nil
}

func commandMapb(config Config) error {
	if locationConfig.locationPrev == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	locations, err := fetchLocations(locationConfig.locationPrev)
	if err != nil {
		return fmt.Errorf("could not fetch locations: %v", err)
	}

	for _, loc := range locations {
		fmt.Println(loc)
	}
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Display the names of 20 location areas in the Pokemon world. Each subsequent call to map displays the next 20 locations.",
			callback:    commandMap,
			cf:          locationConfig,
		},
		"mapb": {
			name:        "mapb",
			description: "Functions the same as map, except shows the previous 20 locations on each subsequent call.",
			callback:    commandMapb,
			cf:          locationConfig,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
			cf:          Config{},
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
			cf:          Config{},
		},
	}
}
