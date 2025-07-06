package commands

import (
	"fmt"

	"github.com/D3rise/pokedexcli/internal/context"
)

func helpCommand(_ *context.Context, _ ...string) error {
	registry := GetCommandRegistry()

	fmt.Print("Usage: \n\n")
	for _, c := range registry {
		fmt.Printf("%s: %s\n", c.Name, c.Description)
	}

	return nil
}
