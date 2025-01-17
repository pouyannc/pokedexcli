package main

import (
	"time"

	"github.com/pouyannc/pokedexcli/internal/pokeapi"
	"github.com/pouyannc/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(10 * time.Second)
	cfg := &Config{
		pokeapiClient: pokeClient,
		pokeCache:     pokeCache,
	}

	startRepl(cfg)
}
