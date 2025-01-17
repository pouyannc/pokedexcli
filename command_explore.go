package main

import (
	"fmt"
)

func commandExplore(config *Config) error {
	fmt.Printf("Exploring %s...\n", config.argument)

	pokearea, err := config.pokeapiClient.FetchArea(config.argument)
	if err != nil {
		return fmt.Errorf("unable to explore this area: %v", err)
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range pokearea.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}
