package commands

import (
	"fmt"
	"strings"

	"github.com/D3rise/pokedexcli/internal/context"
)

var helpCommandMeta = cliCommand{
	Name:        "help",
	Description: "Displays a help message",
	Callback:    helpCommand,
}

func helpCommand(_ *context.Context, _ ...string) error {
	registry := GetCommandRegistry()

	fmt.Print("Usage: \n\n")
	for _, c := range registry {
		var argsInfo string
		if len(c.Args) > 0 {
			argsInfo = fmt.Sprintf(" <%s>", strings.Join(c.Args, "> <"))
		}
		fmt.Printf("%s%s: %s\n", c.Name, argsInfo, c.Description)
	}

	return nil
}
