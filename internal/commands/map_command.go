package commands

import (
	"fmt"
	"github.com/D3rise/pokedexcli/pokedexapi"
	"log"

	"github.com/D3rise/pokedexcli/internal/context"
)

var mapCommandMeta = cliCommand{
	Name:        "map",
	Args:        []string{},
	Description: "Print locations, each usage increases offset by 20",
	Callback:    mapCommand,
}

const mapNextOffset context.ContextKey = "mapNextOffset"

func mapCommand(c *context.Context, _ ...string) error {
	api, ok := c.Get(context.PokedexAPI).(*pokedexapi.PokedexAPI)

	if !ok {
		log.Fatal("pokedex api is not initialized")
	}

	if !c.Has(mapNextOffset) {
		c.Set(mapNextOffset, 0)
	}

	offset := c.Get(mapNextOffset).(int)
	result, err := api.GetLocationAreaList(20, offset)
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
