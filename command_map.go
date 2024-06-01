package main

import (
	"errors"
	"fmt"
)

func commandMap(c *config) error {

	locationResponse, err := c.pokeapiClient.ListLocations(c.nextLocationsURL)
	if err != nil {
		return err
	}

	c.nextLocationsURL = locationResponse.Next
	c.previousLocationsURL = locationResponse.Previous

	for _, location := range locationResponse.Results {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}

func commandMapb(c *config) error {

	if c.previousLocationsURL == nil {
		return errors.New("currently on starting page of locations")
	}

	locationResponse, err := c.pokeapiClient.ListLocations(c.previousLocationsURL)
	if err != nil {
		return err
	}

	c.nextLocationsURL = locationResponse.Next
	c.previousLocationsURL = locationResponse.Previous

	for _, location := range locationResponse.Results {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}
