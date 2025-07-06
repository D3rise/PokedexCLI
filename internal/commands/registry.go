package commands

import "github.com/D3rise/pokedexcli/internal/context"

type cliCommandCallback func(ctx *context.Context, args ...string) error

type cliCommand struct {
	Name        string
	Description string
	Args        []string
	Callback    cliCommandCallback
}

var commandRegistry map[string]cliCommand

func GetCommandRegistry() map[string]cliCommand {
	return commandRegistry
}

func InitializeRegistry() {
	commandRegistry = map[string]cliCommand{
		"exit":    exitCommandMeta,
		"help":    helpCommandMeta,
		"map":     mapCommandMeta,
		"mapb":    mapbCommandMeta,
		"explore": exploreCommandMeta,
	}
}
