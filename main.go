package main

import (
	"time"

	"github.com/pouyannc/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 2*time.Minute)
	cfg := &Config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
