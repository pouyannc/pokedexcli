package main

import "github.com/pouyannc/pokedexcli/internal/pokeapi"

type Pokedex struct {
	pokemon map[string]pokeapi.Pokemon
}
