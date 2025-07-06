package main

import "github.com/D3rise/pokedexcli/internal/commands"

func main() {
	commands.InitializeRegistry()
	repl()
}
