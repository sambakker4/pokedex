package main

import (
	"fmt"
	"errors"
)

func commandPokedex(config *Config, s string) error {
	if len(Pokedex) == 0 {
		return errors.New("Your Pokedex is empty")
	}
	fmt.Println("Your Pokedex:")

	for pokemonName, _ := range Pokedex {
		fmt.Println(" -", pokemonName)
	}
	return nil
}
