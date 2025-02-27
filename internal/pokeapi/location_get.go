package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) GetLocation(locationAreaName string) (LocationResponse, error) {
	url := baseURL + "/location-area/" + locationAreaName

	if cachedData, exists := c.cache.Get(url); exists {
		locationResp := LocationResponse{}
		err := json.Unmarshal(cachedData, &locationResp)
		if err != nil {
			return LocationResponse{}, err
		}

		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request")
		return LocationResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println("Error making request")
		return LocationResponse{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response")
		return LocationResponse{}, err
	}

	locationResp := LocationResponse{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		fmt.Println("Error unmarshalling data")
		return LocationResponse{}, err
	}

	return locationResp, nil
}
