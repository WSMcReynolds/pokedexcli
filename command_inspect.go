package main

import (
	"fmt"
	"pokedexcli/internal/pokeapi"
	"strconv"
)

func commandInspect(c *config) error {
	pokemon, ok := c.caughtPokemon[c.parameters[0]]

	if !ok {
		fmt.Println("you have not caught that pokemon!")
		return nil
	}

	stats := getStats(pokemon)
	types := getTypes(pokemon)

	fmt.Printf("Height: %v\nWeight: %v\nStats:\n%vTypes:\n%v", pokemon.Height, pokemon.Weight, stats, types)
	return nil
}

func getStats(p pokeapi.Pokemon) string {
	stats := ""

	for _, stat := range p.Stats {
		baseStat := strconv.Itoa(stat.BaseStat)
		stats += "  -" + stat.Stat.Name + ": " + baseStat + "\n"
	}

	return stats
}

func getTypes(p pokeapi.Pokemon) string {
	types := ""

	for _, pokemonType := range p.Types {
		types += "  - " + pokemonType.Type.Name + "\n"
	}

	return types
}
