package commands

import "github.com/D3rise/pokedexcli/context"

type cliCommandCallback func(*context.Context) error

type cliCommand struct {
	Name        string
	Description string
	Callback    cliCommandCallback
}

func GetCommandRegistry() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    exitCommand,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    helpCommand,
		},
		"map": {
			Name:        "map",
			Description: "Print locations, each usage increases offset by 20",
			Callback:    mapCommand,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Print locations, each usage decreases offset by 20",
			Callback:    mapbCommand,
		},
	}
}
