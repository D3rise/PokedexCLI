package commands

type cliCommandCallback func() error

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
	}
}
