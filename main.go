package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
