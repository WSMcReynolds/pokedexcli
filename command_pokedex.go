package main

import "fmt"

func commandPokedex(c *config) error {
	if len(c.caughtPokemon) == 0 {
		fmt.Println("No caught pokemon!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range c.caughtPokemon {
		fmt.Printf(" - %v\n", pokemon.Name)
	}

	return nil
}
