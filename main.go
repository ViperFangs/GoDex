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

var commands map[string]cliCommand

func init() {
	commands = make(map[string]cliCommand)
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
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("GoDex > ")

		// Wait for the user to type and press Enter
		scanner.Scan()
		userInput := scanner.Text()
		cleanUserInput := cleanInput(userInput)

		if command, exists := commands[cleanUserInput[0]]; exists {
			command.callback()
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

func commandExit() error {
	fmt.Println("Closing the GoDex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the GoDex!\nUsage:\n\n")

	for _, c := range commands {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}

	return nil
}
