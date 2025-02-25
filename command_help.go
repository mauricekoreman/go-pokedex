package main

import (
	"fmt"

	"github.com/mauricekoreman/go-pokedex/internal/pokecache"
)

func commandHelp(cfg *config, cache *pokecache.Cache) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")

	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
