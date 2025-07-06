package commands

import (
	"fmt"
	"github.com/D3rise/pokedexcli/internal/context"
	"github.com/D3rise/pokedexcli/internal/pokedex"
	"github.com/D3rise/pokedexcli/pokedexapi"
	"math/rand"
	"time"
)

var catchCommandMeta = cliCommand{
	Name: "catch",
	Args: []string{
		"pokemonName",
	},
	Description: "Try to catch a pokemon",
	Callback:    catchCommand,
}

func catchCommand(c *context.Context, args ...string) error {
	api := c.Get(pokedexapi.PokedexApiContextKey).(*pokedexapi.PokedexAPI)
	dex := c.Get(pokedex.PokedexContextKey).(*pokedex.Pokedex)

	pokemonIdOrName := args[0]

	pokemon, err := api.GetPokemonInfo(pokemonIdOrName)
	if err != nil && err.Error() == "not found" {
		fmt.Println("Pokemon Not Found!")
		return nil
	} else if err != nil {
		return err
	}

	fmt.Println("Throwing a Pokeball at " + pokemon.Name + "...")

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	success := r.Intn(pokemon.BaseExperience)

	if success < pokemon.BaseExperience/2 {
		fmt.Println("Oh no... It slipped away! Try again!")
		return nil
	}

	dex.AddNewPokemon(pokemon.Name)
	fmt.Println("Nice! It got caught! Congratulations, new Pokemon is in your Pokedex now.")
	return nil
}
