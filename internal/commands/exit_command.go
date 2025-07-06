package commands

import (
	"fmt"
	"os"

	"github.com/D3rise/pokedexcli/internal/context"
)

func exitCommand(_ *context.Context, _ ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
