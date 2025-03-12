package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex")
	fmt.Printf("Usage:\n\n")
	
	for _, command := range getCommands(){
		fmt.Println(command.name, ":", command.description)
	}
	fmt.Println()
	return nil
}
