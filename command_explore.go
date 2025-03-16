package main

import (
	"net/http"
	"io"
	"encoding/json"
	"fmt"
)


func commandExplore(config *Config, location string) error {
	url := config.URL + "location-area/" + location

	var exploring Explore
	if val, ok := config.Cache.Get(url); ok {
		err := json.Unmarshal(val, &exploring)
		if err != nil {
			return err
		}

	} else {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		config.Cache.Add(url, data)

		err = json.Unmarshal(data, &exploring)
		if err != nil {
			return err
		}
	}
	fmt.Printf("Exploring " + location + "...\n")
	fmt.Println("Found Pokemon:")

	for _, pokemon := range exploring.PokemonEncounters {
		fmt.Println(" - " + pokemon.Pokemon.Name)
	}

	return nil
}
