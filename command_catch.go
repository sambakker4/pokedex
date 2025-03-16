package main

import (
	"net/http"
	"encoding/json"
	"io"
	"math/rand"
	"fmt"
)


func commmandCatch(config *Config, pokemonName string) error {
	url := config.URL + "pokemon/" + pokemonName
	var pokemon Pokemon

	if val, ok := config.Cache.Get(url); ok {
		err := json.Unmarshal(val, &pokemon)
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

		err = json.Unmarshal(data, &pokemon)
		if err != nil {
			return err
		}
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)
	experience := pokemon.BaseExperience

	var caught bool
	chance := rand.Intn(experience) 
	if chance < 1 {
		chance = 1
	} else if chance > 95{
		chance = 95
	}

	if  chance > 40 {
		caught = false
		fmt.Printf("%v escaped!\n", pokemonName)
	} else {
		caught = true
		fmt.Printf("%v was caught!\n", pokemonName)
	}

	if caught {
		Pokedex[pokemonName] = pokemon
	}

	return nil 
}
