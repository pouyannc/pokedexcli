package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *Config) error {
	pokemonName := config.argument
	pokemonData, err := config.pokeapiClient.FetchPokemon(pokemonName)
	if err != nil {
		return fmt.Errorf("could not fetch pokemon data: %v", err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	chance := 39.0 / float64(pokemonData.BaseExperience) * 100.0
	fmt.Printf("Chance: %.0f%%\n", chance)
	catchInt := rand.Intn(pokemonData.BaseExperience)
	if catchInt <= 39 {
		fmt.Printf("%s was caught!\n", pokemonName)
		pokedex.pokemon[pokemonName] = pokemonData
		fmt.Println("Current pokedex contains:")
		for p := range pokedex.pokemon {
			fmt.Println(p)
		}
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	return nil
}
