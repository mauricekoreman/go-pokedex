package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Response struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

func commandMap(c *Config) error {
	URL := c.next
	if URL == "" {
		URL = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	}

	res, err := http.Get(URL)
	if err != nil {
		fmt.Println("Error fetching data", err)
		return nil
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading results", err)
		return nil
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Error unmarshalling JSON", err)
		return nil
	}

	c.next = response.Next
	c.previous = response.Previous

	for _, el := range response.Results {
		fmt.Println(el.Name)
	}

	return nil
}

func commandMapb(c *Config) error {
	URL := c.previous
	if URL == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	res, err := http.Get(URL)
	if err != nil {
		fmt.Println("Error fetching data", err)
		return nil
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading results", err)
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Error unmarshalling JSON", err)
		return nil
	}

	c.next = response.Next
	c.previous = response.Previous

	for _, el := range response.Results {
		fmt.Println(el.Name)
	}

	return nil
}
