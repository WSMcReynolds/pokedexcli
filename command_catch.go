package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(c *config) error {
	pokemon := c.parameters[0]

	pokemonInfo, err := c.pokeapiClient.GetPokemonInfo(pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	if pokemonCaught(pokemonInfo.BaseExperience) {
		c.caughtPokemon[pokemon] = pokemonInfo
		fmt.Printf("%s was caught!", pokemon)
		return nil
	}

	fmt.Printf("%s escaped!", pokemon)
	return nil
}

func pokemonCaught(baseEXP int) bool {
	chance := rand.Float64()
	return float64(baseEXP)*chance < 50.0
}
