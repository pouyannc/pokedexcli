package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pouyannc/pokedexcli/internal/pokeapi"
)

type Config struct {
	pokeapiClient pokeapi.Client
	locationNext  *string
	locationPrev  *string
	argument      string
}

func startRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, ok := getCommands()[commandName]
		if ok {
			if len(words) > 1 {
				cfg.argument = words[1]
			}
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("Unknown command: %v\n", words[0])
		}
	}
}

func cleanInput(text string) []string {
	loweredText := strings.ToLower(text)
	result := strings.Fields(loweredText)
	return result
}
