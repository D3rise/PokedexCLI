package commands

import (
	"fmt"
	"os"

	"github.com/D3rise/pokedexcli/internal/context"
)

var exitCommandMeta = cliCommand{
	Name:        "exit",
	Args:        []string{},
	Description: "Exit the Pokedex",
	Callback:    exitCommand,
}

func exitCommand(_ *context.Context, _ ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
