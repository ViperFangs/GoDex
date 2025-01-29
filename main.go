package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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

	return commands
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("GoDex > ")

		// Wait for the user to type and press Enter
		scanner.Scan()
		userInput := scanner.Text()
		cleanUserInput := cleanInput(userInput)
		commandName := cleanUserInput[0]

		if command, exists := getCommands()[commandName]; exists {
			err := command.callback()
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
