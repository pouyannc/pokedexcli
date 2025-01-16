package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Location_areas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func fetchLocations(url *string) ([]string, error) {
	if url == nil {
		return nil, fmt.Errorf("No more locations to get")
	}

	res, err := http.Get(*url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var locations Location_areas
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locations)
	if err != nil {
		return nil, err
	}

	locNames := make([]string, len(locations.Results))
	for i, result := range locations.Results {
		locNames[i] = result.Name
	}

	locationConfig.locationNext = locations.Next
	locationConfig.locationPrev = locations.Previous

	return locNames, nil
}
