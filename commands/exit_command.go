package commands

import (
	"fmt"
	"os"

	"github.com/D3rise/pokedexcli/context"
)

func exitCommand(_ *context.Context) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
