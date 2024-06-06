package pokeapi

import (
	"encoding/json"
	"io"
)

func (c *Client) GetPokemonInfo(pokemon string) (Pokemon, error) {

	pokemonURL := baseURL + "/pokemon/" + pokemon

	res := Pokemon{}
	val, ok := c.cache.Get(pokemonURL)

	if ok {
		err := json.Unmarshal(val, &res)

		if err != nil {
			return Pokemon{}, err
		}

		return res, nil
	}

	resp, err := c.httpClient.Get(pokemonURL)
	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(pokemonURL, data)

	err = json.Unmarshal(data, &res)

	if err != nil {
		return Pokemon{}, err
	}

	return res, nil

}
