package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/D3rise/pokedexcli/api"
	"github.com/D3rise/pokedexcli/commands"
	"github.com/D3rise/pokedexcli/context"
)

const (
	POKEDEX_PROMPT = "Pokedex > "
	POKEDEX_MOTD   = "Welcome to the Pokedex!"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	registry := commands.GetCommandRegistry()

	context := context.NewContext()
	initializeContext(context)

	fmt.Println(POKEDEX_MOTD)
	fmt.Print(POKEDEX_PROMPT)
	for scanner.Scan() {
		text := scanner.Text()
		if c, ok := registry[text]; ok {
			c.Callback(context)
		} else {
			fmt.Println("Unknown command")
		}
		fmt.Print(POKEDEX_PROMPT)
	}
}

func initializeContext(c *context.Context) {
	c.Set(context.PokedexAPI, api.NewPokedexAPI(""))
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
