package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/sambakker4/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	Cache    pokeapi.Cache
	Next     string
	Previous string
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
			description: "Displays the help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "maps 20 locations every time its used",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map back",
			description: "maps back 20 locations every time its used (reverse of map)",
			callback:    commandMapBack,
		},
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := &Config{
		Next:  "https://pokeapi.co/api/v2/location-area/",
		Cache: pokeapi.NewCache(20 * time.Second),
	}
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		cleanedInput := cleanInput(scanner.Text())

		if len(cleanedInput) == 0 {
			continue
		}

		command := cleanedInput[0]
		if _, ok := getCommands()[command]; !ok {
			fmt.Println("Unknown Command")
			continue
		}

		err := getCommands()[command].callback(config)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

func cleanInput(text string) []string {
	slice := strings.Fields(text)
	newSlice := make([]string, 0)

	for i := range slice {
		newSlice = append(newSlice, strings.ToLower(slice[i]))
	}
	return newSlice
}
