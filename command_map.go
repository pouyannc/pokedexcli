package main

import (
	"errors"
	"fmt"
)

func commandMap(config *Config) error {
	locationsRes, err := config.pokeapiClient.FetchLocationsRes(config.locationNext)
	if err != nil {
		return fmt.Errorf("could not fetch locations: %v", err)
	}

	config.locationNext = locationsRes.Next
	config.locationPrev = locationsRes.Previous

	for _, loc := range locationsRes.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(config *Config) error {
	if config.locationPrev == nil {
		return errors.New("you're on the first page")
	}

	locationsRes, err := config.pokeapiClient.FetchLocationsRes(config.locationPrev)
	if err != nil {
		return fmt.Errorf("could not fetch locations: %v", err)
	}

	config.locationNext = locationsRes.Next
	config.locationPrev = locationsRes.Previous

	for _, loc := range locationsRes.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
