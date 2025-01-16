package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	locationsUrl := "https://pokeapi.co/api/v2/location-area/"
	locationConfig.locationNext = &locationsUrl

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
			command.callback(command.cf)
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	loweredText := strings.ToLower(text)
	result := strings.Fields(loweredText)
	return result
}
