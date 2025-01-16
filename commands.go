package main

import (
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")

	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	fmt.Println()
	return nil
}

func commandMap(config *Config) error {
	locationsRes, err := config.pokeapiClient.FetchLocationsRes(config.locationNext)
	if err != nil {
		return fmt.Errorf("could not fetch locations: %v", err)
	}

	config.locationNext = locationsRes.Next
	config.locationPrev = locationsRes.Previous

	for _, loc := range locationsRes.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(config *Config) error {
	if config.locationPrev == nil {
		return errors.New("you're on the first page")
	}

	locationsRes, err := config.pokeapiClient.FetchLocationsRes(config.locationPrev)
	if err != nil {
		return fmt.Errorf("could not fetch locations: %v", err)
	}

	config.locationNext = locationsRes.Next
	config.locationPrev = locationsRes.Previous

	for _, loc := range locationsRes.Results {
		fmt.Println(loc.Name)
	}
	return nil
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
