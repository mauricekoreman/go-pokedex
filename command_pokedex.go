package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemonName := range cfg.caughtPokemon {
		fmt.Println("- ", pokemonName.Name)
	}

	return nil
}
