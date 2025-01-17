package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) FetchPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	if val, exists := c.cache.Get(url); exists {
		var pokemon Pokemon
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var pokemon Pokemon
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)

	return pokemon, nil
}
