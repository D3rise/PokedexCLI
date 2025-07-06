package commands

import (
	"fmt"
	"os"
)

func exitCommand() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
