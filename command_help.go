package main

import (
	"fmt"
)

func commandHelp(c *config) error {
	cm := getCommandsMap()
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("---Usage---")
	for _, command := range cm {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
