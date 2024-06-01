package main

import "fmt"

func commandExplore(c *config) error {
	if len(c.parameters) < 1 {
		fmt.Printf("Please include an area name to explore.\nExample: explore fuego-ironworks-area")
	}

	area := c.parameters[0]

	areaInfo, err := c.pokeapiClient.GetAreaInfo(area)

	if err != nil {
		return err
	}

	pokemonEncounters := areaInfo.PokemonEncounters

	for _, pokemon := range pokemonEncounters {
		fmt.Printf("-  %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
