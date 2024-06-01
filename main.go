package main

import (
	"pokedexcli/internal/pokeapi"
	"time"
)

func main() {

	cfg := &config{
		pokeapiClient: pokeapi.NewClient(time.Second*3, time.Minute*3),
	}
	startPokedex(cfg)
}
