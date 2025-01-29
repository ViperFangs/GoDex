package main

import "fmt"

func commandHelp(cfg *Config) error {
	fmt.Printf("Welcome to the GoDex!\nUsage:\n\n")

	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}

	return nil
}
