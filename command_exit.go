package main

import (
	"os"
	"fmt"
)

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
