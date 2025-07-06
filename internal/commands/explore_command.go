package commands

import (
	"fmt"
	"github.com/D3rise/pokedexcli/internal/context"
	"github.com/D3rise/pokedexcli/pokedexapi"
)

var exploreCommandMeta = cliCommand{
	Name:        "explore",
	Description: "Explore an area for pokemons",
	Args:        []string{"areaName"},
	Callback:    exploreCommand,
}

func exploreCommand(c *context.Context, args ...string) error {
	api := c.Get(pokedexapi.PokedexApiContextKey).(*pokedexapi.PokedexAPI)

	areaName := args[0]
	fmt.Printf("Exploring %s...\n", areaName)

	locationInfo, err := api.GetLocationAreaInfo(areaName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")

	for _, v := range locationInfo.PokemonEncounters {
		fmt.Println(" - " + v.Pokemon.Name)
	}

	return nil
}
