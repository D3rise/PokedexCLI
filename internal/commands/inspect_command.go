package commands

import (
	"fmt"
	"github.com/D3rise/pokedexcli/internal/context"
	"github.com/D3rise/pokedexcli/internal/pokedex"
	"github.com/D3rise/pokedexcli/pokedexapi"
)

var inspectCommandMeta = cliCommand{
	Name: "inspect",
	Args: []string{
		"pokemonName",
	},
	Description: "Inspect a pokemon from your Pokedex",
	Callback:    inspectCommand,
}

func inspectCommand(c *context.Context, args ...string) error {
	pokemonName := args[0]

	api := c.Get(pokedexapi.PokedexApiContextKey).(*pokedexapi.PokedexAPI)
	dex := c.Get(pokedex.PokedexContextKey).(*pokedex.Pokedex)

	if _, isInDex := dex.GetCaughtPokemons()[pokemonName]; !isInDex {
		fmt.Println("This pokemon is not in your pokedex yet!\nCatch it using `catch` command, then try again.")
		return nil
	}

	pokemonInfo, err := api.GetPokemonInfo(pokemonName)
	if err != nil {
		return err
	}

	fmt.Println("Name:", pokemonInfo.Name)
	fmt.Println("Height:", pokemonInfo.Height)
	fmt.Println("Weight:", pokemonInfo.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemonInfo.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("\nTypes:")
	for _, typ := range pokemonInfo.Types {
		fmt.Printf(" - %s\n", typ.Type.Name)
	}

	fmt.Println()

	return nil
}
