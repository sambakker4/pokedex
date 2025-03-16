package main

import (
	"errors"
	"fmt"
)

func commandInspect(config *Config, pokemonName string) error {
	if _, ok := Pokedex[pokemonName]; !ok {
		return errors.New("Not a caught Pokemon")
	}
	pokemon := Pokedex[pokemonName]
	fmt.Printf("Name: %v\n", pokemonName)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Printf(" -%v: %v\n", typeInfo.Type.Name, typeInfo.Slot)
	}

	return nil
}
