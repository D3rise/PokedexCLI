package commands

import (
	"fmt"
	"log"

	"github.com/D3rise/pokedexcli/internal/context"
	"github.com/D3rise/pokedexcli/internal/pokedexapi"
)

const mapNextOffset context.ContextKey = "mapNextOffset"

func mapCommand(c *context.Context) error {
	pokedexapi, ok := c.Get(context.PokedexAPI).(*pokedexapi.PokedexAPI)

	if !ok {
		log.Fatal("pokedex api is not initialized")
	}

	if !c.Has(mapNextOffset) {
		c.Set(mapNextOffset, 0)
	}

	offset := c.Get(mapNextOffset).(int)
	result, err := pokedexapi.GetLocationAreaList(20, offset)
	if err != nil {
		log.Fatal(err)
	}

	c.Set(mapNextOffset, offset+20)
	c.Set(mapPreviousOffset, offset-20)

	for _, location := range result.Results {
		fmt.Println(location.Name)
	}

	return nil
}
