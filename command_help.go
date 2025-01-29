package main

import "fmt"

func commandHelp() error {
	fmt.Printf("Welcome to the GoDex!\nUsage:\n\n")

	for _, c := range commands {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}

	return nil
}
