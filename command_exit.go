package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Closing the GoDex... Goodbye!")
	os.Exit(0)
	return nil
}
