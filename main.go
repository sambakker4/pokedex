package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand{
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help" : {
			name: "help",
			description: "Displays the help message",
			callback: commandHelp,
		},
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
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

		err := getCommands()[command].callback()
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


