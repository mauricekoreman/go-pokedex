package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]

	pokemonData, found := cfg.caughtPokemon[pokemonName]
	if !found {
		return errors.New("pokemon not found")
	}

	fmt.Printf("Name: %s\n", pokemonData.Name)
	fmt.Printf("Height: %d\n", pokemonData.Height)
	fmt.Printf("Weight: %d\n", pokemonData.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemonData.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, pType := range pokemonData.Types {
		fmt.Printf("  - %s\n", pType.Type.Name)
	}

	return nil
}
