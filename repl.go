package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mauricekoreman/go-pokedex/internal/pokeapi"
	"github.com/mauricekoreman/go-pokedex/internal/pokecache"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationURL     *string
	previousLocationURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	cache := pokecache.NewCache(20 * time.Second)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		inputWords := cleanInput(scanner.Text())
		if len(inputWords) == 0 {
			continue
		}

		firstWord := inputWords[0]

		var secondWord string
		if len(inputWords) > 1 {
			secondWord = inputWords[1]
		}

		if command, exists := getCommands()[firstWord]; exists {
			err := command.callback(cfg, cache, secondWord)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *pokecache.Cache, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of 20 previous location areas in the Pokemon world.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays the names of the Pokemon species that can be found in a specific location area.",
			callback:    commandExplore,
		},
	}
}
