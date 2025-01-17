package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) FetchLocationsRes(pageURL *string) (LocationAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	cacheData, exists := c.cache.Get(url)
	if exists {
		var locations LocationAreas
		err := json.Unmarshal(cacheData, &locations)
		if err != nil {
			return LocationAreas{}, err
		}

		return locations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreas{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}

	defer res.Body.Close()

	var locations LocationAreas

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	err = json.Unmarshal(data, &locations)
	if err != nil {
		return LocationAreas{}, err
	}

	c.cache.Add(url, data)

	return locations, nil
}
