package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/viperfangs/godex/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	Next          string
	Previous      string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	cfg.Next = "https://pokeapi.co/api/v2/location-area/"
	cfg.Previous = "https://pokeapi.co/api/v2/location-area/"

	for {
		fmt.Print("GoDex > ")

		// Wait for the user to type and press Enter
		scanner.Scan()
		userInput := scanner.Text()
		cleanUserInput := cleanInput(userInput)
		commandName := cleanUserInput[0]

		if command, exists := getCommands()[commandName]; exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	commands := make(map[string]cliCommand)

	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the GoDex",
		callback:    commandExit,
	}
	commands["help"] = cliCommand{
		name:        "help",
		description: "Print help information",
		callback:    commandHelp,
	}
	commands["map"] = cliCommand{
		name:        "map",
		description: "Get the next page of locations",
		callback:    commandMap,
	}
	commands["mapb"] = cliCommand{
		name:        "map",
		description: "Get the previous page of locations",
		callback:    commandMapb,
	}

	return commands
}
