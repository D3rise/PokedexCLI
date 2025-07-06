package commands

import (
	"fmt"
	"github.com/D3rise/pokedexcli/pokedexapi"
	"log"

	"github.com/D3rise/pokedexcli/internal/context"
)

const mapPreviousOffset context.ContextKey = "mapPreviousOffset"

func mapbCommand(c *context.Context, _ ...string) error {
	api, ok := c.Get(context.PokedexAPI).(*pokedexapi.PokedexAPI)

	if !ok {
		log.Fatal("pokedex api is not initialized")
	}

	if !c.Has(mapPreviousOffset) {
		c.Set(mapPreviousOffset, 0)
	}

	offset := c.Get(mapPreviousOffset).(int)
	if offset < 0 {
		offset = 0
	}

	result, err := api.GetLocationAreaList(20, offset)
	if err != nil {
		log.Fatal(err)
	}

	c.Set(mapPreviousOffset, offset-20)
	c.Set(mapNextOffset, offset+20)

	for _, location := range result.Results {
		fmt.Println(location.Name)
	}

	return nil
}
