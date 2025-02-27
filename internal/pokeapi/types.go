package pokeapi

type LocationsResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
	} `json:"pokemon"`
}

type PokemonResponse struct {
	BaseExperience int            `json:"base_experience"`
	Height         int            `json:"height"`
	Name           string         `json:"name"`
	Stats          []PokemonStats `json:"stats"`
}

// weight, height, states, types, base_experience
type PokemonStats struct {
	BaseStat int `json:"base_stat"`
	Stat     struct {
		Name string `json:"name"`
	} `json:"stat"`
}
