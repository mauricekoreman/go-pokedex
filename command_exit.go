package main

import (
	"fmt"
	"os"

	"github.com/mauricekoreman/go-pokedex/internal/pokecache"
)

func commandExit(cfg *config, cache *pokecache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
