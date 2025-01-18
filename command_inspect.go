package main

import "fmt"

func commandInspect(config *Config) error {
	pokemonName := config.argument

	pokeData, ok := pokedex.pokemon[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokeData.Name)
	fmt.Printf("Height: %v\n", pokeData.Height)
	fmt.Printf("Weight: %v\n", pokeData.Weight)
	fmt.Println("Stats:")
	for _, stats := range pokeData.Stats {
		fmt.Printf("\t-%v: %v\n", stats.Stat.Name, stats.BaseStat)
	}
	fmt.Println("Types:")
	for _, types := range pokeData.Types {
		fmt.Printf("\t-%v\n", types.Type.Name)
	}

	return nil
}
