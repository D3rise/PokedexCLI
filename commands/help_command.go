package commands

import "fmt"

func helpCommand() error {
	registry := GetCommandRegistry()

	fmt.Print("Usage: \n\n")
	for _, c := range registry {
		fmt.Printf("%s: %s\n", c.Name, c.Description)
	}

	return nil
}
