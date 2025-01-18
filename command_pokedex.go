package main

import "fmt"

func commandPokedex(config *Config) error {
	fmt.Println("Your Pokedex:")
	for p := range pokedex.pokemon {
		fmt.Printf(" - %s\n", p)
	}

	return nil
}
