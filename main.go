package main

import (
	"time"

	"github.com/pouyannc/pokedexcli/internal/pokeapi"
)

var pokedex Pokedex

func main() {
	pokedex.pokemon = make(map[string]pokeapi.Pokemon)

	pokeClient := pokeapi.NewClient(5*time.Second, 2*time.Minute)
	cfg := &Config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
