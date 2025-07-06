package commands

import (
	"fmt"
	"github.com/D3rise/pokedexcli/internal/context"
	"github.com/D3rise/pokedexcli/internal/pokedex"
)

var pokedexCommandMeta = cliCommand{
	Name:        "pokedex",
	Args:        []string{},
	Description: "Print all caught pokemons in your pokedex",
	Callback:    pokedexCommand,
}

func pokedexCommand(c *context.Context, _ ...string) error {
	dex := c.Get(pokedex.PokedexContextKey).(*pokedex.Pokedex)
	pokemons := dex.GetCaughtPokemons()

	if len(pokemons) == 0 {
		fmt.Println("No pokemons in your pokedex yet!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, p := range pokemons {
		fmt.Println(" - ", p.Name)
	}

	return nil
}
