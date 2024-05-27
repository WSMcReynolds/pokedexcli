package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"strings"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationsURL    *string
	previousLocationURL *string
}

func startPokedex(c *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := sanitizeCommand(scanner.Text())
		cmd := input[0]
		runCommand(cmd, c)
		fmt.Print("\n")
	}
}

func runCommand(cmd string, c *config) {
	cm := getCommandsMap()
	command, ok := cm[cmd]
	if !ok {
		fmt.Println("Invalid command input: " + cmd)
		commandHelp(c)
	} else {
		err := command.callback(c)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func sanitizeCommand(cmd string) []string {
	cmd = strings.ToLower(cmd)
	words := strings.Fields(cmd)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommandsMap() map[string]cliCommand {

	commandMap := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 locations",
			callback:    commandMapb,
		},
	}

	return commandMap
}