package commands

import (
	"fmt"
	"log"

	"github.com/D3rise/pokedexcli/api"
	"github.com/D3rise/pokedexcli/context"
)

const mapNextOffset context.ContextKey = "mapNextOffset"

func mapCommand(c *context.Context) error {
	pokedexapi, ok := c.Get(context.PokedexAPI).(*api.PokedexAPI)

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
